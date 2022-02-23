package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/artifact"
)

var Application *App

var Router *artifact.Router

var Response artifact.ResponseBuilder

type App struct {
	Name     string
	Url      string
	Port     int
	Env      artifact.Env
	Response artifact.ResponseBuilder
	Logger   artifact.Logger
	Router   *artifact.Router
}

func (app App) Run() {
	app.Router.Run(fmt.Sprintf(":%d", app.Port))
}

func NewApp() *App {
	env := artifact.NewEnv()
	logger := artifact.NewLogger(env)

	Router = artifact.NewRouter()
	Response = artifact.ResponseBuilder{}

	gin.SetMode(env.App.GinMode)

	Application = &App{
		Name:     env.App.Name,
		Url:      env.App.Url,
		Port:     env.App.Port,
		Env:      env,
		Logger:   logger,
		Router:   Router,
		Response: Response,
	}

	return Application
}

func Run() *App {

	Application.Run()

	return Application
}
