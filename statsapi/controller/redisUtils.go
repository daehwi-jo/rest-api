package controller

import (
	"bytes"
	"charlie/i3.0.0/cls"
	"encoding/json"
	"fmt"
	stats "hydrawebapi/statsapi/model"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
)

var RedisConnection redis.Conn
var AuccessToken string
var RedisAddr string
var RedisPassword string
var BaseAPI_URL string
var BillingAPI_URL string
var CustomerAPI_URL string
var DomainAPI_URL string
var ExtraAPI_URL string
var ProductAPI_URL string
var StatsAPI_URL string
var SystemAPI_URL string
var UserAPI_URL string
var RCSAPI_URL string
var Refresh_TTL int
var Token_TTL int
var SkipURL string

type ApiTokenParam struct {
	Index string `json:"index"`
	Value string `json:"value"`
}

type RespApiToken struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Token string `json:"token"`
	} `json:"data"`
}

type MessageValue struct {
	MessageType string `json:"mtype"`
	Kr          string `json:"kr"`
	En          string `json:"en"`
	Cn          string `json:"cn"`
}

func Redis_App_conf(fileName string) bool {

	v, r := cls.GetTokenValue("Redis_Address", fileName) //IP:PORT  1.1.1.1:6379
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] Redis_Address paser error, tmp value(%s)", v)
			return false
		}
		RedisAddr = v
	} else {
		lprintf(1, "[ERR ] Redis_Address NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("Redis_Password", fileName) //password
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] Redis_Password paser error, tmp value(%s)", v)
			RedisPassword = ""
		}
		RedisPassword = v
	} else {
		lprintf(1, "[ERR ] Redis_Password NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("BaseAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] BaseAPI_URL  paser error, tmp value(%s)", v)
			return false
		}
		BaseAPI_URL = v
	} else {
		lprintf(1, "[ERR ] BaseAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("BillingAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] BillingAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		BillingAPI_URL = v
	} else {
		lprintf(1, "[ERR ] BillingAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("CustomerAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] CustomerAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		CustomerAPI_URL = v
	} else {
		lprintf(1, "[ERR ] CustomerAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("DomainAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] DomainAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		DomainAPI_URL = v
	} else {
		lprintf(1, "[ERR ] DomainAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("ExtraAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] ExtraAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		ExtraAPI_URL = v
	} else {
		lprintf(1, "[ERR ] ExtraAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("ProductAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] ProductAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		ProductAPI_URL = v
	} else {
		lprintf(1, "[ERR ] ProductAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("StatsAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] StatsAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		StatsAPI_URL = v
	} else {
		lprintf(1, "[ERR ] StatsAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("SystemAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] SystemAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		SystemAPI_URL = v
	} else {
		lprintf(1, "[ERR ] SystemAPI_URL    NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("UserAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] UserAPI_URL    paser error, tmp value(%s)", v)
			return false
		}
		UserAPI_URL = v
	} else {
		lprintf(1, "[ERR ] UserAPI_URL NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("RCSAPI_URL", fileName) //http://IP:PORT OR URL http://www.ex.com full name
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] RCSAPI_URL paser error, tmp value(%s)", v)
			return false
		}
		RCSAPI_URL = v
	} else {
		lprintf(1, "[ERR ] RCSAPI_URL  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("Refresh_TTL", fileName) //JWT made TTL
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] Refresh_TTL paser error, tmp value(%s),default Value setting", v)
			Refresh_TTL = 604800
		}
		va, err2 := strconv.Atoi(v)
		if err2 != nil {
			lprintf(1, "[ERR ] Refresh_TTL paser error, tmp value(%s),default Value setting", v)
			Refresh_TTL = 604800
		} else {
			Refresh_TTL = va
		}
	} else {
		lprintf(1, "[ERR ] Refresh_TTL NOT SETTING, default Value setting")
		Refresh_TTL = 604800
	}

	v, r = cls.GetTokenValue("Token_TTL", fileName) //Token TTL (REDIS Expire)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] Token_TTL paser error, tmp value(%s),default Value setting", v)
			Token_TTL = 21600
		}
		va, err2 := strconv.Atoi(v)
		if err2 != nil {
			lprintf(1, "[ERR ] Token_TTL paser error, tmp value(%s),default Value setting", v)
			Token_TTL = 21600
		} else {
			Token_TTL = va
		}
	} else {
		lprintf(1, "[ERR ] Token_TTL  NOT SETTING, default Value setting")
		Token_TTL = 21600
	}

	v, r = cls.GetTokenValue("SkipURL", fileName) //token auth free pass URL
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] SkipURL paser error, tmp value(%s)", v)
			SkipURL = ""
		}
		SkipURL = v
	} else {
		SkipURL = ""
	}

	return true
}

