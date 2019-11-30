package data

import (
	"strconv"
	"sync"
	"utils/db"
)

type HonbaosTb struct {
	Id       uint   `json:"Id"       gorm:"primary_key:Id"`
	Uid      string `json:"uid"      gorm:"type:varchar(200);column:uid"`
	Nickname string `json:"nickname" gorm:"type:varchar(200);column:nickname"`
	Userimg  string `json:"userimg"   gorm:"type:varchar(200);column:userimg"`
	Money    string `json:"money"     gorm:"type:varchar(200);column:money"`
	HbId     uint   `json:"hbid"      gorm:"type:int(11);column:hbid"`
	Lqtime   string `json:"lqtime"    gorm:"type:varchar(200);column:lqtime"`
	Zlei     int    `json:"zlei"    gorm:"type:int(11);column:zlei"`
}

type HonbaosTbs []HonbaosTb

var syc1 sync.Map

func (this *HonbaosTbs) Save() error {
	for _, v := range *this {
		uni:=v.Uid + strconv.Itoa(int(v.HbId))
		if _, ok := syc1.Load(uni); !ok {
			syc1.Store(uni, struct{}{})
			if !v.Exist() {
				v.Save()
			}
		}
	}
	return nil
}

func (this *HonbaosTb) Exist() bool {
	return db.Get().Model(this).Where("Id=?", this.Id).Limit(1).Find(&HonbaosTb{}).Error == nil
}

func (this *HonbaosTb) Delete() error {
	return db.Get().Model(this).Delete(this).Error
}

func (this *HonbaosTb) Update() error {
	return db.Get().Model(this).Update(this).Error
}

func (this *HonbaosTb) Save() error {
	return db.Get().Model(this).Save(this).Error
}

func (this HonbaosTb) GetMoney() float64 {
	i, _ := strconv.ParseFloat(this.Money, 10)
	return i
}
