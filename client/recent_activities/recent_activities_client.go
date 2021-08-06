// Code generated by go-swagger; DO NOT EDIT.

package recent_activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new recent activities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for recent activities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetRecentActivities(params *GetRecentActivitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRecentActivitiesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetRecentActivities gets recent activities

  Get recent activities
*/
func (a *Client) GetRecentActivities(params *GetRecentActivitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRecentActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecentActivitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRecentActivities",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/stats/{aid}/recent-activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRecentActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRecentActivitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRecentActivities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
