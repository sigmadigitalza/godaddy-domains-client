/*
 * GoDaddy API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DomainAvailableError struct for DomainAvailableError
type DomainAvailableError struct {
	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code"`
	// Domain name
	Domain string `json:"domain"`
	// Human-readable, English description of the error
	Message *string `json:"message,omitempty"`
	// <ul> <li style='margin-left: 12px;'>JSONPath referring to a field containing an error</li> <strong style='margin-left: 12px;'>OR</strong> <li style='margin-left: 12px;'>JSONPath referring to a field that refers to an object containing an error, with more detail in `pathRelated`</li> </ul>
	Path string `json:"path"`
	// HTTP status code that would return for a single check
	Status int32 `json:"status"`
}

// NewDomainAvailableError instantiates a new DomainAvailableError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainAvailableError(code string, domain string, path string, status int32) *DomainAvailableError {
	this := DomainAvailableError{}
	this.Code = code
	this.Domain = domain
	this.Path = path
	this.Status = status
	return &this
}

// NewDomainAvailableErrorWithDefaults instantiates a new DomainAvailableError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainAvailableErrorWithDefaults() *DomainAvailableError {
	this := DomainAvailableError{}
	return &this
}

// GetCode returns the Code field value
func (o *DomainAvailableError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableError) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DomainAvailableError) SetCode(v string) {
	o.Code = v
}

// GetDomain returns the Domain field value
func (o *DomainAvailableError) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableError) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainAvailableError) SetDomain(v string) {
	o.Domain = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DomainAvailableError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAvailableError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DomainAvailableError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DomainAvailableError) SetMessage(v string) {
	o.Message = &v
}

// GetPath returns the Path field value
func (o *DomainAvailableError) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableError) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DomainAvailableError) SetPath(v string) {
	o.Path = v
}

// GetStatus returns the Status field value
func (o *DomainAvailableError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableError) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DomainAvailableError) SetStatus(v int32) {
	o.Status = v
}

func (o DomainAvailableError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDomainAvailableError struct {
	value *DomainAvailableError
	isSet bool
}

func (v NullableDomainAvailableError) Get() *DomainAvailableError {
	return v.value
}

func (v *NullableDomainAvailableError) Set(val *DomainAvailableError) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainAvailableError) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainAvailableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainAvailableError(val *DomainAvailableError) *NullableDomainAvailableError {
	return &NullableDomainAvailableError{value: val, isSet: true}
}

func (v NullableDomainAvailableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainAvailableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

