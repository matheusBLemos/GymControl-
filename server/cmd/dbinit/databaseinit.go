package trainingsystem

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	db         *sql.DB
)