package controller

import (
	"bytes"
	b64 "encoding/base64"
	_ "encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	_ "time"

	extra "hydrawebapi/extraapi/model"

	_ "charlie/i3.0.0/cls"

	"github.com/labstack/echo"
)

// func Test2(c echo.Context) error {

// }

// func Test(c echo.Context) error {

// 	barCode := `<form id="alipaysubmit" name="alipaysubmit" action="` + extra.AliOperUrl + `?" method="GET" style='display:none;'>
// 	   			<input type="hidden" name="service" value="` + param.Service + `">
// 	   			<input type="hidden" name="partner" value="` + param.Partner + `">
// 	   			<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
// 	               <input type="hidden" name="sign_type" value="` + param.SignType + `">
// 	               <input type="hidden" name="sign" value="` + param.Sign + `">
// 	   			<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
// 	   			<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
// 	               <input type="hidden" name="refer_url" value="` + param.ReferUrl + `">
// 	   			<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
// 	   			<input type="hidden" name="subject" value="` + param.Subject + `">
// 	               <input type="hidden" name="currency" value="` + param.Currency + `">
// 	   			<input type="hidden" name="total_fee" value="` + param.TotalFee + `">
// 	   			<input type="hidden" name="body" value="` + param.Body + `">
// 	               <input type="hidden" name="product_code" value="` + param.ProductCode + `">
// 	   		</form>
// 	   		<script>document.forms['alipaysubmit'].submit();</script>`

// 	lprintf(4, "[INFO] paypar payment : [%s]\n", barCode)

// 	return c.HTML(http.StatusOK, barCode)

// }

// 빌링에 넣어서 만드는중. 이룸.
// // Paypal 결제 요청 시 access token 얻어오기
func getAccessTokenfromPaypal() extra.GetPaypalAccessTokenOne {
	returnData := extra.GetPaypalAccessTokenOne{}
	var surl, userInfo string

	if extra.Mode {
		surl = "https://api.paypal.com/v1/oauth2/token"
		userInfo = "AUD6dRCvTv4ahnU3yzsDEhGz5DeGVa6uyMd4wh_bK62WYQ1U-7ZKg8n2ScTK:EB6sJhC4tZ23KscPw5JjJRap23aKp3Dh3EqFKcvGw-rjdrK9S_DgIFtjNaHS"
	} else {
		surl = "https://api.sandbox.paypal.com/v1/oauth2/token"
		userInfo = "ASjy5RBTi-izr5nOHGprh-yftEBN8mD8Rt9gNsZirzBvuraXy1-DsWL6y7ox:ECu5lxAzafKG2mDxFrLdg7N8RC8Gk_f0vv09pS66EYXYu-KUEErtR-o0RsC9"
	}
	//surl := "https://api.sandbox.paypal.com/v1/oauth2/token" // test url
	//surl = "https://api.paypal.com/v1/oauth2/token" // 운영 url

	// make Authorization value
	// paypal 관리자 화면 접속해서 [clientId:secret] 데이터 가져옴
	// https://developer.paypal.com/docs/api-basics/manage-apps/#create-or-edit-sandbox-and-live-apps
	// userInfo := "ASjy5RBTi-izr5nOHGprh-yftEBN8mD8Rt9gNsZirzBvuraXy1-DsWL6y7ox:ECu5lxAzafKG2mDxFrLdg7N8RC8Gk_f0vv09pS66EYXYu-KUEErtR-o0RsC9" // 개발
	//userInfo := "AUD6dRCvTv4ahnU3yzsDEhGz5DeGVa6uyMd4wh_bK62WYQ1U-7ZKg8n2ScTK:EB6sJhC4tZ23KscPw5JjJRap23aKp3Dh3EqFKcvGw-rjdrK9S_DgIFtjNaHS" // 운영
	sEnc := b64.StdEncoding.EncodeToString([]byte(userInfo))
	lprintf(4, "[INFO] sEnc : %s", sEnc)
	authorization := fmt.Sprintf("Basic %s", sEnc)

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	if err != nil {
		lprintf(1, "[ERR ] NewRequest : [%s]\n", err)
		return extra.GetPaypalAccessTokenOne{}
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Language", "en_US")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		return extra.GetPaypalAccessTokenOne{}
	}
	defer respGet.Body.Close()

	resp, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		return extra.GetPaypalAccessTokenOne{}
	}
	lprintf(4, "[INFO] access-token resp : [%s]\n", string(resp))

	if err := json.Unmarshal(resp, &returnData); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		return extra.GetPaypalAccessTokenOne{}
	}

	return returnData
}

