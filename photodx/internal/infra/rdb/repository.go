package rdb

import "database/sql"

type Repository struct {
	Pool *sql.DB
}
