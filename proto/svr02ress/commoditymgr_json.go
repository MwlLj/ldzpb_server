package svr02ress

type CAddCommodityPictureReply struct {
	UrlInfos    map[string][]string `json:"urlinfos"`
	Error       int                 `json:"error"`
	ErrorString string              `json:"errorstring"`
}
