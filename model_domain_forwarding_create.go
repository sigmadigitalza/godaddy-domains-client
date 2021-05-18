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

// DomainForwardingCreate struct for DomainForwardingCreate
type DomainForwardingCreate struct {
	// The type of fowarding to implement<br/><ul><li><strong style='margin-left: 12px;'>MASKED</strong> - Prevents the forwarded domain or subdomain URL from displaying in the browser's address bar.</li><li><strong style='margin-left: 12px;'>REDIRECT_PERMANENT*</strong> - Redirects to the url you specified in the forwardTo field using a `301 Moved Permanently` HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.</li><li><strong style='margin-left: 12px;'>REDIRECT_TEMPORARY</strong> - Redirects to the url you specified in the forwardTo field using a `302 Found` HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.</li></ul>
	Type string `json:"type"`
	// Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/)
	Url string `json:"url"`
	Mask *DomainForwardingMask `json:"mask,omitempty"`
}

// NewDomainForwardingCreate instantiates a new DomainForwardingCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainForwardingCreate(type_ string, url string) *DomainForwardingCreate {
	this := DomainForwardingCreate{}
	this.Type = type_
	this.Url = url
	return &this
}

// NewDomainForwardingCreateWithDefaults instantiates a new DomainForwardingCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainForwardingCreateWithDefaults() *DomainForwardingCreate {
	this := DomainForwardingCreate{}
	var type_ string = "REDIRECT_PERMANENT"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *DomainForwardingCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DomainForwardingCreate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DomainForwardingCreate) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *DomainForwardingCreate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DomainForwardingCreate) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DomainForwardingCreate) SetUrl(v string) {
	o.Url = v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *DomainForwardingCreate) GetMask() DomainForwardingMask {
	if o == nil || o.Mask == nil {
		var ret DomainForwardingMask
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainForwardingCreate) GetMaskOk() (*DomainForwardingMask, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *DomainForwardingCreate) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given DomainForwardingMask and assigns it to the Mask field.
func (o *DomainForwardingCreate) SetMask(v DomainForwardingMask) {
	o.Mask = &v
}

func (o DomainForwardingCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	return json.Marshal(toSerialize)
}

type NullableDomainForwardingCreate struct {
	value *DomainForwardingCreate
	isSet bool
}

func (v NullableDomainForwardingCreate) Get() *DomainForwardingCreate {
	return v.value
}

func (v *NullableDomainForwardingCreate) Set(val *DomainForwardingCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainForwardingCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainForwardingCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainForwardingCreate(val *DomainForwardingCreate) *NullableDomainForwardingCreate {
	return &NullableDomainForwardingCreate{value: val, isSet: true}
}

func (v NullableDomainForwardingCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainForwardingCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


