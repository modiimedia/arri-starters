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
		arri.AppOptions[RpcEvent]{
			OnRequest: func(c *RpcEvent) arri.RpcError {
				c.Writer().Header().Set("Access-Control-Allow-Origin", "*")
				return nil
			},
		},
		// initialize an RpcEvent using the incoming response writer and http request
		// you can extend RpcEvent in ./rpc_event.go
		func(w http.ResponseWriter, r *http.Request) (*RpcEvent, arri.RpcError) {
			return &RpcEvent{w: w, r: r}, nil
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

func SayHello(params GreetingParams, event RpcEvent) (GreetingResponse, arri.RpcError) {
	return GreetingResponse{Message: fmt.Sprintf("Hello %s", params.Name)}, nil
}

func SayGoodbye(params GreetingParams, event RpcEvent) (GreetingResponse, arri.RpcError) {
	return GreetingResponse{Message: fmt.Sprintf("Goodbye %s", params.Name)}, nil
}
