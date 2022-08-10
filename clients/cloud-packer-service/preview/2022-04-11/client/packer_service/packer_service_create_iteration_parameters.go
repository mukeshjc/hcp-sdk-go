// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/models"
)

// NewPackerServiceCreateIterationParams creates a new PackerServiceCreateIterationParams object
// with the default values initialized.
func NewPackerServiceCreateIterationParams() *PackerServiceCreateIterationParams {
	var ()
	return &PackerServiceCreateIterationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceCreateIterationParamsWithTimeout creates a new PackerServiceCreateIterationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceCreateIterationParamsWithTimeout(timeout time.Duration) *PackerServiceCreateIterationParams {
	var ()
	return &PackerServiceCreateIterationParams{

		timeout: timeout,
	}
}

// NewPackerServiceCreateIterationParamsWithContext creates a new PackerServiceCreateIterationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceCreateIterationParamsWithContext(ctx context.Context) *PackerServiceCreateIterationParams {
	var ()
	return &PackerServiceCreateIterationParams{

		Context: ctx,
	}
}

// NewPackerServiceCreateIterationParamsWithHTTPClient creates a new PackerServiceCreateIterationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceCreateIterationParamsWithHTTPClient(client *http.Client) *PackerServiceCreateIterationParams {
	var ()
	return &PackerServiceCreateIterationParams{
		HTTPClient: client,
	}
}

/*PackerServiceCreateIterationParams contains all the parameters to send to the API endpoint
for the packer service create iteration operation typically these are written to a http.Request
*/
type PackerServiceCreateIterationParams struct {

	/*Body*/
	Body *models.HashicorpCloudPacker20220411CreateIterationRequest
	/*BucketSlug
	  Human-readable name for the bucket.

	*/
	BucketSlug string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithTimeout(timeout time.Duration) *PackerServiceCreateIterationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithContext(ctx context.Context) *PackerServiceCreateIterationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithHTTPClient(client *http.Client) *PackerServiceCreateIterationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithBody(body *models.HashicorpCloudPacker20220411CreateIterationRequest) *PackerServiceCreateIterationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetBody(body *models.HashicorpCloudPacker20220411CreateIterationRequest) {
	o.Body = body
}

// WithBucketSlug adds the bucketSlug to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithBucketSlug(bucketSlug string) *PackerServiceCreateIterationParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceCreateIterationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) WithLocationProjectID(locationProjectID string) *PackerServiceCreateIterationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service create iteration params
func (o *PackerServiceCreateIterationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceCreateIterationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
