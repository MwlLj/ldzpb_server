package main

import (
	comm_err "../common"
	cfgs "../proto/svr01cfgs"
	cgws_http "../proto/svr10cgws"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CAddServerInfoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	serverUuid := ""
	for {
		// change to request struct
		userRequest := cgws_http.CAddServerInfoRequest{}
		err := json.Unmarshal(userRequestStr, &userRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv http request json decoding error"
			break
		}
		// cookie is vaild
		cookieId := userRequest.Cookieid
		isVaild, err := server.m_sessionMgr.SessionIsVaild(cookieId)
		if err != nil || isVaild == false {
			errorCode = comm_err.UserCookieIsVaildErrorCode
			errorString = comm_err.UserCookieIsVaildErrorString + "addserverinfo error"
			break
		}
		// send to cfgs
		postServerInfoRequest := cfgs.CPostServerInfoRequest{ServerType: userRequest.Servertype, ServerName: userRequest.Servername, ServerIp: userRequest.Serverip, ServerPort: userRequest.Serverport, ServerDomainname: userRequest.Serverdomainname}
		postServerInfoRequestStr, err := json.Marshal(postServerInfoRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "send to cfgs post serverinfo request json encoding error"
			break
		}
		postCfgsServerInfoReplyStr, err := server.m_mqttComm.Post(cfgs.Server_info, string(postServerInfoRequestStr), 1, mqttTimeoutS)
		if err != nil {
			errorCode = comm_err.MqttCommErrorCode
			errorString = comm_err.MqttCommErrorString + "send to cfgs post serverinfo mqtt comminucation error"
			break
		}
		postCfgsServerInfoReply := cfgs.CPostServerInfoReply{}
		err = json.Unmarshal([]byte(postCfgsServerInfoReplyStr), &postCfgsServerInfoReply)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv cfgs post serverinfo json decoding error"
			break
		}
		if postCfgsServerInfoReply.Error != comm_err.OkErrorCode {
			errorCode = postCfgsServerInfoReply.Error
			errorString = postCfgsServerInfoReply.ErrorString
			break
		}
		serverUuid = postCfgsServerInfoReply.Serveruuid
		break
	}
	// response
	reply := cgws_http.CAddServerInfoReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	reply.Serveruuid = serverUuid
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}

func CGetServerInfoHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	reply := cgws_http.CGetServerInfoReply{}
	for {
		// cookie is vaild
		cookieId := r.Header.Get(cgws_http.GetServerInfoHeaderCookieid)
		serverType := r.Header.Get(cgws_http.GetServerInfoHeaderServertype)
		if cookieId == "" || serverType == "" {
			errorCode = comm_err.ParamErrorCode
			errorString = comm_err.ParamErrorString + "cookieid and servertype is null"
			break
		}
		isVaild, err := server.m_sessionMgr.SessionIsVaild(cookieId)
		if err != nil || isVaild == false {
			errorCode = comm_err.UserCookieIsVaildErrorCode
			errorString = comm_err.UserCookieIsVaildErrorString + "addserverinfo error"
			break
		}
		// send to cfgs
		getServerInfoRequest := cfgs.CGetServerInfoRequest{ServerType: serverType}
		getServerInfoRequestStr, err := json.Marshal(getServerInfoRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "send to cfgs get serverinfo request json encoding error"
			break
		}
		getCfgsServerInfoReplyStr, err := server.m_mqttComm.Get(cfgs.Server_info, string(getServerInfoRequestStr), 1, mqttTimeoutS)
		if err != nil {
			errorCode = comm_err.MqttCommErrorCode
			errorString = comm_err.MqttCommErrorString + "send to cfgs get serverinfo mqtt comminucation error"
			break
		}
		getCfgsServerInfoReply := cfgs.CGetServerInfoReply{}
		err = json.Unmarshal([]byte(getCfgsServerInfoReplyStr), &getCfgsServerInfoReply)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv cfgs get serverinfo json decoding error"
			break
		}
		if getCfgsServerInfoReply.Error != comm_err.OkErrorCode {
			errorCode = getCfgsServerInfoReply.Error
			errorString = getCfgsServerInfoReply.ErrorString
			break
		}
		reply.Serverdomainname = getCfgsServerInfoReply.ServerDomainname
		reply.Serverip = getCfgsServerInfoReply.ServerIp
		reply.Serverport = getCfgsServerInfoReply.ServerPort
		break
	}
	// response
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}
