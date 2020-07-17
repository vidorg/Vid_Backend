package database

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xgorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/vidorg/vid_backend/src/config"
	"github.com/vidorg/vid_backend/src/model/po"
	"log"
)

func SetupMySQLConn(cfg *config.MySQLConfig, logger *logrus.Logger) *GormHelper {
	dbParams := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User, cfg.Password,
		cfg.Host, cfg.Port,
		cfg.Name, cfg.Charset,
	)
	db, err := gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatalln("Failed to connect mysql:", err)
	}

	db.LogMode(cfg.IsLog)
	db.SetLogger(xgorm.NewGormLogrus(logger))

	xgorm.HookDeleteAtField(db, xgorm.DefaultDeleteAtTimeStamp)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	}

	autoMigrateModel(db)
	addFullTextIndex(db, cfg)

	return NewGormHelper(db)
}

func autoMigrateModel(db *gorm.DB) {
	autoMigrate := func(value interface{}) {
		rdb := db.AutoMigrate(value)
		if rdb.Error != nil {
			log.Fatalln(rdb.Error)
		}
	}

	autoMigrate(&po.User{})
	autoMigrate(&po.Account{})
	autoMigrate(&po.Video{})
}

func addFullTextIndex(db *gorm.DB, cfg *config.MySQLConfig) {
	// TODO
	// checkExecIndex := func(tblName string, idxName string, param string) {
	// 	cnt := 0
	// 	rdb := db.Table("INFORMATION_SCHEMA.STATISTICS").Where("TABLE_SCHEMA = ? AND TABLE_NAME = ? AND INDEX_NAME = ?", cfg.Name, tblName, idxName).Count(&cnt)
	// 	if rdb.Error != nil {
	// 		log.Fatalln(rdb.Error)
	// 	}
	// 	if cnt == 0 {
	// 		sql := fmt.Sprintf("CREATE FULLTEXT INDEX `%s` ON `%s` (%s) WITH PARSER `ngram`", idxName, tblName, param)
	// 		rdb := db.Exec(sql)
	// 		if rdb.Error != nil {
	// 			log.Fatalln(rdb.Error)
	// 		}
	// 	}
	// }
	//
	// checkExecIndex("tbl_user", "idx_username_profile_fulltext", "`username`(100), `profile`(20)")
	// checkExecIndex("tbl_video", "idx_title_description_fulltext", "`title`(100), `description`(40)")
}