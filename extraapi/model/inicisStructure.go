package model

// API CALL RESPONSE
type InicisResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        []byte       `json:"data"`
}

type RespGetInicisPaymentModule struct {
	// Data []byte `json:"data"`
	Data string `json:"data"`
}

type RespGetInicisVirtualAccount struct {
	Data string `json:"data"`
}

// SetPaymentCancelOne
type PaymentCancelOneParam struct {
	ProviderCompanyIndex string  `json:"providerCompanyIndex"`
	UserIndex            string  `json:"userIndex"`
	LoginId              string  `json:"loginId"`
	ManageNo             string  `json:"manageNo"`
	Status               string  `json:"status"`
	CurrencyType         string  `json:"currencyType"`
	PaymentWay           string  `json:"paymentWay"`
	PgApprovalNo         string  `json:"pgApprovalNo"`
	Price                float64 `json:"price"`
	RefundAcctNum        string  `json:"refundAcctNum"`
	RefundBankCode       string  `json:"refundBankCode"`
	RefundAcctName       string  `json:"refundAcctName"`
}

// alipay refund
type AlipayRefundParam struct {
	Service      string `json:"service"`
	Partner      string `json:"partner"`
	InputCharset string `json:"_input_charset"`
	OutReturnNo  string `json:"out_return_no"`
	OutTradeNo   string `json:"out_trade_no"`
	ReturnAmount string `json:"return_amount"`
	GmtReturn    string `json:"gmt_return"`
	Currency     string `json:"currency"`
}

type RespGetInicisPaymentRefund struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		ResultCode    string `json:"resultCode"`
		ResultMsg     string `json:"resultMsg"`
		CancelDate    string `json:"cancelDate"`
		CancelTime    string `json:"cancelTime"`
		CshrCancelNum string `json:"cshrCancelNum"`
	} `json:"data"`
}

type RespGetInicisCashReceipt struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		ResultCode      string `json:"resultCode"`
		ResultMsg       string `json:"resultMsg"`
		Tid             string `json:"tid"`
		AuthDate        string `json:"authDate"`
		AuthTime        string `json:"authTime"`
		AuthCode        string `json:"authCode"`
		AuthNo          string `json:"authNo"`
		AuthPrice       string `json:"authPrice"`
		AuthSupplyPrice string `json:"authSupplyPrice"`
		AuthTax         string `json:"authTax"`
		AuthSrvcPrice   string `json:"authSrvcPrice"`
		AuthUseOpt      string `json:"authUseOpt"`
	} `json:"data"`
}
