package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

var deleteMeState *oauth2loginflow.State

func (t *Impl) CreateOAuth2State(
	ctx context.Context,
	state *oauth2loginflow.State,
) (*oauth2loginflow.State, error) {
	deleteMeState = state
	return state, nil
}

func (t *Impl) GetAndDeleteOAuth2State(
	ctx context.Context,
	stateCode oauth2loginflow.StateCode,
) (*oauth2loginflow.State, error) {
	if deleteMeState == nil {
		return nil, terrors.Wrap(&common_repository.NoEntryError{})
	}
	return deleteMeState, nil
}
