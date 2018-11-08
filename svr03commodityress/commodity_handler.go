package main

import (
	comm_err "../common"
	crs "../proto/svr03commodityress/"
	commoditydb "./commoditymgr"
	"encoding/json"
	"github.com/MwlLj/mqtt_comm"
	_ "github.com/go-sql-driver/mysql"
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
		postCommodityClassifitionJson := crs.CPostCommodityClassifitionRequest{}
		err = json.Unmarshal([]byte(request), &postCommodityClassifitionJson)
		if err != nil {
			errorCode = comm_err.JsonParseErrorCode
			errorString = comm_err.JsonParseErrorString + "crs get postCommodityClassifitionRequest struct error"
			break
		}
		// add server to db
		input := commoditydb.CProAddCommodityClassifitionInput{}
		output := commoditydb.CProAddCommodityClassifitionOutput{}
		server.m_commodityDbHandler.ProAddCommodityClassifition(&input, &output)
		break
	}
	reply := crs.CPostCommodityClassifitionReply{Error: errorCode, ErrorString: errorString}
	replayStr, err := json.Marshal(reply)
	return string(replayStr), err
}
