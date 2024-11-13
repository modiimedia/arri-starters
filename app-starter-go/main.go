package main

import (
	"fmt"
	"log"
	"net/http"

	arri "github.com/modiimedia/arri/languages/go/go-server"
)

func main() {
	app := arri.NewApp(
		http.DefaultServeMux,
		arri.AppOptions[RpcContext]{
			OnRequest: func(r *http.Request, c *RpcContext) arri.RpcError {
				r.Header.Set("Access-Control-Allow-Origin", "*")
				return nil
			},
		},
		// create the RpcContext using the incoming response writer and http request
		func(w http.ResponseWriter, r *http.Request) (*RpcContext, arri.RpcError) {
			return &RpcContext{w: w, r: r}, nil
		},
	)

	// register procedures on the app
	arri.Rpc(&app, SayHello, arri.RpcOptions{})
	arri.Rpc(&app, SayGoodbye, arri.RpcOptions{})

	err := app.Run(arri.RunOptions{Port: 3000})
	if err != nil {
		log.Fatal(err)
	}
}

type GreetingParams struct {
	Name string
}

type GreetingResponse struct {
	Message string
}

func SayHello(params GreetingParams, context RpcContext) (GreetingResponse, arri.RpcError) {
	return GreetingResponse{Message: fmt.Sprintf("Hello %s", params.Name)}, nil
}

func SayGoodbye(params GreetingParams, context RpcContext) (GreetingResponse, arri.RpcError) {
	return GreetingResponse{Message: fmt.Sprintf("Goodbye %s", params.Name)}, nil
}
