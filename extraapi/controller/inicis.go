package controller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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

var MatrixF = [30][30]byte{
	{'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'a', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b'},
	{'s', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
	{'d', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'c', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd'},
	{'u', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c'},
	{'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'e', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f'},
	{'w', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e'},
	{'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'g', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h'},
	{'y', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g'},
	{'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'i', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j'},
	{'1', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i'},
	{'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'k', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l'},
	{'3', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k'},
	{'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'm', 'n', '5', 'n', '5', 'n', '5', 'n', '5', 'n', '5', 'n'},
	{'5', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', 'n', 'm', 'n', 'm', 'n', 'm', 'n', 'm', 'n', 'm', 'n', 'm'},
	{'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'o', 'p', '7', 'p', '7', 'p', '7', 'p', '7', 'p', '7', 'p'},
	{'7', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', 'p', 'o', 'p', 'o', 'p', 'o', 'p', 'o', 'p', 'o', 'p', 'o'},
	{'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'q', 'r', '9', 'r', '9', 'r', '9', 'r', '9', 'r', '9', 'r'},
	{'9', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', 'r', 'q', 'r', 'q', 'r', 'q', 'r', 'q', 'r', 'q', 'r', 'q'},
	{'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'a', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b'},
	{'s', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
	{'d', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'c', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd'},
	{'u', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c'},
	{'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'e', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f'},
	{'w', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e'},
	{'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'g', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h'},
	{'y', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g'},
	{'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'i', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j'},
	{'1', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i'},
	{'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'k', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l'},
	{'3', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k'},
}

var MatrixL = [30][30]byte{
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
}

// 결제 모듈 호출
func GetInicisPaymentModuleTest(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	// test url : https://stgstdpay.inicis.com/stdjs/INIStdPay.js

	now := time.Now()

	// 가공이 필요한 데이터
	// oid := fmt.Sprintf("PP-%s", Getyyyymmddhhmmssnnn(now))
	timeStamp := GetMillSecondTimeStamp(now)
	oid := fmt.Sprintf("INIpayTest_%d", timeStamp)
	signature := fmt.Sprintf("oid=%s&price=10000&timestamp=%d", oid, timeStamp)
	price := strconv.FormatInt(int64(10000), 10)
	lprintf(4, "[INFO] oid       : [%s]\n", oid)
	lprintf(4, "[INFO] timeStamp : [%d]\n", timeStamp)
	lprintf(4, "[INFO] signature : [%s]\n", signature)
	lprintf(4, "[INFO] price     : [%s]\n", price)

	// 위변조 검증 데이터
	midHash := GetSha256Encoding("INIpayTest")
	timeStampHash := GetSha256Encoding(strconv.FormatInt(timeStamp, 10))
	signatureHash := GetSha256Encoding(signature)
	mkeyHash := GetSha256Encoding("SU5JTElURV9UUklQTEVERVNfS0VZU1RS") // 테스트 signkey

	// data
	data := url.Values{}
	data.Set("version", "1.0")
	data.Set("mid", midHash)
	data.Set("oid", oid)
	data.Set("goodname", "prime")
	// data.Set("price", priceHash)
	data.Set("price", price)
	data.Set("currency", "WON")
	data.Set("buyername", "securitynet")
	data.Set("buyertel", "02-2672-0102")
	data.Set("timestamp", timeStampHash)
	data.Set("signature", signatureHash)
	data.Set("returnUrl", "https://vueui.securitynetsvc.com:15001")
	data.Set("mKey", mkeyHash)
	data.Set("closeUrl", "https://vueui.securitynetsvc.com:15001")
	data.Set("buyeremail", "info@securitynet.co.kr")
	data.Set("gopaymethod", "Card")
	data.Set("offerPeriod", "20200902-20201001")
	data.Set("languageVuew", "ko")
	data.Set("charset", "UTF-8")
	data.Set("payViewType", "overlay")
	data.Set("merchantData", "")
	data.Set("acceptmethod", "")

	lprintf(4, "[data] %v \n", data)

	surl := "https://stgstdpay.inicis.com/stdjs/INIStdPay.js"
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add(echo.HeaderAcceptEncoding, "gzip, deflate, br")

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
	lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))
	contentLength := int64(len(respBody))
	lprintf(4, "[INFO] contentLength: %d\n", contentLength)

	c.Response().Header().Set(echo.HeaderXXSSProtection, "1; mode=block")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScript)
	c.Response().Header().Set(echo.HeaderContentLength, strconv.FormatInt(contentLength, 10))
	// c.Response().Header().Set("Keep-Alive", "timeout=60")
	return c.HTML(http.StatusOK, string(respBody))
}

// 결제 모듈 호출
func GetInicisPaymentModule3(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	var param extra.GetInicisPaymentModuleParam
	// test url : https://stgstdpay.inicis.com/stdjs/INIStdPay.js

	now := time.Now()

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

	// 가공이 필요한 데이터
	// oid := fmt.Sprintf("PP-%s", Getyyyymmddhhmmssnnn(now))
	timeStamp := GetMillSecondTimeStamp(now)
	oid := fmt.Sprintf("%s_%d", param.Mid, timeStamp)
	signature := fmt.Sprintf("oid=%s&price=%d&timestamp=%d", oid, param.Price, timeStamp)
	price := strconv.FormatInt(int64(param.Price), 10)
	lprintf(4, "[INFO] oid       : [%s]\n", oid)
	lprintf(4, "[INFO] timeStamp : [%d]\n", timeStamp)
	lprintf(4, "[INFO] signature : [%s]\n", signature)
	lprintf(4, "[INFO] price     : [%s]\n", price)

	// 위변조 검증 데이터
	midHash := GetSha256Encoding(param.Mid)
	// priceHash := GetSha256Encoding(string(param.Price))
	// timeStampHash := GetSha256Encoding(string(timeStamp))
	timeStampHash := GetSha256Encoding(strconv.FormatInt(timeStamp, 10))
	signatureHash := GetSha256Encoding(signature)
	mkeyHash := GetSha256Encoding("SU5JTElURV9UUklQTEVERVNfS0VZU1RS") // 테스트 signkey

	// data
	data := url.Values{}
	data.Set("version", param.Version)
	data.Set("mid", midHash)
	data.Set("oid", oid)
	data.Set("goodname", param.GoodName)
	// data.Set("price", priceHash)
	data.Set("price", price)
	data.Set("currency", param.Currency)
	data.Set("buyername", param.BuyerName)
	data.Set("buyertel", param.BuyerTel)
	data.Set("timestamp", timeStampHash)
	data.Set("signature", signatureHash)
	data.Set("returnUrl", param.ReturnUrl)
	data.Set("mKey", mkeyHash)
	data.Set("closeUrl", param.CloseUrl)
	data.Set("popupUrl", param.PopupUrl)
	data.Set("buyeremail", param.BuyerEmail)
	data.Set("gopaymethod", param.GoPayMethod)
	data.Set("offerPeriod", param.OfferPeriod)
	data.Set("languageVuew", param.LanguageView)
	data.Set("charset", param.Charset)
	data.Set("payViewType", param.PayViewType)
	data.Set("merchantData", param.MerchantData)
	data.Set("acceptmethod", param.AcceptMethod)

	lprintf(4, "[data] %v \n", data)

	surl := "https://stgstdpay.inicis.com/stdjs/INIStdPay.js"
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
	//lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))

	contentLength := int64(len(respBody))
	lprintf(4, "[INFO] contentLength: %d\n", contentLength)

	// c.Response().Header().Set(echo.HeaderXXSSProtection, "1; mode=block")
	// c.Response().Header().Set("Content-Type", echo.MIMEApplicationJavaScript)
	// c.Response().Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))

	c.Response().Header().Set(echo.HeaderXXSSProtection, "1; mode=block")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScript)
	c.Response().Header().Set(echo.HeaderContentLength, strconv.FormatInt(contentLength, 10))

	//return c.HTMLBlob(http.StatusOK, respBody)
	// respData.Data = signatureHash
	// lprintf(4, "[INFO] respData : [%s]\n", respData)
	return c.JSON(http.StatusOK, signatureHash)
	// return c.HTML(http.StatusOK, string(respBody))
}

// 결제 승인 응답
// func GetInicisPaymentReturn(c echo.Context) error {
// 	return c.HTML(http.StatusOK, "")
// }

// 결제 취소 응답
func GetInicisPaymentCancel(c echo.Context) error {
	return c.HTML(http.StatusOK, "")
}

// 가상 계좌
func GetInicisVirtualAccount(c echo.Context) error {
	respData := extra.ApiCallResponse{}
	// respJson := extra.RespGetInicisVirtualAccount{}
	// var param extra.GetInicisVirtualAccountParam

	remoteIp := c.Request().RemoteAddr // layer 3 Source: 192.168.41.41
	lprintf(4, "[INFO] remoteIp : [%s]\n", remoteIp)
	realIp := c.RealIP() // layer 7 : X-Forwarded-For: 192.168.30.117, X-Real-IP: 192.168.30.117
	lprintf(4, "[INFO] realIp : [%s]\n", realIp)

	// inicis client ip 유효성 체크 로직 필요!!
	// 현재는 테스트 이므로 주석처리
	// if remoteIp != "203.238.37.15" || remoteIp != "39.115.212.9" || remoteIp != "183.109.71.153" {
	// 	lprintf(1, "[INFO] not inicis client : [%s]\n", remoteIp)
	// }

	// form Data
	noTid := c.FormValue("no_tid")
	noOid := c.FormValue("no_oid")
	lprintf(4, "[INFO] noTid : [%s]\n", noTid)
	lprintf(4, "[INFO] noOid : [%s]\n", noOid)

	// json
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
	// 	lprintf(1, "[ERR ] request body : [%s]\n", err)
	// 	respData.Code = extra.C500
	//
	// 	respData.ServiceName = extra.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }

	return c.JSON(http.StatusOK, respData)
}

// 결제 모듈 호출
func GetInicisPaymentModule(c echo.Context) error {

	type GetInicisPaymentModuleParam struct {
		// 필수 데이터
		Version string `json:"version"`
		//Mid       string `json:"mid"`
		GoodName  string `json:"goodName"`
		Price     int    `json:"price"` // currency "WON" 으로 고정
		Currency  string `json:"currency"`
		BuyerName string `json:"buyerName"`
		BuyerTel  string `json:"buyerTel"`
		Signature string `json:"signature"`
		ReturnUrl string `json:"returnUrl"`
		CloseUrl  string `json:"closeUrl"`
		PopupUrl  string `json:"popupUrl"`
		Platform  string `json:"platform"`
		Oid       string `json:"oid"`
		// timestamp (내부에서 생성)
		// mkey - signkey에 대한 hash 값
		// signature - oid + price + timestamp

		// 옵션 데이터
		BuyerEmail   string `json:"buyerEmail"`
		GoPayMethod  string `json:"goPayMethod"`
		OfferPeriod  string `json:"offerPeriod"`
		LanguageView string `json:"languageView"`
		Charset      string `json:"charset"`
		PayViewType  string `json:"payViewType"`
		MerchantData string `json:"merchantData"`
		AcceptMethod string `json:"acceptMethod"`
		// tax  - 사용안함
		// taxfree - 사용안함
		// parentemail - 사용안함
	}

	respData := extra.ApiCallResponse{}
	var param GetInicisPaymentModuleParam
	// test url : https://stgstdpay.inicis.com/stdjs/INIStdPay.js

	now := time.Now()

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

	// 가공이 필요한 데이터
	// oid := fmt.Sprintf("PP-%s", Getyyyymmddhhmmssnnn(now))
	mid := INICIS_MID
	timeStamp := GetMillSecondTimeStamp(now)
	//oid := fmt.Sprintf("%s_%d", mid, timeStamp)
	signature := fmt.Sprintf("oid=%s&price=%d&timestamp=%d", param.Oid, param.Price, timeStamp)
	price := strconv.FormatInt(int64(param.Price), 10)
	lprintf(4, "[INFO] oid       : [%s]\n", param.Oid)
	lprintf(4, "[INFO] timeStamp : [%d]\n", timeStamp)
	lprintf(4, "[INFO] signature : [%s]\n", signature)
	lprintf(4, "[INFO] price     : [%s]\n", price)

	// 위변조 검증 데이터
	//midHash := GetSha256Encoding(param.Mid)
	// priceHash := GetSha256Encoding(string(param.Price))
	// timeStampHash := GetSha256Encoding(string(timeStamp))
	//timeStampHash := GetSha256Encoding(strconv.FormatInt(timeStamp, 10))
	signatureHash := GetSha256Encoding(signature)
	mkeyHash := GetSha256Encoding(INICIS_SIGN_KEY) // 테스트 signkey

	// data
	// data := url.Values{}
	// data.Set("version", param.Version)
	// data.Set("mid", midHash)
	// data.Set("oid", oid)
	// data.Set("goodname", param.GoodName)
	// // data.Set("price", priceHash)
	// data.Set("price", price)
	// data.Set("currency", param.Currency)
	// data.Set("buyername", param.BuyerName)
	// data.Set("buyertel", param.BuyerTel)
	// data.Set("timestamp", timeStampHash)
	// data.Set("signature", signatureHash)
	// data.Set("returnUrl", param.ReturnUrl)
	// data.Set("mKey", mkeyHash)
	// data.Set("closeUrl", param.CloseUrl)
	// data.Set("popupUrl", param.PopupUrl)
	// data.Set("buyeremail", param.BuyerEmail)
	// data.Set("gopaymethod", param.GoPayMethod)
	// data.Set("offerPeriod", param.OfferPeriod)
	// data.Set("languageVuew", param.LanguageView)
	// data.Set("charset", param.Charset)
	// data.Set("payViewType", param.PayViewType)
	// data.Set("merchantData", param.MerchantData)
	// data.Set("acceptmethod", param.AcceptMethod)

	// lprintf(4, "[data] %v \n", data)

	// surl := "https://stgstdpay.inicis.com/stdjs/INIStdPay.js"
	// req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// client := &http.Client{}
	// respGet, err := client.Do(req)
	// if err != nil {
	// 	lprintf(1, "[ERR ] Do : [%s]\n", err)
	// 	respData.Code = user.C500
	// 	respData.Message = user.FAILMESSAGE
	// 	respData.ServiceName = user.TYPE

	// 	return c.JSON(http.StatusOK, respData)
	// }
	// defer respGet.Body.Close()

	// respBody, err := ioutil.ReadAll(respGet.Body)
	// if err != nil {
	// 	lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
	// 	respData.Code = user.C500
	// 	respData.Message = user.FAILMESSAGE
	// 	respData.ServiceName = user.TYPE
	// 	return c.JSON(http.StatusOK, respData)
	// }
	//lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))

	//contentLength := int64(len(respBody))
	//lprintf(4, "[INFO] contentLength: %d\n", contentLength)

	// c.Response().Header().Set(echo.HeaderXXSSProtection, "1; mode=block")
	// c.Response().Header().Set("Content-Type", echo.MIMEApplicationJavaScript)
	// c.Response().Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))

	//c.Response().Header().Set(echo.HeaderXXSSProtection, "1; mode=block")
	// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScript)
	// c.Response().Header().Set(echo.HeaderContentLength, strconv.FormatInt(contentLength, 10))

	//return c.HTMLBlob(http.StatusOK, respBody)
	// respData.Data = signatureHash
	// lprintf(4, "[INFO] respData : [%s]\n", respData)

	type RespD struct {
		Code        string       `json:"code"`
		Message     MessageValue `json:"message"`
		ServiceName string       `json:"serviceName"`
		Data        struct {
			Oid       string `json:"oid"`
			TimeStamp string `json:"timeStamp"`
			MKey      string `json:"mKey"`
			Signature string `json:"signature"`
			Mid       string `json:"mid"`
		} `json:"data"`
	}

	var respD RespD
	// return c.HTML(http.StatusOK, string(respBody))

	if param.Platform == "PC" {
		respD.Data.Oid = param.Oid
		respD.Data.TimeStamp = strconv.FormatInt(timeStamp, 10)
		respD.Data.MKey = mkeyHash
		respD.Data.Signature = signatureHash
		respD.Data.Mid = mid
	} else if param.Platform == "MOBILE" {
		respD.Data.Oid = param.Oid
		respD.Data.Mid = mid
	}

	lprintf(4, "[INFO] respD.Data      : [%s]\n", respD)
	respD.Code = extra.SUCCESS
	respD.ServiceName = extra.TYPE
	return c.JSON(http.StatusOK, respD)
}

func GetInicisPaymentReturn(c echo.Context) error {
	//url := c.QueryParam("returnurl")
	//lprintf(4, "[INFO] test : 11111111111111111\n")

	//var param map[string]string
	var returnUrl string
	param := make(map[string]string)
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
	}

	lprintf(4, "[INFO] resp : [%s]\n", string(resp))
	//result, _ := url.QueryUnescape(string(resp))

	tmp := strings.Split(strings.TrimSpace(string(resp)), "&")

	for _, val := range tmp {
		tmps := strings.Split(val, "=")
		//param[tmps[0]] = tmps[1]
		//if tmps[0] == "resultMsg" || tmps[0] == "returnUrl" || tmps[0] == "authUrl" || tmps[0] == "checkAckUrl" || tmps[0] == "netCancelUrl" {
		if strings.Contains(tmps[0], "Msg") || strings.Contains(tmps[0], "Url") || strings.Contains(tmps[0], "Token") {
			param[tmps[0]], _ = url.QueryUnescape(tmps[1])
		} else {
			param[tmps[0]] = tmps[1]
		}
	}

	/*
		if err := json.Unmarshal([]byte(result), &param); err != nil {
			lprintf(1, "[ERR ] request body : [%s]\n", err)
		}
	*/
	/*
		for key, value := range param {
			lprintf(4, "[INFO] key(%s), value(%s) \n", key, value)
		}
	*/

	resultCode, exists := param["resultCode"]
	if !exists {
		lprintf(4, "[INFO] PaymentRedirect : [%s]\n", returnUrl)

		// error page 호출
		returnUrl = PAYMENT_REDIRECT_FE
		c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리

		return c.HTML(http.StatusMovedPermanently, "")
	}

	merchantData, exist := param["merchantData"]
	lprintf(4, "[INFO] merchantData : %s  \n", merchantData)
	if !exist {
		// map에 merchantData 없을 때
	}

	if merchantData == "FE" {
		returnUrl = PAYMENT_FAIL_REDIRECT_FE
	} else {
		returnUrl = PAYMENT_FAIL_REDIRECT_BO
	}

	if resultCode == "0000" {
		// process

		// inics 요청
		// 응답 받아와서
		// url 301

		//a := "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:8d52771c64ec4cad75efac84b7f9a7a823156037cc0a4ab8cc738ec935337962"

		lprintf(4, "[INFO] test 1 \n")

		//token := strings.Trim(param["authToken"], " ")
		token, exist := param["authToken"]
		if !exist {
			// map에 token 없을 때
		}

		reqUrl, exist := param["authUrl"]
		if !exist {
			// map에 authUrl 없을 때
		}

		mid, exist := param["mid"]
		if !exist {
			// map에 mid 없을 때
		}

		charset, exist := param["charset"]
		if !exist {
			// map에 charset 없을 때
		}

		// signature hash 생성
		token = strings.TrimSpace(token)
		now := time.Now()
		timeStamp := GetMillSecondTimeStamp(now)
		ts := strconv.FormatInt(timeStamp, 10)
		signature := fmt.Sprintf("authToken=%s&timestamp=%s", token, ts)
		signatureHash := GetSha256Encoding(signature)

		// data form 생성
		data := url.Values{}
		data.Set("mid", mid)
		data.Set("authToken", token)
		data.Set("timestamp", ts)
		data.Set("signature", signatureHash)
		data.Set("charset", charset)
		data.Set("format", "JSON")

		lprintf(4, "[data] %v \n", data)

		req, err := http.NewRequest("POST", reqUrl, strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		respGet, err := client.Do(req)
		if err != nil {
			lprintf(1, "[ERR ] Do : [%s]\n", err)
		}
		defer respGet.Body.Close()

		lprintf(4, "[INFO] req url(%s) \n", reqUrl)

		lprintf(4, "[INFO] mid(%s) \n", mid)
		lprintf(4, "[INFO] authToken(%s) \n", token)
		lprintf(4, "[INFO] timestamp(%s) \n", ts)
		lprintf(4, "[INFO] signature(%s) \n", signatureHash)
		lprintf(4, "[INFO] charset(%s) \n", charset)
		lprintf(4, "[INFO] format(%s) \n", "JSON")

		/*
			// 간단한 http.PostForm 예제
			respGet, err := http.PostForm(reqUrl, url.Values{"mid": {mid}, "authToken": {token}, "timestamp": {ts}, "signature": {signatureHash}, "charset": {charset}, "format": {"JSON"}})
			if err != nil {
				lprintf(1, "[ERR ] Do : [%s]\n", err)
			}
			defer respGet.Body.Close()
		*/

		respBody, err := ioutil.ReadAll(respGet.Body)
		if err != nil {
			lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		}

		lprintf(4, "[INFO] respBody : [%s]\n", respBody)

		var result extra.InicisReturn
		if err := json.Unmarshal(respBody, &result); err != nil {
			lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		}

		//결제 성공
		if result.ResultCode == "0000" {
			var sendParam extra.RespGetAlipayReturn

			expireDate := fmt.Sprintf("%s%s", result.VactDate, result.VactTime)
			// time, err := time.Parse("2006-01-02 15:04:05", expireDate)
			// if err != nil {
			// 	return c.HTML(http.StatusMovedPermanently, "time parsing err")
			// }

			sendParam.OrderId = result.Moid
			sendParam.TradeNumber = result.Tid
			sendParam.PayMethod = result.PayMethod
			sendParam.PgData = string(respBody)
			sendParam.Status = "complete"
			if result.PayMethod == "VBank" {
				sendParam.VirtualAccountBankCode = result.VactBankCode
				sendParam.VirtualAccountBankName = result.VactBankName
				sendParam.VirtualAccountNo = result.VactNum
				// sendParam.VirtualAccountExpireDateTime = time.String()
				sendParam.VirtualAccountExpireDateTime = expireDate
				sendParam.Status = "depositwait"
			}
			lprintf(4, "[INFO] orderId  : [%s]\n", sendParam.OrderId)
			lprintf(4, "[INFO] tradeNo  : [%s]\n", sendParam.TradeNumber)
			if merchantData == "FE" {
				returnUrl = fmt.Sprintf("%s?manageNo=%s", PAYMENT_REDIRECT_FE, result.Moid)
				//returnUrl = fmt.Sprintf("https://hydra.securitynetsvc.com:15001/MyAccount/Complete-payment?manageNo=%s", result.Moid)
			}
			sendPaymentInfo(sendParam)
		} else {
			if merchantData == "FE" {
				returnUrl = fmt.Sprintf("%s?errMsg=%s", PAYMENT_FAIL_REDIRECT_FE, result.ResultMsg)
				// returnUrl = fmt.Sprintf("https://hydra.securitynetsvc.com:15001/MyAccount/payment?errMsg=%s", result.ResultMsg)
			}
		}

	}
	// respHtml := `<html><script text="javascript/text">init();function init() {window.close();}</script></html>`
	// return c.HTML(http.StatusOK, respHtml)

	lprintf(4, "[INFO] complete inicis payment!\n")
	//time.Sleep(1 * time.Millisecond)
	lprintf(4, "[INFO] PaymentRedirect : [%s]\n", returnUrl)
	c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리
	return c.HTML(http.StatusMovedPermanently, "")
}

// 이니시스 모바일 결제 리턴 (F/E만 해당)
func GetInicisPaymentMobileReturn(c echo.Context) error {
	var returnUrl string
	param := make(map[string]string)

	lprintf(4, "[INFO] PaymentRedirect : [%s]\n", PAYMENT_FAIL_REDIRECT_FE)
	lprintf(4, "[INFO] PaymentRedirect : [%s]\n", PAYMENT_REDIRECT_FE)

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
	}
	lprintf(4, "[INFO] resp : [%s]\n", string(resp))

	// resultCode := c.QueryParam("P_STATUS")
	// tid := c.QueryParam("P_TID")
	// reqUrl := c.QueryParam("P_REQ_URL")

	tmp := strings.Split(strings.TrimSpace(string(resp)), "&")
	for _, v := range tmp {
		tmps := strings.Split(v, "=")
		if strings.Contains(tmps[0], "P_REQ_URL") {
			param[tmps[0]], _ = url.QueryUnescape(tmps[1])
		} else {
			param[tmps[0]] = tmps[1]
		}
	}

	resultCode, exist := param["P_STATUS"]
	if !exist {
		// map에 resultCode 없을 때
	}

	tid, exist2 := param["P_TID"]
	if !exist2 {
		// map에 tid 없을 때
	}

	reqUrl, exist3 := param["P_REQ_URL"]
	if !exist3 {
		// map에 reqUrl 없을 때
	}

	returnUrl = PAYMENT_FAIL_REDIRECT_FE

	if resultCode == "00" {

		// data form 생성
		data := url.Values{}
		data.Set("P_MID", INICIS_MID)
		data.Set("P_TID", tid)
		data.Set("accept-charset", "EUC-KR")
		//data.Set("format", "JSON")
		//data.Set("P_CHARSET", "utf-8")

		lprintf(4, "[data] %v \n", data)

		req, err := http.NewRequest("POST", reqUrl, strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("accept-charset", "EUC-KR")

		client := &http.Client{}
		respGet, err := client.Do(req)
		if err != nil {
			lprintf(1, "[ERR ] Do : [%s]\n", err)
		}
		defer respGet.Body.Close()

		respBody2, err := ioutil.ReadAll(respGet.Body)
		if err != nil {
			lprintf(1, "[ERR ] ReadAll : [%s]\n", err)
		}

		respBody := ConvertEucKrDecoding(respBody2)

		lprintf(4, "[INFO] respGet.Header : [%v]\n", respGet.Header)
		lprintf(4, "[INFO] respBody : [%s]\n", respBody)

		//var result extra.InicisReturnMobile
		// if err := json.Unmarshal(respBody, &result); err != nil {
		// 	lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		// }
		var pgData string
		tmp := strings.Split(strings.TrimSpace(respBody), "&")
		for _, v := range tmp {
			tmps := strings.Split(v, "=")
			// if strings.Contains(tmps[0], "P_REQ_URL") || strings.Contains(tmps[0], "P_RMESG1") || strings.Contains(tmps[0], "P_UNAME") || strings.Contains(tmps[0], "P_NEXT_URL") || strings.Contains(tmps[0], "P_VACT_NAME") || strings.Contains(tmps[0], "P_FN_NM") || strings.Contains(tmps[0], "P_CSHR_MSG") {
			// 	// 	param[tmps[0]], _ = url.QueryUnescape(tmps[1])
			// 	// } else {
			// 	// 	param[tmps[0]] = tmps[1]
			// 	// }
			// 	tmps[1] = ConvertEucKrDecoding([]byte(tmps[1]))
			// }
			param[tmps[0]] = tmps[1]
			pgData += fmt.Sprintf("%s:%s ,", tmps[0], tmps[1])
		}
		lprintf(4, "[INFO] inicis data : [%v]\n", param)

		if pgData != "" {
			pgData = strings.TrimRight(pgData, ",")
		}

		//결제 성공
		if param["P_STATUS"] == "00" {
			var sendParam extra.RespGetAlipayReturn

			expireDate := fmt.Sprintf("%s%s", param["P_VACT_DATE"], param["P_VACT_TIME"])

			sendParam.OrderId = param["P_OID"]
			sendParam.TradeNumber = param["P_TID"]
			sendParam.PayMethod = param["P_TYPE"]
			sendParam.PgData = pgData
			sendParam.Status = "complete"
			if param["P_TYPE"] == "VBank" || param["P_TYPE"] == "VBANK" {
				sendParam.VirtualAccountBankCode = param["P_VACT_BANK_CODE"]
				sendParam.VirtualAccountBankName = param["P_VACT_NAME"]
				sendParam.VirtualAccountNo = param["P_VACT_NUM"]
				sendParam.VirtualAccountExpireDateTime = expireDate
				sendParam.Status = "depositwait"
			}
			lprintf(4, "[INFO] orderId  : [%s]\n", sendParam.OrderId)
			lprintf(4, "[INFO] tradeNo  : [%s]\n", sendParam.TradeNumber)

			returnUrl = fmt.Sprintf("%s?manageNo=%s", PAYMENT_REDIRECT_FE, param["P_OID"])
			sendPaymentInfo(sendParam)
		} else {
			returnUrl = fmt.Sprintf("%s?errMsg=%s", PAYMENT_FAIL_REDIRECT_FE, param["P_RMESG1"])
		}

	}

	lprintf(4, "[INFO] complete inicis payment!\n")

	lprintf(4, "[INFO] PaymentRedirect : [%s]\n", returnUrl)
	c.Response().Header().Set(echo.HeaderLocation, returnUrl) // 정상 처리 후 redirect 처리
	return c.HTML(http.StatusMovedPermanently, "")
}

// name     : json api call function
// method   : GET, POST, PUT, PATCH, DELETE
// url      : request url - http://192.168.x.x:7000/v1/system/....
// dataType : json, xml ... 현재는 json만 해당 http header setting 시 필요
// data     : 요청 보낼 데이터
// return   : 정상 - 0, []byte 에러 - 1, nil
func FuncApiCallRequest(method string, url string, dataType string, data interface{}) (int, []byte) {
	var reqBytes []byte
	var err error
	var apiCallResponse extra.ApiCallResponse

	if dataType == "json" || dataType == "JSON" {
		reqBytes, err = json.Marshal(data)
		if err != nil {
			lprintf(1, "[ERR ] Marshal : [%s]\n", err)
			return 1, nil
		}
	} else {
		lprintf(1, "[ERR ] http request method error : [%s]\n", method)
		return 1, nil
	}

	buff := bytes.NewBuffer(reqBytes)

	// create http request
	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		lprintf(1, "[ERR ] http new requset err(%s) \n", err.Error())
		return 1, nil
	}

	req.Header.Add("Authorization", AuccessToken)

	// set http-header
	if dataType == "json" || dataType == "JSON" {
		req.Header.Add("Content-Type", "application/json")
	}

	// send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		return 1, nil
	}
	if resp.StatusCode == 401 { //토큰 만료 시 1회 다시 요청
		resp.Body.Close()
		return retryFuncApiCallRequest(method, url, buff)
	}
	// check response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		resp.Body.Close()
		return 1, nil
	}
	resp.Body.Close()

	if dataType == "json" || dataType == "JSON" {
		if err := json.Unmarshal(respBody, &apiCallResponse); err != nil {
			lprintf(1, "[ERR ] json.Unmarshal : [%s]\n", err)
			return 1, nil
		}

		lprintf(4, "[INFO] apiCallResponse : [%v]\n", apiCallResponse)
		if apiCallResponse.Code != "0000" {
			lprintf(1, "[ERR ] api call response code : [%s]\n", apiCallResponse.Code)
			return 1, nil
		}
	}

	return 0, respBody
}

func retryFuncApiCallRequest(method, url string, buff *bytes.Buffer) (int, []byte) {

	var apiCallResponse extra.ApiCallResponse

	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		lprintf(1, "[ERR ] http new requset err(%s) \n", err.Error())
		return 1, nil
	}

	//
	if !SetApiToken() {
		return 1, nil
	}

	req.Header.Add("Authorization", AuccessToken)

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERR ] client.Do : [%s]\n", err)
		return 1, nil
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lprintf(1, "[ERR ] ioutil.ReadAll : [%s]\n", err)
		resp.Body.Close()
		return 1, nil
	}
	resp.Body.Close()

	if err := json.Unmarshal(respBody, &apiCallResponse); err != nil {
		lprintf(1, "[ERR ] json.Unmarshal : [%s]\n", err)
		return 1, nil
	}

	if apiCallResponse.Code != "0000" {
		lprintf(1, "[ERR ] api call response code : [%s]\n", apiCallResponse.Code)
		return 1, nil
	}

	return 0, respBody

}
func GetTest(c echo.Context) error {
	//var param map[string]string

	// BODY DATA
	returnUrl := c.QueryParam("returnurl")
	// resp, _ := ioutil.ReadAll(c.Request().Body)

	// json.Unmarshal(resp, &param)

	//returnUrl := fmt.Sprintf("%s?id=abc", param["returnurl"])
	lprintf(4, "[INFO] returnUrl 11 : [%s]\n", returnUrl)

	c.Response().Header().Set(echo.HeaderLocation, returnUrl)
	// c.Response().Header().Set(echo.HeaderLocation, "https://hydra.securitynetsvc.com:15001/inicisReturn")

	c.Response().Header().Add("Authorization", AuccessToken)
	//c.Header.Add("Authorization", GetApiToken())
	return c.HTML(http.StatusMovedPermanently, "")
}

func GetInicisPaymentModule4(c echo.Context) error {

	// BODY DATA
	returnUrl := c.QueryParam("returnurl")
	// resp, _ := ioutil.ReadAll(c.Request().Body)

	// json.Unmarshal(resp, &param)

	//returnUrl := fmt.Sprintf("%s?id=abc", param["returnurl"])
	lprintf(4, "[INFO] returnUrl 22 : [%s]\n", returnUrl)

	c.Response().Header().Set(echo.HeaderLocation, returnUrl)
	// c.Response().Header().Set(echo.HeaderLocation, "https://hydra.securitynetsvc.com:15001/inicisReturn")

	c.Response().Header().Add("Authorization", AuccessToken)
	//c.Header.Add("Authorization", GetApiToken())
	return c.HTML(http.StatusMovedPermanently, "")
}

func PostTest(c echo.Context) error {
	var param map[string]string

	// BODY DATA
	resp, _ := ioutil.ReadAll(c.Request().Body)

	json.Unmarshal(resp, &param)

	returnUrl := fmt.Sprintf("%s?id=abc", param["returnUrl"])
	lprintf(4, "[INFO] returnUrl : [%s]\n", returnUrl)

	c.Response().Header().Set(echo.HeaderLocation, "https://hydra.securitynetsvc.com:15001/inicisReturn")
	c.Response().Header().Add("Authorization", AuccessToken)
	//c.Response().Header().Set(echo.HeaderLocation, returnUrl)
	return c.HTML(http.StatusMovedPermanently, "")
}

// INICIS 가상계좌 입금대기 결제 취소
func GetInicisPaymentGVacctRefund(c echo.Context) error {
	respData := extra.RespGetInicisPaymentRefund{}
	//var param extra.GetInicisPaymentRefundParam
	var param extra.PaymentCancelOneParam

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

	typeText := "Refund"                                                                                  // Refund 고정
	paymethod := "GVacct"                                                                                 // 지불수단 코드 (가상계좌일 경우 Vacct 고정)
	timeStamp := time.Now().Format("20060102150405")                                                      // 전문생성시간(YYYYMMDDhhmmss)
	clientIp := "192.168.41.3"                                                                            // 가맹점 요청 서버IP (추후 거래 확인 등에 사용됨)
	mid := INICIS_MID                                                                                     // 가맹점 ID - 바꿔야함
	tid := param.PgApprovalNo                                                                             // 취소요청 TID
	msg := ""                                                                                             // 취소요청사유                                                                                  // 환불계좌 예금주명
	hash := fmt.Sprintf("%s%s%s%s%s%s%s", INICIS_KEY, typeText, paymethod, timeStamp, clientIp, mid, tid) // 전문위변조 HASH KEY+type+paymethod+timestamp+clientIp+mid+tid
	hashData := GetSha512Encoding(hash)

	data := url.Values{}
	data.Set("type", typeText)
	data.Set("paymethod", paymethod)
	data.Set("timestamp", timeStamp)
	data.Set("clientIp", clientIp)
	data.Set("mid", mid)
	data.Set("tid", tid)
	data.Set("msg", msg)
	data.Set("hashData", hashData)

	lprintf(4, "[INFO] data : [%s]\n", data)

	surl := "https://iniapi.inicis.com/api/v1/refund"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

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

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] xml.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if respData.Data.ResultCode != "00" {
		// 실패
		respData.Code = extra.FAIL

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	return c.JSON(http.StatusOK, respData)
}

// INICIS 결제 취소 요청
func GetInicisPaymentRefund(c echo.Context) error {
	respData := extra.RespGetInicisPaymentRefund{}
	//var param extra.GetInicisPaymentRefundParam
	var param extra.PaymentCancelOneParam

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

	typeText := "Refund"                                                                                  // Refund 고정
	paymethod := param.PaymentWay                                                                         // 지불수단 코드 (가상계좌일 경우 Vacct 고정)
	timeStamp := time.Now().Format("20060102150405")                                                      // 전문생성시간(YYYYMMDDhhmmss)
	clientIp := "192.168.41.3"                                                                            // 가맹점 요청 서버IP (추후 거래 확인 등에 사용됨)
	mid := INICIS_MID                                                                                     // 가맹점 ID - 바꿔야함
	tid := param.PgApprovalNo                                                                             // 취소요청 TID
	msg := ""                                                                                             // 취소요청사유                                                                                  // 환불계좌 예금주명
	hash := fmt.Sprintf("%s%s%s%s%s%s%s", INICIS_KEY, typeText, paymethod, timeStamp, clientIp, mid, tid) // 전문위변조 HASH KEY+type+paymethod+timestamp+clientIp+mid+tid
	hashData := GetSha512Encoding(hash)

	data := url.Values{}
	data.Set("type", typeText)
	data.Set("paymethod", paymethod)
	data.Set("timestamp", timeStamp)
	data.Set("clientIp", clientIp)
	data.Set("mid", mid)
	data.Set("tid", tid)
	data.Set("msg", msg)
	data.Set("hashData", hashData)

	lprintf(4, "[INFO] data : [%s]\n", data)

	surl := "https://iniapi.inicis.com/api/v1/refund"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

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

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] xml.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if respData.Data.ResultCode != "00" {
		// 실패
		respData.Code = extra.FAIL

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	return c.JSON(http.StatusOK, respData)
}

// INICIS 가상계좌 결제 취소 요청
func GetInicisPaymentVBankRefund(c echo.Context) error {
	respData := extra.RespGetInicisPaymentRefund{}
	//var param extra.GetInicisPaymentVBankRefundParam
	var param extra.PaymentCancelOneParam

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

	encryptedData := AESEncrypt(param.RefundAcctNum, []byte(INICIS_KEY))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println(encryptedString)

	typeText := "Refund"                             // Refund 고정
	paymethod := "Vacct"                             // 지불수단 코드 (가상계좌일 경우 Vacct 고정)
	timeStamp := time.Now().Format("20060102150405") // 전문생성시간(YYYYMMDDhhmmss)
	clientIp := "192.168.41.3"                       // 가맹점 요청 서버IP (추후 거래 확인 등에 사용됨)
	mid := INICIS_MID                                // 가맹점 ID - 바꿔야함
	tid := param.PgApprovalNo                        // 취소요청 TID
	msg := ""                                        // 취소요청사유
	refundAcctNum := encryptedString                 // 환불계좌번호
	refundBankCode := param.RefundBankCode
	refundAcctName := param.RefundAcctName
	hash := fmt.Sprintf("%s%s%s%s%s%s%s%s", INICIS_KEY, typeText, paymethod, timeStamp, clientIp, mid, tid, refundAcctNum) // 전문위변조 HASH KEY+type+paymethod+timestamp+clientIp+mid+tid+refundAcctNum
	hashData := GetSha512Encoding(hash)

	data := url.Values{}
	data.Set("type", typeText)
	data.Set("paymethod", paymethod)
	data.Set("timestamp", timeStamp)
	data.Set("clientIp", clientIp)
	data.Set("mid", mid)
	data.Set("tid", tid)
	data.Set("msg", msg)
	data.Set("hashData", hashData)
	data.Set("refundAcctNum", refundAcctNum)
	data.Set("refundBankCode", refundBankCode)
	data.Set("refundAcctName", refundAcctName)

	lprintf(4, "[INFO] data : [%s]\n", data)

	surl := "https://iniapi.inicis.com/api/v1/refund"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

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

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] xml.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if respData.Data.ResultCode != "00" {
		// 실패
		respData.Code = extra.FAIL
		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	} else {
		respData.Code = extra.SUCCESS
		respData.ServiceName = extra.TYPE
	}

	return c.JSON(http.StatusOK, respData)
}

// INICIS 가상계좌 입금 noti
func GetInicisVBankNoti(c echo.Context) error {
	// respData := extra.ApiCallResponse{}
	param := make(map[string]string)
	var inicisNoti extra.GetInicisVBankNotiParam

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		return c.HTML(http.StatusOK, "FAIL")
	}

	lprintf(4, "[INFO] resp : [%s]\n", resp)
	//resp : [len=0615&no_tid=ININPGVBNKINIpayTest20201228163833983265&no_oid=O443279371&id_merchant=INIpayTest&cd_bank=00000088&cd_deal=00000088&dt_trans=20201228&tm_trans=163833&no_msgseq=9000000030&cd_joinorg=20000050&dt_transbase=20201228&no_transseq= &type_msg=0200&cl_trans=1100&cl_close=0&cl_kor=2&no_msgmanage= &no_vacct=56211992494252&amt_input=880000&amt_check=0&nm_inputbank=__%C5%D7%BD%BA%C6%AE__&nm_input=%C8%AB%B1%E6%B5%BF&dt_inputstd= &dt_calculstd= &flg_close=0&dt_cshr=20201228&tm_cshr=163833&no_cshr_appl=268897130&no_cshr_tid=StdpayVBNKINIpayTest20201228161804321482&no_req_tid=StdpayVBNKINIpayTest20201228161804321482]
	tmp := strings.Split(strings.TrimSpace(string(resp)), "&")
	for _, v := range tmp {
		tmps := strings.Split(v, "=")
		param[tmps[0]] = tmps[1]
	}

	typeMsg, exist := param["type_msg"]
	if !exist {
		return c.HTML(http.StatusOK, "FAIL")
	}

	if typeMsg != "0200" {
		lprintf(1, "[ERR ] inicis vbank fail : [%s]\n", typeMsg)
		return c.HTML(http.StatusOK, "FAIL")
	} else {
		inicisNoti.NoTid = param["no_tid"]
		inicisNoti.NoOid, _ = url.QueryUnescape(param["no_oid"])
		ret, _ := FuncApiCallRequest("POST", BillingAPI_URL+"/v1/billing/payment/inicis-vbank/return", "json", &inicisNoti)
		if ret == 1 {
			return c.HTML(http.StatusOK, "FAIL")
		}
	}

	return c.HTML(http.StatusOK, "OK")
}

// INICIS 가상계좌 입금 noti - MOBILE
func GetInicisVBankNotiMobile(c echo.Context) error {
	// respData := extra.ApiCallResponse{}
	param := make(map[string]string)
	var inicisNoti extra.GetInicisVBankNotiParam

	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		return c.HTML(http.StatusOK, "FAIL")
	}

	lprintf(4, "[INFO] resp : [%s]\n", resp)
	//resp : [len=0615&no_tid=ININPGVBNKINIpayTest20201228163833983265&no_oid=O443279371&id_merchant=INIpayTest&cd_bank=00000088&cd_deal=00000088&dt_trans=20201228&tm_trans=163833&no_msgseq=9000000030&cd_joinorg=20000050&dt_transbase=20201228&no_transseq= &type_msg=0200&cl_trans=1100&cl_close=0&cl_kor=2&no_msgmanage= &no_vacct=56211992494252&amt_input=880000&amt_check=0&nm_inputbank=__%C5%D7%BD%BA%C6%AE__&nm_input=%C8%AB%B1%E6%B5%BF&dt_inputstd= &dt_calculstd= &flg_close=0&dt_cshr=20201228&tm_cshr=163833&no_cshr_appl=268897130&no_cshr_tid=StdpayVBNKINIpayTest20201228161804321482&no_req_tid=StdpayVBNKINIpayTest20201228161804321482]
	tmp := strings.Split(strings.TrimSpace(string(resp)), "&")
	for _, v := range tmp {
		tmps := strings.Split(v, "=")
		param[tmps[0]] = tmps[1]
	}
	lprintf(4, "[INFO] inicis data : [%s]\n", param)

	status, exist := param["P_STATUS"]
	if !exist {
		return c.HTML(http.StatusOK, "FAIL")
	}

	if status == "00" {
		lprintf(1, "[ERR ] inicis vbank request : [%s]\n", status)
		return c.HTML(http.StatusOK, "OK")
	} else if status != "02" {
		lprintf(1, "[ERR ] inicis vbank fail : [%s]\n", status)
		return c.HTML(http.StatusOK, "FAIL")
	} else {
		inicisNoti.NoTid = param["P_TID"]
		inicisNoti.NoOid = param["P_OID"]
		ret, _ := FuncApiCallRequest("POST", BillingAPI_URL+"/v1/billing/payment/inicis-vbank/return", "json", &inicisNoti)
		if ret == 1 {
			return c.HTML(http.StatusOK, "FAIL")
		}
	}

	return c.HTML(http.StatusOK, "OK")
}

// INICIS 현금영수증 발행
func GetInicisCashReceipt(c echo.Context) error {
	respData := extra.RespGetInicisCashReceipt{}
	var param extra.GetInicisCashReceiptParam

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

	encryptedData := AESEncrypt(param.RegNum, []byte(INICIS_KEY))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)

	typeText := "Issue"                              // Issue 고정
	paymethod := "Receipt"                           // "Receipt" 고정
	timeStamp := time.Now().Format("20060102150405") // 전문생성시간(YYYYMMDDhhmmss)
	clientIp := "192.168.41.3"                       // 가맹점 요청 서버IP (추후 거래 확인 등에 사용됨)
	mid := INICIS_MID                                // 상점에 발급된 가맹점 ID
	goodName := param.GoodName                       // 상품명
	currency := param.Currency                       // 통화코드 (필수아님)
	buyerName := param.BuyerName                     // 구매자명
	buyerEmail := param.BuyerEmail                   // 구매자이메일 ("@", "." 외 특수문자 입력불가)
	buyerTel := param.BuyerTel                       // 구매자 연락처 (필수아님)
	crPrice := fmt.Sprintf("%g", param.CrPrice)      // 결제금액
	supPrice := fmt.Sprintf("%g", param.SupPrice)    // 공급가액
	tax := fmt.Sprintf("%g", param.Tax)              // 부가세
	srcvPrice := "0"                                 // 봉사료
	regNum := encryptedString                        // 현금영수증 식별번호(주민번호,휴대폰번호,사업자번호)
	useOpt := "0"                                    // 현금영수증 발행용도(0:소득공제용, 1:지출증빙)
	compayNumber := param.CompayNumber               // 서브몰사업자번호 (서브몰가맹점 등록 요청 후 사용 가능합니다.) (필수아님)

	hash := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", "ItEQKi3rY7uvDS8l", typeText, paymethod, timeStamp, clientIp, mid, crPrice, supPrice, srcvPrice, regNum) // 	전문위변조(KEY+type+paymethod+timestamp+clientIp+mid+crPrice+supPrice+srcvPrice+regNum)
	hashData := GetSha512Encoding(hash)

	data := url.Values{}
	data.Set("type", typeText)
	data.Set("paymethod", paymethod)
	data.Set("timestamp", timeStamp)
	data.Set("clientIp", clientIp)
	data.Set("mid", mid)
	data.Set("goodName", goodName)
	data.Set("currency", currency)
	data.Set("buyerName", buyerName)
	data.Set("buyerEmail", buyerEmail)
	data.Set("buyerTel", buyerTel)
	data.Set("crPrice", crPrice)
	data.Set("supPrice", supPrice)
	data.Set("tax", tax)
	data.Set("srcvPrice", srcvPrice)
	data.Set("regNum", regNum)
	data.Set("useOpt", useOpt)
	data.Set("compayNumber", compayNumber)
	data.Set("hashData", hashData)

	lprintf(4, "[INFO] data : [%s]\n", data)

	surl := "https://iniapi.inicis.com/api/v1/receipt"
	lprintf(4, "[INFO] surl : [%s]\n", surl)
	req, err := http.NewRequest("POST", surl, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

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

	if err := json.Unmarshal(respBody, &respData.Data); err != nil {
		lprintf(1, "[ERR ] xml.Unmarshal : [%s]\n", err)
		respData.Code = extra.C500

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if respData.Data.ResultCode != "00" {
		// 실패
		respData.Code = extra.FAIL

		respData.ServiceName = extra.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	respData.Code = extra.SUCCESS

	respData.ServiceName = extra.TYPE
	return c.JSON(http.StatusOK, respData)
}

// func Ase256(plaintext string, key string, iv string, blockSize int) string {
// 	bKey := []byte(key)
// 	bIV := []byte(iv)
// 	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
// 	block, _ := aes.NewCipher(bKey)
// 	ciphertext := make([]byte, len(bPlaintext))
// 	mode := cipher.NewCBCEncrypter(block, bIV)
// 	mode.CryptBlocks(ciphertext, bPlaintext)
// 	return hex.EncodeToString(ciphertext)
// }

// var (
// 	initialVector = "HYb3yQ4f65QL89=="
// 	passphrase    = "ItEQKi3rY7uvDS8l"
// )

// func main() {
// 	var plainText = "hello world"

// 	encryptedData := AESEncrypt(plainText, []byte(passphrase))
// 	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
// 	fmt.Println(encryptedString)

// 	encryptedData, _ = base64.StdEncoding.DecodeString(encryptedString)
// 	decryptedText := AESDecrypt(encryptedData, []byte(passphrase))
// 	fmt.Println(string(decryptedText))
// }

func AESEncrypt(src string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(INICIS_IV))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
