package main

import (
	"fmt"
	"github.com/namjagbrawa/go_chat/chat"
	"github.com/namjagbrawa/go_chat/exception"
	crossDomain "github.com/namjagbrawa/go_chat/http"
	"github.com/namjagbrawa/go_chat/picture"
	"github.com/namjagbrawa/go_chat/tools"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
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
	http.HandleFunc("/chat/wechat", chat.WechatHandlerAuth)
	http.HandleFunc("/tools/ebbinghaus", exception.ErrWrapper(crossDomain.CrossDomain(tools.Ebbinghaus)))
	http.HandleFunc(picture.URL_PICTURE_UPLOAD, exception.ErrWrapper(crossDomain.CrossDomain(picture.UpdateHandler)))
	http.HandleFunc(picture.URL_PICTURE_VIEW, exception.ErrWrapper(crossDomain.CrossDomain(picture.ViewHandler)))
	http.ListenAndServe("127.0.0.1:8088", nil)
}
