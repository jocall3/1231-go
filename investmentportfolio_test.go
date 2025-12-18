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

func TestInvestmentPortfolioNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.New(context.TODO(), jocall3.InvestmentPortfolioNewParams{
		Currency:          jocall3.F[any]("USD"),
		InitialInvestment: jocall3.F[any](10000),
		Name:              jocall3.F[any]("My First Growth Portfolio"),
		RiskTolerance:     jocall3.F(jocall3.InvestmentPortfolioNewParamsRiskToleranceConservative),
		Type:              jocall3.F(jocall3.InvestmentPortfolioNewParamsTypeDiversified),
		AIAutoAllocate:    jocall3.F[any](true),
		LinkedAccountID:   jocall3.F[any]("acc_chase_checking_4567"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPortfolioGet(t *testing.T) {
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
	_, err := client.Investments.Portfolios.Get(context.TODO(), "portfolio_equity_growth")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPortfolioUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.Update(
		context.TODO(),
		"portfolio_equity_growth",
		jocall3.InvestmentPortfolioUpdateParams{
			AIRebalancingFrequency: jocall3.F(jocall3.InvestmentPortfolioUpdateParamsAIRebalancingFrequencyQuarterly),
			Name:                   jocall3.F[any]("Revised Growth Portfolio"),
			RiskTolerance:          jocall3.F(jocall3.InvestmentPortfolioUpdateParamsRiskToleranceConservative),
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

func TestInvestmentPortfolioListWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.List(context.TODO(), jocall3.InvestmentPortfolioListParams{
		Limit:  jocall3.F[any](map[string]interface{}{}),
		Offset: jocall3.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPortfolioRebalanceWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.Rebalance(
		context.TODO(),
		"portfolio_equity_growth",
		jocall3.InvestmentPortfolioRebalanceParams{
			TargetRiskTolerance:  jocall3.F(jocall3.InvestmentPortfolioRebalanceParamsTargetRiskToleranceConservative),
			ConfirmationRequired: jocall3.F[any](true),
			DryRun:               jocall3.F[any](true),
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
