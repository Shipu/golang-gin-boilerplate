package routes

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/shipu/artifact"
	todoRoute "github.com/shipu/golang-gin-boilerplate/src/todo/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Register() {
	BaseRoute()
	SwaggerRoute()
	EnableCors()

	todoRoute.TodoSetup()
}

func BaseRoute() {
	// Example Route
	Router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"app": Config.GetString("App.Name"),
		}

		//or
		//data := gin.H{
		//	"message": "Hello World",
		//}

		Res.Code(200).
			Message("success").
			Data(data).Json(c)
	})
}

func EnableCors() {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With", "Referer", "guest", "publicKey", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	Router.Use(cors.New(corsConfig))
}

func SwaggerRoute() {
	Router.GET("swagger/swagger.json", func(c *gin.Context) {
		f, err := os.Open("./docs/swagger.json")
		if err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}
		defer f.Close()

		dec := json.NewDecoder(f)
		data := make(map[string]interface{}, 0)
		if err := dec.Decode(&data); err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}
		appUrl := Config.GetString("App.Url")
		appUrl = strings.Replace(appUrl, "https", "", -1)
		appUrl = strings.Replace(appUrl, "http", "", -1)
		data["host"] = appUrl

		Res.
			Raw(data).
			Json(c)
	})

	Router.GET("swagger/swagger.yaml", func(c *gin.Context) {
		f, err := ioutil.ReadFile("./docs/swagger.yaml")
		if err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}

		data := make(map[string]interface{}, 0)
		err = yaml.Unmarshal([]byte(f), &data)
		if err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}
		appUrl := Config.GetString("App.Url")
		appUrl = strings.Replace(appUrl, "https://", "", -1)
		appUrl = strings.Replace(appUrl, "http://", "", -1)
		data["host"] = appUrl

		Res.
			Raw(data).
			Yaml(c)
	})

	Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(Config.GetString("App.Url")+"/swagger/swagger.yaml"), ginSwagger.DefaultModelsExpandDepth(-1)))
}
