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

// DomainAvailableBulk struct for DomainAvailableBulk
type DomainAvailableBulk struct {
	// Domain available response array
	Domains []DomainAvailableResponse `json:"domains"`
}

// NewDomainAvailableBulk instantiates a new DomainAvailableBulk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainAvailableBulk(domains []DomainAvailableResponse) *DomainAvailableBulk {
	this := DomainAvailableBulk{}
	this.Domains = domains
	return &this
}

// NewDomainAvailableBulkWithDefaults instantiates a new DomainAvailableBulk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainAvailableBulkWithDefaults() *DomainAvailableBulk {
	this := DomainAvailableBulk{}
	return &this
}

// GetDomains returns the Domains field value
func (o *DomainAvailableBulk) GetDomains() []DomainAvailableResponse {
	if o == nil {
		var ret []DomainAvailableResponse
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableBulk) GetDomainsOk() (*[]DomainAvailableResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domains, true
}

// SetDomains sets field value
func (o *DomainAvailableBulk) SetDomains(v []DomainAvailableResponse) {
	o.Domains = v
}

func (o DomainAvailableBulk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domains"] = o.Domains
	}
	return json.Marshal(toSerialize)
}

type NullableDomainAvailableBulk struct {
	value *DomainAvailableBulk
	isSet bool
}

func (v NullableDomainAvailableBulk) Get() *DomainAvailableBulk {
	return v.value
}

func (v *NullableDomainAvailableBulk) Set(val *DomainAvailableBulk) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainAvailableBulk) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainAvailableBulk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainAvailableBulk(val *DomainAvailableBulk) *NullableDomainAvailableBulk {
	return &NullableDomainAvailableBulk{value: val, isSet: true}
}

func (v NullableDomainAvailableBulk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainAvailableBulk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


