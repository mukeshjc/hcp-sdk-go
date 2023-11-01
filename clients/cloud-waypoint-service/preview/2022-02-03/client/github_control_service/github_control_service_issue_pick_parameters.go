// Code generated by go-swagger; DO NOT EDIT.

package github_control_service

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

// NewGithubControlServiceIssuePickParams creates a new GithubControlServiceIssuePickParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGithubControlServiceIssuePickParams() *GithubControlServiceIssuePickParams {
	return &GithubControlServiceIssuePickParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGithubControlServiceIssuePickParamsWithTimeout creates a new GithubControlServiceIssuePickParams object
// with the ability to set a timeout on a request.
func NewGithubControlServiceIssuePickParamsWithTimeout(timeout time.Duration) *GithubControlServiceIssuePickParams {
	return &GithubControlServiceIssuePickParams{
		timeout: timeout,
	}
}

// NewGithubControlServiceIssuePickParamsWithContext creates a new GithubControlServiceIssuePickParams object
// with the ability to set a context for a request.
func NewGithubControlServiceIssuePickParamsWithContext(ctx context.Context) *GithubControlServiceIssuePickParams {
	return &GithubControlServiceIssuePickParams{
		Context: ctx,
	}
}

// NewGithubControlServiceIssuePickParamsWithHTTPClient creates a new GithubControlServiceIssuePickParams object
// with the ability to set a custom HTTPClient for a request.
func NewGithubControlServiceIssuePickParamsWithHTTPClient(client *http.Client) *GithubControlServiceIssuePickParams {
	return &GithubControlServiceIssuePickParams{
		HTTPClient: client,
	}
}

/*
GithubControlServiceIssuePickParams contains all the parameters to send to the API endpoint

	for the github control service issue pick operation.

	Typically these are written to a http.Request.
*/
type GithubControlServiceIssuePickParams struct {

	// InstallationID.
	//
	// Format: int64
	InstallationID *string

	// RepoFullName.
	RepoFullName *string

	// SessionID.
	SessionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the github control service issue pick params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GithubControlServiceIssuePickParams) WithDefaults() *GithubControlServiceIssuePickParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the github control service issue pick params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GithubControlServiceIssuePickParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) WithTimeout(timeout time.Duration) *GithubControlServiceIssuePickParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) WithContext(ctx context.Context) *GithubControlServiceIssuePickParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) WithHTTPClient(client *http.Client) *GithubControlServiceIssuePickParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallationID adds the installationID to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) WithInstallationID(installationID *string) *GithubControlServiceIssuePickParams {
	o.SetInstallationID(installationID)
	return o
}

// SetInstallationID adds the installationId to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) SetInstallationID(installationID *string) {
	o.InstallationID = installationID
}

// WithRepoFullName adds the repoFullName to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) WithRepoFullName(repoFullName *string) *GithubControlServiceIssuePickParams {
	o.SetRepoFullName(repoFullName)
	return o
}

// SetRepoFullName adds the repoFullName to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) SetRepoFullName(repoFullName *string) {
	o.RepoFullName = repoFullName
}

// WithSessionID adds the sessionID to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) WithSessionID(sessionID *string) *GithubControlServiceIssuePickParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the github control service issue pick params
func (o *GithubControlServiceIssuePickParams) SetSessionID(sessionID *string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GithubControlServiceIssuePickParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InstallationID != nil {

		// query param installation_id
		var qrInstallationID string

		if o.InstallationID != nil {
			qrInstallationID = *o.InstallationID
		}
		qInstallationID := qrInstallationID
		if qInstallationID != "" {

			if err := r.SetQueryParam("installation_id", qInstallationID); err != nil {
				return err
			}
		}
	}

	if o.RepoFullName != nil {

		// query param repo_full_name
		var qrRepoFullName string

		if o.RepoFullName != nil {
			qrRepoFullName = *o.RepoFullName
		}
		qRepoFullName := qrRepoFullName
		if qRepoFullName != "" {

			if err := r.SetQueryParam("repo_full_name", qRepoFullName); err != nil {
				return err
			}
		}
	}

	if o.SessionID != nil {

		// query param session_id
		var qrSessionID string

		if o.SessionID != nil {
			qrSessionID = *o.SessionID
		}
		qSessionID := qrSessionID
		if qSessionID != "" {

			if err := r.SetQueryParam("session_id", qSessionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}