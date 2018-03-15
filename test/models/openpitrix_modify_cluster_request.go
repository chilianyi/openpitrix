// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyClusterRequest openpitrix modify cluster request
// swagger:model openpitrixModifyClusterRequest
type OpenpitrixModifyClusterRequest struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// transition status
	TransitionStatus string `json:"transition_status,omitempty"`

	// upgrade status
	UpgradeStatus string `json:"upgrade_status,omitempty"`

	// upgrade time
	UpgradeTime strfmt.DateTime `json:"upgrade_time,omitempty"`
}

// Validate validates this openpitrix modify cluster request
func (m *OpenpitrixModifyClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyClusterRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
