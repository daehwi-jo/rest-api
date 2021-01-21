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

// client 리스트 일
func GetClientListDayAll(c echo.Context) error {
	respData := stats.RespGetClientListDayAll{}
	var param stats.GetClientListDayAllParam

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
		param.SortBy = "client_ip"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"provider_company_index": primitive.Regex{Pattern: param.ProviderCompanyIndex, Options: ""},
			"client_ip":              primitive.Regex{Pattern: param.ClientIp, Options: ""},
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":                  bson.M{"client_ip": "$client_ip"},
				"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
				"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
				"clientIp":             bson.M{"$last": "$client_ip"},
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

	coll := Client.Database("hydraplus").Collection("client_hour")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.ClientList); err != nil {
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

// client 리스트 월
func GetClientListMonthAll(c echo.Context) error {
	respData := stats.RespGetClientListDayAll{}
	var param stats.GetClientListDayAllParam

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
		param.SortBy = "clinet_ip"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"provider_company_index": primitive.Regex{Pattern: param.ProviderCompanyIndex, Options: ""},
			"client_ip":              primitive.Regex{Pattern: param.ClientIp, Options: ""},
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":                  bson.M{"client_ip": "$client_ip"},
				"providerCompanyIndex": bson.M{"$last": "$provider_company_index"},
				"providerCompanyName":  bson.M{"$last": "$provider_company_name"},
				"clientIp":             bson.M{"$last": "$client_ip"},
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

	coll := Client.Database("hydraplus").Collection("client_day")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.ClientList); err != nil {
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

// client 통계 그래프 일
func GetClientGraphDayAll(c echo.Context) error {
	respData := stats.RespGetClientGraphDayAll{}
	var param stats.GetClientGraphDayAllParam

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
				"kst":       bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"client_ip": param.ClientIp,
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

	coll := Client.Database("hydraplus").Collection("client_hour")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.ClientGraph); err != nil {
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

// client 통계 그래프 월
func GetClientGraphMonthAll(c echo.Context) error {
	respData := stats.RespGetClientGraphDayAll{}
	var param stats.GetClientGraphDayAllParam

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
				"kst":       bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"client_ip": param.ClientIp,
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

	coll := Client.Database("hydraplus").Collection("client_day")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.ClientGraph); err != nil {
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
