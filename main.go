package main

import (
	"galleryApp_golang/controller"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {

	staticC := controller.NewStatic()
	userC := controller.NewUser()



	r := mux.NewRouter()
	r.Handle("/", staticC.HomeView).Methods("GET")
	r.Handle("/contact", staticC.ContactView).Methods("GET")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup",userC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)

}

func must(err error) {
	if err != nil{
		panic(err)
	}
}