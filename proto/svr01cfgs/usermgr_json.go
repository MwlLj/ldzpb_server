package svr01cfgs

type CPostUserRequest struct {
	Username string `json:"username"`
	Userpwd  string `json:"userpwd"`
	EmailNo  string `json:"emailno"`
	PhoneNo  string `json:"phoneno"`
}

type CPostUserReply struct {
	Useruuid    string `json:"useruuid"`
	Error       int    `json:"error"`
	ErrorString string `json:errorstring`
}

type CDeleteUserRequest struct {
	Useruuid string `json:"useruuid"`
}

type CDeleteUserReply struct {
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}

type CPostUserLoginRequest struct {
	Userno  string `json:"userno"`
	Userpwd string `json:"userpwd"`
}

type CPostUserLoginReply struct {
	Useruuid    string `json:"useruuid"`
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}
