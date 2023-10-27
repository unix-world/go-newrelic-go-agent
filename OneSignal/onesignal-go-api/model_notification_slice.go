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

// NotificationSlice struct for NotificationSlice
type NotificationSlice struct {
	TotalCount *int32 `json:"total_count,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Notifications []NotificationWithMeta `json:"notifications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationSlice NotificationSlice

// NewNotificationSlice instantiates a new NotificationSlice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSlice() *NotificationSlice {
	this := NotificationSlice{}
	return &this
}

// NewNotificationSliceWithDefaults instantiates a new NotificationSlice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSliceWithDefaults() *NotificationSlice {
	this := NotificationSlice{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *NotificationSlice) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlice) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *NotificationSlice) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *NotificationSlice) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *NotificationSlice) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlice) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *NotificationSlice) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *NotificationSlice) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *NotificationSlice) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlice) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *NotificationSlice) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *NotificationSlice) SetLimit(v int32) {
	o.Limit = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *NotificationSlice) GetNotifications() []NotificationWithMeta {
	if o == nil || o.Notifications == nil {
		var ret []NotificationWithMeta
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlice) GetNotificationsOk() ([]NotificationWithMeta, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *NotificationSlice) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationWithMeta and assigns it to the Notifications field.
func (o *NotificationSlice) SetNotifications(v []NotificationWithMeta) {
	o.Notifications = v
}

func (o NotificationSlice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationSlice) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationSlice := _NotificationSlice{}

	if err = json.Unmarshal(bytes, &varNotificationSlice); err == nil {
		*o = NotificationSlice(varNotificationSlice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total_count")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "notifications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationSlice struct {
	value *NotificationSlice
	isSet bool
}

func (v NullableNotificationSlice) Get() *NotificationSlice {
	return v.value
}

func (v *NullableNotificationSlice) Set(val *NotificationSlice) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSlice) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSlice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSlice(val *NotificationSlice) *NullableNotificationSlice {
	return &NullableNotificationSlice{value: val, isSet: true}
}

func (v NullableNotificationSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSlice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

