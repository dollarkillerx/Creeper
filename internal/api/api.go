package api

import (
	"github.com/dollarkillerx/creeper/internal/conf"
	"github.com/dollarkillerx/creeper/internal/server"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	app *gin.Engine
	ser *server.Server
}

func New(ser *server.Server) *ApiServer {
	app := gin.New()
	app.Use(gin.Recovery())
	return &ApiServer{
		app: app,
		ser: ser,
	}
}

func (a *ApiServer) Run() error {
	a.router()
	return a.app.Run(conf.CONFIG.ListenAddr)
}
