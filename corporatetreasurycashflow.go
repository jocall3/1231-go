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

// CorporateTreasuryCashFlowService contains methods and other services that help
// with interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateTreasuryCashFlowService] method instead.
type CorporateTreasuryCashFlowService struct {
	Options []option.RequestOption
}

// NewCorporateTreasuryCashFlowService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCorporateTreasuryCashFlowService(opts ...option.RequestOption) (r *CorporateTreasuryCashFlowService) {
	r = &CorporateTreasuryCashFlowService{}
	r.Options = opts
	return
}

// Retrieves an advanced AI-driven cash flow forecast for the organization,
// projecting liquidity, identifying potential surpluses or deficits, and providing
// recommendations for optimal treasury management.
func (r *CorporateTreasuryCashFlowService) GetForecast(ctx context.Context, query CorporateTreasuryCashFlowGetForecastParams, opts ...option.RequestOption) (res *CorporateTreasuryCashFlowGetForecastResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/cash-flow/forecast"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CorporateTreasuryCashFlowGetForecastResponse struct {
	// AI-generated recommendations for treasury optimization.
	AIRecommendations []AIInsight `json:"aiRecommendations,required"`
	// The currency of the forecast.
	Currency interface{} `json:"currency,required"`
	// Unique identifier for the cash flow forecast report.
	ForecastID interface{} `json:"forecastId,required"`
	// Forecast of cash inflows by source.
	InflowForecast CorporateTreasuryCashFlowGetForecastResponseInflowForecast `json:"inflowForecast,required"`
	// AI-assessed risk score for liquidity (0-100, lower is better).
	LiquidityRiskScore interface{} `json:"liquidityRiskScore,required"`
	// Forecast of cash outflows by category.
	OutflowForecast CorporateTreasuryCashFlowGetForecastResponseOutflowForecast `json:"outflowForecast,required"`
	// Overall assessment of the projected cash flow.
	OverallStatus CorporateTreasuryCashFlowGetForecastResponseOverallStatus `json:"overallStatus,required"`
	// The period covered by the forecast.
	Period interface{} `json:"period,required"`
	// Projected cash balances at key dates, potentially across different scenarios.
	ProjectedBalances []CorporateTreasuryCashFlowGetForecastResponseProjectedBalance `json:"projectedBalances,required"`
	JSON              corporateTreasuryCashFlowGetForecastResponseJSON               `json:"-"`
}

// corporateTreasuryCashFlowGetForecastResponseJSON contains the JSON metadata for
// the struct [CorporateTreasuryCashFlowGetForecastResponse]
type corporateTreasuryCashFlowGetForecastResponseJSON struct {
	AIRecommendations  apijson.Field
	Currency           apijson.Field
	ForecastID         apijson.Field
	InflowForecast     apijson.Field
	LiquidityRiskScore apijson.Field
	OutflowForecast    apijson.Field
	OverallStatus      apijson.Field
	Period             apijson.Field
	ProjectedBalances  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CorporateTreasuryCashFlowGetForecastResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryCashFlowGetForecastResponseJSON) RawJSON() string {
	return r.raw
}

// Forecast of cash inflows by source.
type CorporateTreasuryCashFlowGetForecastResponseInflowForecast struct {
	BySource       []CorporateTreasuryCashFlowGetForecastResponseInflowForecastBySource `json:"bySource"`
	TotalProjected interface{}                                                          `json:"totalProjected"`
	JSON           corporateTreasuryCashFlowGetForecastResponseInflowForecastJSON       `json:"-"`
}

// corporateTreasuryCashFlowGetForecastResponseInflowForecastJSON contains the JSON
// metadata for the struct
// [CorporateTreasuryCashFlowGetForecastResponseInflowForecast]
type corporateTreasuryCashFlowGetForecastResponseInflowForecastJSON struct {
	BySource       apijson.Field
	TotalProjected apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateTreasuryCashFlowGetForecastResponseInflowForecast) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryCashFlowGetForecastResponseInflowForecastJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryCashFlowGetForecastResponseInflowForecastBySource struct {
	Amount interface{}                                                            `json:"amount"`
	Source interface{}                                                            `json:"source"`
	JSON   corporateTreasuryCashFlowGetForecastResponseInflowForecastBySourceJSON `json:"-"`
}

// corporateTreasuryCashFlowGetForecastResponseInflowForecastBySourceJSON contains
// the JSON metadata for the struct
// [CorporateTreasuryCashFlowGetForecastResponseInflowForecastBySource]
type corporateTreasuryCashFlowGetForecastResponseInflowForecastBySourceJSON struct {
	Amount      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateTreasuryCashFlowGetForecastResponseInflowForecastBySource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryCashFlowGetForecastResponseInflowForecastBySourceJSON) RawJSON() string {
	return r.raw
}

