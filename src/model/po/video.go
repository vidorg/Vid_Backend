package po

import (
	"github.com/Aoi-hosizora/ahlib-web/xgorm"
)

type Video struct {
	Vid         uint64 `gorm:"primary_key; auto_increment"`                      // video id
	Title       string `gorm:"type:varchar(255);  not_null"`                     // video title
	Description string `gorm:"type:varchar(1023); not_null"`                     // video description
	VideoUrl    string `gorm:"type:varchar(255);  not_null"`                     // video source url (oss)
	CoverUrl    string `gorm:"type:varchar(255);  not_null"`                     // video cover url (oss)
	AuthorUid   uint64 `gorm:"not_null"`                                         // video author id
	Author      *User  `gorm:"foreignkey:AuthorUid; association_foreignkey:Uid"` // po.Video belongs to po.User (has many)

	xgorm.GormTime
}
