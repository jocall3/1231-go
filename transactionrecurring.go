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

// TransactionRecurringService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionRecurringService] method instead.
type TransactionRecurringService struct {
	Options []option.RequestOption
}

// NewTransactionRecurringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionRecurringService(opts ...option.RequestOption) (r *TransactionRecurringService) {
	r = &TransactionRecurringService{}
	r.Options = opts
	return
}

// Defines a new recurring transaction pattern for future tracking and budgeting.
func (r *TransactionRecurringService) New(ctx context.Context, body TransactionRecurringNewParams, opts ...option.RequestOption) (res *RecurringTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a list of all detected or user-defined recurring transactions, useful
// for budget tracking and subscription management.
func (r *TransactionRecurringService) List(ctx context.Context, query TransactionRecurringListParams, opts ...option.RequestOption) (res *TransactionRecurringListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details of a detected or user-defined recurring transaction.
type RecurringTransaction struct {
	// Unique identifier for the recurring transaction.
	ID interface{} `json:"id,required"`
	// Amount of the recurring transaction.
	Amount interface{} `json:"amount,required"`
	// Category of the recurring transaction.
	Category interface{} `json:"category,required"`
	// ISO 4217 currency code.
	Currency interface{} `json:"currency,required"`
	// Description of the recurring transaction.
	Description interface{} `json:"description,required"`
	// Frequency of the recurring transaction.
	Frequency RecurringTransactionFrequency `json:"frequency,required"`
	// Current status of the recurring transaction.
	Status RecurringTransactionStatus `json:"status,required"`
	// AI confidence score that this is a recurring transaction (0-1).
	AIConfidenceScore interface{} `json:"aiConfidenceScore"`
	// Date of the last payment for this recurring transaction.
	LastPaidDate interface{} `json:"lastPaidDate"`
	// ID of the account typically used for this recurring transaction.
	LinkedAccountID interface{} `json:"linkedAccountId"`
	// Next scheduled due date for the transaction.
	NextDueDate interface{}              `json:"nextDueDate"`
	JSON        recurringTransactionJSON `json:"-"`
}

// recurringTransactionJSON contains the JSON metadata for the struct
// [RecurringTransaction]
type recurringTransactionJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	Category          apijson.Field
	Currency          apijson.Field
	Description       apijson.Field
	Frequency         apijson.Field
	Status            apijson.Field
	AIConfidenceScore apijson.Field
	LastPaidDate      apijson.Field
	LinkedAccountID   apijson.Field
	NextDueDate       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecurringTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recurringTransactionJSON) RawJSON() string {
	return r.raw
}

// Frequency of the recurring transaction.
type RecurringTransactionFrequency string

const (
	RecurringTransactionFrequencyDaily        RecurringTransactionFrequency = "daily"
	RecurringTransactionFrequencyWeekly       RecurringTransactionFrequency = "weekly"
	RecurringTransactionFrequencyBiWeekly     RecurringTransactionFrequency = "bi_weekly"
	RecurringTransactionFrequencyMonthly      RecurringTransactionFrequency = "monthly"
	RecurringTransactionFrequencyQuarterly    RecurringTransactionFrequency = "quarterly"
	RecurringTransactionFrequencySemiAnnually RecurringTransactionFrequency = "semi_annually"
	RecurringTransactionFrequencyAnnually     RecurringTransactionFrequency = "annually"
)

func (r RecurringTransactionFrequency) IsKnown() bool {
	switch r {
	case RecurringTransactionFrequencyDaily, RecurringTransactionFrequencyWeekly, RecurringTransactionFrequencyBiWeekly, RecurringTransactionFrequencyMonthly, RecurringTransactionFrequencyQuarterly, RecurringTransactionFrequencySemiAnnually, RecurringTransactionFrequencyAnnually:
		return true
	}
	return false
}

// Current status of the recurring transaction.
type RecurringTransactionStatus string

const (
	RecurringTransactionStatusActive    RecurringTransactionStatus = "active"
	RecurringTransactionStatusInactive  RecurringTransactionStatus = "inactive"
	RecurringTransactionStatusCancelled RecurringTransactionStatus = "cancelled"
	RecurringTransactionStatusPaused    RecurringTransactionStatus = "paused"
)

func (r RecurringTransactionStatus) IsKnown() bool {
	switch r {
	case RecurringTransactionStatusActive, RecurringTransactionStatusInactive, RecurringTransactionStatusCancelled, RecurringTransactionStatusPaused:
		return true
	}
	return false
}

type TransactionRecurringListResponse struct {
	Data []RecurringTransaction               `json:"data"`
	JSON transactionRecurringListResponseJSON `json:"-"`
	PaginatedList
}

// transactionRecurringListResponseJSON contains the JSON metadata for the struct
// [TransactionRecurringListResponse]
type transactionRecurringListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionRecurringListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionRecurringListResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionRecurringNewParams struct {
	// Amount of the recurring transaction.
	Amount param.Field[interface{}] `json:"amount,required"`
	// Category of the recurring transaction.
	Category param.Field[interface{}] `json:"category,required"`
	// ISO 4217 currency code.
	Currency param.Field[interface{}] `json:"currency,required"`
	// Description of the recurring transaction.
	Description param.Field[interface{}] `json:"description,required"`
	// Frequency of the recurring transaction.
	Frequency param.Field[TransactionRecurringNewParamsFrequency] `json:"frequency,required"`
	// ID of the account to associate with this recurring transaction.
	LinkedAccountID param.Field[interface{}] `json:"linkedAccountId,required"`
	// The date when this recurring transaction is expected to start.
	StartDate param.Field[interface{}] `json:"startDate,required"`
}

func (r TransactionRecurringNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Frequency of the recurring transaction.
type TransactionRecurringNewParamsFrequency string

const (
	TransactionRecurringNewParamsFrequencyDaily        TransactionRecurringNewParamsFrequency = "daily"
	TransactionRecurringNewParamsFrequencyWeekly       TransactionRecurringNewParamsFrequency = "weekly"
	TransactionRecurringNewParamsFrequencyBiWeekly     TransactionRecurringNewParamsFrequency = "bi_weekly"
	TransactionRecurringNewParamsFrequencyMonthly      TransactionRecurringNewParamsFrequency = "monthly"
	TransactionRecurringNewParamsFrequencyQuarterly    TransactionRecurringNewParamsFrequency = "quarterly"
	TransactionRecurringNewParamsFrequencySemiAnnually TransactionRecurringNewParamsFrequency = "semi_annually"
	TransactionRecurringNewParamsFrequencyAnnually     TransactionRecurringNewParamsFrequency = "annually"
)

func (r TransactionRecurringNewParamsFrequency) IsKnown() bool {
	switch r {
	case TransactionRecurringNewParamsFrequencyDaily, TransactionRecurringNewParamsFrequencyWeekly, TransactionRecurringNewParamsFrequencyBiWeekly, TransactionRecurringNewParamsFrequencyMonthly, TransactionRecurringNewParamsFrequencyQuarterly, TransactionRecurringNewParamsFrequencySemiAnnually, TransactionRecurringNewParamsFrequencyAnnually:
		return true
	}
	return false
}

type TransactionRecurringListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [TransactionRecurringListParams]'s query parameters as
// `url.Values`.
func (r TransactionRecurringListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
