package model

import (
	"oneshop/database"
	table "oneshop/internal/table"
)

// func Select_User_Id(data []interface{}) []map[string]interface{} {
// 	row := db.Query("SELECT user_id FROM users WHERE account = ? AND password = ?", data)
// 	return row
// }

// func Update_User_Password(data []interface{}) int {
// 	id := db.Update("UPDATE users SET password = ? where user_id = ?", data)
// 	return id
// }

// func Insert_user(data []interface{}) int {
// 	id := db.Insert("insert into user (account, password, last_login_ip, login_count) values (?, ?, ?, ?)", data)
// 	return id
// }

// func Insert_login_log(data []interface{}) int {
// 	id := db.Insert("insert into login_log (user_id, account, ip) values (?, ?, ?)", data)
// 	return id
// }

func User_Get_Shop_List(data []interface{}) []table.Shop_detail {
	sql := `SELECT shop.shop_id, IFNULL(shop_name, ''), IFNULL(shop_info, ''), IFNULL(shop_image, ''),
	 IFNULL(corporation_name, ''), IFNULL(post_code, ''), IFNULL(shop_location, ''), 
	 IFNULL(shop_city, ''), IFNULL(open_time, ''), IFNULL(dayoff, ''),
	 IFNULL(phonenumber, ''), IFNULL(email, '') FROM shop LEFT JOIN shop_detail 
	 ON shop.shop_id = shop_detail.shop_id WHERE status = 1`
	rows := database.Query(sql, data)
	result := []table.Shop_detail{}
	for rows.Next() {
		var row table.Shop_detail
		err := rows.Scan(&row.ShopId, &row.ShopName, &row.ShopInfo, &row.ShopImage,
			&row.CorporationName, &row.PostCode, &row.ShopLocation, &row.ShopCity, &row.OpenTime,
			&row.DayOff, &row.PhoneNumber, &row.Email)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

func User_Get_Shop(data []interface{}) []table.Shop_detail {
	sql := `SELECT shop.shop_id, IFNULL(shop_name, ''), IFNULL(shop_info, ''), IFNULL(shop_image, ''),
	 IFNULL(corporation_name, ''), IFNULL(post_code, ''), IFNULL(shop_location, ''), 
	 IFNULL(shop_city, ''), IFNULL(open_time, ''), IFNULL(dayoff, ''),
	 IFNULL(phonenumber, ''), IFNULL(email, '') FROM shop LEFT JOIN shop_detail 
	 ON shop.shop_id = shop_detail.shop_id WHERE shop.shop_id = ? AND status = 1`
	rows := database.Query(sql, data)
	result := []table.Shop_detail{}
	for rows.Next() {
		var row table.Shop_detail
		err := rows.Scan(&row.ShopId, &row.ShopName, &row.ShopInfo, &row.ShopImage,
			&row.CorporationName, &row.PostCode, &row.ShopLocation, &row.ShopCity, &row.OpenTime,
			&row.DayOff, &row.PhoneNumber, &row.Email)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

func User_Get_Shop_Car_List(data []interface{}) []table.Car {
	sql := `SELECT car_id, IFNULL(car_name, ''), IFNULL(car_brand, ''), IFNULL(car_image, ''),
	 IFNULL(car_price, ''), IFNULL(car_fee, ''), IFNULL(car_year, ''), IFNULL(shelves, '')
	 FROM car WHERE shop_id = ? AND status = 1`
	rows := database.Query(sql, data)
	result := []table.Car{}
	for rows.Next() {
		var row table.Car
		err := rows.Scan(&row.CarId, &row.CarName, &row.CarBrand,
			&row.CarImage, &row.CarPrice, &row.CarFee, &row.CarYear, &row.Shelves)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}
