package model

const (
	C500       = "C500"
	C404       = "C404"
	FAIL       = "0099"
	SUCCESS    = "0000"
	RETRY      = "1000"
	VALIDATION = "2000"
)

const (
	SUCCMESSAGE  = "정상"
	FAILMESSAGE  = "비정상"
	NOTFOUND     = "NOTFOUND"
	TYPE         = "EXTRA"
	RETRYMESSAGE = "RETRY"
)

var (
	ApiKey           string
	AuthUserId       string
	WhoisUrl         string
	WhoisCheckUrl    string
	CustomerId       string
	RegContactId     string
	AdminContactId   string
	TechContactId    string
	BillingContactId string
	InvoiceOption    string
	ProtectPrivate   string
	Years            string
	Mode             bool // 외부 api 접속 시 dev - false or live - true
	Ns               []string
)

type MessageValue struct {
	MessageType string `json:"mtype"`
	Kr          string `json:"kr"`
	En          string `json:"en"`
	Cn          string `json:"cn"`
}

type TransferStatusResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        string       `json:"data"`
}

// API CALL RESPONSE
type ApiCallResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        interface{}  `json:"data"`
}

type DomainDeleteResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
}

///////////////////////////////////////
type ApiCallResponseInterface struct {
	Data interface{} `json:"data"`
}

type DomainEntityApiCallResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		EntityId string `json:"entityId"`
	} `json:"data"`
}

type BalanceApiCallResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		AvailableBalanceValue string `json:"availableBalanceValue"`
	} `json:"data"`
}

// GetExtraApiAll - extrpapi.go
type GetExtraApi struct {
	DomainName string `json:"domainName"`
	RegisterYn string `json:"registerYn"`
}

type RespGetExtraApiAll struct {
	Code        string        `json:"code"`
	Message     MessageValue  `json:"message"`
	ServiceName string        `json:"serviceName"`
	Data        []GetExtraApi `json:"data"`
}

// 후이즈에 인증코드 제출
type RespSetSubmitAuthCodeWhoisOne struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

type RespGetDomainDetailOrderIdAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Domsecret string `json:"domsecret"`
	} `json:"data"`
}

type RespGetDomainExpireDataOne struct {
	EndTime string `json:"endtime"`
}

type OrderIdParam struct {
	OrderId int `json:"orderId"`
}

type RespSetDomainRegisterWhois struct {
	Data string `json:"data"`
	// Status           string `json:"status"`
	// Error            string `json:"error"`
	// ActionTypeDesc   string `json:"actionTypeDesc"`
	// ActionStatus     string `json:"actionStatus"`
	// EntityId         string `json:"entityId"`
	// EaqId            string `json:"eaqId"`
	// ActionType       string `json:"actionType"`
	// Description      string `json:"description"`
	// ActionStatusDesc string `json:"actionStatusDesc"`

	/*
		성공일때 who is 응답 형태
		"actiontypedesc":"Registration of wklqdnqwjkdnkqwd.com for 1 year",
		"actionstatus":"Success",
		"entityid":"93606499",
		"status":"Success",
		"eaqid":"600316194",
		"actiontype":"AddNewDomain",
		"description":"wklqdnqwjkdnkqwd.com",
		"actionstatusdesc":"Domain registration completed Successfully"
	*/
}

type RespGetDomainAvailableWhoisAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		DomainList []GetDomainAvailableWhoisAll `json:"data"`
	} `json:"data"`
}

type GetDomainAvailableWhoisAll struct {
	DomainName string `json:"domainName"`
	Status     string `json:"status"`
}

type ABC struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Classkey string `json:"classkey"`
}

type RespGetDomainRecommendWhoisAll struct {
	Data string `json:"data"`
}

type RespGetDomainDetailWhoisAll struct {
	/*
		Status  string `json:"status"`
		Message string `json:"message"`
	*/
	Data string `json:"data"`
}

type WhoisApiCallResponse struct {
	Code        string      `json:"code"`
	Message     string      `json:"message"`
	ServiceName string      `json:"serviceName"`
	Data        interface{} `json:"data"`
}

type searchOrders struct {
	orderId           int    `json:"orderId"`
	autorenew         bool   `json:"autorenew"`
	endtime           int    `json:"endtime"`
	resellerlock      bool   `json:"resellerlock"`
	timeStamp         string `json:"timeStamp"`
	customerLock      bool   `json:"customerLock"`
	transferLock      bool   `json:"transferLock"`
	creationTime      int    `json:"creationTime"`
	privacyProtiction bool   `json:"privacyProtiction"`
	creationdt        int    `json:"creationdt"`
}
type searchEntity struct {
	customerId    int    `json:"customerId"`
	entityId      int    `json:"entityId"`
	entityTypeId  int    `json:"entityTypeId"`
	currentStatus string `json:"currentStatus"`
	description   string `json:"description"`
}
type searchEntityType struct {
	entityTypeKey  string `json:"entityTypeKey"`
	entityTypeName string `json:"entityTypeName"`
}
type RespGetDomainSearchWhoisAll struct {
	Data string `json:"data"`
	// Recsonpage     string `json:"recsonpage"`
	// Recsindb       string `json:"recsindb"`
	// RecsonpageData struct {
	// 	orders     searchOrders     `json:"orders"`
	// 	entity     searchEntity     `json:"entity"`
	// 	entityType searchEntityType `json:"entityType"`
	// } `json:"2"`
	// RecsindbData   struct {
	// 	orders     searchOrders     `json:"orders"`
	// 	entity     searchEntity     `json:"entity"`
	// 	entityType searchEntityType `json:"entityType"`
	// } `json:"1"`
}

