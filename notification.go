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

// NotificationService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options  []option.RequestOption
	Settings *NotificationSettingService
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r *NotificationService) {
	r = &NotificationService{}
	r.Options = opts
	r.Settings = NewNotificationSettingService(opts...)
	return
}

// Retrieves a paginated list of personalized notifications and proactive AI alerts
// for the authenticated user, allowing filtering by status and severity.
func (r *NotificationService) ListUserNotifications(ctx context.Context, query NotificationListUserNotificationsParams, opts ...option.RequestOption) (res *NotificationListUserNotificationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Marks a specific user notification as read.
func (r *NotificationService) MarkAsRead(ctx context.Context, notificationID interface{}, opts ...option.RequestOption) (res *Notification, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("notifications/%v/mark-read", notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type Notification struct {
	// Unique identifier for the notification.
	ID interface{} `json:"id,required"`
	// Full message content of the notification.
	Message interface{} `json:"message,required"`
	// Indicates if the user has read the notification.
	Read interface{} `json:"read,required"`
	// Severity of the notification (AI-assessed).
	Severity NotificationSeverity `json:"severity,required"`
	// Timestamp when the notification was generated.
	Timestamp interface{} `json:"timestamp,required"`
	// Concise title for the notification.
	Title interface{} `json:"title,required"`
	// Type of notification.
	Type NotificationType `json:"type,required"`
	// Optional deep link for the user to take action related to the notification.
	ActionableLink interface{} `json:"actionableLink"`
	// If applicable, the ID of the AIInsight that generated this notification.
	AIInsightID interface{}      `json:"aiInsightId"`
	JSON        notificationJSON `json:"-"`
}

// notificationJSON contains the JSON metadata for the struct [Notification]
type notificationJSON struct {
	ID             apijson.Field
	Message        apijson.Field
	Read           apijson.Field
	Severity       apijson.Field
	Timestamp      apijson.Field
	Title          apijson.Field
	Type           apijson.Field
	ActionableLink apijson.Field
	AIInsightID    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Notification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationJSON) RawJSON() string {
	return r.raw
}

// Severity of the notification (AI-assessed).
type NotificationSeverity string

const (
	NotificationSeverityLow      NotificationSeverity = "low"
	NotificationSeverityMedium   NotificationSeverity = "medium"
	NotificationSeverityHigh     NotificationSeverity = "high"
	NotificationSeverityCritical NotificationSeverity = "critical"
)

func (r NotificationSeverity) IsKnown() bool {
	switch r {
	case NotificationSeverityLow, NotificationSeverityMedium, NotificationSeverityHigh, NotificationSeverityCritical:
		return true
	}
	return false
}

// Type of notification.
type NotificationType string

const (
	NotificationTypeSecurity         NotificationType = "security"
	NotificationTypeFinancialInsight NotificationType = "financial_insight"
	NotificationTypeMarketing        NotificationType = "marketing"
	NotificationTypeSystemUpdate     NotificationType = "system_update"
	NotificationTypeTransaction      NotificationType = "transaction"
)

func (r NotificationType) IsKnown() bool {
	switch r {
	case NotificationTypeSecurity, NotificationTypeFinancialInsight, NotificationTypeMarketing, NotificationTypeSystemUpdate, NotificationTypeTransaction:
		return true
	}
	return false
}

type NotificationListUserNotificationsResponse struct {
	Data []Notification                                `json:"data"`
	JSON notificationListUserNotificationsResponseJSON `json:"-"`
	PaginatedList
}

// notificationListUserNotificationsResponseJSON contains the JSON metadata for the
// struct [NotificationListUserNotificationsResponse]
type notificationListUserNotificationsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationListUserNotificationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationListUserNotificationsResponseJSON) RawJSON() string {
	return r.raw
}

type NotificationListUserNotificationsParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
	// Filter notifications by AI-assigned severity level.
	Severity param.Field[NotificationListUserNotificationsParamsSeverity] `query:"severity"`
	// Filter notifications by their read status.
	Status param.Field[NotificationListUserNotificationsParamsStatus] `query:"status"`
}

// URLQuery serializes [NotificationListUserNotificationsParams]'s query parameters
// as `url.Values`.
func (r NotificationListUserNotificationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter notifications by AI-assigned severity level.
type NotificationListUserNotificationsParamsSeverity string

const (
	NotificationListUserNotificationsParamsSeverityLow      NotificationListUserNotificationsParamsSeverity = "low"
	NotificationListUserNotificationsParamsSeverityMedium   NotificationListUserNotificationsParamsSeverity = "medium"
	NotificationListUserNotificationsParamsSeverityHigh     NotificationListUserNotificationsParamsSeverity = "high"
	NotificationListUserNotificationsParamsSeverityCritical NotificationListUserNotificationsParamsSeverity = "critical"
)

func (r NotificationListUserNotificationsParamsSeverity) IsKnown() bool {
	switch r {
	case NotificationListUserNotificationsParamsSeverityLow, NotificationListUserNotificationsParamsSeverityMedium, NotificationListUserNotificationsParamsSeverityHigh, NotificationListUserNotificationsParamsSeverityCritical:
		return true
	}
	return false
}

// Filter notifications by their read status.
type NotificationListUserNotificationsParamsStatus string

const (
	NotificationListUserNotificationsParamsStatusRead   NotificationListUserNotificationsParamsStatus = "read"
	NotificationListUserNotificationsParamsStatusUnread NotificationListUserNotificationsParamsStatus = "unread"
)

func (r NotificationListUserNotificationsParamsStatus) IsKnown() bool {
	switch r {
	case NotificationListUserNotificationsParamsStatusRead, NotificationListUserNotificationsParamsStatusUnread:
		return true
	}
	return false
}
