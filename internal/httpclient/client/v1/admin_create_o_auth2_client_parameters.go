// Code generated by go-swagger; DO NOT EDIT.

package v1

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

	"github.com/ory/hydra/internal/httpclient/client/v1"
	"github.com/ory/hydra/internal/httpclient/models"
)

// NewAdminCreateOAuth2ClientParams creates a new AdminCreateOAuth2ClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminCreateOAuth2ClientParams() *AdminCreateOAuth2ClientParams {
	return &AdminCreateOAuth2ClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateOAuth2ClientParamsWithTimeout creates a new AdminCreateOAuth2ClientParams object
// with the ability to set a timeout on a request.
func NewAdminCreateOAuth2ClientParamsWithTimeout(timeout time.Duration) *AdminCreateOAuth2ClientParams {
	return &AdminCreateOAuth2ClientParams{
		timeout: timeout,
	}
}

// NewAdminCreateOAuth2ClientParamsWithContext creates a new AdminCreateOAuth2ClientParams object
// with the ability to set a context for a request.
func NewAdminCreateOAuth2ClientParamsWithContext(ctx context.Context) *AdminCreateOAuth2ClientParams {
	return &AdminCreateOAuth2ClientParams{
		Context: ctx,
	}
}

// NewAdminCreateOAuth2ClientParamsWithHTTPClient creates a new AdminCreateOAuth2ClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminCreateOAuth2ClientParamsWithHTTPClient(client *http.Client) *AdminCreateOAuth2ClientParams {
	return &AdminCreateOAuth2ClientParams{
		HTTPClient: client,
	}
}

/* AdminCreateOAuth2ClientParams contains all the parameters to send to the API endpoint
   for the admin create o auth2 client operation.

   Typically these are written to a http.Request.
*/
type AdminCreateOAuth2ClientParams struct {

	// Body.
	Body *models.OAuth2Client

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin create o auth2 client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCreateOAuth2ClientParams) WithDefaults() *AdminCreateOAuth2ClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin create o auth2 client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCreateOAuth2ClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) WithTimeout(timeout time.Duration) *AdminCreateOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) WithContext(ctx context.Context) *AdminCreateOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) WithHTTPClient(client *http.Client) *AdminCreateOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) WithBody(body *models.OAuth2Client) *AdminCreateOAuth2ClientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin create o auth2 client params
func (o *AdminCreateOAuth2ClientParams) SetBody(body *models.OAuth2Client) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}