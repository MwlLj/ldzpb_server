package main

import (
	crs "../proto/svr03commodityress"
	"../tools/configs"
	commoditydb "./commoditymgr"
	"fmt"
	"github.com/MwlLj/mqtt_comm"
	_ "github.com/go-sql-driver/mysql"
)

type CCommodityresServer struct {
	m_mqttComm           mqtt_comm.CMqttComm
	m_commodityDbHandler commoditydb.CDbHandler
}

func (this *CCommodityresServer) Start() {
	// connect db
	dbcfg, err := configs.NewMysqlDbConfig("db.cfg")
	if err != nil {
		panic("[Error] open db config error")
	}
	dbData := dbcfg.GetMysqlDbData()
	// usermgr db conncet
	err = this.m_commodityDbHandler.Connect(dbData.Host, dbData.Port, dbData.Username, dbData.Userpwd, dbData.Dbname, "mysql")
	if err != nil {
		panic("[Error] commodity connect db error")
	}
	defer this.m_commodityDbHandler.Disconnect()
	this.m_commodityDbHandler.Create()
	// serverconfig db connect
	// start mqtt server
	this.m_mqttComm = mqtt_comm.NewMqttComm("commodityress", "1.0", 0)
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

func (this *CCommodityresServer) registerRouter() {
	this.m_mqttComm.Subscribe(mqtt_comm.POST, crs.Commodity_classifition, 0, &CPostCommodityClassifitionHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.PUT, crs.Commodity_classifition, 0, &CPutCommodityClassifitionHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.DELETE, crs.Commodity_classifition, 0, &CDeleteCommodityClassifitionHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.GET, crs.Commodity_classifition, 0, &CGetCommodityClassifitionHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.POST, crs.Commodity_classifition_detail, 0, &CPostCommodityClassifitionDetailHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.PUT, crs.Commodity_classifition_detail, 0, &CPutCommodityClassifitionDetailHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.DELETE, crs.Commodity_classifition_detail, 0, &CDeleteCommodityClassifitionDetailHandle{}, this)
	this.m_mqttComm.Subscribe(mqtt_comm.GET, crs.Commodity_classifition_detail, 0, &CGetCommodityClassifitionDetailHandle{}, this)
}

func main() {
	server := CCommodityresServer{}
	server.Start()
}