type RespAvailableBalance struct {
	AvailableBalanceValue float64 `json:"sellingcurrencybalance"` //예치금으로 보내야 하는 값
	CurrencySymbol        string  `json:"sellingcurrencysymbol"`
	CurrencyLockedBalance float64 `json:"sellingcurrencylockedbalance"`
}
type RespAvailableBalanceValue struct {
	AvailableBalanceValue float64 `json:"sellingcurrencybalance"` //예치금으로 보내야 하는 값
}

type RespModiDomainModifyNsWhoisOne struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	} `json:"data"`
	// ActionTypeDesc   string `json:"actionTypeDesc"`
	// EntityId         string `json:"entityId"`
	// ActionStatus     string `json:"actionStatus"`
	// EaqId            string `json:"eaqId"`
	// CurrentAction    string `json:"currentAction"`
	// Description      string `json:"description"`
	// ActionType       string `json:"actionType"`
	// ActionStatusDesc string `json:"actionStatusDesc"`
}
type MessageOne struct {
	Message string `json:"message"` //메세지
}

type RespGetDomainOrderIdWhoisOne struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Status  string `json:"status"`
		OrderId int    `json:"orderId"`
	} `json:"data"`
}

type RespSetDomainAddChildNameServerAll struct {
	Data string `json:"data"`
}

type RespModiHostChildNameServerAll struct {
	Data string `json:"data"`
}

type RespModiIpChildNameServerOne struct {
	Data string `json:"data"`
}

type RespDelChildNameServerAll struct {
	Data string `json:"data"`
}

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

type RespSetDomainRestoreWhoisOne struct {
	Data string `json:"data"`
}

type RespDelDomainNameWhoisOne struct {
	Data string `json:"data"`
}

type RespGetDomainValidateTransferOne struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Status     string `json:"status"`
		DomainName string `json:"DomainName"`
	} `json:"data"`
}

type RespSetDomainValidateTransferOne struct {
	Data string `json:"data"`
}

type RespModiAuthCodeWhoisOne struct {
	Data string `json:"data"`
}

type RespSetResendRfaWhoisOne struct {
	Data string `json:"data"`
}

type RespDelCancelTransferWhoisOne struct {
	Data string `json:"data"`
}

type RespSetResendVerificationWhoisOne struct {
	Data string `json:"data"`
}

type RespModiContactWhoisOne struct {
	Data string `json:"data"`
}

type RespGGetDomainIdnAvailableWhoisAll struct {
	Data string `json:"data"`
}

type RespGetDomainPremiumWhoisAll struct {
	Data string `json:"data"`
}

type RespGetThirdLevelNameWhoisAll struct {
	Data string `json:"data"`
}

type RespGetUkWhoisAll struct {
	Data string `json:"data"`
}

type RespGetPremiumCheckWhoisOne struct {
	Data string `json:"data"`
}

type RespGetCustomerDefaultNsWhoisOne struct {
	Data string `json:"data"`
}

type RespSetPurchasePrivacyWhois struct {
	Data string `json:"data"`
}

type RespModiPrivacyProtectionWhoisOne struct {
	Data string `json:"data"`
}

type RespSetEnableTheftProtectionWhoisOne struct {
	Data string `json:"data"`
}

type RespSetDisableTheftProtectionWhoisOne struct {
	Data string `json:"data"`
}

type RespGetLockWhoisOne struct {
	Data string `json:"data"`
}

type RespModiTelWhoisPrefWhoisOne struct {
	Data string `json:"data"`
}

type RespSetUkWhoisOne struct {
	Data string `json:"data"`
}

type RespGetRecheckNsWhoisOne struct {
	Data string `json:"data"`
}

type RespGetDotxxxAssociationDetailsOne struct {
	Data string `json:"data"`
}

type RespSetDnsSecWhoisOne struct {
	Data string `json:"data"`
}

type RespDelDnsSecWhoisOne struct {
	Data string `json:"data"`
}

type RespSetPreorderingWhoisOne struct {
	Data string `json:"data"`
}

type RespGetPreorderingWhoisOne struct {
	Data string `json:"data"`
}

type RespGetPreorderingCategoryWhoisAll struct {
	Data string `json:"data"`
}

type RespGetTmNoticeWhoisOne struct {
	Data string `json:"data"`
}

type RespGetTldsInPhaseWhoisAll struct {
	Data string `json:"data"`
}

type RespGetTldInfoWhoisOne struct {
	Data string `json:"data"`
}

type RespSetCustomerV2WhoisOne struct {
	Data string `json:"data"`
}

type WhoisDomainRefund struct {
	OrderID int `json:"orderID"`
}

type RespDelDomainNameWhois struct {
	Status string `json:"status"`
}

// 도메인 상세 정보 가져오기
type RespGetDomainInfoWhoxy struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Data string `json:"data"`
	} `json:"data"`
}

type RespWhoxy struct {
	RawWhois string `json:"raw_whois"`
}
