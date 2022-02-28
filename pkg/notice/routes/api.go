package routes

import (
	"golang-gin-boilerplate/artifact"
	. "golang-gin-boilerplate/pkg/notice/controllers"
)

func Setup() {
	artifact.Router.GET("notices", NoticeIndex())
	artifact.Router.POST("notices", NoticeCreate())
	artifact.Router.GET("notices/:noticeId", NoticeShow())
	artifact.Router.PUT("notices/:noticeId", NoticeUpdate())
	artifact.Router.DELETE("notices/:noticeId", NoticeDelete())
}
