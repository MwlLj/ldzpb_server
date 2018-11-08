package sessionmgr

import (
	"../configs"
	db "./sessionmgr_mysql"
	_ "github.com/go-sql-driver/mysql"
)

type sessionMgrMysql struct {
	m_db db.CMysqlHandler
}

func NewSessionMgrMysql() *sessionMgrMysql {
	return &sessionMgrMysql{}
}

func (this *sessionMgrMysql) connect() {
	dbcfg, err := configs.NewMysqlDbConfig("session_msyql_db.cfg")
	if err != nil {
		panic("[Error] open db config error")
	}
	dbData := dbcfg.GetMysqlDbData()
	err = this.m_db.Connect(dbData.Host, dbData.Port, dbData.Username, dbData.Userpwd, dbData.Dbname)
	if err != nil {
		panic("[Error] connect db config error")
	}
}

func (this *sessionMgrMysql) disConnect() {
	this.m_db.Disconnect()
}

func (this *sessionMgrMysql) ConnectDb() {
	this.connect()
}

func (this *sessionMgrMysql) DisConnectDb() {
	this.disConnect()
}

func (this *sessionMgrMysql) CreateSession(timeout uint64, userdata string) (sessionId string, err error) {
	losevaildTime := timeout + getNowTimeStamp()
	// call mysql procedure
	input := db.CProAddSessionInput{}
	output := db.CProAddSessionOutput{}
	input.SetTimeoutTime(timeout)
	input.SetLosevaildTime(losevaildTime)
	input.SetUserdata(userdata)
	e := this.m_db.ProAddSession(&input, &output)
	if e != nil {
		return "", e
	}
	return output.GetSessionId(), nil
}

func (this *sessionMgrMysql) DestroySession(sessionId string) error {
	input := db.CProDeleteSessionInput{}
	input.SetSessionId(sessionId)
	e := this.m_db.ProDeleteSession(&input)
	if e != nil {
		return e
	}
	return nil
}

func (this *sessionMgrMysql) SessionIsVaild(sessionId string) (isVaild bool, err error) {
	input := db.CProSessionIsexistInput{}
	output := db.CProSessionIsexistOutput{}
	input.SetSessionId(sessionId)
	e := this.m_db.ProSessionIsexist(&input, &output)
	if e != nil {
		return false, e
	}
	isexist := output.GetIsexist()
	return isexist, nil
}

func (this *sessionMgrMysql) ResetLosevaildTime(sessionId string) error {
	input := db.CProUpdateLosevaildtimeInput{}
	input.SetSessionId(sessionId)
	input.SetNowTimeStamp(getNowTimeStamp())
	e := this.m_db.ProUpdateLosevaildtime(&input)
	if e != nil {
		return e
	}
	return nil
}

func (this *sessionMgrMysql) DeleteSessionAfterLosevaild(nowTimeStamp uint64) error {
	input := db.CProDeleteLosevaildSessionsInput{}
	input.SetNowTimeStamp(nowTimeStamp)
	e := this.m_db.ProDeleteLosevaildSessions(&input)
	if e != nil {
		return e
	}
	return nil
}

func (this *sessionMgrMysql) GetUserdata(sessionId string) (string, error) {
	input := db.CProGetSessioninfoBySessionidInput{}
	output := db.CProGetSessioninfoBySessionidOutput{}
	input.SetSessionId(sessionId)
	e := this.m_db.ProGetSessioninfoBySessionid(&input, &output)
	if e != nil {
		return "", e
	}
	return output.GetUserdata(), nil
}
