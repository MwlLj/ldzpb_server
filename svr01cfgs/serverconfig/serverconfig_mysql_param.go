package serverconfig

type CProAddServerInfoInput struct {
	m_servername string
	m_servertype string
	m_serverip string
	m_serverport int
	m_serverdomainname string
}

func (this *CProAddServerInfoInput) SetServername(servername string) {
	this.m_servername = servername
}
func (this *CProAddServerInfoInput) GetServername() (string) {
	return this.m_servername
}
func (this *CProAddServerInfoInput) SetServertype(servertype string) {
	this.m_servertype = servertype
}
func (this *CProAddServerInfoInput) GetServertype() (string) {
	return this.m_servertype
}
func (this *CProAddServerInfoInput) SetServerip(serverip string) {
	this.m_serverip = serverip
}
func (this *CProAddServerInfoInput) GetServerip() (string) {
	return this.m_serverip
}
func (this *CProAddServerInfoInput) SetServerport(serverport int) {
	this.m_serverport = serverport
}
func (this *CProAddServerInfoInput) GetServerport() (int) {
	return this.m_serverport
}
func (this *CProAddServerInfoInput) SetServerdomainname(serverdomainname string) {
	this.m_serverdomainname = serverdomainname
}
func (this *CProAddServerInfoInput) GetServerdomainname() (string) {
	return this.m_serverdomainname
}

type CProAddServerInfoOutput struct {
	m_serveruuid string
}

func (this *CProAddServerInfoOutput) SetServeruuid(serveruuid string) {
	this.m_serveruuid = serveruuid
}
func (this *CProAddServerInfoOutput) GetServeruuid() (string) {
	return this.m_serveruuid
}

type CProGetServerInfoInput struct {
	m_servertype string
}

func (this *CProGetServerInfoInput) SetServertype(servertype string) {
	this.m_servertype = servertype
}
func (this *CProGetServerInfoInput) GetServertype() (string) {
	return this.m_servertype
}

type CProGetServerInfoOutput struct {
	m_serveruuid string
	m_servername string
	m_servertype string
	m_serverip string
	m_serverport int
	m_serverdomainname string
}

func (this *CProGetServerInfoOutput) SetServeruuid(serveruuid string) {
	this.m_serveruuid = serveruuid
}
func (this *CProGetServerInfoOutput) GetServeruuid() (string) {
	return this.m_serveruuid
}
func (this *CProGetServerInfoOutput) SetServername(servername string) {
	this.m_servername = servername
}
func (this *CProGetServerInfoOutput) GetServername() (string) {
	return this.m_servername
}
func (this *CProGetServerInfoOutput) SetServertype(servertype string) {
	this.m_servertype = servertype
}
func (this *CProGetServerInfoOutput) GetServertype() (string) {
	return this.m_servertype
}
func (this *CProGetServerInfoOutput) SetServerip(serverip string) {
	this.m_serverip = serverip
}
func (this *CProGetServerInfoOutput) GetServerip() (string) {
	return this.m_serverip
}
func (this *CProGetServerInfoOutput) SetServerport(serverport int) {
	this.m_serverport = serverport
}
func (this *CProGetServerInfoOutput) GetServerport() (int) {
	return this.m_serverport
}
func (this *CProGetServerInfoOutput) SetServerdomainname(serverdomainname string) {
	this.m_serverdomainname = serverdomainname
}
func (this *CProGetServerInfoOutput) GetServerdomainname() (string) {
	return this.m_serverdomainname
}

