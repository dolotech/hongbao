package data

import (
	"strconv"
	"sync"
	"utils/db"
)

type HonbaoTb struct {
	Id       uint   `json:"Id" gorm:"primary_key:Id"`
	Uid      string `json:"uid" gorm:"type:varchar(200);column:uid"`
	Nickname string `json:"nickname" gorm:"type:varchar(200);column:nickname"`
	Userimg  string `json:"userimg" gorm:"type:varchar(200);column:userimg"`
	Type     string `json:"type" gorm:"type:varchar(200);column:type"`
	Money    string `json:"money" gorm:"type:varchar(200);column:money"`
	Moneys   string `json:"moneys" gorm:"type:varchar(200);column:moneys"`
	Lei      int    `json:"lei" gorm:"type:int(11);column:lei"`
	Number   string `json:"number" gorm:"varchar(200);column:number"`
	Fatime   string `json:"fatime" gorm:"type:varchar(200);column:fatime"`
}

type HonbaoTbs []HonbaoTb

var syc2 sync.Map
func (this *HonbaoTbs) Save() error {
	for _, v := range *this {
		if v.Number == "7" {
			if _, ok := syc2.Load(v.Id); !ok {
				syc2.Store(v.Id, struct{}{})
				if !v.Exist() {
					v.Save()
				}
			}
		}
	}
	return nil
}

func (this *HonbaoTb) Exist() bool {
	return db.Get().Model(this).Where("Id=?", this.Id).Limit(1).Find(&HonbaoTb{}).Error == nil
}

func (this *HonbaoTb) Delete() error {
	return db.Get().Model(this).Delete(this).Error
}

func (this *HonbaoTb) Update() error {
	return db.Get().Model(this).Update(this).Error
}

func (this *HonbaoTb) Save() error {
	return db.Get().Model(this).Save(this).Error
}

func (this HonbaoTb) GetMoney() float64 {
	i, _ := strconv.ParseFloat(this.Money, 10)
	return i
}
