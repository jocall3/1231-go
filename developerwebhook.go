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

// DeveloperWebhookService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeveloperWebhookService] method instead.
type DeveloperWebhookService struct {
	Options []option.RequestOption
}

// NewDeveloperWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeveloperWebhookService(opts ...option.RequestOption) (r *DeveloperWebhookService) {
	r = &DeveloperWebhookService{}
	r.Options = opts
	return
}

// Establishes a new webhook subscription, allowing a developer application to
// receive real-time notifications for specified events (e.g., new transaction,
// account update) via a provided callback URL.
func (r *DeveloperWebhookService) New(ctx context.Context, body DeveloperWebhookNewParams, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "developers/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Modifies an existing webhook subscription, allowing changes to the callback URL,
// subscribed events, or activation status.
func (r *DeveloperWebhookService) Update(ctx context.Context, subscriptionID interface{}, body DeveloperWebhookUpdateParams, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("developers/webhooks/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of all active webhook subscriptions for the authenticated
// developer application, detailing endpoint URLs, subscribed events, and current
// status.
func (r *DeveloperWebhookService) List(ctx context.Context, query DeveloperWebhookListParams, opts ...option.RequestOption) (res *DeveloperWebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "developers/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing webhook subscription, stopping all future event
// notifications to the specified callback URL.
func (r *DeveloperWebhookService) Delete(ctx context.Context, subscriptionID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("developers/webhooks/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type WebhookSubscription struct {
	// Unique identifier for the webhook subscription.
	ID interface{} `json:"id,required"`
	// The URL where webhook events will be sent.
	CallbackURL interface{} `json:"callbackUrl,required"`
	// Timestamp when the subscription was created.
	CreatedAt interface{} `json:"createdAt,required"`
	// List of event types subscribed to.
	Events []interface{} `json:"events,required"`
	// Current status of the webhook subscription.
	Status WebhookSubscriptionStatus `json:"status,required"`
	// Number of consecutive failed delivery attempts.
	FailureCount interface{} `json:"failureCount"`
	// Timestamp of the last successful webhook delivery.
	LastTriggered interface{} `json:"lastTriggered"`
	// The shared secret used to sign webhook payloads, for verification. Only returned
	// on creation.
	Secret interface{}             `json:"secret"`
	JSON   webhookSubscriptionJSON `json:"-"`
}

// webhookSubscriptionJSON contains the JSON metadata for the struct
// [WebhookSubscription]
type webhookSubscriptionJSON struct {
	ID            apijson.Field
	CallbackURL   apijson.Field
	CreatedAt     apijson.Field
	Events        apijson.Field
	Status        apijson.Field
	FailureCount  apijson.Field
	LastTriggered apijson.Field
	Secret        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WebhookSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookSubscriptionJSON) RawJSON() string {
	return r.raw
}

// Current status of the webhook subscription.
type WebhookSubscriptionStatus string

const (
	WebhookSubscriptionStatusActive    WebhookSubscriptionStatus = "active"
	WebhookSubscriptionStatusPaused    WebhookSubscriptionStatus = "paused"
	WebhookSubscriptionStatusSuspended WebhookSubscriptionStatus = "suspended"
)

func (r WebhookSubscriptionStatus) IsKnown() bool {
	switch r {
	case WebhookSubscriptionStatusActive, WebhookSubscriptionStatusPaused, WebhookSubscriptionStatusSuspended:
		return true
	}
	return false
}

type DeveloperWebhookListResponse struct {
	Data []WebhookSubscription            `json:"data"`
	JSON developerWebhookListResponseJSON `json:"-"`
	PaginatedList
}

// developerWebhookListResponseJSON contains the JSON metadata for the struct
// [DeveloperWebhookListResponse]
type developerWebhookListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeveloperWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r developerWebhookListResponseJSON) RawJSON() string {
	return r.raw
}

type DeveloperWebhookNewParams struct {
	// The URL to which webhook events will be sent.
	CallbackURL param.Field[interface{}] `json:"callbackUrl,required"`
	// List of event types to subscribe to.
	Events param.Field[[]interface{}] `json:"events,required"`
	// Optional: A custom shared secret for verifying webhook payloads. If omitted, one
	// will be generated.
	Secret param.Field[interface{}] `json:"secret"`
}

func (r DeveloperWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeveloperWebhookUpdateParams struct {
	// Updated URL where webhook events will be sent.
	CallbackURL param.Field[interface{}] `json:"callbackUrl"`
	// Updated list of event types subscribed to.
	Events param.Field[[]interface{}] `json:"events"`
	// Updated status of the webhook subscription.
	Status param.Field[DeveloperWebhookUpdateParamsStatus] `json:"status"`
}

func (r DeveloperWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated status of the webhook subscription.
type DeveloperWebhookUpdateParamsStatus string

const (
	DeveloperWebhookUpdateParamsStatusActive    DeveloperWebhookUpdateParamsStatus = "active"
	DeveloperWebhookUpdateParamsStatusPaused    DeveloperWebhookUpdateParamsStatus = "paused"
	DeveloperWebhookUpdateParamsStatusSuspended DeveloperWebhookUpdateParamsStatus = "suspended"
)

func (r DeveloperWebhookUpdateParamsStatus) IsKnown() bool {
	switch r {
	case DeveloperWebhookUpdateParamsStatusActive, DeveloperWebhookUpdateParamsStatusPaused, DeveloperWebhookUpdateParamsStatusSuspended:
		return true
	}
	return false
}

type DeveloperWebhookListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [DeveloperWebhookListParams]'s query parameters as
// `url.Values`.
func (r DeveloperWebhookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
