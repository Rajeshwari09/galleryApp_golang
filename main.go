package main

import (
	"fmt"
	"net/http"
)

func handlerFunvc(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprint(w,"<h1>welcome</h1>")
}

func main()  {
	http.HandleFunc("/",handlerFunvc)
	http.ListenAndServe(":3000",nil)

}
