package http

import (
	"github.com/namjagbrawa/go_chat/exception"
	"net/http"
)

func CrossDomain(handler exception.AppHandler) exception.AppHandler {

	return func(writer http.ResponseWriter, request *http.Request) error {
		writer.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		// writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		// writer.Header().Set("content-type", "application/json") //返回数据格式是json
		return handler(writer, request)
	}

}
