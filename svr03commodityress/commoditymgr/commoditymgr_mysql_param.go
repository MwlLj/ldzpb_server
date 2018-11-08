package commoditymgr

type CProAddCommodityClassifitionInput struct {
	m_parentUuid string
	m_name string
}

func (this *CProAddCommodityClassifitionInput) SetParentUuid(parentUuid string) {
	this.m_parentUuid = parentUuid
}
func (this *CProAddCommodityClassifitionInput) GetParentUuid() (string) {
	return this.m_parentUuid
}
func (this *CProAddCommodityClassifitionInput) SetName(name string) {
	this.m_name = name
}
func (this *CProAddCommodityClassifitionInput) GetName() (string) {
	return this.m_name
}

type CProAddCommodityClassifitionOutput struct {
	m_uuid string
}

func (this *CProAddCommodityClassifitionOutput) SetUuid(uuid string) {
	this.m_uuid = uuid
}
func (this *CProAddCommodityClassifitionOutput) GetUuid() (string) {
	return this.m_uuid
}

type CProAddCommodityClassifitionDetailInput struct {
	m_classifitionUuid string
	m_detailType string
	m_detailValue string
	m_detailNo int
}

func (this *CProAddCommodityClassifitionDetailInput) SetClassifitionUuid(classifitionUuid string) {
	this.m_classifitionUuid = classifitionUuid
}
func (this *CProAddCommodityClassifitionDetailInput) GetClassifitionUuid() (string) {
	return this.m_classifitionUuid
}
func (this *CProAddCommodityClassifitionDetailInput) SetDetailType(detailType string) {
	this.m_detailType = detailType
}
func (this *CProAddCommodityClassifitionDetailInput) GetDetailType() (string) {
	return this.m_detailType
}
func (this *CProAddCommodityClassifitionDetailInput) SetDetailValue(detailValue string) {
	this.m_detailValue = detailValue
}
func (this *CProAddCommodityClassifitionDetailInput) GetDetailValue() (string) {
	return this.m_detailValue
}
func (this *CProAddCommodityClassifitionDetailInput) SetDetailNo(detailNo int) {
	this.m_detailNo = detailNo
}
func (this *CProAddCommodityClassifitionDetailInput) GetDetailNo() (int) {
	return this.m_detailNo
}

type CProAddCommodityClassifitionDetailOutput struct {
}

type CProAddCommodityInput struct {
	m_parentUuid string
	m_name string
}

func (this *CProAddCommodityInput) SetParentUuid(parentUuid string) {
	this.m_parentUuid = parentUuid
}
func (this *CProAddCommodityInput) GetParentUuid() (string) {
	return this.m_parentUuid
}
func (this *CProAddCommodityInput) SetName(name string) {
	this.m_name = name
}
func (this *CProAddCommodityInput) GetName() (string) {
	return this.m_name
}

type CProAddCommodityOutput struct {
	m_uuid string
}

func (this *CProAddCommodityOutput) SetUuid(uuid string) {
	this.m_uuid = uuid
}
func (this *CProAddCommodityOutput) GetUuid() (string) {
	return this.m_uuid
}

type CProAddCommodityDetailInput struct {
	m_commodityUuid string
	m_detailType string
	m_detailValue string
	m_detailNo int
}

func (this *CProAddCommodityDetailInput) SetCommodityUuid(commodityUuid string) {
	this.m_commodityUuid = commodityUuid
}
func (this *CProAddCommodityDetailInput) GetCommodityUuid() (string) {
	return this.m_commodityUuid
}
func (this *CProAddCommodityDetailInput) SetDetailType(detailType string) {
	this.m_detailType = detailType
}
func (this *CProAddCommodityDetailInput) GetDetailType() (string) {
	return this.m_detailType
}
func (this *CProAddCommodityDetailInput) SetDetailValue(detailValue string) {
	this.m_detailValue = detailValue
}
func (this *CProAddCommodityDetailInput) GetDetailValue() (string) {
	return this.m_detailValue
}
func (this *CProAddCommodityDetailInput) SetDetailNo(detailNo int) {
	this.m_detailNo = detailNo
}
func (this *CProAddCommodityDetailInput) GetDetailNo() (int) {
	return this.m_detailNo
}

type CProAddCommodityDetailOutput struct {
}

