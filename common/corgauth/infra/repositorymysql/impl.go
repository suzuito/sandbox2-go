package repositorymysql

import "database/sql"

type Impl struct {
	Pool *sql.DB
}
