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

// MarketplaceProductService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketplaceProductService] method instead.
type MarketplaceProductService struct {
	Options []option.RequestOption
}

// NewMarketplaceProductService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMarketplaceProductService(opts ...option.RequestOption) (r *MarketplaceProductService) {
	r = &MarketplaceProductService{}
	r.Options = opts
	return
}

// Retrieves a personalized, AI-curated list of products and services from the
// Plato AI marketplace, tailored to the user's financial profile, goals, and
// spending patterns. Includes options for filtering and advanced search.
func (r *MarketplaceProductService) List(ctx context.Context, query MarketplaceProductListParams, opts ...option.RequestOption) (res *MarketplaceProductListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketplace/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Uses the Quantum Oracle to simulate the long-term financial impact of purchasing
// or subscribing to a specific marketplace product, such as a loan, investment, or
// insurance policy, on the user's overall financial health and goals.
func (r *MarketplaceProductService) SimulateImpact(ctx context.Context, productID interface{}, body MarketplaceProductSimulateImpactParams, opts ...option.RequestOption) (res *MarketplaceProductSimulateImpactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("marketplace/products/%v/impact-simulate", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MarketplaceProductListResponse struct {
	Data []MarketplaceProductListResponseData `json:"data"`
	JSON marketplaceProductListResponseJSON   `json:"-"`
	PaginatedList
}

// marketplaceProductListResponseJSON contains the JSON metadata for the struct
// [MarketplaceProductListResponse]
type marketplaceProductListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketplaceProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceProductListResponseJSON) RawJSON() string {
	return r.raw
}

type MarketplaceProductListResponseData struct {
	// Unique identifier for the marketplace product.
	ID interface{} `json:"id,required"`
	// AI's score for how well this product is personalized to the user (0-1).
	AIPersonalizationScore interface{} `json:"aiPersonalizationScore,required"`
	// Category of the product/service.
	Category MarketplaceProductListResponseDataCategory `json:"category,required"`
	// Detailed description of the product/service.
	Description interface{} `json:"description,required"`
	// URL to an image representing the product.
	ImageURL interface{} `json:"imageUrl,required"`
	// Name of the product/service.
	Name interface{} `json:"name,required"`
	// Pricing information (can be a range or fixed text).
	Price interface{} `json:"price,required"`
	// Provider or vendor of the product/service.
	Provider interface{} `json:"provider,required"`
	// Average user rating for the product (0-5).
	Rating interface{} `json:"rating,required"`
	// AI-generated explanation for recommending this product.
	AIRecommendationReason interface{} `json:"aiRecommendationReason"`
	// Details of any special offers associated with the product.
	OfferDetails MarketplaceProductListResponseDataOfferDetails `json:"offerDetails"`
	// Direct URL to the product on the provider's website.
	ProductURL interface{}                            `json:"productUrl"`
	JSON       marketplaceProductListResponseDataJSON `json:"-"`
}

// marketplaceProductListResponseDataJSON contains the JSON metadata for the struct
// [MarketplaceProductListResponseData]
type marketplaceProductListResponseDataJSON struct {
	ID                     apijson.Field
	AIPersonalizationScore apijson.Field
	Category               apijson.Field
	Description            apijson.Field
	ImageURL               apijson.Field
	Name                   apijson.Field
	Price                  apijson.Field
	Provider               apijson.Field
	Rating                 apijson.Field
	AIRecommendationReason apijson.Field
	OfferDetails           apijson.Field
	ProductURL             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MarketplaceProductListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceProductListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Category of the product/service.
type MarketplaceProductListResponseDataCategory string

const (
	MarketplaceProductListResponseDataCategoryLoans          MarketplaceProductListResponseDataCategory = "loans"
	MarketplaceProductListResponseDataCategoryInsurance      MarketplaceProductListResponseDataCategory = "insurance"
	MarketplaceProductListResponseDataCategoryCreditCards    MarketplaceProductListResponseDataCategory = "credit_cards"
	MarketplaceProductListResponseDataCategoryInvestments    MarketplaceProductListResponseDataCategory = "investments"
	MarketplaceProductListResponseDataCategoryBudgetingTools MarketplaceProductListResponseDataCategory = "budgeting_tools"
	MarketplaceProductListResponseDataCategorySmartHome      MarketplaceProductListResponseDataCategory = "smart_home"
	MarketplaceProductListResponseDataCategoryTravel         MarketplaceProductListResponseDataCategory = "travel"
	MarketplaceProductListResponseDataCategoryEducation      MarketplaceProductListResponseDataCategory = "education"
	MarketplaceProductListResponseDataCategoryHealth         MarketplaceProductListResponseDataCategory = "health"
)

func (r MarketplaceProductListResponseDataCategory) IsKnown() bool {
	switch r {
	case MarketplaceProductListResponseDataCategoryLoans, MarketplaceProductListResponseDataCategoryInsurance, MarketplaceProductListResponseDataCategoryCreditCards, MarketplaceProductListResponseDataCategoryInvestments, MarketplaceProductListResponseDataCategoryBudgetingTools, MarketplaceProductListResponseDataCategorySmartHome, MarketplaceProductListResponseDataCategoryTravel, MarketplaceProductListResponseDataCategoryEducation, MarketplaceProductListResponseDataCategoryHealth:
		return true
	}
	return false
}

// Details of any special offers associated with the product.
type MarketplaceProductListResponseDataOfferDetails struct {
	// Optional redemption code.
	Code  interface{}                                        `json:"code"`
	Type  MarketplaceProductListResponseDataOfferDetailsType `json:"type"`
	Value interface{}                                        `json:"value"`
	JSON  marketplaceProductListResponseDataOfferDetailsJSON `json:"-"`
}

// marketplaceProductListResponseDataOfferDetailsJSON contains the JSON metadata
// for the struct [MarketplaceProductListResponseDataOfferDetails]
type marketplaceProductListResponseDataOfferDetailsJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketplaceProductListResponseDataOfferDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceProductListResponseDataOfferDetailsJSON) RawJSON() string {
	return r.raw
}

type MarketplaceProductListResponseDataOfferDetailsType string

const (
	MarketplaceProductListResponseDataOfferDetailsTypeDiscount    MarketplaceProductListResponseDataOfferDetailsType = "discount"
	MarketplaceProductListResponseDataOfferDetailsTypeSpecialRate MarketplaceProductListResponseDataOfferDetailsType = "special_rate"
	MarketplaceProductListResponseDataOfferDetailsTypeFreeTrial   MarketplaceProductListResponseDataOfferDetailsType = "free_trial"
)

func (r MarketplaceProductListResponseDataOfferDetailsType) IsKnown() bool {
	switch r {
	case MarketplaceProductListResponseDataOfferDetailsTypeDiscount, MarketplaceProductListResponseDataOfferDetailsTypeSpecialRate, MarketplaceProductListResponseDataOfferDetailsTypeFreeTrial:
		return true
	}
	return false
}

type MarketplaceProductSimulateImpactResponse struct {
	// Key financial impacts identified by the AI (e.g., on cash flow, debt-to-income).
	KeyImpacts []MarketplaceProductSimulateImpactResponseKeyImpact `json:"keyImpacts,required"`
	// A natural language summary of the simulation's results for this product.
	NarrativeSummary interface{} `json:"narrativeSummary,required"`
	// The ID of the marketplace product being simulated.
	ProductID interface{} `json:"productId,required"`
	// Unique identifier for the simulation performed.
	SimulationID interface{} `json:"simulationId,required"`
	// Actionable recommendations or advice related to the product and its impact.
	AIRecommendations []AIInsight `json:"aiRecommendations,nullable"`
	// Projected amortization schedule for loan products.
	ProjectedAmortizationSchedule []MarketplaceProductSimulateImpactResponseProjectedAmortizationSchedule `json:"projectedAmortizationSchedule,nullable"`
	JSON                          marketplaceProductSimulateImpactResponseJSON                            `json:"-"`
}

// marketplaceProductSimulateImpactResponseJSON contains the JSON metadata for the
// struct [MarketplaceProductSimulateImpactResponse]
type marketplaceProductSimulateImpactResponseJSON struct {
	KeyImpacts                    apijson.Field
	NarrativeSummary              apijson.Field
	ProductID                     apijson.Field
	SimulationID                  apijson.Field
	AIRecommendations             apijson.Field
	ProjectedAmortizationSchedule apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *MarketplaceProductSimulateImpactResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceProductSimulateImpactResponseJSON) RawJSON() string {
	return r.raw
}

type MarketplaceProductSimulateImpactResponseKeyImpact struct {
	Metric   interface{}                                                `json:"metric"`
	Severity MarketplaceProductSimulateImpactResponseKeyImpactsSeverity `json:"severity"`
	Value    interface{}                                                `json:"value"`
	JSON     marketplaceProductSimulateImpactResponseKeyImpactJSON      `json:"-"`
}

// marketplaceProductSimulateImpactResponseKeyImpactJSON contains the JSON metadata
// for the struct [MarketplaceProductSimulateImpactResponseKeyImpact]
type marketplaceProductSimulateImpactResponseKeyImpactJSON struct {
	Metric      apijson.Field
	Severity    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketplaceProductSimulateImpactResponseKeyImpact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceProductSimulateImpactResponseKeyImpactJSON) RawJSON() string {
	return r.raw
}

type MarketplaceProductSimulateImpactResponseKeyImpactsSeverity string

const (
	MarketplaceProductSimulateImpactResponseKeyImpactsSeverityLow    MarketplaceProductSimulateImpactResponseKeyImpactsSeverity = "low"
	MarketplaceProductSimulateImpactResponseKeyImpactsSeverityMedium MarketplaceProductSimulateImpactResponseKeyImpactsSeverity = "medium"
	MarketplaceProductSimulateImpactResponseKeyImpactsSeverityHigh   MarketplaceProductSimulateImpactResponseKeyImpactsSeverity = "high"
)

func (r MarketplaceProductSimulateImpactResponseKeyImpactsSeverity) IsKnown() bool {
	switch r {
	case MarketplaceProductSimulateImpactResponseKeyImpactsSeverityLow, MarketplaceProductSimulateImpactResponseKeyImpactsSeverityMedium, MarketplaceProductSimulateImpactResponseKeyImpactsSeverityHigh:
		return true
	}
	return false
}

type MarketplaceProductSimulateImpactResponseProjectedAmortizationSchedule struct {
	Interest         interface{}                                                               `json:"interest"`
	Month            interface{}                                                               `json:"month"`
	Payment          interface{}                                                               `json:"payment"`
	Principal        interface{}                                                               `json:"principal"`
	RemainingBalance interface{}                                                               `json:"remainingBalance"`
	JSON             marketplaceProductSimulateImpactResponseProjectedAmortizationScheduleJSON `json:"-"`
}

// marketplaceProductSimulateImpactResponseProjectedAmortizationScheduleJSON
// contains the JSON metadata for the struct
// [MarketplaceProductSimulateImpactResponseProjectedAmortizationSchedule]
type marketplaceProductSimulateImpactResponseProjectedAmortizationScheduleJSON struct {
	Interest         apijson.Field
	Month            apijson.Field
	Payment          apijson.Field
	Principal        apijson.Field
	RemainingBalance apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MarketplaceProductSimulateImpactResponseProjectedAmortizationSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceProductSimulateImpactResponseProjectedAmortizationScheduleJSON) RawJSON() string {
	return r.raw
}

type MarketplaceProductListParams struct {
	// Filter by AI personalization level (e.g., low, medium, high). 'High' means
	// highly relevant to user's specific needs.
	AIPersonalizationLevel param.Field[MarketplaceProductListParamsAIPersonalizationLevel] `query:"aiPersonalizationLevel"`
	// Filter products by category (e.g., loans, insurance, credit_cards, investments).
	Category param.Field[MarketplaceProductListParamsCategory] `query:"category"`
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Minimum user rating for products (0-5).
	MinRating param.Field[interface{}] `query:"minRating"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [MarketplaceProductListParams]'s query parameters as
// `url.Values`.
func (r MarketplaceProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by AI personalization level (e.g., low, medium, high). 'High' means
// highly relevant to user's specific needs.
type MarketplaceProductListParamsAIPersonalizationLevel string

const (
	MarketplaceProductListParamsAIPersonalizationLevelLow    MarketplaceProductListParamsAIPersonalizationLevel = "low"
	MarketplaceProductListParamsAIPersonalizationLevelMedium MarketplaceProductListParamsAIPersonalizationLevel = "medium"
	MarketplaceProductListParamsAIPersonalizationLevelHigh   MarketplaceProductListParamsAIPersonalizationLevel = "high"
)

func (r MarketplaceProductListParamsAIPersonalizationLevel) IsKnown() bool {
	switch r {
	case MarketplaceProductListParamsAIPersonalizationLevelLow, MarketplaceProductListParamsAIPersonalizationLevelMedium, MarketplaceProductListParamsAIPersonalizationLevelHigh:
		return true
	}
	return false
}

// Filter products by category (e.g., loans, insurance, credit_cards, investments).
type MarketplaceProductListParamsCategory string

const (
	MarketplaceProductListParamsCategoryLoans          MarketplaceProductListParamsCategory = "loans"
	MarketplaceProductListParamsCategoryInsurance      MarketplaceProductListParamsCategory = "insurance"
	MarketplaceProductListParamsCategoryCreditCards    MarketplaceProductListParamsCategory = "credit_cards"
	MarketplaceProductListParamsCategoryInvestments    MarketplaceProductListParamsCategory = "investments"
	MarketplaceProductListParamsCategoryBudgetingTools MarketplaceProductListParamsCategory = "budgeting_tools"
	MarketplaceProductListParamsCategorySmartHome      MarketplaceProductListParamsCategory = "smart_home"
	MarketplaceProductListParamsCategoryTravel         MarketplaceProductListParamsCategory = "travel"
	MarketplaceProductListParamsCategoryEducation      MarketplaceProductListParamsCategory = "education"
)

func (r MarketplaceProductListParamsCategory) IsKnown() bool {
	switch r {
	case MarketplaceProductListParamsCategoryLoans, MarketplaceProductListParamsCategoryInsurance, MarketplaceProductListParamsCategoryCreditCards, MarketplaceProductListParamsCategoryInvestments, MarketplaceProductListParamsCategoryBudgetingTools, MarketplaceProductListParamsCategorySmartHome, MarketplaceProductListParamsCategoryTravel, MarketplaceProductListParamsCategoryEducation:
		return true
	}
	return false
}

type MarketplaceProductSimulateImpactParams struct {
	// Dynamic parameters specific to the product type (e.g., loan amount, investment
	// term).
	SimulationParameters param.Field[interface{}] `json:"simulationParameters"`
}

func (r MarketplaceProductSimulateImpactParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
