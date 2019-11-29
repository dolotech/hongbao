package data

import (
	"github.com/golang/glog"
	"testing"
	"utils/db"
)

func TestCreateTableHonbaos(t *testing.T) {
	db.InitMysql("root", "YRj9f#*o4W%e^u%W", "47.112.197.61:3306", "hongbao")

	t.Error(db.Get().CreateTable(&HonbaosTb{}))
	t.Error(db.Get().CreateTable(&HonbaoTb{}))
}

func TestHonbaosTb(t *testing.T) {
	db.InitMysql("root", "YRj9f#*o4W%e^u%W", "47.112.197.61:3306", "hongbao")

	user := HonbaosTb{

	}
	glog.Error(user.Save())
}

func TestCreateTableHonbao(t *testing.T) {
	db.InitMysql("root", "YRj9f#*o4W%e^u%W", "47.112.197.61:3306", "hongbao")

	t.Error(db.Get().CreateTable(&HonbaoTb{}))
}

func TestHonbaoTb(t *testing.T) {
	db.InitMysql("root", "YRj9f#*o4W%e^u%W", "47.112.197.61:3306", "hongbao")

	user := HonbaoTb{
		Uid:"7623c",
		Nickname:   "\u6d41\u6d6a",
		Userimg:"Userimg2",
		Type:"2",
		Money:"10.00",
		Moneys:"0.00",
		Fatime:"2019-11-28 21:58:07",
		Id:         112641,
		Lei:1,
		Number:7,

	}
	glog.Error(user.Exist())
	//glog.Error(user.Save())
}

