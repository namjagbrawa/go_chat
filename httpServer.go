package main

import (
	"encoding/xml"
	"fmt"
	"github.com/namjagbrawa/go_chat/message"
	"io/ioutil"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func WechatHandlerAuth(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == http.MethodGet {
		WechatGet(w, r)
	}
	if method == http.MethodPost {
		WechatPost(w, r)
	}

	fmt.Println("method = ", method)

}

func WechatPost(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	m := message.Message{}
	err := xml.Unmarshal(body, &m)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("receive message parse to xml : ", m)
	switch m.MsgType {
	case message.Text:
		r_msg := message.Message{FromUserName: m.ToUserName,
			ToUserName: m.FromUserName,
			CreateTime: int(time.Now().Unix()),
			MsgType:    message.Text,
			Content:    "text"}
		xml, err := xml.Marshal(r_msg)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		w.Write(xml)
	}
}

func WechatGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.URL.Query()
	echostr := query["echostr"][0]
	fmt.Println("echostr = ", echostr)

	fmt.Fprintln(w, echostr)
}

func postParam(req *http.Request) {
	post_act := req.PostFormValue("act")
	fmt.Println(post_act)
	post_act1 := req.PostForm["username"]
	fmt.Println(post_act1)
}

func getParam(req *http.Request) {
	query := req.URL.Query()
	get_act := query["act"][0]
	fmt.Println(get_act)
}

func main() {
	http.HandleFunc("/chat", IndexHandler)
	http.HandleFunc("/chat/wechat", WechatHandlerAuth)
	http.ListenAndServe("127.0.0.1:8088", nil)
}
