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

func TestCorporateCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.List(context.TODO(), jocall3.CorporateCardListParams{
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

func TestCorporateCardNewVirtualWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.NewVirtual(context.TODO(), jocall3.CorporateCardNewVirtualParams{
		Controls: jocall3.F(jocall3.CorporateCardControlsParam{
			AtmWithdrawals:               jocall3.F[any](false),
			ContactlessPayments:          jocall3.F[any](false),
			DailyLimit:                   jocall3.F[any](500),
			InternationalTransactions:    jocall3.F[any](false),
			MerchantCategoryRestrictions: jocall3.F([]interface{}{"Advertising"}),
			MonthlyLimit:                 jocall3.F[any](1000),
			OnlineTransactions:           jocall3.F[any](true),
			SingleTransactionLimit:       jocall3.F[any](200),
			VendorRestrictions:           jocall3.F([]interface{}{"Facebook Ads", "Google Ads"}),
		}),
		ExpirationDate:       jocall3.F[any]("2025-12-31"),
		HolderName:           jocall3.F[any]("Marketing Campaign Q4"),
		Purpose:              jocall3.F[any]("Online advertising for Q4 campaigns"),
		AssociatedEmployeeID: jocall3.F[any]("emp_marketing_01"),
		SpendingPolicyID:     jocall3.F[any]("policy_marketing_fixed"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateCardFreeze(t *testing.T) {
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
	_, err := client.Corporate.Cards.Freeze(
		context.TODO(),
		"corp_card_xyz987654",
		jocall3.CorporateCardFreezeParams{
			Freeze: jocall3.F[any](true),
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

func TestCorporateCardListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.ListTransactions(
		context.TODO(),
		"corp_card_xyz987654",
		jocall3.CorporateCardListTransactionsParams{
			EndDate:   jocall3.F[any]("2024-12-31"),
			Limit:     jocall3.F[any](map[string]interface{}{}),
			Offset:    jocall3.F[any](map[string]interface{}{}),
			StartDate: jocall3.F[any]("2024-01-01"),
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

func TestCorporateCardUpdateControlsWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.UpdateControls(
		context.TODO(),
		"corp_card_xyz987654",
		jocall3.CorporateCardUpdateControlsParams{
			CorporateCardControls: jocall3.CorporateCardControlsParam{
				AtmWithdrawals:               jocall3.F[any](true),
				ContactlessPayments:          jocall3.F[any](true),
				DailyLimit:                   jocall3.F[any](750),
				InternationalTransactions:    jocall3.F[any](true),
				MerchantCategoryRestrictions: jocall3.F([]interface{}{"Software Subscriptions", "Conferences"}),
				MonthlyLimit:                 jocall3.F[any](3000),
				OnlineTransactions:           jocall3.F[any](true),
				SingleTransactionLimit:       jocall3.F[any](1000),
				VendorRestrictions:           jocall3.F([]interface{}{"Amazon", "Uber"}),
			},
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
