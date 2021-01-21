package controller

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	extra "hydrawebapi/extraapi/model"

	"github.com/labstack/echo"
)

/*
AliPay결제 모듈 호출
Form 데이터의 sign, sign_type을 제외한 값을통해 sign을 생성 후
Form 데이터와 함께 alipay 전송

alipay 개발 서버 url
https://mapi.alipaydev.com/gateway.do?

alipay 운영 서버 url
https://mapi.alipay.com/gateway.do?

환불시 필요 데이터
alipayParam.Add("service", "forex_refund");
alipayParam.Add("partner", this.Partner);
alipayParam.Add("_input_charset", this.InputCharSet);
alipayParam.Add("out_return_no", this.OutTradeNo);
alipayParam.Add("out_trade_no", this.OutTradeNo);
alipayParam.Add("return_amount", this.ReturnAmount.ToString());
alipayParam.Add("gmt_return", TimeZoneInfo.ConvertTimeBySystemTimeZoneId(DateTime.UtcNow, "China Standard Time").ToString("YYYYMMDDHHMMSS"));
alipayParam.Add("currency", this.Currency);

*/

// func GetAlipayCancle() string {

// 	partner := "2088621930006625" // config
// 	orderId := "nJhn6WXaszbmxj3SebsB1ooqx0d7Le" // 화면

// 	param := &extra.AlipayParameters{}
// 	param.Service = "forex_refund"
// 	param.Partner = partner
// 	param.InputCharset = "utf-8"
// 	param.OutTradeNo = orderId
// 	param.OutReturnNo = orderId
// 	param.ReturnAmount = "87.27" // 환불 가격 - 화면
// 	loc, _ := time.LoadLocation("Asia/Shanghai") // 고정
// 	param.GmtReturn = time.Now().In(loc).Format("20060102150405")
// 	param.Currency = "USD"

// 	key := "c1pypzgdv8ln51iytyehtnv6t8esgk4c" // config

// 	param.Sign = sign(param, key)
// 	param.SignType = "MD5"

// 	data := url.Values{}
// 	data.Set("service", param.Service)
// 	data.Set("partner", param.Partner)
// 	data.Set("_input_charset", param.InputCharset)
// 	data.Set("out_return_no", param.OutReturnNo)
// 	data.Set("out_trade_no", param.OutTradeNo)
// 	data.Set("return_amount", param.ReturnAmount)
// 	data.Set("gmt_return", param.GmtReturn)
// 	data.Set("currency", param.Currency)
// 	data.Set("sign", param.Sign)
// 	data.Set("sign_type", param.SignType)

// 	req, err := http.NewRequest("POST", "https://mapi.alipaydev.com/gateway.do?_input_charset=utf-8", strings.NewReader(data.Encode()))
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	client := &http.Client{}
// 	respGet, err := client.Do(req)
// 	if err != nil {
// 		return ""
// 	}
// 	defer respGet.Body.Close()

// 	respBody, err := ioutil.ReadAll(respGet.Body)
// 	if err != nil {
// 		return ""
// 	}

// 	return string(respBody)

// }

