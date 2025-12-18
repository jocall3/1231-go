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

// CorporateRiskFraudRuleService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateRiskFraudRuleService] method instead.
type CorporateRiskFraudRuleService struct {
	Options []option.RequestOption
}

// NewCorporateRiskFraudRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateRiskFraudRuleService(opts ...option.RequestOption) (r *CorporateRiskFraudRuleService) {
	r = &CorporateRiskFraudRuleService{}
	r.Options = opts
	return
}

// Creates a new custom AI-powered fraud detection rule, allowing organizations to
// define specific criteria, risk scores, and automated responses to evolving
// threat landscapes.
func (r *CorporateRiskFraudRuleService) New(ctx context.Context, body CorporateRiskFraudRuleNewParams, opts ...option.RequestOption) (res *FraudRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/fraud/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing custom AI-powered fraud detection rule, modifying its
// criteria, actions, or status.
func (r *CorporateRiskFraudRuleService) Update(ctx context.Context, ruleID interface{}, body CorporateRiskFraudRuleUpdateParams, opts ...option.RequestOption) (res *FraudRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("corporate/risk/fraud/rules/%v", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of AI-powered fraud detection rules currently active for the
// organization, including their parameters, thresholds, and associated actions
// (e.g., flag, block, alert).
func (r *CorporateRiskFraudRuleService) List(ctx context.Context, query CorporateRiskFraudRuleListParams, opts ...option.RequestOption) (res *CorporateRiskFraudRuleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/fraud/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specific custom AI-powered fraud detection rule.
func (r *CorporateRiskFraudRuleService) Delete(ctx context.Context, ruleID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("corporate/risk/fraud/rules/%v", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FraudRule struct {
	// Unique identifier for the fraud detection rule.
	ID interface{} `json:"id,required"`
	// Action to take when a fraud rule is triggered.
	Action FraudRuleAction `json:"action,required"`
	// Timestamp when the rule was created.
	CreatedAt interface{} `json:"createdAt,required"`
	// Identifier of who created the rule (e.g., user ID, 'system:ai-risk-engine').
	CreatedBy interface{} `json:"createdBy,required"`
	// Criteria that define when a fraud rule should trigger.
	Criteria FraudRuleCriteria `json:"criteria,required"`
	// Detailed description of what the rule detects.
	Description interface{} `json:"description,required"`
	// Timestamp when the rule was last updated.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Name of the fraud rule.
	Name interface{} `json:"name,required"`
	// Severity level when this rule is triggered.
	Severity FraudRuleSeverity `json:"severity,required"`
	// Current status of the rule.
	Status FraudRuleStatus `json:"status,required"`
	JSON   fraudRuleJSON   `json:"-"`
}

// fraudRuleJSON contains the JSON metadata for the struct [FraudRule]
type fraudRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	CreatedAt   apijson.Field
	CreatedBy   apijson.Field
	Criteria    apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudRuleJSON) RawJSON() string {
	return r.raw
}

// Severity level when this rule is triggered.
type FraudRuleSeverity string

const (
	FraudRuleSeverityLow      FraudRuleSeverity = "Low"
	FraudRuleSeverityMedium   FraudRuleSeverity = "Medium"
	FraudRuleSeverityHigh     FraudRuleSeverity = "High"
	FraudRuleSeverityCritical FraudRuleSeverity = "Critical"
)

func (r FraudRuleSeverity) IsKnown() bool {
	switch r {
	case FraudRuleSeverityLow, FraudRuleSeverityMedium, FraudRuleSeverityHigh, FraudRuleSeverityCritical:
		return true
	}
	return false
}

// Current status of the rule.
type FraudRuleStatus string

const (
	FraudRuleStatusActive   FraudRuleStatus = "active"
	FraudRuleStatusInactive FraudRuleStatus = "inactive"
	FraudRuleStatusDraft    FraudRuleStatus = "draft"
)

func (r FraudRuleStatus) IsKnown() bool {
	switch r {
	case FraudRuleStatusActive, FraudRuleStatusInactive, FraudRuleStatusDraft:
		return true
	}
	return false
}

// Action to take when a fraud rule is triggered.
type FraudRuleAction struct {
	// Details or instructions for the action.
	Details interface{} `json:"details,required"`
	// Type of action to perform.
	Type FraudRuleActionType `json:"type,required"`
	// The team or department to notify for alerts/reviews.
	TargetTeam interface{}         `json:"targetTeam"`
	JSON       fraudRuleActionJSON `json:"-"`
}

// fraudRuleActionJSON contains the JSON metadata for the struct [FraudRuleAction]
type fraudRuleActionJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TargetTeam  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudRuleAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudRuleActionJSON) RawJSON() string {
	return r.raw
}

// Type of action to perform.
type FraudRuleActionType string

const (
	FraudRuleActionTypeBlock        FraudRuleActionType = "block"
	FraudRuleActionTypeAlert        FraudRuleActionType = "alert"
	FraudRuleActionTypeAutoReview   FraudRuleActionType = "auto_review"
	FraudRuleActionTypeManualReview FraudRuleActionType = "manual_review"
	FraudRuleActionTypeRequestMfa   FraudRuleActionType = "request_mfa"
)

func (r FraudRuleActionType) IsKnown() bool {
	switch r {
	case FraudRuleActionTypeBlock, FraudRuleActionTypeAlert, FraudRuleActionTypeAutoReview, FraudRuleActionTypeManualReview, FraudRuleActionTypeRequestMfa:
		return true
	}
	return false
}

// Action to take when a fraud rule is triggered.
type FraudRuleActionParam struct {
	// Details or instructions for the action.
	Details param.Field[interface{}] `json:"details,required"`
	// Type of action to perform.
	Type param.Field[FraudRuleActionType] `json:"type,required"`
	// The team or department to notify for alerts/reviews.
	TargetTeam param.Field[interface{}] `json:"targetTeam"`
}

func (r FraudRuleActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Criteria that define when a fraud rule should trigger.
type FraudRuleCriteria struct {
	// Number of days an account must be inactive for the rule to apply.
	AccountInactivityDays interface{} `json:"accountInactivityDays"`
	// List of ISO 2-letter country codes for transaction origin.
	CountryOfOrigin []interface{} `json:"countryOfOrigin,nullable"`
	// Minimum geographic distance (in km) from recent activity for anomaly.
	GeographicDistanceKm interface{} `json:"geographicDistanceKm"`
	// Number of days since last user login for anomaly detection.
	LastLoginDays interface{} `json:"lastLoginDays"`
	// If true, rule applies only if no prior travel notification was made.
	NoTravelNotification interface{} `json:"noTravelNotification"`
	// Minimum number of payments in a timeframe.
	PaymentCountMin interface{} `json:"paymentCountMin"`
	// List of risk levels for recipient countries.
	RecipientCountryRiskLevel []FraudRuleCriteriaRecipientCountryRiskLevel `json:"recipientCountryRiskLevel,nullable"`
	// If true, recipient must be a new payee.
	RecipientNew interface{} `json:"recipientNew"`
	// Timeframe in hours for payment count or other event aggregations.
	TimeframeHours interface{} `json:"timeframeHours"`
	// Minimum transaction amount to consider.
	TransactionAmountMin interface{} `json:"transactionAmountMin"`
	// Specific transaction type (e.g., debit, credit).
	TransactionType FraudRuleCriteriaTransactionType `json:"transactionType,nullable"`
	JSON            fraudRuleCriteriaJSON            `json:"-"`
}

// fraudRuleCriteriaJSON contains the JSON metadata for the struct
// [FraudRuleCriteria]
type fraudRuleCriteriaJSON struct {
	AccountInactivityDays     apijson.Field
	CountryOfOrigin           apijson.Field
	GeographicDistanceKm      apijson.Field
	LastLoginDays             apijson.Field
	NoTravelNotification      apijson.Field
	PaymentCountMin           apijson.Field
	RecipientCountryRiskLevel apijson.Field
	RecipientNew              apijson.Field
	TimeframeHours            apijson.Field
	TransactionAmountMin      apijson.Field
	TransactionType           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *FraudRuleCriteria) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudRuleCriteriaJSON) RawJSON() string {
	return r.raw
}

type FraudRuleCriteriaRecipientCountryRiskLevel string

const (
	FraudRuleCriteriaRecipientCountryRiskLevelLow      FraudRuleCriteriaRecipientCountryRiskLevel = "Low"
	FraudRuleCriteriaRecipientCountryRiskLevelMedium   FraudRuleCriteriaRecipientCountryRiskLevel = "Medium"
	FraudRuleCriteriaRecipientCountryRiskLevelHigh     FraudRuleCriteriaRecipientCountryRiskLevel = "High"
	FraudRuleCriteriaRecipientCountryRiskLevelVeryHigh FraudRuleCriteriaRecipientCountryRiskLevel = "Very High"
)

func (r FraudRuleCriteriaRecipientCountryRiskLevel) IsKnown() bool {
	switch r {
	case FraudRuleCriteriaRecipientCountryRiskLevelLow, FraudRuleCriteriaRecipientCountryRiskLevelMedium, FraudRuleCriteriaRecipientCountryRiskLevelHigh, FraudRuleCriteriaRecipientCountryRiskLevelVeryHigh:
		return true
	}
	return false
}

// Specific transaction type (e.g., debit, credit).
type FraudRuleCriteriaTransactionType string

const (
	FraudRuleCriteriaTransactionTypeDebit  FraudRuleCriteriaTransactionType = "debit"
	FraudRuleCriteriaTransactionTypeCredit FraudRuleCriteriaTransactionType = "credit"
)

func (r FraudRuleCriteriaTransactionType) IsKnown() bool {
	switch r {
	case FraudRuleCriteriaTransactionTypeDebit, FraudRuleCriteriaTransactionTypeCredit:
		return true
	}
	return false
}

// Criteria that define when a fraud rule should trigger.
type FraudRuleCriteriaParam struct {
	// Number of days an account must be inactive for the rule to apply.
	AccountInactivityDays param.Field[interface{}] `json:"accountInactivityDays"`
	// List of ISO 2-letter country codes for transaction origin.
	CountryOfOrigin param.Field[[]interface{}] `json:"countryOfOrigin"`
	// Minimum geographic distance (in km) from recent activity for anomaly.
	GeographicDistanceKm param.Field[interface{}] `json:"geographicDistanceKm"`
	// Number of days since last user login for anomaly detection.
	LastLoginDays param.Field[interface{}] `json:"lastLoginDays"`
	// If true, rule applies only if no prior travel notification was made.
	NoTravelNotification param.Field[interface{}] `json:"noTravelNotification"`
	// Minimum number of payments in a timeframe.
	PaymentCountMin param.Field[interface{}] `json:"paymentCountMin"`
	// List of risk levels for recipient countries.
	RecipientCountryRiskLevel param.Field[[]FraudRuleCriteriaRecipientCountryRiskLevel] `json:"recipientCountryRiskLevel"`
	// If true, recipient must be a new payee.
	RecipientNew param.Field[interface{}] `json:"recipientNew"`
	// Timeframe in hours for payment count or other event aggregations.
	TimeframeHours param.Field[interface{}] `json:"timeframeHours"`
	// Minimum transaction amount to consider.
	TransactionAmountMin param.Field[interface{}] `json:"transactionAmountMin"`
	// Specific transaction type (e.g., debit, credit).
	TransactionType param.Field[FraudRuleCriteriaTransactionType] `json:"transactionType"`
}

func (r FraudRuleCriteriaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateRiskFraudRuleListResponse struct {
	Data []FraudRule                            `json:"data"`
	JSON corporateRiskFraudRuleListResponseJSON `json:"-"`
	PaginatedList
}

// corporateRiskFraudRuleListResponseJSON contains the JSON metadata for the struct
// [CorporateRiskFraudRuleListResponse]
type corporateRiskFraudRuleListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateRiskFraudRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskFraudRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskFraudRuleNewParams struct {
	// Action to take when a fraud rule is triggered.
	Action param.Field[FraudRuleActionParam] `json:"action,required"`
	// Criteria that define when a fraud rule should trigger.
	Criteria param.Field[FraudRuleCriteriaParam] `json:"criteria,required"`
	// Detailed description of what the rule detects.
	Description param.Field[interface{}] `json:"description,required"`
	// Name of the new fraud rule.
	Name param.Field[interface{}] `json:"name,required"`
	// Severity level when this rule is triggered.
	Severity param.Field[CorporateRiskFraudRuleNewParamsSeverity] `json:"severity,required"`
	// Initial status of the rule.
	Status param.Field[CorporateRiskFraudRuleNewParamsStatus] `json:"status,required"`
}

func (r CorporateRiskFraudRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Severity level when this rule is triggered.
type CorporateRiskFraudRuleNewParamsSeverity string

const (
	CorporateRiskFraudRuleNewParamsSeverityLow      CorporateRiskFraudRuleNewParamsSeverity = "Low"
	CorporateRiskFraudRuleNewParamsSeverityMedium   CorporateRiskFraudRuleNewParamsSeverity = "Medium"
	CorporateRiskFraudRuleNewParamsSeverityHigh     CorporateRiskFraudRuleNewParamsSeverity = "High"
	CorporateRiskFraudRuleNewParamsSeverityCritical CorporateRiskFraudRuleNewParamsSeverity = "Critical"
)

func (r CorporateRiskFraudRuleNewParamsSeverity) IsKnown() bool {
	switch r {
	case CorporateRiskFraudRuleNewParamsSeverityLow, CorporateRiskFraudRuleNewParamsSeverityMedium, CorporateRiskFraudRuleNewParamsSeverityHigh, CorporateRiskFraudRuleNewParamsSeverityCritical:
		return true
	}
	return false
}

// Initial status of the rule.
type CorporateRiskFraudRuleNewParamsStatus string

const (
	CorporateRiskFraudRuleNewParamsStatusActive   CorporateRiskFraudRuleNewParamsStatus = "active"
	CorporateRiskFraudRuleNewParamsStatusInactive CorporateRiskFraudRuleNewParamsStatus = "inactive"
	CorporateRiskFraudRuleNewParamsStatusDraft    CorporateRiskFraudRuleNewParamsStatus = "draft"
)

func (r CorporateRiskFraudRuleNewParamsStatus) IsKnown() bool {
	switch r {
	case CorporateRiskFraudRuleNewParamsStatusActive, CorporateRiskFraudRuleNewParamsStatusInactive, CorporateRiskFraudRuleNewParamsStatusDraft:
		return true
	}
	return false
}

type CorporateRiskFraudRuleUpdateParams struct {
	// Action to take when a fraud rule is triggered.
	Action param.Field[FraudRuleActionParam] `json:"action"`
	// Criteria that define when a fraud rule should trigger.
	Criteria param.Field[FraudRuleCriteriaParam] `json:"criteria"`
	// Updated description of what the rule detects.
	Description param.Field[interface{}] `json:"description"`
	// Updated name of the fraud rule.
	Name param.Field[interface{}] `json:"name"`
	// Updated severity level.
	Severity param.Field[CorporateRiskFraudRuleUpdateParamsSeverity] `json:"severity"`
	// Updated status of the rule.
	Status param.Field[CorporateRiskFraudRuleUpdateParamsStatus] `json:"status"`
}

func (r CorporateRiskFraudRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated severity level.
type CorporateRiskFraudRuleUpdateParamsSeverity string

const (
	CorporateRiskFraudRuleUpdateParamsSeverityLow      CorporateRiskFraudRuleUpdateParamsSeverity = "Low"
	CorporateRiskFraudRuleUpdateParamsSeverityMedium   CorporateRiskFraudRuleUpdateParamsSeverity = "Medium"
	CorporateRiskFraudRuleUpdateParamsSeverityHigh     CorporateRiskFraudRuleUpdateParamsSeverity = "High"
	CorporateRiskFraudRuleUpdateParamsSeverityCritical CorporateRiskFraudRuleUpdateParamsSeverity = "Critical"
)

func (r CorporateRiskFraudRuleUpdateParamsSeverity) IsKnown() bool {
	switch r {
	case CorporateRiskFraudRuleUpdateParamsSeverityLow, CorporateRiskFraudRuleUpdateParamsSeverityMedium, CorporateRiskFraudRuleUpdateParamsSeverityHigh, CorporateRiskFraudRuleUpdateParamsSeverityCritical:
		return true
	}
	return false
}

// Updated status of the rule.
type CorporateRiskFraudRuleUpdateParamsStatus string

const (
	CorporateRiskFraudRuleUpdateParamsStatusActive   CorporateRiskFraudRuleUpdateParamsStatus = "active"
	CorporateRiskFraudRuleUpdateParamsStatusInactive CorporateRiskFraudRuleUpdateParamsStatus = "inactive"
	CorporateRiskFraudRuleUpdateParamsStatusDraft    CorporateRiskFraudRuleUpdateParamsStatus = "draft"
)

func (r CorporateRiskFraudRuleUpdateParamsStatus) IsKnown() bool {
	switch r {
	case CorporateRiskFraudRuleUpdateParamsStatusActive, CorporateRiskFraudRuleUpdateParamsStatusInactive, CorporateRiskFraudRuleUpdateParamsStatusDraft:
		return true
	}
	return false
}

type CorporateRiskFraudRuleListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [CorporateRiskFraudRuleListParams]'s query parameters as
// `url.Values`.
func (r CorporateRiskFraudRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
