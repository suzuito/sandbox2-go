package mail

import (
	"context"
	"net/url"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type UserMailSender interface {
	SendUserCreationMail(
		ctx context.Context,
		user *common_entity.User,
		verifierURL *url.URL,
	) error
}
