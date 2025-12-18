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

// UserMeService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeService] method instead.
type UserMeService struct {
	Options     []option.RequestOption
	Preferences *UserMePreferenceService
	Devices     *UserMeDeviceService
	Biometrics  *UserMeBiometricService
}

// NewUserMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserMeService(opts ...option.RequestOption) (r *UserMeService) {
	r = &UserMeService{}
	r.Options = opts
	r.Preferences = NewUserMePreferenceService(opts...)
	r.Devices = NewUserMeDeviceService(opts...)
	r.Biometrics = NewUserMeBiometricService(opts...)
	return
}

// Fetches the complete and dynamically updated profile information for the
// currently authenticated user, encompassing personal details, security status,
// gamification level, loyalty points, and linked identity attributes.
func (r *UserMeService) Get(ctx context.Context, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates selected fields of the currently authenticated user's profile
// information.
func (r *UserMeService) Update(ctx context.Context, body UserMeUpdateParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UserMeUpdateParams struct {
	Address param.Field[AddressParam] `json:"address"`
	// Updated full name of the user.
	Name param.Field[interface{}] `json:"name"`
	// Updated primary phone number of the user.
	Phone param.Field[interface{}] `json:"phone"`
	// User's personalized preferences for the platform.
	Preferences param.Field[UserPreferencesParam] `json:"preferences"`
}

func (r UserMeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
