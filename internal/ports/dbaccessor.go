package ports

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type dbaccessor interface {
	Connect(database string, user string, password string) (*sql.DB, error)
}

func Connect(database string, user string, password string) (*sql.DB, error) {
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	defer con.Close()
	return con, err
}
