// Code generated by go-swagger; DO NOT EDIT.

package data_source_registration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data source registration service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data source registration service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDataSourceByID(params *GetDataSourceByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataSourceByIDOK, error)

	OffboardDataSource(params *OffboardDataSourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OffboardDataSourceOK, error)

	OnboardDataSource(params *OnboardDataSourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OnboardDataSourceOK, error)

	UpdateDataSourceToken(params *UpdateDataSourceTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDataSourceTokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDataSourceByID get data source by ID API
*/
func (a *Client) GetDataSourceByID(params *GetDataSourceByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataSourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataSourceByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDataSourceByID",
		Method:             "GET",
		PathPattern:        "/2023-05-01/vault-radar/projects/{location.project_id}/data-source-registrations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDataSourceByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataSourceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDataSourceByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OffboardDataSource offboard data source API
*/
func (a *Client) OffboardDataSource(params *OffboardDataSourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OffboardDataSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOffboardDataSourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OffboardDataSource",
		Method:             "POST",
		PathPattern:        "/2023-05-01/vault-radar/projects/{location.project_id}/data-source-registrations/offboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OffboardDataSourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OffboardDataSourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OffboardDataSourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OnboardDataSource onboard data source API
*/
func (a *Client) OnboardDataSource(params *OnboardDataSourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OnboardDataSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOnboardDataSourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OnboardDataSource",
		Method:             "POST",
		PathPattern:        "/2023-05-01/vault-radar/projects/{location.project_id}/data-source-registrations/onboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OnboardDataSourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OnboardDataSourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OnboardDataSourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateDataSourceToken update data source token API
*/
func (a *Client) UpdateDataSourceToken(params *UpdateDataSourceTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDataSourceTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDataSourceTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDataSourceToken",
		Method:             "PUT",
		PathPattern:        "/2023-05-01/vault-radar/projects/{location.project_id}/data-source-registrations/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDataSourceTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDataSourceTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateDataSourceTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
