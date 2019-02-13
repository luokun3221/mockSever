package business

import (
	. "HttpMockServer/common/loginfo"
	"HttpMockServer/common/res"
	"HttpMockServer/dbOpr"
)

func HttpInterfaceView(data res.QueryInfo)(res.Res){
	result := res.Res{}
	var allSize int
	errCount :=dbOpr.Db.QueryRow("SELECT count(*) From httpRequest where remoteIp = ?",data.RemoteIp).Scan(&allSize)
	if errCount !=nil{
		Log.Error(errCount)
		result.RetCode = -1
		result.Allsize = 0
		return  result
	}



	startpage := (data.CurrentPage-1)*data.PageSize

	rows, err :=dbOpr.Db.Query("SELECT uuid,method,host,requestUrl,requestBoby From httpRequest where remoteIp = ? order by id desc limit ?,?",data.RemoteIp,startpage,data.PageSize)
	defer rows.Close()
	if err != nil {
		Log.Error(err)
		result.RetCode = -1
		return result
	}
	for rows.Next() {
		var uuid string
		var method string
		var host string
		var requestUrl string
		var requestBoby string
		var responeBoby string
		rows.Scan(&uuid, &method, &host,&requestUrl,&requestBoby)

		errRes :=dbOpr.Db.QueryRow("SELECT responeBoby  From httpRespone where uuid =?",uuid).Scan(&responeBoby)
		if errRes != nil {
			Log.Error(errRes)
			result.RetCode = -1
			continue
		}

		httpInfo := res.HttpInfo{host,requestUrl,responeBoby,method,requestBoby}
		result.Data = append(result.Data, httpInfo)
	}
	result.RetCode = 0
	result.Allsize=allSize

	return  result
}

func HttpInterfaceClear(data res.QueryInfo)(res.Res){
	result := res.Res{}
	dbOpr.Db.Exec("delete from httpRespone where uuid in(select uuid from httpRequest where remoteIp=?)",data.RemoteIp)
	dbOpr.Db.Exec("delete from httpRequest where remoteIp=?)",data.RemoteIp)
	result.RetCode=0
	result.Allsize=0
	return result



}