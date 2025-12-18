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

func TestAIOracleSimulateRunAdvancedWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Oracle.Simulate.RunAdvanced(context.TODO(), jocall3.AIOracleSimulateRunAdvancedParams{
		Prompt: jocall3.F[any]("Evaluate the long-term impact of a sudden job loss combined with a variable market downturn, analyzing worst-case and best-case recovery scenarios over a decade."),
		Scenarios: jocall3.F([]jocall3.AIOracleSimulateRunAdvancedParamsScenario{{
			DurationYears: jocall3.F[any](10),
			Events: jocall3.F([]jocall3.AIOracleSimulateRunAdvancedParamsScenariosEvent{{
				Details: jocall3.F[any](map[string]interface{}{
					"durationMonths":       6,
					"severanceAmount":      10000,
					"unemploymentBenefits": 2000,
				}),
				Type: jocall3.F(jocall3.AIOracleSimulateRunAdvancedParamsScenariosEventsTypeJobLoss),
			}, {
				Details: jocall3.F[any](map[string]interface{}{
					"impactPercentage": 0.15,
					"recoveryYears":    3,
				}),
				Type: jocall3.F(jocall3.AIOracleSimulateRunAdvancedParamsScenariosEventsTypeMarketDownturn),
			}}),
			Name: jocall3.F[any]("Job Loss & Mild Market Recovery"),
			SensitivityAnalysisParams: jocall3.F([]jocall3.AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam{{
				Max:       jocall3.F[any](0.07),
				Min:       jocall3.F[any](0.03),
				ParamName: jocall3.F[any]("marketRecoveryRate"),
				Step:      jocall3.F[any](0.01),
			}}),
		}}),
		GlobalEconomicFactors: jocall3.F(jocall3.AIOracleSimulateRunAdvancedParamsGlobalEconomicFactors{
			InflationRate:        jocall3.F[any](0.03),
			InterestRateBaseline: jocall3.F[any](0.05),
		}),
		PersonalAssumptions: jocall3.F(jocall3.AIOracleSimulateRunAdvancedParamsPersonalAssumptions{
			AnnualSavingsRate: jocall3.F[any](0.15),
			RiskTolerance:     jocall3.F(jocall3.AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceAggressive),
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

func TestAIOracleSimulateRunStandardWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Oracle.Simulate.RunStandard(context.TODO(), jocall3.AIOracleSimulateRunStandardParams{
		Prompt: jocall3.F[any]("What if I invest an additional $1,000 per month into my aggressive growth portfolio for the next 5 years?"),
		Parameters: jocall3.F[any](map[string]interface{}{
			"durationYears":           5,
			"monthlyInvestmentAmount": 1000,
			"riskTolerance":           "aggressive",
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
