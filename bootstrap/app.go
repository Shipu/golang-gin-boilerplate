package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/config"
	"net/http"
)

var Res ResponseBuilder

var Log LoggerBuilder

func loadRoute() {

	gin.ForceConsoleColor()

	//gin.SetMode("debug")

	Router = gin.Default()

	//httpRouter.SetTrustedProxies([]string{"0.0.0.0"})

	Router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running"})
	})
}

func loadConfig() {
	Config = NewConfig()
	Config.AddConfig("App", new(config.AppConfig))
	Config.AddConfig("Database", new(config.DatabaseConfig))

	Config.Load()
}

func initializeLogger(isLocal string) LoggerBuilder {
	return NewLogger(isLocal)
}

func init() {
	loadRoute()
	loadConfig()
	isLocal, _ := Config.GetString("App.Environment")
	Log = initializeLogger(isLocal)
}

func Start() {

}

func Run() {
	port, _ := Config.Get("App.Port")

	Router.Run(fmt.Sprintf(":%d", port))
}