func GetAlipayPaymentTest(c echo.Context) error {

	// nicName := "enough Name Servers by each Domain Name and a differentiated DNS service that guarantees Domain Name security with A-Alias that compensated CNAME's security issue." // 고정
	// subject := "DNS Professional" // 화면
	// partner := "2088621930006625" // config
	// notifyurl := "http://220.95.208.178/commerce/json/payment/alipay/callback"
	// // notifyurl := ExtraAPI_URL + "/commerce/json/payment/alipay/callback"
	// orderId := "44IhHnmFzko5BzkOFuS53ZGU8bAlHt" // 자체 생성으로 추측 -  화면
	// // orderId := "m)segoH8Vj94unB6w5VkzjdcPRvEze"

	// param := &extra.AlipayParameters{}
	// param.Service = "create_forex_trade" // forex_refund (환불 시) - 고정
	// param.Partner = "2088621930006625"
	// param.InputCharset = "utf-8" // 고정
	// param.NotifyUrl = "http://220.95.208.178/commerce/json/payment/alipay/callback"
	// param.ReturnUrl = "http://localhost:51584/PaymentComplete"
	// param.ReferUrl = "http://localhost:51584"
	// param.OutTradeNo = "44IhHnmFzko5BzkOFuS53ZGU8bAlHt"
	// param.Subject = "DNS Professional"
	// param.Currency = "USD"   // 고정
	// param.TotalFee = "87.27" // 화면
	// param.Body = "enough Name Servers by each Domain Name and a differentiated DNS service that guarantees Domain Name security with A-Alias that compensated CNAME's security issue."
	// param.ProductCode = "NEW_OVERSEAS_SELLER" // 결제 요청 시 고정값

	// key := "c1pypzgdv8ln51iytyehtnv6t8esgk4c" // config
	// param.Sign = sign(param, key)
	// param.SignType = "MD5"

	// barCode := `
	// 	<form id="alipaysubmit" name="alipaysubmit" action="https://mapi.alipaydev.com/gateway.do?" method="GET" style='display:none;'>
	// 		<input type="hidden" name="service" value="` + param.Service + `">
	// 		<input type="hidden" name="partner" value="` + param.Partner + `">
	// 		<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
	//         <input type="hidden" name="sign_type" value="` + param.SignType + `">
	//         <input type="hidden" name="sign" value="` + param.Sign + `">
	// 		<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
	// 		<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
	//         <input type="hidden" name="refer_url" value="` + param.ReferUrl + `">
	// 		<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
	// 		<input type="hidden" name="subject" value="` + param.Subject + `">
	//         <input type="hidden" name="currency" value="` + param.Currency + `">
	// 		<input type="hidden" name="total_fee" value="` + param.TotalFee + `">
	// 		<input type="hidden" name="body" value="` + param.Body + `">
	//         <input type="hidden" name="product_code" value="` + param.ProductCode + `">
	// 	</form>
	// 	<script>
	// 		document.forms['alipaysubmit'].submit();
	// 	</script>
	// `

	// return c.HTML(http.StatusOK, barCode)

	//respData := extra.RespGetAlipayPaymentRequestAll{}
	var param extra.AlipayParameters

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
	}

	param.Service = "create_forex_trade" // forex_refund (환불 시) - 고정
	// param.Partner = "2088621930006625" // config
	param.Partner = ALIPAY_PARTNER_ID // config
	param.InputCharset = "utf-8"      // 고정
	// param.Currency = "USD" // 화면에서 입력
	param.Body = "enough Name Servers by each Domain Name and a differentiated DNS service that guarantees Domain Name security with A-Alias that compensated CNAME's security issue." // 고정
	param.ProductCode = "NEW_OVERSEAS_SELLER"                                                                                                                                          // 결제 요청 시 - 고정값
	param.NotifyUrl = ExtraAPI_URL + "/v1/extra/alipay/callback"
	// param.ReturnUrl = fmt.Sprintf("%s/v1/extra/alipay/return?pageType=%s", extra.PaymentReturn, param.PageType)
	//param.ReturnUrl = fmt.Sprintf("%s/v1/extra/alipay/return", extra.PaymentReturn)
	param.ReferUrl = "https://extraalpi.securitynetsvc.com:15001/v1/extra/alipay/refer"
	param.TotalFee = strconv.FormatFloat(param.Total, 'f', -1, 32)
	key := extra.AliPayKey // config
	param.OutTradeNo = "ex1"
	param.Sign = sign(&param, key)
	param.SignType = "MD5"

	// 결제 요청 form 데이터

	barCode := `<form id="alipaysubmit" name="alipaysubmit" action="` + extra.AliOperUrl + `?" method="GET" style='display:none;'>
	<input type="hidden" name="service" value="` + param.Service + `">
	<input type="hidden" name="partner" value="` + param.Partner + `">
	<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
	<input type="hidden" name="sign_type" value="` + param.SignType + `">
	<input type="hidden" name="sign" value="` + param.Sign + `">
	<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
	<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
	<input type="hidden" name="refer_url" value="` + param.ReferUrl + `">
	<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
	<input type="hidden" name="subject" value="` + param.Subject + `">
	<input type="hidden" name="currency" value="` + param.Currency + `">
	<input type="hidden" name="total_fee" value="` + param.TotalFee + `">
	<input type="hidden" name="body" value="` + param.Body + `">
	<input type="hidden" name="product_code" value="` + param.ProductCode + `">
</form>
<!--script>document.forms['alipaysubmit'].submit();</script-->`

	barCode = strings.Replace(barCode, "\t", "", -1)
	barCode = strings.Replace(barCode, "\n", "<br>", -1)

	// barCode := `<form id="alipaysubmit" name="alipaysubmit" action="` + extra.AliOperUrl + `?" method="GET" style='display:none;'>
	// 		<input type="hidden" name="service" value="` + param.Service + `">
	// 		<input type="hidden" name="partner" value="` + param.Partner + `">
	// 		<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
	//            <input type="hidden" name="sign_type" value="` + param.SignType + `">
	//            <input type="hidden" name="sign" value="` + param.Sign + `">
	// 		<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
	// 		<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
	//            <input type="hidden" name="refer_url" value="` + param.ReferUrl + `">
	// 		<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
	// 		<input type="hidden" name="subject" value="` + param.Subject + `">
	//            <input type="hidden" name="currency" value="` + param.Currency + `">
	// 		<input type="hidden" name="total_fee" value="` + param.TotalFee + `">
	// 		<input type="hidden" name="body" value="` + param.Body + `">
	//            <input type="hidden" name="product_code" value="` + param.ProductCode + `">
	// 	</form>
	// 	<script>document.forms['alipaysubmit'].submit();</script>`

	lprintf(4, "[INFO] after replace alipay payment : [%s]\n", barCode)

	return c.HTML(http.StatusOK, barCode)

}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

