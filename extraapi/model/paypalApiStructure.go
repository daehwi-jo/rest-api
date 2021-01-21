package model

var PaypalTokenType string
var PaypalAccessToken string

type GetPaypalAccessTokenOne struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppId       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

/////////////////////////////////////////////
/////////////////////////////////////////////
/////////////////////////////////////////////
type PaypalRefundParam struct {
	PgApprovalNo string `json:"pgApprovalNo"`
}

type PaypalPaymentParam struct {
	PageType string `json:"pageType"`
	ManageNo string `json:"manageNo"`
}
type PaypalReturnParam struct {
	OrderId     string `json:"orderId"`
	TradeNumber string `json:"tradeNumber"`
}

// paypal request
type ItemInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Tax         string `json:"tax"`
	Sku         string `json:"sku"`
	Currency    string `json:"currency"`
}

type RespPaypalPayment struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		ReqUrl       string `json:"reqUrl"`
		Cmd          string `json:"cmd"`
		Business     string `json:"business"`
		ReturnUrl    string `json:"returnUrl"`
		NotifyUrl    string `json:"notifyUrl"`
		CancelReturn string `json:"cancelReturn"`
		Charset      string `json:"charset"`
		CurrencyType string `json:"currencyType"`
	} `json:"data"`
}

type TransactionInfo struct {
	Amount struct {
		Total    string `json:"total"`
		Currency string `json:"currency"`
		Detail   struct {
			SubTotal        string `json:"subtotal"`
			Tax             string `json:"tax"`
			Shipping        string `json:"shipping"`
			HandlingFee     string `json:"handling_fee"`
			ShppingDiscount string `json:"shipping_discount"`
			Insurance       string `json:"insurance"`
		} `json:"details"`
	} `json:"amount"`
	Description   string `json:"description"`
	Custom        string `json:"custom"`
	InvoiceNumber string `json:"invoice_number"`
	ItemList      struct {
		Items           []ItemInfo `json:"items"`
		ShippingAddress struct {
			RecipientName string `json:"recipient_name"`
			Line1         string `json:"line1"`
			Line2         string `json:"line2"`
			City          string `json:"city"`
			CountryCode   string `json:"country_code"`
			PostalCode    string `json:"postal_code"`
			Phone         string `json:"phone"`
			State         string `json:"state"` // 현재는 미사용
		} `json:"shipping_address"`
	} `json:"item_list"`
}

type RequestPaypalPayment struct {
	Intent string `json:"intent"`
	Payer  struct {
		PaymentMethod string `json:"payment_method"`
	} `json:"payer"`
	ApplicationContext struct {
		BrandName   string `json:"brand_name"`
		Locale      string `json:"locale"`
		LandingPage string `json:"landing_page"`
	} `json:"application_context"`
	Transactions []TransactionInfo `json:"transactions"`
	NoteToPayer  string            `json:"note_to_payer"`
	RedirectUrls struct {
		ReturnUrl string `json:"return_url"`
		CancelUrl string `json:"cancel_url"`
	} `json:"redirect_urls"`
}

/////////////////////////////////////////////
/////////////////////////////////////////////
/////////////////////////////////////////////
// paypal response
type RespPaypalItems struct {
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Currency    string `json:"currency"`
	Tax         string `json:"tax"`
	Quantity    int    `json:"quantity"`
}

