// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetNamespaceByNameOKCode is the HTTP code returned for type GetNamespaceByNameOK
const GetNamespaceByNameOKCode int = 200

/*GetNamespaceByNameOK Detailed namespace and namespace information.

swagger:response getNamespaceByNameOK
*/
type GetNamespaceByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.NamespaceRequest `json:"body,omitempty"`
}

// NewGetNamespaceByNameOK creates GetNamespaceByNameOK with default headers values
func NewGetNamespaceByNameOK() *GetNamespaceByNameOK {

	return &GetNamespaceByNameOK{}
}

// WithPayload adds the payload to the get namespace by name o k response
func (o *GetNamespaceByNameOK) WithPayload(payload *models.NamespaceRequest) *GetNamespaceByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get namespace by name o k response
func (o *GetNamespaceByNameOK) SetPayload(payload *models.NamespaceRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNamespaceByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNamespaceByNameUnauthorizedCode is the HTTP code returned for type GetNamespaceByNameUnauthorized
const GetNamespaceByNameUnauthorizedCode int = 401

/*GetNamespaceByNameUnauthorized Unauthorized

swagger:response getNamespaceByNameUnauthorized
*/
type GetNamespaceByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNamespaceByNameUnauthorized creates GetNamespaceByNameUnauthorized with default headers values
func NewGetNamespaceByNameUnauthorized() *GetNamespaceByNameUnauthorized {

	return &GetNamespaceByNameUnauthorized{}
}

// WithPayload adds the payload to the get namespace by name unauthorized response
func (o *GetNamespaceByNameUnauthorized) WithPayload(payload *models.Error) *GetNamespaceByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get namespace by name unauthorized response
func (o *GetNamespaceByNameUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNamespaceByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNamespaceByNameNotFoundCode is the HTTP code returned for type GetNamespaceByNameNotFound
const GetNamespaceByNameNotFoundCode int = 404

/*GetNamespaceByNameNotFound Model with the given ID not found.

swagger:response getNamespaceByNameNotFound
*/
type GetNamespaceByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNamespaceByNameNotFound creates GetNamespaceByNameNotFound with default headers values
func NewGetNamespaceByNameNotFound() *GetNamespaceByNameNotFound {

	return &GetNamespaceByNameNotFound{}
}

// WithPayload adds the payload to the get namespace by name not found response
func (o *GetNamespaceByNameNotFound) WithPayload(payload *models.Error) *GetNamespaceByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get namespace by name not found response
func (o *GetNamespaceByNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNamespaceByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}