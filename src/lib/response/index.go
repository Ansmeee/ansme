package response

import (
	"ansme/src/utils/logger"
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Error(w http.ResponseWriter, msg string, code int) {
	response := new(ResponseData)
	response.Code = code
	response.Msg = msg

	jsonData, error := json.Marshal(response)
	if error != nil {
		logger.Error(error.Error())
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	SetHeaders(headers, w)

	w.Write(jsonData)
	return
}

func Success(w http.ResponseWriter, data interface{}) {

	response := new(ResponseData)
	response.Code = 200
	response.Msg = "OK"
	response.Data = data

	jsonData, error := json.Marshal(response)
	if error != nil {
		logger.Error(error.Error())
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	SetHeaders(headers, w)

	w.Write(jsonData)
	return
}

func SetHeaders(headers map[string]string, w http.ResponseWriter) {

	for key, val := range headers {
		w.Header().Set(key, val)
	}

}
