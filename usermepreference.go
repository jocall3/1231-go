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

// UserMePreferenceService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMePreferenceService] method instead.
type UserMePreferenceService struct {
	Options []option.RequestOption
}

// NewUserMePreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserMePreferenceService(opts ...option.RequestOption) (r *UserMePreferenceService) {
	r = &UserMePreferenceService{}
	r.Options = opts
	return
}

// Retrieves the user's deep personalization preferences, including AI
// customization settings, notification channel priorities, thematic choices, and
// data sharing consents.
func (r *UserMePreferenceService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserPreferences, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the user's deep personalization preferences, allowing dynamic control
// over AI behavior, notification delivery, thematic choices, and data privacy
// settings.
func (r *UserMePreferenceService) Update(ctx context.Context, body UserMePreferenceUpdateParams, opts ...option.RequestOption) (res *UserPreferences, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// User's personalized preferences for the platform.
type UserPreferences struct {
	// How the user prefers to interact with AI (proactive advice, balanced, or only on
	// demand).
	AIInteractionMode UserPreferencesAIInteractionMode `json:"aiInteractionMode"`
	// Consent status for sharing anonymized data for AI improvement and personalized
	// offers.
	DataSharingConsent interface{} `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels UserPreferencesNotificationChannels `json:"notificationChannels"`
	// Preferred language for the user interface.
	PreferredLanguage interface{} `json:"preferredLanguage"`
	// Preferred UI theme (e.g., Light-Default, Dark-Quantum).
	Theme interface{} `json:"theme"`
	// Default grouping preference for transaction lists.
	TransactionGrouping UserPreferencesTransactionGrouping `json:"transactionGrouping"`
	JSON                userPreferencesJSON                `json:"-"`
}

// userPreferencesJSON contains the JSON metadata for the struct [UserPreferences]
type userPreferencesJSON struct {
	AIInteractionMode    apijson.Field
	DataSharingConsent   apijson.Field
	NotificationChannels apijson.Field
	PreferredLanguage    apijson.Field
	Theme                apijson.Field
	TransactionGrouping  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserPreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferencesJSON) RawJSON() string {
	return r.raw
}

// How the user prefers to interact with AI (proactive advice, balanced, or only on
// demand).
type UserPreferencesAIInteractionMode string

const (
	UserPreferencesAIInteractionModeProactive UserPreferencesAIInteractionMode = "proactive"
	UserPreferencesAIInteractionModeBalanced  UserPreferencesAIInteractionMode = "balanced"
	UserPreferencesAIInteractionModeOnDemand  UserPreferencesAIInteractionMode = "on_demand"
)

func (r UserPreferencesAIInteractionMode) IsKnown() bool {
	switch r {
	case UserPreferencesAIInteractionModeProactive, UserPreferencesAIInteractionModeBalanced, UserPreferencesAIInteractionModeOnDemand:
		return true
	}
	return false
}

// Preferred channels for receiving notifications.
type UserPreferencesNotificationChannels struct {
	Email interface{}                             `json:"email"`
	InApp interface{}                             `json:"inApp"`
	Push  interface{}                             `json:"push"`
	SMS   interface{}                             `json:"sms"`
	JSON  userPreferencesNotificationChannelsJSON `json:"-"`
}

// userPreferencesNotificationChannelsJSON contains the JSON metadata for the
// struct [UserPreferencesNotificationChannels]
type userPreferencesNotificationChannelsJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPreferencesNotificationChannels) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferencesNotificationChannelsJSON) RawJSON() string {
	return r.raw
}

// Default grouping preference for transaction lists.
type UserPreferencesTransactionGrouping string

const (
	UserPreferencesTransactionGroupingCategory UserPreferencesTransactionGrouping = "category"
	UserPreferencesTransactionGroupingMerchant UserPreferencesTransactionGrouping = "merchant"
	UserPreferencesTransactionGroupingDate     UserPreferencesTransactionGrouping = "date"
	UserPreferencesTransactionGroupingAccount  UserPreferencesTransactionGrouping = "account"
)

func (r UserPreferencesTransactionGrouping) IsKnown() bool {
	switch r {
	case UserPreferencesTransactionGroupingCategory, UserPreferencesTransactionGroupingMerchant, UserPreferencesTransactionGroupingDate, UserPreferencesTransactionGroupingAccount:
		return true
	}
	return false
}

// User's personalized preferences for the platform.
type UserPreferencesParam struct {
	// How the user prefers to interact with AI (proactive advice, balanced, or only on
	// demand).
	AIInteractionMode param.Field[UserPreferencesAIInteractionMode] `json:"aiInteractionMode"`
	// Consent status for sharing anonymized data for AI improvement and personalized
	// offers.
	DataSharingConsent param.Field[interface{}] `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels param.Field[UserPreferencesNotificationChannelsParam] `json:"notificationChannels"`
	// Preferred language for the user interface.
	PreferredLanguage param.Field[interface{}] `json:"preferredLanguage"`
	// Preferred UI theme (e.g., Light-Default, Dark-Quantum).
	Theme param.Field[interface{}] `json:"theme"`
	// Default grouping preference for transaction lists.
	TransactionGrouping param.Field[UserPreferencesTransactionGrouping] `json:"transactionGrouping"`
}

func (r UserPreferencesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Preferred channels for receiving notifications.
type UserPreferencesNotificationChannelsParam struct {
	Email param.Field[interface{}] `json:"email"`
	InApp param.Field[interface{}] `json:"inApp"`
	Push  param.Field[interface{}] `json:"push"`
	SMS   param.Field[interface{}] `json:"sms"`
}

func (r UserPreferencesNotificationChannelsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMePreferenceUpdateParams struct {
	// User's personalized preferences for the platform.
	UserPreferences UserPreferencesParam `json:"user_preferences,required"`
}

func (r UserMePreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UserPreferences)
}
