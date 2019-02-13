package httpServer

import (
	. "HttpMockServer/common/loginfo"
	"HttpMockServer/dbOpr"
	"HttpMockServer/httpProxy"
	"HttpMockServer/mock"
	"bytes"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {


	//req请求处理
	var reqBodyBytes []byte
	if req.Body != nil {
		reqBodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes))

	uid,err := uuid.NewV4()
	Log.Debugf("uid:%s Received request  Method:%s Host:%s RemoteAddr:%s RequestURI:%s \n"+
		"Msg:%s\n",uid, req.Method, req.Host, req.RemoteAddr,req.RequestURI,reqBodyBytes)

	RemoteAddr := strings.Split(req.RemoteAddr, ":")

	if RemoteAddr[0] =="132.232.96.57"{

		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	dbOpr.Db.Exec("insert into httpRequest(uuid,host,method,remoteIp,requestUrl,requestBoby) values(?,?,?,?,?,?)",
		uid,req.Host,req.Method,RemoteAddr[0],req.RequestURI,string(reqBodyBytes))
	//isNeedMock := true  //调用mock判断接口
	//
	respMock :=mock.IsNeedMock(req.Host,req.RequestURI,RemoteAddr[0],req.Method)
	//respMock:=""
	if "" != respMock {
		//根据实际情况来，如果有2种的，后续在页面做改造(增加转JSON和不转JSON)
		//resMsg, errMock := json.Marshal(respMock)
		//if errMock !=nil  {
		//	Log.Errorf("Mock失败:%v",errMock)
		//	return
		//}
		resMsg := []byte(respMock)
		rw.Write(resMsg)
		dbOpr.Db.Exec("insert into httpRespone(uuid,responeBoby) values(?,?)",uid,respMock)


	}else{
		bodyString, errProxxy:=httpProxy.HttpPoxy(rw,req)
		if errProxxy != nil{
			Log.Errorf("代理失败:%v",err)
			return
		}
		dbOpr.Db.Exec("insert into httpRespone(uuid,responeBoby) values(?,?)",uid,bodyString)
	}



	//proxy http




}
