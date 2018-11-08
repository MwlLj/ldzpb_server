package svr02ress

import (
	"../../tools/timetool"
	"strings"
)

var (
	CommodityPictureFormname     string = "commoditypicurl"
	CommodityDescPictureFormname string = "commoditydespiccurl"
	CommodityRoot                string = strings.Join([]string{"commodity", timetool.GetNowDayFormat()}, "/")
)

// add commodity picture [POST]
var AddCommodityPicture string = "/commodity/picture"

// delete commodity picture [DELETE]
var DeleteCommodityPicture string = "/commodity/picture"