type RespRelatedResource struct {
	Authorization struct {
		Id         string `json:"id"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
		Amount     struct {
			Total    string `json:"total"`
			Currency string `json:"currency"`
		} `json:"amount"`
		ParentPayment string            `json:"parent_payment"`
		ValidUntil    string            `json:"valid_until"`
		Links         []RespPaypalLinks `json:"links"`
	} `json:"authorization"`
}

type RespPaypalTransaction struct {
	Amount struct {
		Total    string `json:"total"`
		Currency string `json:"currency"`
		Details  struct {
			SubTotal         string `json:"subtotal"`
			Tax              string `json:"tax"`
			Shipping         string `json:"shipping"`
			HandlingFee      string `json:"handling_fee"`
			Insurance        string `json:"insurance"`
			ShippingDiscount string `json:"shipping_discount"`
		} `json:"details"`
	} `json:"amount"`
	Description    string `json:"description"`
	Custom         string `json:"custom"`
	InvoiceNumber  string `json:"invoice_number"`
	SoftDescriptor string `json:"soft_descriptor"`
	// test 응답시 받음
	PaymentOptions struct {
		AllowedPaymentMethod string `json:"allowed_payment_method"`
		RecurringFlag        bool   `json:"recurring_flag"`
		SkipFmf              bool   `json:"skip_fmf"`
	} `json:"payment_options"`
	ItemList struct {
		Items           []RespPaypalItems `json:"items"`
		ShippingAddress struct {
			RecipientName string `json:"recipient_name"`
			Line1         string `json:"line1"`
			Line2         string `json:"line2"`
			City          string `json:"city"`
			State         string `json:"state"`
			PostalCode    string `json:"postal_code"`
			CountryCode   string `json:"country_code"`
			Phone         string `json:"phone"`
		} `json:"shipping_address"`
		RelatedResource []RespRelatedResource `json:"related_resources"`
	} `json:"item_list"`
}

// type RespPaypalTransaction struct {
// 	Amount struct {
// 		Total    string `json:"total"`
// 		Currency string `json:"currency"`
// 		Details  struct {
// 			SubTotal         string `json:"subtotal"`
// 			Tax              string `json:"tax"`
// 			Shipping         string `json:"shipping"`
// 			HandlingFee      string `json:"handling_fee"`
// 			Insurance        string `json:"insurance"`
// 			ShippingDiscount string `json:"shipping_discount"`
// 		} `json:"details"`
// 	} `json:"amount"`
// 	Description    string `json:"description"`
// 	Custom         string `json:"custom"`
// 	InvoiceNumber  string `json:"invoice_number"`
// 	SoftDescriptor string `json:"soft_descriptor"`
// 	// test 응답시 받음
// 	PaymentOptions struct {
// 		AllowedPaymentMethod string `json:"allowed_payment_method"`
// 		RecurringFlag        bool   `json:"recurring_flag"`
// 		SkipFmf              bool   `json:"skip_fmf"`
// 	} `json:"payment_options"`
// 	ItemList struct {
// 		Items           []RespPaypalItems `json:"items"`
// 		ShippingAddress struct {
// 			RecipientName string `json:"recipient_name"`
// 			Line1         string `json:"line1"`
// 			Line2         string `json:"line2"`
// 			City          string `json:"city"`
// 			State         string `json:"state"`
// 			PostalCode    string `json:"postal_code"`
// 			CountryCode   string `json:"country_code"`
// 			Phone         string `json:"phone"`
// 		} `json:"shipping_address"`
// 		RelatedResource []RespRelatedResource `json:"related_resources"`
// 	} `json:"item_list"`
// }

type RespPaypalLinks struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type RespGetPaypalPaymentRequestAll struct {
	// 테스트 응답을 기준으로 key 정령
	Id     string `json:"id"`
	Intent string `json:"intent"`
	State  string `json:"state"`
	Payer  struct {
		PaymentMethod string `json:"payment_method"`
	} `json:"payer"`
	Transactions []RespPaypalTransaction `json:"transactions"`
	NoteToPayer  string                  `json:"note_to_payer"`
	CreateTime   string                  `json:"create_time"`
	Links        []RespPaypalLinks       `json:"links"`
	UpdateTime   string                  `json:"update_time"` // Test 시 응답 데이터 없음
}

/////////////////////////////////////////////
/////////////////////////////////////////////
/////////////////////////////////////////////

type RespGetPaypalPaymentReturnOne struct {
	Data string `json:"data"`
}

type RespDelPaypalRefundRequestOne struct {
	Data string `json:"data"`
}

type RespGetPaypalPaymentCancelOne struct {
	Data string `json:"data"`
}
