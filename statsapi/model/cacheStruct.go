package model

type MessageValue struct {
	Kr string `json:"kr"`
	En string `json:"en"`
	Cn string `json:"cn"`
}

// 캐시 리스트
type RespGetCacheListDayAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		CacheList []GetCacheListDayAll `json:"data"`
	} `json:"data"`
}

type GetCacheListDayAll struct {
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	ProviderCompanyName  string `json:"providerCompanyName"`
	Fqdn                 string `json:"fqdn"`
	TotalCount           int    `json:"totalCount"`
	CacheCount           int    `json:"cacheCount"`
	OriginCount          int    `json:"originCount"`
	TotalTraffic         int    `json:"totalTraffic"`
	CacheTraffic         int    `json:"cacheTraffic"`
	OriginTraffic        int    `json:"originTraffic"`
}

// 캐시 그래프
type RespGetCacheGraphDayAll struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		CacheGraph []GetCacheGraphDayAll `json:"data"`
	} `json:"data"`
}

type GetCacheGraphDayAll struct {
	Kst           string `json:"kst"`
	TotalCount    int    `json:"totalCount"`
	CacheCount    int    `json:"cacheCount"`
	OriginCount   int    `json:"originCount"`
	TotalTraffic  int    `json:"totalTraffic"`
	CacheTraffic  int    `json:"cacheTraffic"`
	OriginTraffic int    `json:"originTraffic"`
}

//////////////////////////////////////////////F.E 통계///////////////////////////////////////////

// 캐시 리스트
type RespGetCacheListDayAllFe struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		CacheList []GetCacheListDayAllFe `json:"data"`
	} `json:"data"`
}

type GetCacheListDayAllFe struct {
	ProviderCompanyIndex string `json:"providerCompanyIndex"`
	ProviderCompanyName  string `json:"providerCompanyName"`
	Fqdn                 string `json:"fqdn"`
	TotalCount           int    `json:"totalCount"`
	CacheCount           int    `json:"cacheCount"`
	OriginCount          int    `json:"originCount"`
	TotalTraffic         int    `json:"totalTraffic"`
	CacheTraffic         int    `json:"cacheTraffic"`
	OriginTraffic        int    `json:"originTraffic"`
}

// 캐시 그래프
type RespGetCacheGraphDayAllFe struct {
	Code        string       `json:"code"`
	Message     MessageValue `json:"message"`
	ServiceName string       `json:"serviceName"`
	Data        struct {
		CacheGraph []GetCacheGraphDayAllFe `json:"data"`
	} `json:"data"`
}

type GetCacheGraphDayAllFe struct {
	Kst           string `json:"kst"`
	TotalCount    int    `json:"totalCount"`
	CacheCount    int    `json:"cacheCount"`
	OriginCount   int    `json:"originCount"`
	TotalTraffic  int    `json:"totalTraffic"`
	CacheTraffic  int    `json:"cacheTraffic"`
	OriginTraffic int    `json:"originTraffic"`
}
