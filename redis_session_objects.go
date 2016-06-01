package passportjs4go

import (
	"github.com/mediocregopher/radix.v2/pool"
	"strconv"
	"encoding/json"
)

type ConnectSessionInfo struct {
	Cookie ConnectCookieInfo `json:"cookie"`
	Passport PassportSessionInfo `json:"passport"`
}

type PassportSessionInfo struct {
	User string `json:"user"`
}

type ConnectCookieInfo struct {
	OriginalMaxAge int `json:"originalMaxAge"`
	Expires string `json:"expires"`
	Secure bool `json:"secure"`
	HttpOnly bool `json:"httpOnly"`
	Path string `json:"path"`
}

func (csi *ConnectSessionInfo) LoadFromRedisStore(address string, storeIndex int, sessionId string) bool {
	pool, _ := pool.New("tcp", address, 10)
	conn, _ := pool.Get()

	redisSessionKey := "sess:" + sessionId

	conn.Cmd("SELECT", strconv.Itoa(storeIndex))
	redisQueryResult := conn.Cmd("GET", redisSessionKey)
	redisSessionInfo, _ := redisQueryResult.Bytes()

	json.Unmarshal(redisSessionInfo, csi)

	return true
}