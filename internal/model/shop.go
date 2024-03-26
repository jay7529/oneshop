package model

import (
	"oneshop/database"
	table "oneshop/internal/table"
)

// 用帳號密碼查詢shop_id
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

// 查詢shop detail
func Get_Shop_Detail(data []interface{}) []table.Shop_detail {
	sql := `SELECT shop_id, IFNULL(shop_name, ''), IFNULL(shop_info, ''), IFNULL(shop_image, ''),
	 IFNULL(corporation_name, ''), IFNULL(shop_location, ''), IFNULL(open_time, ''), IFNULL(dayoff, ''),
	 IFNULL(phonenumber, ''), IFNULL(email, '') FROM shop_detail WHERE shop_id = ?`
	rows := database.Query(sql, data)
	result := []table.Shop_detail{}
	for rows.Next() {
		var row table.Shop_detail
		err := rows.Scan(&row.ShopId, &row.ShopName, &row.ShopInfo, &row.ShopImage,
			&row.CorporationName, &row.ShopLocation, &row.OpenTime,
			&row.DayOff, &row.PhoneNumber, &row.Email)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return result
}

// 更新shop detail
func Update_Shop_Detail(data []interface{}) []table.Shop {
	sql := `UPDATE shop_detail SET shop_name = ?, shop_info = ?, shop_image = ?, corporation_name = ?,
	 shop_location = ?, open_time = ?, dayoff = ?, phonenumber = ?, email = ? WHERE shop_id = ?`
	database.Update(sql, data)
	result := []table.Shop{}
	return result
}

// 新增car
func Insert_Shop_Car(data []interface{}) []table.Shop {
	sql := `INSERT INTO car (shop_id, car_name, car_brand, car_image, car_price, car_fee, car_year)
	 VALUES (?, ?, ?, ?, ?, ?, ?);`
	database.Insert(sql, data)
	result := []table.Shop{}
	return result
}

// 更新car detail
func Update_Shop_Car(data []interface{}) []table.Shop {
	sql := `UPDATE car SET car_name = ?, car_brand = ?, car_image = ?, car_price = ?,
	 car_fee = ?, car_year = ?, shelves = ? WHERE car_id = ? AND shop_id = ?`
	database.Update(sql, data)
	result := []table.Shop{}
	return result
}

// 刪除car
func Delete_Shop_Car(data []interface{}) []table.Shop {
	sql := `DELETE FROM car WHERE car_id = ? AND shop_id = ?;`
	database.Delete(sql, data)
	result := []table.Shop{}
	return result
}

// 查詢car detail
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

// 查詢car list
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
