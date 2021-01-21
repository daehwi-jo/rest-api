package model

// 장비
// GetDeviceListAll
type GetDeviceListAllParam struct {
	SelectType  string `json:"selectType"`
	SelectValue string `json:"selectValue"`
	SortBy      string `json:"sortBy"`
	SortDesc    int    `json:"sortDesc"`
	CurrentPage int    `json:"currentPage"`
	PerPage     int    `json:"PerPage"`
}

// GetDeviceInfoOne
type GetDeviceInfoOneParam struct {
	Uuid string `json:"uuid"`
}

// GetDeviceUsageOne
type GetDeviceUsageOneParam struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Uuid      string `json:"uuid"`
}

// GetDeviceDayQueryAll
type GetDeviceDayQueryAllParam struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Uuid      string `json:"uuid"`
}

// GetDeviceMonthQueryAll
type GetDeviceMonthQueryAllParam struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Uuid      string `json:"uuid"`
}

// 도메인
// GetDomainDayListAll, GetDomainMonthListAll
type GetDomainListAllParam struct {
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	DomainName           string `json:"domainName"`
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	SortBy               string `json:"sortBy"`
	SortDesc             int    `json:"sortDesc"`
	CurrentPage          int    `json:"currentPage"`
	PerPage              int    `json:"perPage"`
}

// GetDomainFqdnDayListAll, GetDomainFqdnMonthListAll
type GetDomainFqdnListAllParam struct {
	DomainName  string `json:"domainName"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	SortBy      string `json:"sortBy"`
	SortDesc    int    `json:"sortDesc"`
	CurrentPage int    `json:"currentPage"`
	PerPage     int    `json:"perPage"`
}

// GetDnsQueryDayGraphAll, GetDnsQueryMonthGraphAll
type GetDnsQueryGraphAllParam struct {
	Fqdn      string `json:"fqdn"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// GetCacheListDayAll
type GetCacheListDayAllParam struct {
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	Fqdn                 string `json:"fqdn"`
	CurrentPage          int    `json:"currentPage"`
	PerPage              int    `json:"perPage"`
	SortBy               string `json:"sortBy"`
	SortDesc             int    `json:"sortDesc"`
}

// GetCacheGraphDayAll
type GetCacheGraphDayAllParam struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Fqdn      string `json:"fqdn"`
}

// GetClientListDayAll
type GetClientListDayAllParam struct {
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	ClientIp             string `json:"clientIp"`
	CurrentPage          int    `json:"currentPage"`
	PerPage              int    `json:"perPage"`
	SortBy               string `json:"sortBy"`
	SortDesc             int    `json:"sortDesc"`
}

// GetClientGraphDayAll
type GetClientGraphDayAllParam struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	ClientIp  string `json:"clientIp"`
}

// API CALL RESPONSE
type ApiCallResponse struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        interface{}  `json:"data"`
}

//GetCacheGraphDayAllFe
type GetCacheGraphDayAllFeParam struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Fqdn      string `json:"fqdn"`
}

// GetCacheListDayAllFe
type GetCacheListDayAllFeParam struct {
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	UserIndex            string `json:"userIndex"`
	Fqdn                 string `json:"fqdn"`
	CurrentPage          int    `json:"currentPage"`
	PerPage              int    `json:"perPage"`
	SortBy               string `json:"sortBy"`
	SortDesc             int    `json:"sortDesc"`
}

// GetDomainDayListAllFe
type GetDomainDayListAllFeParam struct {
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	DomainName           string `json:"domainName"`
	UserIndex            string `json:"userIndex"`
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	SortBy               string `json:"sortBy"`
	SortDesc             int    `json:"sortDesc"`
	CurrentPage          int    `json:"currentPage"`
	PerPage              int    `json:"perPage"`
}

// GetDomainFqdnDayListAllFe
type GetDomainFqdnDayListAllFeParam struct {
	DomainName  string `json:"domainName"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	SortBy      string `json:"sortBy"`
	SortDesc    int    `json:"sortDesc"`
	CurrentPage int    `json:"currentPage"`
	PerPage     int    `json:"perPage"`
}

// GetDnsQueryDayGraphAllFe
type GetDnsQueryDayGraphAllFeParam struct {
	Fqdn      string `json:"fqdn"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
