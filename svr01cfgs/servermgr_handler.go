package main

import (
	comm_err "../common"
	servermgrjson "../proto/svr01cfgs"
	servermgrdb "./serverconfig"
	"encoding/json"
	"github.com/MwlLj/mqtt_comm"
	_ "github.com/go-sql-driver/mysql"
)

type CPostServerInfoHandle struct {
}

func (this *CPostServerInfoHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	server := user.(*CCfgsServer)
	// response field
	errorCode := 0
	errorString := "success"
	serverInfoUuid := ""
	var err error = nil
	for {
		// get request json
		postServerInfoJson := servermgrjson.CPostServerInfoRequest{}
		err = json.Unmarshal([]byte(request), &postServerInfoJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "cfgs get postServerInfoRequest struct error"
			break
		}
		// add server to db
		postServerInfoInput := servermgrdb.CProAddServerInfoInput{}
		postServerInfoOutput := servermgrdb.CProAddServerInfoOutput{}
		postServerInfoInput.SetServername(postServerInfoJson.ServerName)
		postServerInfoInput.SetServertype(postServerInfoJson.ServerType)
		postServerInfoInput.SetServerip(postServerInfoJson.ServerIp)
		postServerInfoInput.SetServerport(postServerInfoJson.ServerPort)
		postServerInfoInput.SetServerdomainname(postServerInfoJson.ServerDomainname)
		err = server.m_serverconfigDbHandler.ProAddServerInfo(&postServerInfoInput, &postServerInfoOutput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "add server info error"
			break
		}
		serverInfoUuid = postServerInfoOutput.GetServeruuid()
		break
	}
	reply := servermgrjson.CPostServerInfoReply{Error: errorCode, ErrorString: errorString, Serveruuid: serverInfoUuid}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CGetServerInfoHandle struct {
}

func (this *CGetServerInfoHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	server := user.(*CCfgsServer)
	// response field
	errorCode := 0
	errorString := "success"
	reply := servermgrjson.CGetServerInfoReply{}
	var err error = nil
	for {
		// get request json
		getServerInfoJson := servermgrjson.CGetServerInfoRequest{}
		err = json.Unmarshal([]byte(request), &getServerInfoJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "cfgs get getServerInfoRequest struct error"
			break
		}
		// get server from db
		getServerInfoInput := servermgrdb.CProGetServerInfoInput{}
		getServerInfoOutput := servermgrdb.CProGetServerInfoOutput{}
		getServerInfoInput.SetServertype(getServerInfoJson.ServerType)
		err = server.m_serverconfigDbHandler.ProGetServerInfo(&getServerInfoInput, &getServerInfoOutput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "get server info error"
			break
		}
		reply.ServerDomainname = getServerInfoOutput.GetServerdomainname()
		reply.ServerIp = getServerInfoOutput.GetServerip()
		reply.ServerPort = getServerInfoOutput.GetServerport()
		break
	}
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}