// Alipay 결제 요청
func GetAlipayPaymentRequestAll(c echo.Context) error {
	respData := extra.RespGetAlipayPaymentRequestAll{}
	var param extra.AlipayParameters

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
	/*
		param.Service = "create_forex_trade" // forex_refund (환불 시) - 고정
		// param.Partner = "2088621930006625" // config
		param.Partner = extra.AliPayPartnerId // config
		param.InputCharset = "utf-8"          // 고정
		// param.Currency = "USD" // 화면에서 입력
		param.Body = "enough Name Servers by each Domain Name and a differentiated DNS service that guarantees Domain Name security with A-Alias that compensated CNAME's security issue." // 고정
		param.ProductCode = "NEW_OVERSEAS_SELLER"                                                                                                                                          // 결제 요청 시 - 고정값
		param.NotifyUrl = ExtraAPI_URL + "/v1/extra/alipay/callback"
		param.ReturnUrl = fmt.Sprintf("%s/v1/extra/alipay/return?pageType=%s", extra.PaymentReturn, param.PageType)
		param.ReferUrl = "https://extraalpi.securitynetsvc.com:15001/v1/extra/alipay/refer"
		param.TotalFee = strconv.FormatFloat(param.Total, 'f', -1, 32)
		key := extra.AliPayKey // config
		param.Sign = sign(&param, key)
		param.SignType = "MD5"
	*/

	param.Service = "create_forex_trade" // forex_refund (환불 시) - 고정
	// param.Partner = "2088621930006625" // config
	param.Partner = ALIPAY_PARTNER_ID                                                                                                                                                  // config
	param.InputCharset = "utf-8"                                                                                                                                                       // 고정
	param.Body = "enough Name Servers by each Domain Name and a differentiated DNS service that guarantees Domain Name security with A-Alias that compensated CNAME's security issue." // 고정
	param.ProductCode = "NEW_OVERSEAS_SELLER"                                                                                                                                          // 결제 요청 시 - 고정값
	param.NotifyUrl = ExtraAPI_URL + "/v1/extra/alipay/callback"
	param.ReturnUrl = fmt.Sprintf("%s/v1/extra/alipay/return?pageType=%s", ExtraAPI_URL, param.PageType)
	// param.ReturnUrl = fmt.Sprintf("%s/v1/extra/alipay/return", extra.PaymentReturn)
	param.ReferUrl = "https://extraalpi.securitynetsvc.com:15001/v1/extra/alipay/refer"
	param.TotalFee = strconv.FormatFloat(param.Total, 'f', -1, 32)
	/*
		respData.Data.Service = "create_forex_trade" // forex_refund (환불 시) - 고정
		// param.Partner = "2088621930006625" // config
		respData.Data.Partner = extra.AliPayPartnerId // config
		respData.Data.InputCharset = "utf-8"          // 고정
		respData.Data.OutTradeNo = param.OutTradeNo
		// param.Currency = "USD" // 화면에서 입력
		respData.Data.Body = "enough Name Servers by each Domain Name and a differentiated DNS service that guarantees Domain Name security with A-Alias that compensated CNAME's security issue." // 고정
		respData.Data.ProductCode = "NEW_OVERSEAS_SELLER"                                                                                                                                          // 결제 요청 시 - 고정값
		respData.Data.NotifyUrl = ExtraAPI_URL + "/v1/extra/alipay/callback"
		respData.Data.ReturnUrl = fmt.Sprintf("%s/v1/extra/alipay/return?pageType=%s", extra.PaymentReturn, param.PageType)
		respData.Data.ReferUrl = "https://extraalpi.securitynetsvc.com:15001/v1/extra/alipay/refer"
		respData.Data.TotalFee = strconv.FormatFloat(param.Total, 'f', -1, 32)
	*/

	respData.Data.Service = param.Service // forex_refund (환불 시) - 고정
	// param.Partner = "2088621930006625" // config
	respData.Data.Partner = ALIPAY_PARTNER_ID // config
	respData.Data.InputCharset = "utf-8"      // 고정
	respData.Data.OutTradeNo = param.OutTradeNo
	respData.Data.Body = param.Body
	respData.Data.ProductCode = param.ProductCode // 결제 요청 시 - 고정값
	respData.Data.NotifyUrl = param.NotifyUrl
	respData.Data.ReturnUrl = param.ReturnUrl
	respData.Data.ReferUrl = param.ReferUrl
	respData.Data.TotalFee = param.TotalFee

	key := ALIPAY_KEY // config
	respData.Data.Sign = sign(param, key)
	respData.Data.SignType = "MD5"

	// barCode = strings.Replace(barCode, "\t", "", -1)
	// barCode = strings.Replace(barCode, "\n", "<br>", -1)
	/*
		// 결제 요청 form 데이터

		barCode := `<form id="alipaysubmit" name="alipaysubmit" action="` + extra.AliOperUrl + `?" method="GET" style='display:none;'>
							<input type="hidden" name="service" value="` + param.Service + `">
							<input type="hidden" name="partner" value="` + param.Partner + `">
							<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
							<input type="hidden" name="sign_type" value="` + param.SignType + `">
							<input type="hidden" name="sign" value="` + param.Sign + `">
							<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
							<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
							<input type="hidden" name="refer_url" value="` + param.ReferUrl + `">
							<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
							<input type="hidden" name="subject" value="` + param.Subject + `">
							<input type="hidden" name="currency" value="` + param.Currency + `">
							<input type="hidden" name="total_fee" value="` + param.TotalFee + `">
							<input type="hidden" name="body" value="` + param.Body + `">
							<input type="hidden" name="product_code" value="` + param.ProductCode + `">
						</form>
						<!--script>document.forms['alipaysubmit'].submit();</script-->`

		barCode = strings.Replace(barCode, "\t", "", -1)
		barCode = strings.Replace(barCode, "\n", "<br>", -1)

		// barCode := `<form id="alipaysubmit" name="alipaysubmit" action="` + extra.AliOperUrl + `?" method="GET" style='display:none;'>
		// 		<input type="hidden" name="service" value="` + param.Service + `">
		// 		<input type="hidden" name="partner" value="` + param.Partner + `">
		// 		<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
		//            <input type="hidden" name="sign_type" value="` + param.SignType + `">
		//            <input type="hidden" name="sign" value="` + param.Sign + `">
		// 		<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
		// 		<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
		//            <input type="hidden" name="refer_url" value="` + param.ReferUrl + `">
		// 		<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
		// 		<input type="hidden" name="subject" value="` + param.Subject + `">
		//            <input type="hidden" name="currency" value="` + param.Currency + `">
		// 		<input type="hidden" name="total_fee" value="` + param.TotalFee + `">
		// 		<input type="hidden" name="body" value="` + param.Body + `">
		//            <input type="hidden" name="product_code" value="` + param.ProductCode + `">
		// 	</form>
		// 	<script>document.forms['alipaysubmit'].submit();</script>`

		lprintf(4, "[INFO] after replace alipay payment : [%s]\n", barCode)
	*/

	// make response data
	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE

	return c.JSON(http.StatusOK, respData)
	// return c.HTML(http.StatusOK, barCode)
}

