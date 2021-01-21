package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	extra "hydrawebapi/extraapi/model"

	"charlie/i3.0.0/cls"

	"github.com/labstack/echo"
)

var lprintf func(int, string, ...interface{}) = cls.Lprintf

func testSleep(data string) {
	idx := 0
	for {
		lprintf(4, "[INFO] %s\n", data)
		idx++
		if idx == 10 {
			break
		}
		time.Sleep(time.Second * 1)
	}
	lprintf(4, "[INFO] end\n")
}

// 등록한 외부 API 검색
func GetExtraApiAll(c echo.Context) error {

	respData := extra.RespGetExtraApiAll{}
	// var param map[string][]string
	// var extraDomainList string

	// // BODY DATA
	// resp, err := ioutil.ReadAll(c.Request().Body)
	// if err != nil {
	// 	lprintf(1, "[ERR ] request body : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// lprintf(4, "[INFO] resp : [%s]\n", resp)

	// if err := json.Unmarshal(resp, &param); err != nil {
	// 	lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	// for _, v := range param["extraDomainArray"] {
	//     extraDomainList += fmt.Sprintf("%s,", v)
	// }
	// if extraDomainList != "" {
	//     extraDomainList = strings.TrimRight(extraDomainList, ",")
	// }

	// procName := fmt.Sprintf("CALL proc_bo_domain_get_extradomainlist_test('%s');", extraDomainList)
	// lprintf(4, "[INFO] procName : [%s]\n", procName)

	// rows, err := cls.QueryDB(procName)
	// if err != nil {
	// 	lprintf(1, "[ERR ] procName : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// defer rows.Close()

	// // 응답코드
	// if result := cls.GetRespCode(rows, procName); result == 99 {
	// 	respData.Code = extra.FAIL
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	// if rows.NextResultSet() {
	//     for rows.Next() {
	//         ed := extra.GetExtraApi{}
	//         if err := rows.Scan(&ed.DomainName, &ed.RegisterYn); err != nil {
	// 			lprintf(1, "[ERR ] %s second return scan error : %s\n", procName, err)
	// 			respData.Code = extra.C500
	//
	// 			respData.ServiceName = extra.TYPE
	// 			return c.JSON(http.StatusOK, respData)
	//         }

	//         respData.Data = append(respData.Data, ed)
	//     }
	// }

	// // make response
	// respData.Code = extra.SUCCESS
	//
	// respData.ServiceName = extra.TYPE

	// return c.JSON(http.StatusOK, respData)

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)

	// c.JSON(http.StatusOK, "success1233")

	// go testSleep("123123123")

	// idx := 0
	// for {
	// 	lprintf(4, "[INFO] 11111111\n")
	// 	if idx == 10 {
	// 		break
	// 	}
	// 	time.Sleep(time.Second * 1)
	// 	idx++
	// }

	// return c.JSON(http.StatusOK, "success1233")
}

// 도메인 등록
func SetDomainRegisterWhoisOne(c echo.Context) error {
	respData := extra.DomainEntityApiCallResponse{}
	var param extra.SetDomainRegisterWhoisOneParam
	var respJson extra.RespSetDomainRegisterWhois

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

	formData := url.Values{
		"domain-name":        {param.DomainName},
		"customer-id":        {extra.CustomerId},
		"reg-contact-id":     {extra.RegContactId},
		"admin-contact-id":   {extra.AdminContactId},
		"tech-contact-id":    {extra.TechContactId},
		"billing-contact-id": {extra.BillingContactId},
		"invoice-option":     {extra.InvoiceOption},
		"protect-privacy":    {extra.ProtectPrivate},
		"years":              {param.YearNum},
		"ns":                 param.NsArray,
	}
	surl := fmt.Sprintf("%s/api/domains/register.xml?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] Domain_name : [%s], response: [%v]\n", param.DomainName, string(respBody))
	respJson.Data = string(respBody)

	if strings.Contains(respJson.Data, "<string>entityid</string>") {
		respData.Code = extra.SUCCESS
		respData.Message.Cn = "Success"
		respData.Message.Kr = "Success"
		respData.Message.En = "Success"
		entityIDArry1 := strings.Split(respJson.Data, "<string>entityid</string>") //두번째에 EntityID 들어있음
		lprintf(4, "[INFO] entityIdArry 1th : [%v]\n", entityIDArry1)

		entityIDArry2 := strings.Split(entityIDArry1[1], "<string>") //개행 <string>
		lprintf(4, "[INFO] entityIdArry 2th : [%v]\n", entityIDArry2)
		entityIDArry3 := strings.Split(entityIDArry2[1], "</string>") //entityID</string>
		lprintf(4, "[INFO] entityIdArry 3th : [%v]\n", entityIDArry3)

		//	entityId := strings.Split(entityIdArry[1], "<string>")
		respData.Data.EntityId = entityIDArry3[0]
		lprintf(4, "[INFO] EntityId : [%s]\n", respData.Data.EntityId)

	} else if strings.Contains(respJson.Data, "<string>error</string>") {
		respData.Code = extra.C500
		respData.Message.Cn = "Whois Error"
		respData.Message.Kr = "Whois Error"
		respData.Message.En = "Whois Error"
	} else {
		respData.Code = extra.C500
		respData.Message.Cn = "Undefined Message" //후이즈 응답을 읽을 수 없음
		respData.Message.Kr = "Undefined Message" //후이즈 응답을 읽을 수 없음
		respData.Message.En = "Undefined Message" //후이즈 응답을 읽을 수 없음
	}

	// 응답 데이터
	// if err := json.Unmarshal(respBody, &respJson); err != nil {
	// 	lprintf(1, "[ERR ] request body : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	// lprintf(4, "[INFO] Status : [%s], [%v]\n", respJson.Status, respJson)
	// if respJson.Status == "ERROR" {
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	respData.ServiceName = extra.TYPE
	lprintf(4, "[INFO] respData : [%v]\n", respData)
	return c.JSON(http.StatusOK, respData)
}

