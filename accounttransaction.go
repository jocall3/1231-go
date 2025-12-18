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

// AccountTransactionService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTransactionService] method instead.
type AccountTransactionService struct {
	Options []option.RequestOption
}

// NewAccountTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTransactionService(opts ...option.RequestOption) (r *AccountTransactionService) {
	r = &AccountTransactionService{}
	r.Options = opts
	return
}

// Retrieves a list of pending transactions that have not yet cleared for a
// specific financial account.
func (r *AccountTransactionService) GetPending(ctx context.Context, accountID interface{}, query AccountTransactionGetPendingParams, opts ...option.RequestOption) (res *AccountTransactionGetPendingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("accounts/%v/transactions/pending", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTransactionGetPendingResponse struct {
	Data []Transaction                            `json:"data"`
	JSON accountTransactionGetPendingResponseJSON `json:"-"`
	PaginatedList
}

// accountTransactionGetPendingResponseJSON contains the JSON metadata for the
// struct [AccountTransactionGetPendingResponse]
type accountTransactionGetPendingResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransactionGetPendingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionGetPendingResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionGetPendingParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [AccountTransactionGetPendingParams]'s query parameters as
// `url.Values`.
func (r AccountTransactionGetPendingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
