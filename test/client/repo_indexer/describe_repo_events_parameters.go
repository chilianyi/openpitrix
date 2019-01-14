// Code generated by go-swagger; DO NOT EDIT.

package repo_indexer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeRepoEventsParams creates a new DescribeRepoEventsParams object
// with the default values initialized.
func NewDescribeRepoEventsParams() *DescribeRepoEventsParams {
	var ()
	return &DescribeRepoEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRepoEventsParamsWithTimeout creates a new DescribeRepoEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeRepoEventsParamsWithTimeout(timeout time.Duration) *DescribeRepoEventsParams {
	var ()
	return &DescribeRepoEventsParams{

		timeout: timeout,
	}
}

// NewDescribeRepoEventsParamsWithContext creates a new DescribeRepoEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeRepoEventsParamsWithContext(ctx context.Context) *DescribeRepoEventsParams {
	var ()
	return &DescribeRepoEventsParams{

		Context: ctx,
	}
}

// NewDescribeRepoEventsParamsWithHTTPClient creates a new DescribeRepoEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeRepoEventsParamsWithHTTPClient(client *http.Client) *DescribeRepoEventsParams {
	var ()
	return &DescribeRepoEventsParams{
		HTTPClient: client,
	}
}

/*DescribeRepoEventsParams contains all the parameters to send to the API endpoint
for the describe repo events operation typically these are written to a http.Request
*/
type DescribeRepoEventsParams struct {

	/*Limit
	  default is 20, max value is 200.

	*/
	Limit *int64
	/*Offset
	  default is 0.

	*/
	Offset *int64
	/*OwnerPath*/
	OwnerPath []string
	/*RepoEventID*/
	RepoEventID []string
	/*RepoID*/
	RepoID []string
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe repo events params
func (o *DescribeRepoEventsParams) WithTimeout(timeout time.Duration) *DescribeRepoEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe repo events params
func (o *DescribeRepoEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe repo events params
func (o *DescribeRepoEventsParams) WithContext(ctx context.Context) *DescribeRepoEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe repo events params
func (o *DescribeRepoEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe repo events params
func (o *DescribeRepoEventsParams) WithHTTPClient(client *http.Client) *DescribeRepoEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe repo events params
func (o *DescribeRepoEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe repo events params
func (o *DescribeRepoEventsParams) WithLimit(limit *int64) *DescribeRepoEventsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe repo events params
func (o *DescribeRepoEventsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe repo events params
func (o *DescribeRepoEventsParams) WithOffset(offset *int64) *DescribeRepoEventsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe repo events params
func (o *DescribeRepoEventsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerPath adds the ownerPath to the describe repo events params
func (o *DescribeRepoEventsParams) WithOwnerPath(ownerPath []string) *DescribeRepoEventsParams {
	o.SetOwnerPath(ownerPath)
	return o
}

// SetOwnerPath adds the ownerPath to the describe repo events params
func (o *DescribeRepoEventsParams) SetOwnerPath(ownerPath []string) {
	o.OwnerPath = ownerPath
}

// WithRepoEventID adds the repoEventID to the describe repo events params
func (o *DescribeRepoEventsParams) WithRepoEventID(repoEventID []string) *DescribeRepoEventsParams {
	o.SetRepoEventID(repoEventID)
	return o
}

// SetRepoEventID adds the repoEventId to the describe repo events params
func (o *DescribeRepoEventsParams) SetRepoEventID(repoEventID []string) {
	o.RepoEventID = repoEventID
}

// WithRepoID adds the repoID to the describe repo events params
func (o *DescribeRepoEventsParams) WithRepoID(repoID []string) *DescribeRepoEventsParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the describe repo events params
func (o *DescribeRepoEventsParams) SetRepoID(repoID []string) {
	o.RepoID = repoID
}

// WithStatus adds the status to the describe repo events params
func (o *DescribeRepoEventsParams) WithStatus(status []string) *DescribeRepoEventsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe repo events params
func (o *DescribeRepoEventsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRepoEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwnerPath := o.OwnerPath

	joinedOwnerPath := swag.JoinByFormat(valuesOwnerPath, "multi")
	// query array param owner_path
	if err := r.SetQueryParam("owner_path", joinedOwnerPath...); err != nil {
		return err
	}

	valuesRepoEventID := o.RepoEventID

	joinedRepoEventID := swag.JoinByFormat(valuesRepoEventID, "multi")
	// query array param repo_event_id
	if err := r.SetQueryParam("repo_event_id", joinedRepoEventID...); err != nil {
		return err
	}

	valuesRepoID := o.RepoID

	joinedRepoID := swag.JoinByFormat(valuesRepoID, "multi")
	// query array param repo_id
	if err := r.SetQueryParam("repo_id", joinedRepoID...); err != nil {
		return err
	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
