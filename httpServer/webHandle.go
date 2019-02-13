package httpServer

import (
	"HttpMockServer/business"
	. "HttpMockServer/common/loginfo"
	"HttpMockServer/common/res"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ModifyMock(rw http.ResponseWriter, req *http.Request){
	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	body, _ := ioutil.ReadAll(req.Body)
	data := res.MockInfo{}
	err := json.Unmarshal(body, &data)
	if nil != err {
		Log.Errorf("LoadMockData json 失败;%v",err)
		return
	}

	result := business.ModifyMockInfo(data)
	resMsg, err := json.Marshal(result)
	if err != nil {

		Log.Errorf("LoadMockData Umarshal failed:%v", err)
		return
	}
	if resMsg != nil {
		Log.Infof("LoadMockData textMsg %s",string(resMsg))
		fmt.Fprintf(rw, string(resMsg))
	}
}


func DeleteMock(rw http.ResponseWriter, req *http.Request){
	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	body, _ := ioutil.ReadAll(req.Body)
	data := res.MockInfo{}
	err := json.Unmarshal(body, &data)
	if nil != err {
		Log.Errorf("AddMock json 失败;%v",err)
		return
	}

	result := business.DeleteMockInfo(data)
	resMsg, err := json.Marshal(result)
	if err != nil {

		Log.Errorf("AddMock Umarshal failed:%v", err)
		return
	}
	if resMsg != nil {
		Log.Infof("AddMock textMsg %s",string(resMsg))
		fmt.Fprintf(rw, string(resMsg))
	}
}


func LoadMockData(rw http.ResponseWriter, req *http.Request){
	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	body, _ := ioutil.ReadAll(req.Body)
	data := res.QueryInfo{}
	err := json.Unmarshal(body, &data)
	if nil != err {
		Log.Errorf("LoadMockData json 失败;%v",err)
		return
	}

	result := business.LoadHttpMockInfo(data)
	resMsg, err := json.Marshal(result)
	if err != nil {

		Log.Errorf("LoadMockData Umarshal failed:%v", err)
		return
	}
	if resMsg != nil {
		Log.Infof("LoadMockData textMsg %s",string(resMsg))
		fmt.Fprintf(rw, string(resMsg))
	}
}


func AddMock(rw http.ResponseWriter, req *http.Request){
	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	body, _ := ioutil.ReadAll(req.Body)
	data := res.MockInfo{}
	err := json.Unmarshal(body, &data)
	if nil != err {
		Log.Errorf("AddMock json 失败;%v",err)
		return
	}

	result := business.HttpMockAdd(data)
	resMsg, err := json.Marshal(result)
	if err != nil {

		Log.Errorf("AddMock Umarshal failed:%v", err)
		return
	}
	if resMsg != nil {
		Log.Infof("AddMock textMsg %s",string(resMsg))
		fmt.Fprintf(rw, string(resMsg))
	}
}




func ClearFilter(rw http.ResponseWriter, req *http.Request){
	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	body, _ := ioutil.ReadAll(req.Body)
	data := res.QueryInfo{}
	err := json.Unmarshal(body, &data)
	if nil != err {
		Log.Errorf("ClearFilter json 失败;%v",err)
		return
	}

	result := business.HttpInterfaceClear(data)
	resMsg, err := json.Marshal(result)
	if err != nil {

		Log.Errorf("ClearFilter Umarshal failed:%v", err)
		return
	}
	if resMsg != nil {
		Log.Infof("ClearFilter textMsg %s",string(resMsg))
		fmt.Fprintf(rw, string(resMsg))
	}

}


func ResetDataFilter(rw http.ResponseWriter, req *http.Request){
	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
//	rw.Header().Set("content-type", "application/json")             //返回数据格式是json
	body, _ := ioutil.ReadAll(req.Body)
	data := res.QueryInfo{}
	err := json.Unmarshal(body, &data)
	if nil != err {
		Log.Errorf("ResetDataFilter json 失败;%v",err)
		return
	}

	result := business.HttpInterfaceView(data)
	resMsg, err := json.Marshal(result)
	if err != nil {

		Log.Errorf("ResetDataFilter Umarshal failed:%v", err)
		return
	}

	if resMsg != nil {
		Log.Infof("ResetDataFilter textMsg %s",string(resMsg))
		fmt.Fprintf(rw, string(resMsg))
	}
}