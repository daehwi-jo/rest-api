package model

var AliPayPartnerId string
var AliPayKey string
var AliOperUrl string
var PaymentRedirect string

var PaymentReturn string

type AlipayParameters struct {
	PageType     string `json:"pageType"`
	Service      string `json:"service"`        // 인터페이스 이름 - alipay 요청 시 사용
	Partner      string `json:"partner"`        // 제휴자 ID - alipay 요청 시 사용
	InputCharset string `json:"_input_charset"` // 网站编码
	NotifyUrl    string `json:"notify_url"`     // 비동기 알림 페이지
	ReturnUrl    string `json:"return_url"`
	ReferUrl     string `json:"refer_url"`
	Currency     string `json:"currency"`     // USD
	OutTradeNo   string `json:"out_trade_no"` // 주문번호
	Subject      string `json:"subject"`      // 제목
	TotalFee     string `json:"total_fee"`    // 가격 - alipay 요청 시 사용
	Body         string `json:"body"`         // 订单描述
	ProductCode  string `json:"product_code"`
	GmtReturn    string `json:"gmt_return"`
	ReturnAmount string `json:"return_amount"`
	OutReturnNo  string `json:"out_return_no"`

	SignType string  `json:"sign_type`
	Sign     string  `json:"sign`
	Total    float64 `json:"total"` // 가격 - 승인/취소 화면 요청
	// OrderId  string  `json:"orderId"`  // 주문번호 - 취소 화면 요청
}

type RespGetAlipayPaymentRefundOne struct {
	IsSuccess string `xml:"is_success"`
	Error     string `xml:"error"`
}

type InicisReturn struct {
	ResultCode    string `json:"resultCode"`
	ResultMsg     string `json:"resultMsg"`
	Tid           string `json:"tid"`
	GoodName      string `json:"goodName"`
	EventCode     string `json:"EventCode"`
	TotPrice      string `json:"TotPrice"`
	Moid          string `json:"MOID"`
	PayMethod     string `json:"payMethod"`
	ApplNum       string `json:"applNum"`
	ApplDate      string `json:"applDate"`
	ApplTime      string `json:"applTime"`
	BuyerEmail    string `json:"buyerEmail"`
	CustEmail     string `json:"custEmail"`
	BuyerTel      string `json:"buyerTel"`
	BuyerName     string `json:"buyerName"`
	VactNum       string `json:"VACT_Num"`
	VactBankCode  string `json:"VACT_BankCode"`
	VactName      string `json:"VACT_Name"`
	VactInputName string `json:"VACT_InputName"`
	VactDate      string `json:"VACT_Date"`
	VactTime      string `json:"VACT_Time"`
	VactBankName  string `json:"vactBankName"`
}

type InicisReturnMobile struct {
	Stats      string `json:"P_STATUS"`
	ResultMsg  string `json:"P_RMESG1"`
	Tid        string `json:"P_TID"`
	Type       string `json:"P_TYPE"`
	AuthDate   string `json:"P_AUTH_DT"`
	Mid        string `json:"P_MID"`
	Oid        string `json:"P_OID"`
	Amt        string `json:"P_AMT"`
	Uname      string `json:"P_UNAME"`
	Mname      string `json:"P_MNAME"`
	Noti       string `json:"P_NOTI"`
	NotieUrl   string `json:"P_NOTEURL"`
	NextUrk    string `json:"P_NEXT_URL"`
	VNum       string `json:"P_VACT_NUM"`
	VDate      string `json:"P_VACT_DATE"`
	VTime      string `json:"P_VACT_TIME"`
	VName      string `json:"P_VACT_NAME"`
	VBankeCode string `json:"P_VACT_BANK_CODE"`
}

type RespGetAlipayReturn struct {
	OrderId                      string  `json:"orderId"`
	Total                        float64 `json:"total"`
	Status                       string  `json:"status"`
	Sign                         string  `json:"sign"`
	TradeNumber                  string  `json:"tradeNumber"`
	Currency                     string  `json:"currency"`
	SignType                     string  `json:"signType"`
	ReturnUrl                    string  `json:"returnUrl"`
	PayMethod                    string  `json:"payMethod"`
	VirtualAccountBankCode       string  `json:"virtualAccountBankCode"`
	VirtualAccountBankName       string  `json:"virtualAccountBankName"`
	VirtualAccountNo             string  `json:"virtualAccountNo"`
	VirtualAccountExpireDateTime string  `json:"virtualAccountExpireDateTime"`
	PgData                       string  `json:"pgData"`
	VBankStatus                  string  `json:"VBankStatus"`
}

// GetAlipayPaymentRequestAll
type RespGetAlipayPaymentRequestAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Service      string `json:"service"`        // 인터페이스 이름 - alipay 요청 시 사용
		Partner      string `json:"partner"`        // 제휴자 ID - alipay 요청 시 사용
		InputCharset string `json:"_input_charset"` // 网站编码
		NotifyUrl    string `json:"notify_url"`     // 비동기 알림 페이지
		ReturnUrl    string `json:"return_url"`
		ReferUrl     string `json:"refer_url"`
		OutTradeNo   string `json:"out_trade_no"` // 주문번호
		TotalFee     string `json:"total_fee"`    // 가격 - alipay 요청 시 사용
		Body         string `json:"body"`         // 订单描述
		ProductCode  string `json:"product_code"`
		SignType     string `json:"sign_type`
		Sign         string `json:"sign`
	} `json:"data"`
}
