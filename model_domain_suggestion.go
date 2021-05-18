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

// DomainSuggestion struct for DomainSuggestion
type DomainSuggestion struct {
	// Suggested domain name
	Domain string `json:"domain"`
}

// NewDomainSuggestion instantiates a new DomainSuggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainSuggestion(domain string) *DomainSuggestion {
	this := DomainSuggestion{}
	this.Domain = domain
	return &this
}

// NewDomainSuggestionWithDefaults instantiates a new DomainSuggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainSuggestionWithDefaults() *DomainSuggestion {
	this := DomainSuggestion{}
	return &this
}

// GetDomain returns the Domain field value
func (o *DomainSuggestion) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainSuggestion) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainSuggestion) SetDomain(v string) {
	o.Domain = v
}

func (o DomainSuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableDomainSuggestion struct {
	value *DomainSuggestion
	isSet bool
}

func (v NullableDomainSuggestion) Get() *DomainSuggestion {
	return v.value
}

func (v *NullableDomainSuggestion) Set(val *DomainSuggestion) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainSuggestion) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainSuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainSuggestion(val *DomainSuggestion) *NullableDomainSuggestion {
	return &NullableDomainSuggestion{value: val, isSet: true}
}

func (v NullableDomainSuggestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainSuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

