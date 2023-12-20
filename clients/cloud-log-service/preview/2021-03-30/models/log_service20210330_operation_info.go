// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogService20210330OperationInfo OperationInfo contains information about the operation a control-plane event
// is part of.
//
// swagger:model log_service_20210330OperationInfo
type LogService20210330OperationInfo struct {

	// operation_id is the ID of the operation.
	OperationID string `json:"operation_id,omitempty"`
}

// Validate validates this log service 20210330 operation info
func (m *LogService20210330OperationInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log service 20210330 operation info based on context it is used
func (m *LogService20210330OperationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330OperationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330OperationInfo) UnmarshalBinary(b []byte) error {
	var res LogService20210330OperationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
