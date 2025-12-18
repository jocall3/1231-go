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

func TestCorporateAnomalyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Anomalies.List(context.TODO(), jocall3.CorporateAnomalyListParams{
		EndDate:    jocall3.F[any]("2024-12-31"),
		EntityType: jocall3.F(jocall3.CorporateAnomalyListParamsEntityTypeTransaction),
		Limit:      jocall3.F[any](map[string]interface{}{}),
		Offset:     jocall3.F[any](map[string]interface{}{}),
		Severity:   jocall3.F(jocall3.CorporateAnomalyListParamsSeverityCritical),
		StartDate:  jocall3.F[any]("2024-01-01"),
		Status:     jocall3.F(jocall3.CorporateAnomalyListParamsStatusNew),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateAnomalyUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Anomalies.UpdateStatus(
		context.TODO(),
		"anom_risk-2024-07-21-D1E2F3",
		jocall3.CorporateAnomalyUpdateStatusParams{
			Status:          jocall3.F(jocall3.CorporateAnomalyUpdateStatusParamsStatusResolved),
			ResolutionNotes: jocall3.F[any]("Confirmed legitimate transaction after contacting vendor. Marked as resolved."),
		},
	)
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
