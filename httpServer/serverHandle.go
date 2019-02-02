package httpServer

import (
	. "HttpMockServer/common/loginfo"
	"HttpMockServer/httpProxy"
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
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

	//isNeedMock := true  //调用mock判断接口
	//if isNeedMock{
	//	responeBoby := mock.MockRespone()
	//	rw.Write(responeBoby)
	//	return
	//}


	//proxy http
	bodyString,err:=httpProxy.HttpPoxy(rw,req)
	fmt.Print(bodyString)


	//respone请求处理
	if err!=nil {
		Log.Debugf("uid:%s Received respone  errMsg:%s\n",uid,err)
	}else{
		Log.Debugf("uid:%s Received respone  Msg:%s\n",uid,bodyString)
	}

}
