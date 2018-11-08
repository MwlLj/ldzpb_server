package sessionmgr_mysql

type CProAddSessionInput struct {
	m_timeoutTime uint64
	m_losevaildTime uint64
	m_userdata string
}

func (this *CProAddSessionInput) SetTimeoutTime(timeoutTime uint64) {
	this.m_timeoutTime = timeoutTime
}
func (this *CProAddSessionInput) GetTimeoutTime() (uint64) {
	return this.m_timeoutTime
}
func (this *CProAddSessionInput) SetLosevaildTime(losevaildTime uint64) {
	this.m_losevaildTime = losevaildTime
}
func (this *CProAddSessionInput) GetLosevaildTime() (uint64) {
	return this.m_losevaildTime
}
func (this *CProAddSessionInput) SetUserdata(userdata string) {
	this.m_userdata = userdata
}
func (this *CProAddSessionInput) GetUserdata() (string) {
	return this.m_userdata
}

type CProAddSessionOutput struct {
	m_sessionId string
}

func (this *CProAddSessionOutput) SetSessionId(sessionId string) {
	this.m_sessionId = sessionId
}
func (this *CProAddSessionOutput) GetSessionId() (string) {
	return this.m_sessionId
}

type CProDeleteSessionInput struct {
	m_sessionId string
}

func (this *CProDeleteSessionInput) SetSessionId(sessionId string) {
	this.m_sessionId = sessionId
}
func (this *CProDeleteSessionInput) GetSessionId() (string) {
	return this.m_sessionId
}

type CProDeleteSessionOutput struct {
}

type CProSessionIsexistInput struct {
	m_sessionId string
}

func (this *CProSessionIsexistInput) SetSessionId(sessionId string) {
	this.m_sessionId = sessionId
}
func (this *CProSessionIsexistInput) GetSessionId() (string) {
	return this.m_sessionId
}

type CProSessionIsexistOutput struct {
	m_isexist bool
}

func (this *CProSessionIsexistOutput) SetIsexist(isexist bool) {
	this.m_isexist = isexist
}
func (this *CProSessionIsexistOutput) GetIsexist() (bool) {
	return this.m_isexist
}

type CProUpdateLosevaildtimeInput struct {
	m_sessionId string
	m_nowTimeStamp uint64
}

func (this *CProUpdateLosevaildtimeInput) SetSessionId(sessionId string) {
	this.m_sessionId = sessionId
}
func (this *CProUpdateLosevaildtimeInput) GetSessionId() (string) {
	return this.m_sessionId
}
func (this *CProUpdateLosevaildtimeInput) SetNowTimeStamp(nowTimeStamp uint64) {
	this.m_nowTimeStamp = nowTimeStamp
}
func (this *CProUpdateLosevaildtimeInput) GetNowTimeStamp() (uint64) {
	return this.m_nowTimeStamp
}

type CProUpdateLosevaildtimeOutput struct {
}

type CProGetSessioninfoBySessionidInput struct {
	m_sessionId string
}

func (this *CProGetSessioninfoBySessionidInput) SetSessionId(sessionId string) {
	this.m_sessionId = sessionId
}
func (this *CProGetSessioninfoBySessionidInput) GetSessionId() (string) {
	return this.m_sessionId
}

type CProGetSessioninfoBySessionidOutput struct {
	m_id uint64
	m_sessionId string
	m_timeoutTime uint64
	m_losevaildTime uint64
	m_userdata string
}

func (this *CProGetSessioninfoBySessionidOutput) SetId(id uint64) {
	this.m_id = id
}
func (this *CProGetSessioninfoBySessionidOutput) GetId() (uint64) {
	return this.m_id
}
func (this *CProGetSessioninfoBySessionidOutput) SetSessionId(sessionId string) {
	this.m_sessionId = sessionId
}
func (this *CProGetSessioninfoBySessionidOutput) GetSessionId() (string) {
	return this.m_sessionId
}
func (this *CProGetSessioninfoBySessionidOutput) SetTimeoutTime(timeoutTime uint64) {
	this.m_timeoutTime = timeoutTime
}
func (this *CProGetSessioninfoBySessionidOutput) GetTimeoutTime() (uint64) {
	return this.m_timeoutTime
}
func (this *CProGetSessioninfoBySessionidOutput) SetLosevaildTime(losevaildTime uint64) {
	this.m_losevaildTime = losevaildTime
}
func (this *CProGetSessioninfoBySessionidOutput) GetLosevaildTime() (uint64) {
	return this.m_losevaildTime
}
func (this *CProGetSessioninfoBySessionidOutput) SetUserdata(userdata string) {
	this.m_userdata = userdata
}
func (this *CProGetSessioninfoBySessionidOutput) GetUserdata() (string) {
	return this.m_userdata
}

type CProDeleteLosevaildSessionsInput struct {
	m_nowTimeStamp uint64
}

func (this *CProDeleteLosevaildSessionsInput) SetNowTimeStamp(nowTimeStamp uint64) {
	this.m_nowTimeStamp = nowTimeStamp
}
func (this *CProDeleteLosevaildSessionsInput) GetNowTimeStamp() (uint64) {
	return this.m_nowTimeStamp
}

type CProDeleteLosevaildSessionsOutput struct {
}

