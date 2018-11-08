package svr01cfgs

type CGetServerInfoRequest struct {
	ServerType string `json:"servertype"`
}

type CGetServerInfoReply struct {
	ServerDomainname string `json:"serverdomainname"`
	ServerIp         string `json:"serverip"`
	ServerPort       int    `json:"serverport"`
	Error            int    `json:"error"`
	ErrorString      string `json:errorstring`
}

type CPostServerInfoRequest struct {
	ServerType       string `json:"servertype"`
	ServerName       string `json:"servername"`
	ServerIp         string `json:"serverip"`
	ServerPort       int    `json:"serverport"`
	ServerDomainname string `json:"serverdomainname"`
}

type CPostServerInfoReply struct {
	Serveruuid  string `json:"serveruuid"`
	Error       int    `json:"error"`
	ErrorString string `json:errorstring`
}

type CPutServerInfoRequest struct {
	ServerType       string `json:"servertype"`
	ServerName       string `json:"servername"`
	ServerIp         string `json:"serverip"`
	ServerPort       int    `json:"serverport"`
	ServerDomainname string `json:"serverdomainname"`
}

type CPutServerInfoReply struct {
	Error       int    `json:"error"`
	ErrorString string `json:errorstring`
}
