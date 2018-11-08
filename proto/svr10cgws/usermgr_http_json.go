package svr10cgws

type CRegisterUserRequest struct {
	Username string `json:"username"`
	Userpwd  string `json:"userpwd"`
	EmailNo  string `json:"emailno"`
	PhoneNo  string `json:"phoneno"`
}

type CRegisterUserReply struct {
	Useruuid    string `json:"useruuid"`
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}

type CUnRegisterUserRequest struct {
	Cookieid string `json:"cookieid"`
	Useruuid string `json:"useruuid"`
}

type CUnRegisterUserReply struct {
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}

type CUserLoginRequest struct {
	Userno   string `json:"userno"`
	Userpwd  string `json:"userpwd"`
	Cookieid string `json:"cookieid"`
}

type CUserLoginReply struct {
	Sessionid   string `json:"sessionid"`
	Useruuid    string `json:"useruuid"`
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}

type CLogoutRequest struct {
	Cookieid string `json:"cookieid"`
}

type CLogoutReply struct {
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}
