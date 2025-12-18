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

func TestCorporateComplianceAuditRequestWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Compliance.Audits.Request(context.TODO(), jocall3.CorporateComplianceAuditRequestParams{
		AuditScope:           jocall3.F(jocall3.CorporateComplianceAuditRequestParamsAuditScopeAllTransactions),
		EndDate:              jocall3.F[any]("2024-06-30"),
		RegulatoryFrameworks: jocall3.F([]jocall3.CorporateComplianceAuditRequestParamsRegulatoryFramework{jocall3.CorporateComplianceAuditRequestParamsRegulatoryFrameworkAml, jocall3.CorporateComplianceAuditRequestParamsRegulatoryFrameworkPciDss}),
		StartDate:            jocall3.F[any]("2024-01-01"),
		AdditionalContext:    jocall3.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateComplianceAuditGetReport(t *testing.T) {
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
	_, err := client.Corporate.Compliance.Audits.GetReport(context.TODO(), "audit_corp_xyz789")
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
