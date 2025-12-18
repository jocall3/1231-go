// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3_test

import (
	"context"
	"os"
	"testing"

	"github.com/jocall3/1231-go"
	"github.com/jocall3/1231-go/internal/testutil"
	"github.com/jocall3/1231-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	user, err := client.Users.Register(context.TODO(), jocall3.UserRegisterParams{
		Email:    jocall3.F[any]("alice.w@example.com"),
		Name:     jocall3.F[any]("Alice Wonderland"),
		Password: jocall3.F[any]("SecureP@ssw0rd2024!"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", user.ID)
}
