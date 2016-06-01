package passportjs4go

import (
	"goji.io"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"fmt"
)

func AuthorizationMiddleware(key string, redisAddress string, redisStoreIndex int) func (goji.Handler) goji.Handler {
	return func (next goji.Handler) goji.Handler {
		middleware := func (ctx context.Context, res http.ResponseWriter, req * http.Request) {
		var newCtx context.Context

		rawCookie, cookieErr := req.Cookie("connect.sid")

		if cookieErr == nil {
		cookie := rawCookie.Value
		rawCookieSessionId := CookieSessionId(cookie)

		if strings.HasPrefix(string(rawCookieSessionId), "s:") {
		sessionId, reportedSignature, err := rawCookieSessionId.SplitSessionIdCookie()

		if err == nil && IsSessionIdOk(key, sessionId, reportedSignature) {
		sessionInfo := &ConnectSessionInfo{}

		sessionInfo.LoadFromRedisStore(redisAddress, redisStoreIndex, sessionId)

		newCtx = context.WithValue(ctx, "IsAuthorized", true)
		} else {
		fmt.Println("The session ID has been tampered with.")
		newCtx = context.WithValue(ctx, "IsAuthorized", false)
		}
		}
		} else {
		newCtx = context.WithValue(ctx, "IsAuthorized", false)
		}
		next.ServeHTTPC(newCtx, res, req)
		}

		return goji.HandlerFunc(middleware)
	}
}