// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutEnvironmentsEnvIDParams creates a new PutEnvironmentsEnvIDParams object
// with the default values initialized.
func NewPutEnvironmentsEnvIDParams() *PutEnvironmentsEnvIDParams {
	var ()
	return &PutEnvironmentsEnvIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEnvironmentsEnvIDParamsWithTimeout creates a new PutEnvironmentsEnvIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEnvironmentsEnvIDParamsWithTimeout(timeout time.Duration) *PutEnvironmentsEnvIDParams {
	var ()
	return &PutEnvironmentsEnvIDParams{

		timeout: timeout,
	}
}

// NewPutEnvironmentsEnvIDParamsWithContext creates a new PutEnvironmentsEnvIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEnvironmentsEnvIDParamsWithContext(ctx context.Context) *PutEnvironmentsEnvIDParams {
	var ()
	return &PutEnvironmentsEnvIDParams{

		Context: ctx,
	}
}

// NewPutEnvironmentsEnvIDParamsWithHTTPClient creates a new PutEnvironmentsEnvIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEnvironmentsEnvIDParamsWithHTTPClient(client *http.Client) *PutEnvironmentsEnvIDParams {
	var ()
	return &PutEnvironmentsEnvIDParams{
		HTTPClient: client,
	}
}

/*PutEnvironmentsEnvIDParams contains all the parameters to send to the API endpoint
for the put environments env ID operation typically these are written to a http.Request
*/
type PutEnvironmentsEnvIDParams struct {

	/*Body
	  Environment definition

	*/
	Body *Environment
	/*EnvID*/
	EnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) WithTimeout(timeout time.Duration) *PutEnvironmentsEnvIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) WithContext(ctx context.Context) *PutEnvironmentsEnvIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) WithHTTPClient(client *http.Client) *PutEnvironmentsEnvIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) WithBody(body *Environment) *PutEnvironmentsEnvIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) SetBody(body *Environment) {
	o.Body = body
}

// WithEnvID adds the envID to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) WithEnvID(envID strfmt.UUID) *PutEnvironmentsEnvIDParams {
	o.SetEnvID(envID)
	return o
}

// SetEnvID adds the envId to the put environments env ID params
func (o *PutEnvironmentsEnvIDParams) SetEnvID(envID strfmt.UUID) {
	o.EnvID = envID
}

// WriteToRequest writes these params to a swagger request
func (o *PutEnvironmentsEnvIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param envId
	if err := r.SetPathParam("envId", o.EnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
