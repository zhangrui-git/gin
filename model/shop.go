package model

import (
	"by/user/bean/shop"
	"by/user/container"
	"gorm.io/gorm"
)

type Shop struct {
	ShopId   uint   `json:"shopId" gorm:"primaryKey"`
	ShopName string `json:"shopName"`
}

func (*Shop) TableName() string {
	return "shop"
}

func (s *Shop) Create() *gorm.DB {
	db := container.Database()
	return db.Create(s)
}

func (s *Shop) Delete() *gorm.DB {
	db := container.Database()
	return db.Delete(s, s.ShopId)
}

func (s *Shop) Update(fields []string) *gorm.DB {
	db := container.Database()
	return db.Select(fields).Updates(s)
}

func (s *Shop) GetBySid(sid shop.Pk) *gorm.DB {
	db := container.Database()
	return db.First(s, sid.Val)
}

type Shops []Shop

func (ss *Shops) GetListBySid(sid shop.PkSet) *gorm.DB {
	db := container.Database()
	return db.Find(ss, sid.Val)
}
