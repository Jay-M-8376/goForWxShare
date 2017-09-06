package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	appid  string
	secret string
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("you have to set the appkey&appSecret!!!")
	}
}

func main() {
	if len(os.Args) > 1 {
		appid = os.Args[1]
		secret = os.Args[2]
		_buildServer()
	} else {
		return
	}

}

func _buildServer() {
	http.HandleFunc("/api/share", HandleShareAction)
	http.ListenAndServe(":8999", nil)
}