/*
//connection 함수
func RedisCon() bool {
	passwordOption := redis.DialPassword(RedisPassword)                  // conf 에서 받아야함
	redisConnection, err := redis.Dial("tcp", RedisAddr, passwordOption) // conf 에서 받아야함
	if err != nil || redisConnection == nil {
		lprintf(1, "[ERR ] Token authentication Server Connection Error: Redis disconnect : message : [%s] ,err: [%s]\n", redisConnection, err)
		return false
	}
	RedisConnection = redisConnection

	return true
}
*/
//connection 함수
func RedisCon() (int, redis.Conn) {
	passwordOption := redis.DialPassword(RedisPassword)                  // conf 에서 받아야함
	redisConnection, err := redis.Dial("tcp", RedisAddr, passwordOption) // conf 에서 받아야함
	if err != nil || redisConnection == nil {
		lprintf(1, "[ERR ] Token authentication Server Connection Error: Redis disconnect : message : [%s] ,err: [%s]\n", redisConnection, err)
		return -1, redisConnection
	}
	return 1, redisConnection
}

//검증함수
func ConfirmTK(tk string) bool {
	rst, RedisConnection := RedisCon()
	if rst < 0 {
		return false
	}
	defer RedisConnection.Close()
	nowtime := uint32(time.Now().Unix())
	interfaceValue, err := RedisConnection.Do("get", tk)
	if interfaceValue == nil || err != nil {
		lprintf(1, "[ERR ] Token authentication, Get error: message: [%s] ,err: [%s]\n", interfaceValue, err)
		return false
	}

	byteValue, _ := interfaceValue.([]byte)
	TokenTimeValue, _ := strconv.Atoi(string(byteValue))

	if uint32(TokenTimeValue+21600) > nowtime {
		lprintf(4, "[INFO] Token Authentication Success \n")
		if (uint32(TokenTimeValue+Token_TTL) - nowtime) < uint32(Token_TTL/2) {
			go RedisConnection.Do("set", tk, nowtime, "EX", Token_TTL)
		}
		return true
	} else {
		lprintf(1, "[ERR ] Token authentication, Expire \n")
		return false
	}
	lprintf(1, "[ERR ] Token authentication, Expire \n")
	return false
}

//발급 저장 함수
func MakeTK(tk string) bool {
	rst, RedisConnection := RedisCon()
	if rst < 0 {
		return false
	}
	defer RedisConnection.Close()
	nowtime := uint32(time.Now().Unix())

	result, err := RedisConnection.Do("set", tk, nowtime, "EX", Token_TTL)
	if result == "OK" {
		lprintf(4, "[INFO] Access Token Redis Save \n")
		return true
	} else {
		lprintf(1, "[ERR ] Access Token Redis Error [%s] \n", err)
	}

	return false
}

// 스케줄러 (토큰 갱신)
func SchedulerLoop() {
	for {
		time.Sleep(time.Duration(Token_TTL-300) * time.Second)
		if !SetApiToken() {
			SetApiToken()
		}
	}
}

// 만료된 토큰 제거
func DeleteTK(tk string) {
	rst, RedisConnection := RedisCon()
	if rst < 0 {
		return
	}
	defer RedisConnection.Close()
	auth := strings.Split(tk, " ")
	//
	var tokenString string
	if len(auth) > 1 {
		tokenString = auth[1]
	} else {
		tokenString = auth[0]
	}
	RedisConnection.Do("del", tokenString)

}

