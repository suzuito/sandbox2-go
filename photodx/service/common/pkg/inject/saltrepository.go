package inject

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/internal/infra/saltrepository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func NewSaltRepository(version string) repository.SaltRepository {
	return &saltrepository.Impl{
		Version: version,
	}
}
