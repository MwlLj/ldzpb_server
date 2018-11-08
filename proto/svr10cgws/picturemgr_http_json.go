package svr10cgws

var (
	AddPictureResourceHeaderCookieid string = "cookieid"
	AddPictureResourceHeaderPictype  string = "pictype"
	AddPictureResourceHeaderFormdata string = "pictureurl"
)

type CAddPictureReply struct {
	UrlInfos    map[string][]string `json:"urlinfos"`
	Error       int                 `json:"error"`
	ErrorString string              `json:"errorstring"`
}

var (
	DeletePictureResourceHeaderCookid     string = "cookieid"
	DeletePictureResourceHeaderPictureurl string = "picurl"
)

type CDeletePictureReply struct {
	Error       int    `json:"error"`
	ErrorString string `json:"errorstring"`
}
