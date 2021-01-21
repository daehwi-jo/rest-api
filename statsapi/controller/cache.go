package controller

import (
	"context"
	"encoding/json"
	stats "hydrawebapi/statsapi/model"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 캐시 리스트 day
func GetCacheListDayAll(c echo.Context) error {
	respData := stats.RespGetCacheListDayAll{}
	var param stats.GetCacheListDayAllParam
	var pipeline []bson.M

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// default
	if param.PerPage == 0 {
		lprintf(4, "[INFO] no value, so set default : [%d]\n", stats.PerRowCount)
		param.PerPage = stats.PerRowCount
	}
	if param.CurrentPage == 0 {
		lprintf(4, "[INFO] no value, so set default : [1]\n")
		param.CurrentPage = 1
	}
	if param.SortBy == "" {
		param.SortBy = "fqdn"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	if param.ProviderCompanyIndex == "" && param.Fqdn == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst": bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	} else if param.ProviderCompanyIndex != "" && param.Fqdn == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"provider_company_index": param.ProviderCompanyIndex,
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	} else if param.ProviderCompanyIndex == "" && param.Fqdn != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":  bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"fqdn": primitive.Regex{Pattern: param.Fqdn, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	} else {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"provider_company_index": param.ProviderCompanyIndex,
				"fqdn":                   primitive.Regex{Pattern: param.Fqdn, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_hour_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 캐시 리스트 Month
func GetCacheListMonthAll(c echo.Context) error {
	respData := stats.RespGetCacheListDayAll{}
	var param stats.GetCacheListDayAllParam
	var pipeline []bson.M

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// default
	if param.PerPage == 0 {
		lprintf(4, "[INFO] no value, so set default : [%d]\n", stats.PerRowCount)
		param.PerPage = stats.PerRowCount
	}
	if param.CurrentPage == 0 {
		lprintf(4, "[INFO] no value, so set default : [1]\n")
		param.CurrentPage = 1
	}
	if param.SortBy == "" {
		param.SortBy = "fqdn"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	if param.ProviderCompanyIndex == "" && param.Fqdn == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst": bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	} else if param.ProviderCompanyIndex != "" && param.Fqdn == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"provider_company_index": param.ProviderCompanyIndex,
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	} else if param.ProviderCompanyIndex == "" && param.Fqdn != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":  bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"fqdn": primitive.Regex{Pattern: param.Fqdn, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	} else {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"provider_company_index": param.ProviderCompanyIndex,
				"fqdn":                   primitive.Regex{Pattern: param.Fqdn, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                  bson.M{"fqdn": "$fqdn"},
					"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
					"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
					"fqdn":                 bson.M{"$last": "$fqdn"},
					"totalCount":           bson.M{"$sum": "$total_count"},
					"cacheCount":           bson.M{"$sum": "$cache_count"},
					"originCount":          bson.M{"$sum": "$origin_count"},
					"totalTraffic":         bson.M{"$sum": "$total_traffic"},
					"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
					"originTraffic":        bson.M{"$sum": "$origin_traffic"},
				},
			},
			bson.M{
				"$project": bson.M{"_id": 0},
			},
			bson.M{
				"$sort": bson.M{param.SortBy: param.SortDesc},
			},
			bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
			bson.M{"$limit": param.PerPage},
		}
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_day_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 캐시 통계 그래프 일
func GetCacheGraphDayAll(c echo.Context) error {
	respData := stats.RespGetCacheGraphDayAll{}
	var param stats.GetCacheGraphDayAllParam

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"kst":  bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"fqdn": param.Fqdn,
			},
		},
		bson.M{
			"$project": bson.M{
				"kst": bson.M{
					"$concat": bson.A{bson.M{"$substr": bson.A{"$kst", 0, 14}}, "00:00"},
				},
				"totalCount":    "$total_count",
				"cacheCount":    "$cache_count",
				"originCount":   "$origin_count",
				"totalTraffic":  "$total_traffic",
				"cacheTraffic":  "$cache_traffic",
				"originTraffic": "$origin_traffic",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_hour_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheGraph); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 캐시 통계 그래프 월
func GetCacheGraphMonthAll(c echo.Context) error {
	respData := stats.RespGetCacheGraphDayAll{}
	var param stats.GetCacheGraphDayAllParam

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"kst":  bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"fqdn": param.Fqdn,
			},
		},
		bson.M{
			"$project": bson.M{
				"kst": bson.M{
					"$concat": bson.A{bson.M{"$substr": bson.A{"$kst", 0, 14}}, "00:00"},
				},
				"totalCount":    "$total_count",
				"cacheCount":    "$cache_count",
				"originCount":   "$origin_count",
				"totalTraffic":  "$total_traffic",
				"cacheTraffic":  "$cache_traffic",
				"originTraffic": "$origin_traffic",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_day_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheGraph); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

//////////////////////////////////////////////F.E 통계///////////////////////////////////////////
// 캐시 일 리스트
func GetCacheListDayAllFe(c echo.Context) error {
	respData := stats.RespGetCacheListDayAllFe{}
	var param stats.GetCacheListDayAllFeParam

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"provider_company_index": param.ProviderCompanyIndex,
			"fqdn":                   primitive.Regex{Pattern: param.Fqdn, Options: ""},
			"user_index":             param.UserIndex,
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":                  bson.M{"fqdn": "$fqdn"},
				"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
				"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
				"fqdn":                 bson.M{"$last": "$fqdn"},
				"totalCount":           bson.M{"$sum": "$total_count"},
				"cacheCount":           bson.M{"$sum": "$cache_count"},
				"originCount":          bson.M{"$sum": "$origin_count"},
				"totalTraffic":         bson.M{"$sum": "$total_traffic"},
				"cacheTraffic":         bson.M{"$sum": "$cache_traffic"},
				"originTraffic":        bson.M{"$sum": "$origin_traffic"},
			},
		},
		bson.M{
			"$project": bson.M{"_id": 0},
		},
		bson.M{
			"$sort": bson.M{param.SortBy: param.SortDesc},
		},
		bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
		bson.M{"$limit": param.PerPage},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_hour_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 캐시 월 리스트
func GetCacheListMonthAllFe(c echo.Context) error {
	respData := stats.RespGetCacheListDayAllFe{}
	var param stats.GetCacheListDayAllFeParam

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"provider_company_index": param.ProviderCompanyIndex,
			"fqdn":                   primitive.Regex{Pattern: param.Fqdn, Options: ""},
			"user_index":             param.UserIndex,
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":                  bson.M{"fqdn": "$fqdn"},
				"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
				"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
				"fqdn":                 bson.M{"$last": "$fqdn"},
				"total_count":          bson.M{"$sum": "$total_count"},
				"cache_count":          bson.M{"$sum": "$cache_count"},
				"origin_count":         bson.M{"$sum": "$origin_count"},
				"total_traffic":        bson.M{"$sum": "$total_traffic"},
				"cache_traffic":        bson.M{"$sum": "$cache_traffic"},
				"origin_traffic":       bson.M{"$sum": "$origin_traffic"},
			},
		},
		bson.M{
			"$project": bson.M{"_id": 0},
		},
		bson.M{
			"$sort": bson.M{param.SortBy: param.SortDesc},
		},
		bson.M{"$skip": (param.CurrentPage - 1) * param.PerPage},
		bson.M{"$limit": param.PerPage},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_day_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 캐시 통계 일 그래프
