package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	httpURL := "http://localhost:18080/hello"

	//初始化hystrix配置
	hystrix.ConfigureCommand("svc1", hystrix.CommandConfig{
		MaxConcurrentRequests : 5,  //设置请求的并发数量
		RequestVolumeThreshold : 4,  //配置断路器断开需要的最少请求失败数量
	})


	for i := 0 ; i < 10 ; i++ {
		for j := 0 ; j< 10 ; j++ {
			go func() {
				req,err := httpDo(httpURL)
				if err != nil {
					log.Println("http request error :",err.Error())
				}else {
					log.Println("http request success:",string(req))
				}
			}()
		}
		time.Sleep(time.Second*5)
	}

}

//http请求
func httpDo(url string) (respBytes []byte, err error) {
	var resp *http.Response
	err = hystrix.Do("svc1",func() error {
		resp,err = http.Get(url)
		if err != nil {
			return err
		}
		respBytes,err = ioutil.ReadAll(resp.Body)
		if err !=nil {
			return err
		}

		return nil
	}, func(error) error {
		return errors.New("http request error")
	})

	return
}