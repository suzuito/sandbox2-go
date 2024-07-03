package repository

import "database/sql"

type Impl struct {
	Pool *sql.DB
}
