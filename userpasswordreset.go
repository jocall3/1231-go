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

// UserPasswordResetService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPasswordResetService] method instead.
type UserPasswordResetService struct {
	Options []option.RequestOption
}

// NewUserPasswordResetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserPasswordResetService(opts ...option.RequestOption) (r *UserPasswordResetService) {
	r = &UserPasswordResetService{}
	r.Options = opts
	return
}

// Confirms the password reset using the received verification code and sets a new
// password.
func (r *UserPasswordResetService) Confirm(ctx context.Context, body UserPasswordResetConfirmParams, opts ...option.RequestOption) (res *UserPasswordResetConfirmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/password-reset/confirm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Starts the password reset flow by sending a verification code or link to the
// user's registered email or phone.
func (r *UserPasswordResetService) Initiate(ctx context.Context, body UserPasswordResetInitiateParams, opts ...option.RequestOption) (res *UserPasswordResetInitiateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/password-reset/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserPasswordResetConfirmResponse struct {
	Message interface{}                          `json:"message"`
	JSON    userPasswordResetConfirmResponseJSON `json:"-"`
}

// userPasswordResetConfirmResponseJSON contains the JSON metadata for the struct
// [UserPasswordResetConfirmResponse]
type userPasswordResetConfirmResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPasswordResetConfirmResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPasswordResetConfirmResponseJSON) RawJSON() string {
	return r.raw
}

type UserPasswordResetInitiateResponse struct {
	Message interface{}                           `json:"message"`
	JSON    userPasswordResetInitiateResponseJSON `json:"-"`
}

// userPasswordResetInitiateResponseJSON contains the JSON metadata for the struct
// [UserPasswordResetInitiateResponse]
type userPasswordResetInitiateResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPasswordResetInitiateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPasswordResetInitiateResponseJSON) RawJSON() string {
	return r.raw
}

type UserPasswordResetConfirmParams struct {
	// User's email or phone number used for verification.
	Identifier param.Field[interface{}] `json:"identifier,required"`
	// The new password for the user account.
	NewPassword param.Field[interface{}] `json:"newPassword,required"`
	// The verification code received via email or SMS.
	VerificationCode param.Field[interface{}] `json:"verificationCode,required"`
}

func (r UserPasswordResetConfirmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserPasswordResetInitiateParams struct {
	// User's email or phone number for verification.
	Identifier param.Field[interface{}] `json:"identifier,required"`
}

func (r UserPasswordResetInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
