// Code generated by go-swagger; DO NOT EDIT.

package log_service

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

// NewLogServiceGetArchiveParams creates a new LogServiceGetArchiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceGetArchiveParams() *LogServiceGetArchiveParams {
	return &LogServiceGetArchiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceGetArchiveParamsWithTimeout creates a new LogServiceGetArchiveParams object
// with the ability to set a timeout on a request.
func NewLogServiceGetArchiveParamsWithTimeout(timeout time.Duration) *LogServiceGetArchiveParams {
	return &LogServiceGetArchiveParams{
		timeout: timeout,
	}
}

// NewLogServiceGetArchiveParamsWithContext creates a new LogServiceGetArchiveParams object
// with the ability to set a context for a request.
func NewLogServiceGetArchiveParamsWithContext(ctx context.Context) *LogServiceGetArchiveParams {
	return &LogServiceGetArchiveParams{
		Context: ctx,
	}
}

// NewLogServiceGetArchiveParamsWithHTTPClient creates a new LogServiceGetArchiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceGetArchiveParamsWithHTTPClient(client *http.Client) *LogServiceGetArchiveParams {
	return &LogServiceGetArchiveParams{
		HTTPClient: client,
	}
}

/*
LogServiceGetArchiveParams contains all the parameters to send to the API endpoint

	for the log service get archive operation.

	Typically these are written to a http.Request.
*/
type LogServiceGetArchiveParams struct {

	/* ID.

	   id represents the database id of the archive record.
	*/
	ID string

	/* ResourceDescription.

	     description is a human-friendly description for this link. This is
	used primarily for informational purposes such as error messages.
	*/
	ResourceDescription *string

	/* ResourceID.

	   id is the identifier for this resource.
	*/
	ResourceID string

	/* ResourceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	ResourceLocationOrganizationID string

	/* ResourceLocationProjectID.

	   project_id is the projects id.
	*/
	ResourceLocationProjectID string

	/* ResourceLocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	ResourceLocationRegionProvider *string

	/* ResourceLocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	ResourceLocationRegionRegion *string

	/* ResourceType.

	     type is the unique type of the resource. Each service publishes a
	unique set of types. The type value is recommended to be formatted
	in "<org>.<type>" such as "hashicorp.hvn". This is to prevent conflicts
	in the future, but any string value will work.
	*/
	ResourceType *string

	/* ResourceUUID.

	   uuid is the unique UUID for this resource.
	*/
	ResourceUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log service get archive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceGetArchiveParams) WithDefaults() *LogServiceGetArchiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service get archive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceGetArchiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service get archive params
func (o *LogServiceGetArchiveParams) WithTimeout(timeout time.Duration) *LogServiceGetArchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service get archive params
func (o *LogServiceGetArchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service get archive params
func (o *LogServiceGetArchiveParams) WithContext(ctx context.Context) *LogServiceGetArchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service get archive params
func (o *LogServiceGetArchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service get archive params
func (o *LogServiceGetArchiveParams) WithHTTPClient(client *http.Client) *LogServiceGetArchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service get archive params
func (o *LogServiceGetArchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the log service get archive params
func (o *LogServiceGetArchiveParams) WithID(id string) *LogServiceGetArchiveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the log service get archive params
func (o *LogServiceGetArchiveParams) SetID(id string) {
	o.ID = id
}

// WithResourceDescription adds the resourceDescription to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceDescription(resourceDescription *string) *LogServiceGetArchiveParams {
	o.SetResourceDescription(resourceDescription)
	return o
}

// SetResourceDescription adds the resourceDescription to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceDescription(resourceDescription *string) {
	o.ResourceDescription = resourceDescription
}

// WithResourceID adds the resourceID to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceID(resourceID string) *LogServiceGetArchiveParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceLocationOrganizationID adds the resourceLocationOrganizationID to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceLocationOrganizationID(resourceLocationOrganizationID string) *LogServiceGetArchiveParams {
	o.SetResourceLocationOrganizationID(resourceLocationOrganizationID)
	return o
}

// SetResourceLocationOrganizationID adds the resourceLocationOrganizationId to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceLocationOrganizationID(resourceLocationOrganizationID string) {
	o.ResourceLocationOrganizationID = resourceLocationOrganizationID
}