// 도메인 사용 가능 여부
func GetDomainAvailableWhoisAll(c echo.Context) error {
	respData := extra.RespGetDomainAvailableWhoisAll{}
	var param extra.GetDomainAvailableWhoisAllParam
	//var respJson extra.ApiCallResponseInterface
	var domainName, tlds string

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

	for _, v := range param.DomainName {
		domainName += fmt.Sprintf("&domain-name=%s", v)
	}
	for _, v := range param.Tlds {
		tlds += fmt.Sprintf("&tlds=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/available.json?api-key=%s&auth-userid=%s%s%s", extra.WhoisCheckUrl, extra.ApiKey, extra.AuthUserId, domainName, tlds)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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

	var aa map[string]interface{}
	lprintf(4, "[INFO] aa : [%s]\n", aa)
	lprintf(4, "[INFO] respBody : [%s]\n", respBody)
	if err := json.Unmarshal(respBody, &aa); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)

	}

	lprintf(4, "[INFO] aa : %v [%s]\n", aa)

	for i, _ := range param.DomainName {
		for j, _ := range param.Tlds {
			dl := extra.GetDomainAvailableWhoisAll{}
			domain := fmt.Sprintf("%s.%s", param.DomainName[i], param.Tlds[j])
			lprintf(4, "[INFO] domain : [%s]\n", domain)

			apiRespMap := aa[domain].(interface{})
			dimainMap, ok := apiRespMap.(map[string]interface{})
			lprintf(4, "[INFO] ok: [%v], dimainMap : [%s]\n", ok, dimainMap)

			dl.DomainName = domain
			dl.Status = dimainMap["status"].(string)

			respData.Data.DomainList = append(respData.Data.DomainList, dl)
		}
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
	/*


		lprintf(4, "[INFO] respBody : [%s]\n", respBody)
		if err := json.Unmarshal(respBody, &respJson.Data); err != nil {
			lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		}

		// domain := fmt.Sprintf("%s.%s", param.DomainName[0], param.Tlds[0])
		// lprintf(4, "[INFO] domain : [%v]\n", domain)

		apiResp, ok := respJson.Data.(map[string]interface{})
		lprintf(4, "[INFO] ok: [%v], apiResp : [%s]\n", ok, apiResp)

		//make key
		// if len(param.DomainName) == len(param.Tlds) {
		// 	for i := range param.DomainName {
		// 		dl := extra.GetDomainAvailableWhoisAll{}
		// 		domain := fmt.Sprintf("%s.%s", param.DomainName[i], param.Tlds[i])
		// 		lprintf(4, "[INFO] domain : [%s]\n", domain)

		// 		apiRespMap := apiResp[domain].(interface{})
		// 		dimainMap, ok := apiRespMap.(map[string]interface{})
		// 		lprintf(4, "[INFO] ok: [%v], apiResp : [%s]\n", ok, apiResp)

		// 		dl.DomainName = domain
		// 		dl.Status = dimainMap["status"].(string)

		// 		respData.Data.DomainList = append(respData.Data.DomainList, dl)
		// 	}
		// } else {
		// 	lprintf(1, "[FAIL] DomainName Length [%d], Tlds [%d] are not same length\n", len(param.DomainName), len(param.Tlds))
		// 	respData.Code = extra.C500
		//
		// 	respData.ServiceName = extra.TYPE
		// 	return c.JSON(http.StatusOK, respData)
		// }

		for i, _ := range param.DomainName {
			for j, _ := range param.Tlds {
				dl := extra.GetDomainAvailableWhoisAll{}
				domain := fmt.Sprintf("%s.%s", param.DomainName[i], param.Tlds[j])
				lprintf(4, "[INFO] domain : [%s]\n", domain)

				apiRespMap := apiResp[domain].(interface{})
				dimainMap, ok := apiRespMap.(map[string]interface{})
				lprintf(4, "[INFO] ok: [%v], apiResp : [%s]\n", ok, apiResp)

				dl.DomainName = domain
				dl.Status = dimainMap["status"].(string)

				respData.Data.DomainList = append(respData.Data.DomainList, dl)
			}
		}

		respData.Code = extra.SUCCESS

		respData.ServiceName = extra.TYPE

		return c.JSON(http.StatusOK, respData)*/
}

