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

// AIOracleSimulateService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOracleSimulateService] method instead.
type AIOracleSimulateService struct {
	Options []option.RequestOption
}

// NewAIOracleSimulateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOracleSimulateService(opts ...option.RequestOption) (r *AIOracleSimulateService) {
	r = &AIOracleSimulateService{}
	r.Options = opts
	return
}

// Engages the Quantum Oracle for highly complex, multi-variable simulations,
// allowing precise control over numerous financial parameters, market conditions,
// and personal events to generate deep, predictive insights and sensitivity
// analysis.
func (r *AIOracleSimulateService) RunAdvanced(ctx context.Context, body AIOracleSimulateRunAdvancedParams, opts ...option.RequestOption) (res *AdvancedSimulationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulate/advanced"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submits a hypothetical scenario to the Quantum Oracle AI for standard financial
// impact analysis. The AI simulates the effect on the user's current financial
// state and provides a summary.
func (r *AIOracleSimulateService) RunStandard(ctx context.Context, body AIOracleSimulateRunStandardParams, opts ...option.RequestOption) (res *SimulationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AdvancedSimulationResponse struct {
	// A high-level summary of findings across all scenarios.
	OverallSummary  interface{}                                `json:"overallSummary,required"`
	ScenarioResults []AdvancedSimulationResponseScenarioResult `json:"scenarioResults,required"`
	// Unique identifier for the completed advanced simulation.
	SimulationID interface{} `json:"simulationId,required"`
	// Overarching strategic recommendations derived from the comparison of scenarios.
	StrategicRecommendations []AIInsight                    `json:"strategicRecommendations,nullable"`
	JSON                     advancedSimulationResponseJSON `json:"-"`
}

// advancedSimulationResponseJSON contains the JSON metadata for the struct
// [AdvancedSimulationResponse]
type advancedSimulationResponseJSON struct {
	OverallSummary           apijson.Field
	ScenarioResults          apijson.Field
	SimulationID             apijson.Field
	StrategicRecommendations apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AdvancedSimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r AdvancedSimulationResponse) implementsAIOracleSimulationGetResponse() {}

type AdvancedSimulationResponseScenarioResult struct {
	// Summary of results for this specific scenario.
	NarrativeSummary interface{} `json:"narrativeSummary,required"`
	// Name of the individual scenario.
	ScenarioName interface{} `json:"scenarioName,required"`
	// Specific AI insights for this scenario.
	AIInsights []AIInsight `json:"aiInsights,nullable"`
	// Projected net worth at the end of the simulation period for this scenario.
	FinalNetWorthProjected interface{}                                               `json:"finalNetWorthProjected"`
	LiquidityMetrics       AdvancedSimulationResponseScenarioResultsLiquidityMetrics `json:"liquidityMetrics,nullable"`
	// Data for generating sensitivity analysis charts (e.g., how net worth changes as
	// a variable is adjusted).
	SensitivityAnalysisGraphs []AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph `json:"sensitivityAnalysisGraphs,nullable"`
	JSON                      advancedSimulationResponseScenarioResultJSON                        `json:"-"`
}

// advancedSimulationResponseScenarioResultJSON contains the JSON metadata for the
// struct [AdvancedSimulationResponseScenarioResult]
type advancedSimulationResponseScenarioResultJSON struct {
	NarrativeSummary          apijson.Field
	ScenarioName              apijson.Field
	AIInsights                apijson.Field
	FinalNetWorthProjected    apijson.Field
	LiquidityMetrics          apijson.Field
	SensitivityAnalysisGraphs apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultJSON) RawJSON() string {
	return r.raw
}

type AdvancedSimulationResponseScenarioResultsLiquidityMetrics struct {
	// Minimum cash balance reached during the scenario.
	MinCashBalance interface{} `json:"minCashBalance"`
	// Time in months to recover to pre-event financial state.
	RecoveryTimeMonths interface{}                                                   `json:"recoveryTimeMonths"`
	JSON               advancedSimulationResponseScenarioResultsLiquidityMetricsJSON `json:"-"`
}

// advancedSimulationResponseScenarioResultsLiquidityMetricsJSON contains the JSON
// metadata for the struct
// [AdvancedSimulationResponseScenarioResultsLiquidityMetrics]
type advancedSimulationResponseScenarioResultsLiquidityMetricsJSON struct {
	MinCashBalance     apijson.Field
	RecoveryTimeMonths apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResultsLiquidityMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultsLiquidityMetricsJSON) RawJSON() string {
	return r.raw
}

type AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph struct {
	Data      []AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData `json:"data"`
	ParamName interface{}                                                              `json:"paramName"`
	JSON      advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON    `json:"-"`
}

// advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON contains
// the JSON metadata for the struct
// [AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph]
type advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON struct {
	Data        apijson.Field
	ParamName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON) RawJSON() string {
	return r.raw
}

type AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData struct {
	OutcomeValue interface{}                                                                `json:"outcomeValue"`
	ParamValue   interface{}                                                                `json:"paramValue"`
	JSON         advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON `json:"-"`
}

// advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON
// contains the JSON metadata for the struct
// [AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData]
type advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON struct {
	OutcomeValue apijson.Field
	ParamValue   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON) RawJSON() string {
	return r.raw
}

type SimulationResponse struct {
	// Key quantitative and qualitative impacts identified by the AI.
	KeyImpacts []SimulationResponseKeyImpact `json:"keyImpacts,required"`
	// A natural language summary of the simulation's results and key findings.
	NarrativeSummary interface{} `json:"narrativeSummary,required"`
	// Unique identifier for the completed simulation.
	SimulationID interface{} `json:"simulationId,required"`
	// Actionable recommendations derived from the simulation.
	Recommendations []AIInsight `json:"recommendations,nullable"`
	// AI-driven risk assessment of the simulated scenario.
	RiskAnalysis SimulationResponseRiskAnalysis `json:"riskAnalysis"`
	// Optional: URLs to generated visualization data or images.
	Visualizations []SimulationResponseVisualization `json:"visualizations,nullable"`
	JSON           simulationResponseJSON            `json:"-"`
}

// simulationResponseJSON contains the JSON metadata for the struct
// [SimulationResponse]
type simulationResponseJSON struct {
	KeyImpacts       apijson.Field
	NarrativeSummary apijson.Field
	SimulationID     apijson.Field
	Recommendations  apijson.Field
	RiskAnalysis     apijson.Field
	Visualizations   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r SimulationResponse) implementsAIOracleSimulationGetResponse() {}

type SimulationResponseKeyImpact struct {
	Metric   interface{}                          `json:"metric"`
	Severity SimulationResponseKeyImpactsSeverity `json:"severity"`
	Value    interface{}                          `json:"value"`
	JSON     simulationResponseKeyImpactJSON      `json:"-"`
}

// simulationResponseKeyImpactJSON contains the JSON metadata for the struct
// [SimulationResponseKeyImpact]
type simulationResponseKeyImpactJSON struct {
	Metric      apijson.Field
	Severity    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SimulationResponseKeyImpact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseKeyImpactJSON) RawJSON() string {
	return r.raw
}

type SimulationResponseKeyImpactsSeverity string

const (
	SimulationResponseKeyImpactsSeverityLow    SimulationResponseKeyImpactsSeverity = "low"
	SimulationResponseKeyImpactsSeverityMedium SimulationResponseKeyImpactsSeverity = "medium"
	SimulationResponseKeyImpactsSeverityHigh   SimulationResponseKeyImpactsSeverity = "high"
)

func (r SimulationResponseKeyImpactsSeverity) IsKnown() bool {
	switch r {
	case SimulationResponseKeyImpactsSeverityLow, SimulationResponseKeyImpactsSeverityMedium, SimulationResponseKeyImpactsSeverityHigh:
		return true
	}
	return false
}

// AI-driven risk assessment of the simulated scenario.
type SimulationResponseRiskAnalysis struct {
	// Maximum potential loss from peak to trough (e.g., 0.25 for 25%).
	MaxDrawdown interface{} `json:"maxDrawdown"`
	// Measure of market volatility associated with the scenario.
	VolatilityIndex interface{}                        `json:"volatilityIndex"`
	JSON            simulationResponseRiskAnalysisJSON `json:"-"`
}

// simulationResponseRiskAnalysisJSON contains the JSON metadata for the struct
// [SimulationResponseRiskAnalysis]
type simulationResponseRiskAnalysisJSON struct {
	MaxDrawdown     apijson.Field
	VolatilityIndex apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SimulationResponseRiskAnalysis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseRiskAnalysisJSON) RawJSON() string {
	return r.raw
}

type SimulationResponseVisualization struct {
	DataUri interface{}                          `json:"dataUri"`
	Title   interface{}                          `json:"title"`
	Type    SimulationResponseVisualizationsType `json:"type"`
	JSON    simulationResponseVisualizationJSON  `json:"-"`
}

// simulationResponseVisualizationJSON contains the JSON metadata for the struct
// [SimulationResponseVisualization]
type simulationResponseVisualizationJSON struct {
	DataUri     apijson.Field
	Title       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SimulationResponseVisualization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseVisualizationJSON) RawJSON() string {
	return r.raw
}

type SimulationResponseVisualizationsType string

const (
	SimulationResponseVisualizationsTypeLineChart SimulationResponseVisualizationsType = "line_chart"
	SimulationResponseVisualizationsTypeBarChart  SimulationResponseVisualizationsType = "bar_chart"
	SimulationResponseVisualizationsTypeTable     SimulationResponseVisualizationsType = "table"
)

func (r SimulationResponseVisualizationsType) IsKnown() bool {
	switch r {
	case SimulationResponseVisualizationsTypeLineChart, SimulationResponseVisualizationsTypeBarChart, SimulationResponseVisualizationsTypeTable:
		return true
	}
	return false
}

type AIOracleSimulateRunAdvancedParams struct {
	// A natural language prompt describing the complex, multi-variable scenario.
	Prompt    param.Field[interface{}]                                 `json:"prompt,required"`
	Scenarios param.Field[[]AIOracleSimulateRunAdvancedParamsScenario] `json:"scenarios,required"`
	// Optional: Global economic conditions to apply to all scenarios.
	GlobalEconomicFactors param.Field[AIOracleSimulateRunAdvancedParamsGlobalEconomicFactors] `json:"globalEconomicFactors"`
	// Optional: Personal financial assumptions to override defaults.
	PersonalAssumptions param.Field[AIOracleSimulateRunAdvancedParamsPersonalAssumptions] `json:"personalAssumptions"`
}

func (r AIOracleSimulateRunAdvancedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenario struct {
	// The duration in years over which this scenario is simulated.
	DurationYears param.Field[interface{}] `json:"durationYears,required"`
	// A list of discrete or continuous events that define this scenario.
	Events param.Field[[]AIOracleSimulateRunAdvancedParamsScenariosEvent] `json:"events,required"`
	// A descriptive name for this specific scenario.
	Name param.Field[interface{}] `json:"name,required"`
	// Parameters for multi-variable sensitivity analysis within this scenario.
	SensitivityAnalysisParams param.Field[[]AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam] `json:"sensitivityAnalysisParams"`
}

func (r AIOracleSimulateRunAdvancedParamsScenario) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenariosEvent struct {
	// Specific parameters for the event (e.g., durationMonths, impactPercentage).
	Details param.Field[interface{}]                                          `json:"details"`
	Type    param.Field[AIOracleSimulateRunAdvancedParamsScenariosEventsType] `json:"type"`
}

func (r AIOracleSimulateRunAdvancedParamsScenariosEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenariosEventsType string

const (
	AIOracleSimulateRunAdvancedParamsScenariosEventsTypeJobLoss          AIOracleSimulateRunAdvancedParamsScenariosEventsType = "job_loss"
	AIOracleSimulateRunAdvancedParamsScenariosEventsTypeMarketDownturn   AIOracleSimulateRunAdvancedParamsScenariosEventsType = "market_downturn"
	AIOracleSimulateRunAdvancedParamsScenariosEventsTypeLargePurchase    AIOracleSimulateRunAdvancedParamsScenariosEventsType = "large_purchase"
	AIOracleSimulateRunAdvancedParamsScenariosEventsTypeIncomeIncrease   AIOracleSimulateRunAdvancedParamsScenariosEventsType = "income_increase"
	AIOracleSimulateRunAdvancedParamsScenariosEventsTypeMedicalEmergency AIOracleSimulateRunAdvancedParamsScenariosEventsType = "medical_emergency"
)

func (r AIOracleSimulateRunAdvancedParamsScenariosEventsType) IsKnown() bool {
	switch r {
	case AIOracleSimulateRunAdvancedParamsScenariosEventsTypeJobLoss, AIOracleSimulateRunAdvancedParamsScenariosEventsTypeMarketDownturn, AIOracleSimulateRunAdvancedParamsScenariosEventsTypeLargePurchase, AIOracleSimulateRunAdvancedParamsScenariosEventsTypeIncomeIncrease, AIOracleSimulateRunAdvancedParamsScenariosEventsTypeMedicalEmergency:
		return true
	}
	return false
}

type AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam struct {
	// Maximum value for the parameter.
	Max param.Field[interface{}] `json:"max"`
	// Minimum value for the parameter.
	Min param.Field[interface{}] `json:"min"`
	// The name of the parameter to vary for sensitivity analysis (e.g.,
	// 'interestRate', 'inflationRate', 'marketRecoveryRate').
	ParamName param.Field[interface{}] `json:"paramName"`
	// Step increment for varying the parameter.
	Step param.Field[interface{}] `json:"step"`
}

func (r AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional: Global economic conditions to apply to all scenarios.
type AIOracleSimulateRunAdvancedParamsGlobalEconomicFactors struct {
	InflationRate        param.Field[interface{}] `json:"inflationRate"`
	InterestRateBaseline param.Field[interface{}] `json:"interestRateBaseline"`
}

func (r AIOracleSimulateRunAdvancedParamsGlobalEconomicFactors) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional: Personal financial assumptions to override defaults.
type AIOracleSimulateRunAdvancedParamsPersonalAssumptions struct {
	AnnualSavingsRate param.Field[interface{}]                                                       `json:"annualSavingsRate"`
	RiskTolerance     param.Field[AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskTolerance] `json:"riskTolerance"`
}

func (r AIOracleSimulateRunAdvancedParamsPersonalAssumptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskTolerance string

const (
	AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceConservative AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskTolerance = "conservative"
	AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceModerate     AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskTolerance = "moderate"
	AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceAggressive   AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskTolerance = "aggressive"
)

func (r AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskTolerance) IsKnown() bool {
	switch r {
	case AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceConservative, AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceModerate, AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceAggressive:
		return true
	}
	return false
}

type AIOracleSimulateRunStandardParams struct {
	// A natural language prompt describing the 'what-if' scenario.
	Prompt param.Field[interface{}] `json:"prompt,required"`
	// Optional structured parameters to guide the simulation (e.g., duration, amount,
	// risk tolerance).
	Parameters param.Field[interface{}] `json:"parameters"`
}

func (r AIOracleSimulateRunStandardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
