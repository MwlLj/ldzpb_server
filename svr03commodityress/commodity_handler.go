package main

import (
	comm_err "../common"
	crs "../proto/svr03commodityress/"
	commoditydb "./commoditymgr"
	"encoding/json"
	"github.com/MwlLj/mqtt_comm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
)

type CPostCommodityClassifitionHandle struct {
}

func (this *CPostCommodityClassifitionHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		// get request json
		postCommodityClassifitionJson := []crs.CPostCommodityClassifitionRequest{}
		err = json.Unmarshal([]byte(request), &postCommodityClassifitionJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs get postCommodityClassifitionRequest struct error"
			break
		}
		input := []commoditydb.CAddCommodityClassifitionInput{}
		for _, req := range postCommodityClassifitionJson {
			in := commoditydb.CAddCommodityClassifitionInput{}
			ud, err := uuid.NewV4()
			if err != nil {
				errorCode = comm_err.UuidGeneralErrorCode
				errorString = comm_err.UuidGeneralErrorString + "crs add commodityclassify"
				break
			}
			in.Uuid = ud.String()
			in.ParentUuid = req.ParentUuid
			in.Name = req.Name
			input = append(input, in)
		}
		err = server.m_commodityDbHandler.AddCommodityClassifition(&input)
		if err != nil {
			errorCode = comm_err.DbExecuteErrorCode
			errorString = comm_err.DbExecuteErrorString + "crs add commodityClassify error"
			break
		}
		break
	}
	reply := crs.CPostCommodityClassifitionReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CPutCommodityClassifitionHandle struct {
}

func (this *CPutCommodityClassifitionHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		// get request json
		putCommodityClassifitionJson := []crs.CPutCommodityClassifitionRequest{}
		err = json.Unmarshal([]byte(request), &putCommodityClassifitionJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs get putCommodityClassifitionRequest struct error"
			break
		}
		break
	}
	reply := crs.CPutCommodityClassifitionReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CDeleteCommodityClassifitionHandle struct {
}

func (this *CDeleteCommodityClassifitionHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		deleteCommodityClassifitionJson := []crs.CDeleteCommodityClassifitionRequest{}
		err = json.Unmarshal([]byte(request), &deleteCommodityClassifitionJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs get deleteCommodityClassifitionRequest struct error"
			break
		}
		break
	}
	reply := crs.CDeleteCommodityClassifitionReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CGetCommodityClassifitionHandle struct {
}

func (this *CGetCommodityClassifitionHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		getCommodityClassifitionJson := []crs.CGetCommodityClassifitionRequest{}
		err = json.Unmarshal([]byte(request), &getCommodityClassifitionJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs get commodityClassifitionRequest struct error"
			break
		}
		break
	}
	reply := crs.CGetCommodityClassifitionReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CPostCommodityClassifitionDetailHandle struct {
}

func (this *CPostCommodityClassifitionDetailHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		postCommodityClassifitionDetailJson := []crs.CPostCommodityClassifitionDetailRequest{}
		err = json.Unmarshal([]byte(request), &postCommodityClassifitionDetailJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs postCommodityClassifitionDetailJson struct error"
			break
		}
		break
	}
	reply := crs.CPostCommodityClassifitionDetailReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CPutCommodityClassifitionDetailHandle struct {
}

func (this *CPutCommodityClassifitionDetailHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		putCommodityClassifitionDetailJson := []crs.CPutCommodityClassifitionDetailRequest{}
		err = json.Unmarshal([]byte(request), &putCommodityClassifitionDetailJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs putCommodityClassifitionDetailJson struct error"
			break
		}
		break
	}
	reply := crs.CPutCommodityClassifitionDetailReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CDeleteCommodityClassifitionDetailHandle struct {
}

func (this *CDeleteCommodityClassifitionDetailHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		deleteCommodityClassifitionDetailJson := []crs.CDeleteCommodityClassifitionDetailRequest{}
		err = json.Unmarshal([]byte(request), &deleteCommodityClassifitionDetailJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs deleteCommodityClassifitionDetailJson struct error"
			break
		}
		break
	}
	reply := crs.CDeleteCommodityClassifitionDetailReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}

type CGetCommodityClassifitionDetailHandle struct {
}

func (this *CGetCommodityClassifitionDetailHandle) Handle(topic string, request string, mc mqtt_comm.CMqttComm, user interface{}) (string, error) {
	// server := user.(*CCommodityresServer)
	// response field
	errorCode := 0
	errorString := "success"
	var err error = nil
	for {
		getCommodityClassifitionDetailJson := []crs.CGetCommodityClassifitionDetailRequest{}
		err = json.Unmarshal([]byte(request), &getCommodityClassifitionDetailJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs getCommodityClassifitionDetailJson struct error"
			break
		}
		break
	}
	reply := crs.CGetCommodityClassifitionDetailReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}
