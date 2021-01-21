package model

var GabiaUrl string
var GabiaID string
var GabiaPass string
var GabiaKey string
var GabiaKeyTTL int
var GabiaKeyInTime string

// GABIA API CALL RESPONSE
type GabiaApiCallResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
}

type RespGabiaNSModifyWhoisOne struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	} `json:"data"`
}

// GABIA API CALL RESPONSE
type TestGabiaApiCallResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        string       `json:"data"`
}

type RespGetGabiaHost struct {
	Data string `json:"data"`
}

type RespGetGabiaAuthToken struct {
	GabiaCode       bool   `json:"code"`
	ResultCode      string `json:"http_code"`
	AccessToken     string `json:"access_token"`
	ExpireTTL       int    `json:"expires_in"`
	ACCESSTokenTime string `json:"access_token_time"`
	MakeTime        string `json:"created_on"`
	Msg             string `json:"msg"`
}

type RespGetGabiaDomainCheck struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
	PunyDomain string `json:"puny_domain"`
}

type RespSetGabiaDomain struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
	Epiredate  string `json:"expiredate"`
}

type RespDelGabiaDomain struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespGetGabiaDomain struct {
	GabiaCode    string `json:"code"`
	Msg          string `json:"msg"`
	DomainName   string `json:"domain_name"`
	Gubun        string `json:"gubun"`
	CreateDate   string `json:"createdate"`
	ExpireDate   string `json:"expiredate"`
	Status       string `json:"status"`
	AuthInfo     string `json:"auth_info"`
	Hanname      string `json:"hanname"`
	Engname      string `json:"Engname"`
	Hanorg       string `json:"Hanorg"`
	Engorg       string `json:"Engorg"`
	CtfyCode     string `json:"ctfy_code"`
	CtfyNo       string `json:"ctfy_no"`
	OpenInfoYn   string `json:"open_info_yn"`
	Zip          string `json:"Zip"`
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
	Ahanname     string `json:"a_hanname"`
	Aengname     string `json:"a_engname"`
	Azip         string `json:"a_zip"`
	Aaddr1       string `json:"a_addr1"`
	Aaddr2       string `json:"a_addr2"`
	Aeaddr1      string `json:"a_eaddr1"`
	Aeaddr2      string `json:"a_eaddr2"`
	Acity        string `json:"a_city"`
	Acountrycode string `json:"a_countrycode"`
	Aphone       string `json:"a_phone"`
	Afax         string `json:"a_fax"`
	Ahp          string `json:"a_hp"`
	SmsYn        string `json:"sms_yn"`
	Aemail       string `json:"a_email"`
	NS1          string `json:"ns1"`
	NS2          string `json:"ns2"`
	NS3          string `json:"ns3"`
	NS4          string `json:"ns4"`
}

//도메인 정보 조회, 필요하면 밑에 구조체 따로 파서 필요한 데이터만 담아야 함
type ApiRespGetGabiaDomain struct {
	Code        string             `json:"code"`
	Message     MessageValue       `json:"message"`
	ServiceName string             `json:"serviceName"`
	Data        RespGetGabiaDomain `json:"data"`
}

type RespModiGabiaDomainAdmin struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespModiGabiaDomainNameServer struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespModiGabiaDomainOwner struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespModiGabiaDomainStatusLock struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespSetGabiaDomainIncrease struct {
	GabiaCode          string `json:"code"`
	Msg                string `json:"msg"`
	DomainName         string `json:"domain_name"`
	PreviousExpiredate string `json:"previous_expiredate"`
	Expiredate         string `json:"expiredate"`
}

type ApiRespSetGabiaDomainIncrease struct {
	Code        string                     `json:"code"`
	Message     MessageValue               `json:"message"`
	ServiceName string                     `json:"serviceName"`
	Data        RespSetGabiaDomainIncrease `json:"data"`
}

