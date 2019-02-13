package main

import (
	. "HttpMockServer/common/loginfo"
	"HttpMockServer/dbOpr"
	"HttpMockServer/httpServer"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"os"
	"path/filepath"
)




func inithttpServer(){
	ports := []string{":80", ":443",":8088"}
	for _, v := range ports {
		go func(port string) { //每个端口都扔进一个goroutine中去监听
			if ":80" == port {
				mux := http.NewServeMux()
				mux.HandleFunc("/", httpServer.ServeHTTP)
				http.ListenAndServe(port, mux)
			}else if  ":8088" == port {
				mux := http.NewServeMux()
				mux.HandleFunc("/resetDataFilter", httpServer.ResetDataFilter)
				mux.HandleFunc("/clearFilter", httpServer.ClearFilter)
				mux.HandleFunc("/addMock", httpServer.AddMock)
				mux.HandleFunc("/loadMockData", httpServer.LoadMockData)
				mux.HandleFunc("/modifyMock", httpServer.ModifyMock)
				mux.HandleFunc("/deleteMock", httpServer.DeleteMock)

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
	//初始化日志
	path, _ := filepath.Abs("./log.txt")
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	defer logFile.Close()
	if err != nil{
		fmt.Println(err)
	}
	InitLog(logFile,logging.DEBUG)

	//初始化数据库连接
	dbCreateError := dbOpr.CreateDbLink()
	if dbCreateError != nil{
		Log.Error(dbCreateError)
		os.Exit(-1)
	}
	defer dbOpr.Db.Close()

	//初始化httpServer
	inithttpServer()




	select {
	}


}

