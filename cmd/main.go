package main

import (
	"database/sql"
	"log"

	"github.com/MrTeacheer/ecom/cmd/api"
	"github.com/MrTeacheer/ecom/common/errs"
	"github.com/MrTeacheer/ecom/config"
	"github.com/MrTeacheer/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, db_err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	errs.Check(db_err)
	checkConnDB(db)

	server := api.NewAPIServer(":8000", db)
	err := server.Run()
	errs.Check(err)
}

func checkConnDB(db *sql.DB) {
	err := db.Ping()
	errs.Check(err)
	log.Println("Database connected succesfully")
}
