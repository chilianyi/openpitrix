// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetRuntimeStatisticsResponse openpitrix get runtime statistics response
// swagger:model openpitrixGetRuntimeStatisticsResponse
type OpenpitrixGetRuntimeStatisticsResponse struct {

	// last two week created
	LastTwoWeekCreated map[string]int64 `json:"last_two_week_created,omitempty"`

	// provider count
	ProviderCount int64 `json:"provider_count,omitempty"`

	// runtime count
	RuntimeCount int64 `json:"runtime_count,omitempty"`

	// top ten providers
	TopTenProviders map[string]int64 `json:"top_ten_providers,omitempty"`
}

// Validate validates this openpitrix get runtime statistics response
func (m *OpenpitrixGetRuntimeStatisticsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetRuntimeStatisticsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetRuntimeStatisticsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetRuntimeStatisticsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
