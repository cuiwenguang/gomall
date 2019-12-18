package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gomall/pkg/settings"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID       int `gorm:"primary_key"`
	CreateAt time.Time
	UpdateAt time.Time
}

func Setup() {
	var err error
	db, err = gorm.Open(settings.AppConfig.Database.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			settings.AppConfig.Database.User,
			settings.AppConfig.Database.Password,
			settings.AppConfig.Database.Host,
			settings.AppConfig.Database.Name))
	if err != nil {
		log.Fatalf("models setup fail : %v", err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&Account{}, &Profile{})
}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `updateat`, `updateat` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreateAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdateAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdateAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateAt", time.Now())
	}
}
