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

// LogService20210330GetArchiveResponse GetArchiveResponse represents the response for an archive request.
//
// swagger:model log_service_20210330GetArchiveResponse
type LogService20210330GetArchiveResponse struct {

	// archive_requested_end_time represents the end time which the archive record was specified.
	// Format: date-time
	ArchiveEndTime strfmt.DateTime `json:"archive_end_time,omitempty"`

	// archive_requested_time represents the time which the archive record was created.
	// Format: date-time
	ArchiveRequestedTime strfmt.DateTime `json:"archive_requested_time,omitempty"`

	// archive_requested_start_time represents the start time which the archive record was specified.
	// Format: date-time
	ArchiveStartTime strfmt.DateTime `json:"archive_start_time,omitempty"`

	// error represents the error returned for an archive.
	Error string `json:"error,omitempty"`

	// id represents the database id of the archive record.
	ID string `json:"id,omitempty"`

	// status represents the status of the archive record.
	Status *LogService20210330Status `json:"status,omitempty"`

	// url represents the url from which to download the archive record.
	URL string `json:"url,omitempty"`

	// url_expiration represents the timestamp of when the url expires.
	// Format: date-time
	URLExpiration strfmt.DateTime `json:"url_expiration,omitempty"`
}

// Validate validates this log service 20210330 get archive response
func (m *LogService20210330GetArchiveResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchiveEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchiveRequestedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchiveStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLExpiration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330GetArchiveResponse) validateArchiveEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchiveEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("archive_end_time", "body", "date-time", m.ArchiveEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogService20210330GetArchiveResponse) validateArchiveRequestedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchiveRequestedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("archive_requested_time", "body", "date-time", m.ArchiveRequestedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogService20210330GetArchiveResponse) validateArchiveStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchiveStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("archive_start_time", "body", "date-time", m.ArchiveStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogService20210330GetArchiveResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330GetArchiveResponse) validateURLExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.URLExpiration) { // not required
		return nil
	}

	if err := validate.FormatOf("url_expiration", "body", "date-time", m.URLExpiration.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this log service 20210330 get archive response based on the context it is used
func (m *LogService20210330GetArchiveResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330GetArchiveResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330GetArchiveResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330GetArchiveResponse) UnmarshalBinary(b []byte) error {
	var res LogService20210330GetArchiveResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
