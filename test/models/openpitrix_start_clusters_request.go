// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixStartClustersRequest openpitrix start clusters request
// swagger:model openpitrixStartClustersRequest
type OpenpitrixStartClustersRequest struct {

	// advanced param
	AdvancedParam []string `json:"advanced_param"`

	// cluster id
	ClusterID []string `json:"cluster_id"`
}

// Validate validates this openpitrix start clusters request
func (m *OpenpitrixStartClustersRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedParam(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixStartClustersRequest) validateAdvancedParam(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvancedParam) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixStartClustersRequest) validateClusterID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixStartClustersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixStartClustersRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixStartClustersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
