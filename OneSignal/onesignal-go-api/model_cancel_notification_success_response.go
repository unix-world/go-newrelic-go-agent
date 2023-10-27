/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// CancelNotificationSuccessResponse struct for CancelNotificationSuccessResponse
type CancelNotificationSuccessResponse struct {
	Success *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelNotificationSuccessResponse CancelNotificationSuccessResponse

// NewCancelNotificationSuccessResponse instantiates a new CancelNotificationSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelNotificationSuccessResponse() *CancelNotificationSuccessResponse {
	this := CancelNotificationSuccessResponse{}
	return &this
}

// NewCancelNotificationSuccessResponseWithDefaults instantiates a new CancelNotificationSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelNotificationSuccessResponseWithDefaults() *CancelNotificationSuccessResponse {
	this := CancelNotificationSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CancelNotificationSuccessResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelNotificationSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CancelNotificationSuccessResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CancelNotificationSuccessResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o CancelNotificationSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CancelNotificationSuccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCancelNotificationSuccessResponse := _CancelNotificationSuccessResponse{}

	if err = json.Unmarshal(bytes, &varCancelNotificationSuccessResponse); err == nil {
		*o = CancelNotificationSuccessResponse(varCancelNotificationSuccessResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelNotificationSuccessResponse struct {
	value *CancelNotificationSuccessResponse
	isSet bool
}

func (v NullableCancelNotificationSuccessResponse) Get() *CancelNotificationSuccessResponse {
	return v.value
}

func (v *NullableCancelNotificationSuccessResponse) Set(val *CancelNotificationSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelNotificationSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelNotificationSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelNotificationSuccessResponse(val *CancelNotificationSuccessResponse) *NullableCancelNotificationSuccessResponse {
	return &NullableCancelNotificationSuccessResponse{value: val, isSet: true}
}

func (v NullableCancelNotificationSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelNotificationSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

