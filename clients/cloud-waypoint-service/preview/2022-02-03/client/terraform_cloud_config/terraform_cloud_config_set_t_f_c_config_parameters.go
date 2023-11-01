// Code generated by go-swagger; DO NOT EDIT.

package terraform_cloud_config

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// NewTerraformCloudConfigSetTFCConfigParams creates a new TerraformCloudConfigSetTFCConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTerraformCloudConfigSetTFCConfigParams() *TerraformCloudConfigSetTFCConfigParams {
	return &TerraformCloudConfigSetTFCConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTerraformCloudConfigSetTFCConfigParamsWithTimeout creates a new TerraformCloudConfigSetTFCConfigParams object
// with the ability to set a timeout on a request.
func NewTerraformCloudConfigSetTFCConfigParamsWithTimeout(timeout time.Duration) *TerraformCloudConfigSetTFCConfigParams {
	return &TerraformCloudConfigSetTFCConfigParams{
		timeout: timeout,
	}
}

// NewTerraformCloudConfigSetTFCConfigParamsWithContext creates a new TerraformCloudConfigSetTFCConfigParams object
// with the ability to set a context for a request.
func NewTerraformCloudConfigSetTFCConfigParamsWithContext(ctx context.Context) *TerraformCloudConfigSetTFCConfigParams {
	return &TerraformCloudConfigSetTFCConfigParams{
		Context: ctx,
	}
}

// NewTerraformCloudConfigSetTFCConfigParamsWithHTTPClient creates a new TerraformCloudConfigSetTFCConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewTerraformCloudConfigSetTFCConfigParamsWithHTTPClient(client *http.Client) *TerraformCloudConfigSetTFCConfigParams {
	return &TerraformCloudConfigSetTFCConfigParams{
		HTTPClient: client,
	}
}

/*
TerraformCloudConfigSetTFCConfigParams contains all the parameters to send to the API endpoint

	for the terraform cloud config set t f c config operation.

	Typically these are written to a http.Request.
*/
type TerraformCloudConfigSetTFCConfigParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointSetTFCConfigRequest

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the terraform cloud config set t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TerraformCloudConfigSetTFCConfigParams) WithDefaults() *TerraformCloudConfigSetTFCConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the terraform cloud config set t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TerraformCloudConfigSetTFCConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) WithTimeout(timeout time.Duration) *TerraformCloudConfigSetTFCConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) WithContext(ctx context.Context) *TerraformCloudConfigSetTFCConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) WithHTTPClient(client *http.Client) *TerraformCloudConfigSetTFCConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) WithBody(body *models.HashicorpCloudWaypointSetTFCConfigRequest) *TerraformCloudConfigSetTFCConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) SetBody(body *models.HashicorpCloudWaypointSetTFCConfigRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) WithNamespaceID(namespaceID string) *TerraformCloudConfigSetTFCConfigParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the terraform cloud config set t f c config params
func (o *TerraformCloudConfigSetTFCConfigParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *TerraformCloudConfigSetTFCConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}