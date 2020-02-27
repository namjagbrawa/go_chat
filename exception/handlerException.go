package exception

import (
	"log"
	"net/http"
)

type UserError string

func (e UserError) Error() string {
	return e.Message()
}

func (e UserError) Message() string {
	return string(e)
}

type AppHandler func(writer http.ResponseWriter, request *http.Request) error

func ErrWrapper(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	//闭包
	return func(writer http.ResponseWriter, request *http.Request) {
		// 如果 发现 panic，判断错误输出错误，否则 继续往上层 panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(
					writer, //向writer汇报错误
					http.StatusText(http.StatusInternalServerError), //错误描述信息（字符串）
					http.StatusInternalServerError)                  //系统内部错误
			}
		}()
		// 执行业务代码操作，上面定义的 defer 就是防止 业务代码中出现 panic
		err := handler(writer, request)

		// 如果业务代码执行出错
		if err != nil {
			//日志输出错误信息
			log.Printf("Error occurred handling request: %s", err.Error())

			//判断错误类型是否为 自定义错误
			if userErr, ok := err.(UserError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			//向writer 中写入错误信息
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}