// Forecast of cash outflows by category.
type CorporateTreasuryCashFlowGetForecastResponseOutflowForecast struct {
	ByCategory     []CorporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategory `json:"byCategory"`
	TotalProjected interface{}                                                             `json:"totalProjected"`
	JSON           corporateTreasuryCashFlowGetForecastResponseOutflowForecastJSON         `json:"-"`
}

// corporateTreasuryCashFlowGetForecastResponseOutflowForecastJSON contains the
// JSON metadata for the struct
// [CorporateTreasuryCashFlowGetForecastResponseOutflowForecast]
type corporateTreasuryCashFlowGetForecastResponseOutflowForecastJSON struct {
	ByCategory     apijson.Field
	TotalProjected apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateTreasuryCashFlowGetForecastResponseOutflowForecast) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryCashFlowGetForecastResponseOutflowForecastJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategory struct {
	Amount   interface{}                                                               `json:"amount"`
	Category interface{}                                                               `json:"category"`
	JSON     corporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategoryJSON `json:"-"`
}

// corporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategoryJSON
// contains the JSON metadata for the struct
// [CorporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategory]
type corporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategoryJSON struct {
	Amount      apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryCashFlowGetForecastResponseOutflowForecastByCategoryJSON) RawJSON() string {
	return r.raw
}

// Overall assessment of the projected cash flow.
type CorporateTreasuryCashFlowGetForecastResponseOverallStatus string

const (
	CorporateTreasuryCashFlowGetForecastResponseOverallStatusPositiveOutlook CorporateTreasuryCashFlowGetForecastResponseOverallStatus = "positive_outlook"
	CorporateTreasuryCashFlowGetForecastResponseOverallStatusNegativeOutlook CorporateTreasuryCashFlowGetForecastResponseOverallStatus = "negative_outlook"
	CorporateTreasuryCashFlowGetForecastResponseOverallStatusStable          CorporateTreasuryCashFlowGetForecastResponseOverallStatus = "stable"
	CorporateTreasuryCashFlowGetForecastResponseOverallStatusUncertain       CorporateTreasuryCashFlowGetForecastResponseOverallStatus = "uncertain"
)

func (r CorporateTreasuryCashFlowGetForecastResponseOverallStatus) IsKnown() bool {
	switch r {
	case CorporateTreasuryCashFlowGetForecastResponseOverallStatusPositiveOutlook, CorporateTreasuryCashFlowGetForecastResponseOverallStatusNegativeOutlook, CorporateTreasuryCashFlowGetForecastResponseOverallStatusStable, CorporateTreasuryCashFlowGetForecastResponseOverallStatusUncertain:
		return true
	}
	return false
}

type CorporateTreasuryCashFlowGetForecastResponseProjectedBalance struct {
	Date          interface{}                                                           `json:"date"`
	ProjectedCash interface{}                                                           `json:"projectedCash"`
	Scenario      CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenario `json:"scenario"`
	JSON          corporateTreasuryCashFlowGetForecastResponseProjectedBalanceJSON      `json:"-"`
}

// corporateTreasuryCashFlowGetForecastResponseProjectedBalanceJSON contains the
// JSON metadata for the struct
// [CorporateTreasuryCashFlowGetForecastResponseProjectedBalance]
type corporateTreasuryCashFlowGetForecastResponseProjectedBalanceJSON struct {
	Date          apijson.Field
	ProjectedCash apijson.Field
	Scenario      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CorporateTreasuryCashFlowGetForecastResponseProjectedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryCashFlowGetForecastResponseProjectedBalanceJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenario string

const (
	CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenarioMostLikely CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenario = "most_likely"
	CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenarioBestCase   CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenario = "best_case"
	CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenarioWorstCase  CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenario = "worst_case"
)

func (r CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenario) IsKnown() bool {
	switch r {
	case CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenarioMostLikely, CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenarioBestCase, CorporateTreasuryCashFlowGetForecastResponseProjectedBalancesScenarioWorstCase:
		return true
	}
	return false
}

type CorporateTreasuryCashFlowGetForecastParams struct {
	// The number of days into the future for which to generate the cash flow forecast
	// (e.g., 30, 90, 180).
	ForecastHorizonDays param.Field[interface{}] `query:"forecastHorizonDays"`
	// If true, the forecast will include best-case and worst-case scenario analysis
	// alongside the most likely projection.
	IncludeScenarioAnalysis param.Field[interface{}] `query:"includeScenarioAnalysis"`
}

// URLQuery serializes [CorporateTreasuryCashFlowGetForecastParams]'s query
// parameters as `url.Values`.
func (r CorporateTreasuryCashFlowGetForecastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
