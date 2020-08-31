package param

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/vidorg/vid_backend/src/model/po"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("InsertVideoParam", "视频请求参数").
			Properties(
				goapidoc.NewProperty("title", "string", true, "标题，长度在 [1, 100] 之间"),
				goapidoc.NewProperty("description", "string", true, "简介，长度在 [0, 255] 之间").AllowEmpty(true),
				goapidoc.NewProperty("cover_url", "string", true, "封面"),
				goapidoc.NewProperty("video_url", "string", true, "资源"),
			),
	)
}

type InsertVideoParam struct {
	Title       string  `form:"title"       json:"title"       binding:"required,min=1,max=100"`
	Description *string `form:"description" json:"description" binding:"required,min=0,max=255"`
	CoverUrl    string  `form:"cover_url"   json:"cover_url"   binding:"required,url"`
	VideoUrl    string  `form:"video_url"   json:"video_url"   binding:"required,url"`
}

type UpdateVideoParam struct {
	Title       string  `form:"title"       json:"title"       binding:"required,min=1,max=100"`
	Description *string `form:"description" json:"description" binding:"required,min=0,max=255"`
	CoverUrl    string  `form:"cover_url"   json:"cover_url"   binding:"required,url"`
	VideoUrl    string  `form:"video_url"   json:"video_url"   binding:"required,url"`
}

func (i *InsertVideoParam) ToPo() *po.Video {
	return &po.Video{
		Title:       i.Title,
		Description: *i.Description,
		VideoUrl:    i.VideoUrl,
		CoverUrl:    i.CoverUrl,
	}
}

func (u *UpdateVideoParam) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":       u.Title,
		"description": *u.Description,
		"cover_url":   u.CoverUrl,
		"video_url":   u.VideoUrl,
	}
}
