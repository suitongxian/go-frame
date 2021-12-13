package main

import (
	"engine"
	"fmt"
	"net/http"
)

func main() {
	e := engine.New()
	e.Get("/get", func(ctx *engine.Context) {
		ctx.String(http.StatusOK, "this is get")
	})
	group1 := e.Group("/prefix")
	group1.Get("/get", func(ctx *engine.Context) {
		ctx.String(http.StatusOK, "this is prefix/get")
	})
	e.Post("/post", func(ctx *engine.Context) {
		fmt.Println("this is prefix/post")
	})
	e.Run("127.0.0.1:8080")
}
