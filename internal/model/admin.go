package model

import (
	"oneshop/database"
	table "oneshop/internal/table"
)

func Select_Admin_Id(data []interface{}) []table.Admin {
	sql := `SELECT admin_id FROM admin WHERE account = ? AND password = ?`
	rows := database.Query(sql, data)
	result := []table.Admin{}
	for rows.Next() {
		var row table.Admin
		err := rows.Scan(&row.AdminId)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

func Insert_Admin_LoginLog(data []interface{}) int {
	sql := `INSERT INTO admin_login_log (admin_id, ip) VALUES (?, ?);`
	id := database.Insert(sql, data)
	return id
}

func Update_Shop_Status(data []interface{}) int {
	sql := `UPDATE shop SET status = ? WHERE shop_id = ?`
	id := database.Update(sql, data)
	return id
}

func Select_Shop_List(data []interface{}) []table.Shop {
	sql := `SELECT shop_id, account, status FROM shop WHERE status = "1, 2, 3, 4"`
	rows := database.Query(sql, data)
	result := []table.Shop{}
	for rows.Next() {
		var row table.Shop
		err := rows.Scan(&row.ShopId, &row.Account, &row.Status)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}
