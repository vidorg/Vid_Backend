package po

import (
	"github.com/vidorg/vid_backend/lib/xgorm"
)

type Video struct {
	Vid         uint64 `gorm:"                    not null; primaryKey; autoIncrement"` // video id
	Title       string `gorm:"type:varchar(255);  not null"`                            // video title
	Description string `gorm:"type:varchar(1023); not null"`                            // video description
	VideoUrl    string `gorm:"type:varchar(255);  not null"`                            // video source url
	CoverUrl    string `gorm:"type:varchar(255);  not null"`                            // video cover url (oss)
	AuthorUid   uint64 `gorm:"type:bigint;        not null"`                            // video author id

	Author   *User   `gorm:"                        foreignKey:AuthorUid; references:Uid"`                                         // po.Video belongs to po.User
	Favoreds []*User `gorm:"many2many:tbl_favorite; foreignKey:Vid;       references:Uid; joinForeignKey:vid; JoinReferences:uid"` // tbl_favorite

	xgorm.Model
}
