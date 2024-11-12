package main

import (
	"net/http"

	arri "github.com/modiimedia/arri/languages/go/go-server"
)

// extend this with any other context you want to be available to Arri procedures
// this must fullfil the arri.Context interface
type RpcContext struct {
	w http.ResponseWriter
	r *http.Request
}

func (c RpcContext) Request() *http.Request {
	return c.r
}

func (c RpcContext) Writer() http.ResponseWriter {
	return c.w
}

// create the RpcContext using the incoming response writer and http request
func CreateRpcContext(w http.ResponseWriter, r *http.Request) (*RpcContext, arri.RpcError) {
	return &RpcContext{w: w, r: r}, nil
}
