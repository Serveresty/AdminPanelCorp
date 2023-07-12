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

// PostAuthLoginFormReader is a Reader for the PostAuthLoginForm structure.
type PostAuthLoginFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthLoginFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthLoginFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthLoginFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthLoginFormInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/login-form] PostAuthLoginForm", response, response.Code())
	}
}

// NewPostAuthLoginFormOK creates a PostAuthLoginFormOK with default headers values
func NewPostAuthLoginFormOK() *PostAuthLoginFormOK {
	return &PostAuthLoginFormOK{}
}

/*
PostAuthLoginFormOK describes a response with status code 200, with default header values.

success
*/
type PostAuthLoginFormOK struct {

	/* Authorization
	 */
	Authorization string

	Payload string
}

// IsSuccess returns true when this post auth login form o k response has a 2xx status code
func (o *PostAuthLoginFormOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post auth login form o k response has a 3xx status code
func (o *PostAuthLoginFormOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth login form o k response has a 4xx status code
func (o *PostAuthLoginFormOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth login form o k response has a 5xx status code
func (o *PostAuthLoginFormOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth login form o k response a status code equal to that given
func (o *PostAuthLoginFormOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post auth login form o k response
func (o *PostAuthLoginFormOK) Code() int {
	return 200
}

func (o *PostAuthLoginFormOK) Error() string {
	return fmt.Sprintf("[POST /auth/login-form][%d] postAuthLoginFormOK  %+v", 200, o.Payload)
}

func (o *PostAuthLoginFormOK) String() string {
	return fmt.Sprintf("[POST /auth/login-form][%d] postAuthLoginFormOK  %+v", 200, o.Payload)
}

func (o *PostAuthLoginFormOK) GetPayload() string {
	return o.Payload
}

func (o *PostAuthLoginFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Authorization
	hdrAuthorization := response.GetHeader("Authorization")

	if hdrAuthorization != "" {
		o.Authorization = hdrAuthorization
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthLoginFormBadRequest creates a PostAuthLoginFormBadRequest with default headers values
func NewPostAuthLoginFormBadRequest() *PostAuthLoginFormBadRequest {
	return &PostAuthLoginFormBadRequest{}
}

/*
PostAuthLoginFormBadRequest describes a response with status code 400, with default header values.

error
*/
type PostAuthLoginFormBadRequest struct {

	/* Authorization
	 */
	Authorization string

	Payload string
}

// IsSuccess returns true when this post auth login form bad request response has a 2xx status code
func (o *PostAuthLoginFormBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post auth login form bad request response has a 3xx status code
func (o *PostAuthLoginFormBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth login form bad request response has a 4xx status code
func (o *PostAuthLoginFormBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post auth login form bad request response has a 5xx status code
func (o *PostAuthLoginFormBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth login form bad request response a status code equal to that given
func (o *PostAuthLoginFormBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post auth login form bad request response
func (o *PostAuthLoginFormBadRequest) Code() int {
	return 400
}

func (o *PostAuthLoginFormBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/login-form][%d] postAuthLoginFormBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthLoginFormBadRequest) String() string {
	return fmt.Sprintf("[POST /auth/login-form][%d] postAuthLoginFormBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthLoginFormBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostAuthLoginFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Authorization
	hdrAuthorization := response.GetHeader("Authorization")

	if hdrAuthorization != "" {
		o.Authorization = hdrAuthorization
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthLoginFormInternalServerError creates a PostAuthLoginFormInternalServerError with default headers values
func NewPostAuthLoginFormInternalServerError() *PostAuthLoginFormInternalServerError {
	return &PostAuthLoginFormInternalServerError{}
}

/*
PostAuthLoginFormInternalServerError describes a response with status code 500, with default header values.

error
*/
type PostAuthLoginFormInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post auth login form internal server error response has a 2xx status code
func (o *PostAuthLoginFormInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post auth login form internal server error response has a 3xx status code
func (o *PostAuthLoginFormInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth login form internal server error response has a 4xx status code
func (o *PostAuthLoginFormInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth login form internal server error response has a 5xx status code
func (o *PostAuthLoginFormInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post auth login form internal server error response a status code equal to that given
func (o *PostAuthLoginFormInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post auth login form internal server error response
func (o *PostAuthLoginFormInternalServerError) Code() int {
	return 500
}

func (o *PostAuthLoginFormInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/login-form][%d] postAuthLoginFormInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthLoginFormInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/login-form][%d] postAuthLoginFormInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthLoginFormInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostAuthLoginFormInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}