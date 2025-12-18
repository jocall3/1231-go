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

// UserMeDeviceService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeDeviceService] method instead.
type UserMeDeviceService struct {
	Options []option.RequestOption
}

// NewUserMeDeviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeDeviceService(opts ...option.RequestOption) (r *UserMeDeviceService) {
	r = &UserMeDeviceService{}
	r.Options = opts
	return
}

// Retrieves a list of all devices linked to the user's account, including mobile
// phones, tablets, and desktops, indicating their last active status and security
// posture.
func (r *UserMeDeviceService) List(ctx context.Context, query UserMeDeviceListParams, opts ...option.RequestOption) (res *UserMeDeviceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes a specific device from the user's linked devices, revoking its access
// and requiring re-registration for future use. Useful for lost or compromised
// devices.
func (r *UserMeDeviceService) Deregister(ctx context.Context, deviceID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("users/me/devices/%v", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Registers a new device for secure access and multi-factor authentication,
// associating it with the user's profile. This typically initiates a biometric or
// MFA enrollment flow.
func (r *UserMeDeviceService) Register(ctx context.Context, body UserMeDeviceRegisterParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Information about a connected device.
type Device struct {
	// Unique identifier for the device.
	ID interface{} `json:"id,required"`
	// Last known IP address of the device.
	IPAddress interface{} `json:"ipAddress,required"`
	// Timestamp of the last activity from this device.
	LastActive interface{} `json:"lastActive,required"`
	// Model of the device.
	Model interface{} `json:"model,required"`
	// Operating system of the device.
	Os interface{} `json:"os,required"`
	// Security trust level of the device.
	TrustLevel DeviceTrustLevel `json:"trustLevel,required"`
	// Type of the device.
	Type DeviceType `json:"type,required"`
	// User-assigned name for the device.
	DeviceName interface{} `json:"deviceName"`
	// Push notification token for the device.
	PushToken interface{} `json:"pushToken"`
	JSON      deviceJSON  `json:"-"`
}

// deviceJSON contains the JSON metadata for the struct [Device]
type deviceJSON struct {
	ID          apijson.Field
	IPAddress   apijson.Field
	LastActive  apijson.Field
	Model       apijson.Field
	Os          apijson.Field
	TrustLevel  apijson.Field
	Type        apijson.Field
	DeviceName  apijson.Field
	PushToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Device) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceJSON) RawJSON() string {
	return r.raw
}

// Security trust level of the device.
type DeviceTrustLevel string

const (
	DeviceTrustLevelTrusted             DeviceTrustLevel = "trusted"
	DeviceTrustLevelPendingVerification DeviceTrustLevel = "pending_verification"
	DeviceTrustLevelUntrusted           DeviceTrustLevel = "untrusted"
	DeviceTrustLevelBlocked             DeviceTrustLevel = "blocked"
)

func (r DeviceTrustLevel) IsKnown() bool {
	switch r {
	case DeviceTrustLevelTrusted, DeviceTrustLevelPendingVerification, DeviceTrustLevelUntrusted, DeviceTrustLevelBlocked:
		return true
	}
	return false
}

// Type of the device.
type DeviceType string

const (
	DeviceTypeMobile     DeviceType = "mobile"
	DeviceTypeDesktop    DeviceType = "desktop"
	DeviceTypeTablet     DeviceType = "tablet"
	DeviceTypeSmartWatch DeviceType = "smart_watch"
)

func (r DeviceType) IsKnown() bool {
	switch r {
	case DeviceTypeMobile, DeviceTypeDesktop, DeviceTypeTablet, DeviceTypeSmartWatch:
		return true
	}
	return false
}

type PaginatedList struct {
	// The maximum number of items returned in the current page.
	Limit interface{} `json:"limit,required"`
	// The number of items skipped before the current page.
	Offset interface{} `json:"offset,required"`
	// The total number of items available across all pages.
	Total interface{} `json:"total,required"`
	// The offset for the next page of results, if available. Null if no more pages.
	NextOffset interface{}       `json:"nextOffset"`
	JSON       paginatedListJSON `json:"-"`
}

// paginatedListJSON contains the JSON metadata for the struct [PaginatedList]
type paginatedListJSON struct {
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedListJSON) RawJSON() string {
	return r.raw
}

type UserMeDeviceListResponse struct {
	Data []Device                     `json:"data"`
	JSON userMeDeviceListResponseJSON `json:"-"`
	PaginatedList
}

// userMeDeviceListResponseJSON contains the JSON metadata for the struct
// [UserMeDeviceListResponse]
type userMeDeviceListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeDeviceListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [UserMeDeviceListParams]'s query parameters as `url.Values`.
func (r UserMeDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserMeDeviceRegisterParams struct {
	// Type of the device being registered.
	DeviceType param.Field[UserMeDeviceRegisterParamsDeviceType] `json:"deviceType,required"`
	// Model of the device.
	Model param.Field[interface{}] `json:"model,required"`
	// Operating system of the device.
	Os param.Field[interface{}] `json:"os,required"`
	// Optional: Base64 encoded biometric signature for initial enrollment (e.g., for
	// Passkey registration).
	BiometricSignature param.Field[interface{}] `json:"biometricSignature"`
	// Optional: A friendly name for the device.
	DeviceName param.Field[interface{}] `json:"deviceName"`
	// Optional: Push notification token for the device.
	PushToken param.Field[interface{}] `json:"pushToken"`
}

func (r UserMeDeviceRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the device being registered.
type UserMeDeviceRegisterParamsDeviceType string

const (
	UserMeDeviceRegisterParamsDeviceTypeMobile     UserMeDeviceRegisterParamsDeviceType = "mobile"
	UserMeDeviceRegisterParamsDeviceTypeDesktop    UserMeDeviceRegisterParamsDeviceType = "desktop"
	UserMeDeviceRegisterParamsDeviceTypeTablet     UserMeDeviceRegisterParamsDeviceType = "tablet"
	UserMeDeviceRegisterParamsDeviceTypeSmartWatch UserMeDeviceRegisterParamsDeviceType = "smart_watch"
)

func (r UserMeDeviceRegisterParamsDeviceType) IsKnown() bool {
	switch r {
	case UserMeDeviceRegisterParamsDeviceTypeMobile, UserMeDeviceRegisterParamsDeviceTypeDesktop, UserMeDeviceRegisterParamsDeviceTypeTablet, UserMeDeviceRegisterParamsDeviceTypeSmartWatch:
		return true
	}
	return false
}
