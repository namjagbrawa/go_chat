package chat

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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
	m := Message{}
	err := xml.Unmarshal(body, &m)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("receive message parse to xml : ", m)
	switch m.MsgType {
	case Text:
		r_msg := Message{FromUserName: m.ToUserName,
			ToUserName: m.FromUserName,
			CreateTime: int(time.Now().Unix()),
			MsgType:    Text,
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