func GetCacheGraphDayAllFe(c echo.Context) error {
	respData := stats.RespGetCacheGraphDayAllFe{}
	var param stats.GetCacheGraphDayAllFeParam

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"kst":  bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"fqdn": param.Fqdn,
			},
		},
		bson.M{
			"$project": bson.M{
				"kst": bson.M{
					"$concat": bson.A{bson.M{"$substr": bson.A{"$kst", 0, 14}}, "00:00"},
				},
				"totalCount":    "$total_count",
				"cacheCount":    "$cache_count",
				"originCount":   "$origin_count",
				"totalTraffic":  "$total_traffic",
				"cacheTraffic":  "$cache_traffic",
				"originTraffic": "$origin_traffic",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_hour_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheGraph); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 캐시 통계 월 그래프
func GetCacheGraphMonthAllFe(c echo.Context) error {
	respData := stats.RespGetCacheGraphDayAllFe{}
	var param stats.GetCacheGraphDayAllFeParam

	// BODY DATA
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		lprintf(1, "[ERR ] request body : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] resp : [%s]\n", resp)

	if err := json.Unmarshal(resp, &param); err != nil {
		lprintf(1, "[ERR ] Unmarshal : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"kst":  bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"fqdn": param.Fqdn,
			},
		},
		bson.M{
			"$project": bson.M{
				"kst": bson.M{
					"$concat": bson.A{bson.M{"$substr": bson.A{"$kst", 0, 11}}, "00:00:00"},
				},
				"totalCount":    "$total_count",
				"cacheCount":    "$cache_count",
				"originCount":   "$origin_count",
				"totalTraffic":  "$total_traffic",
				"cacheTraffic":  "$cache_traffic",
				"originTraffic": "$origin_traffic",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("cache_day_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.CacheGraph); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}
