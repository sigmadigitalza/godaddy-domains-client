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

// DomainAvailableBulkMixed struct for DomainAvailableBulkMixed
type DomainAvailableBulkMixed struct {
	// Domain available response array
	Domains []DomainAvailableResponse `json:"domains"`
	// Errors encountered while performing a domain available check
	Errors *[]DomainAvailableError `json:"errors,omitempty"`
}

// NewDomainAvailableBulkMixed instantiates a new DomainAvailableBulkMixed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainAvailableBulkMixed(domains []DomainAvailableResponse) *DomainAvailableBulkMixed {
	this := DomainAvailableBulkMixed{}
	this.Domains = domains
	return &this
}

// NewDomainAvailableBulkMixedWithDefaults instantiates a new DomainAvailableBulkMixed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainAvailableBulkMixedWithDefaults() *DomainAvailableBulkMixed {
	this := DomainAvailableBulkMixed{}
	return &this
}

// GetDomains returns the Domains field value
func (o *DomainAvailableBulkMixed) GetDomains() []DomainAvailableResponse {
	if o == nil {
		var ret []DomainAvailableResponse
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableBulkMixed) GetDomainsOk() (*[]DomainAvailableResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domains, true
}

// SetDomains sets field value
func (o *DomainAvailableBulkMixed) SetDomains(v []DomainAvailableResponse) {
	o.Domains = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DomainAvailableBulkMixed) GetErrors() []DomainAvailableError {
	if o == nil || o.Errors == nil {
		var ret []DomainAvailableError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAvailableBulkMixed) GetErrorsOk() (*[]DomainAvailableError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DomainAvailableBulkMixed) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []DomainAvailableError and assigns it to the Errors field.
func (o *DomainAvailableBulkMixed) SetErrors(v []DomainAvailableError) {
	o.Errors = &v
}

func (o DomainAvailableBulkMixed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domains"] = o.Domains
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableDomainAvailableBulkMixed struct {
	value *DomainAvailableBulkMixed
	isSet bool
}

func (v NullableDomainAvailableBulkMixed) Get() *DomainAvailableBulkMixed {
	return v.value
}

func (v *NullableDomainAvailableBulkMixed) Set(val *DomainAvailableBulkMixed) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainAvailableBulkMixed) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainAvailableBulkMixed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainAvailableBulkMixed(val *DomainAvailableBulkMixed) *NullableDomainAvailableBulkMixed {
	return &NullableDomainAvailableBulkMixed{value: val, isSet: true}
}

func (v NullableDomainAvailableBulkMixed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainAvailableBulkMixed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


