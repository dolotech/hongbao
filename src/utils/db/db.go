package db

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
)

func Get() *gorm.DB {
	if db == nil {
		glog.Fatalln("数据库未连接")
	}
	return db
}

type MysqlConnectiPool struct {
}

var instance *MysqlConnectiPool
var once sync.Once

var db *gorm.DB
var err_db error

func InitMysql(user, password, host, dbname string) {
	once.Do(func() {
		instance = &MysqlConnectiPool{}
		b := instance.initPool(user, password, host, dbname)
		if b {
			glog.Info("数据库连接成功")
		} else {
			glog.Info("数据库连接失败")
		}
	})
}

/*
* @fuc 初始化数据库连接(可在mail()适当位置调用)
 */
func (m *MysqlConnectiPool) initPool(user, password, host, dbname string) (issucc bool) {
	db, err_db = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", user, password, host, dbname))
	if err_db != nil {
		glog.Error(err_db)
		return false
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(time.Hour)
	//db.Callback().Create().Replace("gorm:create_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	return true
}
func (m *MysqlConnectiPool) GetMysqlDB() (db_con *gorm.DB) {
	return db
}

// // 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	glog.Error("create_time")
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("create_time"); ok {
			glog.Error("create_time")
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("update_time"); ok {
			glog.Error("update_time")
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_time"); !ok {

		glog.Error("update_time")
		scope.SetColumn("update_time", time.Now().Unix())
	}
}

// 注册删除钩子在删除之前
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedTime")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
