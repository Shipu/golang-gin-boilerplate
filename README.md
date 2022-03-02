# golang-gin-boilerplate
Golang [Artifact](https://github.com/shipu/artifact) Boilerplate for Rest Api. It's based on [Gin](https://github.com/gin-gonic/gin)

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

Congratulation !!! You have successfully generated rest crud api. 

Now you can run your application with `go run main.go` command or `air`.

## Config :

Suppose your config is `config/db.go`:
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
and your `.env` is:
```dotenv
DB_CONNECTION=mongodb
DB_HOST=
DB_PORT=
DB_USER=
DB_PASS=
DB_DATABASE=
```

For initialization `DatabaseConfig` config. 

add below example code in `config/main.go`
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

func Boot() {
    ...
    ...
}
```

if you want to run something when application start, you can add anything in `Boot` function.

### To get config:
```go
Config.GetString("DB.Host")
```

Config Method List:
```go
GetString("key")
GetInt("key")
Get("key")
```

## Route

Suppose example route is `notice/routes/api.go`:
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

For Register example route. in `routes/main.go`

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

## Response
In [Gin](https://github.com/gin-gonic/gin)

Where `c` is the `*gin.Context` context.

```go
data := map[string]interface{}{
    "app": "Golang",
}
c.JSON(200, gin.H{
    "status_code":  200,
    "message": "Success",
    "data": data,
})
```
In artifact boilerplate, you can use `Res`.
```go
data := map[string]interface{}{
    "app": "Golang",
}

Res.Status(200).
    Message("Success").
    Data(data).
	Json(c)
```

`Res` Api Methods:
```go
Json(c *gin.Context)
AbortWithStatusJSON(c *gin.Context)
```

## Mongo Collection

```go
var TodoCollection artifact.MongoCollection = artifact.Mongo.Collection("todos")

TodoCollection.Find(bson.M{})
```

All [Go Mongo Driver](https://docs.mongodb.com/drivers/go/current/) Support.