package main

import (
	"net/http"
)

func main() {
	// 指定路由
	http.Handle("/home", &HomeHandler{})

	// 启动http服务
	http.ListenAndServe(":8000", nil)
}

type HomeHandler struct {}

// 实现ServeHTTP
func (h *HomeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello World"))
}

