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

func TestPaymentFxConvertWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Fx.Convert(context.TODO(), jocall3.PaymentFxConvertParams{
		SourceAccountID: jocall3.F[any]("acc_chase_checking_4567"),
		SourceAmount:    jocall3.F[any](1000),
		SourceCurrency:  jocall3.F[any]("USD"),
		TargetCurrency:  jocall3.F[any]("EUR"),
		FxRateLock:      jocall3.F[any](true),
		TargetAccountID: jocall3.F[any]("acc_euro_savings_9876"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentFxGetRatesWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Fx.GetRates(context.TODO(), jocall3.PaymentFxGetRatesParams{
		BaseCurrency:   jocall3.F[any]("USD"),
		TargetCurrency: jocall3.F[any]("EUR"),
		ForecastDays:   jocall3.F[any](7),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
