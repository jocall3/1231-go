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

// MarketplaceOfferService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketplaceOfferService] method instead.
type MarketplaceOfferService struct {
	Options []option.RequestOption
}

// NewMarketplaceOfferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMarketplaceOfferService(opts ...option.RequestOption) (r *MarketplaceOfferService) {
	r = &MarketplaceOfferService{}
	r.Options = opts
	return
}

// Redeems a personalized, exclusive offer from the Plato AI marketplace, often
// resulting in a discount, special rate, or credit to the user's account.
func (r *MarketplaceOfferService) Redeem(ctx context.Context, offerID interface{}, body MarketplaceOfferRedeemParams, opts ...option.RequestOption) (res *MarketplaceOfferRedeemResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("marketplace/offers/%v/redeem", offerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MarketplaceOfferRedeemResponse struct {
	// If applicable, the ID of any associated transaction (e.g., a credit or initial
	// payment).
	AssociatedTransactionID interface{} `json:"associatedTransactionId"`
	// A descriptive message about the redemption.
	Message interface{} `json:"message"`
	// The ID of the redeemed offer.
	OfferID        interface{} `json:"offerId"`
	RedemptionDate interface{} `json:"redemptionDate"`
	// Unique ID for this redemption.
	RedemptionID interface{} `json:"redemptionId"`
	// Status of the redemption.
	Status MarketplaceOfferRedeemResponseStatus `json:"status"`
	JSON   marketplaceOfferRedeemResponseJSON   `json:"-"`
}

// marketplaceOfferRedeemResponseJSON contains the JSON metadata for the struct
// [MarketplaceOfferRedeemResponse]
type marketplaceOfferRedeemResponseJSON struct {
	AssociatedTransactionID apijson.Field
	Message                 apijson.Field
	OfferID                 apijson.Field
	RedemptionDate          apijson.Field
	RedemptionID            apijson.Field
	Status                  apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MarketplaceOfferRedeemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceOfferRedeemResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the redemption.
type MarketplaceOfferRedeemResponseStatus string

const (
	MarketplaceOfferRedeemResponseStatusSuccess MarketplaceOfferRedeemResponseStatus = "success"
	MarketplaceOfferRedeemResponseStatusPending MarketplaceOfferRedeemResponseStatus = "pending"
	MarketplaceOfferRedeemResponseStatusFailed  MarketplaceOfferRedeemResponseStatus = "failed"
)

func (r MarketplaceOfferRedeemResponseStatus) IsKnown() bool {
	switch r {
	case MarketplaceOfferRedeemResponseStatusSuccess, MarketplaceOfferRedeemResponseStatusPending, MarketplaceOfferRedeemResponseStatusFailed:
		return true
	}
	return false
}

type MarketplaceOfferRedeemParams struct {
	// Optional: The ID of the account to use for any associated payment or credit.
	PaymentAccountID param.Field[interface{}] `json:"paymentAccountId"`
}

func (r MarketplaceOfferRedeemParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
