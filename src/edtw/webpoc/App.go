package main

import (
	"net/http"
	"edtw/templates"
	"log"
	"fmt"
)

type ErrorData struct {
	Msg string
}

func main() {
	http.HandleFunc("/",reqHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func reqHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet: doGet(w,r)
		case http.MethodPost: doPost(w,r)
		default: handleError(w,r,&ErrorData{
			Msg: fmt.Sprintf("Method '%s' for ulr '%s' is not supported",
				r.Method, r.URL),
		})
	}
}


func doGet(w http.ResponseWriter, r *http.Request){
	data := new(struct{})
	templates.Cache.ExecuteTemplate(w,"form.html", &data)
}

func doPost(w http.ResponseWriter, r *http.Request){
	data := new(struct{})
	templates.Cache.ExecuteTemplate(w,"form.html", &data)
}

func handleError(w http.ResponseWriter, r *http.Request,data *ErrorData ){
	templates.Cache.ExecuteTemplate(w,"error.html", data )
}