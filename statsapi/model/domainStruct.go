package model

// 도메인 리스트
type RespGetDomainListAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		DomainList []GetDomainListAll `json:"data"`
	} `json:"data"`
}

type GetDomainListAll struct {
	ProviderCompanyName string `json:"providerCompanyName"`
	DomainName          string `json:"domainName"`
	Count               int    `json:"count"`
}

// FQDN 리스트
type RespGetDomainFqdnListAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		FqdnList []GetDomainFqdnListAll `json:"data"`
	} `json:"data"`
}

type GetDomainFqdnListAll struct {
	Fqdn  string `json:"fqdn"`
	Count int    `json:"count"`
}

// Dns 쿼리 그래프
type RespGetDnsQueryGraphAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		QueryList []GetDnsQueryGraphAll `json:"data"`
	} `json:"data"`
}

type GetDnsQueryGraphAll struct {
	Kst   string `json:"kst"`
	Count int    `json:"count"`
}

///////////////////////////////////////F.E 통계////////////////////////////////////////
// 도메인 리스트
type RespGetDomainDayListAllFe struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		DomainList []GetDomainDayListAllFe `json:"data"`
	} `json:"data"`
}

type GetDomainDayListAllFe struct {
	ProviderCompanyName string `json:"providerCompanyName"`
	DomainName          string `json:"domainName"`
	Count               int    `json:"count"`
}

// 도메인 통계 상세 fqdn 카운트
type RespGetDomainFqdnDayListAllFe struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		FqdnList []GetDomainFqdnDayListAllFe `json:"data"`
	} `json:"data"`
}

type GetDomainFqdnDayListAllFe struct {
	Fqdn  string `json:"fqdn"`
	Count int    `json:"count"`
}

// Dns 쿼리 그래프
type RespGetDnsQueryDayGraphAllFe struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		QueryList []GetDnsQueryDayGraphAllFe `json:"data"`
	} `json:"data"`
}

type GetDnsQueryDayGraphAllFe struct {
	Kst   string `json:"kst"`
	Count int    `json:"count"`
}
