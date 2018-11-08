package main

import (
	cgws "../proto/svr10cgws"
	"../tools/configs"
	sessionmgr "../tools/sessionmgr"
	"fmt"
	"github.com/MwlLj/mqtt_comm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

type CCgwsServer struct {
	m_mqttComm   mqtt_comm.CMqttComm
	m_httpRouter *httprouter.Router
	m_sessionMgr sessionmgr.CSessionMgr
}

func (this *CCgwsServer) Start() {
	go func() {
		// start mqtt server
		this.m_mqttComm = mqtt_comm.NewMqttComm("cgws", "1.0", 0)
		// read messagebus config
		mbcfg, err := configs.NewMessagebusConfig("messagebus.cfg")
		if err != nil {
			panic("[Error] read messagebus config error")
		}
		messagebusData := mbcfg.GetMessagebusData()
		this.m_mqttComm.SetMessageBus(messagebusData.Host, int(messagebusData.Port), messagebusData.Username, messagebusData.Userpwd)
		fmt.Println("start success cgws")
		this.m_mqttComm.Connect(true)
	}()
	// session mgr
	this.m_sessionMgr = sessionmgr.New(sessionmgr.Memory_type_mysql)
	defer sessionmgr.Destroy(this.m_sessionMgr)
	// read http config
	httpConfig, err := configs.NewHttpConfig("http.cfg")
	if err != nil {
		panic("[Error] read http config error")
	}
	httpData := httpConfig.GetHttpConfigData()
	// start http server
	this.m_httpRouter = httprouter.New()
	this.registerHttpRouter()
	http.ListenAndServe(strings.Join([]string{httpData.Host, ":", strconv.FormatInt(int64(httpData.Port), 10)}, ""), this.m_httpRouter)
}

func (this *CCgwsServer) registerHttpRouter() {
	this.m_httpRouter.POST(cgws.UserRegister, CRegisterUserHandler)
	this.m_httpRouter.DELETE(cgws.UserUnRegister, CUnRegisterUserHandler)
	this.m_httpRouter.POST(cgws.UserLogin, CLoginHandler)
	this.m_httpRouter.DELETE(cgws.UserLogout, CLogoutHandler)
	this.m_httpRouter.POST(cgws.ServerInfo, CAddServerInfoHandler)
	this.m_httpRouter.GET(cgws.ServerInfo, CGetServerInfoHandle)
	this.m_httpRouter.POST(cgws.ResourcePicture, CAddPictureHandler)
	this.m_httpRouter.DELETE(cgws.ResourcePicture, CDeletePictureHandler)
}

var server *CCgwsServer

func main() {
	server = &CCgwsServer{}
	server.Start()
}
