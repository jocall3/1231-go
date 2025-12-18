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

func TestBudgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.New(context.TODO(), jocall3.BudgetNewParams{
		EndDate:        jocall3.F[any]("2024-09-30"),
		Name:           jocall3.F[any]("September Living Expenses"),
		Period:         jocall3.F(jocall3.BudgetNewParamsPeriodMonthly),
		StartDate:      jocall3.F[any]("2024-09-01"),
		TotalAmount:    jocall3.F[any](2800),
		AIAutoPopulate: jocall3.F[any](true),
		AlertThreshold: jocall3.F[any](75),
		Categories: jocall3.F([]jocall3.BudgetNewParamsCategory{{
			Allocated: jocall3.F[any](1500),
			Name:      jocall3.F[any]("Rent"),
		}, {
			Allocated: jocall3.F[any](400),
			Name:      jocall3.F[any]("Groceries"),
		}}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetGet(t *testing.T) {
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
	_, err := client.Budgets.Get(context.TODO(), "budget_monthly_aug")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Update(
		context.TODO(),
		"budget_monthly_aug",
		jocall3.BudgetUpdateParams{
			AlertThreshold: jocall3.F[any](85),
			Categories: jocall3.F([]jocall3.BudgetUpdateParamsCategory{{
				Allocated: jocall3.F[any](550),
				Name:      jocall3.F[any]("Groceries"),
			}}),
			EndDate:     jocall3.F[any]("2024-08-31"),
			Name:        jocall3.F[any]("August 2024 Revised Household Budget"),
			StartDate:   jocall3.F[any]("2024-08-01"),
			Status:      jocall3.F(jocall3.BudgetUpdateParamsStatusActive),
			TotalAmount: jocall3.F[any](3200),
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

func TestBudgetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.List(context.TODO(), jocall3.BudgetListParams{
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

func TestBudgetDelete(t *testing.T) {
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
	err := client.Budgets.Delete(context.TODO(), "budget_monthly_aug")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
