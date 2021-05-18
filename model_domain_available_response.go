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

// DomainAvailableResponse struct for DomainAvailableResponse
type DomainAvailableResponse struct {
	// Whether or not the domain name is available
	Available bool `json:"available"`
	// Currency in which the `price` is listed. Only returned if tld is offered
	Currency *string `json:"currency,omitempty"`
	// Whether or not the `available` answer has been definitively verified with the registry
	Definitive bool `json:"definitive"`
	// Domain name
	Domain string `json:"domain"`
	// Number of years included in the price. Only returned if tld is offered
	Period *int32 `json:"period,omitempty"`
	// Price of the domain excluding taxes or fees. Only returned if tld is offered
	Price *int32 `json:"price,omitempty"`
}

// NewDomainAvailableResponse instantiates a new DomainAvailableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainAvailableResponse(available bool, definitive bool, domain string) *DomainAvailableResponse {
	this := DomainAvailableResponse{}
	this.Available = available
	var currency string = "USD"
	this.Currency = &currency
	this.Definitive = definitive
	this.Domain = domain
	return &this
}

// NewDomainAvailableResponseWithDefaults instantiates a new DomainAvailableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainAvailableResponseWithDefaults() *DomainAvailableResponse {
	this := DomainAvailableResponse{}
	var currency string = "USD"
	this.Currency = &currency
	return &this
}

// GetAvailable returns the Available field value
func (o *DomainAvailableResponse) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableResponse) GetAvailableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *DomainAvailableResponse) SetAvailable(v bool) {
	o.Available = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *DomainAvailableResponse) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAvailableResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *DomainAvailableResponse) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *DomainAvailableResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetDefinitive returns the Definitive field value
func (o *DomainAvailableResponse) GetDefinitive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Definitive
}

// GetDefinitiveOk returns a tuple with the Definitive field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableResponse) GetDefinitiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Definitive, true
}

// SetDefinitive sets field value
func (o *DomainAvailableResponse) SetDefinitive(v bool) {
	o.Definitive = v
}

// GetDomain returns the Domain field value
func (o *DomainAvailableResponse) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainAvailableResponse) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainAvailableResponse) SetDomain(v string) {
	o.Domain = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DomainAvailableResponse) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAvailableResponse) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DomainAvailableResponse) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *DomainAvailableResponse) SetPeriod(v int32) {
	o.Period = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *DomainAvailableResponse) GetPrice() int32 {
	if o == nil || o.Price == nil {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAvailableResponse) GetPriceOk() (*int32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *DomainAvailableResponse) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *DomainAvailableResponse) SetPrice(v int32) {
	o.Price = &v
}

func (o DomainAvailableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["available"] = o.Available
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["definitive"] = o.Definitive
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableDomainAvailableResponse struct {
	value *DomainAvailableResponse
	isSet bool
}

func (v NullableDomainAvailableResponse) Get() *DomainAvailableResponse {
	return v.value
}

func (v *NullableDomainAvailableResponse) Set(val *DomainAvailableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainAvailableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainAvailableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainAvailableResponse(val *DomainAvailableResponse) *NullableDomainAvailableResponse {
	return &NullableDomainAvailableResponse{value: val, isSet: true}
}

func (v NullableDomainAvailableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainAvailableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

