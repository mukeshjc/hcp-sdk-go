// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListWorkloadIdentityProviderParams creates a new ListWorkloadIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListWorkloadIdentityProviderParams() *ListWorkloadIdentityProviderParams {
	return &ListWorkloadIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListWorkloadIdentityProviderParamsWithTimeout creates a new ListWorkloadIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewListWorkloadIdentityProviderParamsWithTimeout(timeout time.Duration) *ListWorkloadIdentityProviderParams {
	return &ListWorkloadIdentityProviderParams{
		timeout: timeout,
	}
}

// NewListWorkloadIdentityProviderParamsWithContext creates a new ListWorkloadIdentityProviderParams object
// with the ability to set a context for a request.
func NewListWorkloadIdentityProviderParamsWithContext(ctx context.Context) *ListWorkloadIdentityProviderParams {
	return &ListWorkloadIdentityProviderParams{
		Context: ctx,
	}
}

// NewListWorkloadIdentityProviderParamsWithHTTPClient creates a new ListWorkloadIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewListWorkloadIdentityProviderParamsWithHTTPClient(client *http.Client) *ListWorkloadIdentityProviderParams {
	return &ListWorkloadIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
ListWorkloadIdentityProviderParams contains all the parameters to send to the API endpoint

	for the list workload identity provider operation.

	Typically these are written to a http.Request.
*/
type ListWorkloadIdentityProviderParams struct {

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* ParentResourceName.

	     parent_resource_name is the resource name of the service principal to list
	workload identity providers from.
	*/
	ParentResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWorkloadIdentityProviderParams) WithDefaults() *ListWorkloadIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWorkloadIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithTimeout(timeout time.Duration) *ListWorkloadIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithContext(ctx context.Context) *ListWorkloadIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithHTTPClient(client *http.Client) *ListWorkloadIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListWorkloadIdentityProviderParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithPaginationPageSize(paginationPageSize *int64) *ListWorkloadIdentityProviderParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListWorkloadIdentityProviderParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithParentResourceName adds the parentResourceName to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) WithParentResourceName(parentResourceName string) *ListWorkloadIdentityProviderParams {
	o.SetParentResourceName(parentResourceName)
	return o
}

// SetParentResourceName adds the parentResourceName to the list workload identity provider params
func (o *ListWorkloadIdentityProviderParams) SetParentResourceName(parentResourceName string) {
	o.ParentResourceName = parentResourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ListWorkloadIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	// path param parent_resource_name
	if err := r.SetPathParam("parent_resource_name", o.ParentResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}