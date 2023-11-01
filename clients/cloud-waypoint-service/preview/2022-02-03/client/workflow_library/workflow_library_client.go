// Code generated by go-swagger; DO NOT EDIT.

package workflow_library

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflow library API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow library API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	WorkflowLibraryAddToLibrary(params *WorkflowLibraryAddToLibraryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryAddToLibraryOK, error)

	WorkflowLibraryGetFromLibrary(params *WorkflowLibraryGetFromLibraryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryGetFromLibraryOK, error)

	WorkflowLibraryGetFromLibrary2(params *WorkflowLibraryGetFromLibrary2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryGetFromLibrary2OK, error)

	WorkflowLibraryGetFromLibrary3(params *WorkflowLibraryGetFromLibrary3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryGetFromLibrary3OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
WorkflowLibraryAddToLibrary bs e t a add to library adds a workflow program to the organizations library
*/
func (a *Client) WorkflowLibraryAddToLibrary(params *WorkflowLibraryAddToLibraryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryAddToLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowLibraryAddToLibraryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowLibrary_AddToLibrary",
		Method:             "PUT",
		PathPattern:        "/waypoint/2022-02-03/namespace/workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WorkflowLibraryAddToLibraryReader{formats: a.formats},
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
	success, ok := result.(*WorkflowLibraryAddToLibraryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkflowLibraryAddToLibraryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WorkflowLibraryGetFromLibrary bs e t a get from library returns workflow programs in the library that match the requested fields
*/
func (a *Client) WorkflowLibraryGetFromLibrary(params *WorkflowLibraryGetFromLibraryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryGetFromLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowLibraryGetFromLibraryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowLibrary_GetFromLibrary",
		Method:             "GET",
		PathPattern:        "/waypoint/2022-02-03/namespace/workflows/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WorkflowLibraryGetFromLibraryReader{formats: a.formats},
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
	success, ok := result.(*WorkflowLibraryGetFromLibraryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkflowLibraryGetFromLibraryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WorkflowLibraryGetFromLibrary2 bs e t a get from library returns workflow programs in the library that match the requested fields
*/
func (a *Client) WorkflowLibraryGetFromLibrary2(params *WorkflowLibraryGetFromLibrary2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryGetFromLibrary2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowLibraryGetFromLibrary2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowLibrary_GetFromLibrary2",
		Method:             "GET",
		PathPattern:        "/waypoint/2022-02-03/namespace/workflows/search/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WorkflowLibraryGetFromLibrary2Reader{formats: a.formats},
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
	success, ok := result.(*WorkflowLibraryGetFromLibrary2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkflowLibraryGetFromLibrary2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WorkflowLibraryGetFromLibrary3 bs e t a get from library returns workflow programs in the library that match the requested fields
*/
func (a *Client) WorkflowLibraryGetFromLibrary3(params *WorkflowLibraryGetFromLibrary3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WorkflowLibraryGetFromLibrary3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowLibraryGetFromLibrary3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowLibrary_GetFromLibrary3",
		Method:             "GET",
		PathPattern:        "/waypoint/2022-02-03/namespace/workflows/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WorkflowLibraryGetFromLibrary3Reader{formats: a.formats},
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
	success, ok := result.(*WorkflowLibraryGetFromLibrary3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkflowLibraryGetFromLibrary3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}