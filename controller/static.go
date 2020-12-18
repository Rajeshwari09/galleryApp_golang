package controller

import "galleryApp_golang/view"

type Static struct {
	HomeView    *view.View
	ContactView *view.View
}

func NewStatic() *Static  {
	return &Static{
		HomeView: view.NewView("bootstrap","static/home"),
		ContactView: view.NewView("bootstrap","static/contact"),
	}
}