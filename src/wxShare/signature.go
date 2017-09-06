package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type authData struct {
	AppId     string
	Timestamp int64
	NonceStr  string
	Signature string
}

func (a *authData) getAuthData() string {
	res, _ := json.Marshal(a)
	return string(res)
}

func getSignature(id string, sec string, url string) *authData {
	ticket := getTicket(id, sec)

	randomStr := getRandom(16)
	timestamp := time.Now().UnixNano() / 1e6

	ticket_origin := "jsapi_ticket=" + ticket + "&noncestr=" + randomStr + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&url=" + url

	sign := hashSha1(ticket_origin)

	fmt.Println(sign)

	data := &authData{
		AppId:     id,
		Timestamp: timestamp,
		NonceStr:  randomStr,
		Signature: sign,
	}

	return data
}
