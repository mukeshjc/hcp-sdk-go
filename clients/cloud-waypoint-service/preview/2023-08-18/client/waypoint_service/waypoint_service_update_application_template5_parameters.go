// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// NewWaypointServiceUpdateApplicationTemplate5Params creates a new WaypointServiceUpdateApplicationTemplate5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplicationTemplate5Params() *WaypointServiceUpdateApplicationTemplate5Params {
	return &WaypointServiceUpdateApplicationTemplate5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplate5ParamsWithTimeout creates a new WaypointServiceUpdateApplicationTemplate5Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplicationTemplate5ParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplate5Params {
	return &WaypointServiceUpdateApplicationTemplate5Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplate5ParamsWithContext creates a new WaypointServiceUpdateApplicationTemplate5Params object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplicationTemplate5ParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplate5Params {
	return &WaypointServiceUpdateApplicationTemplate5Params{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplicationTemplate5ParamsWithHTTPClient creates a new WaypointServiceUpdateApplicationTemplate5Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplicationTemplate5ParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplate5Params {
	return &WaypointServiceUpdateApplicationTemplate5Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplicationTemplate5Params contains all the parameters to send to the API endpoint

	for the waypoint service update application template5 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplicationTemplate5Params struct {

	/* ApplicationTemplate.

	     application_template resembles the desired updated state of the existing
	application template.
	*/
	ApplicationTemplate *models.HashicorpCloudWaypointApplicationTemplate

	/* ExistingApplicationTemplateID.

	   ID of the ApplicationTemplate
	*/
	ExistingApplicationTemplateID string

	/* ExistingApplicationTemplateName.

	   Name of the ApplicationTemplate
	*/
	ExistingApplicationTemplateName *string

	// NamespaceID.
	NamespaceID string

	/* UseModuleReadme.

	     If true, will auto-import the readme from the Terraform module used
	rather than the one specified in application_template.
	*/
	UseModuleReadme *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update application template5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithDefaults() *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update application template5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationTemplate adds the applicationTemplate to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithApplicationTemplate(applicationTemplate *models.HashicorpCloudWaypointApplicationTemplate) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetApplicationTemplate(applicationTemplate)
	return o
}

// SetApplicationTemplate adds the applicationTemplate to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetApplicationTemplate(applicationTemplate *models.HashicorpCloudWaypointApplicationTemplate) {
	o.ApplicationTemplate = applicationTemplate
}

// WithExistingApplicationTemplateID adds the existingApplicationTemplateID to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithExistingApplicationTemplateID(existingApplicationTemplateID string) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetExistingApplicationTemplateID(existingApplicationTemplateID)
	return o
}

// SetExistingApplicationTemplateID adds the existingApplicationTemplateId to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetExistingApplicationTemplateID(existingApplicationTemplateID string) {
	o.ExistingApplicationTemplateID = existingApplicationTemplateID
}

// WithExistingApplicationTemplateName adds the existingApplicationTemplateName to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithExistingApplicationTemplateName(existingApplicationTemplateName *string) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetExistingApplicationTemplateName(existingApplicationTemplateName)
	return o
}

// SetExistingApplicationTemplateName adds the existingApplicationTemplateName to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetExistingApplicationTemplateName(existingApplicationTemplateName *string) {
	o.ExistingApplicationTemplateName = existingApplicationTemplateName
}

// WithNamespaceID adds the namespaceID to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithNamespaceID(namespaceID string) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithUseModuleReadme adds the useModuleReadme to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) WithUseModuleReadme(useModuleReadme *bool) *WaypointServiceUpdateApplicationTemplate5Params {
	o.SetUseModuleReadme(useModuleReadme)
	return o
}

// SetUseModuleReadme adds the useModuleReadme to the waypoint service update application template5 params
func (o *WaypointServiceUpdateApplicationTemplate5Params) SetUseModuleReadme(useModuleReadme *bool) {
	o.UseModuleReadme = useModuleReadme
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplicationTemplate5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ApplicationTemplate != nil {
		if err := r.SetBodyParam(o.ApplicationTemplate); err != nil {
			return err
		}
	}

	// path param existing_application_template.id
	if err := r.SetPathParam("existing_application_template.id", o.ExistingApplicationTemplateID); err != nil {
		return err
	}

	if o.ExistingApplicationTemplateName != nil {

		// query param existing_application_template.name
		var qrExistingApplicationTemplateName string

		if o.ExistingApplicationTemplateName != nil {
			qrExistingApplicationTemplateName = *o.ExistingApplicationTemplateName
		}
		qExistingApplicationTemplateName := qrExistingApplicationTemplateName
		if qExistingApplicationTemplateName != "" {

			if err := r.SetQueryParam("existing_application_template.name", qExistingApplicationTemplateName); err != nil {
				return err
			}
		}
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if o.UseModuleReadme != nil {

		// query param use_module_readme
		var qrUseModuleReadme bool

		if o.UseModuleReadme != nil {
			qrUseModuleReadme = *o.UseModuleReadme
		}
		qUseModuleReadme := swag.FormatBool(qrUseModuleReadme)
		if qUseModuleReadme != "" {

			if err := r.SetQueryParam("use_module_readme", qUseModuleReadme); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
