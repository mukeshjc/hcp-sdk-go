// Code generated by go-swagger; DO NOT EDIT.

package network_service

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

// NewGetHVNRouteParams creates a new GetHVNRouteParams object
// with the default values initialized.
func NewGetHVNRouteParams() *GetHVNRouteParams {
	var ()
	return &GetHVNRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHVNRouteParamsWithTimeout creates a new GetHVNRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHVNRouteParamsWithTimeout(timeout time.Duration) *GetHVNRouteParams {
	var ()
	return &GetHVNRouteParams{

		timeout: timeout,
	}
}

// NewGetHVNRouteParamsWithContext creates a new GetHVNRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHVNRouteParamsWithContext(ctx context.Context) *GetHVNRouteParams {
	var ()
	return &GetHVNRouteParams{

		Context: ctx,
	}
}

// NewGetHVNRouteParamsWithHTTPClient creates a new GetHVNRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHVNRouteParamsWithHTTPClient(client *http.Client) *GetHVNRouteParams {
	var ()
	return &GetHVNRouteParams{
		HTTPClient: client,
	}
}

/*GetHVNRouteParams contains all the parameters to send to the API endpoint
for the get h v n route operation typically these are written to a http.Request
*/
type GetHVNRouteParams struct {

	/*HvnDescription
	  description is a human-friendly description for this link. This is
	used primarily for informational purposes such as error messages.

	*/
	HvnDescription *string
	/*HvnID
	  id is the identifier for this resource.

	*/
	HvnID string
	/*HvnLocationOrganizationID
	  organization_id is the id of the organization.

	*/
	HvnLocationOrganizationID string
	/*HvnLocationProjectID
	  project_id is the projects id.

	*/
	HvnLocationProjectID string
	/*HvnLocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	HvnLocationRegionProvider *string
	/*HvnLocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	HvnLocationRegionRegion *string
	/*HvnType
	  type is the unique type of the resource. Each service publishes a
	unique set of types. The type value is recommended to be formatted
	in "<org>.<type>" such as "hashicorp.hvn". This is to prevent conflicts
	in the future, but any string value will work.

	*/
	HvnType *string
	/*HvnUUID
	  uuid is the unique UUID for this resource.

	*/
	HvnUUID *string
	/*ID
	  id of the HVN route to fetch

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get h v n route params
func (o *GetHVNRouteParams) WithTimeout(timeout time.Duration) *GetHVNRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get h v n route params
func (o *GetHVNRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get h v n route params
func (o *GetHVNRouteParams) WithContext(ctx context.Context) *GetHVNRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get h v n route params
func (o *GetHVNRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get h v n route params
func (o *GetHVNRouteParams) WithHTTPClient(client *http.Client) *GetHVNRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get h v n route params
func (o *GetHVNRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHvnDescription adds the hvnDescription to the get h v n route params
func (o *GetHVNRouteParams) WithHvnDescription(hvnDescription *string) *GetHVNRouteParams {
	o.SetHvnDescription(hvnDescription)
	return o
}

// SetHvnDescription adds the hvnDescription to the get h v n route params
func (o *GetHVNRouteParams) SetHvnDescription(hvnDescription *string) {
	o.HvnDescription = hvnDescription
}

// WithHvnID adds the hvnID to the get h v n route params
func (o *GetHVNRouteParams) WithHvnID(hvnID string) *GetHVNRouteParams {
	o.SetHvnID(hvnID)
	return o
}

// SetHvnID adds the hvnId to the get h v n route params
func (o *GetHVNRouteParams) SetHvnID(hvnID string) {
	o.HvnID = hvnID
}

// WithHvnLocationOrganizationID adds the hvnLocationOrganizationID to the get h v n route params
func (o *GetHVNRouteParams) WithHvnLocationOrganizationID(hvnLocationOrganizationID string) *GetHVNRouteParams {
	o.SetHvnLocationOrganizationID(hvnLocationOrganizationID)
	return o
}

// SetHvnLocationOrganizationID adds the hvnLocationOrganizationId to the get h v n route params
func (o *GetHVNRouteParams) SetHvnLocationOrganizationID(hvnLocationOrganizationID string) {
	o.HvnLocationOrganizationID = hvnLocationOrganizationID
}

// WithHvnLocationProjectID adds the hvnLocationProjectID to the get h v n route params
func (o *GetHVNRouteParams) WithHvnLocationProjectID(hvnLocationProjectID string) *GetHVNRouteParams {
	o.SetHvnLocationProjectID(hvnLocationProjectID)
	return o
}

// SetHvnLocationProjectID adds the hvnLocationProjectId to the get h v n route params
func (o *GetHVNRouteParams) SetHvnLocationProjectID(hvnLocationProjectID string) {
	o.HvnLocationProjectID = hvnLocationProjectID
}

// WithHvnLocationRegionProvider adds the hvnLocationRegionProvider to the get h v n route params
func (o *GetHVNRouteParams) WithHvnLocationRegionProvider(hvnLocationRegionProvider *string) *GetHVNRouteParams {
	o.SetHvnLocationRegionProvider(hvnLocationRegionProvider)
	return o
}

// SetHvnLocationRegionProvider adds the hvnLocationRegionProvider to the get h v n route params
func (o *GetHVNRouteParams) SetHvnLocationRegionProvider(hvnLocationRegionProvider *string) {
	o.HvnLocationRegionProvider = hvnLocationRegionProvider
}

// WithHvnLocationRegionRegion adds the hvnLocationRegionRegion to the get h v n route params
func (o *GetHVNRouteParams) WithHvnLocationRegionRegion(hvnLocationRegionRegion *string) *GetHVNRouteParams {
	o.SetHvnLocationRegionRegion(hvnLocationRegionRegion)
	return o
}

// SetHvnLocationRegionRegion adds the hvnLocationRegionRegion to the get h v n route params
func (o *GetHVNRouteParams) SetHvnLocationRegionRegion(hvnLocationRegionRegion *string) {
	o.HvnLocationRegionRegion = hvnLocationRegionRegion
}

// WithHvnType adds the hvnType to the get h v n route params
func (o *GetHVNRouteParams) WithHvnType(hvnType *string) *GetHVNRouteParams {
	o.SetHvnType(hvnType)
	return o
}

// SetHvnType adds the hvnType to the get h v n route params
func (o *GetHVNRouteParams) SetHvnType(hvnType *string) {
	o.HvnType = hvnType
}

// WithHvnUUID adds the hvnUUID to the get h v n route params
func (o *GetHVNRouteParams) WithHvnUUID(hvnUUID *string) *GetHVNRouteParams {
	o.SetHvnUUID(hvnUUID)
	return o
}

// SetHvnUUID adds the hvnUuid to the get h v n route params
func (o *GetHVNRouteParams) SetHvnUUID(hvnUUID *string) {
	o.HvnUUID = hvnUUID
}

// WithID adds the id to the get h v n route params
func (o *GetHVNRouteParams) WithID(id string) *GetHVNRouteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get h v n route params
func (o *GetHVNRouteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetHVNRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HvnDescription != nil {

		// query param hvn.description
		var qrHvnDescription string
		if o.HvnDescription != nil {
			qrHvnDescription = *o.HvnDescription
		}
		qHvnDescription := qrHvnDescription
		if qHvnDescription != "" {
			if err := r.SetQueryParam("hvn.description", qHvnDescription); err != nil {
				return err
			}
		}

	}

	// path param hvn.id
	if err := r.SetPathParam("hvn.id", o.HvnID); err != nil {
		return err
	}

	// path param hvn.location.organization_id
	if err := r.SetPathParam("hvn.location.organization_id", o.HvnLocationOrganizationID); err != nil {
		return err
	}

	// path param hvn.location.project_id
	if err := r.SetPathParam("hvn.location.project_id", o.HvnLocationProjectID); err != nil {
		return err
	}

	if o.HvnLocationRegionProvider != nil {

		// query param hvn.location.region.provider
		var qrHvnLocationRegionProvider string
		if o.HvnLocationRegionProvider != nil {
			qrHvnLocationRegionProvider = *o.HvnLocationRegionProvider
		}
		qHvnLocationRegionProvider := qrHvnLocationRegionProvider
		if qHvnLocationRegionProvider != "" {
			if err := r.SetQueryParam("hvn.location.region.provider", qHvnLocationRegionProvider); err != nil {
				return err
			}
		}

	}

	if o.HvnLocationRegionRegion != nil {

		// query param hvn.location.region.region
		var qrHvnLocationRegionRegion string
		if o.HvnLocationRegionRegion != nil {
			qrHvnLocationRegionRegion = *o.HvnLocationRegionRegion
		}
		qHvnLocationRegionRegion := qrHvnLocationRegionRegion
		if qHvnLocationRegionRegion != "" {
			if err := r.SetQueryParam("hvn.location.region.region", qHvnLocationRegionRegion); err != nil {
				return err
			}
		}

	}

	if o.HvnType != nil {

		// query param hvn.type
		var qrHvnType string
		if o.HvnType != nil {
			qrHvnType = *o.HvnType
		}
		qHvnType := qrHvnType
		if qHvnType != "" {
			if err := r.SetQueryParam("hvn.type", qHvnType); err != nil {
				return err
			}
		}

	}

	if o.HvnUUID != nil {

		// query param hvn.uuid
		var qrHvnUUID string
		if o.HvnUUID != nil {
			qrHvnUUID = *o.HvnUUID
		}
		qHvnUUID := qrHvnUUID
		if qHvnUUID != "" {
			if err := r.SetQueryParam("hvn.uuid", qHvnUUID); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
