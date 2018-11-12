package commoditymgr

import (
	"bufio"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type CDbHandler struct {
	m_db *sql.DB
}

func (this *CDbHandler) Connect(host string, port uint, username string, userpwd string, dbname string, dbtype string) (err error) {
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
	var name string
	if dbtype == "mysql" {
		name = b.String()
	} else if dbtype == "sqlite" {
		name = dbname
	} else {
		return errors.New("dbtype not support")
	}
	this.m_db, err = sql.Open(dbtype, name)
	if err != nil {
		return err
	}
	this.m_db.SetMaxOpenConns(2000)
	this.m_db.SetMaxIdleConns(1000)
	this.m_db.Ping()
	return nil
}

func (this *CDbHandler) ConnectByCfg(path string) error {
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
	var dbtype string = "mysql"
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
		case "dbtype":
			dbtype = v
		}
	}
	return this.Connect(host, port, username, userpwd, dbname, dbtype)
}

func (this *CDbHandler) Disconnect() {
	this.m_db.Close()
}

func (this *CDbHandler) Create() error {
	var err error = nil
	_, err = this.m_db.Exec(`create table if not exists t_object_detail(
uuid varchar(64),
-- text / picture / video
detailtype varchar(64),
-- desc / picurl / ...
detailvalue text,
detailno int,
objectuuid varchar(64)
);`)
	if err != nil {
		return err
	}
	_, err = this.m_db.Exec(`
create index UK_objectdetail_uuid on t_object_detail (uuid);`)
	if err != nil {
		return err
	}
	_, err = this.m_db.Exec(`
create table if not exists t_commodity_classifition(
uuid varchar(64),
parentuuid varchar(64),
name varchar(256)
);`)
	if err != nil {
		return err
	}
	_, err = this.m_db.Exec(`
create index UK_commodityclassifition_uuid on t_commodity_classifition (uuid);`)
	if err != nil {
		return err
	}
	_, err = this.m_db.Exec(`
create table if not exists t_commodity_info(
uuid varchar(64),
name varchar(256),
price decimal(10, 2),
classifyuuid varchar(64)
);`)
	if err != nil {
		return err
	}
	_, err = this.m_db.Exec(`
create index UK_commodityinfo_uuid on t_commodity_info (uuid);`)
	if err != nil {
		return err
	}
	return nil
}

func (this *CDbHandler) AddCommodityClassifition(input *[]CAddCommodityClassifitionInput) error {
	stmt, err := this.m_db.Prepare(fmt.Sprintf(`insert into t_commodity_classifition values(?, ?, ?);`))
	if err != nil {
		return err
	}
	defer stmt.Close()
	tx, _ := this.m_db.Begin()
	for _, v := range *input {
		rows, err := stmt.Query(v.Uuid, v.ParentUuid, v.Name)
		if err != nil {
			continue
		}
		defer rows.Close()
		for rows.Next() {
		}
	}
	tx.Commit()
	return nil
}

func (this *CDbHandler) AddCommodityClassifitionDetailInfo(input *[]CAddCommodityClassifitionDetailInfoInput) error {
	stmt, err := this.m_db.Prepare(fmt.Sprintf(`insert into t_object_detail values(?, ?, ?, ?, ?);`))
	if err != nil {
		return err
	}
	defer stmt.Close()
	tx, _ := this.m_db.Begin()
	for _, v := range *input {
		rows, err := stmt.Query(v.DetailUuid, v.DetailType, v.DetailValue, v.DetailNo, v.ClassifyUuid)
		if err != nil {
			continue
		}
		defer rows.Close()
		for rows.Next() {
		}
	}
	tx.Commit()
	return nil
}

func (this *CDbHandler) AddCommodityInfo(input *[]CAddCommodityInfoInput) error {
	stmt, err := this.m_db.Prepare(fmt.Sprintf(`insert into t_commodity_info values(?, ?, ?, ?);`))
	if err != nil {
		return err
	}
	defer stmt.Close()
	tx, _ := this.m_db.Begin()
	for _, v := range *input {
		rows, err := stmt.Query(v.Uuid, v.Name, v.Price, v.ClassifyUuid)
		if err != nil {
			continue
		}
		defer rows.Close()
		for rows.Next() {
		}
	}
	tx.Commit()
	return nil
}

func (this *CDbHandler) AddCommodityDetailInfo(input *[]CAddCommodityDetailInfoInput) error {
	stmt, err := this.m_db.Prepare(fmt.Sprintf(`insert into t_object_detail values(?, ?, ?, ?, ?);`))
	if err != nil {
		return err
	}
	defer stmt.Close()
	tx, _ := this.m_db.Begin()
	for _, v := range *input {
		rows, err := stmt.Query(v.Uuid, v.DetailType, v.DetailValue, v.DetialNo, v.CommodityUuid)
		if err != nil {
			continue
		}
		defer rows.Close()
		for rows.Next() {
		}
	}
	tx.Commit()
	return nil
}
