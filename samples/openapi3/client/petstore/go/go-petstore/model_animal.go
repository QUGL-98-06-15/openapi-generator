/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// checks if the Animal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Animal{}

// Animal struct for Animal
type Animal struct {
	ClassName string `json:"className"`
	Color *string `json:"color,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Animal Animal

// NewAnimal instantiates a new Animal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimal(className string) *Animal {
	this := Animal{}
	this.ClassName = className
	var color string = "red"
	this.Color = &color
	return &this
}

// NewAnimalWithDefaults instantiates a new Animal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimalWithDefaults() *Animal {
	this := Animal{}
	var color string = "red"
	this.Color = &color
	return &this
}

// GetClassName returns the ClassName field value
func (o *Animal) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *Animal) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *Animal) SetClassName(v string) {
	o.ClassName = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Animal) GetColor() string {
	if o == nil || isNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animal) GetColorOk() (*string, bool) {
	if o == nil || isNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Animal) HasColor() bool {
	if o != nil && !isNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Animal) SetColor(v string) {
	o.Color = &v
}

func (o Animal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Animal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["className"] = o.ClassName
	if !isNil(o.Color) {
		toSerialize["color"] = o.Color
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Animal) UnmarshalJSON(bytes []byte) (err error) {
	varAnimal := _Animal{}

	if err = json.Unmarshal(bytes, &varAnimal); err == nil {
		*o = Animal(varAnimal)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "className")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnimal struct {
	value *Animal
	isSet bool
}

func (v NullableAnimal) Get() *Animal {
	return v.value
}

func (v *NullableAnimal) Set(val *Animal) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimal) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimal(val *Animal) *NullableAnimal {
	return &NullableAnimal{value: val, isSet: true}
}

func (v NullableAnimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


