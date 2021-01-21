package controller

import (
	"charlie/i3.0.0/cls"
	"encoding/json"
	"fmt"

	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
)

//var RedisConnection redis.Conn
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

///////
var API_KEY string
var AUTH_USER_ID string
var WHOIS_URL string
var WHOIS_CHECK_URL string
var CUSTOMER_ID string
var REG_CONTACT_ID string
var ADMIN_CONTACT_ID string
var TECH_CONTACT_ID string
var BILLING_CONTACT_ID string
var INVOICE_OPTION string
var PROTECT_PRIVATE string
var YEARS string
var GABIA_URL string
var GABIA_ID string
var GABIA_PASS string
var ALIPAY_PARTNER_ID string
var ALIPAY_KEY string
var ALIPAY_OPER_URL string
var PAYPAL_PAYMENT_URL string
var PAYPAL_REFUND_URL string
var PAYPAL_ID string
var PAYMENT_REDIRECT_BO string
var PAYMENT_REDIRECT_FE string
var PAYMENT_FAIL_REDIRECT_FE string
var PAYMENT_FAIL_REDIRECT_BO string
var INICIS_MID string
var INICIS_SIGN_KEY string
var INICIS_KEY string
var INICIS_IV string
var WHOXY_KEY string

//////

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
	//////
	v, r = cls.GetTokenValue("API_KEY", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] API_KEY paser error, tmp value(%s)", v)
			return false
		}
		API_KEY = v
	} else {
		lprintf(1, "[ERR ] API_KEY  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("AUTH_USER_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] AUTH_USER_ID paser error, tmp value(%s)", v)
			return false
		}
		AUTH_USER_ID = v
	} else {
		lprintf(1, "[ERR ] AUTH_USER_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("WHOIS_URL", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] WHOIS_URL paser error, tmp value(%s)", v)
			return false
		}
		WHOIS_URL = v
	} else {
		lprintf(1, "[ERR ] WHOIS_URL  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("WHOIS_CHECK_URL", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] WHOIS_CHECK_URL paser error, tmp value(%s)", v)
			return false
		}
		WHOIS_CHECK_URL = v
	} else {
		lprintf(1, "[ERR ] WHOIS_CHECK_URL  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("CUSTOMER_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] CUSTOMER_ID paser error, tmp value(%s)", v)
			return false
		}
		CUSTOMER_ID = v
	} else {
		lprintf(1, "[ERR ] CUSTOMER_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("REG_CONTACT_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] REG_CONTACT_ID paser error, tmp value(%s)", v)
			return false
		}
		REG_CONTACT_ID = v
	} else {
		lprintf(1, "[ERR ] REG_CONTACT_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("ADMIN_CONTACT_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] ADMIN_CONTACT_ID paser error, tmp value(%s)", v)
			return false
		}
		ADMIN_CONTACT_ID = v
	} else {
		lprintf(1, "[ERR ] ADMIN_CONTACT_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("TECH_CONTACT_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] TECH_CONTACT_ID paser error, tmp value(%s)", v)
			return false
		}
		TECH_CONTACT_ID = v
	} else {
		lprintf(1, "[ERR ] TECH_CONTACT_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("BILLING_CONTACT_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] BILLING_CONTACT_ID paser error, tmp value(%s)", v)
			return false
		}
		BILLING_CONTACT_ID = v
	} else {
		lprintf(1, "[ERR ] BILLING_CONTACT_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("INVOICE_OPTION", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] INVOICE_OPTION paser error, tmp value(%s)", v)
			return false
		}
		INVOICE_OPTION = v
	} else {
		lprintf(1, "[ERR ] INVOICE_OPTION  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("PROTECT_PRIVATE", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PROTECT_PRIVATE paser error, tmp value(%s)", v)
			return false
		}
		PROTECT_PRIVATE = v
	} else {
		lprintf(1, "[ERR ] PROTECT_PRIVATE  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("YEARS", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] YEARS paser error, tmp value(%s)", v)
			return false
		}
		YEARS = v
	} else {
		lprintf(1, "[ERR ] YEARS  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("GABIA_URL", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] GABIA_URL paser error, tmp value(%s)", v)
			return false
		}
		GABIA_URL = v
	} else {
		lprintf(1, "[ERR ] GABIA_URL  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("GABIA_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] GABIA_ID paser error, tmp value(%s)", v)
			return false
		}
		GABIA_ID = v
	} else {
		lprintf(1, "[ERR ] GABIA_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("GABIA_PASS", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] GABIA_PASS paser error, tmp value(%s)", v)
			return false
		}
		GABIA_PASS = v
	} else {
		lprintf(1, "[ERR ] GABIA_PASS  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("ALIPAY_PARTNER_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] ALIPAY_PARTNER_ID paser error, tmp value(%s)", v)
			return false
		}
		ALIPAY_PARTNER_ID = v
	} else {
		lprintf(1, "[ERR ] ALIPAY_PARTNER_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("ALIPAY_KEY", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] ALIPAY_KEY paser error, tmp value(%s)", v)
			return false
		}
		ALIPAY_KEY = v
	} else {
		lprintf(1, "[ERR ] ALIPAY_KEY  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("ALIPAY_OPER_URL", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] ALIPAY_OPER_URL paser error, tmp value(%s)", v)
			return false
		}
		ALIPAY_OPER_URL = v
	} else {
		lprintf(1, "[ERR ] ALIPAY_OPER_URL  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("PAYPAL_PAYMENT_URL", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYPAL_PAYMENT_URL paser error, tmp value(%s)", v)
			return false
		}
		PAYPAL_PAYMENT_URL = v
	} else {
		lprintf(1, "[ERR ] PAYPAL_PAYMENT_URL  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("PAYPAL_REFUND_URL", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYPAL_REFUND_URL paser error, tmp value(%s)", v)
			return false
		}
		PAYPAL_REFUND_URL = v
	} else {
		lprintf(1, "[ERR ] PAYPAL_REFUND_URL  NOT SETTING")
		return false
	}
	v, r = cls.GetTokenValue("PAYPAL_ID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYPAL_ID paser error, tmp value(%s)", v)
			return false
		}
		PAYPAL_ID = v
	} else {
		lprintf(1, "[ERR ] PAYPAL_ID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("PAYMENT_REDIRECT_BO", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYMENT_REDIRECT_BO paser error, tmp value(%s)", v)
			return false
		}
		PAYMENT_REDIRECT_BO = v
	} else {
		lprintf(1, "[ERR ] PAYMENT_REDIRECT_BO  NOT SETTING")
		return false
	}
	v, r = cls.GetTokenValue("PAYMENT_REDIRECT_FE", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYMENT_REDIRECT_FE paser error, tmp value(%s)", v)
			return false
		}
		PAYMENT_REDIRECT_FE = v
	} else {
		lprintf(1, "[ERR ] PAYMENT_REDIRECT_FE  NOT SETTING")
		return false
	}
	v, r = cls.GetTokenValue("PAYMENT_FAIL_REDIRECT_FE", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYMENT_FAIL_REDIRECT_FE paser error, tmp value(%s)", v)
			return false
		}
		PAYMENT_FAIL_REDIRECT_FE = v
	} else {
		lprintf(1, "[ERR ] PAYMENT_FAIL_REDIRECT_FE  NOT SETTING")
		return false
	}
	v, r = cls.GetTokenValue("PAYMENT_FAIL_REDIRECT_BO", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] PAYMENT_FAIL_REDIRECT_BO paser error, tmp value(%s)", v)
			return false
		}
		PAYMENT_FAIL_REDIRECT_BO = v
	} else {
		lprintf(1, "[ERR ] PAYMENT_FAIL_REDIRECT_BO  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("INICIS_MID", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] INICIS_MID paser error, tmp value(%s)", v)
			return false
		}
		INICIS_MID = v
	} else {
		lprintf(1, "[ERR ] INICIS_MID  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("INICIS_SIGN_KEY", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] INICIS_SIGN_KEY paser error, tmp value(%s)", v)
			return false
		}
		INICIS_SIGN_KEY = v
	} else {
		lprintf(1, "[ERR ] INICIS_SIGN_KEY  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("INICIS_KEY", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] INICIS_KEY paser error, tmp value(%s)", v)
			return false
		}
		INICIS_KEY = v
	} else {
		lprintf(1, "[ERR ] INICIS_KEY  NOT SETTING")
		return false
	}
	v, r = cls.GetTokenValue("INICIS_IV", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] INICIS_IV paser error, tmp value(%s)", v)
			return false
		}
		INICIS_IV = v
	} else {
		lprintf(1, "[ERR ] INICIS_IV  NOT SETTING")
		return false
	}

	v, r = cls.GetTokenValue("WHOXY_KEY", fileName)
	if r != cls.CONF_ERR {
		if len(v) == 0 {
			lprintf(1, "[ERR ] WHOXY_KEY paser error, tmp value(%s)", v)
			return false
		}
		WHOXY_KEY = v
	} else {
		lprintf(1, "[ERR ] WHOXY_KEY  NOT SETTING")
		return false
	}
	//////

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
			SkipURL = "/v1/extra/paypal/payment-return /v1/extra/alipay/return /v1/inicisReturn "
		}
		SkipURL = v
	} else {
		SkipURL = "/v1/extra/paypal/payment-return /v1/extra/alipay/return /v1/inicisReturn "
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

// 스케줄러 (토큰 갱신)
func SchedulerLoop() {
	for {
		time.Sleep(time.Duration(Token_TTL-300) * time.Second)
		if !SetApiToken() {
			SetApiToken()
		}
	}
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
