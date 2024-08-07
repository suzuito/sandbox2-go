package mail

import (
	"context"
	"net/url"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type MailTo struct {
	Email string
	Name  string
}

type UserMailSender interface {
	SendUserCreationCode(
		ctx context.Context,
		req common_entity.UserCreationRequest,
		userRegisterURL *url.URL,
	) error
}
