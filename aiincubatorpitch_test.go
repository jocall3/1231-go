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

func TestAIIncubatorPitchGetDetails(t *testing.T) {
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
	_, err := client.AI.Incubator.Pitch.GetDetails(context.TODO(), "pitch_qw_synergychain-xyz")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIIncubatorPitchSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Incubator.Pitch.Submit(context.TODO(), jocall3.AIIncubatorPitchSubmitParams{
		BusinessPlan: jocall3.F[any]("Quantum-AI powered financial advisor platform leveraging neural networks for predictive analytics and hyper-personalized advice..."),
		FinancialProjections: jocall3.F(jocall3.AIIncubatorPitchSubmitParamsFinancialProjections{
			ProfitabilityEstimate: jocall3.F[any]("Achieve profitability within 18 months."),
			ProjectionYears:       jocall3.F[any](3),
			RevenueForecast:       jocall3.F([]interface{}{500000, 2000000, 6000000}),
			SeedRoundAmount:       jocall3.F[any](2500000),
			ValuationPreMoney:     jocall3.F[any](10000000),
		}),
		FoundingTeam: jocall3.F([]jocall3.AIIncubatorPitchSubmitParamsFoundingTeam{{
			Experience: jocall3.F[any]("15+ years in AI/ML, PhD in Quantum Computing, ex-Google Brain"),
			Name:       jocall3.F[any]("Dr. Eleanor Vance"),
			Role:       jocall3.F[any]("CEO & Lead AI Scientist"),
		}, {
			Experience: jocall3.F[any]("20+ years in Fintech, ex-Goldman Sachs"),
			Name:       jocall3.F[any]("Marcus Thorne"),
			Role:       jocall3.F[any]("COO & Finance Expert"),
		}}),
		MarketOpportunity: jocall3.F[any]("The booming digital finance market coupled with demand for truly personalized, AI-driven financial guidance presents a multi-billion dollar opportunity. Our unique quantum-AI approach provides unparalleled accuracy and foresight."),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIIncubatorPitchSubmitFeedbackWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Incubator.Pitch.SubmitFeedback(
		context.TODO(),
		"pitch_qw_synergychain-xyz",
		jocall3.AIIncubatorPitchSubmitFeedbackParams{
			Answers: jocall3.F([]jocall3.AIIncubatorPitchSubmitFeedbackParamsAnswer{{
				Answer:     jocall3.F[any]("Our mitigation strategy includes dedicated R&D and new hires with specific expertise."),
				QuestionID: jocall3.F[any]("q_qa-team-001"),
			}, {
				Answer:     jocall3.F[any]("Our CAC projections are based on pilot program results showing $500 per enterprise client with a conversion rate of 10% from trials."),
				QuestionID: jocall3.F[any]("q_qa-market-002"),
			}}),
			Feedback: jocall3.F[any]("Regarding the technical challenges, our team has allocated 3 months for R&D on quantum-resistant cryptography, mitigating the risk. We've also brought in Dr. Elena Petrova, a leading expert in secure multi-party computation."),
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
