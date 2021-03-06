// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "mlss-controlcenter-go/pkg/models"
)

// NewUpdateUserGroupParams creates a new UpdateUserGroupParams object
// no default values defined in spec.
func NewUpdateUserGroupParams() UpdateUserGroupParams {

	return UpdateUserGroupParams{}
}

// UpdateUserGroupParams contains all the bound params for the update user group operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateUserGroup
type UpdateUserGroupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The group Request
	  Required: true
	  In: body
	*/
	UserGroup *models.UserGroup
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateUserGroupParams() beforehand.
func (o *UpdateUserGroupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserGroup
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("userGroup", "body"))
			} else {
				res = append(res, errors.NewParseError("userGroup", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UserGroup = &body
			}
		}
	} else {
		res = append(res, errors.Required("userGroup", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
