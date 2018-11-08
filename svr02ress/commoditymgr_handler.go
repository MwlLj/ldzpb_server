package main

import (
	comm_err "../common"
	ress_proto "../proto/svr02ress"
	"../tools/httpfile"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	// "io/ioutil"
	"net/http"
)

func CAddCommodityPicture(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	urlInfos := make(map[string][]string)
	var urls []string
	var err error
	for {
		// download commodity picture
		urls, err = httpfile.DownloadFile(r, ress_proto.CommodityPictureFormname, ress_proto.CommodityRoot, true)
		if err != nil {
			errorCode = comm_err.DownloadPictureFailedErrorCode
			errorString = comm_err.DownloadPictureFailedErrorString
			break
		}
		urlInfos[ress_proto.CommodityPictureFormname] = urls
		// download commodity desc picture
		urls, err = httpfile.DownloadFile(r, ress_proto.CommodityDescPictureFormname, ress_proto.CommodityRoot, true)
		if err != nil {
			errorCode = comm_err.DownloadPictureFailedErrorCode
			errorString = comm_err.DownloadPictureFailedErrorString
			break
		}
		urlInfos[ress_proto.CommodityDescPictureFormname] = urls
		break
	}
	// response
	reply := ress_proto.CAddCommodityPictureReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	reply.UrlInfos = urlInfos
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}

func CDeleteCommodityPicture(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	for {
		break
	}
	// response
	reply := ress_proto.CAddCommodityPictureReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}
