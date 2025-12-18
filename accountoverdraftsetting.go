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

// AccountOverdraftSettingService contains methods and other services that help
// with interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountOverdraftSettingService] method instead.
type AccountOverdraftSettingService struct {
	Options []option.RequestOption
}

// NewAccountOverdraftSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountOverdraftSettingService(opts ...option.RequestOption) (r *AccountOverdraftSettingService) {
	r = &AccountOverdraftSettingService{}
	r.Options = opts
	return
}

// Retrieves the current overdraft protection settings for a specific account.
func (r *AccountOverdraftSettingService) GetOverdraftSettings(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *OverdraftSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("accounts/%v/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the overdraft protection settings for a specific account, enabling or
// disabling protection and configuring preferences.
func (r *AccountOverdraftSettingService) UpdateOverdraftSettings(ctx context.Context, accountID interface{}, body AccountOverdraftSettingUpdateOverdraftSettingsParams, opts ...option.RequestOption) (res *OverdraftSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("accounts/%v/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OverdraftSettings struct {
	// The account ID these overdraft settings apply to.
	AccountID interface{} `json:"accountId,required"`
	// If true, overdraft protection is enabled.
	Enabled interface{} `json:"enabled,required"`
	// User's preference for how overdraft fees are handled or if transactions should
	// be declined.
	FeePreference OverdraftSettingsFeePreference `json:"feePreference,required"`
	// The ID of the linked savings account, if `linkToSavings` is true.
	LinkedSavingsAccountID interface{} `json:"linkedSavingsAccountId"`
	// If true, attempts to draw funds from a linked savings account.
	LinkToSavings interface{} `json:"linkToSavings"`
	// The maximum amount that can be covered by overdraft protection.
	ProtectionLimit interface{}           `json:"protectionLimit"`
	JSON            overdraftSettingsJSON `json:"-"`
}

// overdraftSettingsJSON contains the JSON metadata for the struct
// [OverdraftSettings]
type overdraftSettingsJSON struct {
	AccountID              apijson.Field
	Enabled                apijson.Field
	FeePreference          apijson.Field
	LinkedSavingsAccountID apijson.Field
	LinkToSavings          apijson.Field
	ProtectionLimit        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *OverdraftSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overdraftSettingsJSON) RawJSON() string {
	return r.raw
}

// User's preference for how overdraft fees are handled or if transactions should
// be declined.
type OverdraftSettingsFeePreference string

const (
	OverdraftSettingsFeePreferenceAlwaysPay          OverdraftSettingsFeePreference = "always_pay"
	OverdraftSettingsFeePreferenceDeclineIfOverLimit OverdraftSettingsFeePreference = "decline_if_over_limit"
	OverdraftSettingsFeePreferenceAskMeFirst         OverdraftSettingsFeePreference = "ask_me_first"
)

func (r OverdraftSettingsFeePreference) IsKnown() bool {
	switch r {
	case OverdraftSettingsFeePreferenceAlwaysPay, OverdraftSettingsFeePreferenceDeclineIfOverLimit, OverdraftSettingsFeePreferenceAskMeFirst:
		return true
	}
	return false
}

type AccountOverdraftSettingUpdateOverdraftSettingsParams struct {
	// Enable or disable overdraft protection.
	Enabled param.Field[interface{}] `json:"enabled"`
	// New preference for how overdraft fees are handled.
	FeePreference param.Field[AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference] `json:"feePreference"`
	// New ID of the linked savings account, if `linkToSavings` is true. Set to null to
	// unlink.
	LinkedSavingsAccountID param.Field[interface{}] `json:"linkedSavingsAccountId"`
	// Enable or disable linking to a savings account for overdraft coverage.
	LinkToSavings param.Field[interface{}] `json:"linkToSavings"`
	// New maximum amount for overdraft protection. Set to null to remove limit.
	ProtectionLimit param.Field[interface{}] `json:"protectionLimit"`
}

func (r AccountOverdraftSettingUpdateOverdraftSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New preference for how overdraft fees are handled.
type AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference string

const (
	AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceAlwaysPay          AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference = "always_pay"
	AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceDeclineIfOverLimit AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference = "decline_if_over_limit"
	AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceAskMeFirst         AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference = "ask_me_first"
)

func (r AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference) IsKnown() bool {
	switch r {
	case AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceAlwaysPay, AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceDeclineIfOverLimit, AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceAskMeFirst:
		return true
	}
	return false
}
