// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInterfacesBulkDeleteReader is a Reader for the DcimInterfacesBulkDelete structure.
type DcimInterfacesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInterfacesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesBulkDeleteNoContent creates a DcimInterfacesBulkDeleteNoContent with default headers values
func NewDcimInterfacesBulkDeleteNoContent() *DcimInterfacesBulkDeleteNoContent {
	return &DcimInterfacesBulkDeleteNoContent{}
}

/*
DcimInterfacesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInterfacesBulkDeleteNoContent dcim interfaces bulk delete no content
*/
type DcimInterfacesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim interfaces bulk delete no content response has a 2xx status code
func (o *DcimInterfacesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces bulk delete no content response has a 3xx status code
func (o *DcimInterfacesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces bulk delete no content response has a 4xx status code
func (o *DcimInterfacesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces bulk delete no content response has a 5xx status code
func (o *DcimInterfacesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces bulk delete no content response a status code equal to that given
func (o *DcimInterfacesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimInterfacesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interfaces/][%d] dcimInterfacesBulkDeleteNoContent ", 204)
}

func (o *DcimInterfacesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/interfaces/][%d] dcimInterfacesBulkDeleteNoContent ", 204)
}

func (o *DcimInterfacesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimInterfacesBulkDeleteDefault creates a DcimInterfacesBulkDeleteDefault with default headers values
func NewDcimInterfacesBulkDeleteDefault(code int) *DcimInterfacesBulkDeleteDefault {
	return &DcimInterfacesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimInterfacesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimInterfacesBulkDeleteDefault dcim interfaces bulk delete default
*/
type DcimInterfacesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces bulk delete default response
func (o *DcimInterfacesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interfaces bulk delete default response has a 2xx status code
func (o *DcimInterfacesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interfaces bulk delete default response has a 3xx status code
func (o *DcimInterfacesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interfaces bulk delete default response has a 4xx status code
func (o *DcimInterfacesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interfaces bulk delete default response has a 5xx status code
func (o *DcimInterfacesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interfaces bulk delete default response a status code equal to that given
func (o *DcimInterfacesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfacesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interfaces/][%d] dcim_interfaces_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/interfaces/][%d] dcim_interfaces_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
