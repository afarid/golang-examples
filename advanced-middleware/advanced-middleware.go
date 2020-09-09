package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Logging() Middleware  {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {log.Println(request.URL.Path, time.Since(start))}()
			handlerFunc(writer, request)
		}	
	}
}

func Method(m string) Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if request.Method != m {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
			handlerFunc(writer,request)
		}
	}
}

func Chain(handlerFunc http.HandlerFunc, middlewares ...Middleware)  http.HandlerFunc{
	for _, middleware := range middlewares {
		handlerFunc = middleware(handlerFunc)
	}
	return handlerFunc
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello world")
}
func main() {
	http.HandleFunc("/", Chain(hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}