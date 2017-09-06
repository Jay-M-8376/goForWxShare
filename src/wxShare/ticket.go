package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AccessToken struct {
	Access_token string `json:"access_token"`
	Expires_in   int64  `json:"expires_in"`
	Errcode      int32  `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

type JsTicket struct {
	Errcode    int32  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Ticket     string `json:"ticket"`
	Expires_in int64  `json:"expires_in"`
}

func (a *AccessToken) getTokenAndDate() string {
	return a.Access_token
}

func (jt *JsTicket) getTicketResult() string {
	return jt.Ticket
}

func getTicket(id string, sec string) string {
	ticketData := getData("ticket")
	if len(ticketData) != 0 {
		ticketDataSlice := strings.Split(ticketData, "|")
		currentTime := time.Now().Unix()
		extime, _ := strconv.ParseInt(ticketDataSlice[1], 10, 64)
		if extime-currentTime > 0 {
			return ticketDataSlice[0]
		}
	}
	token := getToken(id, sec)
	ticket := getJsTicket(token)
	expirse := strconv.FormatInt(time.Now().Unix()+7200, 10)
	setData("ticket", ticket+"|"+expirse)
	return ticket
}

func getJsTicket(token string) string {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + token + "&type=jsapi")
	if err != nil {
		fmt.Println("fail to get js_ticket")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var ticketCont JsTicket
	err = json.Unmarshal([]byte(body), &ticketCont)
	if err != nil {
		fmt.Println("something wrong in decode Request at ticket step")
	}
	ticket := ticketCont.getTicketResult()
	return ticket
}

func getToken(id string, sec string) string {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + id + "&secret=" + sec)
	if err != nil {
		fmt.Println("fail to get access_token")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var token AccessToken
	err = json.Unmarshal([]byte(body), &token)
	if err != nil {
		fmt.Println("something wrong in decode Request")
	}
	tok := token.getTokenAndDate()
	return tok
}
