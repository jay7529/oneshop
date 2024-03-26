package model

import (
	table "oneshop/internal/table"
	"oneshop/tools"
)

// 用帳號查詢admin_id
func Select_Admin_Id(data []interface{}) []table.Admin {
	sql := `SELECT admin_id FROM admin WHERE account = ? AND password = ?`
	rows := tools.Query(sql, data)
	result := []table.Admin{}
	for rows.Next() {
		var row table.Admin
		err := rows.Scan(&row.Admin_id)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

// 查詢admin detail
func Select_Admin_Detail(data []interface{}) []table.Admin_detail {
	sql := `SELECT shop_name, shop_info, shop_image, corporation_name,`
	sql += ` shop_location, open_time, dayoff, phonenumber, email`
	sql += ` FROM admin_detail WHERE admin_id = ?`
	rows := tools.Query(sql, data)
	result := []table.Admin_detail{}
	for rows.Next() {
		var row table.Admin_detail
		err := rows.Scan(&row.ShopName, &row.ShopInfo, &row.ShopImage,
			&row.CorporationName, &row.ShopLocation, &row.OpenTime,
			&row.DayOff, &row.PhoneNumber, &row.Email)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

// 更新admin detail
func Update_Admin_Detail(data []interface{}) []table.Admin {
	sql := `UPDATE admin_detail SET`
	sql += ` shop_name = ?, shop_info = ?, shop_image = ?, corporation_name = ?,`
	sql += ` shop_location = ?, open_time = ?, dayoff = ?, phonenumber = ?, email = ?`
	sql += ` WHERE admin_id = ?`
	tools.Update(sql, data)
	result := []table.Admin{}
	return result
}
