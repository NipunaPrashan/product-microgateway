// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wso2/micro-gw/internal/api/models"
)

// PostOauth2TokenOKCode is the HTTP code returned for type PostOauth2TokenOK
const PostOauth2TokenOKCode int = 200

/*PostOauth2TokenOK Authentication successful.
Returns an access token.


swagger:response postOauth2TokenOK
*/
type PostOauth2TokenOK struct {

	/*
	  In: Body
	*/
	Payload *PostOauth2TokenOKBody `json:"body,omitempty"`
}

// NewPostOauth2TokenOK creates PostOauth2TokenOK with default headers values
func NewPostOauth2TokenOK() *PostOauth2TokenOK {

	return &PostOauth2TokenOK{}
}

// WithPayload adds the payload to the post oauth2 token o k response
func (o *PostOauth2TokenOK) WithPayload(payload *PostOauth2TokenOKBody) *PostOauth2TokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post oauth2 token o k response
func (o *PostOauth2TokenOK) SetPayload(payload *PostOauth2TokenOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOauth2TokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOauth2TokenUnauthorizedCode is the HTTP code returned for type PostOauth2TokenUnauthorized
const PostOauth2TokenUnauthorizedCode int = 401

/*PostOauth2TokenUnauthorized Unauthorized. Invalid authentication credentials.

swagger:response postOauth2TokenUnauthorized
*/
type PostOauth2TokenUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostOauth2TokenUnauthorized creates PostOauth2TokenUnauthorized with default headers values
func NewPostOauth2TokenUnauthorized() *PostOauth2TokenUnauthorized {

	return &PostOauth2TokenUnauthorized{}
}

// WithPayload adds the payload to the post oauth2 token unauthorized response
func (o *PostOauth2TokenUnauthorized) WithPayload(payload *models.Error) *PostOauth2TokenUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post oauth2 token unauthorized response
func (o *PostOauth2TokenUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOauth2TokenUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOauth2TokenInternalServerErrorCode is the HTTP code returned for type PostOauth2TokenInternalServerError
const PostOauth2TokenInternalServerErrorCode int = 500

/*PostOauth2TokenInternalServerError Internal Server Error.

swagger:response postOauth2TokenInternalServerError
*/
type PostOauth2TokenInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostOauth2TokenInternalServerError creates PostOauth2TokenInternalServerError with default headers values
func NewPostOauth2TokenInternalServerError() *PostOauth2TokenInternalServerError {

	return &PostOauth2TokenInternalServerError{}
}

// WithPayload adds the payload to the post oauth2 token internal server error response
func (o *PostOauth2TokenInternalServerError) WithPayload(payload *models.Error) *PostOauth2TokenInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post oauth2 token internal server error response
func (o *PostOauth2TokenInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOauth2TokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}