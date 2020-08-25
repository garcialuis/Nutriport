// Code generated by go-swagger; DO NOT EDIT.

package food_item

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

// NewCreateFoodItemParams creates a new CreateFoodItemParams object
// with the default values initialized.
func NewCreateFoodItemParams() *CreateFoodItemParams {

	return &CreateFoodItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFoodItemParamsWithTimeout creates a new CreateFoodItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFoodItemParamsWithTimeout(timeout time.Duration) *CreateFoodItemParams {

	return &CreateFoodItemParams{

		timeout: timeout,
	}
}

// NewCreateFoodItemParamsWithContext creates a new CreateFoodItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFoodItemParamsWithContext(ctx context.Context) *CreateFoodItemParams {

	return &CreateFoodItemParams{

		Context: ctx,
	}
}

// NewCreateFoodItemParamsWithHTTPClient creates a new CreateFoodItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFoodItemParamsWithHTTPClient(client *http.Client) *CreateFoodItemParams {

	return &CreateFoodItemParams{
		HTTPClient: client,
	}
}

/*CreateFoodItemParams contains all the parameters to send to the API endpoint
for the create food item operation typically these are written to a http.Request
*/
type CreateFoodItemParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create food item params
func (o *CreateFoodItemParams) WithTimeout(timeout time.Duration) *CreateFoodItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create food item params
func (o *CreateFoodItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create food item params
func (o *CreateFoodItemParams) WithContext(ctx context.Context) *CreateFoodItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create food item params
func (o *CreateFoodItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create food item params
func (o *CreateFoodItemParams) WithHTTPClient(client *http.Client) *CreateFoodItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create food item params
func (o *CreateFoodItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFoodItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
