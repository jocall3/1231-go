// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// PaymentInternationalService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentInternationalService] method instead.
type PaymentInternationalService struct {
	Options []option.RequestOption
}

// NewPaymentInternationalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentInternationalService(opts ...option.RequestOption) (r *PaymentInternationalService) {
	r = &PaymentInternationalService{}
	r.Options = opts
	return
}

// Facilitates the secure initiation of an international wire transfer to a
// beneficiary in another country and currency, leveraging optimal FX rates and
// tracking capabilities.
func (r *PaymentInternationalService) Initiate(ctx context.Context, body PaymentInternationalInitiateParams, opts ...option.RequestOption) (res *InternationalPaymentStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/international/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the current processing status and details of an initiated
// international payment.
func (r *PaymentInternationalService) GetStatus(ctx context.Context, paymentID interface{}, opts ...option.RequestOption) (res *InternationalPaymentStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("payments/international/%v/status", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InternationalPaymentStatus struct {
	// The foreign exchange rate applied (target per source currency).
	FxRateApplied interface{} `json:"fxRateApplied,required"`
	// Unique identifier for the international payment.
	PaymentID interface{} `json:"paymentId,required"`
	// The amount sent in the source currency.
	SourceAmount interface{} `json:"sourceAmount,required"`
	// The source currency code.
	SourceCurrency interface{} `json:"sourceCurrency,required"`
	// Current processing status of the payment.
	Status InternationalPaymentStatusStatus `json:"status,required"`
	// The amount received by the beneficiary in the target currency.
	TargetAmount interface{} `json:"targetAmount,required"`
	// The target currency code.
	TargetCurrency interface{} `json:"targetCurrency,required"`
	// Estimated time when the payment will be completed.
	EstimatedCompletionTime interface{} `json:"estimatedCompletionTime"`
	// Total fees applied to the payment.
	FeesApplied interface{} `json:"feesApplied"`
	// An optional message providing more context on the status (e.g., reason for
	// hold).
	Message interface{} `json:"message"`
	// URL to track the payment's progress.
	TrackingURL interface{}                    `json:"trackingUrl"`
	JSON        internationalPaymentStatusJSON `json:"-"`
}

// internationalPaymentStatusJSON contains the JSON metadata for the struct
// [InternationalPaymentStatus]
type internationalPaymentStatusJSON struct {
	FxRateApplied           apijson.Field
	PaymentID               apijson.Field
	SourceAmount            apijson.Field
	SourceCurrency          apijson.Field
	Status                  apijson.Field
	TargetAmount            apijson.Field
	TargetCurrency          apijson.Field
	EstimatedCompletionTime apijson.Field
	FeesApplied             apijson.Field
	Message                 apijson.Field
	TrackingURL             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InternationalPaymentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internationalPaymentStatusJSON) RawJSON() string {
	return r.raw
}

// Current processing status of the payment.
type InternationalPaymentStatusStatus string

const (
	InternationalPaymentStatusStatusInProgress    InternationalPaymentStatusStatus = "in_progress"
	InternationalPaymentStatusStatusHeldForReview InternationalPaymentStatusStatus = "held_for_review"
	InternationalPaymentStatusStatusCompleted     InternationalPaymentStatusStatus = "completed"
	InternationalPaymentStatusStatusFailed        InternationalPaymentStatusStatus = "failed"
	InternationalPaymentStatusStatusCancelled     InternationalPaymentStatusStatus = "cancelled"
)

func (r InternationalPaymentStatusStatus) IsKnown() bool {
	switch r {
	case InternationalPaymentStatusStatusInProgress, InternationalPaymentStatusStatusHeldForReview, InternationalPaymentStatusStatusCompleted, InternationalPaymentStatusStatusFailed, InternationalPaymentStatusStatusCancelled:
		return true
	}
	return false
}

type PaymentInternationalInitiateParams struct {
	// The amount to send in the source currency.
	Amount param.Field[interface{}] `json:"amount,required"`
	// Details of the payment beneficiary.
	Beneficiary param.Field[PaymentInternationalInitiateParamsBeneficiary] `json:"beneficiary,required"`
	// Purpose of the payment.
	Purpose param.Field[interface{}] `json:"purpose,required"`
	// The ID of the user's source account for the payment.
	SourceAccountID param.Field[interface{}] `json:"sourceAccountId,required"`
	// The ISO 4217 currency code of the source funds.
	SourceCurrency param.Field[interface{}] `json:"sourceCurrency,required"`
	// The ISO 4217 currency code for the beneficiary's currency.
	TargetCurrency param.Field[interface{}] `json:"targetCurrency,required"`
	// If true, attempts to lock the quoted FX rate for a short period.
	FxRateLock param.Field[interface{}] `json:"fxRateLock"`
	// Indicates whether to use AI-optimized FX rates or standard market rates.
	FxRateProvider param.Field[PaymentInternationalInitiateParamsFxRateProvider] `json:"fxRateProvider"`
	// Optional: Your internal reference for this payment.
	Reference param.Field[interface{}] `json:"reference"`
}

func (r PaymentInternationalInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the payment beneficiary.
type PaymentInternationalInitiateParamsBeneficiary struct {
	// Full address of the beneficiary.
	Address param.Field[interface{}] `json:"address,required"`
	// Name of the beneficiary's bank.
	BankName param.Field[interface{}] `json:"bankName,required"`
	// Full name of the beneficiary.
	Name param.Field[interface{}] `json:"name,required"`
	// Account number (if IBAN/SWIFT not applicable).
	AccountNumber param.Field[interface{}] `json:"accountNumber"`
	// IBAN for Eurozone transfers.
	Iban param.Field[interface{}] `json:"iban"`
	// Routing number (if applicable, e.g., for US transfers).
	RoutingNumber param.Field[interface{}] `json:"routingNumber"`
	// SWIFT/BIC code for international transfers.
	SwiftBic param.Field[interface{}] `json:"swiftBic"`
}

func (r PaymentInternationalInitiateParamsBeneficiary) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates whether to use AI-optimized FX rates or standard market rates.
type PaymentInternationalInitiateParamsFxRateProvider string

const (
	PaymentInternationalInitiateParamsFxRateProviderProprietaryAI PaymentInternationalInitiateParamsFxRateProvider = "proprietary_ai"
	PaymentInternationalInitiateParamsFxRateProviderMarketRate    PaymentInternationalInitiateParamsFxRateProvider = "market_rate"
)

func (r PaymentInternationalInitiateParamsFxRateProvider) IsKnown() bool {
	switch r {
	case PaymentInternationalInitiateParamsFxRateProviderProprietaryAI, PaymentInternationalInitiateParamsFxRateProviderMarketRate:
		return true
	}
	return false
}
