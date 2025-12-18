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

func TestPaymentInternationalInitiateWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.International.Initiate(context.TODO(), jocall3.PaymentInternationalInitiateParams{
		Amount: jocall3.F[any](5000),
		Beneficiary: jocall3.F(jocall3.PaymentInternationalInitiateParamsBeneficiary{
			Address:       jocall3.F[any]("Hauptstrasse 1, 10115 Berlin, Germany"),
			BankName:      jocall3.F[any]("Deutsche Bank"),
			Name:          jocall3.F[any]("Maria Schmidt"),
			AccountNumber: jocall3.F[any](map[string]interface{}{}),
			Iban:          jocall3.F[any]("DE89370400440532013000"),
			RoutingNumber: jocall3.F[any](map[string]interface{}{}),
			SwiftBic:      jocall3.F[any]("DEUTDEFF"),
		}),
		Purpose:         jocall3.F[any]("Vendor payment for Q2 services."),
		SourceAccountID: jocall3.F[any]("acc_chase_checking_4567"),
		SourceCurrency:  jocall3.F[any]("USD"),
		TargetCurrency:  jocall3.F[any]("EUR"),
		FxRateLock:      jocall3.F[any](true),
		FxRateProvider:  jocall3.F(jocall3.PaymentInternationalInitiateParamsFxRateProviderProprietaryAI),
		Reference:       jocall3.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentInternationalGetStatus(t *testing.T) {
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
	_, err := client.Payments.International.GetStatus(context.TODO(), "int_pmt_xyz7890")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
