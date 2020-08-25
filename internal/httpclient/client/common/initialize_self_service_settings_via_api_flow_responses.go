// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// InitializeSelfServiceSettingsViaAPIFlowReader is a Reader for the InitializeSelfServiceSettingsViaAPIFlow structure.
type InitializeSelfServiceSettingsViaAPIFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeSelfServiceSettingsViaAPIFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInitializeSelfServiceSettingsViaAPIFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInitializeSelfServiceSettingsViaAPIFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeSelfServiceSettingsViaAPIFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitializeSelfServiceSettingsViaAPIFlowOK creates a InitializeSelfServiceSettingsViaAPIFlowOK with default headers values
func NewInitializeSelfServiceSettingsViaAPIFlowOK() *InitializeSelfServiceSettingsViaAPIFlowOK {
	return &InitializeSelfServiceSettingsViaAPIFlowOK{}
}

/*InitializeSelfServiceSettingsViaAPIFlowOK handles this case with default header values.

settingsFlow
*/
type InitializeSelfServiceSettingsViaAPIFlowOK struct {
	Payload *models.SettingsFlow
}

func (o *InitializeSelfServiceSettingsViaAPIFlowOK) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/api][%d] initializeSelfServiceSettingsViaApiFlowOK  %+v", 200, o.Payload)
}

func (o *InitializeSelfServiceSettingsViaAPIFlowOK) GetPayload() *models.SettingsFlow {
	return o.Payload
}

func (o *InitializeSelfServiceSettingsViaAPIFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeSelfServiceSettingsViaAPIFlowBadRequest creates a InitializeSelfServiceSettingsViaAPIFlowBadRequest with default headers values
func NewInitializeSelfServiceSettingsViaAPIFlowBadRequest() *InitializeSelfServiceSettingsViaAPIFlowBadRequest {
	return &InitializeSelfServiceSettingsViaAPIFlowBadRequest{}
}

/*InitializeSelfServiceSettingsViaAPIFlowBadRequest handles this case with default header values.

genericError
*/
type InitializeSelfServiceSettingsViaAPIFlowBadRequest struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceSettingsViaAPIFlowBadRequest) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/api][%d] initializeSelfServiceSettingsViaApiFlowBadRequest  %+v", 400, o.Payload)
}

func (o *InitializeSelfServiceSettingsViaAPIFlowBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceSettingsViaAPIFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeSelfServiceSettingsViaAPIFlowInternalServerError creates a InitializeSelfServiceSettingsViaAPIFlowInternalServerError with default headers values
func NewInitializeSelfServiceSettingsViaAPIFlowInternalServerError() *InitializeSelfServiceSettingsViaAPIFlowInternalServerError {
	return &InitializeSelfServiceSettingsViaAPIFlowInternalServerError{}
}

/*InitializeSelfServiceSettingsViaAPIFlowInternalServerError handles this case with default header values.

genericError
*/
type InitializeSelfServiceSettingsViaAPIFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceSettingsViaAPIFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/api][%d] initializeSelfServiceSettingsViaApiFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *InitializeSelfServiceSettingsViaAPIFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceSettingsViaAPIFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
