// Code generated by go-swagger; DO NOT EDIT.

package billing_account_service

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

// NewBillingAccountServiceRemoveOnDemandPaymentMethodParams creates a new BillingAccountServiceRemoveOnDemandPaymentMethodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingAccountServiceRemoveOnDemandPaymentMethodParams() *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	return &BillingAccountServiceRemoveOnDemandPaymentMethodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingAccountServiceRemoveOnDemandPaymentMethodParamsWithTimeout creates a new BillingAccountServiceRemoveOnDemandPaymentMethodParams object
// with the ability to set a timeout on a request.
func NewBillingAccountServiceRemoveOnDemandPaymentMethodParamsWithTimeout(timeout time.Duration) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	return &BillingAccountServiceRemoveOnDemandPaymentMethodParams{
		timeout: timeout,
	}
}

// NewBillingAccountServiceRemoveOnDemandPaymentMethodParamsWithContext creates a new BillingAccountServiceRemoveOnDemandPaymentMethodParams object
// with the ability to set a context for a request.
func NewBillingAccountServiceRemoveOnDemandPaymentMethodParamsWithContext(ctx context.Context) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	return &BillingAccountServiceRemoveOnDemandPaymentMethodParams{
		Context: ctx,
	}
}

// NewBillingAccountServiceRemoveOnDemandPaymentMethodParamsWithHTTPClient creates a new BillingAccountServiceRemoveOnDemandPaymentMethodParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingAccountServiceRemoveOnDemandPaymentMethodParamsWithHTTPClient(client *http.Client) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	return &BillingAccountServiceRemoveOnDemandPaymentMethodParams{
		HTTPClient: client,
	}
}

/*
BillingAccountServiceRemoveOnDemandPaymentMethodParams contains all the parameters to send to the API endpoint

	for the billing account service remove on demand payment method operation.

	Typically these are written to a http.Request.
*/
type BillingAccountServiceRemoveOnDemandPaymentMethodParams struct {

	// BillingAccountID.
	BillingAccountID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing account service remove on demand payment method params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WithDefaults() *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing account service remove on demand payment method params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WithTimeout(timeout time.Duration) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WithContext(ctx context.Context) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WithHTTPClient(client *http.Client) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WithBillingAccountID(billingAccountID string) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) SetBillingAccountID(billingAccountID string) {
	o.BillingAccountID = billingAccountID
}

// WithOrganizationID adds the organizationID to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WithOrganizationID(organizationID string) *BillingAccountServiceRemoveOnDemandPaymentMethodParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the billing account service remove on demand payment method params
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