// Alipay 결제 승인 결과 (화면에서 바코드로 결제 했을때)
func GetAlipayReturn(c echo.Context) error {
	var respJson extra.RespGetAlipayReturn
	var pageType, returnUrl string

	// https://global.alipay.com/docs/ac/gr/trade_status
	// /v1/extra/alipay/return?out_trade_no=IVN7WZfjGD)8x2qH4z01R4I1BdwQdt&total_fee=87.27&trade_status=TRADE_FINISHED&sign=13840fd7788a35bb95c0bd039115f2a9
	// &trade_no=2020091722001306350504519291&currency=USD&sign_type=MD5
	pageType = c.QueryParam("pageType")
	respJson.OrderId = c.QueryParam("out_trade_no")
	respJson.Total, _ = strconv.ParseFloat(c.QueryParam("total_fee"), 64)
	respJson.Status = c.QueryParam("trade_status")
	respJson.Sign = c.QueryParam("sign")
	respJson.TradeNumber = c.QueryParam("trade_no")
	respJson.Currency = c.QueryParam("currency")
	respJson.SignType = c.QueryParam("sign_type")
	respJson.ReturnUrl = c.QueryParam("return_url")

	lprintf(4, "[INFO] orderId  : [%s]\n", respJson.OrderId)
	//orderId := respJson.OrderId
	lprintf(4, "[INFO] total    : [%g]\n", respJson.Total)
	lprintf(4, "[INFO] status   : [%s]\n", respJson.Status)
	lprintf(4, "[INFO] sign     : [%s]\n", respJson.Sign)
	lprintf(4, "[INFO] tradeNo  : [%s]\n", respJson.TradeNumber)
	lprintf(4, "[INFO] currnecy : [%s]\n", respJson.Currency)
	lprintf(4, "[INFO] signType : [%s]\n", respJson.SignType)
	lprintf(4, "[INFO] returnUrl : [%s]\n", respJson.ReturnUrl)

	//go sendPaymentInfo(respJson) // billing api 호출
	if respJson.Status == "TRADE_FINISHED" {
		// 결제 성공
		respJson.Status = "complete"
		sendPaymentInfo(respJson)
	}

	//redirect := fmt.Sprintf("%s/%s", extra.PaymentRedirect, orderId)

	if pageType == "FE" {
		returnUrl = fmt.Sprintf(PAYMENT_REDIRECT_FE+"?manageNo=%s", respJson.OrderId)
	} else {
		returnUrl = PAYMENT_REDIRECT_BO
	}
	lprintf(4, "[INFO] PaymentRedirect : [%s]\n", returnUrl)
	c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리

	// respHtml := `<html><script text="javascript/text">init();function init() {window.close();}</script></html>`
	// return c.HTML(http.StatusOK, respHtml)

	lprintf(4, "[INFO] complete alipay payment\n")
	//time.Sleep(1 * time.Millisecond)
	return c.HTML(http.StatusMovedPermanently, "")
}

