// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewListConfluentIntegrationsParams creates a new ListConfluentIntegrationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListConfluentIntegrationsParams() *ListConfluentIntegrationsParams {
	return &ListConfluentIntegrationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListConfluentIntegrationsParamsWithTimeout creates a new ListConfluentIntegrationsParams object
// with the ability to set a timeout on a request.
func NewListConfluentIntegrationsParamsWithTimeout(timeout time.Duration) *ListConfluentIntegrationsParams {
	return &ListConfluentIntegrationsParams{
		timeout: timeout,
	}
}

// NewListConfluentIntegrationsParamsWithContext creates a new ListConfluentIntegrationsParams object
// with the ability to set a context for a request.
func NewListConfluentIntegrationsParamsWithContext(ctx context.Context) *ListConfluentIntegrationsParams {
	return &ListConfluentIntegrationsParams{
		Context: ctx,
	}
}

// NewListConfluentIntegrationsParamsWithHTTPClient creates a new ListConfluentIntegrationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListConfluentIntegrationsParamsWithHTTPClient(client *http.Client) *ListConfluentIntegrationsParams {
	return &ListConfluentIntegrationsParams{
		HTTPClient: client,
	}
}

/*
ListConfluentIntegrationsParams contains all the parameters to send to the API endpoint

	for the list confluent integrations operation.

	Typically these are written to a http.Request.
*/
type ListConfluentIntegrationsParams struct {

	// OrganizationID.
	OrganizationID string

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

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list confluent integrations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListConfluentIntegrationsParams) WithDefaults() *ListConfluentIntegrationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list confluent integrations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListConfluentIntegrationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithTimeout(timeout time.Duration) *ListConfluentIntegrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithContext(ctx context.Context) *ListConfluentIntegrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithHTTPClient(client *http.Client) *ListConfluentIntegrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithOrganizationID(organizationID string) *ListConfluentIntegrationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListConfluentIntegrationsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithPaginationPageSize(paginationPageSize *int64) *ListConfluentIntegrationsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListConfluentIntegrationsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithProjectID adds the projectID to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) WithProjectID(projectID string) *ListConfluentIntegrationsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list confluent integrations params
func (o *ListConfluentIntegrationsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListConfluentIntegrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}