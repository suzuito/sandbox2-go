package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/notifier/internal/notifierimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

func TestNewConnpassEventNotifier(t *testing.T) {
	testCases := []struct {
		desc                     string
		inputDef                 *notifier.NotifierDefinition
		inputSetting             *factorysetting.NotifierFactorySetting
		expectedDiscordChannelID string
		expectedError            string
	}{
		{
			desc: "Success",
			inputDef: &notifier.NotifierDefinition{
				ID:       notifier.NotifierID("discordconnpassevent"),
				Argument: map[string]any{"DiscordChannelID": "dummy"},
			},
			inputSetting:             &factorysetting.NotifierFactorySetting{},
			expectedDiscordChannelID: "dummy",
		},
		{
			desc: "Failed (not discordblogfeed)",
			inputDef: &notifier.NotifierDefinition{
				ID: notifier.NotifierID("dummy"),
			},
			inputSetting:  &factorysetting.NotifierFactorySetting{},
			expectedError: "NoMatchedNotifierID",
		},
		{
			desc: "Failed (no DiscordChannelID)",
			inputDef: &notifier.NotifierDefinition{
				ID:       notifier.NotifierID("discordconnpassevent"),
				Argument: map[string]any{},
			},
			inputSetting:  &factorysetting.NotifierFactorySetting{},
			expectedError: "Key 'DiscordChannelID' is not found in AgumentDefinition",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			n, err := NewDiscordConnpassEventNotifier(tC.inputDef, tC.inputSetting)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedDiscordChannelID, n.(*notifierimpl.DiscordConnpassEventNotifier).DiscordChannelID)
			}
		})
	}
}
