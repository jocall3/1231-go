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

// UserService contains methods and other services that help with interacting with
// the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	PasswordReset *UserPasswordResetService
	Me            *UserMeService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.PasswordReset = NewUserPasswordResetService(opts...)
	r.Me = NewUserMeService(opts...)
	return
}

// Authenticates a user and creates a secure session, returning access tokens. May
// require MFA depending on user settings.
func (r *UserService) Login(ctx context.Context, body UserLoginParams, opts ...option.RequestOption) (res *UserLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Registers a new user account with , initiating the onboarding process. Requires
// basic user details.
func (r *UserService) Register(ctx context.Context, body UserRegisterParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Address struct {
	City    interface{} `json:"city"`
	Country interface{} `json:"country"`
	State   interface{} `json:"state"`
	Street  interface{} `json:"street"`
	Zip     interface{} `json:"zip"`
	JSON    addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressParam struct {
	City    param.Field[interface{}] `json:"city"`
	Country param.Field[interface{}] `json:"country"`
	State   param.Field[interface{}] `json:"state"`
	Street  param.Field[interface{}] `json:"street"`
	Zip     param.Field[interface{}] `json:"zip"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type User struct {
	// Unique identifier for the user.
	ID interface{} `json:"id,required"`
	// Primary email address of the user.
	Email interface{} `json:"email,required"`
	// Indicates if the user's identity has been verified (e.g., via KYC).
	IdentityVerified interface{} `json:"identityVerified,required"`
	// Full name of the user.
	Name    interface{} `json:"name,required"`
	Address Address     `json:"address"`
	// AI-identified financial persona for tailored advice.
	AIPersona interface{} `json:"aiPersona"`
	// Date of birth of the user (YYYY-MM-DD).
	DateOfBirth interface{} `json:"dateOfBirth"`
	// Current gamification level.
	GamificationLevel interface{} `json:"gamificationLevel"`
	// Current balance of loyalty points.
	LoyaltyPoints interface{} `json:"loyaltyPoints"`
	// Current loyalty program tier.
	LoyaltyTier interface{} `json:"loyaltyTier"`
	// Primary phone number of the user.
	Phone interface{} `json:"phone"`
	// User's personalized preferences for the platform.
	Preferences UserPreferences `json:"preferences"`
	// Security-related status for the user account.
	SecurityStatus UserSecurityStatus `json:"securityStatus"`
	JSON           userJSON           `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID                apijson.Field
	Email             apijson.Field
	IdentityVerified  apijson.Field
	Name              apijson.Field
	Address           apijson.Field
	AIPersona         apijson.Field
	DateOfBirth       apijson.Field
	GamificationLevel apijson.Field
	LoyaltyPoints     apijson.Field
	LoyaltyTier       apijson.Field
	Phone             apijson.Field
	Preferences       apijson.Field
	SecurityStatus    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

// Security-related status for the user account.
type UserSecurityStatus struct {
	// Indicates if biometric authentication is enrolled.
	BiometricsEnrolled interface{} `json:"biometricsEnrolled"`
	// Timestamp of the last successful login.
	LastLogin interface{} `json:"lastLogin"`
	// IP address of the last successful login.
	LastLoginIP interface{} `json:"lastLoginIp"`
	// Indicates if two-factor authentication (2FA) is enabled.
	TwoFactorEnabled interface{}            `json:"twoFactorEnabled"`
	JSON             userSecurityStatusJSON `json:"-"`
}

// userSecurityStatusJSON contains the JSON metadata for the struct
// [UserSecurityStatus]
type userSecurityStatusJSON struct {
	BiometricsEnrolled apijson.Field
	LastLogin          apijson.Field
	LastLoginIP        apijson.Field
	TwoFactorEnabled   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserLoginResponse struct {
	// JWT access token to authenticate subsequent API requests.
	AccessToken interface{} `json:"accessToken,required"`
	// Lifetime of the access token in seconds.
	ExpiresIn interface{} `json:"expiresIn,required"`
	// Token used to obtain new access tokens without re-authenticating.
	RefreshToken interface{} `json:"refreshToken,required"`
	// Type of the access token.
	TokenType interface{}           `json:"tokenType,required"`
	JSON      userLoginResponseJSON `json:"-"`
}

// userLoginResponseJSON contains the JSON metadata for the struct
// [UserLoginResponse]
type userLoginResponseJSON struct {
	AccessToken  apijson.Field
	ExpiresIn    apijson.Field
	RefreshToken apijson.Field
	TokenType    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserLoginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoginResponseJSON) RawJSON() string {
	return r.raw
}

type UserLoginParams struct {
	// User's email address.
	Email param.Field[interface{}] `json:"email,required"`
	// User's password.
	Password param.Field[interface{}] `json:"password,required"`
	// Optional: Multi-factor authentication code, if required.
	MfaCode param.Field[interface{}] `json:"mfaCode"`
}

func (r UserLoginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRegisterParams struct {
	// Email address for registration and login.
	Email param.Field[interface{}] `json:"email,required"`
	// Full name of the user.
	Name param.Field[interface{}] `json:"name,required"`
	// User's chosen password.
	Password param.Field[interface{}]  `json:"password,required"`
	Address  param.Field[AddressParam] `json:"address"`
	// Optional date of birth (YYYY-MM-DD).
	DateOfBirth param.Field[interface{}] `json:"dateOfBirth"`
	// Optional phone number for MFA or recovery.
	Phone param.Field[interface{}] `json:"phone"`
}

func (r UserRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
