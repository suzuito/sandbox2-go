package inject

import (
	"database/sql"

	internal_repository "github.com/suzuito/sandbox2-go/photodx/service/common/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func NewRepository(pool *sql.DB) repository.Repository {
	return &internal_repository.Impl{
		Pool: pool,
	}
}
