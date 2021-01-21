package controller

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	b64 "encoding/base64"
	"encoding/hex"
	"time"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

// EUC-KR -> UTF-8 로 Decoding
func ConvertEucKrDecoding(data []byte) string {
	var bufs bytes.Buffer

	wr := transform.NewWriter(&bufs, korean.EUCKR.NewDecoder())
	wr.Write(data)
	wr.Close()
	rData := bufs.String()
	lprintf(4, "[INFO] euc-kr decoding : [%s]\n", rData)
	return rData
}

// UTF-8 -> EUC-KR 로 encdoing
func ConvertEucKrEncoding(data []byte) string {
	var bufs bytes.Buffer

	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	wr.Write(data)
	wr.Close()
	rData := bufs.String()
	lprintf(4, "[INFO] euc-kr encoding : [%x]\n", rData)
	return rData
}

// timestamp - millisecond - 13자리
func GetMillSecondTimeStamp(t time.Time) int64 {
	//now := time.Now()
	nanos := t.UnixNano()
	millis := nanos / 1000000
	lprintf(4, "[INFO] timestamp : [%d]\n", millis)
	return millis
}

// yyyymmddhhmmssnnn
func Getyyyymmddhhmmssnnn(t time.Time) string {
	//t := time.Now()
	tFormat := t.Format("20060102150405000")
	return tFormat
}

// sha 256 encoding
func GetSha256Encoding(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	//lprintf(4, "[INFO] data : [%s] sha256 : [%s]\n", data, mdStr)

	return mdStr
}

// sha 512 encoding
func GetSha512Encoding(data string) string {
	hash := sha512.New()     // SHA512 해시 인스턴스 생성
	hash.Write([]byte(data)) // 해시 인스턴스에 데이터 추가
	md := hash.Sum(nil)      // 해시 인스턴스에 저장된 데이터의 SHA512 해시 값 추출
	mdStr := hex.EncodeToString(md)

	return mdStr
}

// base64 256 encoding
func GetBase64Encoding(data string) string {

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}
