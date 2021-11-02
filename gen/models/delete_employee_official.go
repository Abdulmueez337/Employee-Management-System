// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteEmployeeOfficial delete employee official
//
// swagger:model deleteEmployeeOfficial
type DeleteEmployeeOfficial struct {

	// job type
	// Example: Resign,Terminate
	// Required: true
	JobType *string `json:"job_type"`
}

// Validate validates this delete employee official
func (m *DeleteEmployeeOfficial) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteEmployeeOfficial) validateJobType(formats strfmt.Registry) error {

	if err := validate.Required("job_type", "body", m.JobType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete employee official based on context it is used
func (m *DeleteEmployeeOfficial) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteEmployeeOfficial) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteEmployeeOfficial) UnmarshalBinary(b []byte) error {
	var res DeleteEmployeeOfficial
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
