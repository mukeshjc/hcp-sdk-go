// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewListClustersParams creates a new ListClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListClustersParams() *ListClustersParams {
	return &ListClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListClustersParamsWithTimeout creates a new ListClustersParams object
// with the ability to set a timeout on a request.
func NewListClustersParamsWithTimeout(timeout time.Duration) *ListClustersParams {
	return &ListClustersParams{
		timeout: timeout,
	}
}

// NewListClustersParamsWithContext creates a new ListClustersParams object
// with the ability to set a context for a request.
func NewListClustersParamsWithContext(ctx context.Context) *ListClustersParams {
	return &ListClustersParams{
		Context: ctx,
	}
}

// NewListClustersParamsWithHTTPClient creates a new ListClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListClustersParamsWithHTTPClient(client *http.Client) *ListClustersParams {
	return &ListClustersParams{
		HTTPClient: client,
	}
}

/*
ListClustersParams contains all the parameters to send to the API endpoint

	for the list clusters operation.

	Typically these are written to a http.Request.
*/
type ListClustersParams struct {

	/* Name.

	   Query param filter: `name` of the cluster
	*/
	Name *string

	/* OrderBy.

	   Sorts the services based on a field. Allowed fields: `name`, `state`, `created_at`. The value needs to be of the format <{name/state/created_at} {asc/desc}>. For example: `name asc`, `state desc`.
	*/
	OrderBy []string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this parameter to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The maximum number of results per page to return. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned, which you can use to get the next page of results in subsequent
	requests. A value of zero causes `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this parameter to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* ParentResourceName.

	   Parent Resource name will be of the form `project/<project_id>`
	*/
	ParentResourceName string

	/* State.

	   Query param filter: `state` of the cluster. This can be combination of `creating`, 'pending', `running`, `disconnected`.
	*/
	State []string

	/* Type.

	   Query param filter: `type` of the cluster. This can be combination of `self_managed`, `hcp-managed`.
	*/
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClustersParams) WithDefaults() *ListClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list clusters params
func (o *ListClustersParams) WithTimeout(timeout time.Duration) *ListClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list clusters params
func (o *ListClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list clusters params
func (o *ListClustersParams) WithContext(ctx context.Context) *ListClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list clusters params
func (o *ListClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list clusters params
func (o *ListClustersParams) WithHTTPClient(client *http.Client) *ListClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list clusters params
func (o *ListClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the list clusters params
func (o *ListClustersParams) WithName(name *string) *ListClustersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list clusters params
func (o *ListClustersParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the list clusters params
func (o *ListClustersParams) WithOrderBy(orderBy []string) *ListClustersParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list clusters params
func (o *ListClustersParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list clusters params
func (o *ListClustersParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListClustersParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list clusters params
func (o *ListClustersParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list clusters params
func (o *ListClustersParams) WithPaginationPageSize(paginationPageSize *int64) *ListClustersParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list clusters params
func (o *ListClustersParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list clusters params
func (o *ListClustersParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListClustersParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list clusters params
func (o *ListClustersParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithParentResourceName adds the parentResourceName to the list clusters params
func (o *ListClustersParams) WithParentResourceName(parentResourceName string) *ListClustersParams {
	o.SetParentResourceName(parentResourceName)
	return o
}

// SetParentResourceName adds the parentResourceName to the list clusters params
func (o *ListClustersParams) SetParentResourceName(parentResourceName string) {
	o.ParentResourceName = parentResourceName
}

// WithState adds the state to the list clusters params
func (o *ListClustersParams) WithState(state []string) *ListClustersParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the list clusters params
func (o *ListClustersParams) SetState(state []string) {
	o.State = state
}

// WithType adds the typeVar to the list clusters params
func (o *ListClustersParams) WithType(typeVar []string) *ListClustersParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list clusters params
func (o *ListClustersParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
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

	// path param parent_resource_name
	if err := r.SetPathParam("parent_resource_name", o.ParentResourceName); err != nil {
		return err
	}

	if o.State != nil {

		// binding items for state
		joinedState := o.bindParamState(reg)

		// query array param state
		if err := r.SetQueryParam("state", joinedState...); err != nil {
			return err
		}
	}

	if o.Type != nil {

		// binding items for type
		joinedType := o.bindParamType(reg)

		// query array param type
		if err := r.SetQueryParam("type", joinedType...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListClusters binds the parameter order_by
func (o *ListClustersParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "multi"
	orderByIS := swag.JoinByFormat(orderByIC, "multi")

	return orderByIS
}

// bindParamListClusters binds the parameter state
func (o *ListClustersParams) bindParamState(formats strfmt.Registry) []string {
	stateIR := o.State

	var stateIC []string
	for _, stateIIR := range stateIR { // explode []string

		stateIIV := stateIIR // string as string
		stateIC = append(stateIC, stateIIV)
	}

	// items.CollectionFormat: "multi"
	stateIS := swag.JoinByFormat(stateIC, "multi")

	return stateIS
}

// bindParamListClusters binds the parameter type
func (o *ListClustersParams) bindParamType(formats strfmt.Registry) []string {
	typeIR := o.Type

	var typeIC []string
	for _, typeIIR := range typeIR { // explode []string

		typeIIV := typeIIR // string as string
		typeIC = append(typeIC, typeIIV)
	}

	// items.CollectionFormat: "multi"
	typeIS := swag.JoinByFormat(typeIC, "multi")

	return typeIS
}
