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

// CorporateCardService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateCardService] method instead.
type CorporateCardService struct {
	Options []option.RequestOption
}

// NewCorporateCardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateCardService(opts ...option.RequestOption) (r *CorporateCardService) {
	r = &CorporateCardService{}
	r.Options = opts
	return
}

// Retrieves a comprehensive list of all physical and virtual corporate cards
// associated with the user's organization, including their status, assigned
// holder, and current spending controls.
func (r *CorporateCardService) List(ctx context.Context, query CorporateCardListParams, opts ...option.RequestOption) (res *CorporateCardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Creates and issues a new virtual corporate card with specified spending limits,
// merchant restrictions, and expiration dates, ideal for secure online purchases
// and temporary projects.
func (r *CorporateCardService) NewVirtual(ctx context.Context, body CorporateCardNewVirtualParams, opts ...option.RequestOption) (res *CorporateCard, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards/virtual"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Immediately changes the frozen status of a corporate card, preventing or
// allowing transactions in real-time, critical for security and expense
// management.
func (r *CorporateCardService) Freeze(ctx context.Context, cardID interface{}, body CorporateCardFreezeParams, opts ...option.RequestOption) (res *CorporateCard, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("corporate/cards/%v/freeze", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of transactions made with a specific corporate card,
// including AI categorization and compliance flags.
func (r *CorporateCardService) ListTransactions(ctx context.Context, cardID interface{}, query CorporateCardListTransactionsParams, opts ...option.RequestOption) (res *PaginatedTransactions, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("corporate/cards/%v/transactions", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the sophisticated spending controls, limits, and policy overrides for a
// specific corporate card, enabling real-time adjustments for security and budget
// adherence.
func (r *CorporateCardService) UpdateControls(ctx context.Context, cardID interface{}, body CorporateCardUpdateControlsParams, opts ...option.RequestOption) (res *CorporateCard, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("corporate/cards/%v/controls", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CorporateCard struct {
	// Unique identifier for the corporate card.
	ID interface{} `json:"id,required"`
	// Masked card number for display purposes.
	CardNumberMask interface{} `json:"cardNumberMask,required"`
	// Type of the card (physical or virtual).
	CardType CorporateCardCardType `json:"cardType,required"`
	// Granular spending controls for a corporate card.
	Controls CorporateCardControls `json:"controls,required"`
	// Timestamp when the card was created.
	CreatedDate interface{} `json:"createdDate,required"`
	// Currency of the card's limits and transactions.
	Currency interface{} `json:"currency,required"`
	// Expiration date of the card (YYYY-MM-DD).
	ExpirationDate interface{} `json:"expirationDate,required"`
	// If true, the card is temporarily frozen and cannot be used.
	Frozen interface{} `json:"frozen,required"`
	// Name of the card holder.
	HolderName interface{} `json:"holderName,required"`
	// Current status of the card.
	Status CorporateCardStatus `json:"status,required"`
	// Optional: ID of the employee associated with this card.
	AssociatedEmployeeID interface{} `json:"associatedEmployeeId"`
	// Optional: ID of the overarching spending policy applied to this card.
	SpendingPolicyID interface{}       `json:"spendingPolicyId"`
	JSON             corporateCardJSON `json:"-"`
}

// corporateCardJSON contains the JSON metadata for the struct [CorporateCard]
type corporateCardJSON struct {
	ID                   apijson.Field
	CardNumberMask       apijson.Field
	CardType             apijson.Field
	Controls             apijson.Field
	CreatedDate          apijson.Field
	Currency             apijson.Field
	ExpirationDate       apijson.Field
	Frozen               apijson.Field
	HolderName           apijson.Field
	Status               apijson.Field
	AssociatedEmployeeID apijson.Field
	SpendingPolicyID     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CorporateCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardJSON) RawJSON() string {
	return r.raw
}

// Type of the card (physical or virtual).
type CorporateCardCardType string

const (
	CorporateCardCardTypePhysical CorporateCardCardType = "physical"
	CorporateCardCardTypeVirtual  CorporateCardCardType = "virtual"
)

func (r CorporateCardCardType) IsKnown() bool {
	switch r {
	case CorporateCardCardTypePhysical, CorporateCardCardTypeVirtual:
		return true
	}
	return false
}

// Current status of the card.
type CorporateCardStatus string

const (
	CorporateCardStatusActive            CorporateCardStatus = "Active"
	CorporateCardStatusSuspended         CorporateCardStatus = "Suspended"
	CorporateCardStatusDeactivated       CorporateCardStatus = "Deactivated"
	CorporateCardStatusPendingActivation CorporateCardStatus = "Pending Activation"
)

func (r CorporateCardStatus) IsKnown() bool {
	switch r {
	case CorporateCardStatusActive, CorporateCardStatusSuspended, CorporateCardStatusDeactivated, CorporateCardStatusPendingActivation:
		return true
	}
	return false
}

// Granular spending controls for a corporate card.
type CorporateCardControls struct {
	// If true, ATM cash withdrawals are allowed.
	AtmWithdrawals interface{} `json:"atmWithdrawals"`
	// If true, contactless payments are allowed.
	ContactlessPayments interface{} `json:"contactlessPayments"`
	// Maximum spending limit per day (null for no limit).
	DailyLimit interface{} `json:"dailyLimit"`
	// If true, international transactions are allowed.
	InternationalTransactions interface{} `json:"internationalTransactions"`
	// List of allowed merchant categories. If empty, all are allowed unless explicitly
	// denied.
	MerchantCategoryRestrictions []interface{} `json:"merchantCategoryRestrictions,nullable"`
	// Maximum spending limit per month (null for no limit).
	MonthlyLimit interface{} `json:"monthlyLimit"`
	// If true, online transactions are allowed.
	OnlineTransactions interface{} `json:"onlineTransactions"`
	// Maximum amount for a single transaction (null for no limit).
	SingleTransactionLimit interface{} `json:"singleTransactionLimit"`
	// List of allowed vendors/merchants by name.
	VendorRestrictions []interface{}             `json:"vendorRestrictions,nullable"`
	JSON               corporateCardControlsJSON `json:"-"`
}

// corporateCardControlsJSON contains the JSON metadata for the struct
// [CorporateCardControls]
type corporateCardControlsJSON struct {
	AtmWithdrawals               apijson.Field
	ContactlessPayments          apijson.Field
	DailyLimit                   apijson.Field
	InternationalTransactions    apijson.Field
	MerchantCategoryRestrictions apijson.Field
	MonthlyLimit                 apijson.Field
	OnlineTransactions           apijson.Field
	SingleTransactionLimit       apijson.Field
	VendorRestrictions           apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CorporateCardControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardControlsJSON) RawJSON() string {
	return r.raw
}

// Granular spending controls for a corporate card.
type CorporateCardControlsParam struct {
	// If true, ATM cash withdrawals are allowed.
	AtmWithdrawals param.Field[interface{}] `json:"atmWithdrawals"`
	// If true, contactless payments are allowed.
	ContactlessPayments param.Field[interface{}] `json:"contactlessPayments"`
	// Maximum spending limit per day (null for no limit).
	DailyLimit param.Field[interface{}] `json:"dailyLimit"`
	// If true, international transactions are allowed.
	InternationalTransactions param.Field[interface{}] `json:"internationalTransactions"`
	// List of allowed merchant categories. If empty, all are allowed unless explicitly
	// denied.
	MerchantCategoryRestrictions param.Field[[]interface{}] `json:"merchantCategoryRestrictions"`
	// Maximum spending limit per month (null for no limit).
	MonthlyLimit param.Field[interface{}] `json:"monthlyLimit"`
	// If true, online transactions are allowed.
	OnlineTransactions param.Field[interface{}] `json:"onlineTransactions"`
	// Maximum amount for a single transaction (null for no limit).
	SingleTransactionLimit param.Field[interface{}] `json:"singleTransactionLimit"`
	// List of allowed vendors/merchants by name.
	VendorRestrictions param.Field[[]interface{}] `json:"vendorRestrictions"`
}

func (r CorporateCardControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardListResponse struct {
	Data []CorporateCard               `json:"data"`
	JSON corporateCardListResponseJSON `json:"-"`
	PaginatedList
}

// corporateCardListResponseJSON contains the JSON metadata for the struct
// [CorporateCardListResponse]
type corporateCardListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateCardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateCardListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [CorporateCardListParams]'s query parameters as
// `url.Values`.
func (r CorporateCardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateCardNewVirtualParams struct {
	// Granular spending controls for a corporate card.
	Controls param.Field[CorporateCardControlsParam] `json:"controls,required"`
	// Expiration date for the virtual card (YYYY-MM-DD).
	ExpirationDate param.Field[interface{}] `json:"expirationDate,required"`
	// Name to appear on the virtual card.
	HolderName param.Field[interface{}] `json:"holderName,required"`
	// Brief description of the virtual card's purpose.
	Purpose param.Field[interface{}] `json:"purpose,required"`
	// Optional: ID of the employee or department this card is for.
	AssociatedEmployeeID param.Field[interface{}] `json:"associatedEmployeeId"`
	// Optional: ID of a spending policy to link with this virtual card.
	SpendingPolicyID param.Field[interface{}] `json:"spendingPolicyId"`
}

func (r CorporateCardNewVirtualParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardFreezeParams struct {
	// Set to `true` to freeze the card, `false` to unfreeze.
	Freeze param.Field[interface{}] `json:"freeze,required"`
}

func (r CorporateCardFreezeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardListTransactionsParams struct {
	// End date for filtering results (inclusive, YYYY-MM-DD).
	EndDate param.Field[interface{}] `query:"endDate"`
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
	// Start date for filtering results (inclusive, YYYY-MM-DD).
	StartDate param.Field[interface{}] `query:"startDate"`
}

// URLQuery serializes [CorporateCardListTransactionsParams]'s query parameters as
// `url.Values`.
func (r CorporateCardListTransactionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateCardUpdateControlsParams struct {
	// Granular spending controls for a corporate card.
	CorporateCardControls CorporateCardControlsParam `json:"corporate_card_controls,required"`
}

func (r CorporateCardUpdateControlsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CorporateCardControls)
}
