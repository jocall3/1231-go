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

// AccountService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options           []option.RequestOption
	Transactions      *AccountTransactionService
	OverdraftSettings *AccountOverdraftSettingService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Transactions = NewAccountTransactionService(opts...)
	r.OverdraftSettings = NewAccountOverdraftSettingService(opts...)
	return
}

// Begins the secure process of linking a new external financial institution (e.g.,
// another bank, investment platform) to the user's profile, typically involving a
// third-party tokenized flow.
func (r *AccountService) Link(ctx context.Context, body AccountLinkParams, opts ...option.RequestOption) (res *AccountLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves comprehensive analytics for a specific financial account, including
// historical balance trends, projected cash flow, and AI-driven insights into
// spending patterns.
func (r *AccountService) GetDetails(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *AccountGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("accounts/%v/details", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a comprehensive, real-time list of all external financial accounts
// linked to the user's profile, including consolidated balances and institutional
// details.
func (r *AccountService) GetMe(ctx context.Context, query AccountGetMeParams, opts ...option.RequestOption) (res *AccountGetMeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetches digital statements for a specific account, allowing filtering by date
// range and format.
func (r *AccountService) GetStatements(ctx context.Context, accountID interface{}, query AccountGetStatementsParams, opts ...option.RequestOption) (res *AccountGetStatementsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("accounts/%v/statements", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Summary information for a linked financial account.
type LinkedAccount struct {
	// Unique identifier for the linked account within .
	ID interface{} `json:"id,required"`
	// ISO 4217 currency code of the account.
	Currency interface{} `json:"currency,required"`
	// Current balance of the account.
	CurrentBalance interface{} `json:"currentBalance,required"`
	// Name of the financial institution where the account is held.
	InstitutionName interface{} `json:"institutionName,required"`
	// Timestamp when the account balance was last synced.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Display name of the account.
	Name interface{} `json:"name,required"`
	// General type of the account.
	Type LinkedAccountType `json:"type,required"`
	// Available balance (after pending transactions) of the account.
	AvailableBalance interface{} `json:"availableBalance"`
	// Optional: Identifier from the external data provider (e.g., Plaid).
	ExternalID interface{} `json:"externalId"`
	// Masked account number (e.g., last 4 digits).
	Mask interface{} `json:"mask"`
	// Specific subtype of the account (e.g., checking, savings, IRA, 401k).
	Subtype interface{}       `json:"subtype"`
	JSON    linkedAccountJSON `json:"-"`
}

// linkedAccountJSON contains the JSON metadata for the struct [LinkedAccount]
type linkedAccountJSON struct {
	ID               apijson.Field
	Currency         apijson.Field
	CurrentBalance   apijson.Field
	InstitutionName  apijson.Field
	LastUpdated      apijson.Field
	Name             apijson.Field
	Type             apijson.Field
	AvailableBalance apijson.Field
	ExternalID       apijson.Field
	Mask             apijson.Field
	Subtype          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LinkedAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedAccountJSON) RawJSON() string {
	return r.raw
}

// General type of the account.
type LinkedAccountType string

const (
	LinkedAccountTypeDepository LinkedAccountType = "depository"
	LinkedAccountTypeCredit     LinkedAccountType = "credit"
	LinkedAccountTypeLoan       LinkedAccountType = "loan"
	LinkedAccountTypeInvestment LinkedAccountType = "investment"
	LinkedAccountTypeOther      LinkedAccountType = "other"
)

func (r LinkedAccountType) IsKnown() bool {
	switch r {
	case LinkedAccountTypeDepository, LinkedAccountTypeCredit, LinkedAccountTypeLoan, LinkedAccountTypeInvestment, LinkedAccountTypeOther:
		return true
	}
	return false
}

type AccountLinkResponse struct {
	// The URI to redirect the user to complete authentication with the external
	// institution.
	AuthUri interface{} `json:"authUri,required"`
	// Unique session ID for the account linking process.
	LinkSessionID interface{} `json:"linkSessionId,required"`
	// Current status of the linking process.
	Status AccountLinkResponseStatus `json:"status,required"`
	// A descriptive message regarding the next steps.
	Message interface{}             `json:"message"`
	JSON    accountLinkResponseJSON `json:"-"`
}

// accountLinkResponseJSON contains the JSON metadata for the struct
// [AccountLinkResponse]
type accountLinkResponseJSON struct {
	AuthUri       apijson.Field
	LinkSessionID apijson.Field
	Status        apijson.Field
	Message       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the linking process.
type AccountLinkResponseStatus string

const (
	AccountLinkResponseStatusPendingUserAction AccountLinkResponseStatus = "pending_user_action"
	AccountLinkResponseStatusCompleted         AccountLinkResponseStatus = "completed"
	AccountLinkResponseStatusFailed            AccountLinkResponseStatus = "failed"
)

func (r AccountLinkResponseStatus) IsKnown() bool {
	switch r {
	case AccountLinkResponseStatusPendingUserAction, AccountLinkResponseStatusCompleted, AccountLinkResponseStatusFailed:
		return true
	}
	return false
}

// Summary information for a linked financial account.
type AccountGetDetailsResponse struct {
	// Name of the primary holder for this account.
	AccountHolder interface{} `json:"accountHolder"`
	// Historical daily balance data.
	BalanceHistory []AccountGetDetailsResponseBalanceHistory `json:"balanceHistory"`
	// Annual interest rate (if applicable).
	InterestRate interface{} `json:"interestRate"`
	// Date the account was opened.
	OpenedDate        interface{}                                `json:"openedDate"`
	ProjectedCashFlow AccountGetDetailsResponseProjectedCashFlow `json:"projectedCashFlow"`
	// Total number of transactions in this account.
	TransactionsCount interface{}                   `json:"transactionsCount"`
	JSON              accountGetDetailsResponseJSON `json:"-"`
	LinkedAccount
}

// accountGetDetailsResponseJSON contains the JSON metadata for the struct
// [AccountGetDetailsResponse]
type accountGetDetailsResponseJSON struct {
	AccountHolder     apijson.Field
	BalanceHistory    apijson.Field
	InterestRate      apijson.Field
	OpenedDate        apijson.Field
	ProjectedCashFlow apijson.Field
	TransactionsCount apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetDetailsResponseBalanceHistory struct {
	Balance interface{}                                 `json:"balance"`
	Date    interface{}                                 `json:"date"`
	JSON    accountGetDetailsResponseBalanceHistoryJSON `json:"-"`
}

// accountGetDetailsResponseBalanceHistoryJSON contains the JSON metadata for the
// struct [AccountGetDetailsResponseBalanceHistory]
type accountGetDetailsResponseBalanceHistoryJSON struct {
	Balance     apijson.Field
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetDetailsResponseBalanceHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetDetailsResponseBalanceHistoryJSON) RawJSON() string {
	return r.raw
}

type AccountGetDetailsResponseProjectedCashFlow struct {
	// AI confidence score for the cash flow projection (0-100).
	ConfidenceScore interface{} `json:"confidenceScore"`
	// Projected cash flow for the next 30 days.
	Days30 interface{} `json:"days30"`
	// Projected cash flow for the next 90 days.
	Days90 interface{}                                    `json:"days90"`
	JSON   accountGetDetailsResponseProjectedCashFlowJSON `json:"-"`
}

// accountGetDetailsResponseProjectedCashFlowJSON contains the JSON metadata for
// the struct [AccountGetDetailsResponseProjectedCashFlow]
type accountGetDetailsResponseProjectedCashFlowJSON struct {
	ConfidenceScore apijson.Field
	Days30          apijson.Field
	Days90          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountGetDetailsResponseProjectedCashFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetDetailsResponseProjectedCashFlowJSON) RawJSON() string {
	return r.raw
}

type AccountGetMeResponse struct {
	Data []LinkedAccount          `json:"data"`
	JSON accountGetMeResponseJSON `json:"-"`
	PaginatedList
}

// accountGetMeResponseJSON contains the JSON metadata for the struct
// [AccountGetMeResponse]
type accountGetMeResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetMeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetStatementsResponse struct {
	// The account ID the statement belongs to.
	AccountID interface{} `json:"accountId,required"`
	// Map of available download URLs for different formats.
	DownloadURLs AccountGetStatementsResponseDownloadURLs `json:"downloadUrls,required"`
	// The period covered by the statement.
	Period interface{} `json:"period,required"`
	// Unique identifier for the statement.
	StatementID interface{}                      `json:"statementId,required"`
	JSON        accountGetStatementsResponseJSON `json:"-"`
}

// accountGetStatementsResponseJSON contains the JSON metadata for the struct
// [AccountGetStatementsResponse]
type accountGetStatementsResponseJSON struct {
	AccountID    apijson.Field
	DownloadURLs apijson.Field
	Period       apijson.Field
	StatementID  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGetStatementsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetStatementsResponseJSON) RawJSON() string {
	return r.raw
}

// Map of available download URLs for different formats.
type AccountGetStatementsResponseDownloadURLs struct {
	// Signed URL to download the statement in CSV format.
	Csv interface{} `json:"csv"`
	// Signed URL to download the statement in PDF format.
	Pdf  interface{}                                  `json:"pdf"`
	JSON accountGetStatementsResponseDownloadURLsJSON `json:"-"`
}

// accountGetStatementsResponseDownloadURLsJSON contains the JSON metadata for the
// struct [AccountGetStatementsResponseDownloadURLs]
type accountGetStatementsResponseDownloadURLsJSON struct {
	Csv         apijson.Field
	Pdf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetStatementsResponseDownloadURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetStatementsResponseDownloadURLsJSON) RawJSON() string {
	return r.raw
}

type AccountLinkParams struct {
	// Two-letter ISO country code of the institution.
	CountryCode param.Field[interface{}] `json:"countryCode,required"`
	// Name of the financial institution to link.
	InstitutionName param.Field[interface{}] `json:"institutionName,required"`
	// Optional: Specific identifier for a third-party linking provider (e.g., 'plaid',
	// 'finicity').
	ProviderIdentifier param.Field[interface{}] `json:"providerIdentifier"`
	// Optional: URI to redirect the user after completing the external authentication
	// flow.
	RedirectUri param.Field[interface{}] `json:"redirectUri"`
}

func (r AccountLinkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGetMeParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [AccountGetMeParams]'s query parameters as `url.Values`.
func (r AccountGetMeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountGetStatementsParams struct {
	// Month for the statement (1-12).
	Month param.Field[interface{}] `query:"month,required"`
	// Year for the statement.
	Year param.Field[interface{}] `query:"year,required"`
	// Desired format for the statement. Use 'application/json' Accept header for
	// download links.
	Format param.Field[AccountGetStatementsParamsFormat] `query:"format"`
}

// URLQuery serializes [AccountGetStatementsParams]'s query parameters as
// `url.Values`.
func (r AccountGetStatementsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Desired format for the statement. Use 'application/json' Accept header for
// download links.
type AccountGetStatementsParamsFormat string

const (
	AccountGetStatementsParamsFormatPdf AccountGetStatementsParamsFormat = "pdf"
	AccountGetStatementsParamsFormatCsv AccountGetStatementsParamsFormat = "csv"
)

func (r AccountGetStatementsParamsFormat) IsKnown() bool {
	switch r {
	case AccountGetStatementsParamsFormatPdf, AccountGetStatementsParamsFormatCsv:
		return true
	}
	return false
}
