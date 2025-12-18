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

func TestNotificationSettingGet(t *testing.T) {
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
	_, err := client.Notifications.Settings.Get(context.TODO())
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.Settings.Update(context.TODO(), jocall3.NotificationSettingUpdateParams{
		ChannelPreferences: jocall3.F(jocall3.NotificationSettingUpdateParamsChannelPreferences{
			Email: jocall3.F[any](true),
			InApp: jocall3.F[any](true),
			Push:  jocall3.F[any](true),
			SMS:   jocall3.F[any](true),
		}),
		EventPreferences: jocall3.F(jocall3.NotificationSettingUpdateParamsEventPreferences{
			AIInsights:        jocall3.F[any](true),
			BudgetAlerts:      jocall3.F[any](true),
			PromotionalOffers: jocall3.F[any](false),
			SecurityAlerts:    jocall3.F[any](true),
			TransactionAlerts: jocall3.F[any](true),
		}),
		QuietHours: jocall3.F(jocall3.NotificationSettingUpdateParamsQuietHours{
			Enabled:   jocall3.F[any](true),
			EndTime:   jocall3.F[any]("08:00"),
			StartTime: jocall3.F[any]("22:00"),
		}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
