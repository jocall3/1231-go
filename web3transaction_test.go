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

func TestWeb3TransactionInitiateTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Web3.Transactions.InitiateTransfer(context.TODO(), jocall3.Web3TransactionInitiateTransferParams{
		Amount:            jocall3.F[any](0.1),
		AssetSymbol:       jocall3.F[any]("ETH"),
		BlockchainNetwork: jocall3.F[any]("Ethereum"),
		RecipientAddress:  jocall3.F[any]("0xdef4567890abcdef1234567890abcdef1234567890"),
		SourceWalletID:    jocall3.F[any]("wallet_conn_eth_0xabc123"),
		GasPriceGwei:      jocall3.F[any](50),
		Memo:              jocall3.F[any]("Payment for services"),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
