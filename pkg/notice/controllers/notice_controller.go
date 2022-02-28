package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/artifact"
	"golang-gin-boilerplate/pkg/notice/models"
	"golang-gin-boilerplate/pkg/notice/services"
	"net/http"
)

func NoticeIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		notices := services.AllNotice()

		artifact.Res.Data(notices).Json(c)
	}
}

func NoticeCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var notice models.Notice

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&notice); err != nil {
			artifact.Res.Status(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		notice = services.CreateANotice(notice)

		artifact.Res.Status(http.StatusCreated).Message("success").Data(notice).Json(c)
	}
}

func NoticeShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		noticeId := c.Param("noticeId")

		notice := services.ANotice(noticeId)

		artifact.Res.Status(http.StatusOK).Message("success").Data(notice).Json(c)
	}
}

func NoticeUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateNotice models.Notice

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		noticeId := c.Param("noticeId")

		if err := c.ShouldBind(&updateNotice); err != nil {
			artifact.Res.Status(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		notice, err := services.UpdateANotice(noticeId, updateNotice)

		if err != nil {
			artifact.Res.Status(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		artifact.Res.Status(http.StatusOK).Message("Successfully Updated !!!").Data(notice).Json(c)
	}
}

func NoticeDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		noticeId := c.Param("noticeId")
		notice, err := services.DeleteANotice(noticeId)

		if !err {
			artifact.Res.Status(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Status(http.StatusOK).Message("Successfully Delete !!!").Data(notice).Json(c)
	}
}
