package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type urlData struct {
	Url string `json:"url"`
}

func (ud *urlData) getUrl() string {
	return ud.Url
}

func HandleShareAction(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("error in request params")
		res.Write([]byte("params error!!!"))
		return
	}
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var ud urlData
	err = json.Unmarshal([]byte(body), &ud)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("content-type", "application/json")
	sign := getSignature(appid, secret, ud.getUrl()).getAuthData()
	res.Write([]byte(sign))
	return
}
