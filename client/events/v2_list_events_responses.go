// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/openshift/assisted-service/models"
)

// V2ListEventsReader is a Reader for the V2ListEvents structure.
type V2ListEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2ListEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2ListEventsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListEventsOK creates a V2ListEventsOK with default headers values
func NewV2ListEventsOK() *V2ListEventsOK {
	return &V2ListEventsOK{}
}

/*
V2ListEventsOK describes a response with status code 200, with default header values.

Success.
*/
type V2ListEventsOK struct {

	/* Count of events with severity 'critical'.
	 */
	SeverityCountCritical int64

	/* Count of events with severity 'error'.
	 */
	SeverityCountError int64

	/* Count of events with severity 'info'.
	 */
	SeverityCountInfo int64

	/* Count of events with severity 'warning'.
	 */
	SeverityCountWarning int64

	Payload models.EventList
}

// IsSuccess returns true when this v2 list events o k response has a 2xx status code
func (o *V2ListEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 list events o k response has a 3xx status code
func (o *V2ListEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list events o k response has a 4xx status code
func (o *V2ListEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list events o k response has a 5xx status code
func (o *V2ListEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list events o k response a status code equal to that given
func (o *V2ListEventsOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2ListEventsOK) Error() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsOK  %+v", 200, o.Payload)
}

func (o *V2ListEventsOK) String() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsOK  %+v", 200, o.Payload)
}

func (o *V2ListEventsOK) GetPayload() models.EventList {
	return o.Payload
}

func (o *V2ListEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Severity-Count-Critical
	hdrSeverityCountCritical := response.GetHeader("Severity-Count-Critical")

	if hdrSeverityCountCritical != "" {
		valseverityCountCritical, err := swag.ConvertInt64(hdrSeverityCountCritical)
		if err != nil {
			return errors.InvalidType("Severity-Count-Critical", "header", "int64", hdrSeverityCountCritical)
		}
		o.SeverityCountCritical = valseverityCountCritical
	}

	// hydrates response header Severity-Count-Error
	hdrSeverityCountError := response.GetHeader("Severity-Count-Error")

	if hdrSeverityCountError != "" {
		valseverityCountError, err := swag.ConvertInt64(hdrSeverityCountError)
		if err != nil {
			return errors.InvalidType("Severity-Count-Error", "header", "int64", hdrSeverityCountError)
		}
		o.SeverityCountError = valseverityCountError
	}

	// hydrates response header Severity-Count-Info
	hdrSeverityCountInfo := response.GetHeader("Severity-Count-Info")

	if hdrSeverityCountInfo != "" {
		valseverityCountInfo, err := swag.ConvertInt64(hdrSeverityCountInfo)
		if err != nil {
			return errors.InvalidType("Severity-Count-Info", "header", "int64", hdrSeverityCountInfo)
		}
		o.SeverityCountInfo = valseverityCountInfo
	}

	// hydrates response header Severity-Count-Warning
	hdrSeverityCountWarning := response.GetHeader("Severity-Count-Warning")

	if hdrSeverityCountWarning != "" {
		valseverityCountWarning, err := swag.ConvertInt64(hdrSeverityCountWarning)
		if err != nil {
			return errors.InvalidType("Severity-Count-Warning", "header", "int64", hdrSeverityCountWarning)
		}
		o.SeverityCountWarning = valseverityCountWarning
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListEventsUnauthorized creates a V2ListEventsUnauthorized with default headers values
func NewV2ListEventsUnauthorized() *V2ListEventsUnauthorized {
	return &V2ListEventsUnauthorized{}
}

/*
V2ListEventsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2ListEventsUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 list events unauthorized response has a 2xx status code
func (o *V2ListEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list events unauthorized response has a 3xx status code
func (o *V2ListEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list events unauthorized response has a 4xx status code
func (o *V2ListEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list events unauthorized response has a 5xx status code
func (o *V2ListEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list events unauthorized response a status code equal to that given
func (o *V2ListEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2ListEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListEventsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListEventsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListEventsForbidden creates a V2ListEventsForbidden with default headers values
func NewV2ListEventsForbidden() *V2ListEventsForbidden {
	return &V2ListEventsForbidden{}
}

/*
V2ListEventsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2ListEventsForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 list events forbidden response has a 2xx status code
func (o *V2ListEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list events forbidden response has a 3xx status code
func (o *V2ListEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list events forbidden response has a 4xx status code
func (o *V2ListEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list events forbidden response has a 5xx status code
func (o *V2ListEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list events forbidden response a status code equal to that given
func (o *V2ListEventsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2ListEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListEventsForbidden) String() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListEventsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListEventsNotFound creates a V2ListEventsNotFound with default headers values
func NewV2ListEventsNotFound() *V2ListEventsNotFound {
	return &V2ListEventsNotFound{}
}

/*
V2ListEventsNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2ListEventsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list events not found response has a 2xx status code
func (o *V2ListEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list events not found response has a 3xx status code
func (o *V2ListEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list events not found response has a 4xx status code
func (o *V2ListEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list events not found response has a 5xx status code
func (o *V2ListEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list events not found response a status code equal to that given
func (o *V2ListEventsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2ListEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsNotFound  %+v", 404, o.Payload)
}

func (o *V2ListEventsNotFound) String() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsNotFound  %+v", 404, o.Payload)
}

func (o *V2ListEventsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListEventsMethodNotAllowed creates a V2ListEventsMethodNotAllowed with default headers values
func NewV2ListEventsMethodNotAllowed() *V2ListEventsMethodNotAllowed {
	return &V2ListEventsMethodNotAllowed{}
}

/*
V2ListEventsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2ListEventsMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list events method not allowed response has a 2xx status code
func (o *V2ListEventsMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list events method not allowed response has a 3xx status code
func (o *V2ListEventsMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list events method not allowed response has a 4xx status code
func (o *V2ListEventsMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list events method not allowed response has a 5xx status code
func (o *V2ListEventsMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list events method not allowed response a status code equal to that given
func (o *V2ListEventsMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2ListEventsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListEventsMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListEventsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListEventsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListEventsInternalServerError creates a V2ListEventsInternalServerError with default headers values
func NewV2ListEventsInternalServerError() *V2ListEventsInternalServerError {
	return &V2ListEventsInternalServerError{}
}

/*
V2ListEventsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ListEventsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list events internal server error response has a 2xx status code
func (o *V2ListEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list events internal server error response has a 3xx status code
func (o *V2ListEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list events internal server error response has a 4xx status code
func (o *V2ListEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list events internal server error response has a 5xx status code
func (o *V2ListEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 list events internal server error response a status code equal to that given
func (o *V2ListEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2ListEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListEventsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/events][%d] v2ListEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListEventsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
