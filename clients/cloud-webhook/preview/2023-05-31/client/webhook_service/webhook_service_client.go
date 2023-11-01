// Code generated by go-swagger; DO NOT EDIT.

package webhook_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new webhook service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for webhook service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	WebhookServiceCreateWebhook(params *WebhookServiceCreateWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceCreateWebhookOK, error)

	WebhookServiceDeleteWebhook(params *WebhookServiceDeleteWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceDeleteWebhookOK, error)

	WebhookServiceGetWebhook(params *WebhookServiceGetWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceGetWebhookOK, error)

	WebhookServiceListWebhookDeliveries(params *WebhookServiceListWebhookDeliveriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceListWebhookDeliveriesOK, error)

	WebhookServiceListWebhooks(params *WebhookServiceListWebhooksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceListWebhooksOK, error)

	WebhookServiceUpdateWebhook(params *WebhookServiceUpdateWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceUpdateWebhookOK, error)

	WebhookServiceUpdateWebhookName(params *WebhookServiceUpdateWebhookNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceUpdateWebhookNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
WebhookServiceCreateWebhook creates an enabled webhook that is optionally subscribed to the resources lifecycle events
*/
func (a *Client) WebhookServiceCreateWebhook(params *WebhookServiceCreateWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceCreateWebhookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceCreateWebhookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_CreateWebhook",
		Method:             "POST",
		PathPattern:        "/2023-05-31/webhook/{parent_resource_name}/webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceCreateWebhookReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceCreateWebhookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceCreateWebhookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WebhookServiceDeleteWebhook deletes a webhook and all of its delivery records
*/
func (a *Client) WebhookServiceDeleteWebhook(params *WebhookServiceDeleteWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceDeleteWebhookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceDeleteWebhookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_DeleteWebhook",
		Method:             "DELETE",
		PathPattern:        "/2023-05-31/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceDeleteWebhookReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceDeleteWebhookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceDeleteWebhookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WebhookServiceGetWebhook gets an existing webhook
*/
func (a *Client) WebhookServiceGetWebhook(params *WebhookServiceGetWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceGetWebhookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceGetWebhookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_GetWebhook",
		Method:             "GET",
		PathPattern:        "/2023-05-31/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceGetWebhookReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceGetWebhookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceGetWebhookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WebhookServiceListWebhookDeliveries lists all the deliveries from the last 90 days for a webhook
*/
func (a *Client) WebhookServiceListWebhookDeliveries(params *WebhookServiceListWebhookDeliveriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceListWebhookDeliveriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceListWebhookDeliveriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_ListWebhookDeliveries",
		Method:             "GET",
		PathPattern:        "/2023-05-31/{parent_resource_name}/deliveries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceListWebhookDeliveriesReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceListWebhookDeliveriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceListWebhookDeliveriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WebhookServiceListWebhooks lists all existing webhooks in an h c p project
*/
func (a *Client) WebhookServiceListWebhooks(params *WebhookServiceListWebhooksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceListWebhooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceListWebhooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_ListWebhooks",
		Method:             "GET",
		PathPattern:        "/2023-05-31/webhook/{parent_resource_name}/webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceListWebhooksReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceListWebhooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceListWebhooksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WebhookServiceUpdateWebhook updates a webhook configuration and subscriptions and sets the status to enabled or disabled
*/
func (a *Client) WebhookServiceUpdateWebhook(params *WebhookServiceUpdateWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceUpdateWebhookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceUpdateWebhookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_UpdateWebhook",
		Method:             "PATCH",
		PathPattern:        "/2023-05-31/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceUpdateWebhookReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceUpdateWebhookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceUpdateWebhookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WebhookServiceUpdateWebhookName updates a webhook name updating a webhook name changes its resource name
*/
func (a *Client) WebhookServiceUpdateWebhookName(params *WebhookServiceUpdateWebhookNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookServiceUpdateWebhookNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookServiceUpdateWebhookNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookService_UpdateWebhookName",
		Method:             "PATCH",
		PathPattern:        "/2023-05-31/{resource_name}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookServiceUpdateWebhookNameReader{formats: a.formats},
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
	success, ok := result.(*WebhookServiceUpdateWebhookNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WebhookServiceUpdateWebhookNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
