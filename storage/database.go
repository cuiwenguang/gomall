package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gomall/pkg/settings"
	"gomall/pkg/web"
	"log"
	"strings"
	"time"
)

var dbs map[string]*gorm.DB

func initDB() {
	var err error
	dbs = make(map[string]*gorm.DB)
	for k, v := range settings.AppConfig.Database.Connections {
		dbs[k], err = gorm.Open(settings.AppConfig.Database.Type, v)
		if err != nil {
			log.Fatalf("models setup fail : %v", err)
		}

		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			if strings.Contains(defaultTableName, "vo") {
				return "v_" + defaultTableName
			} else {
				return "t_" + defaultTableName
			}
		}

		dbs[k].SingularTable(true)
		dbs[k].LogMode(true)
		// 使用自带的Base Model 一下方法不再需要
		//dbs[k].Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		//dbs[k].Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		dbs[k].DB().SetMaxIdleConns(10)
		dbs[k].DB().SetMaxOpenConns(100)

		// 数据迁移工具不成熟，不建议使用，直接手工建表
		// dbs[k].AutoMigrate(&Account{},&Company{})
	}
}

func CloseDB(domain string) {
	defer dbs[domain].Close()
}

func GetDB(req *web.RequestContext) *gorm.DB {
	id := GetDomain(req.Host)
	return dbs[id]
}

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

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateAt", time.Now())
	}
}
