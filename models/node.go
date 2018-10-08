// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Node node
// swagger:model Node
type Node struct {

	// architecture
	Architecture string `json:"architecture,omitempty"`

	// cpu available
	CPUAvailable string `json:"cpu_available,omitempty"`

	// cpu capacity
	CPUCapacity string `json:"cpu_capacity,omitempty"`

	// disk available
	DiskAvailable string `json:"disk_available,omitempty"`

	// disk capacity
	DiskCapacity string `json:"disk_capacity,omitempty"`

	// distribution
	Distribution string `json:"distribution,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// memory available
	MemoryAvailable string `json:"memory_available,omitempty"`

	// memory capacity
	MemoryCapacity string `json:"memory_capacity,omitempty"`

	// region id
	RegionID string `json:"region_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}