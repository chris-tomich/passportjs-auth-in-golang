package passportjs4go

import (
	"strings"
	"errors"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type CookieSessionId string

func (c CookieSessionId) SplitSessionIdCookie() (sessionId string, signature string, err error) {
	if len(string(c)) > 2 {
		sessionIdAndSignature := strings.Split(string(c), ":")[1]
		indexOfSeparator := strings.Index(sessionIdAndSignature, ".")

		sessionId = sessionIdAndSignature[:indexOfSeparator]
		signature = sessionIdAndSignature[(indexOfSeparator + 1):]
	} else {
		err = errors.New("Cookie is unexpectedly short.")
	}

	return
}

func IsSessionIdOk(key string, sessionId string, reportedSignature string) bool {
	sessionIdBytes := []byte(sessionId)
	keyBytes := []byte(key)

	sha256Signer := hmac.New(sha256.New, keyBytes)
	sha256Signer.Write(sessionIdBytes)
	calculatedSignature := sha256Signer.Sum(nil)
	calculatedBase64Signature := strings.TrimRight(base64.StdEncoding.EncodeToString(calculatedSignature), "=")

	return reportedSignature == calculatedBase64Signature
}

