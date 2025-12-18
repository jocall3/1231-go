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

// AIIncubatorService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorService] method instead.
type AIIncubatorService struct {
	Options []option.RequestOption
	Pitch   *AIIncubatorPitchService
}

// NewAIIncubatorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIIncubatorService(opts ...option.RequestOption) (r *AIIncubatorService) {
	r = &AIIncubatorService{}
	r.Options = opts
	r.Pitch = NewAIIncubatorPitchService(opts...)
	return
}

// Retrieves a summary list of all business pitches submitted by the authenticated
// user to Quantum Weaver.
func (r *AIIncubatorService) ListPitches(ctx context.Context, query AIIncubatorListPitchesParams, opts ...option.RequestOption) (res *AIIncubatorListPitchesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AIIncubatorListPitchesResponse struct {
	Data []QuantumWeaverState               `json:"data"`
	JSON aiIncubatorListPitchesResponseJSON `json:"-"`
	PaginatedList
}

// aiIncubatorListPitchesResponseJSON contains the JSON metadata for the struct
// [AIIncubatorListPitchesResponse]
type aiIncubatorListPitchesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorListPitchesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorListPitchesResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorListPitchesParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
	// Filter pitches by their current stage.
	Status param.Field[AIIncubatorListPitchesParamsStatus] `query:"status"`
}

// URLQuery serializes [AIIncubatorListPitchesParams]'s query parameters as
// `url.Values`.
func (r AIIncubatorListPitchesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter pitches by their current stage.
type AIIncubatorListPitchesParamsStatus string

const (
	AIIncubatorListPitchesParamsStatusSubmitted          AIIncubatorListPitchesParamsStatus = "submitted"
	AIIncubatorListPitchesParamsStatusInitialReview      AIIncubatorListPitchesParamsStatus = "initial_review"
	AIIncubatorListPitchesParamsStatusAIAnalysis         AIIncubatorListPitchesParamsStatus = "ai_analysis"
	AIIncubatorListPitchesParamsStatusFeedbackRequired   AIIncubatorListPitchesParamsStatus = "feedback_required"
	AIIncubatorListPitchesParamsStatusTestPhase          AIIncubatorListPitchesParamsStatus = "test_phase"
	AIIncubatorListPitchesParamsStatusFinalReview        AIIncubatorListPitchesParamsStatus = "final_review"
	AIIncubatorListPitchesParamsStatusApprovedForFunding AIIncubatorListPitchesParamsStatus = "approved_for_funding"
	AIIncubatorListPitchesParamsStatusRejected           AIIncubatorListPitchesParamsStatus = "rejected"
	AIIncubatorListPitchesParamsStatusIncubatedGraduated AIIncubatorListPitchesParamsStatus = "incubated_graduated"
)

func (r AIIncubatorListPitchesParamsStatus) IsKnown() bool {
	switch r {
	case AIIncubatorListPitchesParamsStatusSubmitted, AIIncubatorListPitchesParamsStatusInitialReview, AIIncubatorListPitchesParamsStatusAIAnalysis, AIIncubatorListPitchesParamsStatusFeedbackRequired, AIIncubatorListPitchesParamsStatusTestPhase, AIIncubatorListPitchesParamsStatusFinalReview, AIIncubatorListPitchesParamsStatusApprovedForFunding, AIIncubatorListPitchesParamsStatusRejected, AIIncubatorListPitchesParamsStatusIncubatedGraduated:
		return true
	}
	return false
}
