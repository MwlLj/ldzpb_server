package svr10cgws

type CAddServerInfoRequest struct {
	Cookieid         string `json:"cookieid"`
	Servertype       string `json:"servertype"`
	Servername       string `json:"servername"`
	Serverip         string `json:"serverip"`
	Serverport       int    `json:"serverport"`
	Serverdomainname string `json:"serverdomainname"`
}

type CAddServerInfoReply struct {
	Serveruuid  string `json:"serveruuid"`
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}

var (
	GetServerInfoHeaderCookieid   string = "cookieid"
	GetServerInfoHeaderServertype string = "servertype"
)

type CGetServerInfoReply struct {
	Serverdomainname string `json:"serverdomainname"`
	Serverip         string `json:"serverip"`
	Serverport       int    `json:"serverport"`
	Error            int    `json:"error"`
	ErrorString      string `json:errorstring`
}
