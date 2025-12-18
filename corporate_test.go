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

func TestCorporatePerformSanctionScreeningWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.PerformSanctionScreening(context.TODO(), jocall3.CorporatePerformSanctionScreeningParams{
		Country:    jocall3.F[any]("US"),
		EntityType: jocall3.F(jocall3.CorporatePerformSanctionScreeningParamsEntityTypeIndividual),
		Name:       jocall3.F[any]("John Doe"),
		Address: jocall3.F(jocall3.AddressParam{
			City:    jocall3.F[any]("Anytown"),
			Country: jocall3.F[any]("USA"),
			State:   jocall3.F[any]("CA"),
			Street:  jocall3.F[any]("123 Main St"),
			Zip:     jocall3.F[any]("90210"),
		}),
		DateOfBirth:          jocall3.F[any]("1970-01-01"),
		IdentificationNumber: jocall3.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
