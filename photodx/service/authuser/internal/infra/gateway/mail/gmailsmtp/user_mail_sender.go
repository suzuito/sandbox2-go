package gmailsmtp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type UserMailSender struct {
	Host     string
	Port     int
	Account  string
	Password string

	FromName  string
	FromEmail string
}

func (t *UserMailSender) sendUserMail(
	ctx context.Context,
	user *common_entity.User,
	subject string,
	body string,
) error {
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", t.Host, t.Port))
	if err != nil {
		return terrors.Wrap(err)
	}
	defer c.Close()
	if err := c.StartTLS(&tls.Config{
		ServerName: t.Host,
	}); err != nil {
		return terrors.Wrap(err)
	}
	auth := smtp.PlainAuth("", t.Account, t.Password, t.Host)
	if err := c.Auth(auth); err != nil {
		return terrors.Wrap(err)
	}

	// 送信元
	if err := c.Mail(t.FromEmail); err != nil {
		return terrors.Wrap(err)
	}
	// 送信先
	if err := c.Rcpt(user.Email); err != nil {
		return terrors.Wrap(err)
	}
	// if err := c.Rcpt("hoge@example.com"); err != nil {
	// 	return terrors.Wrap(err)
	// }
	// if err := c.Rcpt("fuga@example.com"); err != nil {
	// 	return terrors.Wrap(err)
	// }

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return terrors.Wrap(err)
	}
	defer wc.Close()
	messageBody := ""
	messageBody += fmt.Sprintf("To: %s<%s>\r\n", user.Name, user.Email)
	// messageBody += "Cc: ウルトラマン1000号<hoge@example.com>\r\n"
	// messageBody += "Bcc: ウルトラマン2000号<fuga@example.com>\r\n"
	messageBody += fmt.Sprintf("From: %s<replaced by google smtp server>\r\n", t.FromName)
	messageBody += fmt.Sprintf("Subject: %s\r\n", subject)
	messageBody += body
	// messageBody += "こんにちは\r\n私はウルトラマンである\r\nhttps://www.example.com\r\n"
	messageBody += "\r\n"
	if _, err := wc.Write([]byte(messageBody)); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *UserMailSender) SendUserCreationMail(
	ctx context.Context,
	user *common_entity.User,
	verifierURL *url.URL,
) error {
	return t.sendUserMail(
		ctx,
		user,
		"ユーザー登録確認メール",
		fmt.Sprintf(
			"メールアドレス %s を持つユーザーが仮登録されました。本登録する場合は下記リンクへアクセスしてください。\r\n%s\r\n本メールが身に覚えのないものである場合は無視してください。\r\n",
			user.Email,
			verifierURL.String(),
		),
	)
}
