package usecase

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

var ErrVerifyLINEWebhookSignature = fmt.Errorf("ErrVerifyLINEWebhookSignature")

func (t *Impl) APIPostLineMessagingAPIWebhook(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	body []byte,
	xLINESignature string,
	skipVerifySignature bool,
) error {
	lineLink, err := t.BusinessLogic.GetActiveLineLink(ctx, photoStudioID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if !skipVerifySignature {
		// Verify sig
		signatureLeft, err := base64.StdEncoding.DecodeString(xLINESignature)
		if err != nil {
			t.L.Error("", "err", err)
			return terrors.Wrapf("%s : %w", err.Error(), ErrVerifyLINEWebhookSignature)
		}
		h := hmac.New(sha256.New, []byte(lineLink.MessagingAPIChannelSecret))
		h.Write(body)
		signatureRight := h.Sum(nil)
		t.L.Info(
			"Verify header",
			"matched", bytes.Equal(signatureLeft, signatureRight),
		)
		if !bytes.Equal(signatureLeft, signatureRight) {
			return terrors.Wrapf("signature is not equals : %w", ErrVerifyLINEWebhookSignature)
		}
	}
	return t.ProcLINEMessagingAPIWebhook(ctx, lineLink, body)
}

func (t *Impl) ProcLINEMessagingAPIWebhook(
	ctx context.Context,
	lineLinkInfo *entity.LineLinkInfo,
	messageBytes []byte,
) error {
	message := struct {
		Destination string `json:"destination"`
		Events      []any  `json:"events"`
	}{}
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		return terrors.Wrap(err)
	}
	for _, event := range message.Events {
		if err := t.ProcLINEMessagingAPIWebhook_anyEvent(ctx, lineLinkInfo, event); err != nil {
			t.L.Error("ProcLINEMessagingAPIWebhook_anyEvent is failed", "err", err)
			continue
		}
	}
	return nil
}

func (t *Impl) ProcLINEMessagingAPIWebhook_anyEvent(
	ctx context.Context,
	lineLinkInfo *entity.LineLinkInfo,
	event any,
) error {
	parsedEvent := entity.LINEWebhookEvent{}
	eventBytes, _ := json.Marshal(event)
	if err := json.Unmarshal(eventBytes, &parsedEvent); err != nil {
		return terrors.Wrap(err)
	}
	switch parsedEvent.Type {
	case entity.LINEWebhookEventTypeFollow:
		return t.procLINEMessagingAPIWebhook_Follow(ctx, lineLinkInfo, eventBytes)
	case entity.LINEWebhookEventTypeMessage:
		return t.procLINEMessagingAPIWebhook_Message(ctx, lineLinkInfo, eventBytes)
	}
	return nil
}
