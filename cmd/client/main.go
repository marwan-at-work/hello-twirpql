package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	hellotwirpql "marwan.io/hello-twirpql"
)

func main() {
	client := hellotwirpql.NewGreeterJSONClient("http://localhost:8080", http.DefaultClient)
	resp, err := client.Hello(context.Background(), &hellotwirpql.HelloReq{
		Name: "Gophers",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetText())
}
