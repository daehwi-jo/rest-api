package model

type DomainNameParam struct {
	DomainName string `json:"domainName"`
}

type HostNameParam struct {
	HostName string `json:"host_name"`
}
type SetDomainRegisterWhoisOneParam struct {
	DomainName string   `json:"domainName"`
	YearNum    string   `json:"yearNum"`
	NsArray    []string `json:"ns"`
}

type GetDomainAvailableWhoisAllParam struct {
	DomainName []string `json:"domainName"`
	Tlds       []string `json:"tlds"`
}

type GetDomainDetailWhoisAllParam struct {
	DomainName string   `json:"domainName"`
	Options    []string `json:"options"`
}

type GetDomainSearchWhoisAllParam struct {
	// 파라미터 참고 https://demo.myorderbox.com/kb/node/771
	// 그 외 파라미터 사용시 추가 및 surl 수정 필요
	NoOfRecords int `json:"noOfRecords"`
	PageNo      int `json:"pageNo"`
}

type GetAvailableBalanceParam struct {
	ResellerId int `json:"resellerId"`
}

type GetDomainAvailableSunriseWhoisOneParam struct {
	DomainName string `json:"domainName"`
	Tld        string `json:"tld"`
	Smd        string `json:"smd"`
}

type ModiDomainModifyNsWhoisOneParam struct {
	OrderId int      `json:"orderId"`
	Ns      []string `json:"ns"`
}

type SetDomainAddChildNameServerAllParam struct {
	OrderId int      `json:"orderId"`
	Cns     string   `json:"cns"`
	Ip      []string `json:"ip"`
}

type ModiHostChildNameServerAllParam struct {
	OrderId int    `json:"orderId"`
	OldCns  string `json:"oldCns"`
	NewCns  string `json:"newCns"`
}

type ModiIpChildNameServerOneParam struct {
	OrderId int    `json:"orderId"`
	Cns     string `json:"cns"`
	OldIp   string `json:"oldIp"`
	NewIp   string `json:"newIp"`
}

type DelChildNameServerAllParam struct {
	OrderId int      `json:"orderId"`
	Cns     string   `json:"cns"`
	Ip      []string `json:"ip"`
}

type GetDomainDetailOrderIdAllParam struct {
	OrderId int      `json:"orderId"`
	Options []string `json:"options"`
}

type SetDomainRenewWhoisOneParam struct {
	OrderId        int     `json:"orderId"`
	Years          int     `json:"years"`
	ExpDate        int     `json:"expDate"`
	InvoiceOption  string  `json:"invoiceOption"`
	DiscountAmount float64 `json:"discountAmount"`
}

type SetDomainRestoreWhoisOneParam struct {
	OrderId       int    `json:"orderId"`
	InvoiceOption string `json:"invoiceOption"`
}

type SetDomainValidateTransferOneParam struct {
	DomainName string   `json:"domainName"`
	AuthCode   string   `json:"authCode"`
	Ns         []string `json:"ns"`
	// CustomerId       int      `json:"customerId"`
	// RegContactId     int      `json:"regContactId"`
	// AdminContact     int      `json:"adminContact"`
	// TechContactId    int      `json:"techContactId"`
	// BillingContactId int      `json:"billingContactId"`
	// InvoiceOption    string   `json:"invoiceOption"`
	PurchasePrivacy string `json:"purchasePrivacy"`
}

type ModiAuthCodeWhoisOneParam struct {
	OrderId  int    `json:"orderId"`
	AuthCode string `json:"authCode"`
}

type ModiContactWhoisOneParam struct {
	OrderId          int `json:"orderId"`
	RegContactId     int `json:"regContactId"`
	AdminContactId   int `json:"adminContactId"`
	TechContactId    int `json:"techContactId"`
	BillingContactId int `json:"billingContactId"`
}

type GetDomainIdnAvailableWhoisAllParam struct {
	DomainName      []string `json:"domainName"`
	Tld             string   `json:"tld"`
	IdnLanguageCode string   `json:"idnLanguageCode"`
}

type GetDomainPremiumWhoisAllParam struct {
	KeyWord    string   `json:"keyWord"`
	Tlds       []string `json:"tlds"`
	PriceLow   int      `json:"priceLow"`
	NoOfResult int      `json:"noOfResult"`
}

