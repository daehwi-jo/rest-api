package model

type RespGetClientListDayAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		ClientList []GetClientListDayAll `json:"data"`
	} `json:"data"`
}

type GetClientListDayAll struct {
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	ProviderCompanyName  string `json:"providerCompanyName"`
	ClientIp             string `json:"clientIp"`
	TotalCount           int    `json:"totalCount"`
	CacheCount           int    `json:"cacheCount"`
	OriginCount          int    `json:"originCount"`
	TotalTraffic         int    `json:"totalTraffic"`
	CacheTraffic         int    `json:"cacheTraffic"`
	OriginTraffic        int    `json:"originTraffic"`
}

// GetClientGraphDayAll
type RespGetClientGraphDayAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		ClientGraph []GetClientGraphDayAll `json:"data"`
	} `json:"data"`
}

type GetClientGraphDayAll struct {
	Kst           string `json:"kst"`
	TotalCount    int    `json:"totalCount"`
	CacheCount    int    `json:"cacheCount"`
	OriginCount   int    `json:"originCount"`
	TotalTraffic  int    `json:"totalTraffic"`
	CacheTraffic  int    `json:"cacheTraffic"`
	OriginTraffic int    `json:"originTraffic"`
}
