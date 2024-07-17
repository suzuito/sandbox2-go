package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type SetLineLinkInfoArgument struct {
	MessagingAPIChannelSecret *string `json:"messagingApiChannelSecret"`
	LongAccessToken           *string `json:"longAccessToken"`
}

type Repository interface {
	GetLineLinkInfo(ctx context.Context, photoStudioID common_entity.PhotoStudioID) (*entity.LineLinkInfo, error)
	CreateLineLinkInfo(ctx context.Context, info *entity.LineLinkInfo) (*entity.LineLinkInfo, error)
	DeleteLineLinkInfo(ctx context.Context, photoStudioID common_entity.PhotoStudioID) error
	SetLineLinkInfoActive(ctx context.Context, photoStudioID common_entity.PhotoStudioID, active bool) (*entity.LineLinkInfo, error)
	SetLineLinkInfo(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		arg *SetLineLinkInfoArgument,
	) (*entity.LineLinkInfo, error)
}