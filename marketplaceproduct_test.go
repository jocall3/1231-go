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

func TestMarketplaceProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketplace.Products.List(context.TODO(), jocall3.MarketplaceProductListParams{
		AIPersonalizationLevel: jocall3.F(jocall3.MarketplaceProductListParamsAIPersonalizationLevelHigh),
		Category:               jocall3.F(jocall3.MarketplaceProductListParamsCategoryInsurance),
		Limit:                  jocall3.F[any](map[string]interface{}{}),
		MinRating:              jocall3.F[any](4),
		Offset:                 jocall3.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMarketplaceProductSimulateImpactWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketplace.Products.SimulateImpact(
		context.TODO(),
		"prod_home_insurance_quantum",
		jocall3.MarketplaceProductSimulateImpactParams{
			SimulationParameters: jocall3.F[any](map[string]interface{}{
				"loanAmount":          20000,
				"repaymentTermMonths": 48,
			}),
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