func sendPaymentInfo(data extra.RespGetAlipayReturn) {
	lprintf(4, "[INFO] sendPaymentInfo start\n")
	var apiCallResponse extra.ApiCallResponse
	reqBytes, err := json.Marshal(&data)
	if err != nil {
		lprintf(1, "[ERR ] json.Marshar : [%s]\n", err)
		return
	}

	buff := bytes.NewBuffer(reqBytes)
	surl := BillingAPI_URL + "/v1/billing/payment/alipay/return" // 담당자 전지연씨와 협의하여 billing api 에 명시된 url로 변경 필요
	req, err := http.NewRequest("POST", surl, buff)
	// req.Header.Add("Content-Type", "application/json")
	lprintf(4, "[INFO] surl : [%s]\n", surl)

	// send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		return
	}

	// check response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		return
	}
	lprintf(4, "[INFO] respBody : [%s]\n", respBody)

	if err := json.Unmarshal(respBody, &apiCallResponse); err != nil {
		lprintf(1, "[ERR ] json.Unmarshal : [%s]\n", err)
		return
	}

	if apiCallResponse.Code != "0000" {
		lprintf(1, "[ERR ] api call response code : [%s]\n", apiCallResponse.Code)
		return
	}

	return
}

func GetAlipayCallback(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	return c.JSON(http.StatusOK, respData)
}

