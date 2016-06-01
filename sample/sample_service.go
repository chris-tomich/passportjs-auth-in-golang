package main

import (
	"goji.io"
	"github.com/chris-tomich/passportjs4go"
	"golang.org/x/net/context"
	"net/http"
	"goji.io/pat"
)

const (
	key = "My-Secret-Key"
	redisAddress = "172.16.1.2:6379"
	redisStoreIndex = 0
)

func main() {
	mux := goji.NewMux()
	mux.UseC(passportjs4go.AuthorizationMiddleware(key, redisAddress, redisStoreIndex))
	mux.HandleFuncC(pat.Get("/"), HelloAuthorizedUser)
	http.ListenAndServe(":8080", mux)
}

func HelloAuthorizedUser(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	isAuthorized, found := ctx.Value("IsAuthorized").(bool)

	if found && isAuthorized {
		res.Write([]byte("Hello authorized user!"))
	} else {
		res.Write([]byte("You don't have authorization!"))
	}
}