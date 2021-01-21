package controller

import (
	_ "encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "strconv"
	"strings"

	extra "hydrawebapi/extraapi/model"

	_ "charlie/i3.0.0/cls"

	"github.com/labstack/echo"
)

// 자체 인증키 요청
func getFreeBillSelfAuthKey(id string) string {
	//url := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/GETAUTHKEYBYPW/%s/%s", id, password)
	url := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/GETAUTHKEY/%s", id)
	lprintf(4, "[INFO] url : [%s]\n", url)

	respGet, err := http.Get(url)
	if err != nil {
		lprintf(1, "[ERR ] http.Get : [%s]\n", err)
		return ""
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		return ""
	}

	respString := string(respBody)
	lprintf(4, "[INFO] auth-key respString : [%s]\n", respString)

	infoArray := strings.Split(respString, "|")
	if infoArray[0] == "N" {
		infoArray[1] = ConvertEucKrDecoding([]byte(infoArray[1]))
		//lprintf(1, "[ERR ] %s\n", infoArray)
	}

	return infoArray[1]
}

// Freebill 인증키 발급
func GetFreebillAuth(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetFreebillAuthParam
	var respJson extra.RespGetFreebillAuth

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	url := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/GETAUTHKEYBYPW/%s/%s", param.Id, param.Password)
	lprintf(4, "[INFO] url : [%s]\n", url)

	respGet, err := http.Get(url)
	if err != nil {
		lprintf(1, "[ERR ] Get : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill auth key respString : [%s]\n", respString)

	infoArray := strings.Split(respString, "|")
	if infoArray[0] == "N" {
		infoArray[1] = ConvertEucKrDecoding([]byte(infoArray[1]))
		//lprintf(1, "[ERR ] %s\n", infoArray)
	}
	respJson.Status = infoArray[0]
	respJson.Message = infoArray[1]

	// make response
	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 전자 세금 계산서 등록
func SetFreeBillRegister(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetFreeBillRegisterParam
	var respJson extra.RespSetFreeBillRegister
	var item string

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/LOGIN")
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	// IS_RECEIVE 정의
	// 1 - 매입 : 공급자(판매자), 2 - 매출 : 공급 받는자(구매자)
	// 회사는 판매자 이므로 is_receive 를 2로 셋팅
	// param.IsReceive = "2"

	// open 값 체크
	if param.Open == "" {
		param.Open = "n" // Y 는 화면 바로가기 이므로, 값이 없을 경우 default로 N으로 요청한다. 응답데이터 문서번호
	}

	// items 가공
	for _, v := range param.Item {
		// total로 세금계산서를 등록 할때는 규격 및 수량을 고정
		if v.Standard == "" {
			v.Standard = "-"
		}
		if v.UnitPrice == 0 {
			v.UnitPrice = 1
		}

		productName := ConvertEucKrEncoding([]byte(strings.TrimSpace(v.ProductName)))
		standard := ConvertEucKrEncoding([]byte(strings.TrimSpace(v.Standard)))

		item += fmt.Sprintf("%s||%g||%g||%s||%d||%g||", productName, v.Price, v.PriceVat, standard, v.UnitPrice, v.PriceVatSum)
	}
	if item != "" {
		item = strings.TrimRight(item, "||")
		lprintf(4, "[INFO] item : [%s]\n", item)
	}

	// 요청 시 한글은 euc-kr로 encoding 해야 한다.
	taxType := ConvertEucKrEncoding([]byte(param.TaxType))
	typeText := ConvertEucKrEncoding([]byte(param.TypeText))
	company := ConvertEucKrEncoding([]byte(param.Company))
	president := ConvertEucKrEncoding([]byte(param.President))
	addr := ConvertEucKrEncoding([]byte(param.Addr))
	btype := ConvertEucKrEncoding([]byte(param.Btype))
	bclass := ConvertEucKrEncoding([]byte(param.Bclass))
	name := ConvertEucKrEncoding([]byte(param.Name))
	message := ConvertEucKrEncoding([]byte(param.Message))
	description := ConvertEucKrEncoding([]byte(param.Description))

	data := url.Values{}
	// data.Set("a", param.A)
	data.Set("a", a)
	data.Set("tpf", param.Tpf)
	data.Set("open", param.Open)
	data.Set("is_receive", param.IsReceive)
	data.Set("date", param.Date)
	data.Set("tax_type", taxType)
	data.Set("type_text", typeText)
	data.Set("item", item)
	data.Set("number", param.Number)
	//data.Set("fnumber", param.Fnumber)
	data.Set("company", company)
	data.Set("president", president)
	data.Set("addr", addr)
	data.Set("btype", btype)
	data.Set("bclass", bclass)
	data.Set("name", name)
	data.Set("hp", param.Hp)
	data.Set("fax", param.Fax)
	data.Set("email", param.Email)
	data.Set("message", message)
	data.Set("volume", param.Volume)
	data.Set("issue", param.Issue)
	data.Set("sequence", param.Sequence)
	data.Set("description", description)
	data.Set("payment_type", param.PaymentType)
	data.Set("uniq_code", param.UniqCode)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill register respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 전자 세금 계산서 확인
func GetFreeBillSearch(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetFreeBillSearchParam
	var respJson extra.RespGetFreeBillSearch

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if param.Uniq == "" {
		lprintf(4, "[INFO] uniq Number is empty, so set [n]")
		param.Uniq = "n" // default(프리빌 문서 고유 번호 방식)
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/SEARCH/%s/%s", a, param.Dnumber)
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	data := url.Values{}
	data.Set("uniq", param.Uniq)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill search respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 전자 세금 계산서 상태
func GetFreeBillSearchStatus(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetFreeBillSearchStatusParam
	var respJson extra.RespGetFreeBillSearchStatus

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if param.Uniq == "" {
		lprintf(4, "[INFO] uniq Number is empty, so set [n]")
		param.Uniq = "n"
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/SEARCHSTATUS/%s/%s", a, param.Dnumber)
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	data := url.Values{}
	data.Set("uniq", param.Uniq)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill search status respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		// 응답 : "SUCCESS:0|N|���۴���"
		respSucc := strings.Split(respString, "|")
		respMsg := ConvertEucKrDecoding([]byte(respSucc[2]))
		respString = fmt.Sprintf("%s|%s|%s", respSucc[0], respSucc[1], respMsg)
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)

}

// 전자 세금 계산서 삭제
func DelFreeBill(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.DelFreeBillParam
	var respJson extra.RespDelFreeBill

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if param.Uniq == "" {
		lprintf(4, "[INFO] uniq Number is empty, so set [n]")
		param.Uniq = "n"
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/DELETE/%s/%s", a, param.Dnumber)
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	data := url.Values{}
	data.Set("uniq", param.Uniq)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill delete respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 문서 보기(VIEW)
func GetFreeBillView(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetFreeBillViewParam
	var respJson extra.RespGetFreeBillView

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if param.Uniq == "" {
		lprintf(4, "[INFO] uniq Number is empty, so set [n]")
		param.Uniq = "n" // default(프리빌 문서 고유 번호 방식)
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/VIEW/%s/%s", a, param.Dnumber)
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	data := url.Values{}
	data.Set("uniq", param.Uniq)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill view respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 문서 취소(CANCEL)
func DelFreeBillCancel(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.DelFreeBillCancelParam
	var respJson extra.RespDelFreeBillCancel

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if param.Uniq == "" {
		lprintf(4, "[INFO] uniq Number is empty, so set [n]")
		param.Uniq = "n"
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/CANCEL/%s/%s", a, param.Dnumber)
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	message := ConvertEucKrEncoding([]byte(param.Message))

	data := url.Values{}
	data.Set("uniq", param.Uniq)
	data.Set("msg", message)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill cancel respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 기등록 문서 전자발행 요청
func GetFreeBillPreviousRegister(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetFreeBillPreviousRegisterParam
	var respJson extra.RespGetFreeBillPreviousRegister
	var listCode string

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	for _, v := range param.Dnumber {
		listCode += fmt.Sprintf("%s|", v)
	}
	if listCode != "" {
		listCode = strings.TrimRight(listCode, "|")
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/PUBLISHNOW")
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	data := url.Values{}
	data.Set("a", a)
	data.Set("list_code", listCode)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill previous register respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 세금계산서 등록 + 전자발행 요청
func SetFreeBillPublishNow(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetFreeBillPublishNowParam
	var respJson extra.RespSetFreeBillPublishNow
	var item string

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/PUBLISHNOW")
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	// IS_RECEIVE 정의
	// 1 - 매입 : 공급자(판매자), 2 - 매출 : 공급 받는자(구매자)
	// 회사는 판매자 이므로 is_receive 를 2로 셋팅
	// param.IsReceive = "2"

	// open 값 체크
	if param.Open == "" {
		param.Open = "n" // Y 는 화면 바로가기 이므로, 값이 없을 경우 default로 N으로 요청한다. 응답데이터 문서번호
	}

	// items 가공
	for _, v := range param.Item {
		// total로 세금계산서를 등록 할때는 규격 및 수량을 고정
		if v.Standard == "" {
			v.Standard = "-"
		}
		if v.UnitPrice == 0 {
			v.UnitPrice = 1
		}

		productName := ConvertEucKrEncoding([]byte(strings.TrimSpace(v.ProductName)))
		standard := ConvertEucKrEncoding([]byte(strings.TrimSpace(v.Standard)))

		// item += fmt.Sprintf("%s||%g||%g||%s||%d||%g||", v.ProductName, v.Price, v.PriceVat, v.Standard, v.UnitPrice, v.PriceVatSum)
		item += fmt.Sprintf("%s||%g||%g||%s||%d||%g||", productName, v.Price, v.PriceVat, standard, v.UnitPrice, v.PriceVatSum)
	}
	if item != "" {
		item = strings.TrimRight(item, "||")
		lprintf(4, "[INFO] item : [%s]\n", item)
	}

	// 요청 시 한글은 euc-kr로 encoding 해야 한다.
	taxType := ConvertEucKrEncoding([]byte("과세"))
	typeText := ConvertEucKrEncoding([]byte("영수"))
	company := ConvertEucKrEncoding([]byte(param.Company))
	president := ConvertEucKrEncoding([]byte(param.President))
	addr := ConvertEucKrEncoding([]byte(param.Addr))
	btype := ConvertEucKrEncoding([]byte(param.Btype))
	bclass := ConvertEucKrEncoding([]byte(param.Bclass))
	name := ConvertEucKrEncoding([]byte(param.Name))
	message := ConvertEucKrEncoding([]byte(param.Message))
	description := ConvertEucKrEncoding([]byte(param.Description))

	data := url.Values{}
	data.Set("a", a)
	data.Set("sign_type", param.SignType)
	data.Set("tpf", "2")
	data.Set("open", param.Open)
	data.Set("is_receive", "2")
	data.Set("date", param.Date)
	data.Set("tax_type", taxType)
	data.Set("type_text", typeText)
	data.Set("item", item)
	data.Set("number", param.Number)
	//data.Set("fnumber", param.Fnumber)
	data.Set("company", company)
	data.Set("president", president)
	data.Set("addr", addr)
	data.Set("btype", btype)
	data.Set("bclass", bclass)
	data.Set("name", name)
	data.Set("hp", param.Hp)
	data.Set("fax", param.Fax)
	data.Set("email", param.Email)
	data.Set("message", message)
	data.Set("volume", param.Volume)
	data.Set("issue", param.Issue)
	data.Set("sequence", param.Sequence)
	data.Set("description", description)
	data.Set("payment_type", param.PaymentType)
	data.Set("uniq_code", param.UniqCode)

	// data := url.Values{}
	// data.Set("a", a)
	// data.Set("sign_type", param.SignType)
	// data.Set("tpf", param.Tpf)
	// data.Set("open", param.Open)
	// data.Set("is_receive", param.IsReceive)
	// data.Set("date", param.Date)
	// data.Set("tax_type", param.TaxType)
	// data.Set("type_text", param.TypeText)
	// data.Set("item", item)
	// data.Set("number", param.Number)
	// //data.Set("fnumber", param.Fnumber)
	// data.Set("company", param.Company)
	// data.Set("president", param.President)
	// data.Set("addr", param.Addr)
	// data.Set("btype", param.Btype)
	// data.Set("bclass", param.Bclass)
	// data.Set("name", param.Name)
	// data.Set("hp", param.Hp)
	// data.Set("fax", param.Fax)
	// data.Set("email", param.Email)
	// data.Set("message", param.Message)
	// data.Set("volume", param.Volume)
	// data.Set("issue", param.Issue)
	// data.Set("sequence", param.Sequence)
	// data.Set("description", param.Description)
	// data.Set("payment_type", param.PaymentType)
	// data.Set("uniq_code", param.UniqCode)

	lprintf(4, "[INFO] data : [%s]\n", data)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill register & publish respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 전자 세금 계산서 국세청 수동 신고
func SetFreeBillSend(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetFreeBillSendParam
	var respJson extra.RespSetFreeBillSend
	var listCode string

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	for _, v := range param.Dnumber {
		listCode += fmt.Sprintf("%s|", v)
	}
	if listCode != "" {
		listCode = strings.TrimRight(listCode, "|")
	}

	// 인증키 요청
	a := getFreeBillSelfAuthKey(param.Id)
	if a == "" {
		lprintf(1, "[FAIL] get not auth key\n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pUrl := fmt.Sprintf("http://playauto.freebill.co.kr/web/api/SEND/%s/%s", a, listCode)
	lprintf(4, "[INFO] url : [%s]\n", pUrl)

	data := url.Values{}
	data.Set("uniq", param.Uniq)

	req, err := http.NewRequest("POST", pUrl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "EUC-KR")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respString := string(respBody)
	lprintf(4, "[INFO] freebill send respString : [%s]\n", respString)
	// error check
	if strings.Contains(respString, "ERROR:") {
		respJson.Data = ConvertEucKrDecoding(respBody)
		// 응답 처리는 일단 작업을 완료 후 다시 정의하자!!
	} else {
		respJson.Data = respString
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}
