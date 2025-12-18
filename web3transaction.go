// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// Web3TransactionService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3TransactionService] method instead.
type Web3TransactionService struct {
	Options []option.RequestOption
}

// NewWeb3TransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3TransactionService(opts ...option.RequestOption) (r *Web3TransactionService) {
	r = &Web3TransactionService{}
	r.Options = opts
	return
}

// Prepares and initiates a cryptocurrency transfer from a connected wallet to a
// specified recipient address. Requires user confirmation (e.g., via wallet
// signature).
func (r *Web3TransactionService) InitiateTransfer(ctx context.Context, body Web3TransactionInitiateTransferParams, opts ...option.RequestOption) (res *Web3TransactionInitiateTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/transactions/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Web3TransactionInitiateTransferResponse struct {
	// Current status of the transfer.
	Status Web3TransactionInitiateTransferResponseStatus `json:"status,required"`
	// Unique identifier for this cryptocurrency transfer operation.
	TransferID interface{} `json:"transferId,required"`
	// The blockchain transaction hash, if available and confirmed.
	BlockchainTxnHash interface{} `json:"blockchainTxnHash"`
	// A descriptive message about the transfer status.
	Message interface{}                                 `json:"message"`
	JSON    web3TransactionInitiateTransferResponseJSON `json:"-"`
}

// web3TransactionInitiateTransferResponseJSON contains the JSON metadata for the
// struct [Web3TransactionInitiateTransferResponse]
type web3TransactionInitiateTransferResponseJSON struct {
	Status            apijson.Field
	TransferID        apijson.Field
	BlockchainTxnHash apijson.Field
	Message           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Web3TransactionInitiateTransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3TransactionInitiateTransferResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the transfer.
type Web3TransactionInitiateTransferResponseStatus string

const (
	Web3TransactionInitiateTransferResponseStatusPendingSignature              Web3TransactionInitiateTransferResponseStatus = "pending_signature"
	Web3TransactionInitiateTransferResponseStatusPendingBlockchainConfirmation Web3TransactionInitiateTransferResponseStatus = "pending_blockchain_confirmation"
	Web3TransactionInitiateTransferResponseStatusCompleted                     Web3TransactionInitiateTransferResponseStatus = "completed"
	Web3TransactionInitiateTransferResponseStatusFailed                        Web3TransactionInitiateTransferResponseStatus = "failed"
	Web3TransactionInitiateTransferResponseStatusCancelled                     Web3TransactionInitiateTransferResponseStatus = "cancelled"
)

func (r Web3TransactionInitiateTransferResponseStatus) IsKnown() bool {
	switch r {
	case Web3TransactionInitiateTransferResponseStatusPendingSignature, Web3TransactionInitiateTransferResponseStatusPendingBlockchainConfirmation, Web3TransactionInitiateTransferResponseStatusCompleted, Web3TransactionInitiateTransferResponseStatusFailed, Web3TransactionInitiateTransferResponseStatusCancelled:
		return true
	}
	return false
}

type Web3TransactionInitiateTransferParams struct {
	// The amount of cryptocurrency to transfer.
	Amount param.Field[interface{}] `json:"amount,required"`
	// Symbol of the crypto asset to transfer (e.g., ETH, USDC).
	AssetSymbol param.Field[interface{}] `json:"assetSymbol,required"`
	// The blockchain network for the transfer.
	BlockchainNetwork param.Field[interface{}] `json:"blockchainNetwork,required"`
	// The recipient's blockchain address.
	RecipientAddress param.Field[interface{}] `json:"recipientAddress,required"`
	// ID of the connected wallet from which to send funds.
	SourceWalletID param.Field[interface{}] `json:"sourceWalletId,required"`
	// Optional: Gas price in Gwei for Ethereum-based transactions.
	GasPriceGwei param.Field[interface{}] `json:"gasPriceGwei"`
	// Optional: A short memo or note for the transaction.
	Memo param.Field[interface{}] `json:"memo"`
}

func (r Web3TransactionInitiateTransferParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
