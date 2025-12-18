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

// AIAdvisorChatService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdvisorChatService] method instead.
type AIAdvisorChatService struct {
	Options []option.RequestOption
}

// NewAIAdvisorChatService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdvisorChatService(opts ...option.RequestOption) (r *AIAdvisorChatService) {
	r = &AIAdvisorChatService{}
	r.Options = opts
	return
}

// Fetches the full conversation history with the Quantum AI Advisor for a given
// session or user.
func (r *AIAdvisorChatService) GetHistory(ctx context.Context, query AIAdvisorChatGetHistoryParams, opts ...option.RequestOption) (res *AIAdvisorChatGetHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/chat/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiates or continues a sophisticated conversation with Quantum, the AI
// Advisor. Quantum can provide advanced financial insights, execute complex tasks
// via an expanding suite of intelligent tools, and learn from user interactions to
// offer hyper-personalized guidance.
func (r *AIAdvisorChatService) SendMessage(ctx context.Context, body AIAdvisorChatSendMessageParams, opts ...option.RequestOption) (res *AIAdvisorChatSendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIAdvisorChatGetHistoryResponse struct {
	Data []AIAdvisorChatGetHistoryResponseData `json:"data"`
	JSON aiAdvisorChatGetHistoryResponseJSON   `json:"-"`
	PaginatedList
}

// aiAdvisorChatGetHistoryResponseJSON contains the JSON metadata for the struct
// [AIAdvisorChatGetHistoryResponse]
type aiAdvisorChatGetHistoryResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatGetHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatGetHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatGetHistoryResponseData struct {
	// The textual content of the message.
	Content interface{} `json:"content,required"`
	// Role of the speaker (user, assistant, or tool interaction).
	Role AIAdvisorChatGetHistoryResponseDataRole `json:"role,required"`
	// Timestamp of the message.
	Timestamp interface{} `json:"timestamp,required"`
	// If role is 'tool_call', details of the tool function called by the AI.
	FunctionCall AIAdvisorChatGetHistoryResponseDataFunctionCall `json:"functionCall"`
	// If role is 'tool_response', the output from the tool function.
	FunctionResponse AIAdvisorChatGetHistoryResponseDataFunctionResponse `json:"functionResponse"`
	JSON             aiAdvisorChatGetHistoryResponseDataJSON             `json:"-"`
}

// aiAdvisorChatGetHistoryResponseDataJSON contains the JSON metadata for the
// struct [AIAdvisorChatGetHistoryResponseData]
type aiAdvisorChatGetHistoryResponseDataJSON struct {
	Content          apijson.Field
	Role             apijson.Field
	Timestamp        apijson.Field
	FunctionCall     apijson.Field
	FunctionResponse apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIAdvisorChatGetHistoryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatGetHistoryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Role of the speaker (user, assistant, or tool interaction).
type AIAdvisorChatGetHistoryResponseDataRole string

const (
	AIAdvisorChatGetHistoryResponseDataRoleUser         AIAdvisorChatGetHistoryResponseDataRole = "user"
	AIAdvisorChatGetHistoryResponseDataRoleAssistant    AIAdvisorChatGetHistoryResponseDataRole = "assistant"
	AIAdvisorChatGetHistoryResponseDataRoleToolCall     AIAdvisorChatGetHistoryResponseDataRole = "tool_call"
	AIAdvisorChatGetHistoryResponseDataRoleToolResponse AIAdvisorChatGetHistoryResponseDataRole = "tool_response"
)

func (r AIAdvisorChatGetHistoryResponseDataRole) IsKnown() bool {
	switch r {
	case AIAdvisorChatGetHistoryResponseDataRoleUser, AIAdvisorChatGetHistoryResponseDataRoleAssistant, AIAdvisorChatGetHistoryResponseDataRoleToolCall, AIAdvisorChatGetHistoryResponseDataRoleToolResponse:
		return true
	}
	return false
}

// If role is 'tool_call', details of the tool function called by the AI.
type AIAdvisorChatGetHistoryResponseDataFunctionCall struct {
	Args interface{}                                         `json:"args"`
	Name interface{}                                         `json:"name"`
	JSON aiAdvisorChatGetHistoryResponseDataFunctionCallJSON `json:"-"`
}

// aiAdvisorChatGetHistoryResponseDataFunctionCallJSON contains the JSON metadata
// for the struct [AIAdvisorChatGetHistoryResponseDataFunctionCall]
type aiAdvisorChatGetHistoryResponseDataFunctionCallJSON struct {
	Args        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatGetHistoryResponseDataFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatGetHistoryResponseDataFunctionCallJSON) RawJSON() string {
	return r.raw
}

// If role is 'tool_response', the output from the tool function.
type AIAdvisorChatGetHistoryResponseDataFunctionResponse struct {
	Name     interface{}                                             `json:"name"`
	Response interface{}                                             `json:"response"`
	JSON     aiAdvisorChatGetHistoryResponseDataFunctionResponseJSON `json:"-"`
}

// aiAdvisorChatGetHistoryResponseDataFunctionResponseJSON contains the JSON
// metadata for the struct [AIAdvisorChatGetHistoryResponseDataFunctionResponse]
type aiAdvisorChatGetHistoryResponseDataFunctionResponseJSON struct {
	Name        apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatGetHistoryResponseDataFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatGetHistoryResponseDataFunctionResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatSendMessageResponse struct {
	// The active conversation session ID.
	SessionID interface{} `json:"sessionId,required"`
	// A list of tool functions the AI wants the system to execute.
	FunctionCalls []AIAdvisorChatSendMessageResponseFunctionCall `json:"functionCalls,nullable"`
	// A list of proactive AI insights or recommendations generated by Quantum.
	ProactiveInsights []AIInsight `json:"proactiveInsights,nullable"`
	// Indicates if the AI's response implies that the user needs to take a specific
	// action (e.g., provide more input, confirm a tool call).
	RequiresUserAction interface{} `json:"requiresUserAction"`
	// The AI Advisor's textual response.
	Text interface{}                          `json:"text"`
	JSON aiAdvisorChatSendMessageResponseJSON `json:"-"`
}

// aiAdvisorChatSendMessageResponseJSON contains the JSON metadata for the struct
// [AIAdvisorChatSendMessageResponse]
type aiAdvisorChatSendMessageResponseJSON struct {
	SessionID          apijson.Field
	FunctionCalls      apijson.Field
	ProactiveInsights  apijson.Field
	RequiresUserAction apijson.Field
	Text               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AIAdvisorChatSendMessageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatSendMessageResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatSendMessageResponseFunctionCall struct {
	// Unique ID for this tool call, used to link with `functionResponse`.
	ID interface{} `json:"id"`
	// Key-value pairs representing the arguments to pass to the tool function.
	Args interface{} `json:"args"`
	// The name of the tool function to call.
	Name interface{}                                      `json:"name"`
	JSON aiAdvisorChatSendMessageResponseFunctionCallJSON `json:"-"`
}

// aiAdvisorChatSendMessageResponseFunctionCallJSON contains the JSON metadata for
// the struct [AIAdvisorChatSendMessageResponseFunctionCall]
type aiAdvisorChatSendMessageResponseFunctionCallJSON struct {
	ID          apijson.Field
	Args        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatSendMessageResponseFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatSendMessageResponseFunctionCallJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatGetHistoryParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
	// Optional: Filter history by a specific session ID. If omitted, recent
	// conversations will be returned.
	SessionID param.Field[interface{}] `query:"sessionId"`
}

// URLQuery serializes [AIAdvisorChatGetHistoryParams]'s query parameters as
// `url.Values`.
func (r AIAdvisorChatGetHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAdvisorChatSendMessageParams struct {
	// Optional: The output from a tool function that the AI previously requested to be
	// executed.
	FunctionResponse param.Field[AIAdvisorChatSendMessageParamsFunctionResponse] `json:"functionResponse"`
	// The user's textual input to the AI Advisor.
	Message param.Field[interface{}] `json:"message"`
	// Optional: Session ID to continue a conversation. If omitted, a new session is
	// started.
	SessionID param.Field[interface{}] `json:"sessionId"`
}

func (r AIAdvisorChatSendMessageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional: The output from a tool function that the AI previously requested to be
// executed.
type AIAdvisorChatSendMessageParamsFunctionResponse struct {
	// The name of the tool function for which this is a response.
	Name param.Field[interface{}] `json:"name"`
	// The JSON output from the execution of the tool function.
	Response param.Field[interface{}] `json:"response"`
}

func (r AIAdvisorChatSendMessageParamsFunctionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
