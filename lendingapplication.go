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

// LendingApplicationService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendingApplicationService] method instead.
type LendingApplicationService struct {
	Options []option.RequestOption
}

// NewLendingApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLendingApplicationService(opts ...option.RequestOption) (r *LendingApplicationService) {
	r = &LendingApplicationService{}
	r.Options = opts
	return
}

// Retrieves the current status and detailed information for a submitted loan
// application, including AI underwriting outcomes, approved terms, and next steps.
func (r *LendingApplicationService) Get(ctx context.Context, applicationID interface{}, opts ...option.RequestOption) (res *LoanApplicationStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("lending/applications/%v", applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits a new loan application, which is instantly processed and underwritten by
// our Quantum AI, providing rapid decisions and personalized loan offers based on
// real-time financial health data.
func (r *LendingApplicationService) Submit(ctx context.Context, body LendingApplicationSubmitParams, opts ...option.RequestOption) (res *LoanApplicationStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lending/applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LoanApplicationStatus struct {
	// Timestamp when the application was submitted.
	ApplicationDate interface{} `json:"applicationDate,required"`
	// Unique identifier for the loan application.
	ApplicationID interface{} `json:"applicationId,required"`
	// The amount originally requested in the application.
	LoanAmountRequested interface{} `json:"loanAmountRequested,required"`
	// The purpose of the loan.
	LoanPurpose LoanApplicationStatusLoanPurpose `json:"loanPurpose,required"`
	// Guidance on the next actions for the user.
	NextSteps interface{} `json:"nextSteps,required"`
	// Current status of the loan application.
	Status               LoanApplicationStatusStatus               `json:"status,required"`
	AIUnderwritingResult LoanApplicationStatusAIUnderwritingResult `json:"aiUnderwritingResult"`
	OfferDetails         LoanOffer                                 `json:"offerDetails"`
	JSON                 loanApplicationStatusJSON                 `json:"-"`
}

// loanApplicationStatusJSON contains the JSON metadata for the struct
// [LoanApplicationStatus]
type loanApplicationStatusJSON struct {
	ApplicationDate      apijson.Field
	ApplicationID        apijson.Field
	LoanAmountRequested  apijson.Field
	LoanPurpose          apijson.Field
	NextSteps            apijson.Field
	Status               apijson.Field
	AIUnderwritingResult apijson.Field
	OfferDetails         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoanApplicationStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanApplicationStatusJSON) RawJSON() string {
	return r.raw
}

// The purpose of the loan.
type LoanApplicationStatusLoanPurpose string

const (
	LoanApplicationStatusLoanPurposeHomeImprovement   LoanApplicationStatusLoanPurpose = "home_improvement"
	LoanApplicationStatusLoanPurposeDebtConsolidation LoanApplicationStatusLoanPurpose = "debt_consolidation"
	LoanApplicationStatusLoanPurposeMedicalExpense    LoanApplicationStatusLoanPurpose = "medical_expense"
	LoanApplicationStatusLoanPurposeEducation         LoanApplicationStatusLoanPurpose = "education"
	LoanApplicationStatusLoanPurposeAutoPurchase      LoanApplicationStatusLoanPurpose = "auto_purchase"
	LoanApplicationStatusLoanPurposeOther             LoanApplicationStatusLoanPurpose = "other"
)

func (r LoanApplicationStatusLoanPurpose) IsKnown() bool {
	switch r {
	case LoanApplicationStatusLoanPurposeHomeImprovement, LoanApplicationStatusLoanPurposeDebtConsolidation, LoanApplicationStatusLoanPurposeMedicalExpense, LoanApplicationStatusLoanPurposeEducation, LoanApplicationStatusLoanPurposeAutoPurchase, LoanApplicationStatusLoanPurposeOther:
		return true
	}
	return false
}

// Current status of the loan application.
type LoanApplicationStatusStatus string

const (
	LoanApplicationStatusStatusSubmitted         LoanApplicationStatusStatus = "submitted"
	LoanApplicationStatusStatusUnderwriting      LoanApplicationStatusStatus = "underwriting"
	LoanApplicationStatusStatusApproved          LoanApplicationStatusStatus = "approved"
	LoanApplicationStatusStatusDeclined          LoanApplicationStatusStatus = "declined"
	LoanApplicationStatusStatusPendingAcceptance LoanApplicationStatusStatus = "pending_acceptance"
	LoanApplicationStatusStatusFunded            LoanApplicationStatusStatus = "funded"
	LoanApplicationStatusStatusCancelled         LoanApplicationStatusStatus = "cancelled"
)

func (r LoanApplicationStatusStatus) IsKnown() bool {
	switch r {
	case LoanApplicationStatusStatusSubmitted, LoanApplicationStatusStatusUnderwriting, LoanApplicationStatusStatusApproved, LoanApplicationStatusStatusDeclined, LoanApplicationStatusStatusPendingAcceptance, LoanApplicationStatusStatusFunded, LoanApplicationStatusStatusCancelled:
		return true
	}
	return false
}

type LoanApplicationStatusAIUnderwritingResult struct {
	// AI's confidence in its underwriting decision (0-1).
	AIConfidence interface{} `json:"aiConfidence,required"`
	// The AI's underwriting decision.
	Decision LoanApplicationStatusAIUnderwritingResultDecision `json:"decision,required"`
	// Reasoning for the AI's decision.
	Reason interface{} `json:"reason,required"`
	// The maximum amount the AI is willing to approve.
	MaxApprovedAmount interface{} `json:"maxApprovedAmount"`
	// The interest rate recommended by the AI.
	RecommendedInterestRate interface{}                                   `json:"recommendedInterestRate"`
	JSON                    loanApplicationStatusAIUnderwritingResultJSON `json:"-"`
}

// loanApplicationStatusAIUnderwritingResultJSON contains the JSON metadata for the
// struct [LoanApplicationStatusAIUnderwritingResult]
type loanApplicationStatusAIUnderwritingResultJSON struct {
	AIConfidence            apijson.Field
	Decision                apijson.Field
	Reason                  apijson.Field
	MaxApprovedAmount       apijson.Field
	RecommendedInterestRate apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LoanApplicationStatusAIUnderwritingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanApplicationStatusAIUnderwritingResultJSON) RawJSON() string {
	return r.raw
}

// The AI's underwriting decision.
type LoanApplicationStatusAIUnderwritingResultDecision string

const (
	LoanApplicationStatusAIUnderwritingResultDecisionApproved        LoanApplicationStatusAIUnderwritingResultDecision = "approved"
	LoanApplicationStatusAIUnderwritingResultDecisionDeclined        LoanApplicationStatusAIUnderwritingResultDecision = "declined"
	LoanApplicationStatusAIUnderwritingResultDecisionReferredToHuman LoanApplicationStatusAIUnderwritingResultDecision = "referred_to_human"
)

func (r LoanApplicationStatusAIUnderwritingResultDecision) IsKnown() bool {
	switch r {
	case LoanApplicationStatusAIUnderwritingResultDecisionApproved, LoanApplicationStatusAIUnderwritingResultDecisionDeclined, LoanApplicationStatusAIUnderwritingResultDecisionReferredToHuman:
		return true
	}
	return false
}

type LendingApplicationSubmitParams struct {
	// The desired loan amount.
	LoanAmount param.Field[interface{}] `json:"loanAmount,required"`
	// The purpose of the loan.
	LoanPurpose param.Field[LendingApplicationSubmitParamsLoanPurpose] `json:"loanPurpose,required"`
	// The desired repayment term in months.
	RepaymentTermMonths param.Field[interface{}] `json:"repaymentTermMonths,required"`
	// Optional notes or details for the application.
	AdditionalNotes param.Field[interface{}] `json:"additionalNotes"`
	// Optional: Details of a co-applicant for the loan.
	CoApplicant param.Field[LendingApplicationSubmitParamsCoApplicant] `json:"coApplicant"`
}

func (r LendingApplicationSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The purpose of the loan.
type LendingApplicationSubmitParamsLoanPurpose string

const (
	LendingApplicationSubmitParamsLoanPurposeHomeImprovement   LendingApplicationSubmitParamsLoanPurpose = "home_improvement"
	LendingApplicationSubmitParamsLoanPurposeDebtConsolidation LendingApplicationSubmitParamsLoanPurpose = "debt_consolidation"
	LendingApplicationSubmitParamsLoanPurposeMedicalExpense    LendingApplicationSubmitParamsLoanPurpose = "medical_expense"
	LendingApplicationSubmitParamsLoanPurposeEducation         LendingApplicationSubmitParamsLoanPurpose = "education"
	LendingApplicationSubmitParamsLoanPurposeAutoPurchase      LendingApplicationSubmitParamsLoanPurpose = "auto_purchase"
	LendingApplicationSubmitParamsLoanPurposeOther             LendingApplicationSubmitParamsLoanPurpose = "other"
)

func (r LendingApplicationSubmitParamsLoanPurpose) IsKnown() bool {
	switch r {
	case LendingApplicationSubmitParamsLoanPurposeHomeImprovement, LendingApplicationSubmitParamsLoanPurposeDebtConsolidation, LendingApplicationSubmitParamsLoanPurposeMedicalExpense, LendingApplicationSubmitParamsLoanPurposeEducation, LendingApplicationSubmitParamsLoanPurposeAutoPurchase, LendingApplicationSubmitParamsLoanPurposeOther:
		return true
	}
	return false
}

// Optional: Details of a co-applicant for the loan.
type LendingApplicationSubmitParamsCoApplicant struct {
	Email  param.Field[interface{}] `json:"email"`
	Income param.Field[interface{}] `json:"income"`
	Name   param.Field[interface{}] `json:"name"`
}

func (r LendingApplicationSubmitParamsCoApplicant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
