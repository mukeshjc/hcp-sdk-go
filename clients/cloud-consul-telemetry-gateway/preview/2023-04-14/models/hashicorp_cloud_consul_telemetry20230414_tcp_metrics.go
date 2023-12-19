// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulTelemetry20230414TCPMetrics hashicorp cloud consul telemetry 20230414 Tcp metrics
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.TcpMetrics
type HashicorpCloudConsulTelemetry20230414TCPMetrics struct {

	// cx_connect_fail_rate is the per-second rate of connection failures to a service.
	CxConnectFailRate float32 `json:"cx_connect_fail_rate,omitempty"`

	// cx_none_healthy_rate is the per-second rate of connection failures to a service for lack of any healthy host.
	CxNoneHealthyRate float32 `json:"cx_none_healthy_rate,omitempty"`

	// cx_rate is the per-second rate of connections to a service.
	CxRate float32 `json:"cx_rate,omitempty"`

	// rx_bits_rate is the per-second rate of bits recieved by a service.
	RxBitsRate float32 `json:"rx_bits_rate,omitempty"`

	// tc_bits_rate is the per-second rate of bits transmitted by a service.
	TxBitsRate float32 `json:"tx_bits_rate,omitempty"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 Tcp metrics
func (m *HashicorpCloudConsulTelemetry20230414TCPMetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul telemetry 20230414 Tcp metrics based on context it is used
func (m *HashicorpCloudConsulTelemetry20230414TCPMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414TCPMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414TCPMetrics) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414TCPMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}