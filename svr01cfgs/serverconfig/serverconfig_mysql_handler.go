package serverconfig

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

func (this *CMysqlHandler) ProAddServerInfo(input *CProAddServerInfoInput, output *CProAddServerInfoOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_add_server_info(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetServername(), input.GetServertype(), input.GetServerip(), input.GetServerport(), input.GetServerdomainname())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var serveruuid string
		scanErr := rows.Scan(&serveruuid)
		if scanErr != nil {
			continue
		}
		output.SetServeruuid(serveruuid)
	}
	return nil
}

func (this *CMysqlHandler) ProGetServerInfo(input *CProGetServerInfoInput, output *CProGetServerInfoOutput) (error) {
	stmt, err := this.m_db.Prepare("call pro_get_server_info(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	conn, err := this.m_db.Begin()
	if err != nil {
		return err
	}
	rows, err := stmt.Query(input.GetServertype())
	conn.Commit()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var serveruuid string
		var servername string
		var servertype string
		var serverip string
		var serverport int
		var serverdomainname string
		scanErr := rows.Scan(&serveruuid, &servername, &servertype, &serverip, &serverport, &serverdomainname)
		if scanErr != nil {
			continue
		}
		output.SetServeruuid(serveruuid)
		output.SetServername(servername)
		output.SetServertype(servertype)
		output.SetServerip(serverip)
		output.SetServerport(serverport)
		output.SetServerdomainname(serverdomainname)
	}
	return nil
}

