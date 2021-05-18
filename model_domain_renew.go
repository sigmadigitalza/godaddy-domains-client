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

// DomainRenew struct for DomainRenew
type DomainRenew struct {
	// Number of years to extend the Domain. Must not exceed maximum for TLD. When omitted, defaults to `period` specified during original purchase
	Period *int32 `json:"period,omitempty"`
}

// NewDomainRenew instantiates a new DomainRenew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRenew() *DomainRenew {
	this := DomainRenew{}
	return &this
}

// NewDomainRenewWithDefaults instantiates a new DomainRenew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRenewWithDefaults() *DomainRenew {
	this := DomainRenew{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DomainRenew) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRenew) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DomainRenew) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *DomainRenew) SetPeriod(v int32) {
	o.Period = &v
}

func (o DomainRenew) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableDomainRenew struct {
	value *DomainRenew
	isSet bool
}

func (v NullableDomainRenew) Get() *DomainRenew {
	return v.value
}

func (v *NullableDomainRenew) Set(val *DomainRenew) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRenew) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRenew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRenew(val *DomainRenew) *NullableDomainRenew {
	return &NullableDomainRenew{value: val, isSet: true}
}

func (v NullableDomainRenew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRenew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