// WithResourceLocationProjectID adds the resourceLocationProjectID to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceLocationProjectID(resourceLocationProjectID string) *LogServiceGetArchiveParams {
	o.SetResourceLocationProjectID(resourceLocationProjectID)
	return o
}

// SetResourceLocationProjectID adds the resourceLocationProjectId to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceLocationProjectID(resourceLocationProjectID string) {
	o.ResourceLocationProjectID = resourceLocationProjectID
}

// WithResourceLocationRegionProvider adds the resourceLocationRegionProvider to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceLocationRegionProvider(resourceLocationRegionProvider *string) *LogServiceGetArchiveParams {
	o.SetResourceLocationRegionProvider(resourceLocationRegionProvider)
	return o
}

// SetResourceLocationRegionProvider adds the resourceLocationRegionProvider to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceLocationRegionProvider(resourceLocationRegionProvider *string) {
	o.ResourceLocationRegionProvider = resourceLocationRegionProvider
}

// WithResourceLocationRegionRegion adds the resourceLocationRegionRegion to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceLocationRegionRegion(resourceLocationRegionRegion *string) *LogServiceGetArchiveParams {
	o.SetResourceLocationRegionRegion(resourceLocationRegionRegion)
	return o
}

// SetResourceLocationRegionRegion adds the resourceLocationRegionRegion to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceLocationRegionRegion(resourceLocationRegionRegion *string) {
	o.ResourceLocationRegionRegion = resourceLocationRegionRegion
}

// WithResourceType adds the resourceType to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceType(resourceType *string) *LogServiceGetArchiveParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithResourceUUID adds the resourceUUID to the log service get archive params
func (o *LogServiceGetArchiveParams) WithResourceUUID(resourceUUID *string) *LogServiceGetArchiveParams {
	o.SetResourceUUID(resourceUUID)
	return o
}

// SetResourceUUID adds the resourceUuid to the log service get archive params
func (o *LogServiceGetArchiveParams) SetResourceUUID(resourceUUID *string) {
	o.ResourceUUID = resourceUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceGetArchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ResourceDescription != nil {

		// query param resource.description
		var qrResourceDescription string

		if o.ResourceDescription != nil {
			qrResourceDescription = *o.ResourceDescription
		}
		qResourceDescription := qrResourceDescription
		if qResourceDescription != "" {

			if err := r.SetQueryParam("resource.description", qResourceDescription); err != nil {
				return err
			}
		}
	}

	// path param resource.id
	if err := r.SetPathParam("resource.id", o.ResourceID); err != nil {
		return err
	}

	// path param resource.location.organization_id
	if err := r.SetPathParam("resource.location.organization_id", o.ResourceLocationOrganizationID); err != nil {
		return err
	}

	// path param resource.location.project_id
	if err := r.SetPathParam("resource.location.project_id", o.ResourceLocationProjectID); err != nil {
		return err
	}

	if o.ResourceLocationRegionProvider != nil {

		// query param resource.location.region.provider
		var qrResourceLocationRegionProvider string

		if o.ResourceLocationRegionProvider != nil {
			qrResourceLocationRegionProvider = *o.ResourceLocationRegionProvider
		}
		qResourceLocationRegionProvider := qrResourceLocationRegionProvider
		if qResourceLocationRegionProvider != "" {

			if err := r.SetQueryParam("resource.location.region.provider", qResourceLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.ResourceLocationRegionRegion != nil {

		// query param resource.location.region.region
		var qrResourceLocationRegionRegion string

		if o.ResourceLocationRegionRegion != nil {
			qrResourceLocationRegionRegion = *o.ResourceLocationRegionRegion
		}
		qResourceLocationRegionRegion := qrResourceLocationRegionRegion
		if qResourceLocationRegionRegion != "" {

			if err := r.SetQueryParam("resource.location.region.region", qResourceLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if o.ResourceType != nil {

		// query param resource.type
		var qrResourceType string

		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {

			if err := r.SetQueryParam("resource.type", qResourceType); err != nil {
				return err
			}
		}
	}

	if o.ResourceUUID != nil {

		// query param resource.uuid
		var qrResourceUUID string

		if o.ResourceUUID != nil {
			qrResourceUUID = *o.ResourceUUID
		}
		qResourceUUID := qrResourceUUID
		if qResourceUUID != "" {

			if err := r.SetQueryParam("resource.uuid", qResourceUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