// 후이즈에 키워드에 대한 추천 도메인을 조회
func GetDomainRecommendWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespGetDomainRecommendWhoisAll
	var keyword string

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

	keyword = param["keyword"]
	surl := fmt.Sprintf("%s/api/domains/v5/suggest-names.json?api-key=%s&auth-userid=%s&keyword=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, keyword)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	respJson.Data = string(respBody)
	lprintf(4, "[INFO] respBody : [%s]\n", respJson.Data)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 도메인 등록 순서 세부 정보
func GetDomainDetailWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetDomainDetailWhoisAllParam
	var respJson extra.RespGetDomainDetailWhoisAll
	var options string

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

	for _, v := range param.Options {
		options += fmt.Sprintf("&options=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/details-by-name.json?api-key=%s&auth-userid=%s&domain-name=%s%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.DomainName, options)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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

	// if err := json.Unmarshal(respBody, &respJson); err != nil {
	// 	lprintf(1, "[ERR ] request body : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))

	// if respJson.Status != "SUCCESS" {
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	// 요청시 Options 값에 따라서 status를 판단하는 hask key 명이 다르므로 화면에서 사용시 협의를 통해 응답 데이터를 가공해야 한다.
	//
	// options="Orderdetails" 인 경우 status key 명이 "orderstatus"
	// options="DomainStatus" 인 경우 status key 명이 "domainstatus"
	respJson.Data = string(respBody)
	lprintf(4, "[INFO] respBody : [%s]\n", respJson.Data)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 도메인 등록 순서 세부 정보 (스케줄러가 사용하는 API)
func GetDomainDetailWhoisAllS(c echo.Context) error {
	respData := extra.TransferStatusResponse{}
	var param extra.GetDomainDetailWhoisAllParam
	var respJson extra.RespGetDomainDetailWhoisAll
	var options string

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

	for _, v := range param.Options {
		options += fmt.Sprintf("&options=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/details-by-name.json?api-key=%s&auth-userid=%s&domain-name=%s%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.DomainName, options)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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

	//respJson.Data = string(respBody)
	// data parsing을 위해 whois 응답값을 전부 소문자로 변환 작업함
	respJson.Data = strings.ToLower(string(respBody))
	lprintf(4, "[INFO] respBody : [%s]\n", respJson.Data)

	//"currentstatus":"Active"

	if strings.Contains(respJson.Data, "\"currentstatus\":\"active\"") {
		respData.Data = "active"
	} else if strings.Contains(respJson.Data, "\"actionstatus\":\"adminonsapprove\"") {
		respData.Data = "error"
	} else if strings.Contains(respJson.Data, "\"actionstatus\":\"transferrequestsent\"") {
		respData.Data = "pending"
	} else if strings.Contains(respJson.Data, "\"actionstatus\":\"validateauthinfo\"") {
		respData.Data = "stalled"
	} else if strings.Contains(respJson.Data, "\"message\":\"website doesn’t\"") {
		respData.Data = "fail"
	} else {
		lprintf(4, "[INFO] WhoIs API Response, unkown data \n")
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 Search
func GetDomainSearchWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetDomainSearchWhoisAllParam
	var respJson extra.RespGetDomainSearchWhoisAll

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

	surl := fmt.Sprintf("%s/api/domains/search.json?api-key=%s&auth-userid=%s&no-of-records=%d&page-no=%d",
		extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.NoOfRecords, param.PageNo)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	respJson.Data = string(respBody)
	lprintf(4, "[INFO] respBody : [%s]\n", respJson.Data)

	// lprintf(4, "[INFO] respBody : [%s]\n", respBody)
	// if err := json.Unmarshal(respBody, &respJson); err != nil {
	// 	lprintf(1, "[ERR ] request body : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))

	// lprintf(4, "[INFO] Recsonpage     : [%s]", respJson.Recsonpage)
	// lprintf(4, "[INFO] RecsonpageData : [%v]", respJson.RecsonpageData)
	// lprintf(4, "[INFO] Recsindb       : [%s]", respJson.Recsindb)
	// lprintf(4, "[INFO] RecsindbData   : [%v]", respJson.RecsindbData)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 이용 가능 여부 확인
func GetDomainAvailableSunriseWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetDomainAvailableSunriseWhoisOneParam

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

	formData := url.Values{
		"domainname": {param.DomainName},
		"tld":        {param.Tld},
		"smd":        {param.Smd},
	}

	surl := fmt.Sprintf("%s/api/domains/available-sunrise.json?api-key=%s&auth-userid=%s", extra.WhoisCheckUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// surl := fmt.Sprintf("%s/api/domains/available-sunrise.json?api-key=%s&auth-userid=%s&domainname=%s&tld=%s&smd=%s",
	//   extra.WhoisCheckUrl, extra.ApiKey, extra.AuthUserId, param.DomainName, param.Tld, param.Smd)
	// lprintf(4, "[INFO] surl : [%s]\n", surl)

	// respGet, err := http.Get(surl)
	// if err != nil {
	// 	lprintf(1, "[ERR ] Get : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// defer respGet.Body.Close()

	// respBody, err := ioutil.ReadAll(respGet.Body)
	// if err != nil {
	// 	lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 네임서버 수정
func ModiDomainModifyNsWhoisOne(c echo.Context) error {
	respData := extra.RespModiDomainModifyNsWhoisOne{}
	var param extra.ModiDomainModifyNsWhoisOneParam
	//var respJson extra.ApiCallResponseInterface

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
		"ns":       param.Ns,
	}

	surl := fmt.Sprintf("%s/api/domains/modify-ns.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respBodyData := string(respBody)

	lprintf(4, "[INFO] respBody : [%s]\n", respBodyData)
	//respJson.Data = string(respBody)

	if respData.Data.Status == "Success" {
		respData.Code = extra.SUCCESS

		respData.Data.Status = "Success"
		respData.Data.Msg = "Success"
	} else {
		var respMessage extra.MessageOne
		if err := json.Unmarshal(respBody, &respMessage); err != nil {
			lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
			respData.Code = extra.C500

			respData.ServiceName = extra.TYPE
			return c.JSON(http.StatusOK, respData)
		}
		respData.Code = extra.FAIL

		respData.ServiceName = extra.TYPE
		respData.Data.Status = "ERROR"
		respData.Data.Msg = respMessage.Message

		if strings.Contains(respMessage.Message, "Same value for new and old NameServers.") {
			respData.Code = extra.SUCCESS
			respData.Data.Status = "Success"
			respData.Data.Msg = "Success"
		}

	}

	respData.ServiceName = extra.TYPE
	//respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 orderid 가져오기
func GetDomainOrderIdWhoisOne(c echo.Context) error {
	respData := extra.RespGetDomainOrderIdWhoisOne{}
	var param extra.DomainNameParam
	var respJson extra.ApiCallResponseInterface
	//var domainName string

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

	surl := fmt.Sprintf("%s/api/domains/orderid.json?api-key=%s&auth-userid=%s&domain-name=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.DomainName)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)

	respGet, err := http.Get(surl)
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

	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))

	if err := json.Unmarshal(respBody, &respJson.Data); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	whoisResp, ok := respJson.Data.(map[string]interface{})
	lprintf(4, "[INFO] whoisResp: [%s], ok: [%v]\n", whoisResp, ok)
	var whoisMessage string
	if ok == true {
		respData.Data.Status = whoisResp["status"].(string)
		whoisMessage = whoisResp["message"].(string)

	} else {
		if intRespBody, err := strconv.Atoi(string(respBody)); err != nil {
			lprintf(1, "[ERR ] strconv error : [%s]\n", err)
			respData.Code = extra.C500
			respData.ServiceName = extra.TYPE
			return c.JSON(http.StatusOK, respData)
		} else {
			respData.Data.Status = "success"
			respData.Data.OrderId = intRespBody
		}
	}

	if strings.Contains(whoisMessage, "Website doesn't exist for") {
		respData.Code = extra.FAIL
		respData.ServiceName = extra.TYPE

		return c.JSON(http.StatusOK, respData)

	}

	// if respData.Data.Status == "ERROR" {
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	//respJson.Data = string(respBody)
	//respJson.Data = string(respBody)

	// 응답 데이터
	// if err := json.Unmarshal(respBody, &respJson); err != nil {
	// 	lprintf(4, "[INFO] if success is return orderId, check json error but it is normal [%s]\n", err)
	// 	// 정상인 경우 json 형태가 아닌 orderId 만 리턴하므로 이 경우에는 정상으로 처리
	// 	respData.Code = extra.SUCCESS
	//
	// 	respData.ServiceName = extra.TYPE
	// 	respData.Data = binary.BigEndian.Uint64(respBody)
	// 	return c.JSON(http.StatusOK, respData)
	// }

	// if respJson.Status == "ERROR" {
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	//respData.Data = &respJson
	// }

	// if strings.Contains(respJson.Data, "Website doesn't exist for") {
	// 	respData.Code = extra.FAIL
	//
	// } else {
	// 	respData.Code = extra.SUCCESS
	//
	// }

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 차일드 네임서버 추가
func SetDomainAddChildNameServerAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetDomainAddChildNameServerAllParam
	var respJson extra.RespSetDomainAddChildNameServerAll

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
		"cns":      {param.Cns},
		"ip":       param.Ip,
	}

	surl := fmt.Sprintf("%s/api/domains/add-cns.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	if strings.Contains(respJson.Data, "status\":\"Success") {
		respData.Code = extra.SUCCESS

	} else {
		respData.Code = extra.FAIL

	}

	//respData.Code = extra.SUCCESS
	//
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 차일드 네임서버 호스트 이름 수정
func ModiHostChildNameServerAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiHostChildNameServerAllParam
	var respJson extra.RespModiHostChildNameServerAll

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
		"old-cns":  {param.OldCns},
		"new-cns":  {param.NewCns},
	}

	surl := fmt.Sprintf("%s/api/domains/modify-cns-name.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 차일드 네임서버 IP 수정
func ModiIpChildNameServerOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiIpChildNameServerOneParam
	var respJson extra.RespModiIpChildNameServerOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
		"cns":      {param.Cns},
		"old-ip":   {param.OldIp},
		"new-ip":   {param.NewIp},
	}

	surl := fmt.Sprintf("%s/api/domains/modify-cns-ip.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	if strings.Contains(respJson.Data, "status\":\"Success") {
		respData.Code = extra.SUCCESS

	} else {
		respData.Code = extra.FAIL

	}

	//respData.Code = extra.SUCCESS
	//
	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 차일드 네임서버 삭제
func DelChildNameServerAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.DelChildNameServerAllParam
	var respJson extra.RespDelChildNameServerAll

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
		"cns":      {param.Cns},
		"ip":       param.Ip,
	}

	//surl := fmt.Sprintf("%s/api/domains/modify-cns-ip.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	surl := fmt.Sprintf("%s/api/domains/delete-cns-ip.xml?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	if strings.Contains(respJson.Data, "<string>status</string>\n<string>Success</string>") {
		respData.Code = extra.SUCCESS

	} else {
		respData.Code = extra.FAIL

	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 주문ID 사용하여 도메인 세부정보
func GetDomainDetailOrderIdAll(c echo.Context) error {
	respData := extra.RespGetDomainDetailOrderIdAll{}
	var param extra.GetDomainDetailOrderIdAllParam
	//var respJson extra.ApiCallResponseInterface
	var options string

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
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// https://demo.myorderbox.com/kb/node/770
	// for _, v := range param.Options {
	// 	options += fmt.Sprintf("&options=%s", v)
	// }
	options = "&options=OrderDetails"
	surl := fmt.Sprintf("%s/api/domains/details.json?api-key=%s&auth-userid=%s&order-id=%d%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.OrderId, options)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respData.Data : [%s]\n", respData.Data)

	// respMap, ok := respJson.Data.(map[string]interface{})
	// lprintf(4, "[INFO] respMap : [%s], ok : [%v]\n", respMap, ok)

	// respData.Data.Domsecret = respMap["domsecret"].(string)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 Renew
func SetDomainRenewWhoisOne(c echo.Context) error {
	respData := extra.DomainEntityApiCallResponse{}
	var param extra.SetDomainRenewWhoisOneParam
	var respJson extra.RespSetDomainRenewWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	years := strconv.FormatInt(int64(param.Years), 10)

	expireData, Getresult := GetDomainExpireDataOne(param.OrderId)
	if !Getresult {
		lprintf(1, "[ERR ] Renew domain, Get Expire Date Fail \n")
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	expDate := strconv.FormatInt(int64(expireData), 10)
	discountAmount := strconv.FormatFloat(param.DiscountAmount, 'f', -1, 32)
	autoNum := 0
	autorenew := strconv.FormatInt(int64(autoNum), 10)
	formData := url.Values{
		"order-id":        {orderId},
		"years":           {years},
		"exp-date":        {expDate},
		"invoice-option":  {extra.InvoiceOption},
		"discount-amount": {discountAmount},
		"auto-renew":      {autorenew},
	}

	surl := fmt.Sprintf("%s/api/domains/renew.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))

	/*
	   type RespSetDomainRenewWhoisOne struct {
	   	Actiontypedesc   string `json:"actiontypedesc"`
	   	Actionstatus     string `json:"actionstatus"`
	   	Entityid         string `json:"entityid"`
	   	Status           string `json:"status"`
	   	Eaqid            string `json:"eaqid"`
	   	Actiontype       string `json:"actiontype"`
	   	Description      string `json:"description"`
	   	Actionstatusdesc string `json:"actionstatusdesc"`
	   }
	*/

	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.Message.Cn = "Whois Error"
		respData.Message.Kr = "Whois Error"
		respData.Message.En = "Whois Error"
	} else {
		if respJson.Status == "Success" {
			respData.Code = extra.SUCCESS
			respData.Message.Cn = "Success"
			respData.Message.Kr = "Success"
			respData.Message.En = "Success"
			respData.Data.EntityId = respJson.Entityid
		} else {
			respData.Code = extra.C500
			respData.Message.Cn = "Whois Error"
			respData.Message.Kr = "Whois Error"
			respData.Message.En = "Whois Error"
		}
	}

	respData.ServiceName = extra.TYPE
	lprintf(4, "[INFO] respData : [%v]\n", respData)

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 도메인 복원
func SetDomainRestoreWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetDomainRestoreWhoisOneParam
	var respJson extra.RespSetDomainRestoreWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id":       {orderId},
		"invoice-option": {param.InvoiceOption},
	}

	surl := fmt.Sprintf("%s/api/domains/restore.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 도메인 이름 삭제
func DelDomainNameWhoisOne(c echo.Context) error {
	var respData extra.DomainDeleteResponse
	var param extra.WhoisDomainRefund
	var respJson extra.RespDelDomainNameWhois

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

	orderId := strconv.FormatInt(int64(param.OrderID), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/delete.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	/*
		if strings.Contains(respJson.Data, "<string>entityid</string>") {
			respData.Code = extra.SUCCESS
			respData.Message = "Success"
			entityIDArry1 := strings.Split(respJson.Data, "<string>entityid</string>") //두번째에 EntityID 들어있음
			lprintf(4, "[INFO] entityIdArry 1th : [%v]\n", entityIDArry1)

			entityIDArry2 := strings.Split(entityIDArry1[1], "<string>") //개행 <string>
			lprintf(4, "[INFO] entityIdArry 2th : [%v]\n", entityIDArry2)
			entityIDArry3 := strings.Split(entityIDArry2[1], "</string>") //entityID</string>
			lprintf(4, "[INFO] entityIdArry 3th : [%v]\n", entityIDArry3)

			//	entityId := strings.Split(entityIdArry[1], "<string>")
			respData.Data.EntityId = entityIDArry3[0]
			lprintf(4, "[INFO] EntityId : [%s]\n", respData.Data.EntityId)

		}
	*/

	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", string(respBody))

	if strings.Contains(string(respBody), "ERROR") {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
	} else {
		respData.Code = extra.SUCCESS

		respData.ServiceName = extra.TYPE
	}

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 기관 이전 요청 유효성 검사
func GetDomainValidateTransferOne(c echo.Context) error {
	respData := extra.RespGetDomainValidateTransferOne{}
	var param map[string]string
	var respJson extra.ApiCallResponseInterface
	var domainName string

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

	domainName = param["domainName"]
	surl := fmt.Sprintf("%s/api/domains/validate-transfer.json?api-key=%s&auth-userid=%s&domain-name=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, domainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))

	if err := json.Unmarshal(respBody, &respJson.Data); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	whoisResp, ok := respJson.Data.(map[string]interface{})
	lprintf(4, "[INFO] whoisResp: [%s], ok: [%v]\n", whoisResp, ok)

	if ok == true {
		respData.Data.Status = "false"
	} else {
		respData.Data.Status = string(respBody)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data.DomainName = domainName

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 기관 이전
func SetDomainValidateTransferOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetDomainValidateTransferOneParam
	var respJson extra.RespSetDomainValidateTransferOne

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

	// 요청 데이터 참고 : https://demo.myorderbox.com/kb/node/758
	// customerId := strconv.FormatInt(int64(param.CustomerId), 10)
	// regContactId := strconv.FormatInt(int64(param.RegContactId), 10)
	// adminContactId := strconv.FormatInt(int64(param.AdminContact), 10)
	// techContactId := strconv.FormatInt(int64(param.TechContactId), 10)
	// billingContactId := strconv.FormatInt(int64(param.BillingContactId), 10)

	formData := url.Values{}
	formData.Set("domain-name", param.DomainName)
	if param.AuthCode != "" {
		formData.Set("auth-code", param.AuthCode)
	}
	formData.Set("customer-id", extra.CustomerId)
	formData.Set("reg-contact-id", extra.RegContactId)
	formData.Set("admin-contact-id", extra.AdminContactId)
	formData.Set("tech-contact-id", extra.TechContactId)
	formData.Set("billing-contact-id", extra.BillingContactId)
	formData.Set("invoice-option", extra.InvoiceOption)

	// 일단은 고정으로 "NO", 화면에서 요청됨
	if param.PurchasePrivacy != "" {
		formData.Set("purchase-privacy", param.PurchasePrivacy)
	}
	if param.Ns != nil {
		for _, v := range param.Ns {
			formData.Add("ns", v)
		}
	}

	surl := fmt.Sprintf("%s/api/domains/transfer.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 인증코드 제출
func SetSubmitAuthCodeWhoisOne(c echo.Context) error {
	respData := extra.RespSetSubmitAuthCodeWhoisOne{}
	var param extra.SetSubmitAuthCodeWhoisOneParam
	var respJson interface{}

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id":  {orderId},
		"auth-code": {param.AuthCode},
	}

	surl := fmt.Sprintf("%s/api/domains/transfer/submit-auth-code.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))

	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	whoisResp, ok := respJson.(map[string]interface{})
	lprintf(4, "[INFO] whoisResp: [%s], ok: [%v]\n", whoisResp, ok)

	if ok == true {
		respData.Data.Status = whoisResp["status"].(string)
		respData.Data.Message = whoisResp["message"].(string)
	} else {
		respData.Data.Status = "success"
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 인증코드 수정
func ModiAuthCodeWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiAuthCodeWhoisOneParam
	var respJson extra.RespModiAuthCodeWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id":  {orderId},
		"auth-code": {param.AuthCode},
	}

	surl := fmt.Sprintf("%s/api/domains/modify-auth-code.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 이전 승인 메일 재발송
func SetResendRfaWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespSetResendRfaWhoisOne

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

	orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/resend-rfa.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 이전 취소
func DelCancelTransferWhoisOne(c echo.Context) error {
	var respData extra.DomainDeleteResponse
	var param extra.WhoisDomainRefund
	var respJson extra.RespDelDomainNameWhois

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

	orderId := strconv.FormatInt(int64(param.OrderID), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/delete.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	/*
		if strings.Contains(respJson.Data, "<string>entityid</string>") {
			respData.Code = extra.SUCCESS
			respData.Message = "Success"
			entityIDArry1 := strings.Split(respJson.Data, "<string>entityid</string>") //두번째에 EntityID 들어있음
			lprintf(4, "[INFO] entityIdArry 1th : [%v]\n", entityIDArry1)

			entityIDArry2 := strings.Split(entityIDArry1[1], "<string>") //개행 <string>
			lprintf(4, "[INFO] entityIdArry 2th : [%v]\n", entityIDArry2)
			entityIDArry3 := strings.Split(entityIDArry2[1], "</string>") //entityID</string>
			lprintf(4, "[INFO] entityIdArry 3th : [%v]\n", entityIDArry3)

			//	entityId := strings.Split(entityIdArry[1], "<string>")
			respData.Data.EntityId = entityIDArry3[0]
			lprintf(4, "[INFO] EntityId : [%s]\n", respData.Data.EntityId)

		}
	*/

	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if strings.Contains(string(respBody), "ERROR") {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
	} else {
		respData.Code = extra.SUCCESS

		respData.ServiceName = extra.TYPE
	}

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 등록자 연락처, 이메일 주소 확인 이메일 재전송
func SetResendVerificationWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespSetResendVerificationWhoisOne

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

	orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/raa/resend-verification.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 연락처 수정
func ModiContactWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiContactWhoisOneParam
	var respJson extra.RespModiContactWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	regContactId := strconv.FormatInt(int64(param.RegContactId), 10)
	adminContactId := strconv.FormatInt(int64(param.AdminContactId), 10)
	techContactId := strconv.FormatInt(int64(param.TechContactId), 10)
	billingContactId := strconv.FormatInt(int64(param.BillingContactId), 10)

	formData := url.Values{
		"order-id":           {orderId},
		"reg-contact-id":     {regContactId},
		"admin-contact-id":   {adminContactId},
		"tech-contact-id":    {techContactId},
		"billing-contact-id": {billingContactId},
	}

	surl := fmt.Sprintf("%s/api/domains/modify-contact.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 예약가능 여부 확인-IDN
func GetDomainIdnAvailableWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetDomainIdnAvailableWhoisAllParam
	var respJson extra.RespGGetDomainIdnAvailableWhoisAll
	var domainName string

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

	for _, v := range param.DomainName {
		domainName += fmt.Sprintf("&domain-name=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/idn-available.json?api-key=%s&auth-userid=%s%s&tld=%s&idnLanguageCode=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, domainName, param.Tld, param.IdnLanguageCode)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 예약가능 여부 확인-Premium Domains
func GetDomainPremiumWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetDomainPremiumWhoisAllParam
	var respJson extra.RespGetDomainPremiumWhoisAll
	var tlds string

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

	for _, v := range param.Tlds {
		tlds += fmt.Sprintf("&tlds=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/premium/available.json?api-key=%s&auth-userid=%s&key-word=%s%s&price-high=%d&no-of-result=%d",
		extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.KeyWord, tlds, param.PriceLow, param.NoOfResult)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 예약가능 여부 확인- 3rd level .NAME
func GetThirdLevelNameWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetThirdLevelNameWhoisAllParam
	var respJson extra.RespGetThirdLevelNameWhoisAll
	var domainName, tlds string

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

	for _, v := range param.DomainName {
		domainName += fmt.Sprintf("&domain-name=%s", v)
	}
	for _, v := range param.Tlds {
		tlds += fmt.Sprintf("&tlds=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/thirdlevelname/available.json?api-key=%s&auth-userid=%s%s%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, domainName, tlds)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 UK 도메인 이름에 대한 연락처 정보 확인
func GetUkWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetUkWhoisAllParam
	var respJson extra.RespGetUkWhoisAll
	var domainName, tlds string

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

	for _, v := range param.DomainName {
		domainName += fmt.Sprintf("&domain-name=%s", v)
	}
	for _, v := range param.Tlds {
		tlds += fmt.Sprintf("&tlds=%s", v)
	}

	surl := fmt.Sprintf("%s/api/domains/uk/available.json?api-key=%s&auth-userid=%s%s%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, domainName, tlds)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 프리미엄 도메인 확인
func GetPremiumCheckWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespGetPremiumCheckWhoisOne
	var domainName string

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

	domainName = param["domainName"]
	surl := fmt.Sprintf("%s/api/domains/premium-check.json?api-key=%s&auth-userid=%s&domain-name=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, domainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 고객 기본 서버 이름 가져오기
func GetCustomerDefaultNsWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespGetCustomerDefaultNsWhoisOne
	var customerId int

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

	customerId = param["customerId"]
	surl := fmt.Sprintf("%s/api/domains/customer-default-ns.json?api-key=%s&auth-userid=%s&customer-id=%d", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, customerId)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 Purchasing / Renewing Privacy Protection
func SetPurchasePrivacyWhois(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetPurchasePrivacyWhoisParam
	var respJson extra.RespSetPurchasePrivacyWhois

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	discountAmount := strconv.FormatFloat(param.DiscountAmount, 'f', -1, 32)
	formData := url.Values{
		"order-id":        {orderId},
		"invoice-option":  {param.InvoiceOption},
		"discount-amount": {discountAmount},
	}

	surl := fmt.Sprintf("%s/api/domains/purchase-privacy.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 개인정보보호 상태수정
func ModiPrivacyProtectionWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiPrivacyProtectionWhoisOneParam
	var respJson extra.RespModiPrivacyProtectionWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id":        {orderId},
		"protect-privacy": {param.ProtectPrivacy},
		"reason":          {param.Reason},
	}

	surl := fmt.Sprintf("%s/api/domains/modify-privacy-protection.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 도난방지 잠금 사용
func SetEnableTheftProtectionWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.OrderIdParam
	var respJson extra.RespSetEnableTheftProtectionWhoisOne

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

	// orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/enable-theft-protection.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	if strings.Contains(string(respBody), "Failed") {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 도난방지 잠금 해제
func SetDisableTheftProtectionWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.OrderIdParam
	var respJson extra.RespSetDisableTheftProtectionWhoisOne

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

	// orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/disable-theft-protection.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	if strings.Contains(string(respBody), "Failed") {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈에 도메인 잠금 목록 가져오기
func GetLockWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespGetLockWhoisOne
	var orderId int

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

	orderId = param["orderId"]
	surl := fmt.Sprintf("%s/api/domains/locks.json?api-key=%s&auth-userid=%s&order-id=%d", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, orderId)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 TEL Whois Preference 수정하기
func ModiTelWhoisPrefWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiTelWhoisPrefWhoisOneParam
	var respJson extra.RespModiTelWhoisPrefWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id":   {orderId},
		"whois-type": {param.WhoisType},
		"publish":    {param.Publish},
	}

	surl := fmt.Sprintf("%s/api/domains/tel/modify-whois-pref.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 UK 도메인 이름 해제
func SetUkWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetUkWhoisOneParam
	var respJson extra.RespSetUkWhoisOne

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

	orderId := strconv.FormatInt(int64(param.OrderId), 10)
	formData := url.Values{
		"order-id": {orderId},
		"new-tag":  {param.NewTag},
	}

	surl := fmt.Sprintf("%s/api/domains/uk/release.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 Rechecking NS with .DE Registry
func GetRecheckNsWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespSetUkWhoisOne

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

	orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/de/recheck-ns.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 DotxxxAssociationDetails
func GetDotxxxAssociationDetailsOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespGetDotxxxAssociationDetailsOne

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

	orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/dotxxx/association-details.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 위임자 서명자 (DS) 레코드 추가
func SetDnsSecWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespSetDnsSecWhoisOne

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

	orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/add-dnssec.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 위임자 서명자 (DS) 레코드 삭제
func DelDnsSecWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]int
	var respJson extra.RespSetDnsSecWhoisOne

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

	orderId := strconv.FormatInt(int64(param["orderId"]), 10)
	formData := url.Values{
		"order-id": {orderId},
	}

	surl := fmt.Sprintf("%s/api/domains/del-dnssec.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 희망 목록에 도메인 이름 추가
func SetPreorderingWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespSetPreorderingWhoisOne

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

	formData := url.Values{
		"customer-id": {param["customerId"]},
	}

	surl := fmt.Sprintf("%s/api/domains/preordering/add.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 희망 목록에 도메인 이름 삭제
func DelPreorderingWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespSetPreorderingWhoisOne

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

	formData := url.Values{
		"customer-id": {param["customerId"]},
	}

	surl := fmt.Sprintf("%s/api/domains/preordering/delete.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 희망 목록 가져 오기
func GetPreorderingWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetPreorderingWhoisOneParam
	var respJson extra.RespGetPreorderingWhoisOne

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

	surl := fmt.Sprintf("%s/api/domains/preordering/fetch.json?api-key=%s&auth-userid=%s&customer-id=%d&page-no=%d&no-of-records=%d",
		extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, param.CustomerId, param.PageNo, param.NoOfRecords)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 카테고리를 기반으로 희망목록 TLD 가져 오기
func GetPreorderingCategoryWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespGetPreorderingCategoryWhoisAll
	var category string

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

	category = param["category"]
	surl := fmt.Sprintf("%s/api/domains/preordering/fetchtldlist.json?api-key=%s&auth-userid=%s&category=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, category)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 상표권 주장 데이터 가져 오기
func GetTmNoticeWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespGetTmNoticeWhoisOne
	var lookupKey string

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

	lookupKey = param["lookupKey"]
	surl := fmt.Sprintf("%s/api/domains/get-tm-notice.json?api-key=%s&auth-userid=%s&lookup-key=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, lookupKey)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 Fetching the List of TLDs in Sunrise / Landrush Period
func GetTldsInPhaseWhoisAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespGetTldsInPhaseWhoisAll
	var phase string

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

	phase = param["phase"]
	surl := fmt.Sprintf("%s/api/domains/tlds-in-phase.json?api-key=%s&auth-userid=%s&phase=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, phase)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 단계별로 서명 된 TLD 세부 정보 얻기
func GetTldInfoWhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var respJson extra.RespGetTldInfoWhoisOne

	formData := url.Values{}

	surl := fmt.Sprintf("%s/api/domains/tld-info.json?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 후이즈 희망 목록에 도메인 이름 추가
func SetCustomerV2WhoisOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespSetCustomerV2WhoisOne

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

	formData := url.Values{
		"auth-userid": {param["authUserid"]},
	}

	surl := fmt.Sprintf("%s/api/customers/v2/signup.xml?api-key=%s&auth-userid=%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)
	lprintf(4, "[INFO] formData : [%v]\n", formData)
	respForm, err := http.PostForm(surl, formData)
	if err != nil {
		lprintf(1, "[ERR ] PostForm : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respForm.Body.Close()

	respBody, err := ioutil.ReadAll(respForm.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))
	respJson.Data = string(respBody)

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

//후이즈 예치금 조회
func GetAvailableBalanceOne(c echo.Context) error {
	respData := extra.BalanceApiCallResponse{}
	var respJson extra.RespAvailableBalance

	surl := fmt.Sprintf("%s/api/billing/reseller-balance.json?api-key=%s&auth-userid=%s&reseller-id=%s",
		extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, extra.AuthUserId)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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

	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respData: [%v]\n", respJson)
	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data.AvailableBalanceValue = strconv.FormatFloat(respJson.AvailableBalanceValue, 'f', -1, 64)

	return c.JSON(http.StatusOK, respData)
}

// main - domain 정보 상세보기
func GetDomainInfoWhoxy(c echo.Context) error {
	respData := extra.RespGetDomainInfoWhoxy{}
	rawWhois := extra.RespWhoxy{}

	var param map[string]string

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

	domainName := param["domainName"]
	surl := fmt.Sprintf("http://api.whoxy.com/?key=%s&whois=%s", WHOXY_KEY, domainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	respGet, err := http.Get(surl)
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
	lprintf(4, "[INFO] respData: [%v]\n", string(respBody))

	if err := json.Unmarshal(respBody, &rawWhois); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Data.Data = rawWhois.RawWhois

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	return c.JSON(http.StatusOK, respData)

}

// 후이즈에 주문ID 사용하여 내부적으로 도메인 만료일 구하기
func GetDomainExpireDataOne(oid int) (int, bool) {
	respData := extra.RespGetDomainExpireDataOne{}
	//var respJson extra.ApiCallResponseInterface
	var options string

	lprintf(4, "[INFO] request Expire Data get order ID : [%s]\n", oid)

	options = "&options=OrderDetails"
	surl := fmt.Sprintf("%s/api/domains/details.json?api-key=%s&auth-userid=%s&order-id=%d%s", extra.WhoisUrl, extra.ApiKey, extra.AuthUserId, oid, options)
	lprintf(4, "[INFO] surl     : [%s]\n", surl)

	respGet, err := http.Get(surl)
	if err != nil {
		lprintf(1, "[ERR ] Get : [%s]\n", err)
		return 0, false
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		return 0, false
	}
	lprintf(4, "[INFO] respBody: [%s]\n", string(respBody))

	if err := json.Unmarshal(respBody, &respData); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		return 0, false
	}
	lprintf(4, "[INFO] respEndTime : [%s]\n", respData.EndTime)

	if len(respData.EndTime) != 0 {
		expireData, err3 := strconv.Atoi(respData.EndTime)
		if err3 != nil {
			lprintf(1, "[ERR ] Atoi err  : [%s]\n", err3)
			return 0, false
		}
		return expireData, true

	}

	// respMap, ok := respJson.Data.(map[string]interface{})
	// lprintf(4, "[INFO] respMap : [%s], ok : [%v]\n", respMap, ok)
	// respData.Data.Domsecret = respMap["domsecret"].(string)
	return 0, false
}
