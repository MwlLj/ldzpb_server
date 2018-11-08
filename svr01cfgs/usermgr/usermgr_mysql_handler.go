package usermgr

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

func (this *CMysqlHandler) ProUserIsexistsByEmail(input *CProUserIsexistsByEmailInput, output *CProUserIsexistsByEmailOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_user_isexists_by_email(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetUserEmail())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userUuid string
		scanErr := rows.Scan(&userUuid)
		if scanErr != nil {
			continue
		}
		output.SetUserUuid(userUuid)
	}
	return nil
}

func (this *CMysqlHandler) ProUserIsexistsByPhone(input *CProUserIsexistsByPhoneInput, output *CProUserIsexistsByPhoneOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_user_isexists_by_phone(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetUserPhone())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userUuid string
		scanErr := rows.Scan(&userUuid)
		if scanErr != nil {
			continue
		}
		output.SetUserUuid(userUuid)
	}
	return nil
}

func (this *CMysqlHandler) ProUserIsexistsByEmailOrPhone(input *CProUserIsexistsByEmailOrPhoneInput, output *CProUserIsexistsByEmailOrPhoneOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_user_isexists_by_email_or_phone(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetUserNo())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userUuid string
		scanErr := rows.Scan(&userUuid)
		if scanErr != nil {
			continue
		}
		output.SetUserUuid(userUuid)
	}
	return nil
}

func (this *CMysqlHandler) ProAddUser(input *CProAddUserInput, output *CProAddUserOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_add_user(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetUserName(), input.GetUserPwd(), input.GetUserEmailNo(), input.GetUserPhoneNo())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userUuid string
		scanErr := rows.Scan(&userUuid)
		if scanErr != nil {
			continue
		}
		output.SetUserUuid(userUuid)
	}
	return nil
}

func (this *CMysqlHandler) ProPasswordIstrueByUseruuid(input *CProPasswordIstrueByUseruuidInput, output *CProPasswordIstrueByUseruuidOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_password_istrue_by_useruuid(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetUserUuid(), input.GetUserPwd())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var isTrue bool
		scanErr := rows.Scan(&isTrue)
		if scanErr != nil {
			continue
		}
		output.SetIsTrue(isTrue)
	}
	return nil
}

func (this *CMysqlHandler) ProPasswordIstrueByNo(input *CProPasswordIstrueByNoInput, output *CProPasswordIstrueByNoOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_password_istrue_by_no(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetNo(), input.GetUserPwd())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var isTrue bool
		var userUuid string
		scanErr := rows.Scan(&isTrue, &userUuid)
		if scanErr != nil {
			continue
		}
		output.SetIsTrue(isTrue)
		output.SetUserUuid(userUuid)
	}
	return nil
}

func (this *CMysqlHandler) ProDeleteUser(input *CProDeleteUserInput) (error) {
	stmt, err := this.m_db.Prepare("call pro_delete_user(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetUserUuid())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
	}
	return nil
}

