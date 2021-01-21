package controller

import (
	"bytes"
	_ "encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	extra "hydrawebapi/extraapi/model"

	_ "charlie/i3.0.0/cls"

	"github.com/labstack/echo"
)

// 도메인 등록 여부 (등록가능하면 Success, 그 외에는 C500)
func GetGabiaDomainCheck(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.DomainNameParam
	var respJson extra.RespGetGabiaDomainCheck

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

	lprintf(4, "[INFO] DomainName 2: [%s]\n", param.DomainName)

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/check?domain_name=%s", extra.GabiaUrl, param.DomainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "32000" {
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 등록 (등록 성공하면 Success, 실패하면 C500)
func SetGabiaDomain(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.SetGabiaDomainParam
	//var param extra.SetGabiaDomainJsonParam
	var respJson extra.RespSetGabiaDomain

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
	lprintf(4, "[INFO] Request body : [%s]\n", param)

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	var param2 extra.SetGabiaDomainJsonParam

	param2.DomainName = param.DomainName
	param2.Data.Period = param.Period
	param2.Data.AutoLockYn = param.AutoLockYn
	param2.Data.HanName = param.HanName
	param2.Data.EngName = param.EngName
	param2.Data.HanOrg = param.HanOrg
	param2.Data.EngOrg = param.EngOrg
	param2.Data.Gubun = param.Gubun
	param2.Data.CtfyCode = param.CtfyCode
	param2.Data.CtfyNo = param.CtfyNo
	param2.Data.OpenInfoYn = param.OpenInfoYn
	param2.Data.Zip = param.Zip
	param2.Data.Addr1 = param.Addr1
	param2.Data.Addr2 = param.Addr2
	param2.Data.Eaddr1 = param.Eaddr1
	param2.Data.Eaddr2 = param.Eaddr2
	param2.Data.City = param.City
	param2.Data.CountryCode = param.CountryCode
	param2.Data.Phone = param.Phone
	param2.Data.Fax = param.Fax
	param2.Data.Hp = param.Hp
	param2.Data.Email = param.Email
	param2.Data.AhanName = param.AhanName
	param2.Data.AengName = param.AengName
	param2.Data.Azip = param.Azip
	param2.Data.Aaddr1 = param.Aaddr1
	param2.Data.Aaddr2 = param.Aaddr2
	param2.Data.Aeaddr1 = param.Aeaddr1
	param2.Data.Aeaddr2 = param.Aeaddr2
	param2.Data.Acity = param.Acity
	param2.Data.AcountryCode = param.AcountryCode
	param2.Data.Aphone = param.Aphone
	param2.Data.Afax = param.Afax
	param2.Data.Ahp = param.Ahp
	param2.Data.SmsYn = param.SmsYn
	param2.Data.Aemail = param.Aemail

	for z := 0; z < len(param.Ns); z++ {
		if z == 0 {
			param2.Data.Ns1 = param.Ns[z]
		} else if z == 1 {
			param2.Data.Ns2 = param.Ns[z]
		} else if z == 2 {
			param2.Data.Ns3 = param.Ns[z]
		} else if z == 3 {
			param2.Data.Ns4 = param.Ns[z]
		} else if z == 4 {
			param2.Data.Ns5 = param.Ns[z]
		} else if z == 5 {
			param2.Data.Ns6 = param.Ns[z]
		} else if z == 6 {
			param2.Data.Ns7 = param.Ns[z]
		} else if z == 7 {
			param2.Data.Ns8 = param.Ns[z]
		} else if z == 8 {
			param2.Data.Ns9 = param.Ns[z]
		} else if z == 9 {
			param2.Data.Ns10 = param.Ns[z]
		} else if z == 10 {
			param2.Data.Ns11 = param.Ns[z]
		} else if z == 11 {
			param2.Data.Ns12 = param.Ns[z]
		} else if z == 12 {
			param2.Data.Ns13 = param.Ns[z]
		}

	}

	reqBytes, err3 := json.Marshal(param2)
	if err3 != nil {
		lprintf(1, "[ERR ] Marshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	buff := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequest("POST", surl, buff)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)
	lprintf(4, "[INFO] point...test : \n")

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 삭제( 성공 Suceess, 실패 c500)
func DelGabiaDomain(c echo.Context) error {
	respData := extra.DomainDeleteResponse{}
	var param extra.DomainNameParam
	var respJson extra.RespDelGabiaDomain

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("domain_name", param.DomainName)

	req, err := http.NewRequest("DELETE", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 정보 조회
func GetGabiaDomain(c echo.Context) error {
	respData := extra.ApiRespGetGabiaDomain{}
	var param extra.DomainNameParam
	var respJson extra.RespGetGabiaDomain

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains?domain_name=%s", extra.GabiaUrl, param.DomainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)
	lprintf(4, "[INFO] Gabia Token Value : [%s]\n", extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = respJson

	return c.JSON(http.StatusOK, respData)
}

// 관리자 정보 변경
func ModiGabiaDomainAdmin(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.ModiGabiaDomainAdminParam
	var respJson extra.RespModiGabiaDomainAdmin

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/admins", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	var param2 extra.ModiGabiaDomainAdminJsonParam

	param2.DomainName = param.DomainName
	param2.Data.AhanName = param.AhanName
	param2.Data.AengName = param.AengName
	param2.Data.Azip = param.Azip
	param2.Data.Aaddr1 = param.Aaddr1
	param2.Data.Aaddr2 = param.Aaddr2
	param2.Data.Aeaddr1 = param.Aeaddr1
	param2.Data.Aeaddr2 = param.Aeaddr2
	param2.Data.Acity = param.Acity
	param2.Data.AcountryCode = param.AcountryCode
	param2.Data.Aphone = param.Aphone
	param2.Data.Afax = param.Afax
	param2.Data.Ahp = param.Ahp
	param2.Data.SmsYn = param.SmsYn
	param2.Data.Aemail = param.Aemail

	reqBytes, err3 := json.Marshal(param2)
	if err3 != nil {
		lprintf(1, "[ERR ] Marshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	buff := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequest("PUT", surl, buff)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 네임서버 변경
func ModiGabiaDomainNameServer(c echo.Context) error {
	respData := extra.RespGabiaNSModifyWhoisOne{}
	var param extra.ModiGabiaDomainNameServerParam
	var respJson extra.RespModiGabiaDomainNameServer

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/nameservers", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	/*
		data := url.Values{}
		data.Set("domain_name", param.DomainName)
		nulLen := 0
		keyValue := ""
		for i := 0; i < len(param.Ns); i++ {
			keyValue = fmt.Sprintf("ns%d", i+1)
			data.Set(keyValue, param.Ns[i])
			nulLen = nulLen + 1
			keyValue = fmt.Sprintf("nsip%d", i+1)
			data.Set(keyValue, "")
		}

		for j := nulLen; j < 13; j++ {
			keyValue = fmt.Sprintf("ns%d", j+1)
			data.Set(keyValue, "")
			keyValue = fmt.Sprintf("nsip%d", j+1)
			data.Set(keyValue, "")
		}
	*/

	var param2 extra.ModiGabiaDomainNameServerJsonParam
	param2.DomainName = param.DomainName

	for z := 0; z < len(param.Ns); z++ {
		if z == 0 {
			param2.Data.Ns1 = param.Ns[z]
		} else if z == 1 {
			param2.Data.Ns2 = param.Ns[z]
		} else if z == 2 {
			param2.Data.Ns3 = param.Ns[z]
		} else if z == 3 {
			param2.Data.Ns4 = param.Ns[z]
		} else if z == 4 {
			param2.Data.Ns5 = param.Ns[z]
		} else if z == 5 {
			param2.Data.Ns6 = param.Ns[z]
		} else if z == 6 {
			param2.Data.Ns7 = param.Ns[z]
		} else if z == 7 {
			param2.Data.Ns8 = param.Ns[z]
		} else if z == 8 {
			param2.Data.Ns9 = param.Ns[z]
		} else if z == 9 {
			param2.Data.Ns10 = param.Ns[z]
		} else if z == 10 {
			param2.Data.Ns11 = param.Ns[z]
		} else if z == 11 {
			param2.Data.Ns12 = param.Ns[z]
		} else if z == 12 {
			param2.Data.Ns13 = param.Ns[z]
		}

	}

	reqBytes, err3 := json.Marshal(param2)
	if err3 != nil {
		lprintf(1, "[ERR ] Marshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	buff := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequest("PUT", surl, buff)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

	/*
		req, err := http.NewRequest("PUT", surl, strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "Basic "+extra.GabiaKey)
	*/
	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		respData.Data.Status = "ERROR"
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		respData.Data.Status = "ERROR"
		return c.JSON(http.StatusOK, respData)
	}

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		respData.Data.Status = "ERROR"
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		respData.Data.Status = "ERROR"
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.FAIL

		respData.ServiceName = extra.TYPE
		respData.Data.Status = "ERROR"
		respData.Data.Msg = respJson.Msg
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data.Status = "Success"
	respData.Data.Msg = respJson.Msg

	return c.JSON(http.StatusOK, respData)
}

// 소유자 정보 변경
func ModiGabiaDomainOwner(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.ModiGabiaDomainOwnerParam
	var respJson extra.RespModiGabiaDomainOwner

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/owners", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	var param2 extra.ModiGabiaDomainOwnerJsonParam
	param2.DomainName = param.DomainName
	param2.Data.HanName = param.HanName
	param2.Data.EngName = param.EngName
	param2.Data.HanOrg = param.HanOrg
	param2.Data.EngOrg = param.EngOrg
	param2.Data.Gubun = param.Gubun
	param2.Data.CtfyCode = param.CtfyCode
	param2.Data.CtfyNo = param.CtfyNo
	param2.Data.OpenInfoYn = param.OpenInfoYn
	param2.Data.Zip = param.Zip
	param2.Data.Addr1 = param.Addr1
	param2.Data.Addr2 = param.Addr2
	param2.Data.Eaddr1 = param.Eaddr1
	param2.Data.Eaddr2 = param.Eaddr2
	param2.Data.City = param.City
	param2.Data.CountryCode = param.CountryCode
	param2.Data.Phone = param.Phone
	param2.Data.Fax = param.Fax
	param2.Data.Hp = param.Hp
	param2.Data.Email = param.Email

	reqBytes, err3 := json.Marshal(param2)
	if err3 != nil {
		lprintf(1, "[ERR ] Marshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	buff := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequest("PUT", surl, buff)

	//req, err := http.NewRequest("PUT", surl, strings.NewReader(data.Encode()))
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 잠금 상태 변경
func ModiGabiaDomainStatusLock(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.ModiGabiaDomainStatusLockParam
	var respJson extra.RespModiGabiaDomainStatusLock

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/statuses/lock", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("domain_name", param.DomainName)
	data.Set("auto_lock_yn", param.AutoLockYn)

	req, err := http.NewRequest("PUT", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 연장
func SetGabiaDomainIncrease(c echo.Context) error {
	respData := extra.ApiRespSetGabiaDomainIncrease{}
	var param extra.SetGabiaDomainIncreaseParam
	var respJson extra.RespSetGabiaDomainIncrease

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/increases", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	period := strconv.FormatInt(int64(param.Period), 10)

	data := url.Values{}
	data.Set("domain_name", param.DomainName)
	data.Set("period", period)

	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = respJson

	return c.JSON(http.StatusOK, respData)
}

// KR 도메인 연장 취소
func DelGabiaDomainIncrease(c echo.Context) error {
	respData := extra.ApiRespDelGabiaDomainIncrease{}
	var param extra.DomainNameParam
	var respJson extra.RespDelGabiaDomainIncrease

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/increases", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("domain_name", param.DomainName)

	req, err := http.NewRequest("DELETE", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = respJson

	return c.JSON(http.StatusOK, respData)
}

// 도메인 삭제 복구 신청
func SetGabiaDomainRestore(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.DomainNameParam
	var respJson extra.RespSetGabiaDomainRestore

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/restores", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("domain_name", param.DomainName)

	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 삭제 복구 신청 가능 여부 조회
func GetGabiaDomainRestoreCheck(c echo.Context) error {
	respData := extra.GabiaApiCallResponse{}
	var param extra.DomainNameParam
	var respJson extra.RespGetGabiaDomainRestoreCheck

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/restores/check?domain_name=%s", extra.GabiaUrl, param.DomainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 안이전
func SetGabiaDomainTransfer(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetGabiaDomainTransferParam
	var respJson extra.RespSetGabiaDomainTransfer

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/transfers", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	var param2 extra.SetGabiaDomainTransferJsonParam

	param2.DomainName = param.DomainName
	param2.Data.AuthInfo = param.AuthInfo
	param2.Data.HanName = param.HanName
	param2.Data.EngName = param.EngName
	param2.Data.HanOrg = param.HanOrg
	param2.Data.EngOrg = param.EngOrg
	param2.Data.Gubun = param.Gubun
	param2.Data.CtfyCode = param.CtfyCode
	param2.Data.CtfyNo = param.CtfyNo
	param2.Data.OpenInfoYn = param.OpenInfoYn
	param2.Data.Zip = param.Zip
	param2.Data.Addr1 = param.Addr1
	param2.Data.Addr2 = param.Addr2
	param2.Data.Eaddr1 = param.Eaddr1
	param2.Data.Eaddr2 = param.Eaddr2
	param2.Data.City = param.City
	param2.Data.CountryCode = param.CountryCode
	param2.Data.Phone = param.Phone
	param2.Data.Fax = param.Fax
	param2.Data.Hp = param.Hp
	param2.Data.Email = param.Email
	param2.Data.AhanName = param.AhanName
	param2.Data.AengName = param.AengName
	param2.Data.Azip = param.Azip
	param2.Data.Aaddr1 = param.Aaddr1
	param2.Data.Aaddr2 = param.Aaddr2
	param2.Data.Aeaddr1 = param.Aeaddr1
	param2.Data.Aeaddr2 = param.Aeaddr2
	param2.Data.Acity = param.Acity
	param2.Data.AcountryCode = param.AcountryCode
	param2.Data.Aphone = param.Aphone
	param2.Data.Afax = param.Afax
	param2.Data.Ahp = param.Ahp
	param2.Data.SmsYn = param.SmsYn
	param2.Data.Aemail = param.Aemail

	reqBytes, err3 := json.Marshal(param2)
	if err3 != nil {
		lprintf(1, "[ERR ] Marshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	buff := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequest("POST", surl, buff)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 안이전 조회
func GetGabiaDomainTransferCheck(c echo.Context) error {
	respData := extra.ApiRespGetGabiaDomainTransferCheck{}
	var param extra.GetGabiaDomainTransferCheckParam
	var respJson extra.RespGetGabiaDomainTransferCheck

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/domains/transfers/check?domain_name=%s&type=%s", extra.GabiaUrl, param.DomainName, "interior")
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = respJson

	return c.JSON(http.StatusOK, respData)
}

// 호스트 등록 여부
func GetGabiaDomainHostCheck(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetGabiaDomainHostCheckParam
	var respJson extra.RespGetGabiaDomainHostCheck

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/hosts/check?host_name=%s&type=%s", extra.GabiaUrl, param.HostName, param.Type) //kr인경우 ".kr"
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 호스트 생성
func SetGabiaHost(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.SetGabiaHostParam
	var respJson extra.RespSetGabiaHost

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/hosts", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("host_name", param.HostName)
	data.Set("host_ip", param.HostIp)
	data.Set("type", param.Type)

	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 호스트 정보 조회
func GetGabiaHost(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param map[string]string
	var respJson extra.RespGetGabiaDomainHostCheck

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/hosts?host_name=%s", extra.GabiaUrl, param["hostName"])
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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
	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 호스트 정보 변경
func ModiGabiaHost(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.ModiGabiaHostParam
	var respJson extra.RespModiGabiaHost

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/hosts", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("host_name", param.HostName)
	data.Set("old_host_ip", param.OldHostIp)
	data.Set("new_host_ip", param.NewHostIp)

	req, err := http.NewRequest("PUT", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 호스트 삭제
func DelGabiaHost(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.DelGabiaHostParam
	var respJson extra.RespDelGabiaHost

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

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/hosts", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	data := url.Values{}
	data.Set("domain_name", param.HostName)
	data.Set("domain_name", param.Type)

	req, err := http.NewRequest("DELETE", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 예치금 조회
func GetGabiaDeposit(c echo.Context) error {
	respData := extra.ApiRespGetGabiaDeposit{}
	var respJson extra.RespGetGabiaDeposit

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := fmt.Sprintf("%s/deposits", extra.GabiaUrl)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err1 := json.Unmarshal(respBody, &respJson); err1 != nil {
		lprintf(1, "[ERR ] Unmarshal Error: [%s]\n", err1)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] respBody : [%v]\n", respJson)

	if respJson.GabiaCode != "10000" {
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = respJson

	return c.JSON(http.StatusOK, respData)
}

//가비아 토큰 요청 함수 (dhlqn)
func GetGabiaAuthToken2(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var respJson extra.RespGetGabiaAuthToken

	keyValue := extra.GabiaID + ":" + extra.GabiaPass
	extra.GabiaKey = GetBase64Encoding(keyValue)
	lprintf(4, "[INFO] after replace Authorization key : [%s]\n", extra.GabiaKey)

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := "https://domainpartnerapi.gabia.com/oauth/authen/get_auth_token"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	reqBody := bytes.NewBufferString("grant_type=client_credentials")

	req, err := http.NewRequest("POST", surl, reqBody)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	if strings.Contains(string(respBody), "token is expire") {
		lprintf(1, "[ERR ] TOKEN IS EXpir \n")
		go GetGabiaAuthToken()
		respData.Code = extra.RETRY

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] RespJson : [%v]\n", respJson)
	if respJson.GabiaCode == true {
		lprintf(4, "[INFO] Get Token Value \n")
		extra.GabiaKeyTTL = respJson.ExpireTTL
		lprintf(4, "[INFO] Get Token TTL(%d) \n", extra.GabiaKeyTTL)
		extra.GabiaKey = respJson.AccessToken
		lprintf(4, "[INFO] Get Token Key(%s) \n", extra.GabiaKey)
		extra.GabiaKeyInTime = respJson.ACCESSTokenTime
		lprintf(4, "[INFO] Get Token Key InTime(%s) \n", extra.GabiaKeyInTime)

	} else {
		lprintf(4, "[INFO] Get Token Value Fail, Msg: %s \n", respJson.Msg)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = extra.GabiaKey

	return c.JSON(http.StatusOK, respData)
}

//가비아 토큰 요청 함수
func GetGabiaAuthToken() bool {
	var respJson extra.RespGetGabiaAuthToken

	keyValue := extra.GabiaID + ":" + extra.GabiaPass
	extra.GabiaKey = GetBase64Encoding(keyValue)
	lprintf(4, "[INFO] after replace Authorization key : [%s]\n", extra.GabiaKey)

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	surl := "https://domainpartnerapi.gabia.com/oauth/authen/get_auth_token"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	reqBody := bytes.NewBufferString("grant_type=client_credentials")

	req, err := http.NewRequest("POST", surl, reqBody)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

	lprintf(4, "[INFO] Gabia Token Value : [%s]\n", extra.GabiaKey)

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		return false
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		return false
	}

	if err := json.Unmarshal(respBody, &respJson); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		return false
	}

	lprintf(4, "[INFO] RespJson : [%v]\n", respJson)
	if respJson.GabiaCode == true {
		lprintf(4, "[INFO] Get Token Value \n")
		extra.GabiaKeyTTL = respJson.ExpireTTL
		lprintf(4, "[INFO] Get Token TTL(%d) \n", extra.GabiaKeyTTL)
		extra.GabiaKey = GetBase64Encoding(extra.GabiaID + ":" + respJson.AccessToken)
		lprintf(4, "[INFO] Get Token Key(%s) \n", extra.GabiaKey)
		extra.GabiaKeyInTime = respJson.ACCESSTokenTime
		lprintf(4, "[INFO] Get Token Key InTime(%s) \n", extra.GabiaKeyInTime)
	} else {
		lprintf(4, "[INFO] Get Token Value Fail, Msg: %s \n", respJson.Msg)
		return false
	}

	return true
}

//가비아 토큰 갱신 스케줄러(주기는 TTL의 2분의 1)
func GetGabiaAuthTokenScheduler() {

	for {
		lprintf(4, "[INFO] Gabia Auth Token Get Scheduler start \n")
		result := GetGabiaAuthToken()
		if result == true {
			lprintf(4, "[INFO] Gabia Scheduler TTL (%d) roop \n", extra.GabiaKeyTTL/2)
			time.Sleep(time.Duration(extra.GabiaKeyTTL/2) * time.Second)
		} else {
			lprintf(4, "[ERR ] Gabia Get auth Token value fail, \n")
		}

	}
}

// 도메인 등록 여부 테스트
func TestGetGabiaDomainCheck(c echo.Context) error {
	respData := extra.TestGabiaApiCallResponse{}

	//var respJson extra.RespGetGabiaDomainCheck

	surl := fmt.Sprintf("%s/domains/check?domain_name=securitynetsvc.com,securitynet-test.com", extra.GabiaUrl)

	// ex) https://domainpartnerapi.gabia.com/domains/check?domain_name=amur21-20180122-qa11818.co.kr
	//surl := fmt.Sprintf("%s/domains/check?domain_name=%s", extra.GabiaUrl, param.DomainName)
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+extra.GabiaKey)

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

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = string(respBody)

	return c.JSON(http.StatusOK, respData)
}
