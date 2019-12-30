package group

import (
	"github.com/gin-gonic/gin"
	. "vid/app/controller"
)

func SetupRawGroup(router *gin.Engine) {

	rawGroup := router.Group("/raw")
	{
		rawGroup.GET("/image/:uid/:filename", RawController.RawImage)
	}
}
