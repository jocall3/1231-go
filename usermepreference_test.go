// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/jocall3/1231-go"
	"github.com/jocall3/1231-go/internal/testutil"
	"github.com/jocall3/1231-go/option"
)

func TestUserMePreferenceGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jocall3.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Users.Me.Preferences.Get(context.TODO())
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserMePreferenceUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jocall3.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Users.Me.Preferences.Update(context.TODO(), jocall3.UserMePreferenceUpdateParams{
		UserPreferences: jocall3.UserPreferencesParam{
			AIInteractionMode:  jocall3.F(jocall3.UserPreferencesAIInteractionModeProactive),
			DataSharingConsent: jocall3.F[any](true),
			NotificationChannels: jocall3.F(jocall3.UserPreferencesNotificationChannelsParam{
				Email: jocall3.F[any](true),
				InApp: jocall3.F[any](true),
				Push:  jocall3.F[any](true),
				SMS:   jocall3.F[any](false),
			}),
			PreferredLanguage:   jocall3.F[any]("en-US"),
			Theme:               jocall3.F[any]("Dark-Quantum"),
			TransactionGrouping: jocall3.F(jocall3.UserPreferencesTransactionGroupingCategory),
		},
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
