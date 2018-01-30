package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CnHcbuyGoods struct {
	gorm.Model
	Id				uint
	Goods_id		int
	Name			string  `gorm:"size:255"`
	Price			float32
	Market_price	float32
	Mall_price		float32
	Goods_num		int
	Img_path		string
	Original_img	string
	Goods_thumb_12	string
	Goods_thumb_20	string
	State			int
	Brand			string
	Brand_logo		string
	Weight			float32
	Unit			string
	Category_id		int
	Introduction	string
	Param			string
	Is_best			int
	Is_hot			int
	Is_new			int
	Amount			int
	Product_area	string
	Created_at		time.Time
	Updated_at		time.Time
	Deleted_at		time.Time
}

func (CnHcbuyGoods) TableNmae() string {
	return "cn_hcbuy_goods"
}
