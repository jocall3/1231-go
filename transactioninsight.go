// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// TransactionInsightService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionInsightService] method instead.
type TransactionInsightService struct {
	Options []option.RequestOption
}

// NewTransactionInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionInsightService(opts ...option.RequestOption) (r *TransactionInsightService) {
	r = &TransactionInsightService{}
	r.Options = opts
	return
}

// Retrieves AI-generated insights into user spending trends over time, identifying
// patterns and anomalies.
func (r *TransactionInsightService) GetSpendingTrends(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetSpendingTrendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/spending-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// An AI-generated insight, alert, or recommendation.
type AIInsight struct {
	// Unique identifier for the AI insight.
	ID interface{} `json:"id,required"`
	// Category of the insight (e.g., spending, saving, investing).
	Category AIInsightCategory `json:"category,required"`
	// Detailed explanation of the insight.
	Description interface{} `json:"description,required"`
	// AI-assessed severity or importance of the insight.
	Severity AIInsightSeverity `json:"severity,required"`
	// Timestamp when the insight was generated.
	Timestamp interface{} `json:"timestamp,required"`
	// A concise title for the insight.
	Title interface{} `json:"title,required"`
	// Optional: A concrete action the user can take based on the insight.
	ActionableRecommendation interface{} `json:"actionableRecommendation"`
	// Optional: A programmatic trigger or deep link to initiate the recommended
	// action.
	ActionTrigger interface{}   `json:"actionTrigger"`
	JSON          aiInsightJSON `json:"-"`
}

// aiInsightJSON contains the JSON metadata for the struct [AIInsight]
type aiInsightJSON struct {
	ID                       apijson.Field
	Category                 apijson.Field
	Description              apijson.Field
	Severity                 apijson.Field
	Timestamp                apijson.Field
	Title                    apijson.Field
	ActionableRecommendation apijson.Field
	ActionTrigger            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AIInsight) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInsightJSON) RawJSON() string {
	return r.raw
}

// Category of the insight (e.g., spending, saving, investing).
type AIInsightCategory string

const (
	AIInsightCategorySpending          AIInsightCategory = "spending"
	AIInsightCategorySaving            AIInsightCategory = "saving"
	AIInsightCategoryInvesting         AIInsightCategory = "investing"
	AIInsightCategoryBudgeting         AIInsightCategory = "budgeting"
	AIInsightCategorySecurity          AIInsightCategory = "security"
	AIInsightCategoryFinancialGoals    AIInsightCategory = "financial_goals"
	AIInsightCategorySustainability    AIInsightCategory = "sustainability"
	AIInsightCategoryCorporateTreasury AIInsightCategory = "corporate_treasury"
	AIInsightCategoryCompliance        AIInsightCategory = "compliance"
	AIInsightCategoryOther             AIInsightCategory = "other"
)

func (r AIInsightCategory) IsKnown() bool {
	switch r {
	case AIInsightCategorySpending, AIInsightCategorySaving, AIInsightCategoryInvesting, AIInsightCategoryBudgeting, AIInsightCategorySecurity, AIInsightCategoryFinancialGoals, AIInsightCategorySustainability, AIInsightCategoryCorporateTreasury, AIInsightCategoryCompliance, AIInsightCategoryOther:
		return true
	}
	return false
}

// AI-assessed severity or importance of the insight.
type AIInsightSeverity string

const (
	AIInsightSeverityLow      AIInsightSeverity = "low"
	AIInsightSeverityMedium   AIInsightSeverity = "medium"
	AIInsightSeverityHigh     AIInsightSeverity = "high"
	AIInsightSeverityCritical AIInsightSeverity = "critical"
)

func (r AIInsightSeverity) IsKnown() bool {
	switch r {
	case AIInsightSeverityLow, AIInsightSeverityMedium, AIInsightSeverityHigh, AIInsightSeverityCritical:
		return true
	}
	return false
}

type TransactionInsightGetSpendingTrendsResponse struct {
	// AI-driven insights and recommendations related to spending.
	AIInsights []AIInsight `json:"aiInsights,required"`
	// AI-projected total spending for the next month.
	ForecastNextMonth interface{} `json:"forecastNextMonth,required"`
	// Overall trend of spending (increasing, decreasing, stable).
	OverallTrend TransactionInsightGetSpendingTrendsResponseOverallTrend `json:"overallTrend,required"`
	// Percentage change in spending over the period.
	PercentageChange interface{} `json:"percentageChange,required"`
	// The period over which the spending trend is analyzed.
	Period interface{} `json:"period,required"`
	// Categories with the most significant changes in spending.
	TopCategoriesByChange []TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange `json:"topCategoriesByChange,required"`
	JSON                  transactionInsightGetSpendingTrendsResponseJSON                    `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseJSON contains the JSON metadata for
// the struct [TransactionInsightGetSpendingTrendsResponse]
type transactionInsightGetSpendingTrendsResponseJSON struct {
	AIInsights            apijson.Field
	ForecastNextMonth     apijson.Field
	OverallTrend          apijson.Field
	PercentageChange      apijson.Field
	Period                apijson.Field
	TopCategoriesByChange apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseJSON) RawJSON() string {
	return r.raw
}

// Overall trend of spending (increasing, decreasing, stable).
type TransactionInsightGetSpendingTrendsResponseOverallTrend string

const (
	TransactionInsightGetSpendingTrendsResponseOverallTrendIncreasing TransactionInsightGetSpendingTrendsResponseOverallTrend = "increasing"
	TransactionInsightGetSpendingTrendsResponseOverallTrendDecreasing TransactionInsightGetSpendingTrendsResponseOverallTrend = "decreasing"
	TransactionInsightGetSpendingTrendsResponseOverallTrendStable     TransactionInsightGetSpendingTrendsResponseOverallTrend = "stable"
)

func (r TransactionInsightGetSpendingTrendsResponseOverallTrend) IsKnown() bool {
	switch r {
	case TransactionInsightGetSpendingTrendsResponseOverallTrendIncreasing, TransactionInsightGetSpendingTrendsResponseOverallTrendDecreasing, TransactionInsightGetSpendingTrendsResponseOverallTrendStable:
		return true
	}
	return false
}

type TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange struct {
	AbsoluteChange   interface{}                                                          `json:"absoluteChange"`
	Category         interface{}                                                          `json:"category"`
	PercentageChange interface{}                                                          `json:"percentageChange"`
	JSON             transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON contains
// the JSON metadata for the struct
// [TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange]
type transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON struct {
	AbsoluteChange   apijson.Field
	Category         apijson.Field
	PercentageChange apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON) RawJSON() string {
	return r.raw
}
