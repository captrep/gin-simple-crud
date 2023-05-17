package mysql

import (
	"database/sql"
	"fmt"

	"github.com/captrep/gin-simple-crud/config"
	_ "github.com/go-sql-driver/mysql"
)

var Mysql *sql.DB

func init() {
	Mysql, err := sql.Open(config.Conf.DBDriver, config.Conf.DBSource)
	if err != nil {
		panic(err)
	}
	errPing := Mysql.Ping()
	if errPing != nil {
		panic(errPing)
	}

	fmt.Println("Database succesfully configured!")
}
