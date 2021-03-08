package router

import (
	"ansme/src/controller/home"
	"net/http"
)

type Router struct {
	Method  string
	Path    string
	Handler func(writer http.ResponseWriter, request *http.Request)
}

//register routers
var RouterGroup = []Router{
	{"Get", "/home", home.Info},
}
