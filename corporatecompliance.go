// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"github.com/jocall3/1231-go/option"
)

// CorporateComplianceService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateComplianceService] method instead.
type CorporateComplianceService struct {
	Options []option.RequestOption
	Audits  *CorporateComplianceAuditService
}

// NewCorporateComplianceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateComplianceService(opts ...option.RequestOption) (r *CorporateComplianceService) {
	r = &CorporateComplianceService{}
	r.Options = opts
	r.Audits = NewCorporateComplianceAuditService(opts...)
	return
}
