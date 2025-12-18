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

// UserMeBiometricService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeBiometricService] method instead.
type UserMeBiometricService struct {
	Options []option.RequestOption
}

// NewUserMeBiometricService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeBiometricService(opts ...option.RequestOption) (r *UserMeBiometricService) {
	r = &UserMeBiometricService{}
	r.Options = opts
	return
}

// Removes all enrolled biometric data associated with the user's account for
// security reasons.
func (r *UserMeBiometricService) Deregister(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Initiates the enrollment process for biometric authentication (e.g.,
// fingerprint, facial scan) to enable secure and convenient access to sensitive
// features. Requires a biometric signature for initial proof.
func (r *UserMeBiometricService) Enroll(ctx context.Context, body UserMeBiometricEnrollParams, opts ...option.RequestOption) (res *BiometricStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics/enroll"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the current status of biometric enrollments for the authenticated
// user.
func (r *UserMeBiometricService) Status(ctx context.Context, opts ...option.RequestOption) (res *BiometricStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Performs real-time biometric verification to authorize sensitive actions or
// access protected resources, using a one-time biometric signature.
func (r *UserMeBiometricService) Verify(ctx context.Context, body UserMeBiometricVerifyParams, opts ...option.RequestOption) (res *UserMeBiometricVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Current biometric enrollment status for a user.
type BiometricStatus struct {
	// Overall status indicating if any biometrics are enrolled.
	BiometricsEnrolled interface{} `json:"biometricsEnrolled,required"`
	// List of specific biometric types and devices enrolled.
	EnrolledBiometrics []BiometricStatusEnrolledBiometric `json:"enrolledBiometrics,required"`
	// Timestamp of the last successful biometric authentication.
	LastUsed interface{}         `json:"lastUsed"`
	JSON     biometricStatusJSON `json:"-"`
}

// biometricStatusJSON contains the JSON metadata for the struct [BiometricStatus]
type biometricStatusJSON struct {
	BiometricsEnrolled apijson.Field
	EnrolledBiometrics apijson.Field
	LastUsed           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BiometricStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r biometricStatusJSON) RawJSON() string {
	return r.raw
}

type BiometricStatusEnrolledBiometric struct {
	DeviceID       interface{}                           `json:"deviceId"`
	EnrollmentDate interface{}                           `json:"enrollmentDate"`
	Type           BiometricStatusEnrolledBiometricsType `json:"type"`
	JSON           biometricStatusEnrolledBiometricJSON  `json:"-"`
}

// biometricStatusEnrolledBiometricJSON contains the JSON metadata for the struct
// [BiometricStatusEnrolledBiometric]
type biometricStatusEnrolledBiometricJSON struct {
	DeviceID       apijson.Field
	EnrollmentDate apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BiometricStatusEnrolledBiometric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r biometricStatusEnrolledBiometricJSON) RawJSON() string {
	return r.raw
}

type BiometricStatusEnrolledBiometricsType string

const (
	BiometricStatusEnrolledBiometricsTypeFingerprint       BiometricStatusEnrolledBiometricsType = "fingerprint"
	BiometricStatusEnrolledBiometricsTypeFacialRecognition BiometricStatusEnrolledBiometricsType = "facial_recognition"
	BiometricStatusEnrolledBiometricsTypeVoiceRecognition  BiometricStatusEnrolledBiometricsType = "voice_recognition"
)

func (r BiometricStatusEnrolledBiometricsType) IsKnown() bool {
	switch r {
	case BiometricStatusEnrolledBiometricsTypeFingerprint, BiometricStatusEnrolledBiometricsTypeFacialRecognition, BiometricStatusEnrolledBiometricsTypeVoiceRecognition:
		return true
	}
	return false
}

type UserMeBiometricVerifyResponse struct {
	// A descriptive message for the verification result.
	Message interface{} `json:"message"`
	// Status of the biometric verification.
	VerificationStatus UserMeBiometricVerifyResponseVerificationStatus `json:"verificationStatus"`
	JSON               userMeBiometricVerifyResponseJSON               `json:"-"`
}

// userMeBiometricVerifyResponseJSON contains the JSON metadata for the struct
// [UserMeBiometricVerifyResponse]
type userMeBiometricVerifyResponseJSON struct {
	Message            apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMeBiometricVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeBiometricVerifyResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the biometric verification.
type UserMeBiometricVerifyResponseVerificationStatus string

const (
	UserMeBiometricVerifyResponseVerificationStatusSuccess UserMeBiometricVerifyResponseVerificationStatus = "success"
	UserMeBiometricVerifyResponseVerificationStatusFailed  UserMeBiometricVerifyResponseVerificationStatus = "failed"
)

func (r UserMeBiometricVerifyResponseVerificationStatus) IsKnown() bool {
	switch r {
	case UserMeBiometricVerifyResponseVerificationStatusSuccess, UserMeBiometricVerifyResponseVerificationStatusFailed:
		return true
	}
	return false
}

type UserMeBiometricEnrollParams struct {
	// Base64 encoded representation of the biometric template or proof.
	BiometricSignature param.Field[interface{}] `json:"biometricSignature,required"`
	// The type of biometric data being enrolled.
	BiometricType param.Field[UserMeBiometricEnrollParamsBiometricType] `json:"biometricType,required"`
	// The ID of the device on which the biometric is being enrolled.
	DeviceID param.Field[interface{}] `json:"deviceId,required"`
	// Optional: A friendly name for the device, if not already linked.
	DeviceName param.Field[interface{}] `json:"deviceName"`
}

func (r UserMeBiometricEnrollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of biometric data being enrolled.
type UserMeBiometricEnrollParamsBiometricType string

const (
	UserMeBiometricEnrollParamsBiometricTypeFingerprint       UserMeBiometricEnrollParamsBiometricType = "fingerprint"
	UserMeBiometricEnrollParamsBiometricTypeFacialRecognition UserMeBiometricEnrollParamsBiometricType = "facial_recognition"
	UserMeBiometricEnrollParamsBiometricTypeVoiceRecognition  UserMeBiometricEnrollParamsBiometricType = "voice_recognition"
)

func (r UserMeBiometricEnrollParamsBiometricType) IsKnown() bool {
	switch r {
	case UserMeBiometricEnrollParamsBiometricTypeFingerprint, UserMeBiometricEnrollParamsBiometricTypeFacialRecognition, UserMeBiometricEnrollParamsBiometricTypeVoiceRecognition:
		return true
	}
	return false
}

type UserMeBiometricVerifyParams struct {
	// Base64 encoded representation of the one-time biometric proof for verification.
	BiometricSignature param.Field[interface{}] `json:"biometricSignature,required"`
	// The type of biometric data being verified.
	BiometricType param.Field[UserMeBiometricVerifyParamsBiometricType] `json:"biometricType,required"`
	// The ID of the device initiating the biometric verification.
	DeviceID param.Field[interface{}] `json:"deviceId,required"`
}

func (r UserMeBiometricVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of biometric data being verified.
type UserMeBiometricVerifyParamsBiometricType string

const (
	UserMeBiometricVerifyParamsBiometricTypeFingerprint       UserMeBiometricVerifyParamsBiometricType = "fingerprint"
	UserMeBiometricVerifyParamsBiometricTypeFacialRecognition UserMeBiometricVerifyParamsBiometricType = "facial_recognition"
	UserMeBiometricVerifyParamsBiometricTypeVoiceRecognition  UserMeBiometricVerifyParamsBiometricType = "voice_recognition"
)

func (r UserMeBiometricVerifyParamsBiometricType) IsKnown() bool {
	switch r {
	case UserMeBiometricVerifyParamsBiometricTypeFingerprint, UserMeBiometricVerifyParamsBiometricTypeFacialRecognition, UserMeBiometricVerifyParamsBiometricTypeVoiceRecognition:
		return true
	}
	return false
}
