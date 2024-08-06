package dbinit

import (
	"GusLem/gymControll/configs"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Dbinit  *sql.DB
	err error
)

func init() {
	server := configs.Config.DBServer
	user := configs.Config.DBUser
	password := configs.Config.DBPassword
	dbname := configs.Config.DBDataBase
	ConnectionString := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable host=%v", user, password, dbname, server)
	fmt.Printf("user=%v password=%v dbname=%v sslmode=disable host=%v", user, password, dbname, server)
	Dbinit, err = sql.Open("postgres", ConnectionString)
	if err != nil {
		errors.New("Error connecting to database, " + err.Error())
	}
}
