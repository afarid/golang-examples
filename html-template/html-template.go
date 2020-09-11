package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"html/template"
)

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

type Todo struct {
	Title string
	Done bool
}

var (
	log *logrus.Logger
)
func main()  {

	data := TodoPageData{
		PageTitle: "My To do list",
		Todos:  []Todo{
			{Title: "task 1", Done: true},
			{Title: "task 2", Done: false},
			{Title: "task 3", Done: true},
		},
	}

	tmpl, err := template.ParseFiles("html.tmpl")
	if err != nil {
		log.Fatalln("unable to parse template", err)
	}


	router := mux.NewRouter()
	router.HandleFunc("/todo", func(writer http.ResponseWriter, request *http.Request) {
		tmpl.Execute(writer, data)
	} )

	http.ListenAndServe(":8008", router)


}
