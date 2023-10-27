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

// ExportPlayersRequestBody struct for ExportPlayersRequestBody
type ExportPlayersRequestBody struct {
	// Additional fields that you wish to include. Currently supports location, country, rooted, notification_types, ip, external_user_id, web_auth, and web_p256.
	ExtraFields []string `json:"extra_fields,omitempty"`
	// Export all devices with a last_active timestamp greater than this time.  Unixtime in seconds.
	LastActiveSince *string `json:"last_active_since,omitempty"`
	// Export all devices belonging to the segment.
	SegmentName *string `json:"segment_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportPlayersRequestBody ExportPlayersRequestBody

// NewExportPlayersRequestBody instantiates a new ExportPlayersRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportPlayersRequestBody() *ExportPlayersRequestBody {
	this := ExportPlayersRequestBody{}
	return &this
}

// NewExportPlayersRequestBodyWithDefaults instantiates a new ExportPlayersRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportPlayersRequestBodyWithDefaults() *ExportPlayersRequestBody {
	this := ExportPlayersRequestBody{}
	return &this
}

// GetExtraFields returns the ExtraFields field value if set, zero value otherwise.
func (o *ExportPlayersRequestBody) GetExtraFields() []string {
	if o == nil || o.ExtraFields == nil {
		var ret []string
		return ret
	}
	return o.ExtraFields
}

// GetExtraFieldsOk returns a tuple with the ExtraFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPlayersRequestBody) GetExtraFieldsOk() ([]string, bool) {
	if o == nil || o.ExtraFields == nil {
		return nil, false
	}
	return o.ExtraFields, true
}

// HasExtraFields returns a boolean if a field has been set.
func (o *ExportPlayersRequestBody) HasExtraFields() bool {
	if o != nil && o.ExtraFields != nil {
		return true
	}

	return false
}

// SetExtraFields gets a reference to the given []string and assigns it to the ExtraFields field.
func (o *ExportPlayersRequestBody) SetExtraFields(v []string) {
	o.ExtraFields = v
}

// GetLastActiveSince returns the LastActiveSince field value if set, zero value otherwise.
func (o *ExportPlayersRequestBody) GetLastActiveSince() string {
	if o == nil || o.LastActiveSince == nil {
		var ret string
		return ret
	}
	return *o.LastActiveSince
}

// GetLastActiveSinceOk returns a tuple with the LastActiveSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPlayersRequestBody) GetLastActiveSinceOk() (*string, bool) {
	if o == nil || o.LastActiveSince == nil {
		return nil, false
	}
	return o.LastActiveSince, true
}

// HasLastActiveSince returns a boolean if a field has been set.
func (o *ExportPlayersRequestBody) HasLastActiveSince() bool {
	if o != nil && o.LastActiveSince != nil {
		return true
	}

	return false
}

// SetLastActiveSince gets a reference to the given string and assigns it to the LastActiveSince field.
func (o *ExportPlayersRequestBody) SetLastActiveSince(v string) {
	o.LastActiveSince = &v
}

// GetSegmentName returns the SegmentName field value if set, zero value otherwise.
func (o *ExportPlayersRequestBody) GetSegmentName() string {
	if o == nil || o.SegmentName == nil {
		var ret string
		return ret
	}
	return *o.SegmentName
}

// GetSegmentNameOk returns a tuple with the SegmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPlayersRequestBody) GetSegmentNameOk() (*string, bool) {
	if o == nil || o.SegmentName == nil {
		return nil, false
	}
	return o.SegmentName, true
}

// HasSegmentName returns a boolean if a field has been set.
func (o *ExportPlayersRequestBody) HasSegmentName() bool {
	if o != nil && o.SegmentName != nil {
		return true
	}

	return false
}

// SetSegmentName gets a reference to the given string and assigns it to the SegmentName field.
func (o *ExportPlayersRequestBody) SetSegmentName(v string) {
	o.SegmentName = &v
}

func (o ExportPlayersRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraFields != nil {
		toSerialize["extra_fields"] = o.ExtraFields
	}
	if o.LastActiveSince != nil {
		toSerialize["last_active_since"] = o.LastActiveSince
	}
	if o.SegmentName != nil {
		toSerialize["segment_name"] = o.SegmentName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportPlayersRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varExportPlayersRequestBody := _ExportPlayersRequestBody{}

	if err = json.Unmarshal(bytes, &varExportPlayersRequestBody); err == nil {
		*o = ExportPlayersRequestBody(varExportPlayersRequestBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "extra_fields")
		delete(additionalProperties, "last_active_since")
		delete(additionalProperties, "segment_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportPlayersRequestBody struct {
	value *ExportPlayersRequestBody
	isSet bool
}

func (v NullableExportPlayersRequestBody) Get() *ExportPlayersRequestBody {
	return v.value
}

func (v *NullableExportPlayersRequestBody) Set(val *ExportPlayersRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableExportPlayersRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableExportPlayersRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportPlayersRequestBody(val *ExportPlayersRequestBody) *NullableExportPlayersRequestBody {
	return &NullableExportPlayersRequestBody{value: val, isSet: true}
}

func (v NullableExportPlayersRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportPlayersRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


