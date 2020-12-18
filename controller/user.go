package  controller

import (
	"fmt"
	"galleryApp_golang/view"
	"net/http"
)
func NewUser() *User  {
	return &User{
		NewView: view.NewView("bootstrap","users/new"),
	}
}

type User struct {
	NewView *view.View
}

func (u *User) New(w http.ResponseWriter, r *http.Request)  {
	u.NewView.Render(w,nil)

}

type SignupForm struct {
	Email string `schema: "email"`
	Password string `schema:"password"`
}

func (u *User) Create(w http.ResponseWriter, r *http.Request)  {

	var form SignupForm
	if err := parceForm(r,&form);err !=nil{
		panic(err)
	}
	fmt.Fprintln(w,form)

}