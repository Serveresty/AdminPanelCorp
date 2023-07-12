// Code generated by go-swagger; DO NOT EDIT.

package page

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPageAdminReader is a Reader for the GetPageAdmin structure.
type GetPageAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPageAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPageAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPageAdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPageAdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPageAdminInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /page/admin] GetPageAdmin", response, response.Code())
	}
}

// NewGetPageAdminOK creates a GetPageAdminOK with default headers values
func NewGetPageAdminOK() *GetPageAdminOK {
	return &GetPageAdminOK{}
}

/*
GetPageAdminOK describes a response with status code 200, with default header values.

OK
*/
type GetPageAdminOK struct {

	/* Authorization
	 */
	Authorization string

	Payload string
}

// IsSuccess returns true when this get page admin o k response has a 2xx status code
func (o *GetPageAdminOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get page admin o k response has a 3xx status code
func (o *GetPageAdminOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get page admin o k response has a 4xx status code
func (o *GetPageAdminOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get page admin o k response has a 5xx status code
func (o *GetPageAdminOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get page admin o k response a status code equal to that given
func (o *GetPageAdminOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get page admin o k response
func (o *GetPageAdminOK) Code() int {
	return 200
}

func (o *GetPageAdminOK) Error() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminOK  %+v", 200, o.Payload)
}

func (o *GetPageAdminOK) String() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminOK  %+v", 200, o.Payload)
}

func (o *GetPageAdminOK) GetPayload() string {
	return o.Payload
}

func (o *GetPageAdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPageAdminBadRequest creates a GetPageAdminBadRequest with default headers values
func NewGetPageAdminBadRequest() *GetPageAdminBadRequest {
	return &GetPageAdminBadRequest{}
}

/*
GetPageAdminBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPageAdminBadRequest struct {

	/* Authorization
	 */
	Authorization string

	Payload string
}

// IsSuccess returns true when this get page admin bad request response has a 2xx status code
func (o *GetPageAdminBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get page admin bad request response has a 3xx status code
func (o *GetPageAdminBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get page admin bad request response has a 4xx status code
func (o *GetPageAdminBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get page admin bad request response has a 5xx status code
func (o *GetPageAdminBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get page admin bad request response a status code equal to that given
func (o *GetPageAdminBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get page admin bad request response
func (o *GetPageAdminBadRequest) Code() int {
	return 400
}

func (o *GetPageAdminBadRequest) Error() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminBadRequest  %+v", 400, o.Payload)
}

func (o *GetPageAdminBadRequest) String() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminBadRequest  %+v", 400, o.Payload)
}

func (o *GetPageAdminBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetPageAdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPageAdminForbidden creates a GetPageAdminForbidden with default headers values
func NewGetPageAdminForbidden() *GetPageAdminForbidden {
	return &GetPageAdminForbidden{}
}

/*
GetPageAdminForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPageAdminForbidden struct {

	/* Authorization
	 */
	Authorization string

	Payload string
}

// IsSuccess returns true when this get page admin forbidden response has a 2xx status code
func (o *GetPageAdminForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get page admin forbidden response has a 3xx status code
func (o *GetPageAdminForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get page admin forbidden response has a 4xx status code
func (o *GetPageAdminForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get page admin forbidden response has a 5xx status code
func (o *GetPageAdminForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get page admin forbidden response a status code equal to that given
func (o *GetPageAdminForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get page admin forbidden response
func (o *GetPageAdminForbidden) Code() int {
	return 403
}

func (o *GetPageAdminForbidden) Error() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminForbidden  %+v", 403, o.Payload)
}

func (o *GetPageAdminForbidden) String() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminForbidden  %+v", 403, o.Payload)
}

func (o *GetPageAdminForbidden) GetPayload() string {
	return o.Payload
}

func (o *GetPageAdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPageAdminInternalServerError creates a GetPageAdminInternalServerError with default headers values
func NewGetPageAdminInternalServerError() *GetPageAdminInternalServerError {
	return &GetPageAdminInternalServerError{}
}

/*
GetPageAdminInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPageAdminInternalServerError struct {

	/* Authorization
	 */
	Authorization string

	Payload string
}

// IsSuccess returns true when this get page admin internal server error response has a 2xx status code
func (o *GetPageAdminInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get page admin internal server error response has a 3xx status code
func (o *GetPageAdminInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get page admin internal server error response has a 4xx status code
func (o *GetPageAdminInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get page admin internal server error response has a 5xx status code
func (o *GetPageAdminInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get page admin internal server error response a status code equal to that given
func (o *GetPageAdminInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get page admin internal server error response
func (o *GetPageAdminInternalServerError) Code() int {
	return 500
}

func (o *GetPageAdminInternalServerError) Error() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPageAdminInternalServerError) String() string {
	return fmt.Sprintf("[GET /page/admin][%d] getPageAdminInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPageAdminInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetPageAdminInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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