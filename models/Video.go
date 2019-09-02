package models

import "time"

type Video struct {
	Vid         int       `json:"vid" gorm:"primary_key;AUTO_INCREMENT"`
	Title       string    `json:"title" gorm:"type:varchar(80);not_null"`
	Description string    `json:"description" gorm:"type:varchar(250)"`
	VideoUrl    string    `json:"video_url"`
	UploadTime  time.Time `json:"upload_time"`
	AuthorUid   int       `json:"-"`
	Author      *User     `json:"author,omitempty" gorm:"-"`
}

// @override
func (v *Video) CheckValid() bool {
	return v.Vid > 0
}

func (v *Video) Equals(obj *Video) bool {
	return v.Vid == obj.Vid && v.Title == obj.Title && v.Description == obj.Description &&
		v.VideoUrl == obj.VideoUrl && v.AuthorUid == obj.AuthorUid
}