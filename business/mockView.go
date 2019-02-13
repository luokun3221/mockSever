package business

import (
	"HttpMockServer/common/res"
	"HttpMockServer/dbOpr"
	. "HttpMockServer/common/loginfo"

)

func ModifyMockInfo(data res.MockInfo)(res.Res){
	result := res.Res{}
	_,err := dbOpr.Db.Exec("update MockInfo set status=? where id=?",data.Status,data.Id)
	if err !=nil{
		Log.Error(err)
		result.RetCode = -1
		return  result
	}
	result.RetCode=0
	return result
}

func DeleteMockInfo(data res.MockInfo)(res.Res){
	result := res.Res{}
	_,err := dbOpr.Db.Exec("delete from MockInfo where id =?",data.Id)
	if err !=nil{
		Log.Error(err)
		result.RetCode = -1
		return  result
	}
	result.RetCode=0
	return result
}
func HttpMockAdd(data res.MockInfo)(res.Res){
	result := res.Res{}
	var isHaveUrl int
	err :=dbOpr.Db.QueryRow("SELECT count(*) From MockInfo where requestUrl = ? and host= ? and ip= ?",data.RequestUrl,data.Host,data.Ip).Scan(&isHaveUrl)
	if err !=nil{
		Log.Error(err)
		result.RetCode = -1
		return  result
	}

	if isHaveUrl > 0{
		result.RetCode = -1
		result.Msg ="已经存在requsetUrl"
		return  result
	}

	dbOpr.Db.Exec("insert into MockInfo(ip,host,requestUrl,mockResp,status) values(?,?,?,?,?)",
		data.Ip,data.Host,data.RequestUrl,data.MockRespone,"1")
	result.RetCode=0
    return result
}


func LoadHttpMockInfo(data res.QueryInfo)(res.Res){
	result := res.Res{}
	var allSize int
	err :=dbOpr.Db.QueryRow("SELECT count(*) From MockInfo where ip= ?",data.RemoteIp).Scan(&allSize)
	if err !=nil{
		Log.Error(err)
		result.RetCode = -1
		result.Allsize = 0
		return  result
	}

	startpage := (data.CurrentPage-1)*data.PageSize

	rows, err :=dbOpr.Db.Query("SELECT id,host,requestUrl,mockResp,status From MockInfo where ip = ? order by id desc limit ?,?",data.RemoteIp,startpage,data.PageSize)
	defer rows.Close()
	if err != nil {
		Log.Error(err)
		result.RetCode = -1
		return result
	}
	for rows.Next() {
		var id int
		var host string
		var requestUrl string
		var mockResp string
		var status string
		rows.Scan(&id, &host,&requestUrl,&mockResp,&status)

		mockInfo := res.MockInfo{id,data.RemoteIp,host,requestUrl,mockResp,status}
		result.Data = append(result.Data, mockInfo)
	}
	result.RetCode = 0
	result.Allsize=allSize

	return  result

}
