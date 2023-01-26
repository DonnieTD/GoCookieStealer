package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func StealCookie(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	r.Header.Set("X-GoProxy", "yxorPoG-X")
	fmt.Println("I GOT YOU:")
	fmt.Println(r.Cookies())
	return r, nil
}

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	fmt.Println("Started")
	proxy.OnRequest().DoFunc(StealCookie)
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
