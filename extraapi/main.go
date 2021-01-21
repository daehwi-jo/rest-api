package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	ctrl "hydrawebapi/extraapi/controller"
	extra "hydrawebapi/extraapi/model"

	"charlie/i3.0.0/cls"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var lprintf func(int, string, ...interface{}) = cls.Lprintf
var WhiteList []string

func appConf(fname string) {

	// whois
	extra.ApiKey, _ = cls.GetTokenValue("API_KEY", fname)
	extra.AuthUserId, _ = cls.GetTokenValue("AUTH_USER_ID", fname)
	extra.WhoisUrl, _ = cls.GetTokenValue("WHOIS_URL", fname)
	extra.WhoisCheckUrl, _ = cls.GetTokenValue("WHOIS_CHECK_URL", fname)
	extra.CustomerId, _ = cls.GetTokenValue("CUSTOMER_ID", fname)
	extra.RegContactId, _ = cls.GetTokenValue("REG_CONTACT_ID", fname)
	extra.AdminContactId, _ = cls.GetTokenValue("ADMIN_CONTACT_ID", fname)
	extra.TechContactId, _ = cls.GetTokenValue("TECH_CONTACT_ID", fname)
	extra.BillingContactId, _ = cls.GetTokenValue("BILLING_CONTACT_ID", fname)
	extra.InvoiceOption, _ = cls.GetTokenValue("INVOICE_OPTION", fname)
	extra.ProtectPrivate, _ = cls.GetTokenValue("PROTECT_PRIVATE", fname)
	extra.Years, _ = cls.GetTokenValue("YEARS", fname)

	tmp, err := cls.GetTokenValue("MODE", fname)
	if err == cls.CONF_ERR {
		tmp = "live"
	}
	if tmp == "live" || tmp == "LIVE" {
		extra.Mode = true
	} else {
		extra.Mode = false
	}

	tmpNs, _ := cls.GetTokenValue("NS", fname)
	arrayNs := strings.Split(tmpNs, ",")
	for _, v := range arrayNs {
		extra.Ns = append(extra.Ns, strings.TrimSpace(v))
	}

	// gabia
	extra.GabiaUrl, _ = cls.GetTokenValue("GABIA_URL", fname)
	extra.GabiaID, _ = cls.GetTokenValue("GABIA_ID", fname)
	extra.GabiaPass, _ = cls.GetTokenValue("GABIA_PASS", fname)
	go ctrl.GetGabiaAuthTokenScheduler()

	// alipay
	extra.AliPayPartnerId, _ = cls.GetTokenValue("ALIPAY_PARTNER_ID", fname)
	extra.AliPayKey, _ = cls.GetTokenValue("ALIPAY_KEY", fname)
	extra.AliOperUrl, _ = cls.GetTokenValue("ALIPAY_OPER_URL", fname)
	extra.PaymentRedirect, _ = cls.GetTokenValue("PAYMENT_REDIRECT", fname)
	//extra.PaymentReturn, _ = cls.GetTokenValue("PAYMENT_RETURN", fname)
	// white list
	tmpWhiteList, err := cls.GetTokenValue("WHITELIST", fname)
	if err == cls.CONF_ERR {
		fmt.Println("[FAIL] NOT FOUND KEYWORD in *.ini : WHITELIST")
		os.Exit(1)
	}
	tmpWhiteList = strings.TrimRight(tmpWhiteList, ",")
	tmpWhiteArray := strings.Split(tmpWhiteList, ",")
	for _, v := range tmpWhiteArray {
		WhiteList = append(WhiteList, strings.TrimSpace(v))
	}

	lprintf(4, "[INFO] ==================== whois ==========================\n")
	lprintf(4, "[INFO] ApiKey           : [%s]\n", extra.ApiKey)
	lprintf(4, "[INFO] AuthUserId       : [%s]\n", extra.AuthUserId)
	lprintf(4, "[INFO] WhoisUrl         : [%s]\n", extra.WhoisUrl)
	lprintf(4, "[INFO] WhoisCheckUrl    : [%s]\n", extra.WhoisCheckUrl)
	lprintf(4, "[INFO] CustomerId       : [%s]\n", extra.CustomerId)
	lprintf(4, "[INFO] RegContactId     : [%s]\n", extra.RegContactId)
	lprintf(4, "[INFO] AdminContactId   : [%s]\n", extra.AdminContactId)
	lprintf(4, "[INFO] TechContactId    : [%s]\n", extra.TechContactId)
	lprintf(4, "[INFO] BillingContactId : [%s]\n", extra.BillingContactId)
	lprintf(4, "[INFO] InvoiceOption    : [%s]\n", extra.InvoiceOption)
	lprintf(4, "[INFO] ProtectPrivate   : [%s]\n", extra.ProtectPrivate)
	lprintf(4, "[INFO] Years            : [%s]\n", extra.Years)
	lprintf(4, "[INFO] Ns               : [%s]\n", extra.Ns)
	lprintf(4, "[INFO] =================== gabia ===========================\n")
	lprintf(4, "[INFO] GabiaUrl         : [%s]\n", extra.GabiaUrl)
	lprintf(4, "[INFO] GABIA_ID         : [%s]\n", extra.GabiaID)
	lprintf(4, "[INFO] GABIA_PASS         : [%s]\n", extra.GabiaPass)
	lprintf(4, "[INFO] ================== alipay ===========================\n")
	lprintf(4, "[INFO] AliPayPartnerId  : [%s]\n", extra.AliPayPartnerId)
	lprintf(4, "[INFO] AliPayKey        : [%s]\n", extra.AliPayKey)
	lprintf(4, "[INFO] AliOperUrl       : [%s]\n", extra.AliOperUrl)
	lprintf(4, "[INFO] PaymentRedirect  : [%s]\n", extra.PaymentRedirect)
	lprintf(4, "[INFO] PaymentRedirect  : [%s]\n", extra.PaymentReturn)
	lprintf(4, "[INFO] ================== whitelist ========================\n")
	lprintf(4, "[INFO] WhiteList        : [%s]\n", WhiteList)
	lprintf(4, "[INFO] =====================================================\n")
}

