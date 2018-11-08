package main

import (
	comm_err "../common"
	cfgs "../proto/svr01cfgs"
	ress "../proto/svr02ress"
	cgws_http "../proto/svr10cgws"
	"encoding/json"
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	// "io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CAddPictureHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// response field
	var version float64 = 0.0
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	reply := cgws_http.CAddPictureReply{}
	for {
		// cookie is vaild
		cookieId := r.Header.Get(cgws_http.AddPictureResourceHeaderCookieid)
		picType := r.Header.Get(cgws_http.AddPictureResourceHeaderPictype)
		if cookieId == "" || picType == "" {
			errorCode = comm_err.ParamErrorCode
			errorString = comm_err.ParamErrorString + "cookieid pictype formdataname is null"
			break
		}
		isVaild, err := server.m_sessionMgr.SessionIsVaild(cookieId)
		if err != nil || isVaild == false {
			errorCode = comm_err.UserCookieIsVaildErrorCode
			errorString = comm_err.UserCookieIsVaildErrorString + "add picture error"
			break
		}
		// get ress info
		getServerInfoRequest := cfgs.CGetServerInfoRequest{ServerType: cfgs.Server_type_ress}
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
		url := ""
		if getCfgsServerInfoReply.ServerDomainname == "" {
			url = strings.Join([]string{"http://", getCfgsServerInfoReply.ServerIp, ":", strconv.FormatInt(int64(getCfgsServerInfoReply.ServerPort), 10), ress.AddPicture}, "")
		} else {
			url = strings.Join([]string{"http://", getCfgsServerInfoReply.ServerDomainname, ress.AddPicture}, "")
		}
		// redirect
		// http.Redirect(w, r, url, 307)
		w.Header().Set("Location", url)
		versionStr := strings.Join([]string{strconv.FormatInt(int64(r.ProtoMajor), 10), strconv.FormatInt(int64(r.ProtoMinor), 10)}, ".")
		version, err = strconv.ParseFloat(versionStr, len(versionStr))
		if err != nil {
			break
		}
		break
	}
	// response
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	// w.WriteHeader(httpCode)
	if errorCode != comm_err.OkErrorCode {
		w.WriteHeader(httpCode)
		io.WriteString(w, string(replayStr))
		return
	}
	if version < 1.1 {
		w.WriteHeader(http.StatusFound)
	} else {
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
	io.WriteString(w, string(replayStr))
}

func CDeletePictureHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// response field
	var version float64 = 0.0
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	reply := cgws_http.CDeletePictureReply{}
	for {
		// cookie is vaild
		cookieId := r.Header.Get(cgws_http.DeletePictureResourceHeaderCookid)
		pictureUrl := r.Header.Get(cgws_http.DeletePictureResourceHeaderPictureurl)
		if cookieId == "" || pictureUrl == "" {
			errorCode = comm_err.ParamErrorCode
			errorString = comm_err.ParamErrorString + "cookieid pictureurl is null"
			break
		}
		isVaild, err := server.m_sessionMgr.SessionIsVaild(cookieId)
		if err != nil || isVaild == false {
			errorCode = comm_err.UserCookieIsVaildErrorCode
			errorString = comm_err.UserCookieIsVaildErrorString + "add picture error"
			break
		}
		// get ress info
		getServerInfoRequest := cfgs.CGetServerInfoRequest{ServerType: cfgs.Server_type_ress}
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
		url := ""
		if getCfgsServerInfoReply.ServerDomainname == "" {
			url = strings.Join([]string{"http://", getCfgsServerInfoReply.ServerIp, ":", strconv.FormatInt(int64(getCfgsServerInfoReply.ServerPort), 10), ress.DeletePicture}, "")
		} else {
			url = strings.Join([]string{"http://", getCfgsServerInfoReply.ServerDomainname, ress.DeletePicture}, "")
		}
		// redirect
		// http.Redirect(w, r, url, 307)
		w.Header().Set("Location", url)
		versionStr := strings.Join([]string{strconv.FormatInt(int64(r.ProtoMajor), 10), strconv.FormatInt(int64(r.ProtoMinor), 10)}, ".")
		version, err = strconv.ParseFloat(versionStr, len(versionStr))
		if err != nil {
			break
		}
		break
	}
	// response
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	// w.WriteHeader(httpCode)
	if errorCode != comm_err.OkErrorCode {
		w.WriteHeader(httpCode)
		io.WriteString(w, string(replayStr))
		return
	}
	if version < 1.1 {
		w.WriteHeader(http.StatusFound)
	} else {
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
	io.WriteString(w, string(replayStr))
}
