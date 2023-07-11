// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAuthRegistrationFormReader is a Reader for the PostAuthRegistrationForm structure.
type PostAuthRegistrationFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthRegistrationFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthRegistrationFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthRegistrationFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthRegistrationFormInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/registration-form] PostAuthRegistrationForm", response, response.Code())
	}
}

// NewPostAuthRegistrationFormOK creates a PostAuthRegistrationFormOK with default headers values
func NewPostAuthRegistrationFormOK() *PostAuthRegistrationFormOK {
	return &PostAuthRegistrationFormOK{}
}

/*
PostAuthRegistrationFormOK describes a response with status code 200, with default header values.

success
*/
type PostAuthRegistrationFormOK struct {
	Payload string
}

// IsSuccess returns true when this post auth registration form o k response has a 2xx status code
func (o *PostAuthRegistrationFormOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post auth registration form o k response has a 3xx status code
func (o *PostAuthRegistrationFormOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth registration form o k response has a 4xx status code
func (o *PostAuthRegistrationFormOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth registration form o k response has a 5xx status code
func (o *PostAuthRegistrationFormOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth registration form o k response a status code equal to that given
func (o *PostAuthRegistrationFormOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post auth registration form o k response
func (o *PostAuthRegistrationFormOK) Code() int {
	return 200
}

func (o *PostAuthRegistrationFormOK) Error() string {
	return fmt.Sprintf("[POST /auth/registration-form][%d] postAuthRegistrationFormOK  %+v", 200, o.Payload)
}

func (o *PostAuthRegistrationFormOK) String() string {
	return fmt.Sprintf("[POST /auth/registration-form][%d] postAuthRegistrationFormOK  %+v", 200, o.Payload)
}

func (o *PostAuthRegistrationFormOK) GetPayload() string {
	return o.Payload
}

func (o *PostAuthRegistrationFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthRegistrationFormBadRequest creates a PostAuthRegistrationFormBadRequest with default headers values
func NewPostAuthRegistrationFormBadRequest() *PostAuthRegistrationFormBadRequest {
	return &PostAuthRegistrationFormBadRequest{}
}

/*
PostAuthRegistrationFormBadRequest describes a response with status code 400, with default header values.

error
*/
type PostAuthRegistrationFormBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post auth registration form bad request response has a 2xx status code
func (o *PostAuthRegistrationFormBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post auth registration form bad request response has a 3xx status code
func (o *PostAuthRegistrationFormBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth registration form bad request response has a 4xx status code
func (o *PostAuthRegistrationFormBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post auth registration form bad request response has a 5xx status code
func (o *PostAuthRegistrationFormBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth registration form bad request response a status code equal to that given
func (o *PostAuthRegistrationFormBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post auth registration form bad request response
func (o *PostAuthRegistrationFormBadRequest) Code() int {
	return 400
}

func (o *PostAuthRegistrationFormBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/registration-form][%d] postAuthRegistrationFormBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthRegistrationFormBadRequest) String() string {
	return fmt.Sprintf("[POST /auth/registration-form][%d] postAuthRegistrationFormBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthRegistrationFormBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostAuthRegistrationFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthRegistrationFormInternalServerError creates a PostAuthRegistrationFormInternalServerError with default headers values
func NewPostAuthRegistrationFormInternalServerError() *PostAuthRegistrationFormInternalServerError {
	return &PostAuthRegistrationFormInternalServerError{}
}

/*
PostAuthRegistrationFormInternalServerError describes a response with status code 500, with default header values.

error
*/
type PostAuthRegistrationFormInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post auth registration form internal server error response has a 2xx status code
func (o *PostAuthRegistrationFormInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post auth registration form internal server error response has a 3xx status code
func (o *PostAuthRegistrationFormInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth registration form internal server error response has a 4xx status code
func (o *PostAuthRegistrationFormInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth registration form internal server error response has a 5xx status code
func (o *PostAuthRegistrationFormInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post auth registration form internal server error response a status code equal to that given
func (o *PostAuthRegistrationFormInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post auth registration form internal server error response
func (o *PostAuthRegistrationFormInternalServerError) Code() int {
	return 500
}

func (o *PostAuthRegistrationFormInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/registration-form][%d] postAuthRegistrationFormInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthRegistrationFormInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/registration-form][%d] postAuthRegistrationFormInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthRegistrationFormInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostAuthRegistrationFormInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
