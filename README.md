# golang-gin-boilerplate
Golang Gin Boilerplate for Rest Api

## For Generate Crud: 
```go
go run ./art crud package_name crud_module_name
```

### Example: 

Run 
```go 
go run ./art crud github.com/shipu/golang-gin-boilerplate notice
``` 
to generate crud.

Folder structure: 
```go
src/notice
├── controllers
│   └── notice_controller.go
├── models
│   └── notice.go
├── routes
│   └── api.go
└── services
└── notice_service.go
```

#### For Register Routes:

In `routes/main.go`
```go
import (
    ...	
    ...
    ...
    noticeRoute "github.com/shipu/golang-gin-boilerplate/src/notice/routes"
)

```
```go
func Register() {
    ...
    ...
    noticeRoute.Setup()
}
```
#### For Initialize Model Collection:

In `routes/main.go`
```go
import (
    ...	
    ...
    ...
    notice "github.com/shipu/golang-gin-boilerplate/src/notice/models"
)

```
```go
func Boot() {
    ...
    ...
    notice.Setup()
}
```

Done !!! 

And you can run your application with `go run main.go` command or `air`.

## Config :

Suppose your `config/db.go` config is:
```go
package config

type DatabaseConfig struct {
	Username   string `mapstructure:"DB_USER" default:""`
	Password   string `mapstructure:"DB_PASS" default:""`
	Host       string `mapstructure:"DB_HOST" default:""`
	Port       string `mapstructure:"DB_PORT" default:""`
	Database   string `mapstructure:"DB_DATABASE" default:""`
	Connection string `mapstructure:"DB_CONNECTION" default:""`
}
```
and your `.env` file is:
```dotenv
DB_CONNECTION=mongodb
DB_HOST=
DB_PORT=
DB_USER=
DB_PASS=
DB_DATABASE=
```

For initialization `DatabaseConfig` config. 

### Example
in `config/main.go`
```go
import (
    ...	
    ...
    ...
    . "github.com/shipu/artifact"
)

```
```go
func Register() {
    ...
    ...
    Config.AddConfig("DB", new(DatabaseConfig)).Load()
}
```

if you want to run something when application start, you can add anything in `Boot` function.

## Route

Suppose your route `notice/routes/api.go` is:
```go
package routes

import (
	. "github.com/shipu/artifact"
	c "github.com/shipu/golang-gin-boilerplate/src/notice/controllers"
)

func Setup() {
    v1 := Router.Group("api/v1")
    v1.GET("notices", c.NoticeIndex())
    v1.POST("notices", c.NoticeCreate())
    v1.GET("notices/:noticeId", c.NoticeShow())
    v1.PUT("notices/:noticeId", c.NoticeUpdate())
    v1.DELETE("notices/:noticeId", c.NoticeDelete())
}
```

For Register your route. in `routes/main.go`

```go
import (
    ...	
    ...
    ...
    noticeRoute "github.com/shipu/golang-gin-boilerplate/src/notice/routes"
)

```
```go
func Register() {
    ...
    ...
    noticeRoute.Setup()
}
```