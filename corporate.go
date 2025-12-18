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

// CorporateService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateService] method instead.
type CorporateService struct {
	Options    []option.RequestOption
	Cards      *CorporateCardService
	Anomalies  *CorporateAnomalyService
	Compliance *CorporateComplianceService
	Treasury   *CorporateTreasuryService
	Risk       *CorporateRiskService
}

// NewCorporateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateService(opts ...option.RequestOption) (r *CorporateService) {
	r = &CorporateService{}
	r.Options = opts
	r.Cards = NewCorporateCardService(opts...)
	r.Anomalies = NewCorporateAnomalyService(opts...)
	r.Compliance = NewCorporateComplianceService(opts...)
	r.Treasury = NewCorporateTreasuryService(opts...)
	r.Risk = NewCorporateRiskService(opts...)
	return
}

// Executes a real-time screening of an individual or entity against global
// sanction lists and watchlists.
func (r *CorporateService) PerformSanctionScreening(ctx context.Context, body CorporatePerformSanctionScreeningParams, opts ...option.RequestOption) (res *CorporatePerformSanctionScreeningResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/sanction-screening"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporatePerformSanctionScreeningResponse struct {
	// Details of any potential or exact matches found.
	MatchDetails []CorporatePerformSanctionScreeningResponseMatchDetail `json:"matchDetails,required"`
	// True if any potential matches were found on sanction lists.
	MatchFound interface{} `json:"matchFound,required"`
	// Unique identifier for this screening operation.
	ScreeningID interface{} `json:"screeningId,required"`
	// Timestamp when the screening was performed.
	ScreeningTimestamp interface{} `json:"screeningTimestamp,required"`
	// Overall status of the screening result.
	Status CorporatePerformSanctionScreeningResponseStatus `json:"status,required"`
	// An optional message providing more context on the status.
	Message interface{}                                   `json:"message"`
	JSON    corporatePerformSanctionScreeningResponseJSON `json:"-"`
}

// corporatePerformSanctionScreeningResponseJSON contains the JSON metadata for the
// struct [CorporatePerformSanctionScreeningResponse]
type corporatePerformSanctionScreeningResponseJSON struct {
	MatchDetails       apijson.Field
	MatchFound         apijson.Field
	ScreeningID        apijson.Field
	ScreeningTimestamp apijson.Field
	Status             apijson.Field
	Message            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CorporatePerformSanctionScreeningResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporatePerformSanctionScreeningResponseJSON) RawJSON() string {
	return r.raw
}

type CorporatePerformSanctionScreeningResponseMatchDetail struct {
	// Name of the sanction list where a match was found.
	ListName interface{} `json:"listName"`
	// The name on the sanction list that matched.
	MatchedName interface{} `json:"matchedName"`
	// Optional: URL to public record of the sanction list entry.
	PublicURL interface{} `json:"publicUrl"`
	// Reason for the match (e.g., exact name, alias, partial match).
	Reason interface{} `json:"reason"`
	// Match confidence score (0-1).
	Score interface{}                                              `json:"score"`
	JSON  corporatePerformSanctionScreeningResponseMatchDetailJSON `json:"-"`
}

// corporatePerformSanctionScreeningResponseMatchDetailJSON contains the JSON
// metadata for the struct [CorporatePerformSanctionScreeningResponseMatchDetail]
type corporatePerformSanctionScreeningResponseMatchDetailJSON struct {
	ListName    apijson.Field
	MatchedName apijson.Field
	PublicURL   apijson.Field
	Reason      apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporatePerformSanctionScreeningResponseMatchDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporatePerformSanctionScreeningResponseMatchDetailJSON) RawJSON() string {
	return r.raw
}

// Overall status of the screening result.
type CorporatePerformSanctionScreeningResponseStatus string

const (
	CorporatePerformSanctionScreeningResponseStatusClear          CorporatePerformSanctionScreeningResponseStatus = "clear"
	CorporatePerformSanctionScreeningResponseStatusPotentialMatch CorporatePerformSanctionScreeningResponseStatus = "potential_match"
	CorporatePerformSanctionScreeningResponseStatusConfirmedMatch CorporatePerformSanctionScreeningResponseStatus = "confirmed_match"
	CorporatePerformSanctionScreeningResponseStatusError          CorporatePerformSanctionScreeningResponseStatus = "error"
)

func (r CorporatePerformSanctionScreeningResponseStatus) IsKnown() bool {
	switch r {
	case CorporatePerformSanctionScreeningResponseStatusClear, CorporatePerformSanctionScreeningResponseStatusPotentialMatch, CorporatePerformSanctionScreeningResponseStatusConfirmedMatch, CorporatePerformSanctionScreeningResponseStatusError:
		return true
	}
	return false
}

type CorporatePerformSanctionScreeningParams struct {
	// Two-letter ISO country code related to the entity (e.g., country of residence,
	// registration).
	Country param.Field[interface{}] `json:"country,required"`
	// The type of entity being screened.
	EntityType param.Field[CorporatePerformSanctionScreeningParamsEntityType] `json:"entityType,required"`
	// Full name of the individual or organization to screen.
	Name    param.Field[interface{}]  `json:"name,required"`
	Address param.Field[AddressParam] `json:"address"`
	// Date of birth for individuals (YYYY-MM-DD).
	DateOfBirth param.Field[interface{}] `json:"dateOfBirth"`
	// Optional: Any government-issued identification number (e.g., passport, national
	// ID).
	IdentificationNumber param.Field[interface{}] `json:"identificationNumber"`
}

func (r CorporatePerformSanctionScreeningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of entity being screened.
type CorporatePerformSanctionScreeningParamsEntityType string

const (
	CorporatePerformSanctionScreeningParamsEntityTypeIndividual   CorporatePerformSanctionScreeningParamsEntityType = "individual"
	CorporatePerformSanctionScreeningParamsEntityTypeOrganization CorporatePerformSanctionScreeningParamsEntityType = "organization"
)

func (r CorporatePerformSanctionScreeningParamsEntityType) IsKnown() bool {
	switch r {
	case CorporatePerformSanctionScreeningParamsEntityTypeIndividual, CorporatePerformSanctionScreeningParamsEntityTypeOrganization:
		return true
	}
	return false
}
