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

// IdentityKYCService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityKYCService] method instead.
type IdentityKYCService struct {
	Options []option.RequestOption
}

// NewIdentityKYCService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityKYCService(opts ...option.RequestOption) (r *IdentityKYCService) {
	r = &IdentityKYCService{}
	r.Options = opts
	return
}

// Retrieves the current status of the user's Know Your Customer (KYC) verification
// process.
func (r *IdentityKYCService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *KYCStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "identity/kyc/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits Know Your Customer (KYC) documentation, such as identity proofs and
// address verification, for AI-accelerated compliance and identity verification,
// crucial for higher service tiers and regulatory adherence.
func (r *IdentityKYCService) Submit(ctx context.Context, body IdentityKYCSubmitParams, opts ...option.RequestOption) (res *KYCStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "identity/kyc/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type KYCStatus struct {
	// Timestamp of the last KYC document submission.
	LastSubmissionDate interface{} `json:"lastSubmissionDate,required"`
	// Overall status of the KYC verification process.
	OverallStatus KYCStatusOverallStatus `json:"overallStatus,required"`
	// List of actions required from the user if status is 'requires_more_info'.
	RequiredActions []interface{} `json:"requiredActions,required"`
	// The ID of the user whose KYC status is being retrieved.
	UserID interface{} `json:"userId,required"`
	// Reason for rejection if status is 'rejected'.
	RejectionReason interface{} `json:"rejectionReason"`
	// The KYC verification tier achieved (e.g., for different service levels).
	VerifiedTier KYCStatusVerifiedTier `json:"verifiedTier,nullable"`
	JSON         kycStatusJSON         `json:"-"`
}

// kycStatusJSON contains the JSON metadata for the struct [KYCStatus]
type kycStatusJSON struct {
	LastSubmissionDate apijson.Field
	OverallStatus      apijson.Field
	RequiredActions    apijson.Field
	UserID             apijson.Field
	RejectionReason    apijson.Field
	VerifiedTier       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *KYCStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kycStatusJSON) RawJSON() string {
	return r.raw
}

// Overall status of the KYC verification process.
type KYCStatusOverallStatus string

const (
	KYCStatusOverallStatusNotSubmitted     KYCStatusOverallStatus = "not_submitted"
	KYCStatusOverallStatusInReview         KYCStatusOverallStatus = "in_review"
	KYCStatusOverallStatusVerified         KYCStatusOverallStatus = "verified"
	KYCStatusOverallStatusRejected         KYCStatusOverallStatus = "rejected"
	KYCStatusOverallStatusRequiresMoreInfo KYCStatusOverallStatus = "requires_more_info"
)

func (r KYCStatusOverallStatus) IsKnown() bool {
	switch r {
	case KYCStatusOverallStatusNotSubmitted, KYCStatusOverallStatusInReview, KYCStatusOverallStatusVerified, KYCStatusOverallStatusRejected, KYCStatusOverallStatusRequiresMoreInfo:
		return true
	}
	return false
}

// The KYC verification tier achieved (e.g., for different service levels).
type KYCStatusVerifiedTier string

const (
	KYCStatusVerifiedTierBronze   KYCStatusVerifiedTier = "bronze"
	KYCStatusVerifiedTierSilver   KYCStatusVerifiedTier = "silver"
	KYCStatusVerifiedTierGold     KYCStatusVerifiedTier = "gold"
	KYCStatusVerifiedTierPlatinum KYCStatusVerifiedTier = "platinum"
)

func (r KYCStatusVerifiedTier) IsKnown() bool {
	switch r {
	case KYCStatusVerifiedTierBronze, KYCStatusVerifiedTierSilver, KYCStatusVerifiedTierGold, KYCStatusVerifiedTierPlatinum:
		return true
	}
	return false
}

type IdentityKYCSubmitParams struct {
	// The two-letter ISO country code where the document was issued.
	CountryOfIssue param.Field[interface{}] `json:"countryOfIssue,required"`
	// The identification number on the document.
	DocumentNumber param.Field[interface{}] `json:"documentNumber,required"`
	// The type of KYC document being submitted.
	DocumentType param.Field[IdentityKYCSubmitParamsDocumentType] `json:"documentType,required"`
	// The expiration date of the document (YYYY-MM-DD).
	ExpirationDate param.Field[interface{}] `json:"expirationDate,required"`
	// The issue date of the document (YYYY-MM-DD).
	IssueDate param.Field[interface{}] `json:"issueDate,required"`
	// Array of additional documents (e.g., utility bills) as base64 encoded images.
	AdditionalDocuments param.Field[[]interface{}] `json:"additionalDocuments"`
	// Base64 encoded image of the back of the document (if applicable).
	DocumentBackImage param.Field[interface{}] `json:"documentBackImage"`
	// Base64 encoded image of the front of the document. Use 'application/json' with
	// base64 string, or 'multipart/form-data' for direct file upload.
	DocumentFrontImage param.Field[interface{}] `json:"documentFrontImage"`
}

func (r IdentityKYCSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of KYC document being submitted.
type IdentityKYCSubmitParamsDocumentType string

const (
	IdentityKYCSubmitParamsDocumentTypeDriversLicense IdentityKYCSubmitParamsDocumentType = "drivers_license"
	IdentityKYCSubmitParamsDocumentTypePassport       IdentityKYCSubmitParamsDocumentType = "passport"
	IdentityKYCSubmitParamsDocumentTypeNationalID     IdentityKYCSubmitParamsDocumentType = "national_id"
	IdentityKYCSubmitParamsDocumentTypeUtilityBill    IdentityKYCSubmitParamsDocumentType = "utility_bill"
	IdentityKYCSubmitParamsDocumentTypeBankStatement  IdentityKYCSubmitParamsDocumentType = "bank_statement"
	IdentityKYCSubmitParamsDocumentTypeOther          IdentityKYCSubmitParamsDocumentType = "other"
)

func (r IdentityKYCSubmitParamsDocumentType) IsKnown() bool {
	switch r {
	case IdentityKYCSubmitParamsDocumentTypeDriversLicense, IdentityKYCSubmitParamsDocumentTypePassport, IdentityKYCSubmitParamsDocumentTypeNationalID, IdentityKYCSubmitParamsDocumentTypeUtilityBill, IdentityKYCSubmitParamsDocumentTypeBankStatement, IdentityKYCSubmitParamsDocumentTypeOther:
		return true
	}
	return false
}
