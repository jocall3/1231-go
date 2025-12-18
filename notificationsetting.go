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

// NotificationSettingService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationSettingService] method instead.
type NotificationSettingService struct {
	Options []option.RequestOption
}

// NewNotificationSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationSettingService(opts ...option.RequestOption) (r *NotificationSettingService) {
	r = &NotificationSettingService{}
	r.Options = opts
	return
}

// Retrieves the user's granular notification preferences across different channels
// (email, push, SMS, in-app) and event types.
func (r *NotificationSettingService) Get(ctx context.Context, opts ...option.RequestOption) (res *NotificationSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications/settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the user's notification preferences, allowing control over channels,
// event types, and quiet hours.
func (r *NotificationSettingService) Update(ctx context.Context, body NotificationSettingUpdateParams, opts ...option.RequestOption) (res *NotificationSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications/settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type NotificationSettings struct {
	// Preferences for notification delivery channels.
	ChannelPreferences NotificationSettingsChannelPreferences `json:"channelPreferences,required"`
	// Preferences for different types of events.
	EventPreferences NotificationSettingsEventPreferences `json:"eventPreferences,required"`
	// Settings for notification quiet hours.
	QuietHours NotificationSettingsQuietHours `json:"quietHours,required"`
	JSON       notificationSettingsJSON       `json:"-"`
}

// notificationSettingsJSON contains the JSON metadata for the struct
// [NotificationSettings]
type notificationSettingsJSON struct {
	ChannelPreferences apijson.Field
	EventPreferences   apijson.Field
	QuietHours         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Preferences for notification delivery channels.
type NotificationSettingsChannelPreferences struct {
	// Receive notifications via email.
	Email interface{} `json:"email"`
	// Receive notifications within the application.
	InApp interface{} `json:"inApp"`
	// Receive notifications via push notifications.
	Push interface{} `json:"push"`
	// Receive notifications via SMS.
	SMS  interface{}                                `json:"sms"`
	JSON notificationSettingsChannelPreferencesJSON `json:"-"`
}

// notificationSettingsChannelPreferencesJSON contains the JSON metadata for the
// struct [NotificationSettingsChannelPreferences]
type notificationSettingsChannelPreferencesJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationSettingsChannelPreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationSettingsChannelPreferencesJSON) RawJSON() string {
	return r.raw
}

// Preferences for different types of events.
type NotificationSettingsEventPreferences struct {
	// Receive proactive AI-driven financial insights.
	AIInsights interface{} `json:"aiInsights"`
	// Receive alerts for budget progress (e.g., nearing limit, over budget).
	BudgetAlerts interface{} `json:"budgetAlerts"`
	// Receive promotional offers and marketing communications.
	PromotionalOffers interface{} `json:"promotionalOffers"`
	// Receive critical security alerts (e.g., suspicious login).
	SecurityAlerts interface{} `json:"securityAlerts"`
	// Receive alerts for transactions (e.g., large spend, recurring charges).
	TransactionAlerts interface{}                              `json:"transactionAlerts"`
	JSON              notificationSettingsEventPreferencesJSON `json:"-"`
}

// notificationSettingsEventPreferencesJSON contains the JSON metadata for the
// struct [NotificationSettingsEventPreferences]
type notificationSettingsEventPreferencesJSON struct {
	AIInsights        apijson.Field
	BudgetAlerts      apijson.Field
	PromotionalOffers apijson.Field
	SecurityAlerts    apijson.Field
	TransactionAlerts apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *NotificationSettingsEventPreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationSettingsEventPreferencesJSON) RawJSON() string {
	return r.raw
}

// Settings for notification quiet hours.
type NotificationSettingsQuietHours struct {
	// If true, notifications are suppressed during specified quiet hours.
	Enabled interface{} `json:"enabled"`
	// End time for quiet hours (HH:MM format).
	EndTime interface{} `json:"endTime"`
	// Start time for quiet hours (HH:MM format).
	StartTime interface{}                        `json:"startTime"`
	JSON      notificationSettingsQuietHoursJSON `json:"-"`
}

// notificationSettingsQuietHoursJSON contains the JSON metadata for the struct
// [NotificationSettingsQuietHours]
type notificationSettingsQuietHoursJSON struct {
	Enabled     apijson.Field
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationSettingsQuietHours) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationSettingsQuietHoursJSON) RawJSON() string {
	return r.raw
}

type NotificationSettingUpdateParams struct {
	// Updated preferences for notification delivery channels. Only provided fields are
	// updated.
	ChannelPreferences param.Field[NotificationSettingUpdateParamsChannelPreferences] `json:"channelPreferences"`
	// Updated preferences for different types of events. Only provided fields are
	// updated.
	EventPreferences param.Field[NotificationSettingUpdateParamsEventPreferences] `json:"eventPreferences"`
	// Updated settings for notification quiet hours. Only provided fields are updated.
	QuietHours param.Field[NotificationSettingUpdateParamsQuietHours] `json:"quietHours"`
}

func (r NotificationSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated preferences for notification delivery channels. Only provided fields are
// updated.
type NotificationSettingUpdateParamsChannelPreferences struct {
	Email param.Field[interface{}] `json:"email"`
	InApp param.Field[interface{}] `json:"inApp"`
	Push  param.Field[interface{}] `json:"push"`
	SMS   param.Field[interface{}] `json:"sms"`
}

func (r NotificationSettingUpdateParamsChannelPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated preferences for different types of events. Only provided fields are
// updated.
type NotificationSettingUpdateParamsEventPreferences struct {
	AIInsights        param.Field[interface{}] `json:"aiInsights"`
	BudgetAlerts      param.Field[interface{}] `json:"budgetAlerts"`
	PromotionalOffers param.Field[interface{}] `json:"promotionalOffers"`
	SecurityAlerts    param.Field[interface{}] `json:"securityAlerts"`
	TransactionAlerts param.Field[interface{}] `json:"transactionAlerts"`
}

func (r NotificationSettingUpdateParamsEventPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updated settings for notification quiet hours. Only provided fields are updated.
type NotificationSettingUpdateParamsQuietHours struct {
	Enabled   param.Field[interface{}] `json:"enabled"`
	EndTime   param.Field[interface{}] `json:"endTime"`
	StartTime param.Field[interface{}] `json:"startTime"`
}

func (r NotificationSettingUpdateParamsQuietHours) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
