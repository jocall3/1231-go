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

func TestCorporateRiskFraudRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Risk.Fraud.Rules.New(context.TODO(), jocall3.CorporateRiskFraudRuleNewParams{
		Action: jocall3.F(jocall3.FraudRuleActionParam{
			Details:    jocall3.F[any]("Hold payment, notify sender for additional verification, and escalate to compliance."),
			Type:       jocall3.F(jocall3.FraudRuleActionTypeAutoReview),
			TargetTeam: jocall3.F[any]("Fraud Prevention Team"),
		}),
		Criteria: jocall3.F(jocall3.FraudRuleCriteriaParam{
			AccountInactivityDays:     jocall3.F[any](90),
			CountryOfOrigin:           jocall3.F([]interface{}{"US", "CA"}),
			GeographicDistanceKm:      jocall3.F[any](5000),
			LastLoginDays:             jocall3.F[any](7),
			NoTravelNotification:      jocall3.F[any](true),
			PaymentCountMin:           jocall3.F[any](3),
			RecipientCountryRiskLevel: jocall3.F([]jocall3.FraudRuleCriteriaRecipientCountryRiskLevel{jocall3.FraudRuleCriteriaRecipientCountryRiskLevelHigh, jocall3.FraudRuleCriteriaRecipientCountryRiskLevelVeryHigh}),
			RecipientNew:              jocall3.F[any](true),
			TimeframeHours:            jocall3.F[any](24),
			TransactionAmountMin:      jocall3.F[any](5000),
			TransactionType:           jocall3.F(jocall3.FraudRuleCriteriaTransactionTypeDebit),
		}),
		Description: jocall3.F[any]("Detects multiple international payments to new beneficiaries in high-risk countries within a short timeframe."),
		Name:        jocall3.F[any]("Suspicious International Payment Pattern"),
		Severity:    jocall3.F(jocall3.CorporateRiskFraudRuleNewParamsSeverityCritical),
		Status:      jocall3.F(jocall3.CorporateRiskFraudRuleNewParamsStatusActive),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateRiskFraudRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Risk.Fraud.Rules.Update(
		context.TODO(),
		"fraud_rule_high_value_inactive",
		jocall3.CorporateRiskFraudRuleUpdateParams{
			Action: jocall3.F(jocall3.FraudRuleActionParam{
				Details:    jocall3.F[any]("Flag for manual review only, do not block."),
				Type:       jocall3.F(jocall3.FraudRuleActionTypeBlock),
				TargetTeam: jocall3.F[any]("Fraud Prevention Team"),
			}),
			Criteria: jocall3.F(jocall3.FraudRuleCriteriaParam{
				AccountInactivityDays:     jocall3.F[any](60),
				CountryOfOrigin:           jocall3.F([]interface{}{"US", "CA"}),
				GeographicDistanceKm:      jocall3.F[any](5000),
				LastLoginDays:             jocall3.F[any](7),
				NoTravelNotification:      jocall3.F[any](true),
				PaymentCountMin:           jocall3.F[any](3),
				RecipientCountryRiskLevel: jocall3.F([]jocall3.FraudRuleCriteriaRecipientCountryRiskLevel{jocall3.FraudRuleCriteriaRecipientCountryRiskLevelLow}),
				RecipientNew:              jocall3.F[any](true),
				TimeframeHours:            jocall3.F[any](24),
				TransactionAmountMin:      jocall3.F[any](7500),
				TransactionType:           jocall3.F(jocall3.FraudRuleCriteriaTransactionTypeDebit),
			}),
			Description: jocall3.F[any]("Revised logic for flagging high value transactions from dormant accounts."),
			Name:        jocall3.F[any]("Revised High Value Transaction Rule"),
			Severity:    jocall3.F(jocall3.CorporateRiskFraudRuleUpdateParamsSeverityHigh),
			Status:      jocall3.F(jocall3.CorporateRiskFraudRuleUpdateParamsStatusInactive),
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

func TestCorporateRiskFraudRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Risk.Fraud.Rules.List(context.TODO(), jocall3.CorporateRiskFraudRuleListParams{
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

func TestCorporateRiskFraudRuleDelete(t *testing.T) {
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
	err := client.Corporate.Risk.Fraud.Rules.Delete(context.TODO(), "fraud_rule_high_value_inactive")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
