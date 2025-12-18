// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// Web3WalletService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3WalletService] method instead.
type Web3WalletService struct {
	Options []option.RequestOption
}

// NewWeb3WalletService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3WalletService(opts ...option.RequestOption) (r *Web3WalletService) {
	r = &Web3WalletService{}
	r.Options = opts
	return
}

// Retrieves a list of all securely linked cryptocurrency wallets (e.g., MetaMask,
// Ledger integration), showing their addresses, associated networks, and
// verification status.
func (r *Web3WalletService) List(ctx context.Context, query Web3WalletListParams, opts ...option.RequestOption) (res *Web3WalletListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiates the process to securely connect a new cryptocurrency wallet to the
// user's profile, typically involving a signed message or OAuth flow from the
// wallet provider.
func (r *Web3WalletService) Connect(ctx context.Context, body Web3WalletConnectParams, opts ...option.RequestOption) (res *CryptoWalletConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the current balances of all recognized crypto assets within a specific
// connected wallet.
func (r *Web3WalletService) GetBalances(ctx context.Context, walletID interface{}, query Web3WalletGetBalancesParams, opts ...option.RequestOption) (res *Web3WalletGetBalancesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("web3/wallets/%v/balances", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CryptoWalletConnection struct {
	// Unique identifier for this wallet connection.
	ID interface{} `json:"id,required"`
	// The blockchain network this wallet is primarily connected to (e.g., Ethereum,
	// Solana, Polygon).
	BlockchainNetwork interface{} `json:"blockchainNetwork,required"`
	// Timestamp when the wallet's data was last synchronized.
	LastSynced interface{} `json:"lastSynced,required"`
	// Indicates if read access (balances, NFTs) is granted.
	ReadAccessGranted interface{} `json:"readAccessGranted,required"`
	// Current status of the wallet connection.
	Status CryptoWalletConnectionStatus `json:"status,required"`
	// Public address of the connected cryptocurrency wallet.
	WalletAddress interface{} `json:"walletAddress,required"`
	// Name of the wallet provider (e.g., MetaMask, Ledger, Phantom).
	WalletProvider interface{} `json:"walletProvider,required"`
	// Indicates if write access (transactions) is granted. Requires higher
	// permission/security.
	WriteAccessGranted interface{}                `json:"writeAccessGranted,required"`
	JSON               cryptoWalletConnectionJSON `json:"-"`
}

// cryptoWalletConnectionJSON contains the JSON metadata for the struct
// [CryptoWalletConnection]
type cryptoWalletConnectionJSON struct {
	ID                 apijson.Field
	BlockchainNetwork  apijson.Field
	LastSynced         apijson.Field
	ReadAccessGranted  apijson.Field
	Status             apijson.Field
	WalletAddress      apijson.Field
	WalletProvider     apijson.Field
	WriteAccessGranted apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CryptoWalletConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cryptoWalletConnectionJSON) RawJSON() string {
	return r.raw
}

// Current status of the wallet connection.
type CryptoWalletConnectionStatus string

const (
	CryptoWalletConnectionStatusConnected           CryptoWalletConnectionStatus = "connected"
	CryptoWalletConnectionStatusDisconnected        CryptoWalletConnectionStatus = "disconnected"
	CryptoWalletConnectionStatusPendingVerification CryptoWalletConnectionStatus = "pending_verification"
	CryptoWalletConnectionStatusError               CryptoWalletConnectionStatus = "error"
)

func (r CryptoWalletConnectionStatus) IsKnown() bool {
	switch r {
	case CryptoWalletConnectionStatusConnected, CryptoWalletConnectionStatusDisconnected, CryptoWalletConnectionStatusPendingVerification, CryptoWalletConnectionStatusError:
		return true
	}
	return false
}

type Web3WalletListResponse struct {
	Data []CryptoWalletConnection   `json:"data"`
	JSON web3WalletListResponseJSON `json:"-"`
	PaginatedList
}

// web3WalletListResponseJSON contains the JSON metadata for the struct
// [Web3WalletListResponse]
type web3WalletListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletListResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletGetBalancesResponse struct {
	Data []Web3WalletGetBalancesResponseData `json:"data"`
	JSON web3WalletGetBalancesResponseJSON   `json:"-"`
	PaginatedList
}

// web3WalletGetBalancesResponseJSON contains the JSON metadata for the struct
// [Web3WalletGetBalancesResponse]
type web3WalletGetBalancesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletGetBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletGetBalancesResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletGetBalancesResponseData struct {
	// Full name of the crypto asset.
	AssetName interface{} `json:"assetName,required"`
	// Symbol of the crypto asset (e.g., ETH, BTC, USDC).
	AssetSymbol interface{} `json:"assetSymbol,required"`
	// Current balance of the asset in the wallet.
	Balance interface{} `json:"balance,required"`
	// Current USD value of the asset balance.
	UsdValue interface{} `json:"usdValue,required"`
	// The contract address for ERC-20 tokens or similar.
	ContractAddress interface{} `json:"contractAddress"`
	// The blockchain network the asset resides on (if different from wallet's
	// primary).
	Network interface{}                           `json:"network"`
	JSON    web3WalletGetBalancesResponseDataJSON `json:"-"`
}

// web3WalletGetBalancesResponseDataJSON contains the JSON metadata for the struct
// [Web3WalletGetBalancesResponseData]
type web3WalletGetBalancesResponseDataJSON struct {
	AssetName       apijson.Field
	AssetSymbol     apijson.Field
	Balance         apijson.Field
	UsdValue        apijson.Field
	ContractAddress apijson.Field
	Network         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Web3WalletGetBalancesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletGetBalancesResponseDataJSON) RawJSON() string {
	return r.raw
}

type Web3WalletListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [Web3WalletListParams]'s query parameters as `url.Values`.
func (r Web3WalletListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type Web3WalletConnectParams struct {
	// The blockchain network for this wallet (e.g., Ethereum, Solana).
	BlockchainNetwork param.Field[interface{}] `json:"blockchainNetwork,required"`
	// A message cryptographically signed by the wallet owner to prove
	// ownership/intent.
	SignedMessage param.Field[interface{}] `json:"signedMessage,required"`
	// The public address of the cryptocurrency wallet.
	WalletAddress param.Field[interface{}] `json:"walletAddress,required"`
	// The name of the wallet provider (e.g., MetaMask, Phantom).
	WalletProvider param.Field[interface{}] `json:"walletProvider,required"`
	// If true, requests write access to initiate transactions from this wallet.
	RequestWriteAccess param.Field[interface{}] `json:"requestWriteAccess"`
}

func (r Web3WalletConnectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3WalletGetBalancesParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [Web3WalletGetBalancesParams]'s query parameters as
// `url.Values`.
func (r Web3WalletGetBalancesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
