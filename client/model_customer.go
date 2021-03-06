/*
 * Moov API
 *
 * test
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type Customer struct {
	// Customer ID
	Id string `json:"id,omitempty"`
	// The customers email address
	Email string `json:"email,omitempty"`
	// The depository account to be used by default per transfer. ID must be a valid Customer Depository account
	DefaultDepository string `json:"defaultDepository,omitempty"`
	// Defines the status of the Customer
	Status string `json:"status,omitempty"`
	// Additional meta data to be used for display only
	Metadata string `json:"metadata,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}