func GetAlipayRefer(c echo.Context) error {
	lprintf(4, "[INFO] refer : [%s]\n", c.Request())
	return c.HTML(http.StatusOK, "")
}

// Alipay 결제 취소 요청
func GetAlipayPaymentRefundOne(c echo.Context) error {
	//https://global.alipay.com/docs/ac/hkapi/refund.query
	respData := extra.ApiCallResponse{}
	//var param extra.AlipayParameters
	var param extra.PaymentCancelOneParam
	var refundParam extra.AlipayRefundParam
	var respXml extra.RespGetAlipayPaymentRefundOne

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

	// param.Service = "forex_refund"
	// param.Partner = extra.AliPayPartnerId
	// param.InputCharset = "utf-8" // 고정
	// // param.OutTradeNo = param.OrderId // 화면에서 입력
	// param.OutReturnNo = param.
	// param.ReturnAmount = strconv.FormatFloat(param.Total, 'f', -1, 32) // 환불 가격 - 화면
	// loc, _ := time.LoadLocation("Asia/Shanghai")                       // 고정
	// param.GmtReturn = time.Now().In(loc).Format("20060102150405")
	// // param.Currency = "USD" // 화면에서 입력

	// key := extra.AliPayKey // config
	// param.Sign = sign(&param, key)
	// param.SignType = "MD5" // 고정
	// lprintf(4, "[INFO] sign : [%s]\n", param.Sign)

	// service := "forex_refund"
	// partner := extra.AliPayPartnerId
	// inputCharset := "utf-8"          // 고정
	// outTradeNo := param.PgApprovalNo // 화면에서 입력
	// outReturnNo := param.ManageNo
	// returnAmount := strconv.FormatFloat(param.Price, 'f', -1, 64) // 환불 가격 - 화면
	// loc, _ := time.LoadLocation("Asia/Shanghai")                  // 고정
	// gmtReturn := time.Now().In(loc).Format("20060102150405")
	// currency := param.CurrencyType

	// key := extra.AliPayKey // config
	// sign := refundSign(&param, key)
	// signType := "MD5" // 고정
	// lprintf(4, "[INFO] sign : [%s]\n", sign)

	// data := url.Values{}
	// data.Set("service", service)
	// data.Set("partner", partner)
	// data.Set("_input_charset", inputCharset)
	// data.Set("out_return_no", outReturnNo)
	// data.Set("out_trade_no", outTradeNo)
	// data.Set("return_amount", returnAmount)
	// data.Set("gmt_return", gmtReturn)
	// data.Set("currency", currency)
	// data.Set("sign", sign)
	// data.Set("sign_type", signType)
	refundParam.Service = "forex_refund"
	refundParam.Partner = ALIPAY_PARTNER_ID
	refundParam.InputCharset = "utf-8"      // 고정
	refundParam.OutTradeNo = param.ManageNo // 화면에서 입력
	refundParam.OutReturnNo = param.ManageNo
	refundParam.ReturnAmount = strconv.FormatFloat(param.Price, 'f', -1, 64) // 환불 가격 - 화면
	loc, _ := time.LoadLocation("Asia/Shanghai")                             // 고정
	refundParam.GmtReturn = time.Now().In(loc).Format("20060102150405")
	refundParam.Currency = param.CurrencyType

	key := ALIPAY_KEY // config
	sign := refundSign(&refundParam, key)
	signType := "MD5" // 고정
	lprintf(4, "[INFO] sign : [%s]\n", sign)

	data := url.Values{}
	data.Set("service", refundParam.Service)
	data.Set("partner", refundParam.Partner)
	data.Set("_input_charset", refundParam.InputCharset)
	data.Set("out_return_no", refundParam.OutReturnNo)
	data.Set("out_trade_no", refundParam.OutTradeNo)
	data.Set("return_amount", refundParam.ReturnAmount)
	data.Set("gmt_return", refundParam.GmtReturn)
	data.Set("currency", refundParam.Currency)
	data.Set("sign", sign)
	data.Set("sign_type", signType)

	lprintf(4, "[INFO] data : [%s]\n", data)

	surl := fmt.Sprintf("%s?_input_charset=utf-8", ALIPAY_OPER_URL)
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
	lprintf(4, "[INFO] alipay respBody : [%s]\n", respBody)

	if err := xml.Unmarshal(respBody, &respXml); err != nil {
		lprintf(1, "[ERR ] xml.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] IsSuccess : [%s], Error : [%s]\n", respXml.IsSuccess, respXml.Error) // T : 정상, F : 비정상
	result := respXml.IsSuccess
	if result != "T" {
		lprintf(1, "[ERR ] Alipay Refund FAIL : [%s]\n", respXml.Error)
		respData.Code = extra.FAIL
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = extra.SUCCESS
	respData.ServiceName = extra.TYPE
	respData.Data = &respXml

	return c.JSON(http.StatusOK, respData)
}

/*
alipay sign 방법
json을 byte로 변한 한 값에서 key=value& 형태의 url을 생성 후 마지막에 parter private key를 합쳐 md5 sign 값 제작
*/
func refundSign(param interface{}, key string) string {

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return ""
	}

	var sign string
	var signs []string
	oldString := string(paramBytes)

	oldString = strings.Replace(oldString, `\u003c`, "<", -1)
	oldString = strings.Replace(oldString, `\u003e`, ">", -1)

	oldString = strings.Replace(oldString, "\"", "", -1)
	oldString = strings.Replace(oldString, "{", "", -1)
	oldString = strings.Replace(oldString, "}", "", -1)
	paramArray := strings.Split(oldString, ",")

	for _, v := range paramArray {
		detail := strings.SplitN(v, ":", 2)
		lprintf(4, "[INFO] detail : [%s], \n", detail)
		if len(detail[1]) > 0 && detail[0] != "sign" && detail[0] != "sign_type" && detail[0] != "total" && detail[0] != "pageType" { // alipay 키가 아닌 것을 추가해서 제거
			signs = append(signs, detail[0]+"="+detail[1])
		}
	}

	sort.Strings(signs)

	for i := 0; i < len(signs); i++ {
		sign += signs[i] + "&"
	}
	//lprintf(4, "[INFO] sign : [%s]\n", sign)

	m := md5.New()
	m.Write([]byte(sign[:len(sign)-1] + key))

	return hex.EncodeToString(m.Sum(nil))

}
func sign(param interface{}, key string) string {

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return ""
	}

	var sign string
	var signs []string
	oldString := string(paramBytes)

	oldString = strings.Replace(oldString, `\u003c`, "<", -1)
	oldString = strings.Replace(oldString, `\u003e`, ">", -1)

	oldString = strings.Replace(oldString, "\"", "", -1)
	oldString = strings.Replace(oldString, "{", "", -1)
	oldString = strings.Replace(oldString, "}", "", -1)
	paramArray := strings.Split(oldString, ",")

	for _, v := range paramArray {
		detail := strings.SplitN(v, ":", 2)
		lprintf(4, "[INFO] detail : [%s], \n", detail)
		if len(detail[1]) > 0 && detail[0] != "sign" && detail[0] != "sign_type" && detail[0] != "total" && detail[0] != "pageType" { // alipay 키가 아닌 것을 추가해서 제거
			signs = append(signs, detail[0]+"="+detail[1])
		}
	}

	sort.Strings(signs)

	for i := 0; i < len(signs); i++ {
		sign += signs[i] + "&"
	}
	//lprintf(4, "[INFO] sign : [%s]\n", sign)

	m := md5.New()
	m.Write([]byte(sign[:len(sign)-1] + key))

	return hex.EncodeToString(m.Sum(nil))

}
