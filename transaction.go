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

// TransactionService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options   []option.RequestOption
	Recurring *TransactionRecurringService
	Insights  *TransactionInsightService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.Recurring = NewTransactionRecurringService(opts...)
	r.Insights = NewTransactionInsightService(opts...)
	return
}

// Retrieves granular information for a single transaction by its unique ID,
// including AI categorization confidence, merchant details, and associated carbon
// footprint.
func (r *TransactionService) Get(ctx context.Context, transactionID interface{}, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("transactions/%v", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a paginated list of the user's transactions, with extensive options
// for filtering by type, category, date range, amount, and intelligent AI-driven
// sorting and search capabilities.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *PaginatedTransactions, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allows the user to override or refine the AI's categorization for a transaction,
// improving future AI accuracy and personal financial reporting.
func (r *TransactionService) Categorize(ctx context.Context, transactionID interface{}, body TransactionCategorizeParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("transactions/%v/categorize", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Begins the process of disputing a specific transaction, providing details and
// supporting documentation for review by our compliance team and AI.
func (r *TransactionService) Dispute(ctx context.Context, transactionID interface{}, body TransactionDisputeParams, opts ...option.RequestOption) (res *TransactionDisputeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("transactions/%v/dispute", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows the user to add or update personal notes for a specific transaction.
func (r *TransactionService) UpdateNotes(ctx context.Context, transactionID interface{}, body TransactionUpdateNotesParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("transactions/%v/notes", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PaginatedTransactions struct {
	Data []Transaction             `json:"data"`
	JSON paginatedTransactionsJSON `json:"-"`
	PaginatedList
}

// paginatedTransactionsJSON contains the JSON metadata for the struct
// [PaginatedTransactions]
type paginatedTransactionsJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedTransactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedTransactionsJSON) RawJSON() string {
	return r.raw
}

type Transaction struct {
	// Unique identifier for the transaction.
	ID interface{} `json:"id,required"`
	// ID of the account from which the transaction occurred.
	AccountID interface{} `json:"accountId,required"`
	// Amount of the transaction.
	Amount interface{} `json:"amount,required"`
	// AI-assigned or user-defined category of the transaction (e.g., 'Groceries',
	// 'Utilities').
	Category interface{} `json:"category,required"`
	// ISO 4217 currency code of the transaction.
	Currency interface{} `json:"currency,required"`
	// Date the transaction occurred (local date).
	Date interface{} `json:"date,required"`
	// Detailed description of the transaction.
	Description interface{} `json:"description,required"`
	// Type of the transaction.
	Type TransactionType `json:"type,required"`
	// AI confidence score for the assigned category (0-1).
	AICategoryConfidence interface{} `json:"aiCategoryConfidence"`
	// Estimated carbon footprint in kg CO2e for this transaction, derived by AI.
	CarbonFootprint interface{} `json:"carbonFootprint"`
	// Current dispute status of the transaction.
	DisputeStatus TransactionDisputeStatus `json:"disputeStatus"`
	// Geographic location details for a transaction.
	Location TransactionLocation `json:"location"`
	// Detailed information about a merchant associated with a transaction.
	MerchantDetails TransactionMerchantDetails `json:"merchantDetails"`
	// Personal notes added by the user to the transaction.
	Notes interface{} `json:"notes"`
	// Channel through which the payment was made.
	PaymentChannel TransactionPaymentChannel `json:"paymentChannel,nullable"`
	// Date the transaction was posted to the account (local date).
	PostedDate interface{} `json:"postedDate"`
	// URL to a digital receipt for the transaction.
	ReceiptURL interface{} `json:"receiptUrl"`
	// User-defined tags for the transaction.
	Tags []interface{}   `json:"tags,nullable"`
	JSON transactionJSON `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Amount               apijson.Field
	Category             apijson.Field
	Currency             apijson.Field
	Date                 apijson.Field
	Description          apijson.Field
	Type                 apijson.Field
	AICategoryConfidence apijson.Field
	CarbonFootprint      apijson.Field
	DisputeStatus        apijson.Field
	Location             apijson.Field
	MerchantDetails      apijson.Field
	Notes                apijson.Field
	PaymentChannel       apijson.Field
	PostedDate           apijson.Field
	ReceiptURL           apijson.Field
	Tags                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

// Type of the transaction.
type TransactionType string

const (
	TransactionTypeIncome      TransactionType = "income"
	TransactionTypeExpense     TransactionType = "expense"
	TransactionTypeTransfer    TransactionType = "transfer"
	TransactionTypeInvestment  TransactionType = "investment"
	TransactionTypeRefund      TransactionType = "refund"
	TransactionTypeBillPayment TransactionType = "bill_payment"
)

func (r TransactionType) IsKnown() bool {
	switch r {
	case TransactionTypeIncome, TransactionTypeExpense, TransactionTypeTransfer, TransactionTypeInvestment, TransactionTypeRefund, TransactionTypeBillPayment:
		return true
	}
	return false
}

// Current dispute status of the transaction.
type TransactionDisputeStatus string

const (
	TransactionDisputeStatusNone        TransactionDisputeStatus = "none"
	TransactionDisputeStatusPending     TransactionDisputeStatus = "pending"
	TransactionDisputeStatusUnderReview TransactionDisputeStatus = "under_review"
	TransactionDisputeStatusResolved    TransactionDisputeStatus = "resolved"
	TransactionDisputeStatusRejected    TransactionDisputeStatus = "rejected"
)

func (r TransactionDisputeStatus) IsKnown() bool {
	switch r {
	case TransactionDisputeStatusNone, TransactionDisputeStatusPending, TransactionDisputeStatusUnderReview, TransactionDisputeStatusResolved, TransactionDisputeStatusRejected:
		return true
	}
	return false
}

// Geographic location details for a transaction.
type TransactionLocation struct {
	// City where the transaction occurred.
	City interface{} `json:"city"`
	// Latitude coordinate of the transaction.
	Latitude interface{} `json:"latitude"`
	// Longitude coordinate of the transaction.
	Longitude interface{} `json:"longitude"`
	// State where the transaction occurred.
	State interface{} `json:"state"`
	// Zip code where the transaction occurred.
	Zip  interface{}             `json:"zip"`
	JSON transactionLocationJSON `json:"-"`
}

// transactionLocationJSON contains the JSON metadata for the struct
// [TransactionLocation]
type transactionLocationJSON struct {
	City        apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionLocationJSON) RawJSON() string {
	return r.raw
}

// Detailed information about a merchant associated with a transaction.
type TransactionMerchantDetails struct {
	Address Address `json:"address"`
	// URL to the merchant's logo.
	LogoURL interface{} `json:"logoUrl"`
	// Official name of the merchant.
	Name interface{} `json:"name"`
	// Merchant's phone number.
	Phone interface{} `json:"phone"`
	// Merchant's website URL.
	Website interface{}                    `json:"website"`
	JSON    transactionMerchantDetailsJSON `json:"-"`
}

// transactionMerchantDetailsJSON contains the JSON metadata for the struct
// [TransactionMerchantDetails]
type transactionMerchantDetailsJSON struct {
	Address     apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Phone       apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

// Channel through which the payment was made.
type TransactionPaymentChannel string

const (
	TransactionPaymentChannelInStore     TransactionPaymentChannel = "in_store"
	TransactionPaymentChannelOnline      TransactionPaymentChannel = "online"
	TransactionPaymentChannelMobile      TransactionPaymentChannel = "mobile"
	TransactionPaymentChannelAtm         TransactionPaymentChannel = "ATM"
	TransactionPaymentChannelBillPayment TransactionPaymentChannel = "bill_payment"
	TransactionPaymentChannelTransfer    TransactionPaymentChannel = "transfer"
	TransactionPaymentChannelOther       TransactionPaymentChannel = "other"
)

func (r TransactionPaymentChannel) IsKnown() bool {
	switch r {
	case TransactionPaymentChannelInStore, TransactionPaymentChannelOnline, TransactionPaymentChannelMobile, TransactionPaymentChannelAtm, TransactionPaymentChannelBillPayment, TransactionPaymentChannelTransfer, TransactionPaymentChannelOther:
		return true
	}
	return false
}

type TransactionDisputeResponse struct {
	// Unique identifier for the dispute case.
	DisputeID interface{} `json:"disputeId,required"`
	// Timestamp when the dispute status was last updated.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Guidance on what to expect next in the dispute process.
	NextSteps interface{} `json:"nextSteps,required"`
	// Current status of the dispute.
	Status TransactionDisputeResponseStatus `json:"status,required"`
	JSON   transactionDisputeResponseJSON   `json:"-"`
}

// transactionDisputeResponseJSON contains the JSON metadata for the struct
// [TransactionDisputeResponse]
type transactionDisputeResponseJSON struct {
	DisputeID   apijson.Field
	LastUpdated apijson.Field
	NextSteps   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionDisputeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDisputeResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the dispute.
type TransactionDisputeResponseStatus string

const (
	TransactionDisputeResponseStatusPending          TransactionDisputeResponseStatus = "pending"
	TransactionDisputeResponseStatusUnderReview      TransactionDisputeResponseStatus = "under_review"
	TransactionDisputeResponseStatusRequiresMoreInfo TransactionDisputeResponseStatus = "requires_more_info"
	TransactionDisputeResponseStatusResolved         TransactionDisputeResponseStatus = "resolved"
	TransactionDisputeResponseStatusRejected         TransactionDisputeResponseStatus = "rejected"
)

func (r TransactionDisputeResponseStatus) IsKnown() bool {
	switch r {
	case TransactionDisputeResponseStatusPending, TransactionDisputeResponseStatusUnderReview, TransactionDisputeResponseStatusRequiresMoreInfo, TransactionDisputeResponseStatusResolved, TransactionDisputeResponseStatusRejected:
		return true
	}
	return false
}

type TransactionListParams struct {
	// Filter transactions by their AI-assigned or user-defined category.
	Category param.Field[interface{}] `query:"category"`
	// Retrieve transactions up to this date (inclusive).
	EndDate param.Field[interface{}] `query:"endDate"`
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Filter for transactions with an amount less than or equal to this value.
	MaxAmount param.Field[interface{}] `query:"maxAmount"`
	// Filter for transactions with an amount greater than or equal to this value.
	MinAmount param.Field[interface{}] `query:"minAmount"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
	// Free-text search across transaction descriptions, merchants, and notes.
	SearchQuery param.Field[interface{}] `query:"searchQuery"`
	// Retrieve transactions from this date (inclusive).
	StartDate param.Field[interface{}] `query:"startDate"`
	// Filter transactions by type (e.g., income, expense, transfer).
	Type param.Field[TransactionListParamsType] `query:"type"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter transactions by type (e.g., income, expense, transfer).
type TransactionListParamsType string

const (
	TransactionListParamsTypeIncome      TransactionListParamsType = "income"
	TransactionListParamsTypeExpense     TransactionListParamsType = "expense"
	TransactionListParamsTypeTransfer    TransactionListParamsType = "transfer"
	TransactionListParamsTypeInvestment  TransactionListParamsType = "investment"
	TransactionListParamsTypeRefund      TransactionListParamsType = "refund"
	TransactionListParamsTypeBillPayment TransactionListParamsType = "bill_payment"
)

func (r TransactionListParamsType) IsKnown() bool {
	switch r {
	case TransactionListParamsTypeIncome, TransactionListParamsTypeExpense, TransactionListParamsTypeTransfer, TransactionListParamsTypeInvestment, TransactionListParamsTypeRefund, TransactionListParamsTypeBillPayment:
		return true
	}
	return false
}

type TransactionCategorizeParams struct {
	// The new category for the transaction. Can be hierarchical.
	Category param.Field[interface{}] `json:"category,required"`
	// If true, the AI will learn from this correction and try to apply it to similar
	// future transactions.
	ApplyToFuture param.Field[interface{}] `json:"applyToFuture"`
	// Optional notes to add to the transaction.
	Notes param.Field[interface{}] `json:"notes"`
}

func (r TransactionCategorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionDisputeParams struct {
	// Detailed explanation of the dispute.
	Details param.Field[interface{}] `json:"details,required"`
	// The primary reason for disputing the transaction.
	Reason param.Field[TransactionDisputeParamsReason] `json:"reason,required"`
	// URLs to supporting documents (e.g., receipts, communication).
	SupportingDocuments param.Field[[]interface{}] `json:"supportingDocuments"`
}

func (r TransactionDisputeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The primary reason for disputing the transaction.
type TransactionDisputeParamsReason string

const (
	TransactionDisputeParamsReasonUnauthorized        TransactionDisputeParamsReason = "unauthorized"
	TransactionDisputeParamsReasonDuplicateCharge     TransactionDisputeParamsReason = "duplicate_charge"
	TransactionDisputeParamsReasonIncorrectAmount     TransactionDisputeParamsReason = "incorrect_amount"
	TransactionDisputeParamsReasonProductServiceIssue TransactionDisputeParamsReason = "product_service_issue"
	TransactionDisputeParamsReasonOther               TransactionDisputeParamsReason = "other"
)

func (r TransactionDisputeParamsReason) IsKnown() bool {
	switch r {
	case TransactionDisputeParamsReasonUnauthorized, TransactionDisputeParamsReasonDuplicateCharge, TransactionDisputeParamsReasonIncorrectAmount, TransactionDisputeParamsReasonProductServiceIssue, TransactionDisputeParamsReasonOther:
		return true
	}
	return false
}

type TransactionUpdateNotesParams struct {
	// The personal notes to add or update for the transaction.
	Notes param.Field[interface{}] `json:"notes,required"`
}

func (r TransactionUpdateNotesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
