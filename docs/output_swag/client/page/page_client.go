// Code generated by go-swagger; DO NOT EDIT.

package page

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new page API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for page API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPageAdmin(params *GetPageAdminParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPageAdminOK, error)

	GetPageHomepage(params *GetPageHomepageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPageHomepageAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetPageAdmin pings page

Admin page
*/
func (a *Client) GetPageAdmin(params *GetPageAdminParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPageAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageAdminParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPageAdmin",
		Method:             "GET",
		PathPattern:        "/page/admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPageAdminReader{formats: a.formats},
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
	success, ok := result.(*GetPageAdminOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPageAdmin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPageHomepage pings page

home page
*/
func (a *Client) GetPageHomepage(params *GetPageHomepageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPageHomepageAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageHomepageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPageHomepage",
		Method:             "GET",
		PathPattern:        "/page/homepage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPageHomepageReader{formats: a.formats},
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
	success, ok := result.(*GetPageHomepageAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPageHomepage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}