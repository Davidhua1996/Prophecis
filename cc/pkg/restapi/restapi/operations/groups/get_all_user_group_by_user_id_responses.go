// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetAllUserGroupByUserIDOKCode is the HTTP code returned for type GetAllUserGroupByUserIDOK
const GetAllUserGroupByUserIDOKCode int = 200

/*GetAllUserGroupByUserIDOK Detailed userGroup and userGroup information.

swagger:response getAllUserGroupByUserIdOK
*/
type GetAllUserGroupByUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserGroup `json:"body,omitempty"`
}

// NewGetAllUserGroupByUserIDOK creates GetAllUserGroupByUserIDOK with default headers values
func NewGetAllUserGroupByUserIDOK() *GetAllUserGroupByUserIDOK {

	return &GetAllUserGroupByUserIDOK{}
}

// WithPayload adds the payload to the get all user group by user Id o k response
func (o *GetAllUserGroupByUserIDOK) WithPayload(payload *models.UserGroup) *GetAllUserGroupByUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all user group by user Id o k response
func (o *GetAllUserGroupByUserIDOK) SetPayload(payload *models.UserGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllUserGroupByUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllUserGroupByUserIDUnauthorizedCode is the HTTP code returned for type GetAllUserGroupByUserIDUnauthorized
const GetAllUserGroupByUserIDUnauthorizedCode int = 401

/*GetAllUserGroupByUserIDUnauthorized Unauthorized

swagger:response getAllUserGroupByUserIdUnauthorized
*/
type GetAllUserGroupByUserIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllUserGroupByUserIDUnauthorized creates GetAllUserGroupByUserIDUnauthorized with default headers values
func NewGetAllUserGroupByUserIDUnauthorized() *GetAllUserGroupByUserIDUnauthorized {

	return &GetAllUserGroupByUserIDUnauthorized{}
}

// WithPayload adds the payload to the get all user group by user Id unauthorized response
func (o *GetAllUserGroupByUserIDUnauthorized) WithPayload(payload *models.Error) *GetAllUserGroupByUserIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all user group by user Id unauthorized response
func (o *GetAllUserGroupByUserIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllUserGroupByUserIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllUserGroupByUserIDNotFoundCode is the HTTP code returned for type GetAllUserGroupByUserIDNotFound
const GetAllUserGroupByUserIDNotFoundCode int = 404

/*GetAllUserGroupByUserIDNotFound Model with the given ID not found.

swagger:response getAllUserGroupByUserIdNotFound
*/
type GetAllUserGroupByUserIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllUserGroupByUserIDNotFound creates GetAllUserGroupByUserIDNotFound with default headers values
func NewGetAllUserGroupByUserIDNotFound() *GetAllUserGroupByUserIDNotFound {

	return &GetAllUserGroupByUserIDNotFound{}
}

// WithPayload adds the payload to the get all user group by user Id not found response
func (o *GetAllUserGroupByUserIDNotFound) WithPayload(payload *models.Error) *GetAllUserGroupByUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all user group by user Id not found response
func (o *GetAllUserGroupByUserIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllUserGroupByUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
