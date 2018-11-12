package svr03commodityress

type CPostCommodityClassifitionRequest struct {
	Name       string `json:"name"`
	ParentUuid string `json:"parentUuid"`
}

type CPostCommodityClassifitionReply struct {
	Uuid        string `json:"uuid"`
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CPutCommodityClassifitionRequest struct {
	Uuid         string  `json:"uuid"`
	Name         *string `json:"name"`
	ClassifyUuid *string `json:"classifyuuid"`
}

type CPutCommodityClassifitionReply struct {
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CDeleteCommodityClassifitionRequest struct {
	Uuid string `json:"uuid"`
}

type CDeleteCommodityClassifitionReply struct {
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CGetCommodityClassifitionRequest struct {
	Uuid string `json:"uuid"`
}

type CGetCommodityClassifitionReply struct {
	Name        string `json:"name"`
	ParentUuid  string `json:"parentuuid"`
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CPostCommodityClassifitionDetailRequest struct {
	ClassifyUuid string `json:"classifyuuid"`
	DetailType   string `json:"detailtype"`
	DetailValue  string `json:"detailvalue"`
	DetailNo     int    `json:"detailno"`
}

type CPostCommodityClassifitionDetailReply struct {
	DetailUuid  string `json:"detialuuid"`
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CPutCommodityClassifitionDetailRequest struct {
	DetialUuid  string  `json:"detailuuid"`
	DetailType  *string `json:"detailtype"`
	DetailValue *string `json:"detailvalue"`
	DetailNo    *int    `json:"detailno"`
}

type CPutCommodityClassifitionDetailReply struct {
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CDeleteCommodityClassifitionDetailRequest struct {
	DetailUuid string `json:"detialuuid"`
}

type CDeleteCommodityClassifitionDetailReply struct {
	Error       int    `json:"errorcode"`
	ErrorString string `json:"errorstring"`
}

type CGetCommodityClassifitionDetailRequest struct {
	DetailUuid string `json:"detailuuid"`
}

type CGetCommodityClassifitionDetailReply struct {
	ClassifyUuid string `json:"classifyuuid"`
	DetailType   string `json:"detailtype"`
	DetailValue  string `json:"detailvalue"`
	DetailNo     int    `json:"detailno"`
	Error        int    `json:"errorcode"`
	ErrorString  string `json:"errorstring"`
}
