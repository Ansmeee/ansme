package home

import (
	"ansme/src/lib/response"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)

	data["info"] = "Hello Go !"
	response.Success(w, data)
	return
}
