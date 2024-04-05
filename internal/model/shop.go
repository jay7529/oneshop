package model

import (
	"oneshop/database"
	table "oneshop/internal/table"
)

func Select_Shop_Id(data []interface{}) []table.Shop {
	sql := `SELECT shop_id FROM shop WHERE account = ? AND password = ?`
	rows := database.Query(sql, data)
	result := []table.Shop{}
	for rows.Next() {
		var row table.Shop
		err := rows.Scan(&row.ShopId)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

func Insert_Shop_LoginLog(data []interface{}) int {
	sql := `INSERT INTO shop_login_log (shop_id, ip) VALUES (?, ?)`
	id := database.Insert(sql, data)
	return id
}

func Insert_Shop(data []interface{}) int {
	sql := `INSERT INTO shop (account, password) 
	SELECT ?, ? WHERE NOT EXISTS (SELECT account FROM shop WHERE account = ?)`
	id := database.Insert(sql, data)
	return id
}

func Update_Shop_Detail_FirstTime(data []interface{}) int {
	sql := `UPDATE shop_detail SET shop_name = ?, post_code = ?, shop_location = ?,
	phonenumber = ?, email = ? WHERE shop_id = ?`
	id := database.Update(sql, data)
	return id
}

func Select_Shop_Password(data []interface{}) []table.Shop {
	sql := `SELECT password FROM shop WHERE account = ?`
	rows := database.Query(sql, data)
	result := []table.Shop{}
	for rows.Next() {
		var row table.Shop
		err := rows.Scan(&row.Password)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

func Reset_Shop_Password(data []interface{}) int {
	sql := `UPDATE shop SET password = ? WHERE account = ? AND password = ?`
	id := database.Update(sql, data)
	return id
}

func Update_Shop_Password(data []interface{}) int {
	sql := `UPDATE shop SET password = ? WHERE shop_id = ? AND password = ?`
	id := database.Update(sql, data)
	return id
}

func Get_Shop_Detail(data []interface{}) []table.Shop_detail {
	sql := `SELECT shop_id, IFNULL(shop_name, ''), IFNULL(shop_info, ''), IFNULL(shop_image, ''),
	 IFNULL(corporation_name, ''), IFNULL(post_code, ''), IFNULL(shop_location, ''), IFNULL(shop_city, ''), IFNULL(open_time, ''), IFNULL(dayoff, ''),
	 IFNULL(phonenumber, ''), IFNULL(email, '') FROM shop_detail WHERE shop_id = ?`
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

func Update_Shop_Detail(data []interface{}) int {
	sql := `UPDATE shop_detail SET shop_name = ?, shop_info = ?, shop_image = ?, corporation_name = ?,
	 post_code = ?, shop_location = ?, shop_city = ?, open_time = ?, 
	 dayoff = ?, phonenumber = ?, email = ? WHERE shop_id = ?`
	id := database.Update(sql, data)
	if id > 0 {
		sql = `UPDATE shop SET status = 3 WHERE shop_id = ?`
		id = database.Update(sql, data)
		return id
	}
	return id
}

func Insert_Shop_Car(data []interface{}) int {
	sql := `INSERT INTO car (shop_id, car_name, car_brand, car_image, car_price, car_fee, car_year)
	 VALUES (?, ?, ?, ?, ?, ?, ?)`
	id := database.Insert(sql, data)
	return id
}

func Update_Shop_Car(data []interface{}) int {
	sql := `UPDATE car SET car_name = ?, car_brand = ?, car_image = ?, car_price = ?,
	 car_fee = ?, car_year = ?, shelves = ? WHERE car_id = ? AND shop_id = ?`
	id := database.Update(sql, data)
	return id
}

func Delete_Shop_Car(data []interface{}) int {
	sql := `DELETE FROM car WHERE car_id = ? AND shop_id = ?`
	id := database.Delete(sql, data)
	return id
}

func Get_Shop_Car(data []interface{}) []table.Car {
	sql := `SELECT car_id, IFNULL(car_name, ''), IFNULL(car_brand, ''), IFNULL(car_image, ''),
	 IFNULL(car_price, ''), IFNULL(car_fee, ''), IFNULL(car_year, ''), IFNULL(shelves, '')
	 FROM car WHERE car_id = ? AND shop_id = ?`
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

func Get_Shop_Car_List(data []interface{}) []table.Car {
	sql := `SELECT car_id, IFNULL(car_name, ''), IFNULL(car_brand, ''), IFNULL(car_image, ''),
	 IFNULL(car_price, ''), IFNULL(car_fee, ''), IFNULL(car_year, ''), IFNULL(shelves, '')
	 FROM car WHERE shop_id = ?`
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

func Insert_Shop_Staff(data []interface{}) int {
	sql := `INSERT INTO staff (shop_id, staff_name, staff_image, staff_position, staff_introduction)
	 VALUES (?, ?, ?, ?, ?)`
	id := database.Insert(sql, data)
	return id
}

func Update_Shop_Staff(data []interface{}) int {
	sql := `UPDATE staff SET staff_name = ?, staff_image = ?, staff_position = ?, staff_introduction = ?
	 WHERE staff_id = ? AND shop_id = ?`
	id := database.Update(sql, data)
	return id
}

func Delete_Shop_Staff(data []interface{}) int {
	sql := `DELETE FROM staff WHERE staff_id = ? AND shop_id = ?`
	id := database.Delete(sql, data)
	return id
}

func Get_Shop_Staff_List(data []interface{}) []table.Staff {
	sql := `SELECT staff_id, IFNULL(staff_name, ''), IFNULL(staff_image, ''), IFNULL(staff_position, ''),
	 IFNULL(staff_introduction, '') FROM staff WHERE shop_id = ?`
	rows := database.Query(sql, data)
	result := []table.Staff{}
	for rows.Next() {
		var row table.Staff
		err := rows.Scan(&row.StaffId, &row.StaffName, &row.StaffImage,
			&row.StaffPosition, &row.StaffIntroduction)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}
