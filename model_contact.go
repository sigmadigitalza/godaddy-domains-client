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

// Contact struct for Contact
type Contact struct {
	AddressMailing Address `json:"addressMailing"`
	Email string `json:"email"`
	Fax *string `json:"fax,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
	NameFirst string `json:"nameFirst"`
	NameLast string `json:"nameLast"`
	NameMiddle *string `json:"nameMiddle,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Phone string `json:"phone"`
}

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact(addressMailing Address, email string, nameFirst string, nameLast string, phone string) *Contact {
	this := Contact{}
	this.AddressMailing = addressMailing
	this.Email = email
	this.NameFirst = nameFirst
	this.NameLast = nameLast
	this.Phone = phone
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetAddressMailing returns the AddressMailing field value
func (o *Contact) GetAddressMailing() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.AddressMailing
}

// GetAddressMailingOk returns a tuple with the AddressMailing field value
// and a boolean to check if the value has been set.
func (o *Contact) GetAddressMailingOk() (*Address, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddressMailing, true
}

// SetAddressMailing sets field value
func (o *Contact) SetAddressMailing(v Address) {
	o.AddressMailing = v
}

// GetEmail returns the Email field value
func (o *Contact) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Contact) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Contact) SetEmail(v string) {
	o.Email = v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *Contact) GetFax() string {
	if o == nil || o.Fax == nil {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetFaxOk() (*string, bool) {
	if o == nil || o.Fax == nil {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *Contact) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *Contact) SetFax(v string) {
	o.Fax = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *Contact) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Contact) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *Contact) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetNameFirst returns the NameFirst field value
func (o *Contact) GetNameFirst() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameFirst
}

// GetNameFirstOk returns a tuple with the NameFirst field value
// and a boolean to check if the value has been set.
func (o *Contact) GetNameFirstOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameFirst, true
}

// SetNameFirst sets field value
func (o *Contact) SetNameFirst(v string) {
	o.NameFirst = v
}

// GetNameLast returns the NameLast field value
func (o *Contact) GetNameLast() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameLast
}

// GetNameLastOk returns a tuple with the NameLast field value
// and a boolean to check if the value has been set.
func (o *Contact) GetNameLastOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameLast, true
}

// SetNameLast sets field value
func (o *Contact) SetNameLast(v string) {
	o.NameLast = v
}

// GetNameMiddle returns the NameMiddle field value if set, zero value otherwise.
func (o *Contact) GetNameMiddle() string {
	if o == nil || o.NameMiddle == nil {
		var ret string
		return ret
	}
	return *o.NameMiddle
}

// GetNameMiddleOk returns a tuple with the NameMiddle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetNameMiddleOk() (*string, bool) {
	if o == nil || o.NameMiddle == nil {
		return nil, false
	}
	return o.NameMiddle, true
}

// HasNameMiddle returns a boolean if a field has been set.
func (o *Contact) HasNameMiddle() bool {
	if o != nil && o.NameMiddle != nil {
		return true
	}

	return false
}

// SetNameMiddle gets a reference to the given string and assigns it to the NameMiddle field.
func (o *Contact) SetNameMiddle(v string) {
	o.NameMiddle = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Contact) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Contact) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *Contact) SetOrganization(v string) {
	o.Organization = &v
}

// GetPhone returns the Phone field value
func (o *Contact) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *Contact) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *Contact) SetPhone(v string) {
	o.Phone = v
}

func (o Contact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addressMailing"] = o.AddressMailing
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Fax != nil {
		toSerialize["fax"] = o.Fax
	}
	if o.JobTitle != nil {
		toSerialize["jobTitle"] = o.JobTitle
	}
	if true {
		toSerialize["nameFirst"] = o.NameFirst
	}
	if true {
		toSerialize["nameLast"] = o.NameLast
	}
	if o.NameMiddle != nil {
		toSerialize["nameMiddle"] = o.NameMiddle
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableContact struct {
	value *Contact
	isSet bool
}

func (v NullableContact) Get() *Contact {
	return v.value
}

func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

func (v NullableContact) IsSet() bool {
	return v.isSet
}

func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


