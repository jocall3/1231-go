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

func TestAIAdvisorChatGetHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Advisor.Chat.GetHistory(context.TODO(), jocall3.AIAdvisorChatGetHistoryParams{
		Limit:     jocall3.F[any](map[string]interface{}{}),
		Offset:    jocall3.F[any](map[string]interface{}{}),
		SessionID: jocall3.F[any]("session-quantum-xyz-789-alpha"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAdvisorChatSendMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Advisor.Chat.SendMessage(context.TODO(), jocall3.AIAdvisorChatSendMessageParams{
		FunctionResponse: jocall3.F(jocall3.AIAdvisorChatSendMessageParamsFunctionResponse{
			Name: jocall3.F[any]("send_money"),
			Response: jocall3.F[any](map[string]interface{}{
				"status":        "success",
				"transactionId": "pmt_654321",
				"amountSent":    55.5,
				"recipient":     "Alex",
			}),
		}),
		Message:   jocall3.F[any]("Can you analyze my recent spending patterns and suggest areas for saving, focusing on my dining expenses?"),
		SessionID: jocall3.F[any]("session-quantum-xyz-789-alpha"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
