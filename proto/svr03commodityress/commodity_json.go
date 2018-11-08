package svr03commodityress

type CPostCommodityClassifitionRequest struct {
}

type CPostCommodityClassifitionReply struct {
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}
