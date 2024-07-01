package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

type DTOAPIPostPhotoStudioMembers struct {
	Member         *entity.PhotoStudioMember
	SentInvitation bool
}

func (t *Impl) APIPostPhotoStudioMembers(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	name string,
) (*DTOAPIPostPhotoStudioMembers, error) {
	member, _, err := t.S.CreatePhotoStudioMember(ctx, photoStudioID, email, name)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	sentInvitation := true
	if err := t.S.SendPhotoStudioMemberInvitation(ctx, member.ID); err != nil {
		t.L.Error("SendInvitation is failed", "err", err)
		sentInvitation = false
	}
	return &DTOAPIPostPhotoStudioMembers{
		Member:         member,
		SentInvitation: sentInvitation,
	}, nil
}