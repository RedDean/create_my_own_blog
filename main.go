package main

import (
	"MyBlog/controller"
	"log"
	"net/http"
)

func main() {
	router := controller.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
