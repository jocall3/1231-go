// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"github.com/jocall3/1231-go/option"
)

// AIOracleService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOracleService] method instead.
type AIOracleService struct {
	Options     []option.RequestOption
	Simulate    *AIOracleSimulateService
	Simulations *AIOracleSimulationService
}

// NewAIOracleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOracleService(opts ...option.RequestOption) (r *AIOracleService) {
	r = &AIOracleService{}
	r.Options = opts
	r.Simulate = NewAIOracleSimulateService(opts...)
	r.Simulations = NewAIOracleSimulationService(opts...)
	return
}
