package svr02ress

type CAddPictureReply struct {
	UrlInfos    map[string][]string `json:"urlinfos"`
	Error       int                 `json:"error"`
	ErrorString string              `json:"errorstring"`
}

type CDeletePictureReply struct {
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}
