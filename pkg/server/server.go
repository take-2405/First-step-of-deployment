package server

import(
	"github.com/gin-gonic/gin"
	"F-deployment/pkg/server/handler"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init(){
	Server = gin.Default()
	Server.GET("/ping", handler.HandlePing())
}
