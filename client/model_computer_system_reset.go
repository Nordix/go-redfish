/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type ComputerSystemReset struct {
	// The unique identifier for a resource.
	Target string `json:"target,omitempty"`
	ResetTypeRedfishAllowableValues []ResetType `json:"ResetType@Redfish.AllowableValues,omitempty"`
}
