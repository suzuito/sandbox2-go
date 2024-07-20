package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Client interface {
	GetProfile(
		ctx context.Context,
		accessToken string,
		userID string,
	) (*entity.User, error)
}

type Impl struct {
	Cli *http.Client
}

func (t *Impl) GetProfile(
	ctx context.Context,
	accessToken string,
	userID string,
) (*entity.User, error) {
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.line.me/v2/bot/profile/%s", userID),
		nil,
	)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	res, err := t.Cli.Do(req)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, terrors.Wrapf("http error code=%d body=%s", res.StatusCode, string(bodyBytes))
	}
	message := struct {
		DisplayName string `json:"displayName"`
		PictureURL  string `json:"pictureUrl"`
	}{}
	if err := json.Unmarshal(bodyBytes, &message); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &entity.User{
		Name:            message.DisplayName,
		ProfileImageURL: message.PictureURL,
	}, nil
}
