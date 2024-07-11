package web

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

// https://developers.line.biz/ja/reference/messaging-api
type LineWebhookMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		Type      string `json:"type"`
		Mode      string `json:"mode"`
		Timestamp int64  `json:"timestamp"`
		Source    *struct {
			Type    string `json:"type"`
			UserID  string `json:"userId"`
			GroupID string `json:"groupId"`
			RoomID  string `json:"roomId"`
		} `json:"source"`
		WebhookEventID  string `json:"webhookEventId"`
		DeliveryContext struct {
			IsRedelivery bool `json:"isRedelivery"`
		} `json:"deliveryContext"`
		ReplyToken string `json:"replyToken"`
		Message    *struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			QuoteToken string `json:"quoteToken"`
			Text       string `json:"text"`
			Emojis     []struct {
				Index     int    `json:"index"`
				Length    int    `json:"length"`
				ProductID string `json:"productId"`
				EmojiID   string `json:"emojiId"`
			} `json:"emojis"`
			Mention *struct {
				Mentionees []struct {
					Index  int    `json:"index"`
					Length int    `json:"length"`
					Type   string `json:"type"`
					UserID string `json:"userId"`
				} `json:"mentionees"`
			} `json:"mention"`
			QuotedMessageID string `json:"quotedMessageId"`
			ContentProvider *struct {
				Type               string `json:"type"`
				OriginalContentURL string `json:"originalContentUrl"`
				PreviewImageURL    string `json:"previewImageUrl"`
			} `json:"contentProvider"`
			ImageSet *struct {
				ID    string `json:"id"`
				Index int    `json:"index"`
				Total int    `json:"total"`
			} `json:"imageSet"`
			Duration            int64    `json:"duration"`
			FileName            string   `json:"fileName"`
			FileSize            int64    `json:"fileSize"`
			Title               string   `json:"title"`
			Address             string   `json:"address"`
			Latitude            float64  `json:"latitude"`
			Longitude           float64  `json:"longitude"`
			PackageID           string   `json:"packageId"`
			StickerID           string   `json:"stickerId"`
			StickerResourceType string   `json:"stickerResourceType"`
			Keywords            []string `json:"keywords"`
			Unsend              *struct {
				MessageID string `json:"messageId"`
			} `json:"unsend"`
			Follow *struct {
				IsUnblocked bool `json:"isUnblocked"`
			} `json:"follow"`
			Joined *struct {
				Members []struct {
					Type   string `json:"type"`
					UserID string `json:"userId"`
				} `json:"members"`
			} `json:"joined"`
			Left *struct {
				Members []struct {
					Type   string `json:"type"`
					UserID string `json:"userId"`
				} `json:"members"`
			} `json:"left"`
			Postback *struct {
				Data   string         `json:"data"`
				Params map[string]any `json:"params"`
			} `json:"postback"`
			VideoPlayComplete *struct {
				TrackingID string `json:"trackingId"`
			} `json:"videoPlayComplete"`
			Beacon *struct {
				HwID string `json:"hwid"`
				Type string `json:"type"`
				Dm   string `json:"dm"`
			} `json:"beacon"`
			Link *struct {
				Result string `json:"result"`
				Nonce  string `json:"nonce"`
			} `json:"link"`
		} `json:"message"`
	} `json:"events"`
}

func (t *Impl) APIPostLineMessagingAPIWebhook(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		t.P.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
			Message: "%+v",
		})
		return
	}
	// Print json
	bodyMap := map[string]any{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		t.L.Error("", "err", err)
	} else {
		t.L.Info("", "body", string(body))
		t.L.Info("", "bodyMap", bodyMap)
	}
	// Verify header
	channelSecret := "213c14c37f5f4ef0d6ba3ec7daee6585"
	signatureLeft, err := base64.StdEncoding.DecodeString(ctx.GetHeader("x-line-signature"))
	if err != nil {
		t.L.Error("", "err", err)
	} else {
		h := hmac.New(sha256.New, []byte(channelSecret))
		h.Write(body)
		signatureRight := h.Sum(nil)
		t.L.Info(
			"Verify header",
			"matched", bytes.Equal(signatureLeft, signatureRight),
		)
	}
	ctx.Status(200)
}
