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

// InvestmentPortfolioService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentPortfolioService] method instead.
type InvestmentPortfolioService struct {
	Options []option.RequestOption
}

// NewInvestmentPortfolioService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestmentPortfolioService(opts ...option.RequestOption) (r *InvestmentPortfolioService) {
	r = &InvestmentPortfolioService{}
	r.Options = opts
	return
}

// Creates a new investment portfolio, with options for initial asset allocation.
func (r *InvestmentPortfolioService) New(ctx context.Context, body InvestmentPortfolioNewParams, opts ...option.RequestOption) (res *InvestmentPortfolio, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information for a specific investment portfolio, including
// holdings, performance, and AI insights.
func (r *InvestmentPortfolioService) Get(ctx context.Context, portfolioID interface{}, opts ...option.RequestOption) (res *InvestmentPortfolio, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("investments/portfolios/%v", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates high-level details of an investment portfolio, such as name or risk
// tolerance.
func (r *InvestmentPortfolioService) Update(ctx context.Context, portfolioID interface{}, body InvestmentPortfolioUpdateParams, opts ...option.RequestOption) (res *InvestmentPortfolio, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("investments/portfolios/%v", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a summary of all investment portfolios linked to the user's account.
func (r *InvestmentPortfolioService) List(ctx context.Context, query InvestmentPortfolioListParams, opts ...option.RequestOption) (res *InvestmentPortfolioListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Triggers an AI-driven rebalancing process for a specific investment portfolio
// based on a target risk tolerance or strategy.
func (r *InvestmentPortfolioService) Rebalance(ctx context.Context, portfolioID interface{}, body InvestmentPortfolioRebalanceParams, opts ...option.RequestOption) (res *InvestmentPortfolioRebalanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("investments/portfolios/%v/rebalance", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InvestmentPortfolio struct {
	// Unique identifier for the investment portfolio.
	ID interface{} `json:"id,required"`
	// ISO 4217 currency code of the portfolio.
	Currency interface{} `json:"currency,required"`
	// Timestamp when the portfolio data was last updated.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Name of the portfolio.
	Name interface{} `json:"name,required"`
	// User's stated or AI-assessed risk tolerance for this portfolio.
	RiskTolerance InvestmentPortfolioRiskTolerance `json:"riskTolerance,required"`
	// Daily gain or loss on the portfolio.
	TodayGainLoss interface{} `json:"todayGainLoss,required"`
	// Current total market value of the portfolio.
	TotalValue interface{} `json:"totalValue,required"`
	// General type or strategy of the portfolio.
	Type InvestmentPortfolioType `json:"type,required"`
	// Total unrealized gain or loss on the portfolio.
	UnrealizedGainLoss interface{} `json:"unrealizedGainLoss,required"`
	// AI-driven insights into portfolio performance and market outlook.
	AIPerformanceInsights []AIInsight `json:"aiPerformanceInsights,nullable"`
	// Frequency at which AI-driven rebalancing is set to occur.
	AIRebalancingFrequency InvestmentPortfolioAIRebalancingFrequency `json:"aiRebalancingFrequency,nullable"`
	// List of individual assets held in the portfolio.
	Holdings []InvestmentPortfolioHolding `json:"holdings"`
	JSON     investmentPortfolioJSON      `json:"-"`
}

// investmentPortfolioJSON contains the JSON metadata for the struct
// [InvestmentPortfolio]
type investmentPortfolioJSON struct {
	ID                     apijson.Field
	Currency               apijson.Field
	LastUpdated            apijson.Field
	Name                   apijson.Field
	RiskTolerance          apijson.Field
	TodayGainLoss          apijson.Field
	TotalValue             apijson.Field
	Type                   apijson.Field
	UnrealizedGainLoss     apijson.Field
	AIPerformanceInsights  apijson.Field
	AIRebalancingFrequency apijson.Field
	Holdings               apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestmentPortfolio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioJSON) RawJSON() string {
	return r.raw
}

// User's stated or AI-assessed risk tolerance for this portfolio.
type InvestmentPortfolioRiskTolerance string

const (
	InvestmentPortfolioRiskToleranceConservative   InvestmentPortfolioRiskTolerance = "conservative"
	InvestmentPortfolioRiskToleranceModerate       InvestmentPortfolioRiskTolerance = "moderate"
	InvestmentPortfolioRiskToleranceAggressive     InvestmentPortfolioRiskTolerance = "aggressive"
	InvestmentPortfolioRiskToleranceVeryAggressive InvestmentPortfolioRiskTolerance = "very_aggressive"
)

func (r InvestmentPortfolioRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRiskToleranceConservative, InvestmentPortfolioRiskToleranceModerate, InvestmentPortfolioRiskToleranceAggressive, InvestmentPortfolioRiskToleranceVeryAggressive:
		return true
	}
	return false
}

// General type or strategy of the portfolio.
type InvestmentPortfolioType string

const (
	InvestmentPortfolioTypeEquities    InvestmentPortfolioType = "equities"
	InvestmentPortfolioTypeBonds       InvestmentPortfolioType = "bonds"
	InvestmentPortfolioTypeDiversified InvestmentPortfolioType = "diversified"
	InvestmentPortfolioTypeCrypto      InvestmentPortfolioType = "crypto"
	InvestmentPortfolioTypeRetirement  InvestmentPortfolioType = "retirement"
	InvestmentPortfolioTypeOther       InvestmentPortfolioType = "other"
)

func (r InvestmentPortfolioType) IsKnown() bool {
	switch r {
	case InvestmentPortfolioTypeEquities, InvestmentPortfolioTypeBonds, InvestmentPortfolioTypeDiversified, InvestmentPortfolioTypeCrypto, InvestmentPortfolioTypeRetirement, InvestmentPortfolioTypeOther:
		return true
	}
	return false
}

// Frequency at which AI-driven rebalancing is set to occur.
type InvestmentPortfolioAIRebalancingFrequency string

const (
	InvestmentPortfolioAIRebalancingFrequencyMonthly      InvestmentPortfolioAIRebalancingFrequency = "monthly"
	InvestmentPortfolioAIRebalancingFrequencyQuarterly    InvestmentPortfolioAIRebalancingFrequency = "quarterly"
	InvestmentPortfolioAIRebalancingFrequencySemiAnnually InvestmentPortfolioAIRebalancingFrequency = "semi_annually"
	InvestmentPortfolioAIRebalancingFrequencyAnnually     InvestmentPortfolioAIRebalancingFrequency = "annually"
	InvestmentPortfolioAIRebalancingFrequencyNever        InvestmentPortfolioAIRebalancingFrequency = "never"
)

func (r InvestmentPortfolioAIRebalancingFrequency) IsKnown() bool {
	switch r {
	case InvestmentPortfolioAIRebalancingFrequencyMonthly, InvestmentPortfolioAIRebalancingFrequencyQuarterly, InvestmentPortfolioAIRebalancingFrequencySemiAnnually, InvestmentPortfolioAIRebalancingFrequencyAnnually, InvestmentPortfolioAIRebalancingFrequencyNever:
		return true
	}
	return false
}

type InvestmentPortfolioHolding struct {
	// Average cost per unit.
	AverageCost interface{} `json:"averageCost,required"`
	// Current market price per unit.
	CurrentPrice interface{} `json:"currentPrice,required"`
	// Total market value of the holding.
	MarketValue interface{} `json:"marketValue,required"`
	// Full name of the investment asset.
	Name interface{} `json:"name,required"`
	// Percentage of the total portfolio value this holding represents.
	PercentageOfPortfolio interface{} `json:"percentageOfPortfolio,required"`
	// Number of units held.
	Quantity interface{} `json:"quantity,required"`
	// Stock ticker or asset symbol.
	Symbol interface{} `json:"symbol,required"`
	// Overall ESG (Environmental, Social, Governance) score of the asset (0-10).
	EsgScore interface{}                    `json:"esgScore"`
	JSON     investmentPortfolioHoldingJSON `json:"-"`
}

// investmentPortfolioHoldingJSON contains the JSON metadata for the struct
// [InvestmentPortfolioHolding]
type investmentPortfolioHoldingJSON struct {
	AverageCost           apijson.Field
	CurrentPrice          apijson.Field
	MarketValue           apijson.Field
	Name                  apijson.Field
	PercentageOfPortfolio apijson.Field
	Quantity              apijson.Field
	Symbol                apijson.Field
	EsgScore              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvestmentPortfolioHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioHoldingJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioListResponse struct {
	Data []InvestmentPortfolio               `json:"data"`
	JSON investmentPortfolioListResponseJSON `json:"-"`
	PaginatedList
}

// investmentPortfolioListResponseJSON contains the JSON metadata for the struct
// [InvestmentPortfolioListResponse]
type investmentPortfolioListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestmentPortfolioListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioRebalanceResponse struct {
	// ID of the portfolio being rebalanced.
	PortfolioID interface{} `json:"portfolioId,required"`
	// Unique identifier for the rebalancing operation.
	RebalanceID interface{} `json:"rebalanceId,required"`
	// Current status of the rebalancing operation.
	Status InvestmentPortfolioRebalanceResponseStatus `json:"status,required"`
	// A descriptive message about the current rebalance status.
	StatusMessage interface{} `json:"statusMessage,required"`
	// Timestamp when the rebalance confirmation expires, if `confirmationRequired` is
	// true.
	ConfirmationExpiresAt interface{} `json:"confirmationExpiresAt"`
	// Indicates if user confirmation is required to proceed with trades.
	ConfirmationRequired interface{} `json:"confirmationRequired"`
	// AI-estimated impact of the rebalance on the portfolio.
	EstimatedImpact interface{} `json:"estimatedImpact"`
	// List of proposed trades if `dryRun` was true and status is
	// `pending_confirmation`.
	ProposedTrades []InvestmentPortfolioRebalanceResponseProposedTrade `json:"proposedTrades,nullable"`
	JSON           investmentPortfolioRebalanceResponseJSON            `json:"-"`
}

// investmentPortfolioRebalanceResponseJSON contains the JSON metadata for the
// struct [InvestmentPortfolioRebalanceResponse]
type investmentPortfolioRebalanceResponseJSON struct {
	PortfolioID           apijson.Field
	RebalanceID           apijson.Field
	Status                apijson.Field
	StatusMessage         apijson.Field
	ConfirmationExpiresAt apijson.Field
	ConfirmationRequired  apijson.Field
	EstimatedImpact       apijson.Field
	ProposedTrades        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvestmentPortfolioRebalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioRebalanceResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the rebalancing operation.
type InvestmentPortfolioRebalanceResponseStatus string

const (
	InvestmentPortfolioRebalanceResponseStatusAnalyzing           InvestmentPortfolioRebalanceResponseStatus = "analyzing"
	InvestmentPortfolioRebalanceResponseStatusPendingConfirmation InvestmentPortfolioRebalanceResponseStatus = "pending_confirmation"
	InvestmentPortfolioRebalanceResponseStatusExecutingTrades     InvestmentPortfolioRebalanceResponseStatus = "executing_trades"
	InvestmentPortfolioRebalanceResponseStatusCompleted           InvestmentPortfolioRebalanceResponseStatus = "completed"
	InvestmentPortfolioRebalanceResponseStatusFailed              InvestmentPortfolioRebalanceResponseStatus = "failed"
)

func (r InvestmentPortfolioRebalanceResponseStatus) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRebalanceResponseStatusAnalyzing, InvestmentPortfolioRebalanceResponseStatusPendingConfirmation, InvestmentPortfolioRebalanceResponseStatusExecutingTrades, InvestmentPortfolioRebalanceResponseStatusCompleted, InvestmentPortfolioRebalanceResponseStatusFailed:
		return true
	}
	return false
}

type InvestmentPortfolioRebalanceResponseProposedTrade struct {
	Action         InvestmentPortfolioRebalanceResponseProposedTradesAction `json:"action"`
	EstimatedPrice interface{}                                              `json:"estimatedPrice"`
	Quantity       interface{}                                              `json:"quantity"`
	Symbol         interface{}                                              `json:"symbol"`
	JSON           investmentPortfolioRebalanceResponseProposedTradeJSON    `json:"-"`
}

// investmentPortfolioRebalanceResponseProposedTradeJSON contains the JSON metadata
// for the struct [InvestmentPortfolioRebalanceResponseProposedTrade]
type investmentPortfolioRebalanceResponseProposedTradeJSON struct {
	Action         apijson.Field
	EstimatedPrice apijson.Field
	Quantity       apijson.Field
	Symbol         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvestmentPortfolioRebalanceResponseProposedTrade) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioRebalanceResponseProposedTradeJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioRebalanceResponseProposedTradesAction string

const (
	InvestmentPortfolioRebalanceResponseProposedTradesActionBuy  InvestmentPortfolioRebalanceResponseProposedTradesAction = "buy"
	InvestmentPortfolioRebalanceResponseProposedTradesActionSell InvestmentPortfolioRebalanceResponseProposedTradesAction = "sell"
)

func (r InvestmentPortfolioRebalanceResponseProposedTradesAction) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRebalanceResponseProposedTradesActionBuy, InvestmentPortfolioRebalanceResponseProposedTradesActionSell:
		return true
	}
	return false
}

type InvestmentPortfolioNewParams struct {
	// ISO 4217 currency code of the portfolio.
	Currency param.Field[interface{}] `json:"currency,required"`
	// Initial amount to invest into the portfolio.
	InitialInvestment param.Field[interface{}] `json:"initialInvestment,required"`
	// Name for the new investment portfolio.
	Name param.Field[interface{}] `json:"name,required"`
	// Desired risk tolerance for this portfolio.
	RiskTolerance param.Field[InvestmentPortfolioNewParamsRiskTolerance] `json:"riskTolerance,required"`
	// General type or strategy of the portfolio.
	Type param.Field[InvestmentPortfolioNewParamsType] `json:"type,required"`
	// If true, AI will automatically allocate initial investment based on risk
	// tolerance.
	AIAutoAllocate param.Field[interface{}] `json:"aiAutoAllocate"`
	// Optional: ID of a linked account to fund the initial investment.
	LinkedAccountID param.Field[interface{}] `json:"linkedAccountId"`
}

func (r InvestmentPortfolioNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired risk tolerance for this portfolio.
type InvestmentPortfolioNewParamsRiskTolerance string

const (
	InvestmentPortfolioNewParamsRiskToleranceConservative   InvestmentPortfolioNewParamsRiskTolerance = "conservative"
	InvestmentPortfolioNewParamsRiskToleranceModerate       InvestmentPortfolioNewParamsRiskTolerance = "moderate"
	InvestmentPortfolioNewParamsRiskToleranceAggressive     InvestmentPortfolioNewParamsRiskTolerance = "aggressive"
	InvestmentPortfolioNewParamsRiskToleranceVeryAggressive InvestmentPortfolioNewParamsRiskTolerance = "very_aggressive"
)

func (r InvestmentPortfolioNewParamsRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioNewParamsRiskToleranceConservative, InvestmentPortfolioNewParamsRiskToleranceModerate, InvestmentPortfolioNewParamsRiskToleranceAggressive, InvestmentPortfolioNewParamsRiskToleranceVeryAggressive:
		return true
	}
	return false
}

// General type or strategy of the portfolio.
type InvestmentPortfolioNewParamsType string

const (
	InvestmentPortfolioNewParamsTypeEquities    InvestmentPortfolioNewParamsType = "equities"
	InvestmentPortfolioNewParamsTypeBonds       InvestmentPortfolioNewParamsType = "bonds"
	InvestmentPortfolioNewParamsTypeDiversified InvestmentPortfolioNewParamsType = "diversified"
	InvestmentPortfolioNewParamsTypeCrypto      InvestmentPortfolioNewParamsType = "crypto"
	InvestmentPortfolioNewParamsTypeRetirement  InvestmentPortfolioNewParamsType = "retirement"
	InvestmentPortfolioNewParamsTypeOther       InvestmentPortfolioNewParamsType = "other"
)

func (r InvestmentPortfolioNewParamsType) IsKnown() bool {
	switch r {
	case InvestmentPortfolioNewParamsTypeEquities, InvestmentPortfolioNewParamsTypeBonds, InvestmentPortfolioNewParamsTypeDiversified, InvestmentPortfolioNewParamsTypeCrypto, InvestmentPortfolioNewParamsTypeRetirement, InvestmentPortfolioNewParamsTypeOther:
		return true
	}
	return false
}

type InvestmentPortfolioUpdateParams struct {
	// Updated frequency for AI-driven rebalancing.
	AIRebalancingFrequency param.Field[InvestmentPortfolioUpdateParamsAIRebalancingFrequency] `json:"aiRebalancingFrequency"`
	// Updated name of the portfolio.
	Name param.Field[interface{}] `json:"name"`
	// Updated risk tolerance for this portfolio. May trigger rebalancing.
	RiskTolerance param.Field[InvestmentPortfolioUpdateParamsRiskTolerance] `json:"riskTolerance"`
}

func (r InvestmentPortfolioUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated frequency for AI-driven rebalancing.
type InvestmentPortfolioUpdateParamsAIRebalancingFrequency string

const (
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyMonthly      InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "monthly"
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyQuarterly    InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "quarterly"
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencySemiAnnually InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "semi_annually"
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyAnnually     InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "annually"
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyNever        InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "never"
)

func (r InvestmentPortfolioUpdateParamsAIRebalancingFrequency) IsKnown() bool {
	switch r {
	case InvestmentPortfolioUpdateParamsAIRebalancingFrequencyMonthly, InvestmentPortfolioUpdateParamsAIRebalancingFrequencyQuarterly, InvestmentPortfolioUpdateParamsAIRebalancingFrequencySemiAnnually, InvestmentPortfolioUpdateParamsAIRebalancingFrequencyAnnually, InvestmentPortfolioUpdateParamsAIRebalancingFrequencyNever:
		return true
	}
	return false
}

// Updated risk tolerance for this portfolio. May trigger rebalancing.
type InvestmentPortfolioUpdateParamsRiskTolerance string

const (
	InvestmentPortfolioUpdateParamsRiskToleranceConservative   InvestmentPortfolioUpdateParamsRiskTolerance = "conservative"
	InvestmentPortfolioUpdateParamsRiskToleranceModerate       InvestmentPortfolioUpdateParamsRiskTolerance = "moderate"
	InvestmentPortfolioUpdateParamsRiskToleranceAggressive     InvestmentPortfolioUpdateParamsRiskTolerance = "aggressive"
	InvestmentPortfolioUpdateParamsRiskToleranceVeryAggressive InvestmentPortfolioUpdateParamsRiskTolerance = "very_aggressive"
)

func (r InvestmentPortfolioUpdateParamsRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioUpdateParamsRiskToleranceConservative, InvestmentPortfolioUpdateParamsRiskToleranceModerate, InvestmentPortfolioUpdateParamsRiskToleranceAggressive, InvestmentPortfolioUpdateParamsRiskToleranceVeryAggressive:
		return true
	}
	return false
}

type InvestmentPortfolioListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [InvestmentPortfolioListParams]'s query parameters as
// `url.Values`.
func (r InvestmentPortfolioListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentPortfolioRebalanceParams struct {
	// The desired risk tolerance for rebalancing the portfolio.
	TargetRiskTolerance param.Field[InvestmentPortfolioRebalanceParamsTargetRiskTolerance] `json:"targetRiskTolerance,required"`
	// If true, user confirmation is required before executing actual trades after a
	// dry run.
	ConfirmationRequired param.Field[interface{}] `json:"confirmationRequired"`
	// If true, only simulate the rebalance without executing trades. Returns proposed
	// trades.
	DryRun param.Field[interface{}] `json:"dryRun"`
}

func (r InvestmentPortfolioRebalanceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The desired risk tolerance for rebalancing the portfolio.
type InvestmentPortfolioRebalanceParamsTargetRiskTolerance string

const (
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceConservative   InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "conservative"
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceModerate       InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "moderate"
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceAggressive     InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "aggressive"
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceVeryAggressive InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "very_aggressive"
)

func (r InvestmentPortfolioRebalanceParamsTargetRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRebalanceParamsTargetRiskToleranceConservative, InvestmentPortfolioRebalanceParamsTargetRiskToleranceModerate, InvestmentPortfolioRebalanceParamsTargetRiskToleranceAggressive, InvestmentPortfolioRebalanceParamsTargetRiskToleranceVeryAggressive:
		return true
	}
	return false
}
