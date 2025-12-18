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

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "txn_quantum-2024-07-21-A7B8C9")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), jocall3.TransactionListParams{
		Category:    jocall3.F[any]("Groceries"),
		EndDate:     jocall3.F[any]("2024-12-31"),
		Limit:       jocall3.F[any](map[string]interface{}{}),
		MaxAmount:   jocall3.F[any](100),
		MinAmount:   jocall3.F[any](20),
		Offset:      jocall3.F[any](map[string]interface{}{}),
		SearchQuery: jocall3.F[any]("Starbucks"),
		StartDate:   jocall3.F[any]("2024-01-01"),
		Type:        jocall3.F(jocall3.TransactionListParamsTypeExpense),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionCategorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Categorize(
		context.TODO(),
		"txn_quantum-2024-07-21-A7B8C9",
		jocall3.TransactionCategorizeParams{
			Category:      jocall3.F[any]("Home > Groceries"),
			ApplyToFuture: jocall3.F[any](true),
			Notes:         jocall3.F[any]("Bulk purchase for party"),
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

func TestTransactionDisputeWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Dispute(
		context.TODO(),
		"txn_quantum-2024-07-21-A7B8C9",
		jocall3.TransactionDisputeParams{
			Details:             jocall3.F[any]("I did not authorize this purchase. My card may have been compromised and I was traveling internationally on this date."),
			Reason:              jocall3.F(jocall3.TransactionDisputeParamsReasonUnauthorized),
			SupportingDocuments: jocall3.F([]interface{}{"https://demobank.com/uploads/flight_ticket.png"}),
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

func TestTransactionUpdateNotes(t *testing.T) {
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
	_, err := client.Transactions.UpdateNotes(
		context.TODO(),
		"txn_quantum-2024-07-21-A7B8C9",
		jocall3.TransactionUpdateNotesParams{
			Notes: jocall3.F[any]("This was a special coffee for a client meeting."),
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
