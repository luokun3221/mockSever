package res

type Res struct {
	RetCode int `json:"retCode"`
	Data[] interface{}  `json:"data"`
	Msg   string `json:"msg"`
	Allsize int `json:"allsize"`

}

type HttpInfo struct {
	Host string `json:"host"`
	RequestUrl string `json:"requestUrl"`
	ResponeBoby  string `json:"responeBoby"`
	Method string `json:"method"`
	RequestBoby string `json:"requestBoby"`

}


type QueryInfo struct {
	RemoteIp string `json:"remoteip"`
	PageSize int `json:"pagesize"`
	CurrentPage  int `json:"currentpage"`
}


type MockInfo struct {
	Id int `json:"id"`
	Ip string `json:"ip"`
	Host string `json:"host"`
	RequestUrl  string `json:"requesturl"`
	MockRespone  string `json:"mockrespone"`
	Status  string `json:"status"`
}

