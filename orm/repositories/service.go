package repositories

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var db *gorm.DB
type SqlConn interface {
	Conn()
}
func NewMysqlConn() SqlConn {
	return mysqlConn{}
}

type mysqlConn struct{}
// New returns a basic Service with all of the expected middlewares wired in.
func New() SqlConn {
	var svc SqlConn
	{
		svc = NewMysqlConn()
	}
	return svc
}
func (m mysqlConn) Conn() {
	if db == nil {
		var err error
		db, err = gorm.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/xz_local?charset=utf8&parseTime=True&loc=Local")
		db.LogMode(true)
		if err != nil {
			fmt.Println("链接数据库失败%s", err)
		}
		defer db.Close()
	}

	//CnHcbuyGoods := models.CnHcbuyGoods{}
	//row := db.Model(&CnHcbuyGoods).Where("id = ?", 1).First(&CnHcbuyGoods)
	//row.Scan(&CnHcbuyGoods.Name)
	//fmt.Println(CnHcbuyGoods.Name)
	//db.Debug().Where("id = ?", 1).First(&models.CnHcbuyGoods{})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

}
