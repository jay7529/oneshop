package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"oneshop/utils"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

func connectWithConnector() (*sql.DB, error) {
	var (
		dbUser                 = "J"
		dbPwd                  = "j75297529"
		dbName                 = "LAA0989476oneshop"
		instanceConnectionName = "oneshop-418410:asia-northeast1:oneshop"
		usePrivate             = ""
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true", dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}

func Insert(sqlstring string, data []interface{}) int {

	DB, err := connectWithConnector()
	utils.CheckErr(err)

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

	DB, err := connectWithConnector()
	utils.CheckErr(err)

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

	DB, err := connectWithConnector()
	utils.CheckErr(err)

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

	DB, err := connectWithConnector()
	utils.CheckErr(err)

	rows, err := DB.Query(sqlstring, data...)
	utils.CheckErr(err)

	defer DB.Close()

	return rows
}
