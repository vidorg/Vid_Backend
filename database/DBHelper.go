package database

import (
	"fmt"
	"log"
	"vid/config"
	"vid/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func SetupDBConn(cfg config.Config) {
	dbParams := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbName,
	)
	var err error
	DB, err = gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatal(2, err)
	}

	DB.LogMode(true)
	DB.SingularTable(true) // 复数表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	} // 表名前缀

	DB.AutoMigrate(&models.User{})

	DB.AutoMigrate(&models.PassRecord{})
	DB.Model(&models.PassRecord{}).AddForeignKey("uid", "tbl_user(uid)", "cascade", "cascade")

	DB.AutoMigrate(&models.Video{})
	DB.Model(&models.Video{}).AddForeignKey("author_uid", "tbl_user(uid)", "set null", "cascade")

	DB.AutoMigrate(&models.Playlist{})
	DB.Model(&models.Playlist{}).AddForeignKey("author_uid", "tbl_user(uid)", "set null", "cascade")

	DB.AutoMigrate(&models.Videoinlist{})
	DB.Model(&models.Videoinlist{}).AddForeignKey("gid", "tbl_playlist(gid)", "cascade", "cascade")
	// set null
	DB.Model(&models.Videoinlist{}).AddForeignKey("vid", "tbl_video(vid)", "cascade", "cascade")
}