type GetThirdLevelNameWhoisAllParam struct {
	DomainName []string `json:"domainName"`
	Tlds       []string `json:"tlds"`
}

type GetUkWhoisAllParam struct {
	DomainName []string `json:"domainName"`
	Tlds       []string `json:"tlds"`
}

type SetPurchasePrivacyWhoisParam struct {
	OrderId        int     `json:"orderId"`
	InvoiceOption  string  `json:"invoiceOption"`
	DiscountAmount float64 `json:"discountAmount"`
}

type ModiPrivacyProtectionWhoisOneParam struct {
	OrderId        int    `json:"orderId"`
	ProtectPrivacy string `json:"protectPrivacy"`
	Reason         string `json:"reason"`
}

type ModiTelWhoisPrefWhoisOneParam struct {
	OrderId   int    `json:"orderId"`
	WhoisType string `json:"whoisType"`
	Publish   string `json:"publish"`
}

type SetUkWhoisOneParam struct {
	OrderId int    `json:"orderId"`
	NewTag  string `json:"newTag"`
}

type GetPreorderingWhoisOneParam struct {
	CustomerId  int `json:"customerId"`
	PageNo      int `json:"pageNo"`
	NoOfRecords int `json:"noOfRecords"`
}

type SetGabiaDomainParam struct {
	DomainName   string   `json:"domainName"`
	Period       int      `json:"period"`
	AutoLockYn   string   `json:"autoLockYn"`
	HanName      string   `json:"hanName"`
	EngName      string   `json:"engName"`
	HanOrg       string   `json:"hanOrg"`
	EngOrg       string   `json:"engOrg"`
	Gubun        string   `json:"gubun"`
	CtfyCode     string   `json:"ctfyCode"`
	CtfyNo       string   `json:"ctfyNo"`
	OpenInfoYn   string   `json:"openInfoYn"`
	Zip          string   `json:"zip"`
	Addr1        string   `json:"addr1"`
	Addr2        string   `json:"addr2"`
	Eaddr1       string   `json:"eaddr1"`
	Eaddr2       string   `json:"eaddr2"`
	City         string   `json:"city"`
	CountryCode  string   `json:"countryCode"`
	Phone        string   `json:"phone"`
	Fax          string   `json:"fax"`
	Hp           string   `json:"hp"`
	Email        string   `json:"email"`
	AhanName     string   `json:"ahanName"`
	AengName     string   `json:"aengName"`
	Azip         string   `json:"azip"`
	Aaddr1       string   `json:"aaddr1"`
	Aaddr2       string   `json:"aaddr2"`
	Aeaddr1      string   `json:"aeaddr1"`
	Aeaddr2      string   `json:"aeaddr2"`
	Acity        string   `json:"acity"`
	AcountryCode string   `json:"acountryCode"`
	Aphone       string   `json:"aphone"`
	Afax         string   `json:"afax"`
	Ahp          string   `json:"ahp"`
	SmsYn        string   `json:"smsYn"`
	Aemail       string   `json:"aemail"`
	Ns           []string `json:"ns"`
	NsIp         []string `json:"nsIp"`
}

type GetGabiaDomainParam struct {
	DomainName string `json:"domainName"`
}

type ModiGabiaDomainAdminParam struct {
	DomainName   string `json:"domainName"`
	AhanName     string `json:"ahanName"`
	AengName     string `json:"aengName"`
	Azip         string `json:"azip"`
	Aaddr1       string `json:"aaddr1"`
	Aaddr2       string `json:"aaddr2"`
	Aeaddr1      string `json:"aeaddr1"`
	Aeaddr2      string `json:"aeaddr2"`
	Acity        string `json:"acity"`
	AcountryCode string `json:"acountryCode"`
	Aphone       string `json:"aphone"`
	Afax         string `json:"afax"`
	Ahp          string `json:"ahp"`
	SmsYn        string `json:"smsYn"`
	Aemail       string `json:"aemail"`
}

