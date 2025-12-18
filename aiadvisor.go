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

// AIAdvisorService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdvisorService] method instead.
type AIAdvisorService struct {
	Options []option.RequestOption
	Chat    *AIAdvisorChatService
}

// NewAIAdvisorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdvisorService(opts ...option.RequestOption) (r *AIAdvisorService) {
	r = &AIAdvisorService{}
	r.Options = opts
	r.Chat = NewAIAdvisorChatService(opts...)
	return
}

// Retrieves a dynamic manifest of all integrated AI tools that Quantum can invoke
// and execute, providing details on their capabilities, parameters, and access
// requirements.
func (r *AIAdvisorService) ListTools(ctx context.Context, query AIAdvisorListToolsParams, opts ...option.RequestOption) (res *AIAdvisorListToolsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AIAdvisorListToolsResponse struct {
	Data []AIAdvisorListToolsResponseData `json:"data"`
	JSON aiAdvisorListToolsResponseJSON   `json:"-"`
	PaginatedList
}

// aiAdvisorListToolsResponseJSON contains the JSON metadata for the struct
// [AIAdvisorListToolsResponse]
type aiAdvisorListToolsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorListToolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorListToolsResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorListToolsResponseData struct {
	// The OAuth2 scope required to execute this tool.
	AccessScope interface{} `json:"accessScope,required"`
	// A description of what the tool does.
	Description interface{} `json:"description,required"`
	// The unique name of the AI tool (function name).
	Name interface{} `json:"name,required"`
	// OpenAPI schema object defining the input parameters for the tool function.
	Parameters AIAdvisorListToolsResponseDataParameters `json:"parameters,required"`
	JSON       aiAdvisorListToolsResponseDataJSON       `json:"-"`
}

// aiAdvisorListToolsResponseDataJSON contains the JSON metadata for the struct
// [AIAdvisorListToolsResponseData]
type aiAdvisorListToolsResponseDataJSON struct {
	AccessScope apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Parameters  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorListToolsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorListToolsResponseDataJSON) RawJSON() string {
	return r.raw
}

// OpenAPI schema object defining the input parameters for the tool function.
type AIAdvisorListToolsResponseDataParameters struct {
	Properties interface{}                                  `json:"properties"`
	Required   []interface{}                                `json:"required"`
	Type       AIAdvisorListToolsResponseDataParametersType `json:"type"`
	JSON       aiAdvisorListToolsResponseDataParametersJSON `json:"-"`
}

// aiAdvisorListToolsResponseDataParametersJSON contains the JSON metadata for the
// struct [AIAdvisorListToolsResponseDataParameters]
type aiAdvisorListToolsResponseDataParametersJSON struct {
	Properties  apijson.Field
	Required    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorListToolsResponseDataParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorListToolsResponseDataParametersJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorListToolsResponseDataParametersType string

const (
	AIAdvisorListToolsResponseDataParametersTypeObject AIAdvisorListToolsResponseDataParametersType = "object"
)

func (r AIAdvisorListToolsResponseDataParametersType) IsKnown() bool {
	switch r {
	case AIAdvisorListToolsResponseDataParametersTypeObject:
		return true
	}
	return false
}

type AIAdvisorListToolsParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [AIAdvisorListToolsParams]'s query parameters as
// `url.Values`.
func (r AIAdvisorListToolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
