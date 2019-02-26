package main

import (
	"fmt"
	"github.com/KarenLKL/studygolang/exceptionhandler/file"
	"log"
	"net/http"
	"os"
	"time"
)

var logger *log.Logger

type userError interface {
	error
	Message() string
}

type httpHandleFunc func(writer http.ResponseWriter, request *http.Request) error

/**
服务端统一错误处理
 */
func httpHandleFuncWrap(handleFunc httpHandleFunc) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("request url:%s ", request.URL)
		defer func() {
			if r := recover(); r != nil {
				logger.Printf("panic: %v ", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handleFunc(writer, request)
		if err != nil {
			code := http.StatusOK
			if userError, ok := err.(userError); ok {
				http.Error(writer,userError.Message(),http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", httpHandleFuncWrap(file.Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func init() {
	logFile, err := os.Create("./" + time.Now().Format("20060102") + ".txt");
	if err != nil {
		fmt.Println(err);
	}
	logger = log.New(logFile, "test_", log.Ldate|log.Ltime|log.Lshortfile);
}
