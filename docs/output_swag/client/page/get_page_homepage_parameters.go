// Code generated by go-swagger; DO NOT EDIT.

package page

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

// NewGetPageHomepageParams creates a new GetPageHomepageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPageHomepageParams() *GetPageHomepageParams {
	return &GetPageHomepageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPageHomepageParamsWithTimeout creates a new GetPageHomepageParams object
// with the ability to set a timeout on a request.
func NewGetPageHomepageParamsWithTimeout(timeout time.Duration) *GetPageHomepageParams {
	return &GetPageHomepageParams{
		timeout: timeout,
	}
}

// NewGetPageHomepageParamsWithContext creates a new GetPageHomepageParams object
// with the ability to set a context for a request.
func NewGetPageHomepageParamsWithContext(ctx context.Context) *GetPageHomepageParams {
	return &GetPageHomepageParams{
		Context: ctx,
	}
}

// NewGetPageHomepageParamsWithHTTPClient creates a new GetPageHomepageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPageHomepageParamsWithHTTPClient(client *http.Client) *GetPageHomepageParams {
	return &GetPageHomepageParams{
		HTTPClient: client,
	}
}

/*
GetPageHomepageParams contains all the parameters to send to the API endpoint

	for the get page homepage operation.

	Typically these are written to a http.Request.
*/
type GetPageHomepageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get page homepage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPageHomepageParams) WithDefaults() *GetPageHomepageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get page homepage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPageHomepageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get page homepage params
func (o *GetPageHomepageParams) WithTimeout(timeout time.Duration) *GetPageHomepageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get page homepage params
func (o *GetPageHomepageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get page homepage params
func (o *GetPageHomepageParams) WithContext(ctx context.Context) *GetPageHomepageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get page homepage params
func (o *GetPageHomepageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get page homepage params
func (o *GetPageHomepageParams) WithHTTPClient(client *http.Client) *GetPageHomepageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get page homepage params
func (o *GetPageHomepageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPageHomepageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
