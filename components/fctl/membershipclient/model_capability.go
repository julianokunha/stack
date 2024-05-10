/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
	"fmt"
)

// Capability the model 'Capability'
type Capability string

// List of Capability
const (
	MODULE_SELECTION Capability = "MODULE_SELECTION"
)

// All allowed values of Capability enum
var AllowedCapabilityEnumValues = []Capability{
	"MODULE_SELECTION",
}

func (v *Capability) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Capability(value)
	for _, existing := range AllowedCapabilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Capability", value)
}

// NewCapabilityFromValue returns a pointer to a valid Capability
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCapabilityFromValue(v string) (*Capability, error) {
	ev := Capability(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Capability: valid values are %v", v, AllowedCapabilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Capability) IsValid() bool {
	for _, existing := range AllowedCapabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Capability value
func (v Capability) Ptr() *Capability {
	return &v
}

type NullableCapability struct {
	value *Capability
	isSet bool
}

func (v NullableCapability) Get() *Capability {
	return v.value
}

func (v *NullableCapability) Set(val *Capability) {
	v.value = val
	v.isSet = true
}

func (v NullableCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapability(val *Capability) *NullableCapability {
	return &NullableCapability{value: val, isSet: true}
}

func (v NullableCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