func sub_main() {
	fmt.Println("[INFO] start")
	fname := cls.Cls_conf(os.Args)
	domains := cls.WebConf(fname)

	if !ctrl.Redis_App_conf(fname) { //redis set
		lprintf(4, "Redis not setting \n")
		return

	}
	/*
		if !ctrl.RedisCon() { //redis set
			lprintf(4, "Redis Connect setting fail \n")
			return
		}
	*/
	if !ctrl.SetApiToken() { // 기동, 토큰 발급받음
		lprintf(4, "AccessToken not setting \n")
		return
	}
	go ctrl.SchedulerLoop()
	appConf(fname)

	// aliResp := ctrl.GetAlipayPayment()
	// fmt.Println(aliResp)

	// if cls.Db_conf(fname) < 0 {
	// 	lprintf(4, "DataBase not setting \n")
	// }
	// lprintf(4, "[INFO] db connect success")

	e := echo.New()

	/*
		Client 요청
		e.Pre 호출 하고
		next 호출을 해야 -> controller 넘어감
	*/
	e.Use(middleware.CORS()) // set cors

	// 허용된 client만 통신 가능
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		remoteAddr := c.Request().RemoteAddr
	// 		clientIp := strings.Split(remoteAddr, ":")

	// 		cls.Lprintf(4, "[INFO] clientIp : [%s]", clientIp[0])
	// 		checkIp := 0
	// 		for _, v := range WhiteList {
	// 			if clientIp[0] == v {
	// 				checkIp = 1
	// 				break
	// 			}
	// 		}

	// 		if checkIp == 0 {
	// 			return c.HTML(http.StatusForbidden, "") // 허용되지 않은 ip 응답 403
	// 		}
	// 		err := next(c) // controller 호출
	// 		return err
	// 	}
	// })
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			url := c.Request().URL.String()
			lprintf(4, "[INFO] apiUrl : [%s]", url)

			//return next(c) // controller 호출

			if strings.Contains(ctrl.SkipURL, url) {
				return next(c) // controller 호출
			}

			lprintf(4, "[INFO] pre http header : [%v]\n", c.Request().Header)
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

	e.GET("/v1/domain/extra-domain", ctrl.GetExtraApiAll) // 등록한 외부 API 검색

	// Whois
	e.PUT("/v1/extra/whois/domain/domain-register", ctrl.SetDomainRegisterWhoisOne)                   // Whois 도메인 등록
	e.POST("/v1/extra/whois/domain/domain-available", ctrl.GetDomainAvailableWhoisAll)                // Whois 도메인 사용가능 여부 유효성 검증
	e.POST("/v1/extra/whois/domain/domain-recommend", ctrl.GetDomainRecommendWhoisAll)                // 후이즈에 키워드에 대한 추천 도메인을 조회
	e.POST("/v1/extra/whois/domain/domain-detail", ctrl.GetDomainDetailWhoisAll)                      // 후이즈 도메인 등록 순서 세부 정보
	e.POST("/v1/extra/whois/domain/domain-detail/scheduler", ctrl.GetDomainDetailWhoisAllS)           // 후이즈 도메인 등록 순서 세부 정보 (스케줄러 사용)
	e.POST("/v1/extra/whois/domain/domain-search", ctrl.GetDomainSearchWhoisAll)                      // 후이즈에 Search
	e.POST("/v1/extra/whois/domain/domain-available-sunrise", ctrl.GetDomainAvailableSunriseWhoisOne) // 후이즈 이용 가능 여부 확인
	// /api/domains/available-sunrise.json
	// https://demo.myorderbox.com/kb/node/2016 참고 smd 파일 여부 확인 필요!!
	// smd 파일 설명 : https://demo.myorderbox.com/kb/node/2018
	e.PUT("/v1/extra/whois/domain/domain-ns", ctrl.ModiDomainModifyNsWhoisOne)                       // 후이즈에 네임서버 수정
	e.POST("/v1/extra/whois/domain/domain-order-id", ctrl.GetDomainOrderIdWhoisOne)                  // 후이즈 orderid 가져오기
	e.PUT("/v1/extra/whois/domain/domain-child-ns", ctrl.SetDomainAddChildNameServerAll)             // 후이즈에 차일드 네임서버 추가
	e.PATCH("/v1/extra/whois/domain/host-child-ns", ctrl.ModiHostChildNameServerAll)                 // 후이즈에 차일드 네임서버 호스트 이름 수정
	e.PATCH("/v1/extra/whois/domain/ip-child-ns", ctrl.ModiIpChildNameServerOne)                     // 후이즈에 차일드 네임서버 IP 수정
	e.DELETE("/v1/extra/whois/domain/child-ns", ctrl.DelChildNameServerAll)                          // 후이즈에 차일드 네임서버 삭제
	e.POST("/v1/extra/whois/domain/domain-detail/order-id", ctrl.GetDomainDetailOrderIdAll)          // 후이즈에 주문ID 사용하여 도메인 세부정보
	e.PUT("/v1/extra/whois/domain/domain-renew", ctrl.SetDomainRenewWhoisOne)                        // 후이즈에 Renew
	e.PUT("/v1/extra/whois/domain/domain-restore", ctrl.SetDomainRestoreWhoisOne)                    // 후이즈 도메인 복원
	e.DELETE("/v1/extra/whois/domain/domain-name", ctrl.DelDomainNameWhoisOne)                       // 후이즈 도메인 이름 삭제
	e.POST("/v1/extra/whois/domain/domain-validate-transfer", ctrl.GetDomainValidateTransferOne)     // 후이즈 기관 이전 요청 유효성 검사
	e.PUT("/v1/extra/whois/domain/domain-validate-transfer", ctrl.SetDomainValidateTransferOne)      // 후이즈에 기관 이전
	e.PUT("/v1/extra/whois/domain/submit-auth-code", ctrl.SetSubmitAuthCodeWhoisOne)                 // 후이즈에 인증코드 제출
	e.PATCH("/v1/extra/whois/domain/auth-code", ctrl.ModiAuthCodeWhoisOne)                           // 후이즈 인증코드 수정
	e.PUT("/v1/extra/whois/domain/resend-rfa", ctrl.SetResendRfaWhoisOne)                            // 후이즈 이전 승인 메일 재발송
	e.DELETE("/v1/extra/whois/domain/cancel-transfer", ctrl.DelCancelTransferWhoisOne)               // 후이즈 기관 이전 취소
	e.PUT("/v1/extra/whois/domain/resend-verification", ctrl.SetResendVerificationWhoisOne)          // 후이즈 등록자 연락처, 이메일 주소 확인 이메일 재전송
	e.PATCH("/v1/extra/whois/domain/modify-contact", ctrl.ModiContactWhoisOne)                       // 후이즈 연락처 수정
	e.POST("/v1/extra/whois/domain/idn-available", ctrl.GetDomainIdnAvailableWhoisAll)               // 후이즈 예약가능 여부 확인-IDN
	e.POST("/v1/extra/whois/domain/premium", ctrl.GetDomainPremiumWhoisAll)                          // 후이즈 예약가능 여부 확인-Premium Domains
	e.POST("/v1/extra/whois/domain/third-level-name", ctrl.GetThirdLevelNameWhoisAll)                // 후이즈 예약가능 여부 확인- 3rd level .NAME
	e.POST("/v1/extra/whois/domain/uk", ctrl.GetUkWhoisAll)                                          // 후이즈 UK 도메인 이름에 대한 연락처 정보 확인
	e.POST("/v1/extra/whois/domain/premium-check", ctrl.GetPremiumCheckWhoisOne)                     // 후이즈에 프리미엄 도메인 확인
	e.POST("/v1/extra/whois/domain/customer-default-ns", ctrl.GetCustomerDefaultNsWhoisOne)          // 후이즈에 고객 기본 서버 이름 가져오기
	e.PUT("/v1/extra/whois/domain/purchase-privacy", ctrl.SetPurchasePrivacyWhois)                   // 후이즈 Purchasing / Renewing Privacy Protection
	e.PATCH("/v1/extra/whois/domain/privacy-protection", ctrl.ModiPrivacyProtectionWhoisOne)         // 후이즈 개인정보보호 상태수정
	e.PUT("/v1/extra/whois/domain/enable-theft-protection", ctrl.SetEnableTheftProtectionWhoisOne)   // 후이즈 도난방지 잠금 사용
	e.PUT("/v1/extra/whois/domain/disable-theft-protection", ctrl.SetDisableTheftProtectionWhoisOne) // 후이즈 도난방지 잠금 해제
	e.POST("/v1/extra/whois/domain/locks", ctrl.GetLockWhoisOne)                                     // 후이즈에 도메인 잠금 목록 가져오기
	e.PATCH("/v1/extra/whois/domain/tel/whois-pref", ctrl.ModiTelWhoisPrefWhoisOne)                  // 후이즈 TEL Whois Preference 수정하기
	e.PUT("/v1/extra/whois/domain/uk", ctrl.SetUkWhoisOne)                                           // 후이즈 UK 도메인 이름 해제
	e.POST("/v1/extra/whois/domain/recheck-ns", ctrl.GetRecheckNsWhoisOne)                           // 후이즈 Rechecking NS with .DE Registry
	e.POST("/v1/extra/whois/domain/dotxxx/association-details", ctrl.GetDotxxxAssociationDetailsOne) // 후이즈 DotxxxAssociationDetails
	e.PUT("/v1/extra/whois/domain/add-dnssec", ctrl.SetDnsSecWhoisOne)                               // 후이즈 위임자 서명자 (DS) 레코드 추가
	e.DELETE("/v1/extra/whois/domain/del-dnssec", ctrl.DelDnsSecWhoisOne)                            // 후이즈 위임자 서명자 (DS) 레코드 삭제
	e.PUT("/v1/extra/whois/domain/preordering", ctrl.SetPreorderingWhoisOne)                         // 후이즈 희망 목록에 도메인 이름 추가
	e.DELETE("/v1/extra/whois/domain/preordering", ctrl.DelPreorderingWhoisOne)                      // 후이즈 희망 목록에 도메인 이름 삭제
	e.POST("/v1/extra/whois/domain/preordering", ctrl.GetPreorderingWhoisOne)                        // 후이즈 희망 목록 가져 오기
	e.POST("/v1/extra/whois/domain/preordering-category", ctrl.GetPreorderingCategoryWhoisAll)       // 후이즈 카테고리를 기반으로 희망목록 TLD 가져 오기
	e.POST("/v1/extra/whois/domain/get-tm-notice", ctrl.GetTmNoticeWhoisOne)                         // 후이즈 상표권 주장 데이터 가져 오기
	e.POST("/v1/extra/whois/domain/tlds-in-phase", ctrl.GetTldsInPhaseWhoisAll)                      // 후이즈 Fetching the List of TLDs in Sunrise / Landrush Period
	e.POST("/v1/extra/whois/domain/tld-info", ctrl.GetTldInfoWhoisOne)                               // 후이즈 단계별로 서명 된 TLD 세부 정보 얻기
	e.PUT("/v1/extra/whois/domain/customer-v2", ctrl.SetCustomerV2WhoisOne)                          // 후이즈 희망 목록에 도메인 이름 추가
	e.POST("/v1/extra/whois/domain/available-balance", ctrl.GetAvailableBalanceOne)                  // 후이즈 예치금 조회

	// Gabia
	e.POST("/v1/extra/gabia/domain-check", ctrl.GetGabiaDomainCheck)                  // 도메인 등록 여부
	e.PUT("/v1/extra/gabia/domains", ctrl.SetGabiaDomain)                             // 도메인 등록
	e.DELETE("/v1/extra/gabia/domains", ctrl.DelGabiaDomain)                          // 도메인 삭제(취소)
	e.POST("/v1/extra/gabia/domains", ctrl.GetGabiaDomain)                            // 도메인 정보 조회
	e.PATCH("/v1/extra/gabia/domains-admin", ctrl.ModiGabiaDomainAdmin)               // 관리자 정보 변경
	e.PATCH("/v1/extra/gabia/domains-nameservers", ctrl.ModiGabiaDomainNameServer)    // 네임서버 변경
	e.PATCH("/v1/extra/gabia/domains-owners", ctrl.ModiGabiaDomainOwner)              // 소유자 정보 변경
	e.PATCH("/v1/extra/gabia/domains-status-lock", ctrl.ModiGabiaDomainStatusLock)    // 도메인 잠금 상태 변경
	e.PUT("/v1/extra/gabia/domain-increase", ctrl.SetGabiaDomainIncrease)             // 도메인 연장
	e.DELETE("/v1/extra/gabia/domain-increase", ctrl.DelGabiaDomainIncrease)          // KR 도메인 연장 취소
	e.PUT("/v1/extra/gabia/domain-restore", ctrl.SetGabiaDomainRestore)               // 도메인 삭제 복구 신청
	e.POST("/v1/extra/gabia/domain-restore-check", ctrl.GetGabiaDomainRestoreCheck)   // 도메인 삭제 복구 신청 가능 여부 조회
	e.PUT("/v1/extra/gabia/domain-transfer", ctrl.SetGabiaDomainTransfer)             // 도메인 안이전
	e.POST("/v1/extra/gabia/domain-transfer-check", ctrl.GetGabiaDomainTransferCheck) // 도메인 안이전 조회
	e.POST("/v1/extra/gabia/domain-host-check", ctrl.GetGabiaDomainHostCheck)         // 호스트 등록 여부
	e.PUT("/v1/extra/gabia/host", ctrl.SetGabiaHost)                                  // 호스트 생성
	e.POST("/v1/extra/gabia/host", ctrl.GetGabiaHost)                                 // 호스트 정보 조회
	e.PATCH("/v1/extra/gabia/host", ctrl.ModiGabiaHost)                               // 호스트 정보 변경
	e.DELETE("/v1/extra/gabia/host", ctrl.DelGabiaHost)                               // 호스트 삭제
	e.POST("/v1/extra/gabia/deposits", ctrl.GetGabiaDeposit)                          // 예치금 조회
	//e.POST("/v1/extra/gabia/get-auth-token", ctrl.GetGabiaAuthToken)                  // 가비아 토큰 인증
	e.POST("/v1/extra/gabia/testdomaincheck", ctrl.TestGetGabiaDomainCheck)

	// 세금 계산서
	e.POST("/v1/extra/freebill/get-auth", ctrl.GetFreebillAuth)                                // Freebill 인증키 발급
	e.PUT("/v1/extra/freebill/register", ctrl.SetFreeBillRegister)                             // 전자 세금 계산서 등록
	e.POST("/v1/extra/freebill/search", ctrl.GetFreeBillSearch)                                // 전자 세금 계산서 확인
	e.POST("/v1/extra/freebill/search-status", ctrl.GetFreeBillSearchStatus)                   // 전자 세금 계산서 상태
	e.DELETE("/v1/extra/freebill/delete", ctrl.DelFreeBill)                                    // 전자 세금 계산서 삭제
	e.POST("/v1/extra/freebill/view", ctrl.GetFreeBillView)                                    // 문서 보기(VIEW)
	e.DELETE("/v1/extra/freebill/cancel", ctrl.DelFreeBillCancel)                              // 문서 취소(CANCEL)
	e.PUT("/v1/extra/freebill/publishnow/previous-register", ctrl.GetFreeBillPreviousRegister) // 기등록 문서 전자발행 요청
	e.PUT("/v1/extra/freebill/publishnow", ctrl.SetFreeBillPublishNow)                         // 세금계산서 등록 + 전자발행 요청
	e.POST("/v1/extra/freebill/send", ctrl.SetFreeBillSend)                                    // 전자 세금 계산서 국세청 수동 신고

	// 이니시스
	// e.GET("/v1/extra/inicis/payment-module-test", ctrl.GetInicisPaymentModuleTest)    // 결제 모듈 호출 테스트
	// e.POST("/v1/extra/inicis/payment-module", ctrl.GetInicisPaymentModule)            // 결제 모듈 호출
	e.POST("/v1/extra/inicis/virtual-account", ctrl.GetInicisVirtualAccount) // 가상 계좌
	//e.GET("/v1/extra/inicis/payment-return", ctrl.GetInicisPaymentReturn)             // 결제 승인 응답
	e.POST("/v1/extra/inicis/payment-refund", ctrl.GetInicisPaymentRefund)                // 결제 취소 요청
	e.POST("/v1/extra/inicis/payment-vbank-refund", ctrl.GetInicisPaymentVBankRefund)     // 결제 취소 요청 - 가상계좌
	e.POST("/v1/extra/inicis/payment-vbank-return", ctrl.GetInicisVBankNoti)              // INICIS 가상계좌 입금 noti
	e.POST("/v1/extra/inicis/payment-vbank-return-mobile", ctrl.GetInicisVBankNotiMobile) // INICIS 가상계좌 입금 noti - Mobile
	e.GET("/v1/extra/inicis/payment-cancel", ctrl.GetInicisPaymentCancel)                 // 결제 취소 응답
	e.PUT("/v1/extra/inicis/cash-receipt", ctrl.GetInicisCashReceipt)                     // INICIS 현금영수증 발행
	e.POST("/v1/extra/inicis/payment-request", ctrl.GetInicisPaymentModule)               // 결제 모듈 호출
	e.POST("/v1/inicisReturn", ctrl.GetInicisPaymentReturn)                               // 결제 요청 리턴
	e.POST("/v1/mobile-inicisReturn", ctrl.GetInicisPaymentMobileReturn)                  // 모바일 결제 요청 리턴

	// PayPal
	e.POST("/v1/extra/paypal/payment-request", ctrl.GetPaypalPaymentRequestParam) // 페이팔 결제 요청.
	//e.POST("/v1/extra/paypal/payment-request", ctrl.GetPaypalPaymentRequestAll)   // 페이팔 결제 요청.
	//e.GET("/v1/extra/paypal/payment-return", ctrl.GetPaypalPaymentReturnOne) // 페이팔 결제 승인 요청
	//e.DELETE("/v1/extra/paypal/payment-refund", ctrl.DelPaypalRefundRequestOne)   // 페이팔 결체 취소 요청
	//e.GET("/v1/extra/paypal/payment-cancel", ctrl.GetPaypalPaymentCancelOne) // 페이팔 결제 취소 승인 요청
	//e.POST("/v1/extra/paypal/paymentReturnUrl1", ctrl.GetPaypalPaymentReturnOne2)    // test
	e.POST("/v1/extra/paypal/payment-return", ctrl.GetPaypalPaymentReturn) // 페이팔 결제 응답
	e.POST("/v1/extra/paypal/payment-refund", ctrl.GetPaypalRefundRequest) // 페이팔 환불 요청

	// Alipay
	e.POST("/v1/extra/alipay/payment-request", ctrl.GetAlipayPaymentRequestAll) // Alipay 결제 요청
	e.GET("/v1/extra/alipay/return", ctrl.GetAlipayReturn)                      // Alipay 결제 승인 결과
	e.DELETE("/v1/extra/alipay/payment-refund", ctrl.GetAlipayPaymentRefundOne) // Alipay 결제 취소 요청
	e.POST("/v1/extra/alipay/callback", ctrl.GetAlipayCallback)
	e.GET("/v1/extra/alipay/refer", ctrl.GetAlipayRefer)

	//e.POST("/v1/pageTest", ctrl.PostTest) // test
	e.GET("/v1/pageTest", ctrl.GetTest) // test
	// e.GET("/v1/extra/alipay/test", ctrl.Test)
	// e.GET("/v1/extra/alipay/test2", ctrl.Test2)

	// Alipay
	//e.POST("/v1/extra/alipay/payment-request", ctrl.GetAlipayPayment) // alipay 결제
	//e.POST("/v1/extra/alipay/payment-cancle", ctrl.GetAlipayCancle)   // alipay 결제 취소
	// e.POST("/v1/extra/inicis/payment-request", ctrl.GetInicisPaymentModule2) // 결제 모듈 호출
	// e.GET("/v1/inicisReturn", ctrl.GetInicisPaymentModule4)                  // 결제 모듈 호출
	// e.POST("/v1/extra/inicis/payment-response", ctrl.GetInicisPaymentModule3)

	e.PUT("/v1/alipay-test", ctrl.GetAlipayPaymentTest)            // GetAlipayPaymentTest
	e.POST("/v1/extra/whoxy/domain-info", ctrl.GetDomainInfoWhoxy) // 도메인 상세 정보 가져오기

	// Whoxy

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
