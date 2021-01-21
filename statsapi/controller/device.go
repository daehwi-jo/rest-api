package controller

import (
	"context"
	"encoding/json"
	stats "hydrawebapi/statsapi/model"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"charlie/i3.0.1/cls"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

var lprintf func(int, string, ...interface{}) = cls.Lprintf

var Client *mongo.Client

var MatrixF = [30][30]byte{
	{'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'a', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b'},
	{'s', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
	{'d', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'c', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd'},
	{'u', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c'},
	{'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'e', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f'},
	{'w', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e'},
	{'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'g', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h'},
	{'y', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g'},
	{'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'i', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j'},
	{'1', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i'},
	{'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'k', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l'},
	{'3', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k'},
	{'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'n', '6', 'm', 'n', '5', 'n', '5', 'n', '5', 'n', '5', 'n', '5', 'n'},
	{'5', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', '6', '5', 'n', 'm', 'n', 'm', 'n', 'm', 'n', 'm', 'n', 'm', 'n', 'm'},
	{'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'p', '8', 'o', 'p', '7', 'p', '7', 'p', '7', 'p', '7', 'p', '7', 'p'},
	{'7', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', '8', '7', 'p', 'o', 'p', 'o', 'p', 'o', 'p', 'o', 'p', 'o', 'p', 'o'},
	{'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'r', '0', 'q', 'r', '9', 'r', '9', 'r', '9', 'r', '9', 'r', '9', 'r'},
	{'9', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', '0', '9', 'r', 'q', 'r', 'q', 'r', 'q', 'r', 'q', 'r', 'q', 'r', 'q'},
	{'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'b', 't', 'a', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b', 's', 'b'},
	{'s', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 't', 's', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
	{'d', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'd', 'v', 'c', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd', 'u', 'd'},
	{'u', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'v', 'u', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c', 'd', 'c'},
	{'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'f', 'x', 'e', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f', 'w', 'f'},
	{'w', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'x', 'w', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e', 'f', 'e'},
	{'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'h', 'z', 'g', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h', 'y', 'h'},
	{'y', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'z', 'y', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g', 'h', 'g'},
	{'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'j', '2', 'i', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j', '1', 'j'},
	{'1', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', '2', '1', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i', 'j', 'i'},
	{'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'l', '4', 'k', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l', '3', 'l'},
	{'3', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', '4', '3', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k', 'l', 'k'},
}

var MatrixL = [30][30]byte{
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'0', 't', '8', 'v', '6', 'x', '4', 'z', '2', '2', 'z', '4', 'x', '6', 'v', '8', 't', '0', 't', 't', '0', 'v', '8', 'x', '6', 'z', '4', '2', '2', '4'},
	{'t', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3', '5', '5', '7', '7', '9', '9', 's', 's', 'u', 'u', 'w', 'w', 'y', 'y', '1', '1', '3', '3'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
	{'r', 'b', 'p', 'd', 'n', 'f', 'l', 'h', 'j', 'j', 'h', 'l', 'f', 'n', 'd', 'p', 'b', 'r', 'b', 'b', 'r', 'd', 'p', 'f', 'n', 'h', 'l', 'j', 'j', 'l'},
	{'b', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k', 'm', 'm', 'o', 'o', 'q', 'q', 'a', 'a', 'c', 'c', 'e', 'e', 'g', 'g', 'i', 'i', 'k', 'k'},
}

// 장비 리스트
func GetDeviceListAll(c echo.Context) error {
	respData := stats.RespGetDeviceListAll{}
	var param stats.GetDeviceListAllParam
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
		param.SortBy = "kst"
	}
	if param.SortDesc == 0 {
		lprintf(4, "[INFO] no value, so set default : [asc]\n")
		param.SortDesc = 1
	}

	now := time.Now()

	// startDate := now.Format("2006-01-02")
	startDate := now.AddDate(0, -1, 1).Format("2006-01-02")
	endDate := now.AddDate(0, 0, 1).Format("2006-01-02")

	if param.SelectType != "" {
		pipeline = []bson.M{
			{"$match": bson.M{
				"kst":            bson.M{"$gte": startDate, "$lte": endDate},
				param.SelectType: primitive.Regex{Pattern: param.SelectValue, Options: ""},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":        bson.M{"uuid": "$uuId"},
					"uuid":       bson.M{"$last": "$uuId"},
					"deviceName": bson.M{"$last": "$devName"},
					"publicIp":   bson.M{"$last": "$publicIp"},
					"cpu":        bson.M{"$last": bson.M{"$toDouble": "$cpu"}},
					"memory":     bson.M{"$last": bson.M{"$toDouble": "$mem"}},
					"storage":    bson.M{"$last": bson.M{"$toDouble": "$disk"}},
					"network":    bson.M{"$last": bson.M{"$toDouble": "$net"}},
					"time":       bson.M{"$last": "$kst"},
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
				"kst": bson.M{"$gte": startDate, "$lte": endDate},
			},
			},
			bson.M{
				"$group": bson.M{
					"_id":        bson.M{"uuid": "$uuId"},
					"uuid":       bson.M{"$last": "$uuId"},
					"deviceName": bson.M{"$last": "$devName"},
					"publicIp":   bson.M{"$last": "$publicIp"},
					"cpu":        bson.M{"$last": bson.M{"$toDouble": "$cpu"}},
					"memory":     bson.M{"$last": bson.M{"$toDouble": "$mem"}},
					"storage":    bson.M{"$last": bson.M{"$toDouble": "$disk"}},
					"network":    bson.M{"$last": bson.M{"$toDouble": "$net"}},
					"time":       bson.M{"$last": "$kst"},
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

	coll := Client.Database("hydraplus").Collection("device")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData.Data)
	}
	lprintf(4, "[INFO] respData.Data : [%v]\n", respData.Data)

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 장비 정보
func GetDeviceInfoOne(c echo.Context) error {
	respData := stats.RespGetDeviceInfoOne{}
	var param stats.GetDeviceInfoOneParam
	var dataStruct []stats.GetDeviceInfoOne

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

	now := time.Now()

	startDate := now.Format("2006-01-02")
	endDate := now.AddDate(0, 0, 1).Format("2006-01-02")

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"kst":  bson.M{"$gte": startDate, "$lte": endDate},
				"uuId": primitive.Regex{Pattern: param.Uuid, Options: ""},
			},
		},
		bson.M{
			"$sort": bson.M{"kst": -1},
		},
		bson.M{
			"$group": bson.M{
				"_id":        bson.M{"uuid": "$uuId"},
				"uuid":       bson.M{"$last": "$uuId"},
				"deviceName": bson.M{"$last": "$devName"},
				"publicIp":   bson.M{"$last": "$publicIp"},
				"cpu":        bson.M{"$last": bson.M{"$toDouble": "$cpu"}},
				"memory":     bson.M{"$last": bson.M{"$toDouble": "$mem"}},
				"storage":    bson.M{"$last": bson.M{"$toDouble": "$disk"}},
				"network":    bson.M{"$last": bson.M{"$toDouble": "$net"}},
				"time":       bson.M{"$last": "$kst"},
			},
		},
		bson.M{
			"$project": bson.M{"_id": 0},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("device")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &dataStruct); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if dataStruct != nil {
		respData.Data.Uuid = dataStruct[0].Uuid
		respData.Data.DeviceName = dataStruct[0].DeviceName
		respData.Data.PublicIp = dataStruct[0].PublicIp
		respData.Data.Cpu = dataStruct[0].Cpu
		respData.Data.Memory = dataStruct[0].Memory
		respData.Data.Storage = dataStruct[0].Storage
		respData.Data.Network = dataStruct[0].Network
		respData.Data.Time = dataStruct[0].Time
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

// 장비 그래프 평균 최대 최저
func GetDeviceUsageOne(c echo.Context) error {
	respData := stats.RespGetDeviceUsageOne{}
	var param stats.GetDeviceUsageOneParam
	var dataStruct []stats.GetDeviceUsageOne

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
				"uuId": param.Uuid,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":            bson.M{"uuid": "$uuId"},
				"uuid":           bson.M{"$last": "$uuId"},
				"cpuAverage":     bson.M{"$avg": bson.M{"$toDouble": "$cpu"}},
				"cpuMax":         bson.M{"$max": bson.M{"$toDouble": "$cpu"}},
				"cpuMin":         bson.M{"$min": bson.M{"$toDouble": "$cpu"}},
				"memoryAverage":  bson.M{"$avg": bson.M{"$toDouble": "$mem"}},
				"memoryMax":      bson.M{"$max": bson.M{"$toDouble": "$mem"}},
				"memoryMin":      bson.M{"$min": bson.M{"$toDouble": "$mem"}},
				"storageAverage": bson.M{"$avg": bson.M{"$toDouble": "$disk"}},
				"storageMax":     bson.M{"$max": bson.M{"$toDouble": "$disk"}},
				"storageMin":     bson.M{"$min": bson.M{"$toDouble": "$disk"}},
				"networkAverage": bson.M{"$avg": bson.M{"$toDouble": "$net"}},
				"networkMax":     bson.M{"$max": bson.M{"$toDouble": "$net"}},
				"networkMin":     bson.M{"$min": bson.M{"$toDouble": "$net"}},
			},
		},
		bson.M{
			"$project": bson.M{"_id": 0},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("device")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &dataStruct); err != nil {
		lprintf(1, "[ERR ] structure parsing error : [%s]\n", err)
		respData.Code = stats.C500
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}

	if dataStruct != nil {
		respData.Data.Uuid = dataStruct[0].Uuid
		respData.Data.CpuAverage = dataStruct[0].CpuAverage
		respData.Data.CpuMax = dataStruct[0].CpuMax
		respData.Data.CpuMin = dataStruct[0].CpuMin
		respData.Data.MemoryAverage = dataStruct[0].MemoryAverage
		respData.Data.MemoryMax = dataStruct[0].MemoryMax
		respData.Data.MemoryMin = dataStruct[0].MemoryMin
		respData.Data.StorageAverage = dataStruct[0].StorageAverage
		respData.Data.StorageMax = dataStruct[0].StorageMax
		respData.Data.StorageMin = dataStruct[0].StorageMin
		respData.Data.NetworkAverage = dataStruct[0].NetworkAverage
		respData.Data.NetworkMax = dataStruct[0].NetworkMax
		respData.Data.NetworkMin = dataStruct[0].NetworkMin
	}

	// make response
	respData.Code = stats.SUCCESS
	respData.ServiceName = stats.TYPE

	return c.JSON(http.StatusOK, respData)
}

//장비 그래프 일 쿼리
func GetDeviceDayQueryAll(c echo.Context) error {
	respData := stats.RespGetDeviceDayQueryAll{}
	var param stats.GetDeviceDayQueryAllParam

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
				"uuid": param.Uuid,
			},
		},
		bson.M{
			"$project": bson.M{
				"uuid": "$uuid",
				"kst": bson.M{
					"$concat": bson.A{
						bson.M{
							"$substr": bson.A{"$kst", 0, 14},
						}, "00:00"},
				},
				"cpuAverage":     "$cpu_average",
				"memoryAverage":  "$memory_average",
				"storageAverage": "$storage_average",
				"networkAverage": "$network_average",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         0,
				"count":       0,
				"device_name": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("device_hour")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data); err != nil {
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

// 장비 그래프 월 쿼리
func GetDeviceMonthQueryAll(c echo.Context) error {
	respData := stats.RespGetDeviceMonthQueryAll{}
	var param stats.GetDeviceMonthQueryAllParam

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
				"uuid": param.Uuid,
			},
		},
		bson.M{
			"$project": bson.M{
				"uuid": "$uuid",
				"kst": bson.M{
					"$concat": bson.A{
						bson.M{
							"$substr": bson.A{"$kst", 0, 10},
						}, "00:00:00"},
				},
				"cpuAverage":     "$cpu_average",
				"memoryAverage":  "$memory_average",
				"storageAverage": "$storage_average",
				"networkAverage": "$network_average",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         0,
				"count":       0,
				"device_name": 0,
			},
		},
	}

	lprintf(4, "[INFO] pipeline : [%v]\n", pipeline)

	coll := Client.Database("hydraplus").Collection("device_day")

	cur, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		lprintf(1, "[ERR ] collection.Aggregate : [%s]\n", err)
		respData.Code = stats.FAIL
		respData.ServiceName = stats.TYPE
		return c.JSON(http.StatusOK, respData)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &respData.Data); err != nil {
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
