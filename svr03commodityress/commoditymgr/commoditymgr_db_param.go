package commoditymgr

type CAddCommodityClassifitionInput struct {
	Uuid string
	ParentUuid string
	Name string
}

type CAddCommodityClassifitionDetailInfoInput struct {
	DetailUuid string
	DetailType string
	DetailValue string
	DetailNo int
	ClassifyUuid string
}

type CAddCommodityInfoInput struct {
	Uuid string
	Name string
	Price string
	ClassifyUuid string
}

type CAddCommodityDetailInfoInput struct {
	Uuid string
	DetailType string
	DetailValue string
	DetialNo int
	CommodityUuid string
}

