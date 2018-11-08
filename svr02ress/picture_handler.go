package main

import (
	comm_err "../common"
	ress_proto "../proto/svr02ress"
	"../tools/httpfile"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	// "io/ioutil"
	"../tools/timetool"
	// "fmt"
	"net/http"
	"os"
	"strings"
)

func pictype2path(pictype string) string {
	var path string = ""
	switch pictype {
	default:
		path = "unknow"
	}
	return path
}

func CAddPicture(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	urlInfos := make(map[string][]string)
	var urls []string
	var err error
	for {
		pictype := r.Header.Get(ress_proto.AddPictureHeaderPictypeKey)
		if pictype == "" {
			errorCode = comm_err.ParamErrorCode
			errorString = comm_err.ParamErrorString + "picturetype is null"
			break
		}
		// download picture
		path := pictype2path(pictype)
		urls, err = httpfile.DownloadFile(r, ress_proto.PictureFormname, strings.Join([]string{path, timetool.GetNowDayFormat()}, "/"), true)
		if err != nil {
			errorCode = comm_err.DownloadPictureFailedErrorCode
			errorString = comm_err.DownloadPictureFailedErrorString + "addPicture error"
			break
		}
		urlInfos[ress_proto.PictureFormname] = urls
		break
	}
	// response
	reply := ress_proto.CAddPictureReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	reply.UrlInfos = urlInfos
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}

func CDeletePicture(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// userRequestStr, _ := ioutil.ReadAll(r.Body)
	// response field
	httpCode := http.StatusOK
	errorCode := comm_err.OkErrorCode
	errorString := comm_err.OkErrorString
	for {
		picUrl := r.Header.Get(ress_proto.DeletePictureHeaderPicurl)
		if picUrl == "" {
			errorCode = comm_err.DeletePicturePicurlIsnullErrorCode
			errorString = comm_err.DeletePicturePicurlIsnullErrorString + "deletePicture error"
			break
		}
		err := os.Remove(picUrl)
		if err != nil {
			errorCode = comm_err.DeleteFileErrorCode
			errorString = comm_err.DeleteFileErrorString + "deletePicture error"
			break
		}
		break
	}
	// response
	reply := ress_proto.CDeletePictureReply{}
	reply.Error = errorCode
	reply.ErrorString = errorString
	replayStr, _ := json.Marshal(reply)
	w.WriteHeader(httpCode)
	io.WriteString(w, string(replayStr))
}
