// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateEmployeeOKCode is the HTTP code returned for type UpdateEmployeeOK
const UpdateEmployeeOKCode int = 200

/*UpdateEmployeeOK Successfully Update

swagger:response updateEmployeeOK
*/
type UpdateEmployeeOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateEmployeeOK creates UpdateEmployeeOK with default headers values
func NewUpdateEmployeeOK() *UpdateEmployeeOK {

	return &UpdateEmployeeOK{}
}

// WithPayload adds the payload to the update employee o k response
func (o *UpdateEmployeeOK) WithPayload(payload string) *UpdateEmployeeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update employee o k response
func (o *UpdateEmployeeOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateEmployeeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateEmployeeBadRequestCode is the HTTP code returned for type UpdateEmployeeBadRequest
const UpdateEmployeeBadRequestCode int = 400

/*UpdateEmployeeBadRequest Bad Request

swagger:response updateEmployeeBadRequest
*/
type UpdateEmployeeBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateEmployeeBadRequest creates UpdateEmployeeBadRequest with default headers values
func NewUpdateEmployeeBadRequest() *UpdateEmployeeBadRequest {

	return &UpdateEmployeeBadRequest{}
}

// WithPayload adds the payload to the update employee bad request response
func (o *UpdateEmployeeBadRequest) WithPayload(payload string) *UpdateEmployeeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update employee bad request response
func (o *UpdateEmployeeBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateEmployeeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateEmployeeUnauthorizedCode is the HTTP code returned for type UpdateEmployeeUnauthorized
const UpdateEmployeeUnauthorizedCode int = 401

/*UpdateEmployeeUnauthorized Unotherized

swagger:response updateEmployeeUnauthorized
*/
type UpdateEmployeeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateEmployeeUnauthorized creates UpdateEmployeeUnauthorized with default headers values
func NewUpdateEmployeeUnauthorized() *UpdateEmployeeUnauthorized {

	return &UpdateEmployeeUnauthorized{}
}

// WithPayload adds the payload to the update employee unauthorized response
func (o *UpdateEmployeeUnauthorized) WithPayload(payload string) *UpdateEmployeeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update employee unauthorized response
func (o *UpdateEmployeeUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateEmployeeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateEmployeeNotFoundCode is the HTTP code returned for type UpdateEmployeeNotFound
const UpdateEmployeeNotFoundCode int = 404

/*UpdateEmployeeNotFound Not Found

swagger:response updateEmployeeNotFound
*/
type UpdateEmployeeNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateEmployeeNotFound creates UpdateEmployeeNotFound with default headers values
func NewUpdateEmployeeNotFound() *UpdateEmployeeNotFound {

	return &UpdateEmployeeNotFound{}
}

// WithPayload adds the payload to the update employee not found response
func (o *UpdateEmployeeNotFound) WithPayload(payload string) *UpdateEmployeeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update employee not found response
func (o *UpdateEmployeeNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateEmployeeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateEmployeeInternalServerErrorCode is the HTTP code returned for type UpdateEmployeeInternalServerError
const UpdateEmployeeInternalServerErrorCode int = 500

/*UpdateEmployeeInternalServerError Internal Server Error

swagger:response updateEmployeeInternalServerError
*/
type UpdateEmployeeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateEmployeeInternalServerError creates UpdateEmployeeInternalServerError with default headers values
func NewUpdateEmployeeInternalServerError() *UpdateEmployeeInternalServerError {

	return &UpdateEmployeeInternalServerError{}
}

// WithPayload adds the payload to the update employee internal server error response
func (o *UpdateEmployeeInternalServerError) WithPayload(payload string) *UpdateEmployeeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update employee internal server error response
func (o *UpdateEmployeeInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateEmployeeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