type RespDelGabiaDomainIncrease struct {
	GabiaCode          string `json:"code"`
	Msg                string `json:"msg"`
	DomainName         string `json:"domain_name"`
	PreviousExpiredate string `json:"previous_expiredate"`
	Expiredate         string `json:"expiredate"`
}

type ApiRespDelGabiaDomainIncrease struct {
	Code        string                     `json:"code"`
	Message     MessageValue               `json:"message"`
	ServiceName string                     `json:"serviceName"`
	Data        RespDelGabiaDomainIncrease `json:"data"`
}

type RespSetGabiaDomainRestore struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespGetGabiaDomainRestoreCheck struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespSetGabiaDomainTransfer struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	DomainName string `json:"domain_name"`
}

type RespGetGabiaDomainTransferCheck struct {
	GabiaCode    string `json:"code"`
	Msg          string `json:"msg"`
	PunyDomain   string `json:"puny_domain"`
	DomainName   string `json:"domain_name"`
	RequestDate  string `json:"requestdate "`
	CompleteDate string `json:"completedate"`
	Status       string `json:"status"`
}

type ApiRespGetGabiaDomainTransferCheck struct {
	Code        string                          `json:"code"`
	Message     MessageValue                    `json:"message"`
	ServiceName string                          `json:"serviceName"`
	Data        RespGetGabiaDomainTransferCheck `json:"data"`
}

type RespGetGabiaDomainHostCheck struct {
	GabiaCode string `json:"code"`
	Msg       string `json:"msg"`
	HostName  string `json:"host_name"`
}

type RespSetGabiaHost struct {
	GabiaCode string `json:"code"`
	Msg       string `json:"msg"`
	HostName  string `json:"host_name"`
}

type RespModiGabiaHost struct {
	GabiaCode string `json:"code"`
	Msg       string `json:"msg"`
	HostName  string `json:"host_name"`
}

type RespDelGabiaHost struct {
	GabiaCode string `json:"code"`
	Msg       string `json:"msg"`
	HostName  string `json:"host_name"`
}

type RespGetGabiaDeposit struct {
	GabiaCode  string `json:"code"`
	Msg        string `json:"msg"`
	ResellerId string `json:"reseller_id "`
	Deposit    string `json:"deposit"`
}

type ApiRespGetGabiaDeposit struct {
	Code        string              `json:"code"`
	Message     MessageValue        `json:"message"`
	ServiceName string              `json:"serviceName"`
	Data        RespGetGabiaDeposit `json:"data"`
}

// gabia 등록
type SetGabiaDomainJsonParam struct {
	DomainName string      `json:"domain_name"`
	Data       Gabiadatast `json:"data"`
}

type Gabiadatast struct {
	Period       int    `json:"period"`
	AutoLockYn   string `json:"auto_lock_yn"`
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
	Ns1          string `json:"ns1"`
	Ns2          string `json:"ns2"`
	Ns3          string `json:"ns3"`
	Ns4          string `json:"ns4"`
	Ns5          string `json:"ns5"`
	Ns6          string `json:"ns6"`
	Ns7          string `json:"ns7"`
	Ns8          string `json:"ns8"`
	Ns9          string `json:"ns9"`
	Ns10         string `json:"ns10"`
	Ns11         string `json:"ns11"`
	Ns12         string `json:"ns12"`
	Ns13         string `json:"ns13"`
	NsIP1        string `json:"nsip1"`
	NsIP2        string `json:"nsip2"`
	NsIP3        string `json:"nsip3"`
	NsIP4        string `json:"nsip4"`
	NsIP5        string `json:"nsip5"`
	NsIP6        string `json:"nsip6"`
	NsIP7        string `json:"nsip7"`
	NsIP8        string `json:"nsip8"`
	NsIP9        string `json:"nsip9"`
	NsIP10       string `json:"nsip10"`
	NsIP11       string `json:"nsip11"`
	NsIP12       string `json:"nsip12"`
	NsIP13       string `json:"nsip13"`
}
