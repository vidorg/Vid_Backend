package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shomali11/util/xconditions"
	"github.com/vidorg/vid_backend/src/config"
	"github.com/vidorg/vid_backend/src/controller/exception"
	"github.com/vidorg/vid_backend/src/model/dto/common"
	"github.com/vidorg/vid_backend/src/util"
	"net/http"
	"strconv"
)

type rawController struct {
	config *config.ServerConfig
}

func RawController(config *config.ServerConfig) *rawController {
	return &rawController{
		config: config,
	}
}

// @Router 				/raw/image/{uid}/{filename} [GET]
// @Summary 			获取图片
// @Description 		获取用户头像图片以及视频封面
// @Tag					Raw
// @Param 				uid path string true "用户id，或者default"
// @Param 				filename path string true "图片文件名，jpg后缀名"
// @Accept 				multipart/form-data
// @ErrorCode 			400 request param error
// @ErrorCode 			404 image not found
/* @Success 200 		| Key | Value |
						| --- | --- |
 						| Content-Type | image/jpeg | */
func (r *rawController) RawImage(c *gin.Context) {
	uidString := c.Param("uid")
	uid := -1
	if uidString != "default" {
		var err error
		uid, err = strconv.Atoi(uidString)
		if err != nil {
			common.Result{}.Error(http.StatusBadRequest).SetMessage(exception.RequestParamError.Error()).JSON(c)
			return
		}
	}
	filename := c.Param("filename")

	filePath := xconditions.IfThenElse(uid == -1, fmt.Sprintf("./usr/default/%s", filename), fmt.Sprintf("./usr/image/%d/%s", uid, filename)).(string)
	if !util.CommonUtil.IsDirOrFileExist(filePath) {
		common.Result{}.Error(http.StatusBadRequest).SetMessage(exception.ImageNotFoundError.Error()).JSON(c)
		return
	}

	c.Writer.Header().Add("Content-Type", "image/jpeg")
	c.File(filePath)
}