//토큰 요청
func SetApiToken() bool {
	var param ApiTokenParam
	respData := RespApiToken{}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	num1 := random.Intn(30)
	num2 := random.Intn(30)
	var numStr1, numStr2 string

	if num1 < 10 {
		numStr1 = fmt.Sprintf("0%d", num1)
	} else {
		numStr1 = strconv.Itoa(num1)
	}

	if num2 < 10 {
		numStr2 = fmt.Sprintf("0%d", num2)
	} else {
		numStr2 = strconv.Itoa(num2)
	}

	number := fmt.Sprintf("%s%s", numStr1, numStr2)

	key := string(MatrixF[num1][num2]) + string(MatrixF[num2][num1]) + string(MatrixL[num1][num2]) + string(MatrixL[num2][num1])

	param.Index = number
	param.Value = key

	ret, apiCallResp := FuncApiCallRequest("POST", UserAPI_URL+"/v1/user/main/api-token", "json", &param)
	if ret == 1 {

		return false
	}
	if err := json.Unmarshal(apiCallResp, &respData); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)

		return false
	}
	if len(respData.Data.Token) == 0 {

		return false
	}

	AuccessToken = respData.Data.Token
	return true
}

//API 생성
func SetDifApiToken() (string, bool) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	num1 := random.Intn(30)
	num2 := random.Intn(30)
	var numStr1, numStr2 string

	if num1 < 10 {
		numStr1 = fmt.Sprintf("0%d", num1)
	} else {
		numStr1 = strconv.Itoa(num1)
	}

	if num2 < 10 {
		numStr2 = fmt.Sprintf("0%d", num2)
	} else {
		numStr2 = strconv.Itoa(num2)
	}

	number := fmt.Sprintf("%s%s", numStr1, numStr2)

	key := string(MatrixF[num1][num2]) + string(MatrixF[num2][num1]) + string(MatrixL[num1][num2]) + string(MatrixL[num2][num1])

	type jwtCustomClaims struct {
		Number string
		jwt.StandardClaims
	}

	expireTime := time.Hour * 24 * 7
	// Set custom claims
	//expireTime := time.Hour * 1
	claims := &jwtCustomClaims{
		number,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(key)) // sercret key를 사용하여 Access Token 발급
	if err != nil {
		lprintf(1, "[ERR ] JWT token signed error : [%s]\n", err)
		return "", false
	}

	if !MakeTK(accessToken) {
		lprintf(1, "[ERR ] Access Token Redis Error \n")
		return "", false
	}

	return accessToken, true
}

func FuncApiCallRequest(method string, url string, dataType string, data interface{}) (int, []byte) {
	var reqBytes []byte
	var err error
	var apiCallResponse stats.ApiCallResponse

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

	lprintf(4, "[INFO] respBody : [%s]\n", string(respBody))
	if dataType == "json" || dataType == "JSON" {
		if err := json.Unmarshal(respBody, &apiCallResponse); err != nil {
			lprintf(1, "[ERR ] json.Unmarshal : [%s]\n", err)
			return 1, nil
		}

		lprintf(4, "[INFO] apiCallResponse : [%v]\n", apiCallResponse)
		if apiCallResponse.Code == "2000" {
			lprintf(4, "[INFO] api call response code : [%s], Validation Message type \n", apiCallResponse.Code)
			return 2, respBody
		}

		if apiCallResponse.Code != "0000" {
			lprintf(1, "[ERR ] api call response code : [%s]\n", apiCallResponse.Code)
			return 1, nil
		}
	}

	return 0, respBody
}

func retryFuncApiCallRequest(method, url string, buff *bytes.Buffer) (int, []byte) {

	var apiCallResponse stats.ApiCallResponse

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

	if apiCallResponse.Code == "2000" {
		lprintf(4, "[INFO] api call response code : [%s], Validation Message type \n", apiCallResponse.Code)
		return 2, respBody
	}

	if apiCallResponse.Code != "0000" {
		lprintf(1, "[ERR ] api call response code : [%s]\n", apiCallResponse.Code)
		return 1, nil
	}

	return 0, respBody

}
