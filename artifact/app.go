package artifact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/config"
	"net/http"
)

var Res ResponseBuilder

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
	Config.AddConfig("DB", new(config.DatabaseConfig))

	Config.Load()
}

func initializeLogger() LoggerBuilder {
	return NewLogger()
}

func connectDb() {
	Mongo = NewMongoDB()
}

func init() {
	loadRoute()
	loadConfig()
	connectDb()
}

func Start() {
	initializeLogger()

}

func Run() {
	defer Mongo.Client.Disconnect(Mongo.Ctx)
	port, _ := Config.Get("App.Port")

	Router.Run(fmt.Sprintf(":%d", port))
}
