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

func TestIdentityKYCGetStatus(t *testing.T) {
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
	_, err := client.Identity.KYC.GetStatus(context.TODO())
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentityKYCSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.KYC.Submit(context.TODO(), jocall3.IdentityKYCSubmitParams{
		CountryOfIssue:      jocall3.F[any]("US"),
		DocumentNumber:      jocall3.F[any]("ABC12345"),
		DocumentType:        jocall3.F(jocall3.IdentityKYCSubmitParamsDocumentTypeDriversLicense),
		ExpirationDate:      jocall3.F[any]("2030-01-01"),
		IssueDate:           jocall3.F[any]("2020-01-01"),
		AdditionalDocuments: jocall3.F([]interface{}{map[string]interface{}{}}),
		DocumentBackImage:   jocall3.F[any]("base64encoded_image_of_drivers_license_back"),
		DocumentFrontImage:  jocall3.F[any]("base64encoded_image_of_drivers_license_front"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
