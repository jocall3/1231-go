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

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Login(context.TODO(), jocall3.UserLoginParams{
		Email:    jocall3.F[any]("quantum.visionary@demobank.com"),
		Password: jocall3.F[any]("YourSecurePassword123"),
		MfaCode:  jocall3.F[any]("123456"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserRegisterWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Register(context.TODO(), jocall3.UserRegisterParams{
		Email:    jocall3.F[any]("alice.w@example.com"),
		Name:     jocall3.F[any]("Alice Wonderland"),
		Password: jocall3.F[any]("SecureP@ssw0rd2024!"),
		Address: jocall3.F(jocall3.AddressParam{
			City:    jocall3.F[any]("Anytown"),
			Country: jocall3.F[any]("USA"),
			State:   jocall3.F[any]("CA"),
			Street:  jocall3.F[any]("123 Main St"),
			Zip:     jocall3.F[any]("90210"),
		}),
		DateOfBirth: jocall3.F[any]("1990-05-10"),
		Phone:       jocall3.F[any]("+1-555-987-6543"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
