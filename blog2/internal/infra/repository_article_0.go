package infra

import (
	"database/sql"
)

type RepositoryArticle struct {
	Pool *sql.DB
}
