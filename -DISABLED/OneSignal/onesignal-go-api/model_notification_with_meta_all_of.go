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

// NotificationWithMetaAllOf struct for NotificationWithMetaAllOf
type NotificationWithMetaAllOf struct {
	// Number of notifications that have not been sent out yet. This can mean either our system is still processing the notification or you have delayed options set.
	Remaining *int32 `json:"remaining,omitempty"`
	// Number of notifications that were successfully delivered.
	Successful *int32 `json:"successful,omitempty"`
	// Number of notifications that could not be delivered due to those devices being unsubscribed.
	Failed *int32 `json:"failed,omitempty"`
	// Number of notifications that could not be delivered due to an error. You can find more information by viewing the notification in the dashboard.
	Errored *int32 `json:"errored,omitempty"`
	// Number of users who have clicked / tapped on your notification.
	Converted *int32 `json:"converted,omitempty"`
	// Unix timestamp indicating when the notification was created.
	QueuedAt *int64 `json:"queued_at,omitempty"`
	// Unix timestamp indicating when notification delivery should begin.
	SendAfter NullableInt64 `json:"send_after,omitempty"`
	// Unix timestamp indicating when notification delivery completed. The delivery duration from start to finish can be calculated with completed_at - send_after.
	CompletedAt NullableInt64 `json:"completed_at,omitempty"`
	PlatformDeliveryStats *PlatformDeliveryData `json:"platform_delivery_stats,omitempty"`
	// Confirmed Deliveries number of devices that received the push notification. Paid Feature Only. Free accounts will see 0.
	Received NullableInt32 `json:"received,omitempty"`
	// number of push notifications sent per minute. Paid Feature Only. If throttling is not enabled for the app or the notification, and for free accounts, null is returned. Refer to Throttling for more details.
	ThrottleRatePerMinute NullableInt32 `json:"throttle_rate_per_minute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationWithMetaAllOf NotificationWithMetaAllOf

// NewNotificationWithMetaAllOf instantiates a new NotificationWithMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationWithMetaAllOf() *NotificationWithMetaAllOf {
	this := NotificationWithMetaAllOf{}
	return &this
}

// NewNotificationWithMetaAllOfWithDefaults instantiates a new NotificationWithMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithMetaAllOfWithDefaults() *NotificationWithMetaAllOf {
	this := NotificationWithMetaAllOf{}
	return &this
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetRemaining() int32 {
	if o == nil || o.Remaining == nil {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetRemainingOk() (*int32, bool) {
	if o == nil || o.Remaining == nil {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasRemaining() bool {
	if o != nil && o.Remaining != nil {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *NotificationWithMetaAllOf) SetRemaining(v int32) {
	o.Remaining = &v
}

// GetSuccessful returns the Successful field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetSuccessful() int32 {
	if o == nil || o.Successful == nil {
		var ret int32
		return ret
	}
	return *o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetSuccessfulOk() (*int32, bool) {
	if o == nil || o.Successful == nil {
		return nil, false
	}
	return o.Successful, true
}

// HasSuccessful returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasSuccessful() bool {
	if o != nil && o.Successful != nil {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given int32 and assigns it to the Successful field.
func (o *NotificationWithMetaAllOf) SetSuccessful(v int32) {
	o.Successful = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetFailed() int32 {
	if o == nil || o.Failed == nil {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetFailedOk() (*int32, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *NotificationWithMetaAllOf) SetFailed(v int32) {
	o.Failed = &v
}

// GetErrored returns the Errored field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetErrored() int32 {
	if o == nil || o.Errored == nil {
		var ret int32
		return ret
	}
	return *o.Errored
}

// GetErroredOk returns a tuple with the Errored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetErroredOk() (*int32, bool) {
	if o == nil || o.Errored == nil {
		return nil, false
	}
	return o.Errored, true
}

// HasErrored returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasErrored() bool {
	if o != nil && o.Errored != nil {
		return true
	}

	return false
}

// SetErrored gets a reference to the given int32 and assigns it to the Errored field.
func (o *NotificationWithMetaAllOf) SetErrored(v int32) {
	o.Errored = &v
}

// GetConverted returns the Converted field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetConverted() int32 {
	if o == nil || o.Converted == nil {
		var ret int32
		return ret
	}
	return *o.Converted
}

// GetConvertedOk returns a tuple with the Converted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetConvertedOk() (*int32, bool) {
	if o == nil || o.Converted == nil {
		return nil, false
	}
	return o.Converted, true
}

// HasConverted returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasConverted() bool {
	if o != nil && o.Converted != nil {
		return true
	}

	return false
}

// SetConverted gets a reference to the given int32 and assigns it to the Converted field.
func (o *NotificationWithMetaAllOf) SetConverted(v int32) {
	o.Converted = &v
}

// GetQueuedAt returns the QueuedAt field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetQueuedAt() int64 {
	if o == nil || o.QueuedAt == nil {
		var ret int64
		return ret
	}
	return *o.QueuedAt
}

// GetQueuedAtOk returns a tuple with the QueuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetQueuedAtOk() (*int64, bool) {
	if o == nil || o.QueuedAt == nil {
		return nil, false
	}
	return o.QueuedAt, true
}

// HasQueuedAt returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasQueuedAt() bool {
	if o != nil && o.QueuedAt != nil {
		return true
	}

	return false
}

// SetQueuedAt gets a reference to the given int64 and assigns it to the QueuedAt field.
func (o *NotificationWithMetaAllOf) SetQueuedAt(v int64) {
	o.QueuedAt = &v
}

// GetSendAfter returns the SendAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationWithMetaAllOf) GetSendAfter() int64 {
	if o == nil || o.SendAfter.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SendAfter.Get()
}

// GetSendAfterOk returns a tuple with the SendAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationWithMetaAllOf) GetSendAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendAfter.Get(), o.SendAfter.IsSet()
}

// HasSendAfter returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasSendAfter() bool {
	if o != nil && o.SendAfter.IsSet() {
		return true
	}

	return false
}

// SetSendAfter gets a reference to the given NullableInt64 and assigns it to the SendAfter field.
func (o *NotificationWithMetaAllOf) SetSendAfter(v int64) {
	o.SendAfter.Set(&v)
}
// SetSendAfterNil sets the value for SendAfter to be an explicit nil
func (o *NotificationWithMetaAllOf) SetSendAfterNil() {
	o.SendAfter.Set(nil)
}

// UnsetSendAfter ensures that no value is present for SendAfter, not even an explicit nil
func (o *NotificationWithMetaAllOf) UnsetSendAfter() {
	o.SendAfter.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationWithMetaAllOf) GetCompletedAt() int64 {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationWithMetaAllOf) GetCompletedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableInt64 and assigns it to the CompletedAt field.
func (o *NotificationWithMetaAllOf) SetCompletedAt(v int64) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *NotificationWithMetaAllOf) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *NotificationWithMetaAllOf) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetPlatformDeliveryStats returns the PlatformDeliveryStats field value if set, zero value otherwise.
func (o *NotificationWithMetaAllOf) GetPlatformDeliveryStats() PlatformDeliveryData {
	if o == nil || o.PlatformDeliveryStats == nil {
		var ret PlatformDeliveryData
		return ret
	}
	return *o.PlatformDeliveryStats
}

// GetPlatformDeliveryStatsOk returns a tuple with the PlatformDeliveryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWithMetaAllOf) GetPlatformDeliveryStatsOk() (*PlatformDeliveryData, bool) {
	if o == nil || o.PlatformDeliveryStats == nil {
		return nil, false
	}
	return o.PlatformDeliveryStats, true
}

// HasPlatformDeliveryStats returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasPlatformDeliveryStats() bool {
	if o != nil && o.PlatformDeliveryStats != nil {
		return true
	}

	return false
}

// SetPlatformDeliveryStats gets a reference to the given PlatformDeliveryData and assigns it to the PlatformDeliveryStats field.
func (o *NotificationWithMetaAllOf) SetPlatformDeliveryStats(v PlatformDeliveryData) {
	o.PlatformDeliveryStats = &v
}

// GetReceived returns the Received field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationWithMetaAllOf) GetReceived() int32 {
	if o == nil || o.Received.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Received.Get()
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationWithMetaAllOf) GetReceivedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Received.Get(), o.Received.IsSet()
}

// HasReceived returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasReceived() bool {
	if o != nil && o.Received.IsSet() {
		return true
	}

	return false
}

// SetReceived gets a reference to the given NullableInt32 and assigns it to the Received field.
func (o *NotificationWithMetaAllOf) SetReceived(v int32) {
	o.Received.Set(&v)
}
// SetReceivedNil sets the value for Received to be an explicit nil
func (o *NotificationWithMetaAllOf) SetReceivedNil() {
	o.Received.Set(nil)
}

// UnsetReceived ensures that no value is present for Received, not even an explicit nil
func (o *NotificationWithMetaAllOf) UnsetReceived() {
	o.Received.Unset()
}

// GetThrottleRatePerMinute returns the ThrottleRatePerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationWithMetaAllOf) GetThrottleRatePerMinute() int32 {
	if o == nil || o.ThrottleRatePerMinute.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ThrottleRatePerMinute.Get()
}

// GetThrottleRatePerMinuteOk returns a tuple with the ThrottleRatePerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationWithMetaAllOf) GetThrottleRatePerMinuteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThrottleRatePerMinute.Get(), o.ThrottleRatePerMinute.IsSet()
}

// HasThrottleRatePerMinute returns a boolean if a field has been set.
func (o *NotificationWithMetaAllOf) HasThrottleRatePerMinute() bool {
	if o != nil && o.ThrottleRatePerMinute.IsSet() {
		return true
	}

	return false
}

// SetThrottleRatePerMinute gets a reference to the given NullableInt32 and assigns it to the ThrottleRatePerMinute field.
func (o *NotificationWithMetaAllOf) SetThrottleRatePerMinute(v int32) {
	o.ThrottleRatePerMinute.Set(&v)
}
// SetThrottleRatePerMinuteNil sets the value for ThrottleRatePerMinute to be an explicit nil
func (o *NotificationWithMetaAllOf) SetThrottleRatePerMinuteNil() {
	o.ThrottleRatePerMinute.Set(nil)
}

// UnsetThrottleRatePerMinute ensures that no value is present for ThrottleRatePerMinute, not even an explicit nil
func (o *NotificationWithMetaAllOf) UnsetThrottleRatePerMinute() {
	o.ThrottleRatePerMinute.Unset()
}

func (o NotificationWithMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Remaining != nil {
		toSerialize["remaining"] = o.Remaining
	}
	if o.Successful != nil {
		toSerialize["successful"] = o.Successful
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	if o.Errored != nil {
		toSerialize["errored"] = o.Errored
	}
	if o.Converted != nil {
		toSerialize["converted"] = o.Converted
	}
	if o.QueuedAt != nil {
		toSerialize["queued_at"] = o.QueuedAt
	}
	if o.SendAfter.IsSet() {
		toSerialize["send_after"] = o.SendAfter.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.PlatformDeliveryStats != nil {
		toSerialize["platform_delivery_stats"] = o.PlatformDeliveryStats
	}
	if o.Received.IsSet() {
		toSerialize["received"] = o.Received.Get()
	}
	if o.ThrottleRatePerMinute.IsSet() {
		toSerialize["throttle_rate_per_minute"] = o.ThrottleRatePerMinute.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationWithMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationWithMetaAllOf := _NotificationWithMetaAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationWithMetaAllOf); err == nil {
		*o = NotificationWithMetaAllOf(varNotificationWithMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "remaining")
		delete(additionalProperties, "successful")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "errored")
		delete(additionalProperties, "converted")
		delete(additionalProperties, "queued_at")
		delete(additionalProperties, "send_after")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "platform_delivery_stats")
		delete(additionalProperties, "received")
		delete(additionalProperties, "throttle_rate_per_minute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationWithMetaAllOf struct {
	value *NotificationWithMetaAllOf
	isSet bool
}

func (v NullableNotificationWithMetaAllOf) Get() *NotificationWithMetaAllOf {
	return v.value
}

func (v *NullableNotificationWithMetaAllOf) Set(val *NotificationWithMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationWithMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationWithMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationWithMetaAllOf(val *NotificationWithMetaAllOf) *NullableNotificationWithMetaAllOf {
	return &NullableNotificationWithMetaAllOf{value: val, isSet: true}
}

func (v NullableNotificationWithMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationWithMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

