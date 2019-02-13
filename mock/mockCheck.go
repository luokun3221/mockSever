package mock

import (
	"HttpMockServer/dbOpr"
	"strings"
)

func IsNeedMock(host,requestUrl,remoteIp,method string)(string){
	var MockResp string =""
	if "GET" == method {
		//全匹配
		dbOpr.Db.QueryRow("SELECT mockResp From MockInfo where requestUrl = ? and host= ? and ip= ?",requestUrl,host,remoteIp).Scan(&MockResp)
		if ""==MockResp {
			baseUrl := strings.Split(requestUrl, "?")
			dbOpr.Db.QueryRow("SELECT mockResp From MockInfo where requestUrl = ? and host= ? and ip= ?",baseUrl[0],host,remoteIp).Scan(&MockResp)
		}
		return MockResp

	}else {
		dbOpr.Db.QueryRow("SELECT mockResp From MockInfo where requestUrl like  ? and host= ? and ip= ?","%" + requestUrl + "%",host,remoteIp).Scan(&MockResp)
        return  MockResp
	}
}
