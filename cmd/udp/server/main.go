package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"test-station/pkg/workspace"
)

var (
	sockAddr = workspace.InstallPathJoin(workspace.UnixSockSpace, "test.sock")
)
func main(){
	router:=gin.New()
	router.GET("/testGet",handlerGet)
	err := os.Remove(sockAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	unixAddr, err := net.ResolveUnixAddr("unix", sockAddr)
	if err!=nil{
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUnix("unix", unixAddr)
	if err!=nil{
		fmt.Println("listening error:",err)
	}
	fmt.Println("listening...")
	_ = http.Serve(listener, router)
}

func handlerGet(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"resp":"ok",
	})
}

