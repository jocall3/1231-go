// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// SustainabilityService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityService] method instead.
type SustainabilityService struct {
	Options     []option.RequestOption
	Investments *SustainabilityInvestmentService
}

// NewSustainabilityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSustainabilityService(opts ...option.RequestOption) (r *SustainabilityService) {
	r = &SustainabilityService{}
	r.Options = opts
	r.Investments = NewSustainabilityInvestmentService(opts...)
	return
}

// Allows users to purchase carbon offsets to neutralize their estimated carbon
// footprint, supporting environmental initiatives.
func (r *SustainabilityService) PurchaseCarbonOffsets(ctx context.Context, body SustainabilityPurchaseCarbonOffsetsParams, opts ...option.RequestOption) (res *SustainabilityPurchaseCarbonOffsetsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/carbon-offsets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generates a detailed report of the user's estimated carbon footprint based on
// transaction data, lifestyle choices, and AI-driven impact assessments, offering
// insights and reduction strategies.
func (r *SustainabilityService) GetCarbonFootprint(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityGetCarbonFootprintResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/carbon-footprint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityPurchaseCarbonOffsetsResponse struct {
	// The amount of carbon dioxide equivalent offset by this purchase.
	AmountOffsetKgCo2e interface{} `json:"amountOffsetKgCO2e,required"`
	// Timestamp of the purchase.
	PurchaseDate interface{} `json:"purchaseDate,required"`
	// Unique identifier for the carbon offset purchase.
	PurchaseID interface{} `json:"purchaseId,required"`
	// Total cost of the carbon offset purchase in USD.
	TotalCostUsd interface{} `json:"totalCostUSD,required"`
	// URL to the official carbon offset certificate.
	CertificateURL interface{} `json:"certificateUrl"`
	// The carbon offset project supported.
	ProjectSupported interface{} `json:"projectSupported"`
	// The ID of the internal financial transaction for this purchase.
	TransactionID interface{}                                     `json:"transactionId"`
	JSON          sustainabilityPurchaseCarbonOffsetsResponseJSON `json:"-"`
}

// sustainabilityPurchaseCarbonOffsetsResponseJSON contains the JSON metadata for
// the struct [SustainabilityPurchaseCarbonOffsetsResponse]
type sustainabilityPurchaseCarbonOffsetsResponseJSON struct {
	AmountOffsetKgCo2e apijson.Field
	PurchaseDate       apijson.Field
	PurchaseID         apijson.Field
	TotalCostUsd       apijson.Field
	CertificateURL     apijson.Field
	ProjectSupported   apijson.Field
	TransactionID      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SustainabilityPurchaseCarbonOffsetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityPurchaseCarbonOffsetsResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponse struct {
	// AI-driven insights and recommendations for reducing carbon footprint.
	AIInsights []AIInsight `json:"aiInsights,required"`
	// Breakdown of carbon footprint by spending categories.
	BreakdownByCategory []SustainabilityGetCarbonFootprintResponseBreakdownByCategory `json:"breakdownByCategory,required"`
	// The period covered by the report.
	Period interface{} `json:"period,required"`
	// Unique identifier for the carbon footprint report.
	ReportID interface{} `json:"reportId,required"`
	// Total estimated carbon footprint in kilograms of CO2 equivalent.
	TotalCarbonFootprintKgCo2e interface{} `json:"totalCarbonFootprintKgCO2e,required"`
	// Recommendations for purchasing carbon offsets.
	OffsetRecommendations []SustainabilityGetCarbonFootprintResponseOffsetRecommendation `json:"offsetRecommendations,nullable"`
	JSON                  sustainabilityGetCarbonFootprintResponseJSON                   `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseJSON contains the JSON metadata for the
// struct [SustainabilityGetCarbonFootprintResponse]
type sustainabilityGetCarbonFootprintResponseJSON struct {
	AIInsights                 apijson.Field
	BreakdownByCategory        apijson.Field
	Period                     apijson.Field
	ReportID                   apijson.Field
	TotalCarbonFootprintKgCo2e apijson.Field
	OffsetRecommendations      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponseBreakdownByCategory struct {
	CarbonFootprintKgCo2e interface{}                                                     `json:"carbonFootprintKgCO2e"`
	Category              interface{}                                                     `json:"category"`
	Percentage            interface{}                                                     `json:"percentage"`
	JSON                  sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON contains the
// JSON metadata for the struct
// [SustainabilityGetCarbonFootprintResponseBreakdownByCategory]
type sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON struct {
	CarbonFootprintKgCo2e apijson.Field
	Category              apijson.Field
	Percentage            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponseBreakdownByCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponseOffsetRecommendation struct {
	CostPerTonUsd      interface{}                                                      `json:"costPerTonUSD"`
	OffsetAmountKgCo2e interface{}                                                      `json:"offsetAmountKgCO2e"`
	Project            interface{}                                                      `json:"project"`
	TotalCostUsd       interface{}                                                      `json:"totalCostUSD"`
	JSON               sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON contains the
// JSON metadata for the struct
// [SustainabilityGetCarbonFootprintResponseOffsetRecommendation]
type sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON struct {
	CostPerTonUsd      apijson.Field
	OffsetAmountKgCo2e apijson.Field
	Project            apijson.Field
	TotalCostUsd       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponseOffsetRecommendation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON) RawJSON() string {
	return r.raw
}

type SustainabilityPurchaseCarbonOffsetsParams struct {
	// The amount of carbon dioxide equivalent to offset in kilograms.
	AmountKgCo2e param.Field[interface{}] `json:"amountKgCO2e,required"`
	// Optional: The specific carbon offset project to support.
	OffsetProject param.Field[interface{}] `json:"offsetProject,required"`
	// The ID of the user's account to use for payment.
	PaymentAccountID param.Field[interface{}] `json:"paymentAccountId,required"`
}

func (r SustainabilityPurchaseCarbonOffsetsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
