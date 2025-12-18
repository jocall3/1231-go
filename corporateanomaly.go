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

// CorporateAnomalyService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateAnomalyService] method instead.
type CorporateAnomalyService struct {
	Options []option.RequestOption
}

// NewCorporateAnomalyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateAnomalyService(opts ...option.RequestOption) (r *CorporateAnomalyService) {
	r = &CorporateAnomalyService{}
	r.Options = opts
	return
}

// Retrieves a comprehensive list of AI-detected financial anomalies across
// transactions, payments, and corporate cards that require immediate review and
// potential action to mitigate risk and ensure compliance.
func (r *CorporateAnomalyService) List(ctx context.Context, query CorporateAnomalyListParams, opts ...option.RequestOption) (res *CorporateAnomalyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the review status of a specific financial anomaly, allowing compliance
// officers to mark it as dismissed, resolved, or escalate for further
// investigation after thorough AI-assisted and human review.
func (r *CorporateAnomalyService) UpdateStatus(ctx context.Context, anomalyID interface{}, body CorporateAnomalyUpdateStatusParams, opts ...option.RequestOption) (res *FinancialAnomaly, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("corporate/anomalies/%v/status", anomalyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FinancialAnomaly struct {
	// Unique identifier for the detected anomaly.
	ID interface{} `json:"id,required"`
	// AI's confidence in its detection of the anomaly (0-1).
	AIConfidenceScore interface{} `json:"aiConfidenceScore,required"`
	// A brief summary of the anomaly.
	Description interface{} `json:"description,required"`
	// The ID of the specific entity (e.g., transaction, user, card) the anomaly is
	// linked to.
	EntityID interface{} `json:"entityId,required"`
	// The type of financial entity related to the anomaly.
	EntityType FinancialAnomalyEntityType `json:"entityType,required"`
	// AI-recommended immediate action to address the anomaly.
	RecommendedAction interface{} `json:"recommendedAction,required"`
	// AI-assigned risk score (0-100), higher is more risky.
	RiskScore interface{} `json:"riskScore,required"`
	// AI-assessed severity of the anomaly.
	Severity FinancialAnomalySeverity `json:"severity,required"`
	// Current review status of the anomaly.
	Status FinancialAnomalyStatus `json:"status,required"`
	// Timestamp when the anomaly was detected.
	Timestamp interface{} `json:"timestamp,required"`
	// Detailed context and reasoning behind the anomaly detection.
	Details interface{} `json:"details"`
	// List of IDs of other transactions or entities related to this anomaly.
	RelatedTransactions []interface{} `json:"relatedTransactions,nullable"`
	// Notes recorded during the resolution or dismissal of the anomaly.
	ResolutionNotes interface{}          `json:"resolutionNotes"`
	JSON            financialAnomalyJSON `json:"-"`
}

// financialAnomalyJSON contains the JSON metadata for the struct
// [FinancialAnomaly]
type financialAnomalyJSON struct {
	ID                  apijson.Field
	AIConfidenceScore   apijson.Field
	Description         apijson.Field
	EntityID            apijson.Field
	EntityType          apijson.Field
	RecommendedAction   apijson.Field
	RiskScore           apijson.Field
	Severity            apijson.Field
	Status              apijson.Field
	Timestamp           apijson.Field
	Details             apijson.Field
	RelatedTransactions apijson.Field
	ResolutionNotes     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *FinancialAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAnomalyJSON) RawJSON() string {
	return r.raw
}

// The type of financial entity related to the anomaly.
type FinancialAnomalyEntityType string

const (
	FinancialAnomalyEntityTypePaymentOrder  FinancialAnomalyEntityType = "PaymentOrder"
	FinancialAnomalyEntityTypeTransaction   FinancialAnomalyEntityType = "Transaction"
	FinancialAnomalyEntityTypeCounterparty  FinancialAnomalyEntityType = "Counterparty"
	FinancialAnomalyEntityTypeCorporateCard FinancialAnomalyEntityType = "CorporateCard"
	FinancialAnomalyEntityTypeUser          FinancialAnomalyEntityType = "User"
	FinancialAnomalyEntityTypeInvoice       FinancialAnomalyEntityType = "Invoice"
)

func (r FinancialAnomalyEntityType) IsKnown() bool {
	switch r {
	case FinancialAnomalyEntityTypePaymentOrder, FinancialAnomalyEntityTypeTransaction, FinancialAnomalyEntityTypeCounterparty, FinancialAnomalyEntityTypeCorporateCard, FinancialAnomalyEntityTypeUser, FinancialAnomalyEntityTypeInvoice:
		return true
	}
	return false
}

// AI-assessed severity of the anomaly.
type FinancialAnomalySeverity string

const (
	FinancialAnomalySeverityLow      FinancialAnomalySeverity = "Low"
	FinancialAnomalySeverityMedium   FinancialAnomalySeverity = "Medium"
	FinancialAnomalySeverityHigh     FinancialAnomalySeverity = "High"
	FinancialAnomalySeverityCritical FinancialAnomalySeverity = "Critical"
)

func (r FinancialAnomalySeverity) IsKnown() bool {
	switch r {
	case FinancialAnomalySeverityLow, FinancialAnomalySeverityMedium, FinancialAnomalySeverityHigh, FinancialAnomalySeverityCritical:
		return true
	}
	return false
}

// Current review status of the anomaly.
type FinancialAnomalyStatus string

const (
	FinancialAnomalyStatusNew         FinancialAnomalyStatus = "New"
	FinancialAnomalyStatusUnderReview FinancialAnomalyStatus = "Under Review"
	FinancialAnomalyStatusEscalated   FinancialAnomalyStatus = "Escalated"
	FinancialAnomalyStatusDismissed   FinancialAnomalyStatus = "Dismissed"
	FinancialAnomalyStatusResolved    FinancialAnomalyStatus = "Resolved"
)

func (r FinancialAnomalyStatus) IsKnown() bool {
	switch r {
	case FinancialAnomalyStatusNew, FinancialAnomalyStatusUnderReview, FinancialAnomalyStatusEscalated, FinancialAnomalyStatusDismissed, FinancialAnomalyStatusResolved:
		return true
	}
	return false
}

type CorporateAnomalyListResponse struct {
	Data []FinancialAnomaly               `json:"data"`
	JSON corporateAnomalyListResponseJSON `json:"-"`
	PaginatedList
}

// corporateAnomalyListResponseJSON contains the JSON metadata for the struct
// [CorporateAnomalyListResponse]
type corporateAnomalyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateAnomalyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateAnomalyListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateAnomalyListParams struct {
	// End date for filtering results (inclusive, YYYY-MM-DD).
	EndDate param.Field[interface{}] `query:"endDate"`
	// Filter anomalies by the type of financial entity they are related to.
	EntityType param.Field[CorporateAnomalyListParamsEntityType] `query:"entityType"`
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
	// Filter anomalies by their AI-assessed severity level.
	Severity param.Field[CorporateAnomalyListParamsSeverity] `query:"severity"`
	// Start date for filtering results (inclusive, YYYY-MM-DD).
	StartDate param.Field[interface{}] `query:"startDate"`
	// Filter anomalies by their current review status.
	Status param.Field[CorporateAnomalyListParamsStatus] `query:"status"`
}

// URLQuery serializes [CorporateAnomalyListParams]'s query parameters as
// `url.Values`.
func (r CorporateAnomalyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter anomalies by the type of financial entity they are related to.
type CorporateAnomalyListParamsEntityType string

const (
	CorporateAnomalyListParamsEntityTypePaymentOrder  CorporateAnomalyListParamsEntityType = "PaymentOrder"
	CorporateAnomalyListParamsEntityTypeTransaction   CorporateAnomalyListParamsEntityType = "Transaction"
	CorporateAnomalyListParamsEntityTypeCounterparty  CorporateAnomalyListParamsEntityType = "Counterparty"
	CorporateAnomalyListParamsEntityTypeCorporateCard CorporateAnomalyListParamsEntityType = "CorporateCard"
	CorporateAnomalyListParamsEntityTypeInvoice       CorporateAnomalyListParamsEntityType = "Invoice"
)

func (r CorporateAnomalyListParamsEntityType) IsKnown() bool {
	switch r {
	case CorporateAnomalyListParamsEntityTypePaymentOrder, CorporateAnomalyListParamsEntityTypeTransaction, CorporateAnomalyListParamsEntityTypeCounterparty, CorporateAnomalyListParamsEntityTypeCorporateCard, CorporateAnomalyListParamsEntityTypeInvoice:
		return true
	}
	return false
}

// Filter anomalies by their AI-assessed severity level.
type CorporateAnomalyListParamsSeverity string

const (
	CorporateAnomalyListParamsSeverityLow      CorporateAnomalyListParamsSeverity = "Low"
	CorporateAnomalyListParamsSeverityMedium   CorporateAnomalyListParamsSeverity = "Medium"
	CorporateAnomalyListParamsSeverityHigh     CorporateAnomalyListParamsSeverity = "High"
	CorporateAnomalyListParamsSeverityCritical CorporateAnomalyListParamsSeverity = "Critical"
)

func (r CorporateAnomalyListParamsSeverity) IsKnown() bool {
	switch r {
	case CorporateAnomalyListParamsSeverityLow, CorporateAnomalyListParamsSeverityMedium, CorporateAnomalyListParamsSeverityHigh, CorporateAnomalyListParamsSeverityCritical:
		return true
	}
	return false
}

// Filter anomalies by their current review status.
type CorporateAnomalyListParamsStatus string

const (
	CorporateAnomalyListParamsStatusNew         CorporateAnomalyListParamsStatus = "New"
	CorporateAnomalyListParamsStatusUnderReview CorporateAnomalyListParamsStatus = "Under Review"
	CorporateAnomalyListParamsStatusEscalated   CorporateAnomalyListParamsStatus = "Escalated"
	CorporateAnomalyListParamsStatusDismissed   CorporateAnomalyListParamsStatus = "Dismissed"
	CorporateAnomalyListParamsStatusResolved    CorporateAnomalyListParamsStatus = "Resolved"
)

func (r CorporateAnomalyListParamsStatus) IsKnown() bool {
	switch r {
	case CorporateAnomalyListParamsStatusNew, CorporateAnomalyListParamsStatusUnderReview, CorporateAnomalyListParamsStatusEscalated, CorporateAnomalyListParamsStatusDismissed, CorporateAnomalyListParamsStatusResolved:
		return true
	}
	return false
}

type CorporateAnomalyUpdateStatusParams struct {
	// The new status for the financial anomaly.
	Status param.Field[CorporateAnomalyUpdateStatusParamsStatus] `json:"status,required"`
	// Optional notes regarding the resolution or dismissal of the anomaly.
	ResolutionNotes param.Field[interface{}] `json:"resolutionNotes"`
}

func (r CorporateAnomalyUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The new status for the financial anomaly.
type CorporateAnomalyUpdateStatusParamsStatus string

const (
	CorporateAnomalyUpdateStatusParamsStatusDismissed   CorporateAnomalyUpdateStatusParamsStatus = "Dismissed"
	CorporateAnomalyUpdateStatusParamsStatusResolved    CorporateAnomalyUpdateStatusParamsStatus = "Resolved"
	CorporateAnomalyUpdateStatusParamsStatusUnderReview CorporateAnomalyUpdateStatusParamsStatus = "Under Review"
	CorporateAnomalyUpdateStatusParamsStatusEscalated   CorporateAnomalyUpdateStatusParamsStatus = "Escalated"
)

func (r CorporateAnomalyUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case CorporateAnomalyUpdateStatusParamsStatusDismissed, CorporateAnomalyUpdateStatusParamsStatusResolved, CorporateAnomalyUpdateStatusParamsStatusUnderReview, CorporateAnomalyUpdateStatusParamsStatusEscalated:
		return true
	}
	return false
}
