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

// BudgetService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBudgetService] method instead.
type BudgetService struct {
	Options []option.RequestOption
}

// NewBudgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBudgetService(opts ...option.RequestOption) (r *BudgetService) {
	r = &BudgetService{}
	r.Options = opts
	return
}

// Creates a new financial budget for the user, with optional AI auto-population of
// categories and amounts.
func (r *BudgetService) New(ctx context.Context, body BudgetNewParams, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information for a specific budget, including current
// spending, remaining amounts, and AI recommendations.
func (r *BudgetService) Get(ctx context.Context, budgetID interface{}, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("budgets/%v", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the parameters of an existing budget, such as total amount, dates, or
// categories.
func (r *BudgetService) Update(ctx context.Context, budgetID interface{}, body BudgetUpdateParams, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("budgets/%v", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of all active and historical budgets for the authenticated
// user.
func (r *BudgetService) List(ctx context.Context, query BudgetListParams, opts ...option.RequestOption) (res *BudgetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specific budget from the user's profile.
func (r *BudgetService) Delete(ctx context.Context, budgetID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("budgets/%v", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Budget struct {
	// Unique identifier for the budget.
	ID interface{} `json:"id,required"`
	// Percentage threshold at which an alert is triggered (e.g., 80% spent).
	AlertThreshold interface{} `json:"alertThreshold,required"`
	// Breakdown of the budget by categories.
	Categories []BudgetCategory `json:"categories,required"`
	// End date of the budget period.
	EndDate interface{} `json:"endDate,required"`
	// Name of the budget.
	Name interface{} `json:"name,required"`
	// The frequency or period of the budget.
	Period BudgetPeriod `json:"period,required"`
	// Remaining amount in the budget.
	RemainingAmount interface{} `json:"remainingAmount,required"`
	// Total amount spent against this budget so far.
	SpentAmount interface{} `json:"spentAmount,required"`
	// Start date of the budget period.
	StartDate interface{} `json:"startDate,required"`
	// Current status of the budget.
	Status BudgetStatus `json:"status,required"`
	// Total amount allocated for the entire budget.
	TotalAmount interface{} `json:"totalAmount,required"`
	// AI-driven recommendations related to this budget.
	AIRecommendations []AIInsight `json:"aiRecommendations,nullable"`
	JSON              budgetJSON  `json:"-"`
}

// budgetJSON contains the JSON metadata for the struct [Budget]
type budgetJSON struct {
	ID                apijson.Field
	AlertThreshold    apijson.Field
	Categories        apijson.Field
	EndDate           apijson.Field
	Name              apijson.Field
	Period            apijson.Field
	RemainingAmount   apijson.Field
	SpentAmount       apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	TotalAmount       apijson.Field
	AIRecommendations apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Budget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetJSON) RawJSON() string {
	return r.raw
}

type BudgetCategory struct {
	// Amount allocated to this category.
	Allocated interface{} `json:"allocated,required"`
	// Name of the budget category.
	Name interface{} `json:"name,required"`
	// Remaining amount in this category.
	Remaining interface{} `json:"remaining,required"`
	// Amount spent in this category so far.
	Spent interface{}        `json:"spent,required"`
	JSON  budgetCategoryJSON `json:"-"`
}

// budgetCategoryJSON contains the JSON metadata for the struct [BudgetCategory]
type budgetCategoryJSON struct {
	Allocated   apijson.Field
	Name        apijson.Field
	Remaining   apijson.Field
	Spent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetCategoryJSON) RawJSON() string {
	return r.raw
}

// The frequency or period of the budget.
type BudgetPeriod string

const (
	BudgetPeriodWeekly    BudgetPeriod = "weekly"
	BudgetPeriodBiWeekly  BudgetPeriod = "bi_weekly"
	BudgetPeriodMonthly   BudgetPeriod = "monthly"
	BudgetPeriodQuarterly BudgetPeriod = "quarterly"
	BudgetPeriodAnnually  BudgetPeriod = "annually"
	BudgetPeriodCustom    BudgetPeriod = "custom"
)

func (r BudgetPeriod) IsKnown() bool {
	switch r {
	case BudgetPeriodWeekly, BudgetPeriodBiWeekly, BudgetPeriodMonthly, BudgetPeriodQuarterly, BudgetPeriodAnnually, BudgetPeriodCustom:
		return true
	}
	return false
}

// Current status of the budget.
type BudgetStatus string

const (
	BudgetStatusActive   BudgetStatus = "active"
	BudgetStatusArchived BudgetStatus = "archived"
	BudgetStatusEnded    BudgetStatus = "ended"
)

func (r BudgetStatus) IsKnown() bool {
	switch r {
	case BudgetStatusActive, BudgetStatusArchived, BudgetStatusEnded:
		return true
	}
	return false
}

type BudgetListResponse struct {
	Data []Budget               `json:"data"`
	JSON budgetListResponseJSON `json:"-"`
	PaginatedList
}

// budgetListResponseJSON contains the JSON metadata for the struct
// [BudgetListResponse]
type budgetListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetListResponseJSON) RawJSON() string {
	return r.raw
}

type BudgetNewParams struct {
	// End date of the budget period.
	EndDate param.Field[interface{}] `json:"endDate,required"`
	// Name of the new budget.
	Name param.Field[interface{}] `json:"name,required"`
	// The frequency or period of the budget.
	Period param.Field[BudgetNewParamsPeriod] `json:"period,required"`
	// Start date of the budget period.
	StartDate param.Field[interface{}] `json:"startDate,required"`
	// Total amount allocated for the entire budget.
	TotalAmount param.Field[interface{}] `json:"totalAmount,required"`
	// If true, AI will automatically populate categories and amounts based on
	// historical spending.
	AIAutoPopulate param.Field[interface{}] `json:"aiAutoPopulate"`
	// Percentage threshold at which an alert is triggered.
	AlertThreshold param.Field[interface{}] `json:"alertThreshold"`
	// Initial breakdown of the budget by categories.
	Categories param.Field[[]BudgetNewParamsCategory] `json:"categories"`
}

func (r BudgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency or period of the budget.
type BudgetNewParamsPeriod string

const (
	BudgetNewParamsPeriodWeekly    BudgetNewParamsPeriod = "weekly"
	BudgetNewParamsPeriodBiWeekly  BudgetNewParamsPeriod = "bi_weekly"
	BudgetNewParamsPeriodMonthly   BudgetNewParamsPeriod = "monthly"
	BudgetNewParamsPeriodQuarterly BudgetNewParamsPeriod = "quarterly"
	BudgetNewParamsPeriodAnnually  BudgetNewParamsPeriod = "annually"
	BudgetNewParamsPeriodCustom    BudgetNewParamsPeriod = "custom"
)

func (r BudgetNewParamsPeriod) IsKnown() bool {
	switch r {
	case BudgetNewParamsPeriodWeekly, BudgetNewParamsPeriodBiWeekly, BudgetNewParamsPeriodMonthly, BudgetNewParamsPeriodQuarterly, BudgetNewParamsPeriodAnnually, BudgetNewParamsPeriodCustom:
		return true
	}
	return false
}

type BudgetNewParamsCategory struct {
	Allocated param.Field[interface{}] `json:"allocated"`
	Name      param.Field[interface{}] `json:"name"`
}

func (r BudgetNewParamsCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetUpdateParams struct {
	// Updated percentage threshold for alerts.
	AlertThreshold param.Field[interface{}] `json:"alertThreshold"`
	// Updated breakdown of the budget by categories. Existing categories will be
	// updated, new ones added.
	Categories param.Field[[]BudgetUpdateParamsCategory] `json:"categories"`
	// Updated end date of the budget period.
	EndDate param.Field[interface{}] `json:"endDate"`
	// Updated name of the budget.
	Name param.Field[interface{}] `json:"name"`
	// Updated start date of the budget period.
	StartDate param.Field[interface{}] `json:"startDate"`
	// Updated status of the budget.
	Status param.Field[BudgetUpdateParamsStatus] `json:"status"`
	// Updated total amount for the entire budget.
	TotalAmount param.Field[interface{}] `json:"totalAmount"`
}

func (r BudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetUpdateParamsCategory struct {
	Allocated param.Field[interface{}] `json:"allocated"`
	Name      param.Field[interface{}] `json:"name"`
}

func (r BudgetUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated status of the budget.
type BudgetUpdateParamsStatus string

const (
	BudgetUpdateParamsStatusActive   BudgetUpdateParamsStatus = "active"
	BudgetUpdateParamsStatusArchived BudgetUpdateParamsStatus = "archived"
	BudgetUpdateParamsStatusEnded    BudgetUpdateParamsStatus = "ended"
)

func (r BudgetUpdateParamsStatus) IsKnown() bool {
	switch r {
	case BudgetUpdateParamsStatusActive, BudgetUpdateParamsStatusArchived, BudgetUpdateParamsStatusEnded:
		return true
	}
	return false
}

type BudgetListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [BudgetListParams]'s query parameters as `url.Values`.
func (r BudgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
