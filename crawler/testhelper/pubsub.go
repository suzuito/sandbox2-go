package testhelper

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/pubsub"
)

type TestCaseForPubSub struct {
	SetUp    func(ctx context.Context, fcli *pubsub.Client) error
	TearDown func(ctx context.Context, fcli *pubsub.Client) error
}

func (th *TestCaseForPubSub) Run(
	ctx context.Context,
	desc string,
	t *testing.T,
	run func(t *testing.T, pcli *pubsub.Client),
) bool {
	return t.Run(desc, func(t *testing.T) {
		pcli, err := NewPubSubClient(ctx)
		if err != nil {
			t.Errorf("failed to NewPubSubClient : %+v", err)
			t.Fail()
			return
		}
		defer func() {
			if th.TearDown != nil {
				th.TearDown(ctx, pcli)
			}
			pcli.Close()
		}()
		if th.SetUp != nil {
			if err := th.SetUp(ctx, pcli); err != nil {
				t.Errorf("failed to SetUp : %+v", err)
				t.Fail()
				return
			}
		}
		run(t, pcli)
	})
}

func NewPubSubClient(ctx context.Context) (*pubsub.Client, error) {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8083")
	cli, err := pubsub.NewClient(ctx, "dummy-prj")
	if err != nil {
		return nil, err
	}
	return cli, nil
}
