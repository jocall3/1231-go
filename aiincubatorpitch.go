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

// AIIncubatorPitchService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorPitchService] method instead.
type AIIncubatorPitchService struct {
	Options []option.RequestOption
}

// NewAIIncubatorPitchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIIncubatorPitchService(opts ...option.RequestOption) (r *AIIncubatorPitchService) {
	r = &AIIncubatorPitchService{}
	r.Options = opts
	return
}

// Retrieves the granular AI-driven analysis, strategic feedback, market validation
// results, and any outstanding questions from Quantum Weaver for a specific
// business pitch.
func (r *AIIncubatorPitchService) GetDetails(ctx context.Context, pitchID interface{}, opts ...option.RequestOption) (res *AIIncubatorPitchGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/incubator/pitch/%v/details", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits a detailed business plan to the Quantum Weaver AI for rigorous analysis,
// market validation, and seed funding consideration. This initiates the AI-driven
// incubation journey, aiming to transform innovative ideas into commercially
// successful ventures.
func (r *AIIncubatorPitchService) Submit(ctx context.Context, body AIIncubatorPitchSubmitParams, opts ...option.RequestOption) (res *QuantumWeaverState, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows the entrepreneur to respond to specific questions or provide additional
// details requested by Quantum Weaver, moving the pitch forward in the incubation
// process.
func (r *AIIncubatorPitchService) SubmitFeedback(ctx context.Context, pitchID interface{}, body AIIncubatorPitchSubmitFeedbackParams, opts ...option.RequestOption) (res *QuantumWeaverState, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/incubator/pitch/%v/feedback", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type QuantumWeaverState struct {
	// Timestamp of the last status update.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Guidance on the next actions for the user.
	NextSteps interface{} `json:"nextSteps,required"`
	// Unique identifier for the business pitch.
	PitchID interface{} `json:"pitchId,required"`
	// Current stage of the business pitch in the incubation process.
	Stage QuantumWeaverStateStage `json:"stage,required"`
	// A human-readable status message.
	StatusMessage interface{} `json:"statusMessage,required"`
	// AI's estimated funding offer, if the pitch progresses.
	EstimatedFundingOffer interface{} `json:"estimatedFundingOffer"`
	// A summary of AI-generated feedback, if applicable.
	FeedbackSummary interface{} `json:"feedbackSummary"`
	// List of questions from Quantum Weaver requiring the user's input.
	Questions []QuantumWeaverStateQuestion `json:"questions,nullable"`
	JSON      quantumWeaverStateJSON       `json:"-"`
}

// quantumWeaverStateJSON contains the JSON metadata for the struct
// [QuantumWeaverState]
type quantumWeaverStateJSON struct {
	LastUpdated           apijson.Field
	NextSteps             apijson.Field
	PitchID               apijson.Field
	Stage                 apijson.Field
	StatusMessage         apijson.Field
	EstimatedFundingOffer apijson.Field
	FeedbackSummary       apijson.Field
	Questions             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *QuantumWeaverState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quantumWeaverStateJSON) RawJSON() string {
	return r.raw
}

// Current stage of the business pitch in the incubation process.
type QuantumWeaverStateStage string

const (
	QuantumWeaverStateStageSubmitted          QuantumWeaverStateStage = "submitted"
	QuantumWeaverStateStageInitialReview      QuantumWeaverStateStage = "initial_review"
	QuantumWeaverStateStageAIAnalysis         QuantumWeaverStateStage = "ai_analysis"
	QuantumWeaverStateStageFeedbackRequired   QuantumWeaverStateStage = "feedback_required"
	QuantumWeaverStateStageTestPhase          QuantumWeaverStateStage = "test_phase"
	QuantumWeaverStateStageFinalReview        QuantumWeaverStateStage = "final_review"
	QuantumWeaverStateStageApprovedForFunding QuantumWeaverStateStage = "approved_for_funding"
	QuantumWeaverStateStageRejected           QuantumWeaverStateStage = "rejected"
	QuantumWeaverStateStageIncubatedGraduated QuantumWeaverStateStage = "incubated_graduated"
)

func (r QuantumWeaverStateStage) IsKnown() bool {
	switch r {
	case QuantumWeaverStateStageSubmitted, QuantumWeaverStateStageInitialReview, QuantumWeaverStateStageAIAnalysis, QuantumWeaverStateStageFeedbackRequired, QuantumWeaverStateStageTestPhase, QuantumWeaverStateStageFinalReview, QuantumWeaverStateStageApprovedForFunding, QuantumWeaverStateStageRejected, QuantumWeaverStateStageIncubatedGraduated:
		return true
	}
	return false
}

type QuantumWeaverStateQuestion struct {
	ID         interface{}                    `json:"id"`
	Category   interface{}                    `json:"category"`
	IsRequired interface{}                    `json:"isRequired"`
	Question   interface{}                    `json:"question"`
	JSON       quantumWeaverStateQuestionJSON `json:"-"`
}

// quantumWeaverStateQuestionJSON contains the JSON metadata for the struct
// [QuantumWeaverStateQuestion]
type quantumWeaverStateQuestionJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	IsRequired  apijson.Field
	Question    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuantumWeaverStateQuestion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quantumWeaverStateQuestionJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponse struct {
	// AI-generated coaching plan for the entrepreneur.
	AICoachingPlan AIIncubatorPitchGetDetailsResponseAICoachingPlan `json:"aiCoachingPlan,nullable"`
	// AI's detailed financial model analysis.
	AIFinancialModel AIIncubatorPitchGetDetailsResponseAIFinancialModel `json:"aiFinancialModel,nullable"`
	// AI's detailed market analysis.
	AIMarketAnalysis AIIncubatorPitchGetDetailsResponseAIMarketAnalysis `json:"aiMarketAnalysis,nullable"`
	// AI's assessment of risks associated with the venture.
	AIRiskAssessment AIIncubatorPitchGetDetailsResponseAIRiskAssessment `json:"aiRiskAssessment,nullable"`
	// AI's score for how well the pitch matches potential investors in the network
	// (0-1).
	InvestorMatchScore interface{}                            `json:"investorMatchScore"`
	JSON               aiIncubatorPitchGetDetailsResponseJSON `json:"-"`
	QuantumWeaverState
}

// aiIncubatorPitchGetDetailsResponseJSON contains the JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponse]
type aiIncubatorPitchGetDetailsResponseJSON struct {
	AICoachingPlan     apijson.Field
	AIFinancialModel   apijson.Field
	AIMarketAnalysis   apijson.Field
	AIRiskAssessment   apijson.Field
	InvestorMatchScore apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// AI-generated coaching plan for the entrepreneur.
type AIIncubatorPitchGetDetailsResponseAICoachingPlan struct {
	Steps   []AIIncubatorPitchGetDetailsResponseAICoachingPlanStep `json:"steps"`
	Summary interface{}                                            `json:"summary"`
	Title   interface{}                                            `json:"title"`
	JSON    aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON   `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON contains the JSON metadata
// for the struct [AIIncubatorPitchGetDetailsResponseAICoachingPlan]
type aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON struct {
	Steps       apijson.Field
	Summary     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAICoachingPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlanStep struct {
	Description interface{}                                                     `json:"description"`
	Resources   []AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource `json:"resources"`
	Status      AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus     `json:"status"`
	Timeline    interface{}                                                     `json:"timeline"`
	Title       interface{}                                                     `json:"title"`
	JSON        aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON        `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAICoachingPlanStep]
type aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON struct {
	Description apijson.Field
	Resources   apijson.Field
	Status      apijson.Field
	Timeline    apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAICoachingPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource struct {
	Name interface{}                                                       `json:"name"`
	URL  interface{}                                                       `json:"url"`
	JSON aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON contains the
// JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource]
type aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON struct {
	Name        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus string

const (
	AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusPending    AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus = "pending"
	AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusInProgress AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus = "in_progress"
	AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusCompleted  AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus = "completed"
)

func (r AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus) IsKnown() bool {
	switch r {
	case AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusPending, AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusInProgress, AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusCompleted:
		return true
	}
	return false
}

// AI's detailed financial model analysis.
type AIIncubatorPitchGetDetailsResponseAIFinancialModel struct {
	BreakevenPoint        interface{}                                                             `json:"breakevenPoint"`
	CapitalRequirements   interface{}                                                             `json:"capitalRequirements"`
	CostStructureAnalysis interface{}                                                             `json:"costStructureAnalysis"`
	RevenueBreakdown      interface{}                                                             `json:"revenueBreakdown"`
	SensitivityAnalysis   []AIIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysis `json:"sensitivityAnalysis"`
	JSON                  aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON                  `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIFinancialModel]
type aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON struct {
	BreakevenPoint        apijson.Field
	CapitalRequirements   apijson.Field
	CostStructureAnalysis apijson.Field
	RevenueBreakdown      apijson.Field
	SensitivityAnalysis   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIFinancialModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysis struct {
	ProjectedIrr  interface{}                                                               `json:"projectedIRR"`
	Scenario      interface{}                                                               `json:"scenario"`
	TerminalValue interface{}                                                               `json:"terminalValue"`
	JSON          aiIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysisJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysisJSON
// contains the JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysis]
type aiIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysisJSON struct {
	ProjectedIrr  apijson.Field
	Scenario      apijson.Field
	TerminalValue apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIFinancialModelSensitivityAnalysisJSON) RawJSON() string {
	return r.raw
}

// AI's detailed market analysis.
type AIIncubatorPitchGetDetailsResponseAIMarketAnalysis struct {
	CompetitiveAdvantages []interface{}                                          `json:"competitiveAdvantages"`
	GrowthOpportunities   interface{}                                            `json:"growthOpportunities"`
	RiskFactors           interface{}                                            `json:"riskFactors"`
	TargetMarketSize      interface{}                                            `json:"targetMarketSize"`
	JSON                  aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIMarketAnalysis]
type aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON struct {
	CompetitiveAdvantages apijson.Field
	GrowthOpportunities   apijson.Field
	RiskFactors           apijson.Field
	TargetMarketSize      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIMarketAnalysis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON) RawJSON() string {
	return r.raw
}

// AI's assessment of risks associated with the venture.
type AIIncubatorPitchGetDetailsResponseAIRiskAssessment struct {
	MarketRisk    interface{}                                            `json:"marketRisk"`
	TeamRisk      interface{}                                            `json:"teamRisk"`
	TechnicalRisk interface{}                                            `json:"technicalRisk"`
	JSON          aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIRiskAssessment]
type aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON struct {
	MarketRisk    apijson.Field
	TeamRisk      apijson.Field
	TechnicalRisk apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIRiskAssessment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchSubmitParams struct {
	// The user's detailed narrative business plan (e.g., executive summary, vision,
	// strategy).
	BusinessPlan param.Field[interface{}] `json:"businessPlan,required"`
	// Key financial metrics and projections for the next 3-5 years.
	FinancialProjections param.Field[AIIncubatorPitchSubmitParamsFinancialProjections] `json:"financialProjections,required"`
	// Key profiles and expertise of the founding team members.
	FoundingTeam param.Field[[]AIIncubatorPitchSubmitParamsFoundingTeam] `json:"foundingTeam,required"`
	// Detailed analysis of the target market, problem statement, and proposed
	// solution's unique value proposition.
	MarketOpportunity param.Field[interface{}] `json:"marketOpportunity,required"`
}

func (r AIIncubatorPitchSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Key financial metrics and projections for the next 3-5 years.
type AIIncubatorPitchSubmitParamsFinancialProjections struct {
	// Estimated time to profitability.
	ProfitabilityEstimate param.Field[interface{}] `json:"profitabilityEstimate"`
	// Number of years for financial projections.
	ProjectionYears param.Field[interface{}]   `json:"projectionYears"`
	RevenueForecast param.Field[[]interface{}] `json:"revenueForecast"`
	// Requested seed funding in USD.
	SeedRoundAmount param.Field[interface{}] `json:"seedRoundAmount"`
	// Pre-money valuation in USD.
	ValuationPreMoney param.Field[interface{}] `json:"valuationPreMoney"`
}

func (r AIIncubatorPitchSubmitParamsFinancialProjections) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitParamsFoundingTeam struct {
	// Relevant experience.
	Experience param.Field[interface{}] `json:"experience"`
	// Name of the team member.
	Name param.Field[interface{}] `json:"name"`
	// Role of the team member.
	Role param.Field[interface{}] `json:"role"`
}

func (r AIIncubatorPitchSubmitParamsFoundingTeam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitFeedbackParams struct {
	Answers param.Field[[]AIIncubatorPitchSubmitFeedbackParamsAnswer] `json:"answers"`
	// General textual feedback or additional details for Quantum Weaver.
	Feedback param.Field[interface{}] `json:"feedback"`
}

func (r AIIncubatorPitchSubmitFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitFeedbackParamsAnswer struct {
	// The answer to the specific question.
	Answer param.Field[interface{}] `json:"answer,required"`
	// The ID of the question being answered.
	QuestionID param.Field[interface{}] `json:"questionId,required"`
}

func (r AIIncubatorPitchSubmitFeedbackParamsAnswer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
