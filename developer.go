// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"github.com/jocall3/1231-go/option"
)

// DeveloperService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeveloperService] method instead.
type DeveloperService struct {
	Options  []option.RequestOption
	Webhooks *DeveloperWebhookService
	APIKeys  *DeveloperAPIKeyService
}

// NewDeveloperService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeveloperService(opts ...option.RequestOption) (r *DeveloperService) {
	r = &DeveloperService{}
	r.Options = opts
	r.Webhooks = NewDeveloperWebhookService(opts...)
	r.APIKeys = NewDeveloperAPIKeyService(opts...)
	return
}
