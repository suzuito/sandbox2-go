package mail

import (
	"context"
	"net/url"
)

type UserMailSender interface {
	SendUserCreationMail(
		ctx context.Context,
		email string,
		verifierURL *url.URL,
	) error
}
