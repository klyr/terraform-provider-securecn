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

// NewDeleteCdRuleIDDeploymentRuleParams creates a new DeleteCdRuleIDDeploymentRuleParams object
// with the default values initialized.
func NewDeleteCdRuleIDDeploymentRuleParams() *DeleteCdRuleIDDeploymentRuleParams {
	var ()
	return &DeleteCdRuleIDDeploymentRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCdRuleIDDeploymentRuleParamsWithTimeout creates a new DeleteCdRuleIDDeploymentRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCdRuleIDDeploymentRuleParamsWithTimeout(timeout time.Duration) *DeleteCdRuleIDDeploymentRuleParams {
	var ()
	return &DeleteCdRuleIDDeploymentRuleParams{

		timeout: timeout,
	}
}

// NewDeleteCdRuleIDDeploymentRuleParamsWithContext creates a new DeleteCdRuleIDDeploymentRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCdRuleIDDeploymentRuleParamsWithContext(ctx context.Context) *DeleteCdRuleIDDeploymentRuleParams {
	var ()
	return &DeleteCdRuleIDDeploymentRuleParams{

		Context: ctx,
	}
}

// NewDeleteCdRuleIDDeploymentRuleParamsWithHTTPClient creates a new DeleteCdRuleIDDeploymentRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCdRuleIDDeploymentRuleParamsWithHTTPClient(client *http.Client) *DeleteCdRuleIDDeploymentRuleParams {
	var ()
	return &DeleteCdRuleIDDeploymentRuleParams{
		HTTPClient: client,
	}
}

/*DeleteCdRuleIDDeploymentRuleParams contains all the parameters to send to the API endpoint
for the delete cd rule ID deployment rule operation typically these are written to a http.Request
*/
type DeleteCdRuleIDDeploymentRuleParams struct {

	/*RuleID
	  ruleId (uid)

	*/
	RuleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) WithTimeout(timeout time.Duration) *DeleteCdRuleIDDeploymentRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) WithContext(ctx context.Context) *DeleteCdRuleIDDeploymentRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) WithHTTPClient(client *http.Client) *DeleteCdRuleIDDeploymentRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) WithRuleID(ruleID strfmt.UUID) *DeleteCdRuleIDDeploymentRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete cd rule ID deployment rule params
func (o *DeleteCdRuleIDDeploymentRuleParams) SetRuleID(ruleID strfmt.UUID) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCdRuleIDDeploymentRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
