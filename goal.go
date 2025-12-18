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

// GoalService contains methods and other services that help with interacting with
// the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoalService] method instead.
type GoalService struct {
	Options []option.RequestOption
}

// NewGoalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGoalService(opts ...option.RequestOption) (r *GoalService) {
	r = &GoalService{}
	r.Options = opts
	return
}

// Creates a new long-term financial goal, with optional AI plan generation.
func (r *GoalService) New(ctx context.Context, body GoalNewParams, opts ...option.RequestOption) (res *FinancialGoal, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "goals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information for a specific financial goal, including current
// progress, AI strategic plan, and related insights.
func (r *GoalService) Get(ctx context.Context, goalID interface{}, opts ...option.RequestOption) (res *FinancialGoal, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("goals/%v", goalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the parameters of an existing financial goal, such as target amount,
// date, or contributing accounts. This may trigger an AI plan recalculation.
func (r *GoalService) Update(ctx context.Context, goalID interface{}, body GoalUpdateParams, opts ...option.RequestOption) (res *FinancialGoal, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("goals/%v", goalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of all financial goals defined by the user, including their
// progress and associated AI plans.
func (r *GoalService) List(ctx context.Context, query GoalListParams, opts ...option.RequestOption) (res *GoalListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "goals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specific financial goal from the user's profile.
func (r *GoalService) Delete(ctx context.Context, goalID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("goals/%v", goalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FinancialGoal struct {
	// Unique identifier for the financial goal.
	ID interface{} `json:"id,required"`
	// The current amount saved or invested towards the goal.
	CurrentAmount interface{} `json:"currentAmount,required"`
	// Timestamp when the goal's status or progress was last updated.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Name of the financial goal.
	Name interface{} `json:"name,required"`
	// Percentage completion of the goal.
	ProgressPercentage interface{} `json:"progressPercentage,required"`
	// Current status of the goal's progress.
	Status FinancialGoalStatus `json:"status,required"`
	// The target monetary amount for the goal.
	TargetAmount interface{} `json:"targetAmount,required"`
	// The target completion date for the goal.
	TargetDate interface{} `json:"targetDate,required"`
	// Type of financial goal.
	Type FinancialGoalType `json:"type,required"`
	// AI-driven insights and recommendations related to this goal.
	AIInsights []AIInsight `json:"aiInsights,nullable"`
	// AI-generated strategic plan for achieving the goal.
	AIStrategicPlan FinancialGoalAIStrategicPlan `json:"aiStrategicPlan"`
	// List of account IDs contributing to this goal.
	ContributingAccounts []interface{} `json:"contributingAccounts,nullable"`
	// Recommended or chosen risk tolerance for investments related to this goal.
	RiskTolerance FinancialGoalRiskTolerance `json:"riskTolerance,nullable"`
	JSON          financialGoalJSON          `json:"-"`
}

// financialGoalJSON contains the JSON metadata for the struct [FinancialGoal]
type financialGoalJSON struct {
	ID                   apijson.Field
	CurrentAmount        apijson.Field
	LastUpdated          apijson.Field
	Name                 apijson.Field
	ProgressPercentage   apijson.Field
	Status               apijson.Field
	TargetAmount         apijson.Field
	TargetDate           apijson.Field
	Type                 apijson.Field
	AIInsights           apijson.Field
	AIStrategicPlan      apijson.Field
	ContributingAccounts apijson.Field
	RiskTolerance        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FinancialGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialGoalJSON) RawJSON() string {
	return r.raw
}

// Current status of the goal's progress.
type FinancialGoalStatus string

const (
	FinancialGoalStatusOnTrack         FinancialGoalStatus = "on_track"
	FinancialGoalStatusBehindSchedule  FinancialGoalStatus = "behind_schedule"
	FinancialGoalStatusAheadOfSchedule FinancialGoalStatus = "ahead_of_schedule"
	FinancialGoalStatusCompleted       FinancialGoalStatus = "completed"
	FinancialGoalStatusPaused          FinancialGoalStatus = "paused"
	FinancialGoalStatusCancelled       FinancialGoalStatus = "cancelled"
)

func (r FinancialGoalStatus) IsKnown() bool {
	switch r {
	case FinancialGoalStatusOnTrack, FinancialGoalStatusBehindSchedule, FinancialGoalStatusAheadOfSchedule, FinancialGoalStatusCompleted, FinancialGoalStatusPaused, FinancialGoalStatusCancelled:
		return true
	}
	return false
}

// Type of financial goal.
type FinancialGoalType string

const (
	FinancialGoalTypeRetirement    FinancialGoalType = "retirement"
	FinancialGoalTypeHomePurchase  FinancialGoalType = "home_purchase"
	FinancialGoalTypeEducation     FinancialGoalType = "education"
	FinancialGoalTypeLargePurchase FinancialGoalType = "large_purchase"
	FinancialGoalTypeDebtReduction FinancialGoalType = "debt_reduction"
	FinancialGoalTypeOther         FinancialGoalType = "other"
)

func (r FinancialGoalType) IsKnown() bool {
	switch r {
	case FinancialGoalTypeRetirement, FinancialGoalTypeHomePurchase, FinancialGoalTypeEducation, FinancialGoalTypeLargePurchase, FinancialGoalTypeDebtReduction, FinancialGoalTypeOther:
		return true
	}
	return false
}

// AI-generated strategic plan for achieving the goal.
type FinancialGoalAIStrategicPlan struct {
	PlanID  interface{}                        `json:"planId"`
	Steps   []FinancialGoalAIStrategicPlanStep `json:"steps"`
	Summary interface{}                        `json:"summary"`
	JSON    financialGoalAIStrategicPlanJSON   `json:"-"`
}

// financialGoalAIStrategicPlanJSON contains the JSON metadata for the struct
// [FinancialGoalAIStrategicPlan]
type financialGoalAIStrategicPlanJSON struct {
	PlanID      apijson.Field
	Steps       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialGoalAIStrategicPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialGoalAIStrategicPlanJSON) RawJSON() string {
	return r.raw
}

type FinancialGoalAIStrategicPlanStep struct {
	Description interface{}                             `json:"description"`
	Status      FinancialGoalAIStrategicPlanStepsStatus `json:"status"`
	Title       interface{}                             `json:"title"`
	JSON        financialGoalAIStrategicPlanStepJSON    `json:"-"`
}

// financialGoalAIStrategicPlanStepJSON contains the JSON metadata for the struct
// [FinancialGoalAIStrategicPlanStep]
type financialGoalAIStrategicPlanStepJSON struct {
	Description apijson.Field
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialGoalAIStrategicPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialGoalAIStrategicPlanStepJSON) RawJSON() string {
	return r.raw
}

type FinancialGoalAIStrategicPlanStepsStatus string

const (
	FinancialGoalAIStrategicPlanStepsStatusPending    FinancialGoalAIStrategicPlanStepsStatus = "pending"
	FinancialGoalAIStrategicPlanStepsStatusInProgress FinancialGoalAIStrategicPlanStepsStatus = "in_progress"
	FinancialGoalAIStrategicPlanStepsStatusCompleted  FinancialGoalAIStrategicPlanStepsStatus = "completed"
)

func (r FinancialGoalAIStrategicPlanStepsStatus) IsKnown() bool {
	switch r {
	case FinancialGoalAIStrategicPlanStepsStatusPending, FinancialGoalAIStrategicPlanStepsStatusInProgress, FinancialGoalAIStrategicPlanStepsStatusCompleted:
		return true
	}
	return false
}

// Recommended or chosen risk tolerance for investments related to this goal.
type FinancialGoalRiskTolerance string

const (
	FinancialGoalRiskToleranceConservative FinancialGoalRiskTolerance = "conservative"
	FinancialGoalRiskToleranceModerate     FinancialGoalRiskTolerance = "moderate"
	FinancialGoalRiskToleranceAggressive   FinancialGoalRiskTolerance = "aggressive"
)

func (r FinancialGoalRiskTolerance) IsKnown() bool {
	switch r {
	case FinancialGoalRiskToleranceConservative, FinancialGoalRiskToleranceModerate, FinancialGoalRiskToleranceAggressive:
		return true
	}
	return false
}

type GoalListResponse struct {
	Data []FinancialGoal      `json:"data"`
	JSON goalListResponseJSON `json:"-"`
	PaginatedList
}

// goalListResponseJSON contains the JSON metadata for the struct
// [GoalListResponse]
type goalListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r goalListResponseJSON) RawJSON() string {
	return r.raw
}

type GoalNewParams struct {
	// Name of the new financial goal.
	Name param.Field[interface{}] `json:"name,required"`
	// The target monetary amount for the goal.
	TargetAmount param.Field[interface{}] `json:"targetAmount,required"`
	// The target completion date for the goal.
	TargetDate param.Field[interface{}] `json:"targetDate,required"`
	// Type of financial goal.
	Type param.Field[GoalNewParamsType] `json:"type,required"`
	// Optional: List of account IDs initially contributing to this goal.
	ContributingAccounts param.Field[[]interface{}] `json:"contributingAccounts"`
	// If true, AI will automatically generate a strategic plan for the goal.
	GenerateAIPlan param.Field[interface{}] `json:"generateAIPlan"`
	// Optional: Initial amount to contribute to the goal.
	InitialContribution param.Field[interface{}] `json:"initialContribution"`
	// Desired risk tolerance for investments related to this goal.
	RiskTolerance param.Field[GoalNewParamsRiskTolerance] `json:"riskTolerance"`
}

func (r GoalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of financial goal.
type GoalNewParamsType string

const (
	GoalNewParamsTypeRetirement    GoalNewParamsType = "retirement"
	GoalNewParamsTypeHomePurchase  GoalNewParamsType = "home_purchase"
	GoalNewParamsTypeEducation     GoalNewParamsType = "education"
	GoalNewParamsTypeLargePurchase GoalNewParamsType = "large_purchase"
	GoalNewParamsTypeDebtReduction GoalNewParamsType = "debt_reduction"
	GoalNewParamsTypeOther         GoalNewParamsType = "other"
)

func (r GoalNewParamsType) IsKnown() bool {
	switch r {
	case GoalNewParamsTypeRetirement, GoalNewParamsTypeHomePurchase, GoalNewParamsTypeEducation, GoalNewParamsTypeLargePurchase, GoalNewParamsTypeDebtReduction, GoalNewParamsTypeOther:
		return true
	}
	return false
}

// Desired risk tolerance for investments related to this goal.
type GoalNewParamsRiskTolerance string

const (
	GoalNewParamsRiskToleranceConservative GoalNewParamsRiskTolerance = "conservative"
	GoalNewParamsRiskToleranceModerate     GoalNewParamsRiskTolerance = "moderate"
	GoalNewParamsRiskToleranceAggressive   GoalNewParamsRiskTolerance = "aggressive"
)

func (r GoalNewParamsRiskTolerance) IsKnown() bool {
	switch r {
	case GoalNewParamsRiskToleranceConservative, GoalNewParamsRiskToleranceModerate, GoalNewParamsRiskToleranceAggressive:
		return true
	}
	return false
}

type GoalUpdateParams struct {
	// Updated list of account IDs contributing to this goal.
	ContributingAccounts param.Field[[]interface{}] `json:"contributingAccounts"`
	// If true, AI will recalculate and update the strategic plan for the goal.
	GenerateAIPlan param.Field[interface{}] `json:"generateAIPlan"`
	// Updated name of the financial goal.
	Name param.Field[interface{}] `json:"name"`
	// Updated risk tolerance for investments related to this goal.
	RiskTolerance param.Field[GoalUpdateParamsRiskTolerance] `json:"riskTolerance"`
	// Updated status of the goal's progress.
	Status param.Field[GoalUpdateParamsStatus] `json:"status"`
	// The updated target monetary amount for the goal.
	TargetAmount param.Field[interface{}] `json:"targetAmount"`
	// The updated target completion date for the goal.
	TargetDate param.Field[interface{}] `json:"targetDate"`
}

func (r GoalUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated risk tolerance for investments related to this goal.
type GoalUpdateParamsRiskTolerance string

const (
	GoalUpdateParamsRiskToleranceConservative GoalUpdateParamsRiskTolerance = "conservative"
	GoalUpdateParamsRiskToleranceModerate     GoalUpdateParamsRiskTolerance = "moderate"
	GoalUpdateParamsRiskToleranceAggressive   GoalUpdateParamsRiskTolerance = "aggressive"
)

func (r GoalUpdateParamsRiskTolerance) IsKnown() bool {
	switch r {
	case GoalUpdateParamsRiskToleranceConservative, GoalUpdateParamsRiskToleranceModerate, GoalUpdateParamsRiskToleranceAggressive:
		return true
	}
	return false
}

// Updated status of the goal's progress.
type GoalUpdateParamsStatus string

const (
	GoalUpdateParamsStatusOnTrack         GoalUpdateParamsStatus = "on_track"
	GoalUpdateParamsStatusBehindSchedule  GoalUpdateParamsStatus = "behind_schedule"
	GoalUpdateParamsStatusAheadOfSchedule GoalUpdateParamsStatus = "ahead_of_schedule"
	GoalUpdateParamsStatusCompleted       GoalUpdateParamsStatus = "completed"
	GoalUpdateParamsStatusPaused          GoalUpdateParamsStatus = "paused"
	GoalUpdateParamsStatusCancelled       GoalUpdateParamsStatus = "cancelled"
)

func (r GoalUpdateParamsStatus) IsKnown() bool {
	switch r {
	case GoalUpdateParamsStatusOnTrack, GoalUpdateParamsStatusBehindSchedule, GoalUpdateParamsStatusAheadOfSchedule, GoalUpdateParamsStatusCompleted, GoalUpdateParamsStatusPaused, GoalUpdateParamsStatusCancelled:
		return true
	}
	return false
}

type GoalListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [GoalListParams]'s query parameters as `url.Values`.
func (r GoalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
