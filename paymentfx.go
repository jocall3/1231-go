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

// PaymentFxService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentFxService] method instead.
type PaymentFxService struct {
	Options []option.RequestOption
}

// NewPaymentFxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentFxService(opts ...option.RequestOption) (r *PaymentFxService) {
	r = &PaymentFxService{}
	r.Options = opts
	return
}

// Executes an instant currency conversion between two currencies, either from a
// balance or into a specified account.
func (r *PaymentFxService) Convert(ctx context.Context, body PaymentFxConvertParams, opts ...option.RequestOption) (res *PaymentFxConvertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/fx/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves current and AI-predicted future foreign exchange rates for a specified
// currency pair, including bid/ask spreads and historical volatility data for
// informed decisions.
func (r *PaymentFxService) GetRates(ctx context.Context, query PaymentFxGetRatesParams, opts ...option.RequestOption) (res *PaymentFxGetRatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/fx/rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaymentFxConvertResponse struct {
	// Unique identifier for the currency conversion.
	ConversionID interface{} `json:"conversionId,required"`
	// Timestamp when the conversion was completed.
	ConversionTimestamp interface{} `json:"conversionTimestamp,required"`
	// The foreign exchange rate applied (target per source currency).
	FxRateApplied interface{} `json:"fxRateApplied,required"`
	// The amount converted from the source currency.
	SourceAmount interface{} `json:"sourceAmount,required"`
	// The source currency code.
	SourceCurrency interface{} `json:"sourceCurrency,required"`
	// Status of the currency conversion.
	Status PaymentFxConvertResponseStatus `json:"status,required"`
	// The amount converted into the target currency.
	TargetAmount interface{} `json:"targetAmount,required"`
	// Any fees applied to the conversion.
	FeesApplied interface{} `json:"feesApplied"`
	// The target currency code.
	TargetCurrency interface{} `json:"targetCurrency"`
	// The ID of the internal transaction representing this conversion.
	TransactionID interface{}                  `json:"transactionId"`
	JSON          paymentFxConvertResponseJSON `json:"-"`
}

// paymentFxConvertResponseJSON contains the JSON metadata for the struct
// [PaymentFxConvertResponse]
type paymentFxConvertResponseJSON struct {
	ConversionID        apijson.Field
	ConversionTimestamp apijson.Field
	FxRateApplied       apijson.Field
	SourceAmount        apijson.Field
	SourceCurrency      apijson.Field
	Status              apijson.Field
	TargetAmount        apijson.Field
	FeesApplied         apijson.Field
	TargetCurrency      apijson.Field
	TransactionID       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentFxConvertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxConvertResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the currency conversion.
type PaymentFxConvertResponseStatus string

const (
	PaymentFxConvertResponseStatusCompleted PaymentFxConvertResponseStatus = "completed"
	PaymentFxConvertResponseStatusPending   PaymentFxConvertResponseStatus = "pending"
	PaymentFxConvertResponseStatusFailed    PaymentFxConvertResponseStatus = "failed"
)

func (r PaymentFxConvertResponseStatus) IsKnown() bool {
	switch r {
	case PaymentFxConvertResponseStatusCompleted, PaymentFxConvertResponseStatusPending, PaymentFxConvertResponseStatusFailed:
		return true
	}
	return false
}

type PaymentFxGetRatesResponse struct {
	// The base currency code.
	BaseCurrency interface{} `json:"baseCurrency,required"`
	// Real-time foreign exchange rates.
	CurrentRate PaymentFxGetRatesResponseCurrentRate `json:"currentRate,required"`
	// The target currency code.
	TargetCurrency       interface{}                                   `json:"targetCurrency,required"`
	HistoricalVolatility PaymentFxGetRatesResponseHistoricalVolatility `json:"historicalVolatility"`
	// AI-predicted foreign exchange rates for future dates.
	PredictiveRates []PaymentFxGetRatesResponsePredictiveRate `json:"predictiveRates,nullable"`
	JSON            paymentFxGetRatesResponseJSON             `json:"-"`
}

// paymentFxGetRatesResponseJSON contains the JSON metadata for the struct
// [PaymentFxGetRatesResponse]
type paymentFxGetRatesResponseJSON struct {
	BaseCurrency         apijson.Field
	CurrentRate          apijson.Field
	TargetCurrency       apijson.Field
	HistoricalVolatility apijson.Field
	PredictiveRates      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseJSON) RawJSON() string {
	return r.raw
}

// Real-time foreign exchange rates.
type PaymentFxGetRatesResponseCurrentRate struct {
	// Current ask rate (price at which a currency dealer will sell the base currency).
	Ask interface{} `json:"ask"`
	// Current bid rate (price at which a currency dealer will buy the base currency).
	Bid interface{} `json:"bid"`
	// Mid-market rate (average of bid and ask).
	Mid interface{} `json:"mid"`
	// Timestamp of the current rate.
	Timestamp interface{}                              `json:"timestamp"`
	JSON      paymentFxGetRatesResponseCurrentRateJSON `json:"-"`
}

// paymentFxGetRatesResponseCurrentRateJSON contains the JSON metadata for the
// struct [PaymentFxGetRatesResponseCurrentRate]
type paymentFxGetRatesResponseCurrentRateJSON struct {
	Ask         apijson.Field
	Bid         apijson.Field
	Mid         apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponseCurrentRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseCurrentRateJSON) RawJSON() string {
	return r.raw
}

type PaymentFxGetRatesResponseHistoricalVolatility struct {
	// Historical volatility over the past 30 days.
	Past30Days interface{} `json:"past30Days"`
	// Historical volatility over the past 7 days.
	Past7Days interface{}                                       `json:"past7Days"`
	JSON      paymentFxGetRatesResponseHistoricalVolatilityJSON `json:"-"`
}

// paymentFxGetRatesResponseHistoricalVolatilityJSON contains the JSON metadata for
// the struct [PaymentFxGetRatesResponseHistoricalVolatility]
type paymentFxGetRatesResponseHistoricalVolatilityJSON struct {
	Past30Days  apijson.Field
	Past7Days   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponseHistoricalVolatility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseHistoricalVolatilityJSON) RawJSON() string {
	return r.raw
}

type PaymentFxGetRatesResponsePredictiveRate struct {
	// AI model's confidence in the prediction (0-1).
	AIModelConfidence interface{} `json:"aiModelConfidence"`
	// Lower bound of the AI's confidence interval for the predicted rate.
	ConfidenceIntervalLower interface{} `json:"confidenceIntervalLower"`
	// Upper bound of the AI's confidence interval for the predicted rate.
	ConfidenceIntervalUpper interface{} `json:"confidenceIntervalUpper"`
	// Date for the predicted rate.
	Date interface{} `json:"date"`
	// AI-predicted mid-market rate.
	PredictedMidRate interface{}                                 `json:"predictedMidRate"`
	JSON             paymentFxGetRatesResponsePredictiveRateJSON `json:"-"`
}

// paymentFxGetRatesResponsePredictiveRateJSON contains the JSON metadata for the
// struct [PaymentFxGetRatesResponsePredictiveRate]
type paymentFxGetRatesResponsePredictiveRateJSON struct {
	AIModelConfidence       apijson.Field
	ConfidenceIntervalLower apijson.Field
	ConfidenceIntervalUpper apijson.Field
	Date                    apijson.Field
	PredictedMidRate        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponsePredictiveRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponsePredictiveRateJSON) RawJSON() string {
	return r.raw
}

type PaymentFxConvertParams struct {
	// The ID of the account from which funds will be converted.
	SourceAccountID param.Field[interface{}] `json:"sourceAccountId,required"`
	// The amount to convert from the source currency.
	SourceAmount param.Field[interface{}] `json:"sourceAmount,required"`
	// The ISO 4217 currency code of the source funds.
	SourceCurrency param.Field[interface{}] `json:"sourceCurrency,required"`
	// The ISO 4217 currency code for the target currency.
	TargetCurrency param.Field[interface{}] `json:"targetCurrency,required"`
	// If true, attempts to lock the quoted FX rate for a short period.
	FxRateLock param.Field[interface{}] `json:"fxRateLock"`
	// Optional: The ID of the account to deposit the converted funds. If null, funds
	// are held in a wallet/balance.
	TargetAccountID param.Field[interface{}] `json:"targetAccountId"`
}

func (r PaymentFxConvertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentFxGetRatesParams struct {
	// The base currency code (e.g., USD).
	BaseCurrency param.Field[interface{}] `query:"baseCurrency,required"`
	// The target currency code (e.g., EUR).
	TargetCurrency param.Field[interface{}] `query:"targetCurrency,required"`
	// Number of days into the future to provide an AI-driven prediction.
	ForecastDays param.Field[interface{}] `query:"forecastDays"`
}

// URLQuery serializes [PaymentFxGetRatesParams]'s query parameters as
// `url.Values`.
func (r PaymentFxGetRatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
