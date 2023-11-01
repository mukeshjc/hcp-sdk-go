// Code generated by go-swagger; DO NOT EDIT.

package statement_service

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
)

// NewStatementServiceListStatementsParams creates a new StatementServiceListStatementsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatementServiceListStatementsParams() *StatementServiceListStatementsParams {
	return &StatementServiceListStatementsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatementServiceListStatementsParamsWithTimeout creates a new StatementServiceListStatementsParams object
// with the ability to set a timeout on a request.
func NewStatementServiceListStatementsParamsWithTimeout(timeout time.Duration) *StatementServiceListStatementsParams {
	return &StatementServiceListStatementsParams{
		timeout: timeout,
	}
}

// NewStatementServiceListStatementsParamsWithContext creates a new StatementServiceListStatementsParams object
// with the ability to set a context for a request.
func NewStatementServiceListStatementsParamsWithContext(ctx context.Context) *StatementServiceListStatementsParams {
	return &StatementServiceListStatementsParams{
		Context: ctx,
	}
}

// NewStatementServiceListStatementsParamsWithHTTPClient creates a new StatementServiceListStatementsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatementServiceListStatementsParamsWithHTTPClient(client *http.Client) *StatementServiceListStatementsParams {
	return &StatementServiceListStatementsParams{
		HTTPClient: client,
	}
}

/*
StatementServiceListStatementsParams contains all the parameters to send to the API endpoint

	for the statement service list statements operation.

	Typically these are written to a http.Request.
*/
type StatementServiceListStatementsParams struct {

	/* BillingAccountID.

	     id is the user-settable ID that uniquely identifies the Billing Account
	within the organization.
	*/
	BillingAccountID string

	/* OrganizationID.

	     organization_id is the ID of the organization to which the Billing Account
	belongs.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the statement service list statements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatementServiceListStatementsParams) WithDefaults() *StatementServiceListStatementsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the statement service list statements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatementServiceListStatementsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the statement service list statements params
func (o *StatementServiceListStatementsParams) WithTimeout(timeout time.Duration) *StatementServiceListStatementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the statement service list statements params
func (o *StatementServiceListStatementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the statement service list statements params
func (o *StatementServiceListStatementsParams) WithContext(ctx context.Context) *StatementServiceListStatementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the statement service list statements params
func (o *StatementServiceListStatementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the statement service list statements params
func (o *StatementServiceListStatementsParams) WithHTTPClient(client *http.Client) *StatementServiceListStatementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the statement service list statements params
func (o *StatementServiceListStatementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the statement service list statements params
func (o *StatementServiceListStatementsParams) WithBillingAccountID(billingAccountID string) *StatementServiceListStatementsParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the statement service list statements params
func (o *StatementServiceListStatementsParams) SetBillingAccountID(billingAccountID string) {
	o.BillingAccountID = billingAccountID
}

// WithOrganizationID adds the organizationID to the statement service list statements params
func (o *StatementServiceListStatementsParams) WithOrganizationID(organizationID string) *StatementServiceListStatementsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the statement service list statements params
func (o *StatementServiceListStatementsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *StatementServiceListStatementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billing_account_id
	if err := r.SetPathParam("billing_account_id", o.BillingAccountID); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