type ModiGabiaDomainAdminJsonParam struct {
	DomainName string                   `json:"domain_name"`
	Data       ModiGabiaDomainAdminData `json:"data"`
}
type ModiGabiaDomainAdminData struct {
	AhanName     string `json:"a_hanname"`
	AengName     string `json:"a_engname"`
	Azip         string `json:"a_zip"`
	Aaddr1       string `json:"a_addr1"`
	Aaddr2       string `json:"a_addr2"`
	Aeaddr1      string `json:"a_eaddr1"`
	Aeaddr2      string `json:"a_eaddr2"`
	Acity        string `json:"a_city"`
	AcountryCode string `json:"a_countrycode"`
	Aphone       string `json:"a_phone"`
	Afax         string `json:"a_fax"`
	Ahp          string `json:"a_hp"`
	SmsYn        string `json:"sms_yn"`
	Aemail       string `json:"a_email"`
}

type ModiGabiaDomainNameServerParam struct {
	DomainName string   `json:"domainName"`
	Ns         []string `json:"ns"`
}

type ModiGabiaDomainNameServerJsonParam struct {
	DomainName string                            `json:"domain_name"`
	Data       ModiGabiaDomainNameServerJsonData `json:"data"`
}

type ModiGabiaDomainNameServerJsonData struct {
	Ns1    string `json:"ns1"`
	Ns2    string `json:"ns2"`
	Ns3    string `json:"ns3"`
	Ns4    string `json:"ns4"`
	Ns5    string `json:"ns5"`
	Ns6    string `json:"ns6"`
	Ns7    string `json:"ns7"`
	Ns8    string `json:"ns8"`
	Ns9    string `json:"ns9"`
	Ns10   string `json:"ns10"`
	Ns11   string `json:"ns11"`
	Ns12   string `json:"ns12"`
	Ns13   string `json:"ns13"`
	NsIP1  string `json:"nsip1"`
	NsIP2  string `json:"nsip2"`
	NsIP3  string `json:"nsip3"`
	NsIP4  string `json:"nsip4"`
	NsIP5  string `json:"nsip5"`
	NsIP6  string `json:"nsip6"`
	NsIP7  string `json:"nsip7"`
	NsIP8  string `json:"nsip8"`
	NsIP9  string `json:"nsip9"`
	NsIP10 string `json:"nsip10"`
	NsIP11 string `json:"nsip11"`
	NsIP12 string `json:"nsip12"`
	NsIP13 string `json:"nsip13"`
}

type ModiGabiaDomainOwnerJsonParam struct {
	DomainName string                       `json:"domain_name"`
	Data       ModiGabiaDomainOwnerJsonData `json:"data"`
}

type ModiGabiaDomainOwnerJsonData struct {
	HanName     string `json:"hanname"`
	EngName     string `json:"engname"`
	HanOrg      string `json:"hanorg"`
	EngOrg      string `json:"engorg"`
	Gubun       int    `json:"gubun"`
	CtfyCode    string `json:"ctfy_code"`
	CtfyNo      string `json:"ctfy_no"`
	OpenInfoYn  string `json:"open_info_yn"`
	Zip         string `json:"zip"`
	Addr1       string `json:"addr1"`
	Addr2       string `json:"addr2"`
	Eaddr1      string `json:"eaddr1"`
	Eaddr2      string `json:"eaddr2"`
	City        string `json:"city"`
	CountryCode string `json:"countrycode"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	Hp          string `json:"hp"`
	Email       string `json:"email"`
}

type ModiGabiaDomainOwnerParam struct {
	DomainName  string `json:"domainName"`
	HanName     string `json:"hanName"`
	EngName     string `json:"engName"`
	HanOrg      string `json:"hanOrg"`
	EngOrg      string `json:"engOrg"`
	Gubun       int    `json:"gubun"`
	CtfyCode    string `json:"ctfyCode"`
	CtfyNo      string `json:"ctfyNo"`
	OpenInfoYn  string `json:"openInfoYn"`
	Zip         string `json:"zip"`
	Addr1       string `json:"addr1"`
	Addr2       string `json:"addr2"`
	Eaddr1      string `json:"eaddr1"`
	Eaddr2      string `json:"eaddr2"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	Hp          string `json:"hp"`
	Email       string `json:"email"`
}