func getAccessTokenfromPaypal2() string {

	surl := "https://api.sandbox.paypal.com/v1/oauth2/token"
	userInfo := "ASjy5RBTi-izr5nOHGprh-yftEBN8mD8Rt9gNsZirzBvuraXy1-DsWL6y7ox:ECu5lxAzafKG2mDxFrLdg7N8RC8Gk_f0vv09pS66EYXYu-KUEErtR-o0RsC9"

	//surl := "https://api.sandbox.paypal.com/v1/oauth2/token" // test url
	//surl = "https://api.paypal.com/v1/oauth2/token" // 운영 url

	// make Authorization value
	// paypal 관리자 화면 접속해서 [clientId:secret] 데이터 가져옴
	// https://developer.paypal.com/docs/api-basics/manage-apps/#create-or-edit-sandbox-and-live-apps
	// userInfo := "ASjy5RBTi-izr5nOHGprh-yftEBN8mD8Rt9gNsZirzBvuraXy1-DsWL6y7ox:ECu5lxAzafKG2mDxFrLdg7N8RC8Gk_f0vv09pS66EYXYu-KUEErtR-o0RsC9" // 개발
	//userInfo := "AUD6dRCvTv4ahnU3yzsDEhGz5DeGVa6uyMd4wh_bK62WYQ1U-7ZKg8n2ScTK:EB6sJhC4tZ23KscPw5JjJRap23aKp3Dh3EqFKcvGw-rjdrK9S_DgIFtjNaHS" // 운영
	sEnc := b64.StdEncoding.EncodeToString([]byte(userInfo))
	//lprintf(4, "[INFO] sEnc : %s", sEnc)
	authorization := fmt.Sprintf("Basic %s", sEnc)

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	if err != nil {
		lprintf(1, "[ERR ] NewRequest : [%s]\n", err)
		return ""
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Language", "en_US")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] Do : [%s]\n", err)
		return ""
	}
	defer respGet.Body.Close()

	resp, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		return ""
	}
	lprintf(4, "[INFO] access-token resp : [%s]\n", string(resp))

	return string(resp)
}

// paypal 승인 데이터
func makePaypalPaymentCreate(data extra.GetPaypalPaymentRequestAllParam) []byte {
	var paypal extra.RequestPaypalPayment
	/*
		var transArray []extra.TransactionInfo
		var amountSubTotal, amountTax float64
		var amountTotal float64

		// intent
		paypal.Intent = "sale"
		// payer
		paypal.Payer.PaymentMethod = "paypal"

		// application_context
		paypal.ApplicationContext.BrandName = "innogs kor"
		paypal.ApplicationContext.Locale = "ko-KR"
		paypal.ApplicationContext.LandingPage = "Login"

		// transactions - array start
		amountSubTotal = 0
		amountTax = 0
		amountTotal = 0

			for _, v := range data.Transactions {
				var trans extra.TransactionInfo
				trans.Description = "pro domain"
				trans.Custom = "leeroom"
				trans.InvoiceNumber = "48787589673"
				// item list
				// shipping address
				trans.ItemList.ShippingAddress.RecipientName = "leeroom"
				trans.ItemList.ShippingAddress.Line1 = "1 main St"
				trans.ItemList.ShippingAddress.Line2 = ""
				trans.ItemList.ShippingAddress.City = "San Jose"
				trans.ItemList.ShippingAddress.CountryCode = "US"
				trans.ItemList.ShippingAddress.PostalCode = "95131"
				trans.ItemList.ShippingAddress.Phone = "011862212345678"
				// items
				for _, w := range v.ItemsList.Items {
					var item extra.ItemInfo
					item.Name = w.Name
					item.Description = w.Description
					item.Quantity = strconv.Itoa(w.Quantity)
					item.Price = strconv.FormatFloat(w.Price, 'f', -1, 32)
					item.Tax = strconv.FormatFloat(w.Tax, 'f', -1, 32)
					item.Sku = w.Sku
					item.Currency = "USD"
					trans.ItemList.Items = append(trans.ItemList.Items, item)
					// amount 전문 만들시 사용
					amountSubTotal += (float64(w.Quantity) * w.Price)
					amountTax += (float64(w.Quantity) * w.Tax)
					amountTotal += amountSubTotal + amountTax
				}
				// amount
				trans.Amount.Total = strconv.FormatFloat(amountTotal, 'f', -1, 32)
				trans.Amount.Currency = "USD"
				trans.Amount.Detail.SubTotal = strconv.FormatFloat(amountSubTotal, 'f', -1, 32)
				trans.Amount.Detail.Tax = strconv.FormatFloat(amountTax, 'f', -1, 32)
				trans.Amount.Detail.Shipping = "0"
				trans.Amount.Detail.HandlingFee = "0"
				trans.Amount.Detail.ShppingDiscount = "0"
				trans.Amount.Detail.Insurance = "0"

				transArray = append(transArray, trans)
			}
			paypal.Transactions = transArray
			// transactions - array end

			// note to payer
			paypal.NoteToPayer = data.NoteToPayer
			// redirect urls
			paypal.RedirectUrls.ReturnUrl = data.ReturnUrl
			paypal.RedirectUrls.CancelUrl = data.CancelUrl
	*/
	// marshal
	payment, err := json.Marshal(&paypal)
	if err != nil {
		lprintf(1, "[ERR ] Marshal : [%s]\n", err)
		return nil
	}

	return payment
}

