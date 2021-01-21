package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	ctrl "hydrawebapi/statsapi/controller"

	"charlie/i3.0.1/cls"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dgrijalva/jwt-go"
)

var lprintf func(int, string, ...interface{}) = cls.Lprintf

func main() {
	fmt.Println("[INFO] start")
	fname := cls.Cls_conf(os.Args)
	domains := cls.WebConf(fname)
	var err error
	// if cls.Db_conf(fname) < 0 {
	// 	lprintf(4, "DataBase not setting \n")
	// }
	// lprintf(4, "[INFO] db connect success")

	// ------------------------------------------------------------------------------
	dbip, r := cls.GetTokenValue("HOST_DBMS", fname)
	if r == -1 {
		lprintf(1, "[FAIL] DBMS not exist value\n")
		return
	}
	lprintf(4, "[INFO] dbip : %s\n", dbip)

	dbpt, r := cls.GetTokenValue("PORT_DBMS", fname)
	if r == -1 {
		lprintf(1, "[FAIL] DBMS not exist value\n")
		return
	}
	lprintf(4, "[INFO] dbpt : %s\n", dbpt)

	dbid, r := cls.GetTokenValue("ID_DBMS", fname)
	if r == -1 {
		lprintf(1, "[FAIL] DBMS not exist value\n")
		return
	}
	lprintf(4, "[INFO] dbid : %s\n", dbid)

	dbps, r := cls.GetTokenValue("PASSWD_DBMS", fname)
	if r == -1 {
		lprintf(1, "[FAIL] DBMS not exist value\n")
		return
	}
	lprintf(4, "[INFO] dbps : %s\n", dbps)

	dbn, r := cls.GetTokenValue("DATABASE_DBMS", fname)
	if r == -1 {
		lprintf(1, "[FAIL] DBMS not exist value\n")
		return
	}
	lprintf(4, "[INFO] dbn : %s\n", dbn)

	// connect
	credential := options.Credential{
		AuthSource: dbn,
		Username:   dbid,
		Password:   dbps,
	}

	dbAddr := fmt.Sprintf("mongodb://%s:%s", dbip, dbpt)
	lprintf(4, "%s", dbAddr)
	clientOption := options.Client().ApplyURI(dbAddr).SetAuth(credential)
	ctrl.Client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		lprintf(1, "[ERR ] MongoDB connection error\n")
		return
	}

	if !ctrl.Redis_App_conf(fname) { //redis set
		lprintf(4, "Redis not setting \n")
		return

	}

	if !ctrl.SetApiToken() { // 기동, 토큰 발급받음
		lprintf(4, "AccessToken not setting \n")
		return
	}
	go ctrl.SchedulerLoop()

	if !cls.GetValidationData(ctrl.BaseAPI_URL) {
		lprintf(1, "[ERR ] Validation not setting \n")
	}

	// ------------------------------------------------------------------------------

	e := echo.New()
	e.Use(middleware.CORS()) // set cors

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			url := c.Request().URL.String()
			lprintf(4, "[INFO] apiUrl : [%s]", url)

			//return next(c) // controller 호출

			if strings.Contains(ctrl.SkipURL, url) {
				return next(c) // controller 호출
			}

			//lprintf(4, "[INFO] pre http header : [%v]\n", c.Request().Header)
			lprintf(4, "[INFO] pre http header : [%v]\n", c.Request().Header.Get("Authorization"))
			auth := strings.Split(c.Request().Header.Get("Authorization"), " ")

			var tokenString string
			if len(auth) > 1 {
				tokenString = auth[1]
			} else {
				tokenString = auth[0]
			}

			lprintf(4, "[INFO] tokenString : [%s]\n", tokenString)

			type jwtCustomClaims struct {
				Number               string
				UserIndex            string
				LoginId              string
				ProviderCompanyIndex string
				jwt.StandardClaims
			}

			if len(tokenString) < 1 {
				//return c.HTML(http.StatusForbidden, "")
				return next(c) // 임시
			}
			headerArr := strings.Split(tokenString, ".")
			var payload string
			if len(headerArr) > 1 {
				payload = headerArr[1]
			} else {
				return c.HTML(http.StatusUnauthorized, "")
			}

			code, err := jwt.DecodeSegment(payload)
			if err != nil {
				lprintf(1, "[ERR ] token payload decoding err : %s\n", err)
				return c.HTML(http.StatusForbidden, "") //
			}
			lprintf(4, "[INFO] code : %s\n", code)

			var clm jwtCustomClaims

			json.Unmarshal(code, &clm)
			lprintf(4, "[INFO] clm.Number : %s\n", clm.Number)

			num1, _ := strconv.Atoi(clm.Number[0:2])
			num2, _ := strconv.Atoi(clm.Number[2:4])
			key := string(ctrl.MatrixF[num1][num2]) + string(ctrl.MatrixF[num2][num1]) + string(ctrl.MatrixL[num1][num2]) + string(ctrl.MatrixL[num2][num1])
			lprintf(4, "[INFO] num %d, %d : \n", num1, num2)

			_, err3 := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(key), nil
			})
			if err3 != nil {
				lprintf(1, "[ERR ] ParseWithClaims error : %s\n", err3)
				if strings.Contains(err3.Error(), "token is expired") { //리프레시 만료 (1주일)
					ctrl.DeleteTK(tokenString)
					return c.HTML(http.StatusUnauthorized, "")
				}
				return c.HTML(http.StatusUnauthorized, "") //
			}

			if !ctrl.ConfirmTK(tokenString) { //토큰 만료 체크 함수
				lprintf(1, "[ERR ] Token Expire error \n")
				return c.HTML(http.StatusUnauthorized, "")
			}
			/*
				claims := token.Claims.(*jwtCustomClaims)
				if token.Valid { // 토큰 유효성 검사
					lprintf(4, "[INFO] userIndex : [%s], loginId : [%s], providerCompanyIndex : [%s], ExpiresAt : [%s]\n", claims.UserIndex, claims.LoginId, claims.ProviderCompanyIndex, strconv.FormatInt(claims.StandardClaims.ExpiresAt, 10))
					c.Set("userIndex", claims.UserIndex)
					c.Set("loginId", claims.LoginId)
					c.Set("providerCompanyIndex", claims.ProviderCompanyIndex)
				} else {
					return c.HTML(http.StatusForbidden, "") //
				}
			*/
			return next(c)
		}
	})

	// device
	e.POST("/v1/statsapi/device-list", ctrl.GetDeviceListAll)             // 장비 리스트
	e.POST("/v1/statsapi/device-info", ctrl.GetDeviceInfoOne)             // 장비 정보
	e.POST("/v1/statsapi/device-info-usage", ctrl.GetDeviceUsageOne)      // 장비 그래프 평균 최대 최저
	e.POST("/v1/statsapi/device-info-day", ctrl.GetDeviceDayQueryAll)     // 장비 그래프 일 쿼리
	e.POST("/v1/statsapi/device-info-month", ctrl.GetDeviceMonthQueryAll) // 장비 그래프 월 쿼리

	// domain
	e.POST("/v1/statsapi/domain-list-day", ctrl.GetDomainDayListAll)                       // 도메인 일 리스트
	e.POST("/v1/statsapi/domain-list-month", ctrl.GetDomainMonthListAll)                   // 도메인 월 리스트
	e.POST("/v1/statsapi/domain/fqdn-list-day", ctrl.GetDomainFqdnDayListAll)              // FQDN 일 리스트
	e.POST("/v1/statsapi/domain/fqdn-list-month", ctrl.GetDomainFqdnMonthListAll)          // FQDN 월 리스트
	e.POST("/v1/statsapi/domain/fqdn-list/dns-graph-day", ctrl.GetDnsQueryDayGraphAll)     // Dns 쿼리 일 그래프
	e.POST("/v1/statsapi/domain/fqdn-list/dns-graph-month", ctrl.GetDnsQueryMonthGraphAll) // Dns 쿼리 월 그래프

	e.POST("/v1/front/statsapi/domain-list-day", ctrl.GetDomainDayListAllFe)                       // 도메인 일 리스트
	e.POST("/v1/front/statsapi/domain-list-month", ctrl.GetDomainMonthListAllFe)                   // 도메인 월 리스트
	e.POST("/v1/front/statsapi/domain/fqdn-list-day", ctrl.GetDomainFqdnDayListAllFe)              // FQDN 일 리스트
	e.POST("/v1/front/statsapi/domain/fqdn-list-month", ctrl.GetDomainFqdnMonthListAllFe)          // FQDN 월 리스트
	e.POST("/v1/front/statsapi/domain/fqdn-list/dns-graph-day", ctrl.GetDnsQueryDayGraphAllFe)     // Dns 쿼리 일 그래프
	e.POST("/v1/front/statsapi/domain/fqdn-list/dns-graph-month", ctrl.GetDnsQueryMonthGraphAllFe) // Dns 쿼리 월 그래프

	// cache
	e.POST("/v1/statsapi/cache/cache-list-day", ctrl.GetCacheListDayAll)            // 캐시 리스트 일
	e.POST("/v1/statsapi/cache/cache-list-month", ctrl.GetCacheListMonthAll)        // 캐시 리스트 월
	e.POST("/v1/statsapi/cache/cache-list/graph-day", ctrl.GetCacheGraphDayAll)     // 캐시 통계 그래프 일
	e.POST("/v1/statsapi/cache/cache-list/graph-month", ctrl.GetCacheGraphMonthAll) // 캐시 통계 그래프 월

	e.POST("/v1/front/statsapi/cache/cache-list-day", ctrl.GetCacheListDayAllFe)            // 캐시 리스트 일
	e.POST("/v1/front/statsapi/cache/cache-list-month", ctrl.GetCacheListMonthAllFe)        // 캐시 리스트 월
	e.POST("/v1/front/statsapi/cache/cache-list/graph-day", ctrl.GetCacheGraphDayAllFe)     // 캐시 통계 그래프 일
	e.POST("/v1/front/statsapi/cache/cache-list/graph-month", ctrl.GetCacheGraphMonthAllFe) // 캐시 통계 그래프 월

	// client
	e.POST("/v1/statsapi/client/client-list-day", ctrl.GetClientListDayAll)            // 캐시 리스트 일
	e.POST("/v1/statsapi/client/client-list-month", ctrl.GetClientListMonthAll)        // 캐시 리스트 월
	e.POST("/v1/statsapi/client/client-list/graph-day", ctrl.GetClientGraphDayAll)     // client 통계 그래프 일
	e.POST("/v1/statsapi/client/client-list/graph-month", ctrl.GetClientGraphMonthAll) // client 통계 그래프 월

	//	e.POST("/v1/system/user/feuser/")

	/*
		// set url
		e.Renderer = echotemplate.New(echotemplate.TemplateConfig{
			Root:         "templates",
			Extension:    ".html",
			Partials:     []string{},
			DisableCache: true,
			Delims:       echotemplate.Delims{Left: "[[", Right: "]]"},
		})
	*/

	domains[0].EchoData = e
	cls.StartDomain(domains)
	fmt.Println("[INFO] end")
	return
}
