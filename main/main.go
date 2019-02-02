package main

import (
	. "HttpMockServer/common/loginfo"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"os"
	"path/filepath"
	"HttpMockServer/httpServer"
)




func inithttpServer(){
	ports := []string{":80", ":443"}
	for _, v := range ports {
		go func(port string) { //每个端口都扔进一个goroutine中去监听
			if ":80" == port {
				mux := http.NewServeMux()
				mux.HandleFunc("/", httpServer.ServeHTTP)
				http.ListenAndServe(port, mux)
			}else{
				//后续开发
				//mux := http.NewServeMux()
				//mux.Handle("/", &httpProxy.Pxy{})
				//http.ListenAndServeTLS("","",port, mux)
			}

		}(v)
	}
}


func main() {
	//初始化
	path, _ := filepath.Abs("./log.txt")
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	defer logFile.Close()
	if err != nil{
		fmt.Println(err)
	}
	InitLog(logFile,logging.DEBUG)

	//初始化httpServer
	inithttpServer()


	select {
	}


}

