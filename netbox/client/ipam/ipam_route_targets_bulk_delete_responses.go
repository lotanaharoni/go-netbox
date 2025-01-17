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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRouteTargetsBulkDeleteReader is a Reader for the IpamRouteTargetsBulkDelete structure.
type IpamRouteTargetsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRouteTargetsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsBulkDeleteNoContent creates a IpamRouteTargetsBulkDeleteNoContent with default headers values
func NewIpamRouteTargetsBulkDeleteNoContent() *IpamRouteTargetsBulkDeleteNoContent {
	return &IpamRouteTargetsBulkDeleteNoContent{}
}

/*
IpamRouteTargetsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamRouteTargetsBulkDeleteNoContent ipam route targets bulk delete no content
*/
type IpamRouteTargetsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam route targets bulk delete no content response has a 2xx status code
func (o *IpamRouteTargetsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets bulk delete no content response has a 3xx status code
func (o *IpamRouteTargetsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets bulk delete no content response has a 4xx status code
func (o *IpamRouteTargetsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets bulk delete no content response has a 5xx status code
func (o *IpamRouteTargetsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets bulk delete no content response a status code equal to that given
func (o *IpamRouteTargetsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamRouteTargetsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipamRouteTargetsBulkDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipamRouteTargetsBulkDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRouteTargetsBulkDeleteDefault creates a IpamRouteTargetsBulkDeleteDefault with default headers values
func NewIpamRouteTargetsBulkDeleteDefault(code int) *IpamRouteTargetsBulkDeleteDefault {
	return &IpamRouteTargetsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamRouteTargetsBulkDeleteDefault describes a response with status code -1, with default header values.

IpamRouteTargetsBulkDeleteDefault ipam route targets bulk delete default
*/
type IpamRouteTargetsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets bulk delete default response
func (o *IpamRouteTargetsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam route targets bulk delete default response has a 2xx status code
func (o *IpamRouteTargetsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam route targets bulk delete default response has a 3xx status code
func (o *IpamRouteTargetsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam route targets bulk delete default response has a 4xx status code
func (o *IpamRouteTargetsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam route targets bulk delete default response has a 5xx status code
func (o *IpamRouteTargetsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam route targets bulk delete default response a status code equal to that given
func (o *IpamRouteTargetsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRouteTargetsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipam_route-targets_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipam_route-targets_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
