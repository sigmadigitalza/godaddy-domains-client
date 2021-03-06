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

// DNSRecordCreateType struct for DNSRecordCreateType
type DNSRecordCreateType struct {
	Data string `json:"data"`
	Name string `json:"name"`
	// Service port (SRV only)
	Port *int32 `json:"port,omitempty"`
	// Record priority (MX and SRV only)
	Priority *int32 `json:"priority,omitempty"`
	// Service protocol (SRV only)
	Protocol *string `json:"protocol,omitempty"`
	// Service type (SRV only)
	Service *string `json:"service,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	// Record weight (SRV only)
	Weight *int32 `json:"weight,omitempty"`
}

// NewDNSRecordCreateType instantiates a new DNSRecordCreateType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSRecordCreateType(data string, name string) *DNSRecordCreateType {
	this := DNSRecordCreateType{}
	this.Data = data
	this.Name = name
	return &this
}

// NewDNSRecordCreateTypeWithDefaults instantiates a new DNSRecordCreateType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSRecordCreateTypeWithDefaults() *DNSRecordCreateType {
	this := DNSRecordCreateType{}
	return &this
}

// GetData returns the Data field value
func (o *DNSRecordCreateType) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DNSRecordCreateType) SetData(v string) {
	o.Data = v
}

// GetName returns the Name field value
func (o *DNSRecordCreateType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DNSRecordCreateType) SetName(v string) {
	o.Name = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DNSRecordCreateType) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DNSRecordCreateType) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DNSRecordCreateType) SetPort(v int32) {
	o.Port = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DNSRecordCreateType) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DNSRecordCreateType) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *DNSRecordCreateType) SetPriority(v int32) {
	o.Priority = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *DNSRecordCreateType) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *DNSRecordCreateType) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *DNSRecordCreateType) SetProtocol(v string) {
	o.Protocol = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *DNSRecordCreateType) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *DNSRecordCreateType) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *DNSRecordCreateType) SetService(v string) {
	o.Service = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DNSRecordCreateType) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DNSRecordCreateType) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *DNSRecordCreateType) SetTtl(v int32) {
	o.Ttl = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *DNSRecordCreateType) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordCreateType) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *DNSRecordCreateType) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *DNSRecordCreateType) SetWeight(v int32) {
	o.Weight = &v
}

func (o DNSRecordCreateType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableDNSRecordCreateType struct {
	value *DNSRecordCreateType
	isSet bool
}

func (v NullableDNSRecordCreateType) Get() *DNSRecordCreateType {
	return v.value
}

func (v *NullableDNSRecordCreateType) Set(val *DNSRecordCreateType) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSRecordCreateType) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSRecordCreateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSRecordCreateType(val *DNSRecordCreateType) *NullableDNSRecordCreateType {
	return &NullableDNSRecordCreateType{value: val, isSet: true}
}

func (v NullableDNSRecordCreateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSRecordCreateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


