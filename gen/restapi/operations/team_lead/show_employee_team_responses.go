// Code generated by go-swagger; DO NOT EDIT.

package team_lead

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/models"
)

// ShowEmployeeTeamOKCode is the HTTP code returned for type ShowEmployeeTeamOK
const ShowEmployeeTeamOKCode int = 200

/*ShowEmployeeTeamOK Found

swagger:response showEmployeeTeamOK
*/
type ShowEmployeeTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.TeamLeadEmployee `json:"body,omitempty"`
}

// NewShowEmployeeTeamOK creates ShowEmployeeTeamOK with default headers values
func NewShowEmployeeTeamOK() *ShowEmployeeTeamOK {

	return &ShowEmployeeTeamOK{}
}

// WithPayload adds the payload to the show employee team o k response
func (o *ShowEmployeeTeamOK) WithPayload(payload *models.TeamLeadEmployee) *ShowEmployeeTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show employee team o k response
func (o *ShowEmployeeTeamOK) SetPayload(payload *models.TeamLeadEmployee) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowEmployeeTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowEmployeeTeamBadRequestCode is the HTTP code returned for type ShowEmployeeTeamBadRequest
const ShowEmployeeTeamBadRequestCode int = 400

/*ShowEmployeeTeamBadRequest Bad Request

swagger:response showEmployeeTeamBadRequest
*/
type ShowEmployeeTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewShowEmployeeTeamBadRequest creates ShowEmployeeTeamBadRequest with default headers values
func NewShowEmployeeTeamBadRequest() *ShowEmployeeTeamBadRequest {

	return &ShowEmployeeTeamBadRequest{}
}

// WithPayload adds the payload to the show employee team bad request response
func (o *ShowEmployeeTeamBadRequest) WithPayload(payload string) *ShowEmployeeTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show employee team bad request response
func (o *ShowEmployeeTeamBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowEmployeeTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ShowEmployeeTeamUnauthorizedCode is the HTTP code returned for type ShowEmployeeTeamUnauthorized
const ShowEmployeeTeamUnauthorizedCode int = 401

/*ShowEmployeeTeamUnauthorized Unotherized

swagger:response showEmployeeTeamUnauthorized
*/
type ShowEmployeeTeamUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewShowEmployeeTeamUnauthorized creates ShowEmployeeTeamUnauthorized with default headers values
func NewShowEmployeeTeamUnauthorized() *ShowEmployeeTeamUnauthorized {

	return &ShowEmployeeTeamUnauthorized{}
}

// WithPayload adds the payload to the show employee team unauthorized response
func (o *ShowEmployeeTeamUnauthorized) WithPayload(payload string) *ShowEmployeeTeamUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show employee team unauthorized response
func (o *ShowEmployeeTeamUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowEmployeeTeamUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ShowEmployeeTeamNotFoundCode is the HTTP code returned for type ShowEmployeeTeamNotFound
const ShowEmployeeTeamNotFoundCode int = 404

/*ShowEmployeeTeamNotFound Not Found

swagger:response showEmployeeTeamNotFound
*/
type ShowEmployeeTeamNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewShowEmployeeTeamNotFound creates ShowEmployeeTeamNotFound with default headers values
func NewShowEmployeeTeamNotFound() *ShowEmployeeTeamNotFound {

	return &ShowEmployeeTeamNotFound{}
}

// WithPayload adds the payload to the show employee team not found response
func (o *ShowEmployeeTeamNotFound) WithPayload(payload string) *ShowEmployeeTeamNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show employee team not found response
func (o *ShowEmployeeTeamNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowEmployeeTeamNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ShowEmployeeTeamInternalServerErrorCode is the HTTP code returned for type ShowEmployeeTeamInternalServerError
const ShowEmployeeTeamInternalServerErrorCode int = 500

/*ShowEmployeeTeamInternalServerError Internal Server Error

swagger:response showEmployeeTeamInternalServerError
*/
type ShowEmployeeTeamInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewShowEmployeeTeamInternalServerError creates ShowEmployeeTeamInternalServerError with default headers values
func NewShowEmployeeTeamInternalServerError() *ShowEmployeeTeamInternalServerError {

	return &ShowEmployeeTeamInternalServerError{}
}

// WithPayload adds the payload to the show employee team internal server error response
func (o *ShowEmployeeTeamInternalServerError) WithPayload(payload string) *ShowEmployeeTeamInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show employee team internal server error response
func (o *ShowEmployeeTeamInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowEmployeeTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
