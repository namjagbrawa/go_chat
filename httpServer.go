package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func WechatHandlerAuth(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	m := make(map[string]string)

	err := json.Unmarshal(body, &m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}


	fmt.Println("m:", m)

	for k,v :=range m {
		fmt.Println(k, ":", v)
	}
}

func main() {
	http.HandleFunc("/chat", IndexHandler)
	http.HandleFunc("/chat/wechat", WechatHandlerAuth)
	http.ListenAndServe("127.0.0.1:8088", nil)
}
