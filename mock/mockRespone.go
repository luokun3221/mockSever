//只是mock返回，后续可以mock请求
package mock

import "encoding/json"

//随便写的，后续通过前端配置，数据库读取
func addParams()([]byte){
	params := make(map[string]interface{})
	params["aid"] = 1
	params["page"] = 0
	params["status"] = "ongoing"
	params["size"] = 20

	bytesData, _:= json.Marshal(params)
	return bytesData
}

func MockRespone()([]byte){
	bytesData :=addParams()
	return bytesData

}