type ModiGabiaDomainStatusLockParam struct {
	DomainName string `json:"domainName"`
	AutoLockYn string `json:"autoLockYn"`
}

type SetGabiaDomainIncreaseParam struct {
	DomainName string `json:"domainName"`
	Period     int    `json:"peroid"`
}

type SetGabiaDomainTransferParam struct {
	DomainName   string `json:"domainName"`
	AuthInfo     string `json:"authInfo"`
	HanName      string `json:"hanName"`
	EngName      string `json:"engName"`
	HanOrg       string `json:"hanOrg"`
	EngOrg       string `json:"engOrg"`
	Gubun        string `json:"gubun"`
	CtfyCode     string `json:"ctfyCode"`
	CtfyNo       string `json:"ctfyNo"`
	OpenInfoYn   string `json:"openInfoYn"`
	Zip          string `json:"zip"`
	Addr1        string `json:"addr1"`
	Addr2        string `json:"addr2"`
	Eaddr1       string `json:"eaddr1"`
	Eaddr2       string `json:"eaddr2"`
	City         string `json:"city"`
	CountryCode  string `json:"countryCode"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Hp           string `json:"hp"`
	Email        string `json:"email"`
	AhanName     string `json:"ahanName"`
	AengName     string `json:"aengName"`
	Azip         string `json:"azip"`
	Aaddr1       string `json:"aaddr1"`
	Aaddr2       string `json:"aaddr2"`
	Aeaddr1      string `json:"aeaddr1"`
	Aeaddr2      string `json:"aeaddr2"`
	Acity        string `json:"acity"`
	AcountryCode string `json:"acountryCode"`
	Aphone       string `json:"aphone"`
	Afax         string `json:"afax"`
	Ahp          string `json:"ahp"`
	SmsYn        string `json:"smsYn"`
	Aemail       string `json:"aemail"`
}

type SetGabiaDomainTransferJsonParam struct {
	DomainName string                     `json:"domain_name"`
	Data       SetGabiaDomainTransferData `json:"data"`
}

type SetGabiaDomainTransferData struct {
	AuthInfo     string `json:"auth_info"`
	HanName      string `json:"hanname"`
	EngName      string `json:"engname"`
	HanOrg       string `json:"hanorg"`
	EngOrg       string `json:"engorg"`
	Gubun        string `json:"gubun"`
	CtfyCode     string `json:"ctfy_code"`
	CtfyNo       string `json:"ctfy_no"`
	OpenInfoYn   string `json:"open_info_yn"`
	Zip          string `json:"zip"`
	Addr1        string `json:"addr1"`
	Addr2        string `json:"addr2"`
	Eaddr1       string `json:"eaddr1"`
	Eaddr2       string `json:"eaddr2"`
	City         string `json:"city"`
	CountryCode  string `json:"countrycode"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Hp           string `json:"hp"`
	Email        string `json:"email"`
	AhanName     string `json:"a_hanname"`
	AengName     string `json:"a_engname"`
	Azip         string `json:"a_zip"`
	Aaddr1       string `json:"a_addr1"`
	Aaddr2       string `json:"a_addr2"`
	Aeaddr1      string `json:"a_eaddr1"`
	Aeaddr2      string `json:"a_eaddr2"`
	Acity        string `json:"a_city"`
	AcountryCode string `json:"a_countrycode"`
	Aphone       string `json:"a_phone"`
	Afax         string `json:"a_fax"`
	Ahp          string `json:"a_hp"`
	SmsYn        string `json:"sms_yn"`
	Aemail       string `json:"a_email"`
}

type GetGabiaDomainTransferCheckParam struct {
	DomainName string `json:"domainName"`
	//Type       string `json:"type"`
}

type SetGabiaHostParam struct {
	HostName string `json:"hostName"`
	HostIp   string `json:"hostIp"`
	Type     string `json:"type"`
}

type ModiGabiaHostParam struct {
	HostName  string `json:"hostName"`
	OldHostIp string `json:"oldHostIp"`
	NewHostIp string `json:"newHostIp"`
}

type DelGabiaHostParam struct {
	HostName string `json:"hostName"`
	Type     string `json:"type"`
}

