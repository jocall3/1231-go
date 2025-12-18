// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// Web3Service contains methods and other services that help with interacting with
// the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3Service] method instead.
type Web3Service struct {
	Options      []option.RequestOption
	Wallets      *Web3WalletService
	Transactions *Web3TransactionService
}

// NewWeb3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWeb3Service(opts ...option.RequestOption) (r *Web3Service) {
	r = &Web3Service{}
	r.Options = opts
	r.Wallets = NewWeb3WalletService(opts...)
	r.Transactions = NewWeb3TransactionService(opts...)
	return
}

// Fetches a comprehensive list of Non-Fungible Tokens (NFTs) owned by the user
// across all connected wallets and supported blockchain networks, including
// metadata and market values.
func (r *Web3Service) GetNFTs(ctx context.Context, query Web3GetNFTsParams, opts ...option.RequestOption) (res *Web3GetNFTsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/nfts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Web3GetNFTsResponse struct {
	Data []Web3GetNFTsResponseData `json:"data"`
	JSON web3GetNFTsResponseJSON   `json:"-"`
	PaginatedList
}

// web3GetNFTsResponseJSON contains the JSON metadata for the struct
// [Web3GetNFTsResponse]
type web3GetNFTsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3GetNFTsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3GetNFTsResponseJSON) RawJSON() string {
	return r.raw
}

type Web3GetNFTsResponseData struct {
	// Unique identifier for the NFT within the system.
	ID interface{} `json:"id,required"`
	// Blockchain network on which the NFT exists.
	BlockchainNetwork interface{} `json:"blockchainNetwork,required"`
	// Name of the NFT collection.
	CollectionName interface{} `json:"collectionName,required"`
	// Blockchain contract address of the NFT collection.
	ContractAddress interface{} `json:"contractAddress,required"`
	// URL to the NFT's image.
	ImageURL interface{} `json:"imageUrl,required"`
	// Name of the specific NFT.
	Name interface{} `json:"name,required"`
	// Blockchain address of the current owner.
	OwnerAddress interface{} `json:"ownerAddress,required"`
	// Unique ID of the token within its contract.
	TokenID interface{} `json:"tokenId,required"`
	// Key-value attributes of the NFT (e.g., rarity traits).
	Attributes []Web3GetNFTsResponseDataAttribute `json:"attributes,nullable"`
	// Description of the NFT.
	Description interface{} `json:"description"`
	// AI-estimated current market value in USD.
	EstimatedValueUsd interface{} `json:"estimatedValueUSD"`
	// Last known sale price in USD.
	LastSalePriceUsd interface{}                 `json:"lastSalePriceUSD"`
	JSON             web3GetNFTsResponseDataJSON `json:"-"`
}

// web3GetNFTsResponseDataJSON contains the JSON metadata for the struct
// [Web3GetNFTsResponseData]
type web3GetNFTsResponseDataJSON struct {
	ID                apijson.Field
	BlockchainNetwork apijson.Field
	CollectionName    apijson.Field
	ContractAddress   apijson.Field
	ImageURL          apijson.Field
	Name              apijson.Field
	OwnerAddress      apijson.Field
	TokenID           apijson.Field
	Attributes        apijson.Field
	Description       apijson.Field
	EstimatedValueUsd apijson.Field
	LastSalePriceUsd  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Web3GetNFTsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3GetNFTsResponseDataJSON) RawJSON() string {
	return r.raw
}

type Web3GetNFTsResponseDataAttribute struct {
	TraitType interface{}                          `json:"trait_type"`
	Value     interface{}                          `json:"value"`
	JSON      web3GetNFTsResponseDataAttributeJSON `json:"-"`
}

// web3GetNFTsResponseDataAttributeJSON contains the JSON metadata for the struct
// [Web3GetNFTsResponseDataAttribute]
type web3GetNFTsResponseDataAttributeJSON struct {
	TraitType   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3GetNFTsResponseDataAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3GetNFTsResponseDataAttributeJSON) RawJSON() string {
	return r.raw
}

type Web3GetNFTsParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [Web3GetNFTsParams]'s query parameters as `url.Values`.
func (r Web3GetNFTsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
