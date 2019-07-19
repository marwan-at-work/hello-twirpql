package main

import (
	"context"
	"fmt"
	"net/http"

	hellotwirpql "marwan.io/hello-twirpql"
	"marwan.io/hello-twirpql/twirpql"
)

type server struct{}

func (s *server) Hello(ctx context.Context, req *hellotwirpql.HelloReq) (*hellotwirpql.HelloResp, error) {
	return &hellotwirpql.HelloResp{
		Text: "Hello " + req.GetName(),
	}, nil
}

func main() {
	greeterService := &server{}

	restHandler := hellotwirpql.NewGreeterServer(greeterService, nil)
	http.Handle("/", restHandler)

	gqlHandler := twirpql.Handler(greeterService, nil)
	http.Handle("/query", gqlHandler)
	http.Handle("/play", twirpql.Playground("Hello TwirpQL", "/query"))

	fmt.Println("Playground at http://localhost:8080/play")
	http.ListenAndServe(":8080", nil)
}
