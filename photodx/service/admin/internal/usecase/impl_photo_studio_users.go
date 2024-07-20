package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

type DTOPhotoStudioUser struct {
	User            *common_entity.User     `json:"user"`
	PhotoStudioUser *entity.PhotoStudioUser `json:"photoStudioUser"`
}

type DTOAPIGetPhotoStudioUsers struct {
	Results    []*DTOPhotoStudioUser `json:"results"`
	HasNext    bool                  `json:"hasNext"`
	HasPrev    bool                  `json:"hasPrev"`
	NextOffset int                   `json:"nextOffset"`
	PrevOffset int                   `json:"prevOffset"`
}

func (t *Impl) APIGetPhotoStudioUsers(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
	offset int,
) (*DTOAPIGetPhotoStudioUsers, error) {
	if offset < 0 {
		offset = 0
	}
	q := cgorm.ListQuery{
		Offset: offset,
		Limit:  30,
		SortColumns: []cgorm.SortColumn{
			{
				Name: "updated_at",
				Type: cgorm.Desc,
			},
		},
	}
	users, hasNext, err := t.BusinessLogic.GetPhotoStudioUsers(
		ctx,
		principal.GetPhotoStudioID(),
		&q,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	dtoUsers, err := t.attachCommonUsers(ctx, users)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetPhotoStudioUsers{
		Results:    dtoUsers,
		HasNext:    hasNext,
		HasPrev:    q.HasPrev(),
		NextOffset: q.NextOffset(),
		PrevOffset: q.PrevOffset(),
	}, nil
}

func (t *Impl) attachCommonUsers(
	ctx context.Context,
	photoStudioUsers []*entity.PhotoStudioUser,
) ([]*DTOPhotoStudioUser, error) {
	userIDs := arrayutil.Map(
		photoStudioUsers,
		func(v *entity.PhotoStudioUser) common_entity.UserID {
			return v.UserID
		},
	)
	userIDs = arrayutil.Filter(
		userIDs,
		func(v common_entity.UserID) bool {
			return v != common_entity.UserID("")
		},
	)
	commonUsers, err := t.AuthUserBusinessLogic.GetUsers(
		ctx,
		userIDs,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	commonUsersMap := arrayutil.ListToMap(
		commonUsers,
		func(e *common_entity.User) common_entity.UserID {
			return e.ID
		},
	)
	ret := []*DTOPhotoStudioUser{}
	for _, photoStudioUser := range photoStudioUsers {
		commonUsers, exists := commonUsersMap[photoStudioUser.UserID]
		if !exists || len(commonUsers) <= 0 {
			continue
		}
		ret = append(ret, &DTOPhotoStudioUser{
			User:            commonUsers[0],
			PhotoStudioUser: photoStudioUser,
		})
	}
	return ret, nil
}

func (t *Impl) APIGetPhotoStudioUser(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
	userID common_entity.UserID,
) (*DTOPhotoStudioUser, error) {
	photoStudioUser, err := t.BusinessLogic.GetPhotoStudioUser(ctx, principal.GetPhotoStudioID(), userID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	dtos, err := t.attachCommonUsers(ctx, []*entity.PhotoStudioUser{photoStudioUser})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if len(dtos) <= 0 {
		return nil, &common_repository.NoEntryError{
			EntryType: "PhotoStudioUser",
			EntryID:   string(principal.GetPhotoStudioID()) + "-" + string(userID),
		}
	}
	return dtos[0], nil
}
