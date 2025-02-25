package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/hello"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHalloHandler(router)
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	
	fmt.Println("Server is listenong on port 8081")
	server.ListenAndServe()
}