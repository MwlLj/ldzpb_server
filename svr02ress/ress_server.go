package main

import (
	ress "../proto/svr02ress"
	"../tools/configs"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

type CRessServer struct {
	m_httpRouter *httprouter.Router
	m_httpMutex  *http.ServeMux
}

func (this *CRessServer) Start() {
	// read http config
	httpConfig, err := configs.NewHttpConfig("ress.cfg")
	if err != nil {
		panic("[Error] read http config error")
	}
	httpData := httpConfig.GetHttpConfigData()
	// start http server
	this.m_httpRouter = httprouter.New()
	this.registerHttpRouter()
	http.ListenAndServe(strings.Join([]string{httpData.Host, ":", strconv.FormatInt(int64(httpData.Port), 10)}, ""), this.m_httpRouter)
}

func (this *CRessServer) registerHttpRouter() {
	this.m_httpRouter.POST(ress.AddCommodityPicture, CAddCommodityPicture)
	this.m_httpRouter.DELETE(ress.DeleteCommodityPicture, CDeleteCommodityPicture)
	this.m_httpRouter.POST(ress.AddPicture, CAddPicture)
	this.m_httpRouter.DELETE(ress.DeletePicture, CDeletePicture)
}

var server *CRessServer

func main() {
	server = &CRessServer{}
	server.Start()
}
