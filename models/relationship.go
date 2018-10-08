// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Relationship relationship
// swagger:model Relationship
type Relationship struct {

	// Amount of bandwidth available on the link - bit/s
	BandwidthAvailable int64 `json:"bandwidth_available,omitempty"`

	// Amount of bandwidth capacity on the link - bit/s
	BandwidthCapacity int64 `json:"bandwidth_capacity,omitempty"`

	// endpoint a
	EndpointA string `json:"endpoint_a,omitempty"`

	// endpoint b
	EndpointB string `json:"endpoint_b,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Link latency - milliseconds
	Latency int64 `json:"latency,omitempty"`

	// prices
	Prices *RelationshipPrices `json:"prices,omitempty"`

	// region id
	RegionID string `json:"region_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this relationship
func (m *Relationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Relationship) validatePrices(formats strfmt.Registry) error {

	if swag.IsZero(m.Prices) { // not required
		return nil
	}

	if m.Prices != nil {

		if err := m.Prices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prices")
			}
			return err
		}
	}

	return nil
}

var relationshipTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relationshipTypeStatusPropEnum = append(relationshipTypeStatusPropEnum, v)
	}
}

const (
	// RelationshipStatusUp captures enum value "up"
	RelationshipStatusUp string = "up"
	// RelationshipStatusDown captures enum value "down"
	RelationshipStatusDown string = "down"
)

// prop value enum
func (m *Relationship) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, relationshipTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Relationship) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Relationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Relationship) UnmarshalBinary(b []byte) error {
	var res Relationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
