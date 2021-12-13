package main

import (
	"engine"
	"fmt"
	"net/http"
)

func main() {
	e := engine.New()
	group1 := e.Group("/prefix")
	group1.Get("/get", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("111")
	})
	e.Post("/post", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("222")
	})
	e.Run("127.0.0.1:8080")
}
