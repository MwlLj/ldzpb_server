package usermgr

type CProUserIsexistsByEmailInput struct {
	m_userEmail string
}

func (this *CProUserIsexistsByEmailInput) SetUserEmail(userEmail string) {
	this.m_userEmail = userEmail
}
func (this *CProUserIsexistsByEmailInput) GetUserEmail() (string) {
	return this.m_userEmail
}

type CProUserIsexistsByEmailOutput struct {
	m_userUuid string
}

func (this *CProUserIsexistsByEmailOutput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProUserIsexistsByEmailOutput) GetUserUuid() (string) {
	return this.m_userUuid
}

type CProUserIsexistsByPhoneInput struct {
	m_userPhone string
}

func (this *CProUserIsexistsByPhoneInput) SetUserPhone(userPhone string) {
	this.m_userPhone = userPhone
}
func (this *CProUserIsexistsByPhoneInput) GetUserPhone() (string) {
	return this.m_userPhone
}

type CProUserIsexistsByPhoneOutput struct {
	m_userUuid string
}

func (this *CProUserIsexistsByPhoneOutput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProUserIsexistsByPhoneOutput) GetUserUuid() (string) {
	return this.m_userUuid
}

type CProUserIsexistsByEmailOrPhoneInput struct {
	m_userNo string
}

func (this *CProUserIsexistsByEmailOrPhoneInput) SetUserNo(userNo string) {
	this.m_userNo = userNo
}
func (this *CProUserIsexistsByEmailOrPhoneInput) GetUserNo() (string) {
	return this.m_userNo
}

type CProUserIsexistsByEmailOrPhoneOutput struct {
	m_userUuid string
}

func (this *CProUserIsexistsByEmailOrPhoneOutput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProUserIsexistsByEmailOrPhoneOutput) GetUserUuid() (string) {
	return this.m_userUuid
}

type CProAddUserInput struct {
	m_userName string
	m_userPwd string
	m_userEmailNo string
	m_userPhoneNo string
}

func (this *CProAddUserInput) SetUserName(userName string) {
	this.m_userName = userName
}
func (this *CProAddUserInput) GetUserName() (string) {
	return this.m_userName
}
func (this *CProAddUserInput) SetUserPwd(userPwd string) {
	this.m_userPwd = userPwd
}
func (this *CProAddUserInput) GetUserPwd() (string) {
	return this.m_userPwd
}
func (this *CProAddUserInput) SetUserEmailNo(userEmailNo string) {
	this.m_userEmailNo = userEmailNo
}
func (this *CProAddUserInput) GetUserEmailNo() (string) {
	return this.m_userEmailNo
}
func (this *CProAddUserInput) SetUserPhoneNo(userPhoneNo string) {
	this.m_userPhoneNo = userPhoneNo
}
func (this *CProAddUserInput) GetUserPhoneNo() (string) {
	return this.m_userPhoneNo
}

type CProAddUserOutput struct {
	m_userUuid string
}

func (this *CProAddUserOutput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProAddUserOutput) GetUserUuid() (string) {
	return this.m_userUuid
}

type CProPasswordIstrueByUseruuidInput struct {
	m_userUuid string
	m_userPwd string
}

func (this *CProPasswordIstrueByUseruuidInput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProPasswordIstrueByUseruuidInput) GetUserUuid() (string) {
	return this.m_userUuid
}
func (this *CProPasswordIstrueByUseruuidInput) SetUserPwd(userPwd string) {
	this.m_userPwd = userPwd
}
func (this *CProPasswordIstrueByUseruuidInput) GetUserPwd() (string) {
	return this.m_userPwd
}

type CProPasswordIstrueByUseruuidOutput struct {
	m_isTrue bool
}

func (this *CProPasswordIstrueByUseruuidOutput) SetIsTrue(isTrue bool) {
	this.m_isTrue = isTrue
}
func (this *CProPasswordIstrueByUseruuidOutput) GetIsTrue() (bool) {
	return this.m_isTrue
}

type CProPasswordIstrueByNoInput struct {
	m_no string
	m_userPwd string
}

func (this *CProPasswordIstrueByNoInput) SetNo(no string) {
	this.m_no = no
}
func (this *CProPasswordIstrueByNoInput) GetNo() (string) {
	return this.m_no
}
func (this *CProPasswordIstrueByNoInput) SetUserPwd(userPwd string) {
	this.m_userPwd = userPwd
}
func (this *CProPasswordIstrueByNoInput) GetUserPwd() (string) {
	return this.m_userPwd
}

type CProPasswordIstrueByNoOutput struct {
	m_isTrue bool
	m_userUuid string
}

func (this *CProPasswordIstrueByNoOutput) SetIsTrue(isTrue bool) {
	this.m_isTrue = isTrue
}
func (this *CProPasswordIstrueByNoOutput) GetIsTrue() (bool) {
	return this.m_isTrue
}
func (this *CProPasswordIstrueByNoOutput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProPasswordIstrueByNoOutput) GetUserUuid() (string) {
	return this.m_userUuid
}

type CProDeleteUserInput struct {
	m_userUuid string
}

func (this *CProDeleteUserInput) SetUserUuid(userUuid string) {
	this.m_userUuid = userUuid
}
func (this *CProDeleteUserInput) GetUserUuid() (string) {
	return this.m_userUuid
}

type CProDeleteUserOutput struct {
}

