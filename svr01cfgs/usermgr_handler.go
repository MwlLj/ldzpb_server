package main

import (
	comm_err "../common"
	usermgrjson "../proto/svr01cfgs"
	usermgrdb "./usermgr"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/MwlLj/mqtt_comm"
	_ "github.com/go-sql-driver/mysql"
)

type CPostUserHandle struct {
}

func (this *CPostUserHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	server := user.(*CCfgsServer)
	// response field
	errorCode := 0
	errorString := "success"
	userUuid := ""
	var err error = nil
	for {
		// get request json
		postUserJson := usermgrjson.CPostUserRequest{}
		err = json.Unmarshal([]byte(request), &postUserJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "cfgs get postUserRequest struct error"
			break
		}
		// emailno is exist
		enIsExistInput := usermgrdb.CProUserIsexistsByEmailInput{}
		enIsExistOutput := usermgrdb.CProUserIsexistsByEmailOutput{}
		enIsExistInput.SetUserEmail(postUserJson.EmailNo)
		err = server.m_usermgrDbHandler.ProUserIsexistsByEmail(&enIsExistInput, &enIsExistOutput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "user isexist by email error"
			break
		}
		if enIsExistOutput.GetUserUuid() != "" {
			errorCode = comm_err.UserAlreadyExistErrorCode
			errorString = comm_err.UserAlreadyExistErrorString + "user email is exist"
			break
		}
		// phoneno is exist
		pnIsExistInput := usermgrdb.CProUserIsexistsByPhoneInput{}
		pnIsExistOutput := usermgrdb.CProUserIsexistsByPhoneOutput{}
		pnIsExistInput.SetUserPhone(postUserJson.PhoneNo)
		err = server.m_usermgrDbHandler.ProUserIsexistsByPhone(&pnIsExistInput, &pnIsExistOutput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "user isexist by phone error"
			break
		}
		if pnIsExistOutput.GetUserUuid() != "" {
			errorCode = comm_err.UserAlreadyExistErrorCode
			errorString = comm_err.UserAlreadyExistErrorString + "user phone is exist"
			break
		}
		// gen md5 pwd
		h := md5.New()
		h.Write([]byte(postUserJson.Userpwd))
		md5pwd := hex.EncodeToString(h.Sum(nil))
		input := usermgrdb.CProAddUserInput{}
		output := usermgrdb.CProAddUserOutput{}
		input.SetUserName(postUserJson.Username)
		input.SetUserPwd(md5pwd)
		input.SetUserEmailNo(postUserJson.EmailNo)
		input.SetUserPhoneNo(postUserJson.PhoneNo)
		err = server.m_usermgrDbHandler.ProAddUser(&input, &output)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "add user error"
			break
		}
		userUuid = output.GetUserUuid()
		break
	}
	reply := usermgrjson.CPostUserReply{Error: errorCode, ErrorString: errorString, Useruuid: userUuid}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CDeleteUserHandle struct {
}

func (this *CDeleteUserHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	server := user.(*CCfgsServer)
	// server := user.(*CCfgsServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		// get request json
		deleteUserJson := usermgrjson.CDeleteUserRequest{}
		err = json.Unmarshal([]byte(request), &deleteUserJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "cfgs get deleteUserRequest struct error"
			break
		}
		// delete user
		userUuid := deleteUserJson.Useruuid
		deleteUserInput := usermgrdb.CProDeleteUserInput{}
		deleteUserInput.SetUserUuid(userUuid)
		err = server.m_usermgrDbHandler.ProDeleteUser(&deleteUserInput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "delete user error"
			break
		}
		break
	}
	reply := usermgrjson.CDeleteUserReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CPostUserLoginHandle struct {
}

func (this *CPostUserLoginHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	server := user.(*CCfgsServer)
	// response field
	errorCode := 0
	errorString := "success"
	userUuid := ""
	var err error = nil
	for {
		// get request json
		postUserLoginJson := usermgrjson.CPostUserLoginRequest{}
		err = json.Unmarshal([]byte(request), &postUserLoginJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "cfgs get postUserLoginRequest struct error"
			break
		}
		userNo := postUserLoginJson.Userno
		userPwd := postUserLoginJson.Userpwd
		// gen md5 pwd
		h := md5.New()
		h.Write([]byte(userPwd))
		userPwd = hex.EncodeToString(h.Sum(nil))
		// no is exist
		noIsExistInput := usermgrdb.CProUserIsexistsByEmailOrPhoneInput{}
		noIsExistOutput := usermgrdb.CProUserIsexistsByEmailOrPhoneOutput{}
		noIsExistInput.SetUserNo(userNo)
		err = server.m_usermgrDbHandler.ProUserIsexistsByEmailOrPhone(&noIsExistInput, &noIsExistOutput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "user is exist by emailno or phoneno error"
			break
		}
		userUuid = noIsExistOutput.GetUserUuid()
		if userUuid == "" {
			errorCode = comm_err.UserUnRegisterErrorCode
			errorString = comm_err.UserUnRegisterErrorString + "user emailno or phoneno is not exist"
			break
		}
		// pwd is true
		pwdIsTrueInput := usermgrdb.CProPasswordIstrueByUseruuidInput{}
		pwdIsTrueOutput := usermgrdb.CProPasswordIstrueByUseruuidOutput{}
		pwdIsTrueInput.SetUserUuid(userUuid)
		pwdIsTrueInput.SetUserPwd(userPwd)
		err = server.m_usermgrDbHandler.ProPasswordIstrueByUseruuid(&pwdIsTrueInput, &pwdIsTrueOutput)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "password is true by useruuid error"
			break
		}
		isTrue := pwdIsTrueOutput.GetIsTrue()
		if !isTrue {
			errorCode = comm_err.UserPasswordIsFalseErrorCode
			errorString = comm_err.UserPasswordIsFalseErrorString + "user passwors error"
			break
		}
		break
	}
	reply := usermgrjson.CPostUserLoginReply{Error: errorCode, ErrorString: errorString, Useruuid: userUuid}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}
