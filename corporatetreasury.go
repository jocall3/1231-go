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

// CorporateTreasuryService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateTreasuryService] method instead.
type CorporateTreasuryService struct {
	Options  []option.RequestOption
	CashFlow *CorporateTreasuryCashFlowService
}

// NewCorporateTreasuryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateTreasuryService(opts ...option.RequestOption) (r *CorporateTreasuryService) {
	r = &CorporateTreasuryService{}
	r.Options = opts
	r.CashFlow = NewCorporateTreasuryCashFlowService(opts...)
	return
}

// Provides a real-time overview of the organization's liquidity across all
// accounts, currencies, and short-term investments.
func (r *CorporateTreasuryService) GetLiquidityPositions(ctx context.Context, opts ...option.RequestOption) (res *CorporateTreasuryGetLiquidityPositionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/liquidity-positions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CorporateTreasuryGetLiquidityPositionsResponse struct {
	// Breakdown of liquid assets by account type.
	AccountTypeBreakdown []CorporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdown `json:"accountTypeBreakdown,required"`
	// AI's overall assessment of liquidity.
	AILiquidityAssessment CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessment `json:"aiLiquidityAssessment,required"`
	// AI-generated recommendations for liquidity management.
	AIRecommendations []AIInsight `json:"aiRecommendations,required"`
	// Breakdown of liquid assets by currency.
	CurrencyBreakdown []CorporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdown `json:"currencyBreakdown,required"`
	// Details on short-term investments contributing to liquidity.
	ShortTermInvestments CorporateTreasuryGetLiquidityPositionsResponseShortTermInvestments `json:"shortTermInvestments,required"`
	// Timestamp of the liquidity snapshot.
	SnapshotTime interface{} `json:"snapshotTime,required"`
	// Total value of all liquid assets across the organization.
	TotalLiquidAssets interface{}                                        `json:"totalLiquidAssets,required"`
	JSON              corporateTreasuryGetLiquidityPositionsResponseJSON `json:"-"`
}

// corporateTreasuryGetLiquidityPositionsResponseJSON contains the JSON metadata
// for the struct [CorporateTreasuryGetLiquidityPositionsResponse]
type corporateTreasuryGetLiquidityPositionsResponseJSON struct {
	AccountTypeBreakdown  apijson.Field
	AILiquidityAssessment apijson.Field
	AIRecommendations     apijson.Field
	CurrencyBreakdown     apijson.Field
	ShortTermInvestments  apijson.Field
	SnapshotTime          apijson.Field
	TotalLiquidAssets     apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CorporateTreasuryGetLiquidityPositionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetLiquidityPositionsResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdown struct {
	Amount interface{}                                                            `json:"amount"`
	Type   interface{}                                                            `json:"type"`
	JSON   corporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdownJSON `json:"-"`
}

// corporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdownJSON contains
// the JSON metadata for the struct
// [CorporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdown]
type corporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdownJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetLiquidityPositionsResponseAccountTypeBreakdownJSON) RawJSON() string {
	return r.raw
}

// AI's overall assessment of liquidity.
type CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessment struct {
	Message interface{}                                                               `json:"message"`
	Status  CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus `json:"status"`
	JSON    corporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentJSON   `json:"-"`
}

// corporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentJSON contains
// the JSON metadata for the struct
// [CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessment]
type corporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus string

const (
	CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusOptimal    CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus = "optimal"
	CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusSufficient CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus = "sufficient"
	CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusTight      CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus = "tight"
	CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusCritical   CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus = "critical"
)

func (r CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatus) IsKnown() bool {
	switch r {
	case CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusOptimal, CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusSufficient, CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusTight, CorporateTreasuryGetLiquidityPositionsResponseAILiquidityAssessmentStatusCritical:
		return true
	}
	return false
}

type CorporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdown struct {
	Amount     interface{}                                                         `json:"amount"`
	Currency   interface{}                                                         `json:"currency"`
	Percentage interface{}                                                         `json:"percentage"`
	JSON       corporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdownJSON `json:"-"`
}

// corporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdownJSON contains the
// JSON metadata for the struct
// [CorporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdown]
type corporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Percentage  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetLiquidityPositionsResponseCurrencyBreakdownJSON) RawJSON() string {
	return r.raw
}

// Details on short-term investments contributing to liquidity.
type CorporateTreasuryGetLiquidityPositionsResponseShortTermInvestments struct {
	MaturingNext30Days interface{}                                                            `json:"maturingNext30Days"`
	TotalValue         interface{}                                                            `json:"totalValue"`
	JSON               corporateTreasuryGetLiquidityPositionsResponseShortTermInvestmentsJSON `json:"-"`
}

// corporateTreasuryGetLiquidityPositionsResponseShortTermInvestmentsJSON contains
// the JSON metadata for the struct
// [CorporateTreasuryGetLiquidityPositionsResponseShortTermInvestments]
type corporateTreasuryGetLiquidityPositionsResponseShortTermInvestmentsJSON struct {
	MaturingNext30Days apijson.Field
	TotalValue         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CorporateTreasuryGetLiquidityPositionsResponseShortTermInvestments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetLiquidityPositionsResponseShortTermInvestmentsJSON) RawJSON() string {
	return r.raw
}