// 페이팔 결제 데이터
func GetPaypalPaymentRequestParam(c echo.Context) error {
	respData := extra.RespPaypalPayment{}
	var param extra.PaypalPaymentParam

	//테스트 망 url
	//https://www.sandbox.paypal.com/cgi-bin/webscr

	//운영 망 url
	//https://www.paypal.com/cgi-bin/webscr

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

	reqUrl := "https://www.sandbox.paypal.com/cgi-bin/webscr"

	respData.Data.ReqUrl = reqUrl
	respData.Data.Cmd = "_xclick"
	respData.Data.Business = PAYPAL_ID
	//respData.Data.ReturnUrl = ExtraAPI_URL + "/v1/extra/paypal/paymentReturnUrl1"
	if param.PageType == "FE" {
		respData.Data.ReturnUrl = fmt.Sprintf(PAYMENT_REDIRECT_FE+"?manageNo=%s", param.ManageNo)
		//respData.Data.CancelReturn = fmt.Sprintf(ExtraAPI_URL + "/v1/extra/paypal/payment-cancel-return?pageType=%s", param.PageType)
		respData.Data.CancelReturn = fmt.Sprintf(PAYMENT_REDIRECT_FE+"?manageNo=%s", param.ManageNo)
	} else {
		respData.Data.ReturnUrl = PAYMENT_REDIRECT_BO
		respData.Data.CancelReturn = PAYMENT_REDIRECT_BO
	}
	//respData.Data.ReturnUrl = fmt.Sprintf(ExtraAPI_URL + "/v1/extra/paypal/paymentReturnUrl1?manageNo=%s&pageType=%s", param.ManageNo, param.PageType)
	respData.Data.NotifyUrl = fmt.Sprintf(ExtraAPI_URL+"/v1/extra/paypal/payment-return?manageNo=%s", param.ManageNo)
	respData.Data.Charset = "UTF-8"
	respData.Data.CurrencyType = "USD"

	// make response
	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 페이팔 결제 요청
func GetPaypalPaymentRequestAll(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	authToken := extra.GetPaypalAccessTokenOne{}
	respParam := extra.RespGetPaypalPaymentRequestAll{}
	var param extra.GetPaypalPaymentRequestAllParam

	var surl string

	if extra.Mode {
		surl = "https://api.paypal.com/v1/payments/payment"
	} else {
		surl = "https://api.sandbox.paypal.com/v1/payments/payment"
	}

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

	// get access token
	authToken = getAccessTokenfromPaypal()
	if (extra.GetPaypalAccessTokenOne{}) == authToken {
		lprintf(1, "[ERR ] get access-token fail\n")
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] Scope       : [%s]\n", authToken.Scope)
	lprintf(4, "[INFO] AccessToken : [%s]\n", authToken.AccessToken)
	lprintf(4, "[INFO] TokenType   : [%s]\n", authToken.TokenType)
	lprintf(4, "[INFO] AppId       : [%s]\n", authToken.AppId)
	lprintf(4, "[INFO] ExpiresIn   : [%d]\n", authToken.ExpiresIn)
	lprintf(4, "[INFO] Nonce       : [%s]\n", authToken.Nonce)

	// 결제 승인시 필요
	extra.PaypalTokenType = authToken.TokenType
	extra.PaypalAccessToken = authToken.AccessToken

	// paypal 승인 데이터
	payment := makePaypalPaymentCreate(param)
	// lprintf(4, "[INFO] payment : [%s]\n", payment)

	// create payment
	// 참고 : https://developer.paypal.com/docs/api/payments/v1/#payment_create
	//surl := "https://api.sandbox.paypal.com/v1/payments/payment" // 개발
	//surl := "https://api.paypal.com/v1/payments/payment" // 운영
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	// surl := "http://ics.innogs.com"
	// buff := bytes.NewBuffer(resp)
	buff := bytes.NewBuffer(payment)
	lprintf(4, "[INFO] buff : [%s]\n", buff)
	req, err := http.NewRequest("POST", surl, buff)
	if err != nil {
		lprintf(1, "[ERR ] http.NewRequest : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	req.Header.Add("Content-Type", "application/json")
	authorization := fmt.Sprintf("%s %s", authToken.TokenType, authToken.AccessToken)
	req.Header.Add("Authorization", authorization)
	//lprintf(4, "[INFO] req : [%s]\n", req)

	// send
	client := &http.Client{}
	clientDo, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer clientDo.Body.Close()

	// read
	respBody, err := ioutil.ReadAll(clientDo.Body)
	// lprintf(4, "[INFO] clientDo.Body : [%s]\n", clientDo.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody : [%s]\n", respBody)

	if err := json.Unmarshal(respBody, &respParam); err != nil {
		lprintf(1, "[ERR ] json.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respParam : [%s]\n", respParam)

	// make response
	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respParam.Links

	return c.JSON(http.StatusOK, respData)
}

func GetPaypalPaymentReturn(c echo.Context) error {
	param := make(map[string]string)
	var returnParam extra.RespGetAlipayReturn
	manageNo := c.QueryParam("manageNo")

	lprintf(4, "[INFO] paypal payment return url(%s)\n", c.Request().URL.String())

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusOK, nil)
	}

	lprintf(4, "[INFO] paypal payment return data(%s) \n", string(resp))
	tmp := strings.Split(strings.TrimSpace(string(resp)), "&")
	for _, v := range tmp {
		tmps := strings.Split(v, "=")
		param[tmps[0]] = tmps[1]
	}

	// 결제 성공일 때
	if param["payment_status"] == "Completed" {
		returnParam.OrderId = manageNo
		returnParam.TradeNumber = param["txn_id"]
		sendPaymentInfo(returnParam)
	} else {
		// 결제 실패시
	}

	return c.HTML(http.StatusOK, "")
}

func GetPaypalPaymentReturnOne4(c echo.Context) error {

	pageType := c.QueryParam("pageType")

	var returnUrl string

	if pageType == "FE" {
		returnUrl = PAYMENT_REDIRECT_FE
	} else {
		returnUrl = PAYMENT_REDIRECT_BO
	}

	c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리
	return c.HTML(http.StatusMovedPermanently, "")
}

func PostRedirectGet(c echo.Context) error {

	manageNo := c.QueryParam("manageNo")

	var returnUrl string
	if len(manageNo) == 0 {
		returnUrl = PAYMENT_REDIRECT_FE + "?manageNo=PP-20201223184206518"
	} else {
		returnUrl = fmt.Sprintf(PAYMENT_REDIRECT_FE+"?manageNo=%s", manageNo)
	}

	lprintf(4, "[INFO] PaymentRedirect : [%s]\n", returnUrl)

	return c.HTML(http.StatusOK, "OK")

	resp, err := http.Get(returnUrl)
	if err != nil {
		lprintf(1, "[ERR ] http(%s) get err(%s) \n", returnUrl, err.Error())
		c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리
		return c.HTML(http.StatusMovedPermanently, "")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lprintf(1, "[ERR ] http(%s) read err(%s) \n", returnUrl, err.Error())
		c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리
		return c.HTML(http.StatusMovedPermanently, "")
	}

	lprintf(4, "[INFO] resp html(%s) \n", string(data))

	return c.HTMLBlob(http.StatusOK, data)
}

// 페이팔 결제 승인 요청
func GetPaypalPaymentReturnOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var respJson extra.RespGetPaypalPaymentReturnOne
	// 참고 : https://developer.paypal.com/docs/checkout/reference/upgrade-integration/#billing-agreement-integrations-with-v1billing-agreementsagreement-tokens-1
	//https://extraapi.securitynetsvc.com:15001/v1/extra/paypal/payment-return?paymentId=PAYID-L5QYQQQ2P344273EG471680B&token=EC-4KD044345X910652B&PayerID=NDY8MGYPPRY9G

	paymentId := c.QueryParam("paymentId")
	token := c.QueryParam("token")
	lprintf(4, "[INFO] paymentId : [%s]\n", paymentId)
	lprintf(4, "[INFO] token : [%s]\n", token)

	var surl string

	if extra.Mode {
		surl = "https://api.paypal.com/v1/billing-agreements/agreements"
	} else {
		surl = "https://api.sandbox.paypal.com/v1/billing-agreements/agreements"
	}

	// finalize payment
	// surl := "https://api.sandbox.paypal.com/v1/billing-agreements/agreements" // 개발
	//surl := "https://api.paypal.com/v1/billing-agreements/agreements" // 운영
	// surl := "http://ics.innogs.com"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	tokenStruct := map[string]string{"billingToken": token}
	tokenJson, err := json.Marshal(tokenStruct)
	if err != nil {
		lprintf(1, "[ERR ] tokenJson : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
	}
	lprintf(4, "[INFO] tokenJson : [%s]\n", tokenJson)
	buff := bytes.NewBuffer(tokenJson)

	// send
	req, err := http.NewRequest("POST", surl, buff)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
	}
	req.Header.Add("Content-Type", "application/json")
	authorization := fmt.Sprintf("%s %s", extra.PaypalTokenType, extra.PaypalAccessToken)
	lprintf(4, "[INFO] authorization : [%s]\n", authorization)
	req.Header.Add("Authorization", authorization)

	// send
	client := &http.Client{}
	clientDo, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer clientDo.Body.Close()

	// read
	respBody, err := ioutil.ReadAll(clientDo.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody : [%s]\n", respBody)
	respJson.Data = string(respBody)

	//extra.PaymentRedirect = fmt.Sprintf("%s/%s", extra.PaymentRedirect, respJson.Data)
	lprintf(4, "[INFO] paypal redirect url(%s) \n", extra.PaymentRedirect)
	c.Response().Header().Set(echo.HeaderLocation, extra.PaymentRedirect) // 정상 처리 후 redirect 처리

	return c.HTML(http.StatusMovedPermanently, "")

	// make response
	// respData.Code = extra.SUCCESS
	//
	// respData.ServiceName = extra.TYPE
	// respData.Data = &respJson

	// return c.JSON(http.StatusOK, respData)
}

// 페이팔 결체 취소 요청
func DelPaypalRefundRequestOne(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.DelPaypalRefundRequestOneParam
	var respJson extra.RespDelPaypalRefundRequestOne

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

	var surl string

	if extra.Mode {
		surl = fmt.Sprintf("https://api.paypal.com/v1/payments/refund/%s", param.RefundId)
	} else {
		surl = fmt.Sprintf("https://api.sandbox.paypal.com/v1/payments/refund/%s", param.RefundId)
	}

	// finalize payment
	// surl := fmt.Sprintf("https://api.sandbox.paypal.com/v1/payments/refund/%s", param.RefundId) // 개발
	//surl := fmt.Sprintf("https://api.paypal.com/v1/payments/refund/%s", param.RefundId) // 운영
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	req, err := http.NewRequest("GET", surl, nil)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
	}
	req.Header.Add("Content-Type", "application/json")

	// test
	// extra.PaypalTokenType = "Bearer"
	// extra.PaypalAccessToken = "A21AAPYU6pTlhckxuT8N4aMzr8exF_QyKDwmjcy95TIUNZ1s2E2kkAOt3T-nrXVpFN63hs3fXGYFgGH6sHErNiVGeFKn0yFJw"
	authorization := fmt.Sprintf("%s %s", param.PaypalTokenType, param.PaypalAccessToken)
	lprintf(4, "[INFO] authorization : [%s]\n", authorization)
	req.Header.Add("Authorization", authorization)

	// send
	client := &http.Client{}
	clientDo, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer clientDo.Body.Close()

	// read
	respBody, err := ioutil.ReadAll(clientDo.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respBody : [%s]\n", respBody)
	respJson.Data = string(respBody)

	// make response
	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// 페이팔 결제 취소 승인 요청
func GetPaypalPaymentCancelOne(c echo.Context) error {
	// https://extraapi.securitynetsvc.com:15001/v1/extra/paypal/payment-cancel?token=EC-8RR2828827667003R
	respData := extra.ApiCallResponse{}
	// var respJson extra.RespGetPaypalPaymentCancelOne

	token := c.QueryParam("token")
	lprintf(4, "[INFO] token : [%s]\n", token)

	// surl := "https://api.paypal.com/v1/billing-agreements/agreements" // 운영

	// // send
	// req, err := http.NewRequest("POST", surl, nil)
	// if err != nil {
	//     lprintf(1, "[ERR ] client.Do : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// }
	// req.Header.Add("Content-Type", "application/json")

	// // 취소
	// // test
	// extra.PaypalTokenType = "Bearer"
	// extra.PaypalAccessToken = "A21AAOAYqzTmtbVhKltWMisAM4_NYPf_J7Kh9bP2Ly7WBzcHNm4H9WPu652zDYPDPGMycRAIZb7LpCiXW08TsO9famr9OpeC"

	// authorization := fmt.Sprintf("%s %s", extra.PaypalTokenType, extra.PaypalAccessToken)
	// lprintf(4, "[INFO] authorization : [%s]\n", authorization)
	// req.Header.Add("Authorization", authorization)

	// // send
	// client := &http.Client{}
	// clientDo, err := client.Do(req)
	// if err != nil {
	// 	lprintf(1, "[ERR ] client.Do : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// defer clientDo.Body.Close()

	// // read
	// respBody, err := ioutil.ReadAll(clientDo.Body)
	// if err != nil {
	// 	lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	// lprintf(4, "[INFO] respBody : [%s]\n", respBody)
	// respJson.Data = string(respBody)

	// // make response
	// respData.Code = extra.SUCCESS
	//
	// respData.ServiceName = extra.TYPE
	// respData.Data = &respJson

	return c.JSON(http.StatusOK, respData)
}

// paypal 환불 요청
func GetPaypalRefundRequest(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.PaypalRefundParam
	paypalResp := make(map[string]string)

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

	txnId := param.PgApprovalNo

	data := url.Values{}
	data.Set("USER", "innogs_api1.innogs.com")
	data.Set("PWD", "87Z4K77DQPCY3DYN")
	data.Set("SIGNATURE", "ANPqaIIjJFj9K9rndoPcyjmO.Nd-Af.wHwv-r2pX8oYR3y3DyrlpFtAB")
	data.Set("METHOD", "RefundTransaction")
	data.Set("VERSION", "94")
	data.Set("TRANSACTIONID", txnId)
	data.Set("REFUNDTYPE", "Full")

	lprintf(4, "[INFO] data : [%s]\n", data)

	surl := PAYPAL_REFUND_URL
	//운영망 surl := "https://api-3t.paypal.com/nvp"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	respGet, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer respGet.Body.Close()

	respBody, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] paypal respBody : [%s]\n", respBody)

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] xml.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[DH 22 body] %s \n", string(resp))
	tmp := strings.Split(strings.TrimSpace(string(resp)), "&")
	for _, v := range tmp {
		tmps := strings.Split(v, "=")
		paypalResp[tmps[0]] = tmps[1]
	}

	// 환불 실패일 때
	if paypalResp["ACK"] != "Success" || paypalResp["ACK"] != "SuccessWithWarning" {
		respData.Code = extra.FAIL

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	return c.JSON(http.StatusOK, respData)
}
