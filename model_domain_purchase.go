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

// DomainPurchase struct for DomainPurchase
type DomainPurchase struct {
	Consent Consent `json:"consent"`
	ContactAdmin *Contact `json:"contactAdmin,omitempty"`
	ContactBilling *Contact `json:"contactBilling,omitempty"`
	ContactRegistrant *Contact `json:"contactRegistrant,omitempty"`
	ContactTech *Contact `json:"contactTech,omitempty"`
	// For internationalized domain names with non-ascii characters, the domain name is converted to punycode before format and pattern validation rules are checked
	Domain string `json:"domain"`
	NameServers *[]string `json:"nameServers,omitempty"`
	Period *int32 `json:"period,omitempty"`
	Privacy *bool `json:"privacy,omitempty"`
	RenewAuto *bool `json:"renewAuto,omitempty"`
}

// NewDomainPurchase instantiates a new DomainPurchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainPurchase(consent Consent, domain string) *DomainPurchase {
	this := DomainPurchase{}
	this.Consent = consent
	this.Domain = domain
	var privacy bool = false
	this.Privacy = &privacy
	var renewAuto bool = true
	this.RenewAuto = &renewAuto
	return &this
}

// NewDomainPurchaseWithDefaults instantiates a new DomainPurchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainPurchaseWithDefaults() *DomainPurchase {
	this := DomainPurchase{}
	var privacy bool = false
	this.Privacy = &privacy
	var renewAuto bool = true
	this.RenewAuto = &renewAuto
	return &this
}

// GetConsent returns the Consent field value
func (o *DomainPurchase) GetConsent() Consent {
	if o == nil {
		var ret Consent
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetConsentOk() (*Consent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *DomainPurchase) SetConsent(v Consent) {
	o.Consent = v
}

// GetContactAdmin returns the ContactAdmin field value if set, zero value otherwise.
func (o *DomainPurchase) GetContactAdmin() Contact {
	if o == nil || o.ContactAdmin == nil {
		var ret Contact
		return ret
	}
	return *o.ContactAdmin
}

// GetContactAdminOk returns a tuple with the ContactAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetContactAdminOk() (*Contact, bool) {
	if o == nil || o.ContactAdmin == nil {
		return nil, false
	}
	return o.ContactAdmin, true
}

// HasContactAdmin returns a boolean if a field has been set.
func (o *DomainPurchase) HasContactAdmin() bool {
	if o != nil && o.ContactAdmin != nil {
		return true
	}

	return false
}

// SetContactAdmin gets a reference to the given Contact and assigns it to the ContactAdmin field.
func (o *DomainPurchase) SetContactAdmin(v Contact) {
	o.ContactAdmin = &v
}

// GetContactBilling returns the ContactBilling field value if set, zero value otherwise.
func (o *DomainPurchase) GetContactBilling() Contact {
	if o == nil || o.ContactBilling == nil {
		var ret Contact
		return ret
	}
	return *o.ContactBilling
}

// GetContactBillingOk returns a tuple with the ContactBilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetContactBillingOk() (*Contact, bool) {
	if o == nil || o.ContactBilling == nil {
		return nil, false
	}
	return o.ContactBilling, true
}

// HasContactBilling returns a boolean if a field has been set.
func (o *DomainPurchase) HasContactBilling() bool {
	if o != nil && o.ContactBilling != nil {
		return true
	}

	return false
}

// SetContactBilling gets a reference to the given Contact and assigns it to the ContactBilling field.
func (o *DomainPurchase) SetContactBilling(v Contact) {
	o.ContactBilling = &v
}

// GetContactRegistrant returns the ContactRegistrant field value if set, zero value otherwise.
func (o *DomainPurchase) GetContactRegistrant() Contact {
	if o == nil || o.ContactRegistrant == nil {
		var ret Contact
		return ret
	}
	return *o.ContactRegistrant
}

// GetContactRegistrantOk returns a tuple with the ContactRegistrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetContactRegistrantOk() (*Contact, bool) {
	if o == nil || o.ContactRegistrant == nil {
		return nil, false
	}
	return o.ContactRegistrant, true
}

// HasContactRegistrant returns a boolean if a field has been set.
func (o *DomainPurchase) HasContactRegistrant() bool {
	if o != nil && o.ContactRegistrant != nil {
		return true
	}

	return false
}

