package mysql

import (
	"database/sql"
	"log"

	"github.com/captrep/gin-simple-crud/config"
	_ "github.com/go-sql-driver/mysql"
)

var Mysql *sql.DB

func init() {
	var err error
	Mysql, err = sql.Open(config.Conf.DBDriver, config.Conf.DBSource)
	if err != nil {
		log.Println(err)
	}
	errPing := Mysql.Ping()
	if errPing != nil {
		panic(errPing)
	}

	log.Println("Database succesfully configured!")
}
