package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)
func main()  {

	router := mux.NewRouter()
	router.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	})
	router.HandleFunc("/world", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "world")
	})
	http.ListenAndServe(":8080", router)
}