type GetFreebillAuthParam struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type freebillItem struct {
	ProductName string  `json:"productName"` // 품목명
	Price       float32 `json:"price"`       // 공급가
	PriceVat    float32 `json:"priceVat"`    // 세액
	PriceVatSum float32 `json:"priceVatSum"` // 규격
	Standard    string  `json:"standard"`    // 수량
	UnitPrice   int     `json:"unitPrice"`   // 단가
}

type SetFreeBillRegisterParam struct {
	// A        string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Tpf      string `json:"tpf"`
	Open     string `json:"open"` // option
	// Y : 등록된 문서의 화면이 브라우저로 확인 가능
	// N : SUCCESS|문서고유번호
	// N 이면서 UniqCode가 null 이 아닐 때 SUCCESS|문서고유번호|요청자 자체 문서고유번호
	IsReceive string `json:"isReceive"`
	Date      string `json:"date"`
	TaxType   string `json:"taxType"`
	TypeText  string `json:"typeText"`
	// Item        string `json:"item"`
	Item        []freebillItem `json:"item"`
	Number      string         `json:"number"`
	Fnumber     string         `json:"fnumber"`
	Company     string         `json:"company"`   // option
	President   string         `json:"president"` // option
	Addr        string         `json:"addr"`      // option
	Btype       string         `json"btype"`      // option
	Bclass      string         `json:"bclass"`    // option
	Name        string         `json:"name"`      // option
	Hp          string         `json:"hp"`        // option
	Fax         string         `json:"fax"`
	Email       string         `json:"email"`       // option
	Message     string         `json:"message"`     // option
	Volume      string         `json:"volume"`      // option
	Issue       string         `json:"issue"`       // option
	Sequence    string         `json:"sequence"`    // option
	Description string         `json:"description"` // option
	PaymentType string         `json:"paymentType"` // option
	UniqCode    string         `json:"uniqCode"`    // option
}

type GetFreeBillSearchParam struct {
	// A     string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Dnumber  string `json:"dNumber"`
	Uniq     string `json:"uniq"`
}

type GetFreeBillSearchStatusParam struct {
	// A     string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Dnumber  string `json:"dNumber"`
	Uniq     string `json:"uniq"`
}

type DelFreeBillParam struct {
	// A     string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Dnumber  string `json:"dNumber"`
	Uniq     string `json:"uniq"`
}

type GetFreeBillViewParam struct {
	// A     string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Dnumber  string `json:"dNumber"`
	Uniq     string `json:"uniq"`
}

type DelFreeBillCancelParam struct {
	// A     string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Dnumber  string `json:"dNumber"`
	Uniq     string `json:"uniq"`
	Message  string `json:"message"`
}

type GetFreeBillPreviousRegisterParam struct {
	// A     string `json:"a"`
	Id       string   `json:"id"`
	Password string   `json:"password"`
	Dnumber  []string `json:"dNumber"`
}

type SetFreeBillPublishNowParam struct {
	// A        string `json:"a"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	SignType string `json:"signType"`
	Tpf      string `json:"tpf"`
	Open     string `json:"open"` // option
	// Y : 화면 바로 가기 기능 사용함
	// N : 화면 바로가기 기능 사용안함, 문서번호 리턴 요청
	IsReceive string `json:"isReceive"`
	Date      string `json:"date"`
	TaxType   string `json:"taxType"`
	TypeText  string `json:"typeText"`
	// Item        string `json:"item"`
	Item        []freebillItem `json:"item"`
	Number      string         `json:"number"`
	Fnumber     string         `json:"fnumber"`
	Company     string         `json:"company"`   // option
	President   string         `json:"president"` // option
	Addr        string         `json:"addr"`      // option
	Btype       string         `json"btype"`      // option
	Bclass      string         `json:"bclass"`    // option
	Name        string         `json:"name"`      // option
	Hp          string         `json:"hp"`        // option
	Fax         string         `json:"fax"`
	Email       string         `json:"email"`       // option
	Message     string         `json:"message"`     // option
	Volume      string         `json:"volume"`      // option
	Issue       string         `json:"issue"`       // option
	Sequence    string         `json:"sequence"`    // option
	Description string         `json:"description"` // option
	PaymentType string         `json:"paymentType"` // option
	UniqCode    string         `json:"uniqCode"`    // option

}

