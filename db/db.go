package db

import (
	"database/sql"

	"github.com/MrTeacheer/ecom/common/errs"
	"github.com/go-sql-driver/mysql"
)



func NewMySQLStorage(cfg mysql.Config) (*sql.DB,error){
	db,err := sql.Open("mysql",cfg.FormatDSN())
	errs.Check(err)
	return db,nil
}
