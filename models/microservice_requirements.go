// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MicroserviceRequirements microservice requirements
// swagger:model MicroserviceRequirements
type MicroserviceRequirements struct {

	// cpu required
	CPURequired string `json:"cpu_required,omitempty"`

	// deployment descriptor
	DeploymentDescriptor string `json:"deployment_descriptor,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// disk required
	DiskRequired string `json:"disk_required,omitempty"`

	// memory required
	MemoryRequired string `json:"memory_required,omitempty"`

	// Unique FogAtlas internal name, defined by the user in a deployment request
	Name string `json:"name,omitempty"`

	// The price of this micrservice calculated through the negotiator
	PriceComputed float64 `json:"price_computed,omitempty"`

	// Upper bound limit of the price of this microservice fixed by the user/client. -1 means not set
	PriceRequired float64 `json:"price_required,omitempty"`

	// region id
	RegionID string `json:"region_id,omitempty"`

	// region required
	RegionRequired string `json:"region_required,omitempty"`
}

// Validate validates this microservice requirements
func (m *MicroserviceRequirements) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MicroserviceRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MicroserviceRequirements) UnmarshalBinary(b []byte) error {
	var res MicroserviceRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}