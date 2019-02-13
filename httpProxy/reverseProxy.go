package httpProxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)




func addParams()([]byte){
	params := make(map[string]interface{})
	params["aid"] = 1
	params["page"] = 0
	params["status"] = "ongoing"
	params["size"] = 20

	bytesData, _:= json.Marshal(params)
	return bytesData
}


func HttpPoxy(rw http.ResponseWriter, req *http.Request)(string, error){
	transport :=  http.DefaultTransport //超时时间暂时不单独设置


	// step 1
	outReq := new(http.Request)
	*outReq = *req // this only does shallow copies of maps

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	// step 2
	outReq.URL.Scheme ="http"
	outReq.URL.Host=req.Host
	outReq.RequestURI=""
	outReq.Header.Set("Accept-Encoding",  "") //这个如果是python request请求要设置，不然是乱码
	res, err := transport.RoundTrip(outReq)
	//	res, err :=http.DefaultClient.Do(outReq)

	if err != nil {
		fmt.Print(err)
		rw.WriteHeader(http.StatusBadGateway)
		return "",err
	}

	//step 3
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}

	//step 4
	var bodyBytes []byte
	if res.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(res.Body)
	}
	// 写回去
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	io.Copy(rw, res.Body)
	defer res.Body.Close()
	return string(bodyBytes),nil
}

