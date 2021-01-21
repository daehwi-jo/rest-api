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

// 도메인 일 리스트
func GetDomainDayListAll(c echo.Context) error {
	respData := stats.RespGetDomainListAll{}
	var param stats.GetDomainListAllParam
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
		param.SortBy = "count"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	if param.ProviderCompanyIndex == "" && param.DomainName == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name": bson.M{"$ne": ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
	} else if param.ProviderCompanyIndex == "" && param.DomainName != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name": primitive.Regex{Pattern: param.DomainName, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
	} else if param.ProviderCompanyIndex != "" && param.DomainName == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name":            bson.M{"$ne": ""},
				"provider_company_index": primitive.Regex{Pattern: param.ProviderCompanyIndex, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
	} else if param.ProviderCompanyIndex != "" && param.DomainName != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"provider_company_index": primitive.Regex{Pattern: param.ProviderCompanyIndex, Options: ""},
				"domain_name":            primitive.Regex{Pattern: param.DomainName, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_hour_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.DomainList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respData.Data : [%v]\n", respData.Data.DomainList)

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 도메인 월 리스트
func GetDomainMonthListAll(c echo.Context) error {
	respData := stats.RespGetDomainListAll{}
	var param stats.GetDomainListAllParam
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
		param.SortBy = "count"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	if param.ProviderCompanyIndex == "" && param.DomainName == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name": bson.M{"$ne": ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
	} else if param.ProviderCompanyIndex == "" && param.DomainName != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name": primitive.Regex{Pattern: param.DomainName, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
	} else if param.ProviderCompanyIndex != "" && param.DomainName == "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name":            bson.M{"$ne": ""},
				"provider_company_index": primitive.Regex{Pattern: param.ProviderCompanyIndex, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
	} else if param.ProviderCompanyIndex != "" && param.DomainName != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":                    bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"provider_company_index": primitive.Regex{Pattern: param.ProviderCompanyIndex, Options: ""},
				"domain_name":            primitive.Regex{Pattern: param.DomainName, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_day_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.DomainList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respData.Data : [%v]\n", respData.Data.DomainList)

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// FQDN 일 리스트
func GetDomainFqdnDayListAll(c echo.Context) error {
	respData := stats.RespGetDomainFqdnListAll{}
	var param stats.GetDomainListAllParam

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

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"domain_name": param.DomainName,
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":   bson.M{"fqdn": "$fqdn"},
				"fqdn":  bson.M{"$last": "$fqdn"},
				"count": bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_hour_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.FqdnList); err != nil {
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

// FQDN 월 리스트
func GetDomainFqdnMonthListAll(c echo.Context) error {
	respData := stats.RespGetDomainFqdnListAll{}
	var param stats.GetDomainListAllParam

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

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"domain_name": param.DomainName,
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":   bson.M{"fqdn": "$fqdn"},
				"fqdn":  bson.M{"$last": "$fqdn"},
				"count": bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_day_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.FqdnList); err != nil {
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

// Dns 쿼리 일 그래프
func GetDnsQueryDayGraphAll(c echo.Context) error {
	respData := stats.RespGetDnsQueryGraphAll{}
	var param stats.GetDnsQueryGraphAllParam

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
				"count": "$count",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("dns_hour_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.QueryList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", respData.Data.QueryList)

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// Dns 쿼리 월 그래프
func GetDnsQueryMonthGraphAll(c echo.Context) error {
	respData := stats.RespGetDnsQueryGraphAll{}
	var param stats.GetDnsQueryGraphAllParam

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
					"$concat": bson.A{bson.M{"$substr": bson.A{"$kst", 0, 10}}, " 00:00:00"},
				},
				"count": "$count",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("dns_day_bo")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.QueryList); err != nil {
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

////////////////////////////////////////////////////////////////F.E 통계////////////////////////////////////////////////////////////////////
// 도메인 통계 일 목록 리스트
func GetDomainDayListAllFe(c echo.Context) error {
	lprintf(4, "[INFO] test! \n")
	respData := stats.RespGetDomainDayListAllFe{}
	var param stats.GetDomainDayListAllFeParam
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
		param.SortBy = "count"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	if param.DomainName != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst": bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				//"domain_name" : bson.M{ "$ne" : "" },
				"provider_company_index": param.ProviderCompanyIndex,
				"domain_name":            primitive.Regex{Pattern: param.DomainName, Options: ""},
				"user_index":             param.UserIndex,
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
				"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				"domain_name": bson.M{"$ne": ""},
				// "provider_company_index": param.ProviderCompanyIndex,
				// // "domain_name":            primitive.Regex{Pattern: param.DomainName, Options: ""},
				// "user_index": param.UserIndex,
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_hour_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.DomainList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respData.Data : [%v]\n", respData.Data.DomainList)

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)

}

// 도메인 통계 월 목록 리스트
func GetDomainMonthListAllFe(c echo.Context) error {
	respData := stats.RespGetDomainDayListAllFe{}
	var param stats.GetDomainDayListAllFeParam
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
		param.SortBy = "count"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	if param.DomainName != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst": bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
				// domain_name : bson.M{ "$ne" : "" },
				"provider_company_index": param.ProviderCompanyIndex,
				"domain_name":            primitive.Regex{Pattern: param.DomainName, Options: ""},
				"user_index":             param.UserIndex,
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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
				"domain_name":            bson.M{"$ne": ""},
				"provider_company_index": param.ProviderCompanyIndex,
				// "domain_name":            primitive.Regex{Pattern: param.DomainName, Options: ""},
				"user_index": param.UserIndex,
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 bson.M{"domain_name": "$domain_name"},
					"providerCompanyName": bson.M{"$last": "$provider_company_name"},
					"domainName":          bson.M{"$last": "$domain_name"},
					"count":               bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_day_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.DomainList); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	lprintf(4, "[INFO] respData.Data : [%v]\n", respData.Data.DomainList)

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)

}

// 도메인 통계 상세 일 fqdn 카운트
func GetDomainFqdnDayListAllFe(c echo.Context) error {
	respData := stats.RespGetDomainFqdnDayListAllFe{}
	var param stats.GetDomainFqdnDayListAllFeParam

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

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"domain_name": param.DomainName,
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":   bson.M{"fqdn": "$fqdn"},
				"fqdn":  bson.M{"$last": "$fqdn"},
				"count": bson.M{"$sum": "$count"},
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

	// pipeline = []bson.M{
	// 	{"$match": bson.M{
	// 		"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
	// 		"domain_name": param.DomainName,
	// 	},
	// 	},
	// 	bson.M{
	// 		"$group": bson.M{
	// 			"_id":   bson.M{"fqdn": "$fqdn"},
	// 			"fqdn":  bson.M{"$last": "$fqdn"},
	// 			"count": bson.M{"$sum": "$count"},
	// 		},
	// 	},
	// 	bson.M{
	// 		"$project": bson.M{"_id": 0},
	// 	},
	// 	bson.M{
	// 		"$sort": bson.M{"fqdn": -1},
	// 	},
	// 	bson.M{"$skip": (1 - 1) * 10},
	// 	bson.M{"$limit": 10},
	// }

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("dns_hour_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.FqdnList); err != nil {
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

// 도메인 통계 상세 월 fqdn 카운트
func GetDomainFqdnMonthListAllFe(c echo.Context) error {
	respData := stats.RespGetDomainFqdnDayListAllFe{}
	var param stats.GetDomainFqdnDayListAllFeParam

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

	pipeline := []bson.M{
		{"$match": bson.M{
			"kst":         bson.M{"$gte": param.StartDate, "$lte": param.EndDate},
			"domain_name": param.DomainName,
		},
		},
		bson.M{
			"$group": bson.M{
				"_id":   bson.M{"fqdn": "$fqdn"},
				"fqdn":  bson.M{"$last": "$fqdn"},
				"count": bson.M{"$sum": "$count"},
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

	coll := Client.Database("hydraplus").Collection("dns_day_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.FqdnList); err != nil {
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

// Dns 쿼리 일 그래프
func GetDnsQueryDayGraphAllFe(c echo.Context) error {
	respData := stats.RespGetDnsQueryDayGraphAllFe{}
	var param stats.GetDnsQueryDayGraphAllFeParam

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
				"count": "$count",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("dns_hour_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.QueryList); err != nil {
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

// Dns 쿼리 월 그래프
func GetDnsQueryMonthGraphAllFe(c echo.Context) error {
	respData := stats.RespGetDnsQueryDayGraphAllFe{}
	var param stats.GetDnsQueryDayGraphAllFeParam

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
					"$concat": bson.A{bson.M{"$substr": bson.A{"$kst", 0, 10}}, " 00:00:00"},
				},
				"count": "$count",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("dns_day_fe")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data.QueryList); err != nil {
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
