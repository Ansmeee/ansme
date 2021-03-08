package middleware

import (
	"ansme/src/config"
	"ansme/src/utils/logger"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var form = make(map[string][]string)
var maxMemory = int64(1024 * 1024 * 10)

func Form(r *http.Request) map[string][]string {
	requestMethod := strings.ToLower(r.Method)

	switch requestMethod {
	case "get":
		return parseGetForm(r)

	case "put":
	case "post":
		return parsePostForm(r)
	}

	return form
}

func parsePostForm(r *http.Request) map[string][]string {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var jsonForm map[string]string
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&jsonForm)

		for key, val := range jsonForm {
			valSlice := make([]string, 1)
			valSlice[0] = val
			form[key] = valSlice
		}
	}

	r.ParseMultipartForm(maxMemory)
	for key, val := range r.PostForm {
		form[key] = val
	}

	r.ParseForm()
	for key, val := range r.Form {
		form[key] = val
	}

	return form
}

func parseGetForm(r *http.Request) map[string][]string {
	r.ParseForm()

	for key, val := range r.Form {
		form[key] = val
	}

	return form
}

func Authorize(r *http.Request) bool {
	headers := r.Header

	var nonce = ""
	nonceMap := headers["X-Nonce"]
	if len(nonceMap) > 0 {
		nonce = nonceMap[0]
	}

	fmt.Println("auth nonce")
	if false == authorizeNonce(nonce) {
		logger.Error("Authorize authorizeNonce failed")
		return false
	}

	var sign = ""
	signMap := headers["X-Sign"]
	if len(signMap) > 0 {
		sign = signMap[0]
	}

	fmt.Println("auth sign")
	if false == authorizeSign(sign, nonce) {
		return false
	}

	return true
}

func authorizeNonce(formNonce string) bool {

	if "" == formNonce {
		logger.Error("authorizeNonce failed: formNonce is empty")
		return false
	}

	formTimeUnix, error := strconv.Atoi(formNonce)
	if error != nil {
		logger.Error(error.Error())
		return false
	}

	currentTimeStr := strconv.FormatInt(time.Now().Unix(), 10)
	currentTimeUnix, error := strconv.Atoi(currentTimeStr)
	if error != nil {
		logger.Error(error.Error())
		return false
	}

	expiredTime, _ := strconv.Atoi(config.Get("expired_time"))
	if (currentTimeUnix >= formTimeUnix) && ((currentTimeUnix - formTimeUnix) < expiredTime) {
		return true
	}

	logger.Error(fmt.Sprintf("authorizeNonce failed: formTimeUnix: %d, currentTimeUnix: %d", formTimeUnix, currentTimeUnix))

	return false
}

func authorizeSign(formSign, formNonce string) bool {

	appSecret := config.Get("app_secret")
	signData := []byte(appSecret + formNonce)
	sign := fmt.Sprintf("%X", md5.Sum(signData))

	fmt.Println(formSign, sign, formSign == sign)

	if formSign == sign {
		return true
	}


	logger.Error(fmt.Sprintf("authorizeSign failed: sign: %s, formSign: %s", sign, formSign))

	return false
}
