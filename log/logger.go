package log

import (
	"log"
	"os"
)

var Logger *log.Logger
var File *os.File
var err error

func init() {
	File, err = os.OpenFile("go_chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger = log.New(File, "", log.LstdFlags)
	Logger.SetPrefix("go_chat ") // 设置日志前缀
	Logger.SetFlags(log.LstdFlags | log.Lshortfile)
	/*
		const (
			// 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
			// 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
			Ldate         = 1 << iota     // 日期：2009/01/23
			Ltime                         // 时间：01:23:23
			Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
			Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
			Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
			LstdFlags     = Ldate | Ltime // 标准logger的初始值
		)
	*/
}
