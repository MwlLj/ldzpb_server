package main

import (
	cfgs "../proto/svr01cfgs"
	"../tools/configs"
	serverconfigdb "./serverconfig"
	usermgrdb "./usermgr"
	"fmt"
	"github.com/MwlLj/mqtt_comm"
	_ "github.com/go-sql-driver/mysql"
)

type CCfgsServer struct {
	m_mqttComm              mqtt_comm.CMqttComm
	m_usermgrDbHandler      usermgrdb.CMysqlHandler
	m_serverconfigDbHandler serverconfigdb.CMysqlHandler
}

func (this *CCfgsServer) Start() {
	// connect db
	dbcfg, err := configs.NewMysqlDbConfig("db.cfg")
	if err != nil {
		panic("[Error] open db config error")
	}
	dbData := dbcfg.GetMysqlDbData()
	// usermgr db conncet
	err = this.m_usermgrDbHandler.Connect(dbData.Host, dbData.Port, dbData.Username, dbData.Userpwd, dbData.Dbname)
	if err != nil {
		panic("[Error] usermgr connect db error")
	}
	defer this.m_usermgrDbHandler.Disconnect()
	// serverconfig db connect
	err = this.m_serverconfigDbHandler.Connect(dbData.Host, dbData.Port, dbData.Username, dbData.Userpwd, dbData.Dbname)
	if err != nil {
		panic("[Error] serverconfig connect db error")
	}
	defer this.m_serverconfigDbHandler.Disconnect()
	// start mqtt server
	this.m_mqttComm = mqtt_comm.NewMqttComm("cfgs", "1.0", 0)
	// read messagebus config
	mbcfg, err := configs.NewMessagebusConfig("messagebus.cfg")
	if err != nil {
		// panic("read messagebus config error")
	}
	messagebusData := mbcfg.GetMessagebusData()
	this.m_mqttComm.SetMessageBus(messagebusData.Host, int(messagebusData.Port), messagebusData.Username, messagebusData.Userpwd)
	this.registerRouter()
	fmt.Println("start success")
	this.m_mqttComm.Connect(true)
}

func (this *CCfgsServer) registerRouter() {
	this.m_mqttComm.Subscribe(mqtt_comm.POST, cfgs.User, 0, &CPostUserHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.DELETE, cfgs.User, 0, &CDeleteUserHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.POST, cfgs.User_login, 0, &CPostUserLoginHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.POST, cfgs.Server_info, 0, &CPostServerInfoHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.GET, cfgs.Server_info, 0, &CGetServerInfoHandle{}, this)
}

func main() {
	server := CCfgsServer{}
	server.Start()
}
