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

func TestLendingApplicationGet(t *testing.T) {
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
	_, err := client.Lending.Applications.Get(context.TODO(), "loan_app_creditflow-123")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLendingApplicationSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Lending.Applications.Submit(context.TODO(), jocall3.LendingApplicationSubmitParams{
		LoanAmount:          jocall3.F[any](10000),
		LoanPurpose:         jocall3.F(jocall3.LendingApplicationSubmitParamsLoanPurposeHomeImprovement),
		RepaymentTermMonths: jocall3.F[any](36),
		AdditionalNotes:     jocall3.F[any]("Funds needed to replace a broken HVAC system."),
		CoApplicant: jocall3.F(jocall3.LendingApplicationSubmitParamsCoApplicant{
			Email:  jocall3.F[any]("jane.doe@example.com"),
			Income: jocall3.F[any](75000),
			Name:   jocall3.F[any]("Jane Doe"),
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
