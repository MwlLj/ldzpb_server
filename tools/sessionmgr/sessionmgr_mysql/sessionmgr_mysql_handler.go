package sessionmgr_mysql

import (
	"bufio"
	"bytes"
	"database/sql"
	"io"
	"os"
	"regexp"
	"strconv"
)

type CMysqlHandler struct  {
	m_db *sql.DB
}

func (this *CMysqlHandler) Connect(host string, port uint, username string, userpwd string, dbname string) (err error) {
	b := bytes.Buffer{}
	b.WriteString(username)
	b.WriteString(":")
	b.WriteString(userpwd)
	b.WriteString("@tcp(")
	b.WriteString(host)
	b.WriteString(":")
	b.WriteString(strconv.FormatUint(uint64(port), 10))
	b.WriteString(")/")
	b.WriteString(dbname)
	this.m_db, err = sql.Open("mysql", b.String())
	if err != nil {
		return err
	}
	this.m_db.SetMaxOpenConns(2000)
	this.m_db.SetMaxIdleConns(1000)
	this.m_db.Ping()
	return nil
}

func (this *CMysqlHandler) ConnectByCfg(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var host string = "localhost"
	var port uint = 3306
	var username string = "root"
	var userpwd string = "123456"
	var dbname string = "test"
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		content := string(a)
		r, _ := regexp.Compile("(.*)?=(.*)?")
		ret := r.FindStringSubmatch(content)
		if len(ret) != 3 {
			continue
		}
		k := ret[1]
		v := ret[2]
		switch k {
		case "host":
			host = v
		case "port":
			port_tmp, _ := strconv.ParseUint(v, 10, 32)
			port = uint(port_tmp)
		case "username":
			username = v
		case "userpwd":
			userpwd = v
		case "dbname":
			dbname = v
		}
	}
	return this.Connect(host, port, username, userpwd, dbname)
}

func (this *CMysqlHandler) Disconnect() {
	this.m_db.Close()
}

func (this *CMysqlHandler) ProAddSession(input *CProAddSessionInput, output *CProAddSessionOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_add_session(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(input.GetTimeoutTime(), input.GetLosevaildTime(), input.GetUserdata())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var sessionId string
		scanErr := rows.Scan(&sessionId)
		if scanErr != nil {
			continue
		}
		output.SetSessionId(sessionId)
	}
	return nil
}

func (this *CMysqlHandler) ProDeleteSession(input *CProDeleteSessionInput) (error) {
	stmt, err := this.m_db.Prepare("call pro_delete_session(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(input.GetSessionId())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
	}
	return nil
}

func (this *CMysqlHandler) ProSessionIsexist(input *CProSessionIsexistInput, output *CProSessionIsexistOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_session_isexist(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(input.GetSessionId())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var isexist bool
		scanErr := rows.Scan(&isexist)
		if scanErr != nil {
			continue
		}
		output.SetIsexist(isexist)
	}
	return nil
}

func (this *CMysqlHandler) ProUpdateLosevaildtime(input *CProUpdateLosevaildtimeInput) (error) {
	stmt, err := this.m_db.Prepare("call pro_update_losevaildtime(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(input.GetSessionId(), input.GetNowTimeStamp())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
	}
	return nil
}

func (this *CMysqlHandler) ProGetSessioninfoBySessionid(input *CProGetSessioninfoBySessionidInput, output *CProGetSessioninfoBySessionidOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_get_sessioninfo_by_sessionid(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(input.GetSessionId())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id uint64
		var sessionId string
		var timeoutTime uint64
		var losevaildTime uint64
		var userdata string
		scanErr := rows.Scan(&id, &sessionId, &timeoutTime, &losevaildTime, &userdata)
		if scanErr != nil {
			continue
		}
		output.SetId(id)
		output.SetSessionId(sessionId)
		output.SetTimeoutTime(timeoutTime)
		output.SetLosevaildTime(losevaildTime)
		output.SetUserdata(userdata)
	}
	return nil
}

func (this *CMysqlHandler) ProDeleteLosevaildSessions(input *CProDeleteLosevaildSessionsInput) (error) {
	stmt, err := this.m_db.Prepare("call pro_delete_losevaild_sessions(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(input.GetNowTimeStamp())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
	}
	return nil
}

