package database

import (
	"database/sql"
	"oneshop/utils"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {

	DB, err := sql.Open("mysql", "root:Oz444yvyx88@tcp(127.0.0.1:3306)/LAA0989476oneshop")
	utils.CheckErr(err)

	return DB
}

func Insert(sqlstring string, data []interface{}) int {

	DB := GetDB()

	stmt, err := DB.Prepare(sqlstring)
	utils.CheckErr(err)

	res, err := stmt.Exec(data...)
	utils.CheckErr(err)

	id, err := res.LastInsertId()
	utils.CheckErr(err)

	defer DB.Close()

	return int(id)
}

func Update(sqlstring string, data []interface{}) int {

	DB := GetDB()

	stmt, err := DB.Prepare(sqlstring)
	utils.CheckErr(err)

	res, err := stmt.Exec(data...)
	utils.CheckErr(err)

	id, err := res.LastInsertId()
	utils.CheckErr(err)

	defer DB.Close()

	return int(id)
}

func Delete(sqlstring string, data []interface{}) int {

	DB := GetDB()

	stmt, err := DB.Prepare(sqlstring)
	utils.CheckErr(err)

	res, err := stmt.Exec(data...)
	utils.CheckErr(err)

	id, err := res.LastInsertId()
	utils.CheckErr(err)

	defer DB.Close()

	return int(id)
}

func Query(sqlstring string, data []interface{}) *sql.Rows {

	DB := GetDB()

	rows, err := DB.Query(sqlstring, data...)
	utils.CheckErr(err)

	defer DB.Close()

	return rows
}