// SetContactRegistrant gets a reference to the given Contact and assigns it to the ContactRegistrant field.
func (o *DomainPurchase) SetContactRegistrant(v Contact) {
	o.ContactRegistrant = &v
}

// GetContactTech returns the ContactTech field value if set, zero value otherwise.
func (o *DomainPurchase) GetContactTech() Contact {
	if o == nil || o.ContactTech == nil {
		var ret Contact
		return ret
	}
	return *o.ContactTech
}

// GetContactTechOk returns a tuple with the ContactTech field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetContactTechOk() (*Contact, bool) {
	if o == nil || o.ContactTech == nil {
		return nil, false
	}
	return o.ContactTech, true
}

// HasContactTech returns a boolean if a field has been set.
func (o *DomainPurchase) HasContactTech() bool {
	if o != nil && o.ContactTech != nil {
		return true
	}

	return false
}

// SetContactTech gets a reference to the given Contact and assigns it to the ContactTech field.
func (o *DomainPurchase) SetContactTech(v Contact) {
	o.ContactTech = &v
}

// GetDomain returns the Domain field value
func (o *DomainPurchase) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainPurchase) SetDomain(v string) {
	o.Domain = v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise.
func (o *DomainPurchase) GetNameServers() []string {
	if o == nil || o.NameServers == nil {
		var ret []string
		return ret
	}
	return *o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetNameServersOk() (*[]string, bool) {
	if o == nil || o.NameServers == nil {
		return nil, false
	}
	return o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *DomainPurchase) HasNameServers() bool {
	if o != nil && o.NameServers != nil {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *DomainPurchase) SetNameServers(v []string) {
	o.NameServers = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DomainPurchase) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DomainPurchase) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *DomainPurchase) SetPeriod(v int32) {
	o.Period = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *DomainPurchase) GetPrivacy() bool {
	if o == nil || o.Privacy == nil {
		var ret bool
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetPrivacyOk() (*bool, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *DomainPurchase) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given bool and assigns it to the Privacy field.
func (o *DomainPurchase) SetPrivacy(v bool) {
	o.Privacy = &v
}

// GetRenewAuto returns the RenewAuto field value if set, zero value otherwise.
func (o *DomainPurchase) GetRenewAuto() bool {
	if o == nil || o.RenewAuto == nil {
		var ret bool
		return ret
	}
	return *o.RenewAuto
}

// GetRenewAutoOk returns a tuple with the RenewAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurchase) GetRenewAutoOk() (*bool, bool) {
	if o == nil || o.RenewAuto == nil {
		return nil, false
	}
	return o.RenewAuto, true
}

// HasRenewAuto returns a boolean if a field has been set.
func (o *DomainPurchase) HasRenewAuto() bool {
	if o != nil && o.RenewAuto != nil {
		return true
	}

	return false
}

// SetRenewAuto gets a reference to the given bool and assigns it to the RenewAuto field.
func (o *DomainPurchase) SetRenewAuto(v bool) {
	o.RenewAuto = &v
}

func (o DomainPurchase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consent"] = o.Consent
	}
	if o.ContactAdmin != nil {
		toSerialize["contactAdmin"] = o.ContactAdmin
	}
	if o.ContactBilling != nil {
		toSerialize["contactBilling"] = o.ContactBilling
	}
	if o.ContactRegistrant != nil {
		toSerialize["contactRegistrant"] = o.ContactRegistrant
	}
	if o.ContactTech != nil {
		toSerialize["contactTech"] = o.ContactTech
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.NameServers != nil {
		toSerialize["nameServers"] = o.NameServers
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if o.RenewAuto != nil {
		toSerialize["renewAuto"] = o.RenewAuto
	}
	return json.Marshal(toSerialize)
}

type NullableDomainPurchase struct {
	value *DomainPurchase
	isSet bool
}

func (v NullableDomainPurchase) Get() *DomainPurchase {
	return v.value
}

func (v *NullableDomainPurchase) Set(val *DomainPurchase) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainPurchase) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainPurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainPurchase(val *DomainPurchase) *NullableDomainPurchase {
	return &NullableDomainPurchase{value: val, isSet: true}
}

func (v NullableDomainPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainPurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

