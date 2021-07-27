// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new packer service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for packer service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBucket(params *CreateBucketParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBucketOK, error)

	CreateBuild(params *CreateBuildParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBuildOK, error)

	CreateChannel(params *CreateChannelParams, authInfo runtime.ClientAuthInfoWriter) (*CreateChannelOK, error)

	CreateIteration(params *CreateIterationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIterationOK, error)

	DeleteBucket(params *DeleteBucketParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBucketOK, error)

	DeleteBuild(params *DeleteBuildParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBuildOK, error)

	DeleteChannel(params *DeleteChannelParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteChannelOK, error)

	DeleteIteration(params *DeleteIterationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIterationOK, error)

	GetAncestorImages(params *GetAncestorImagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAncestorImagesOK, error)

	GetBucket(params *GetBucketParams, authInfo runtime.ClientAuthInfoWriter) (*GetBucketOK, error)

	GetBuild(params *GetBuildParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildOK, error)

	GetChannel(params *GetChannelParams, authInfo runtime.ClientAuthInfoWriter) (*GetChannelOK, error)

	GetChildImages(params *GetChildImagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetChildImagesOK, error)

	GetIteration(params *GetIterationParams, authInfo runtime.ClientAuthInfoWriter) (*GetIterationOK, error)

	ListBuckets(params *ListBucketsParams, authInfo runtime.ClientAuthInfoWriter) (*ListBucketsOK, error)

	ListBuilds(params *ListBuildsParams, authInfo runtime.ClientAuthInfoWriter) (*ListBuildsOK, error)

	ListChannels(params *ListChannelsParams, authInfo runtime.ClientAuthInfoWriter) (*ListChannelsOK, error)

	ListIterations(params *ListIterationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListIterationsOK, error)

	UpdateBucket(params *UpdateBucketParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBucketOK, error)

	UpdateBuild(params *UpdateBuildParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBuildOK, error)

	UpdateChannel(params *UpdateChannelParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateChannelOK, error)

	UpdateIteration(params *UpdateIterationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIterationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBucket operations realted to images
*/
func (a *Client) CreateBucket(params *CreateBucketParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBucketParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateBucket",
		Method:             "PUT",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateBuild create build API
*/
func (a *Client) CreateBuild(params *CreateBuildParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateBuild",
		Method:             "POST",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{build.iteration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateChannel operations related to channels mutable aliases pointers for an image iteration
*/
func (a *Client) CreateChannel(params *CreateChannelParams, authInfo runtime.ClientAuthInfoWriter) (*CreateChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateChannel",
		Method:             "POST",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChannelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateChannelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateIteration operations related to image iterations
*/
func (a *Client) CreateIteration(params *CreateIterationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIterationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIterationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateIteration",
		Method:             "POST",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateIterationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateIterationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateIterationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteBucket delete bucket API
*/
func (a *Client) DeleteBucket(params *DeleteBucketParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBucketParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteBucket",
		Method:             "DELETE",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteBuild delete build API
*/
func (a *Client) DeleteBuild(params *DeleteBuildParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteBuild",
		Method:             "DELETE",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteChannel delete channel API
*/
func (a *Client) DeleteChannel(params *DeleteChannelParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteChannel",
		Method:             "DELETE",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteChannelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteChannelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteIteration delete iteration API
*/
func (a *Client) DeleteIteration(params *DeleteIterationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIterationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIterationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIteration",
		Method:             "DELETE",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteIterationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIterationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIterationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAncestorImages APIs endpoints to ease UI implementation
*/
func (a *Client) GetAncestorImages(params *GetAncestorImagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAncestorImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAncestorImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAncestorImages",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{incremental_version}/ancestors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAncestorImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAncestorImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAncestorImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBucket get bucket API
*/
func (a *Client) GetBucket(params *GetBucketParams, authInfo runtime.ClientAuthInfoWriter) (*GetBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBucketParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBucket",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBuild ts o d o make it possible to query by iteration incremental version not just u l ID
*/
func (a *Client) GetBuild(params *GetBuildParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBuild",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetChannel get channel API
*/
func (a *Client) GetChannel(params *GetChannelParams, authInfo runtime.ClientAuthInfoWriter) (*GetChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChannel",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetChannelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetChannelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetChildImages get child images API
*/
func (a *Client) GetChildImages(params *GetChildImagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetChildImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChildImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChildImages",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{incremental_version}/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetChildImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetChildImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetChildImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetIteration gets iteration allows to get an iteration by iteration id incremental version fingerprint these are supplied as a query parameter e g

  images/mybucket/iteration?iteration_id=fingerprint

bucket_slug must always be set because it is possible for iterations to
have the same incremental_version or fingerprint across buckets
*/
func (a *Client) GetIteration(params *GetIterationParams, authInfo runtime.ClientAuthInfoWriter) (*GetIterationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIterationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIteration",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iteration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIterationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIterationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIterationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListBuckets list buckets API
*/
func (a *Client) ListBuckets(params *ListBucketsParams, authInfo runtime.ClientAuthInfoWriter) (*ListBucketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBucketsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListBuckets",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBucketsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBucketsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListBucketsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListBuilds list builds API
*/
func (a *Client) ListBuilds(params *ListBuildsParams, authInfo runtime.ClientAuthInfoWriter) (*ListBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListBuilds",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{iteration_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBuildsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBuildsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListBuildsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListChannels list channels API
*/
func (a *Client) ListChannels(params *ListChannelsParams, authInfo runtime.ClientAuthInfoWriter) (*ListChannelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListChannelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListChannels",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListChannelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListChannelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListChannelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListIterations list iterations API
*/
func (a *Client) ListIterations(params *ListIterationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListIterationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListIterationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListIterations",
		Method:             "GET",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListIterationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListIterationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListIterationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateBucket update bucket API
*/
func (a *Client) UpdateBucket(params *UpdateBucketParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBucketParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateBucket",
		Method:             "PATCH",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateBuild update build API
*/
func (a *Client) UpdateBuild(params *UpdateBuildParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateBuild",
		Method:             "PATCH",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateChannel update channel API
*/
func (a *Client) UpdateChannel(params *UpdateChannelParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateChannel",
		Method:             "PATCH",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateChannelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateChannelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateIteration onlies used to set incremental version once all builds are complete otherwise update build is used for specific build updates
*/
func (a *Client) UpdateIteration(params *UpdateIterationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIterationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIterationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateIteration",
		Method:             "PATCH",
		PathPattern:        "/packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateIterationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateIterationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateIterationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
