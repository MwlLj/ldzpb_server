package main

import (
	comm_err "../common"
	cfgs_usermgr "../proto/svr01cfgs"
	cgws_http "../proto/svr10cgws"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	mqttTimeoutS     int    = 30
	sessionTimeoutNS uint64 = 30 * 60 * 1000 * 1000 * 1000
)

func CRegisterUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	userUuid := ""
	for {
		// change to request struct
		userRequest := cgws_http.CRegisterUserRequest{}
		err := json.Unmarshal(userRequestStr, &userRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv http request json decoding error"
			break
		}
		// send to cfgs
		postCfgsUserRequest := cfgs_usermgr.CPostUserRequest{Username: userRequest.Username, Userpwd: userRequest.Userpwd, EmailNo: userRequest.EmailNo, PhoneNo: userRequest.PhoneNo}
		postCfgsUserRequestStr, err := json.Marshal(postCfgsUserRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "send to cfgs post user request json encoding error"
			break
		}
		postCfgsUserReplyStr, err := server.m_mqttComm.Post(cfgs_usermgr.User, string(postCfgsUserRequestStr), 1, mqttTimeoutS)
		if err != nil {
			errorCode = comm_err.MqttCommErrorCode
			errorString = comm_err.MqttCommErrorString + "send to cfgs post user mqtt comminucation error"
			break
		}
		postCfgsUserReply := cfgs_usermgr.CPostUserReply{}
		err = json.Unmarshal([]byte(postCfgsUserReplyStr), &postCfgsUserReply)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv cfgs post user json decoding error"
			break
		}
		if postCfgsUserReply.Error != comm_err.OkErrorCode {
			errorCode = postCfgsUserReply.Error
			errorString = postCfgsUserReply.ErrorString
			break
		}
		userUuid = postCfgsUserReply.Useruuid
		break
	}
	// response
	reply := cgws_http.CRegisterUserReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	reply.Useruuid = userUuid
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}

func CUnRegisterUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	for {
		// change to request struct
		userRequest := cgws_http.CUnRegisterUserRequest{}
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
			errorString = comm_err.UserCookieIsVaildErrorString + "unregister error"
			break
		}
		// send to cfgs
		deleteCfgsUserRequest := cfgs_usermgr.CDeleteUserRequest{Useruuid: userRequest.Useruuid}
		deleteCfgsUserRequestStr, err := json.Marshal(deleteCfgsUserRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "send to cfgs delete user request json encoding error"
			break
		}
		deleteCfgsUserReplyStr, err := server.m_mqttComm.Delete(cfgs_usermgr.User, string(deleteCfgsUserRequestStr), 1, mqttTimeoutS)
		if err != nil {
			errorCode = comm_err.MqttCommErrorCode
			errorString = comm_err.MqttCommErrorString + "send to cfgs delete user mqtt comminucation error"
			break
		}
		deleteCfgsUserReply := cfgs_usermgr.CDeleteUserReply{}
		err = json.Unmarshal([]byte(deleteCfgsUserReplyStr), &deleteCfgsUserReply)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv cfgs delete user json decoding error"
			break
		}
		if deleteCfgsUserReply.Error != comm_err.OkErrorCode {
			errorCode = deleteCfgsUserReply.Error
			errorString = deleteCfgsUserReply.ErrorString
			break
		}
		break
	}
	// response
	reply := cgws_http.CUnRegisterUserReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}

func CLoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	userUuid := ""
	sessionId := ""
	for {
		// change to request struct
		userRequest := cgws_http.CUserLoginRequest{}
		err := json.Unmarshal(userRequestStr, &userRequest)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "recv http request json decoding error"
			break
		}
		// check cookie id is vaild
		cookieId := userRequest.Cookieid
		if cookieId == "" {
			// send to cfgs
			postCfgsUserLoginRequest := cfgs_usermgr.CPostUserLoginRequest{Userno: userRequest.Userno, Userpwd: userRequest.Userpwd}
			postCfgsUserLoginRequestStr, err := json.Marshal(postCfgsUserLoginRequest)
			if err != nil {
				errorCode = comm_err.JsonParseErrorCode
				errorString = comm_err.JsonParseErrorString + "send to cfgs post user login request json encoding error"
				break
			}
			postCfgsUserLoginReplyStr, err := server.m_mqttComm.Post(cfgs_usermgr.User_login, string(postCfgsUserLoginRequestStr), 1, mqttTimeoutS)
			if err != nil {
				errorCode = comm_err.MqttCommErrorCode
				errorString = comm_err.MqttCommErrorString + "send to cfgs post user login mqtt comminucation error"
				break
			}
			postCfgsUserLoginReply := cfgs_usermgr.CPostUserLoginReply{}
			err = json.Unmarshal([]byte(postCfgsUserLoginReplyStr), &postCfgsUserLoginReply)
			if err != nil {
				errorCode = comm_err.JsonParseErrorCode
				errorString = comm_err.JsonParseErrorString + "recv cfgs post user login json decoding error"
				break
			}
			if postCfgsUserLoginReply.Error != comm_err.OkErrorCode {
				errorCode = postCfgsUserLoginReply.Error
				errorString = postCfgsUserLoginReply.ErrorString
				break
			}
			// create session
			userUuid = postCfgsUserLoginReply.Useruuid
			sessionId, err = server.m_sessionMgr.CreateSession(sessionTimeoutNS, userUuid)
			if err != nil {
				errorCode = comm_err.DbExecuteErrorCode
				errorString = comm_err.DbExecuteErrorString + "create session error"
				break
			}
		} else {
			isVaild, e := server.m_sessionMgr.SessionIsVaild(cookieId)
			if e != nil {
				errorCode = comm_err.DbExecuteErrorCode
				errorString = comm_err.DbExecuteErrorString + "get sessionIsVaile error"
				break
			}
			if isVaild == false {
				errorCode = comm_err.UserCookieIsVaildErrorCode
				errorString = comm_err.UserCookieIsVaildErrorString + "user cookie is not exist"
				break
			} else {
				e := server.m_sessionMgr.ResetLosevaildTime(cookieId)
				if e != nil {
					errorCode = comm_err.DbExecuteErrorCode
					errorString = comm_err.DbExecuteErrorString + "reset losevaild time error"
					break
				}
				userdata, e := server.m_sessionMgr.GetUserdata(cookieId)
				if e != nil {
					errorCode = comm_err.DbExecuteErrorCode
					errorString = comm_err.DbExecuteErrorString + "get session userdata error"
					break
				}
				sessionId = cookieId
				userUuid = userdata
			}
		}
		break
	}
	// response
	reply := cgws_http.CUserLoginReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	reply.Sessionid = sessionId
	reply.Useruuid = userUuid
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}

func CLogoutHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	for {
		// change to request struct
		userRequest := cgws_http.CLogoutRequest{}
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
			errorString = comm_err.UserCookieIsVaildErrorString + "user session isvaild error"
			break
		}
		err = server.m_sessionMgr.DestroySession(cookieId)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "destroy session error"
			break
		}
		break
	}
	// response
	reply := cgws_http.CLogoutReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}
