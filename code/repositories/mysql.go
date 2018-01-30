package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/xz_local?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id,goods_id,name FROM cn_hcbuy_goods limit 2")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var goods_id int
		var name string
		err := rows.Scan(&id, &goods_id, &name)
		if  err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ID为：%d，商品ID为：%d，商品名称为：%s\n", id, goods_id, name)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

}