type SetFreeBillSendParam struct {
	// A     string `json:"a"`
	Id       string   `json:"id"`
	Password string   `json:"password"`
	Dnumber  []string `json:"dNumber"`
	Uniq     string   `json:"uniq"`
}

type GetInicisPaymentModuleParam struct {
	// 필수 데이터
	Version   string `json:"version"`
	Mid       string `json:"mid"`
	GoodName  string `json:"goodName"`
	Price     int    `json:"price"` // currency "WON" 으로 고정
	Currency  string `json:"currency"`
	BuyerName string `json:"buyerName"`
	BuyerTel  string `json:"buyerTel"`
	Signature string `json:"signature"`
	ReturnUrl string `json:"returnUrl"`
	CloseUrl  string `json:"closeUrl"`
	PopupUrl  string `json:"popupUrl"`
	// Oid - mid + timestamp (내부에서 만들어서 요청 예정)
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

type GetInicisVirtualAccountParam struct {
	NoTid       string `json:"noTid"`
	NoOid       string `json:"noOid"`
	CdBank      string `json:"cdBank"`
	CdDeal      string `json:"cdDeal"`
	DtTrans     string `json:"dtTrans"`
	TmTrans     string `json:"tmTrans"`
	NoVacct     string `json:"noVacct"`
	AmtInput    int    `json:"amtInput"`
	FlgClose    string `json:"flgClose"`
	ClClose     string `json:"clClose"`
	TypeMsg     string `json:"typeMsg"`
	NmInputbank string `json:"nmInputbank"`
	NmInput     string `json:"nmInput"`
	DtInputstd  int    `json:"dtInputstd"`
	DtCalculstd int    `json:"dtCalculstd"`
	DtTransbase int    `json:"dtTransbase"`
	ClTrans     string `json:"clTrans"`
	ClKor       int    `json:"clKor"`
	DtCshr      int    `json:"dtCshr"`
	TmCshr      int    `json:"tmCshr"`
	NoCshrAppl  int    `json:"noCshrAppl"`
	NoCshrTid   string `json:"noCshrTid"`
}

type SetSubmitAuthCodeWhoisOneParam struct {
	OrderId  int    `json:"orderId"`
	AuthCode string `json:"authCode"`
}

// paypal
type GetPaypalPaymentRequestAllParam struct {
	Transactions []struct {
		Cmd          string `json:"cmd"`
		Business     string `json:"business"`
		NotifyUrl    string `json:"NotifyUrl"`
		Charset      string `json:"vharset"`
		CurrencyType string `json:"vurrencyType"`
		NoteToPayer  string `json:"noteToPayer"` // 판매자에게 보내는 메세지
		ReturnUrl    string `json:"returnUrl"`   // 결제 승인 url
		CancelUrl    string `json:"cancelUrl"`   // 결체 취소 시 url
		ItemsList    []struct {
			Items []struct {
				Name        string  `json:"itemName"`    // 상품명
				Description string  `json:"description"` // 상품 설명서
				Quantity    int     `json:"quantity"`    // 수량
				Price       float64 `json:"price"`       // 가격
				Tax         float64 `json:"tax"`         // 세금
				Sku         string  `json:"sku"`         // 재고 관리 코드
			} `json:"items"`
			RecipientName string `json:"recipientName"` // 받는 사람 이름
			Line1         string `json:"line1"`         // 주소1
			Line2         string `json:"line2"`         // 주소2
			City          string `json:"city"`          // 도시 코드
			CountryCode   string `json:"countryCode"`   // 국가 코드
			PostalCode    string `json:"postalCode"`    // 우편 번호
			Phone         string `json:"phone"`         // 연락처 번호
		} `json:"itemsList"`
	}
}

// type GetPaypalPaymentRequestAllParam struct {
// 	// transactions - array start
// 	Transactions []struct {
// 		Description   string `json:"description"`   // 구매설명서
// 		Custom        string `json:"custom"`        // 구매자 이름
// 		InvoiceNumber string `json:"invoiceNumber"` // 결제 번호
// 		// item_list
// 		ItemsList struct {
// 			Items []struct {
// 				Name        string  `json:"name"`        // 상품명
// 				Description string  `json:"description"` // 상품 설명서
// 				Quantity    int     `json:"quantity"`    // 수량
// 				Price       float64 `json:"price"`       // 가격
// 				Tax         float64 `json:"tax"`         // 세금
// 				Sku         string  `json:"sku"`         // 재고 관리 코드
// 				// Currency    string `json:"currency"`     // 통화
// 			} `json:"items"`
// 			// shipping_address
// 			RecipientName string `json:"recipientName"` // 받는 사람 이름
// 			Line1         string `json:"line1"`         // 주소1
// 			Line2         string `json:"line2"`         // 주소2
// 			City          string `json:"city"`          // 도시 코드
// 			CountryCode   string `json:"countryCode"`   // 국가 코드
// 			PostalCode    string `json:"postalCode"`    // 우편 번호
// 			Phone         string `json:"phone"`         // 연락처 번호
// 		} `json:"itemsList"`
// 	} `json:"transactions"`
// 	// transactions - array end

// 	NoteToPayer string `json:"noteToPayer"` // 판매자에게 보내는 메세지
// 	ReturnUrl   string `json:"returnUrl"`   // 결제 승인 url
// 	CancelUrl   string `json:"cancelUrl"`   // 결체 취소 시 url
// }

type DelPaypalRefundRequestOneParam struct {
	RefundId          string `json:"refundId"`
	PaypalTokenType   string `json:"paypalTokenType"`
	PaypalAccessToken string `json:"paypalAccessToken"`
}

type GetGabiaDomainHostCheckParam struct {
	HostName string `json:"host_name"`
	Type     string `json:"type"`
}

// GetInicisPaymentRefund
type GetInicisPaymentRefundParam struct {
	PayMethod string `json:"payMethod"`
	Tid       string `json:"tid"`
	Msg       string `json:"msg"`
}

// GetInicisPaymentVBankRefund
type GetInicisPaymentVBankRefundParam struct {
	PayMethod      string `json:"payMethod"`
	Tid            string `json:"tid"`
	Msg            string `json:"msg"`
	RefundAccNum   string `json:"refundAccNum"`
	RefundBankCode string `json:"refundBankCode"`
	RefundAccName  string `json:"refundAccName"`
}

// GetInicisVBankNoti
type GetInicisVBankNotiParam struct {
	NoTid       string `json:"no_tid"`
	NoOid       string `json:"no_oid"`
	CdBank      string `json:"cd_bank"`
	CdDeal      string `json:"cd_deal"`
	DtTrans     string `json:"dt_trans"`
	TmTrans     string `json:"tm_trans"`
	NoVacct     string `json:"no_vacct"`
	AmtInput    string `json:"amt_input"`
	FlgClose    string `json:"flg_close"`
	ClClose     string `json:"cl_close"`
	TypeMsg     string `json:"type_msg"`
	NmInputbank string `json:"nm_inputbank"`
	nmInput     string `json:"nm_input"`
	DtInputstd  string `json:"dt_inputstd"`
	DtCalculstd string `json:"dt_transbase"`
	ClTrans     string `json:"cl_trans"`
	ClKor       string `json:"cl_kor"`
	Dtcshr      string `json:"dt_cshr"`
	Tmcshr      string `json:"tm_cshr"`
	NoCshrAppl  string `json:"no_cshr_appl"`
	NoCshrtid   string `json:"no_cshr_tid"`
}

// GetInicisCashReceipt
type GetInicisCashReceiptParam struct {
	GoodName     string  `json:"goodName"`
	Currency     string  `json:"currency"`
	BuyerName    string  `json:"buyerName"`
	BuyerEmail   string  `json:"buyerEmail"`
	BuyerTel     string  `json:"buyerTel"`
	CrPrice      float32 `json:"crPrice"`
	SupPrice     float32 `json:"supPrice"`
	Tax          float32 `json:"tax"`
	RegNum       string  `json:"regNum"`
	CompayNumber string  `json:"compayNumber"`
}
