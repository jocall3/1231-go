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

// DeveloperAPIKeyService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeveloperAPIKeyService] method instead.
type DeveloperAPIKeyService struct {
	Options []option.RequestOption
}

// NewDeveloperAPIKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeveloperAPIKeyService(opts ...option.RequestOption) (r *DeveloperAPIKeyService) {
	r = &DeveloperAPIKeyService{}
	r.Options = opts
	return
}

// Generates a new API key for the developer application with specified scopes and
// an optional expiration.
func (r *DeveloperAPIKeyService) New(ctx context.Context, body DeveloperAPIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "developers/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a list of API keys issued to the authenticated developer application.
func (r *DeveloperAPIKeyService) List(ctx context.Context, query DeveloperAPIKeyListParams, opts ...option.RequestOption) (res *DeveloperAPIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "developers/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Revokes an existing API key, disabling its access immediately.
func (r *DeveloperAPIKeyService) Revoke(ctx context.Context, keyID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("developers/api-keys/%v", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type APIKey struct {
	// Unique identifier for the API key.
	ID interface{} `json:"id,required"`
	// Timestamp when the API key was created.
	CreatedAt interface{} `json:"createdAt,required"`
	// The non-secret prefix of the API key, used for identification.
	Prefix interface{} `json:"prefix,required"`
	// List of permissions granted to this API key.
	Scopes []interface{} `json:"scopes,required"`
	// Current status of the API key.
	Status APIKeyStatus `json:"status,required"`
	// Timestamp when the API key will expire, if set.
	ExpiresAt interface{} `json:"expiresAt"`
	// Timestamp of the last time this API key was used.
	LastUsed interface{} `json:"lastUsed"`
	JSON     apiKeyJSON  `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Prefix      apijson.Field
	Scopes      apijson.Field
	Status      apijson.Field
	ExpiresAt   apijson.Field
	LastUsed    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

// Current status of the API key.
type APIKeyStatus string

const (
	APIKeyStatusActive  APIKeyStatus = "active"
	APIKeyStatusRevoked APIKeyStatus = "revoked"
	APIKeyStatusExpired APIKeyStatus = "expired"
)

func (r APIKeyStatus) IsKnown() bool {
	switch r {
	case APIKeyStatusActive, APIKeyStatusRevoked, APIKeyStatusExpired:
		return true
	}
	return false
}

type DeveloperAPIKeyListResponse struct {
	Data []APIKey                        `json:"data"`
	JSON developerAPIKeyListResponseJSON `json:"-"`
	PaginatedList
}

// developerAPIKeyListResponseJSON contains the JSON metadata for the struct
// [DeveloperAPIKeyListResponse]
type developerAPIKeyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeveloperAPIKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r developerAPIKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type DeveloperAPIKeyNewParams struct {
	// A descriptive name for the API key.
	Name param.Field[interface{}] `json:"name,required"`
	// List of permissions to grant to this API key.
	Scopes param.Field[[]interface{}] `json:"scopes,required"`
	// Optional: Number of days until the API key expires. If omitted, it will not
	// expire.
	ExpiresInDays param.Field[interface{}] `json:"expiresInDays"`
}

func (r DeveloperAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeveloperAPIKeyListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [DeveloperAPIKeyListParams]'s query parameters as
// `url.Values`.
func (r DeveloperAPIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
