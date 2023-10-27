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

// ExportPlayersSuccessResponse struct for ExportPlayersSuccessResponse
type ExportPlayersSuccessResponse struct {
	CsvFileUrl *string `json:"csv_file_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportPlayersSuccessResponse ExportPlayersSuccessResponse

// NewExportPlayersSuccessResponse instantiates a new ExportPlayersSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportPlayersSuccessResponse() *ExportPlayersSuccessResponse {
	this := ExportPlayersSuccessResponse{}
	return &this
}

// NewExportPlayersSuccessResponseWithDefaults instantiates a new ExportPlayersSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportPlayersSuccessResponseWithDefaults() *ExportPlayersSuccessResponse {
	this := ExportPlayersSuccessResponse{}
	return &this
}

// GetCsvFileUrl returns the CsvFileUrl field value if set, zero value otherwise.
func (o *ExportPlayersSuccessResponse) GetCsvFileUrl() string {
	if o == nil || o.CsvFileUrl == nil {
		var ret string
		return ret
	}
	return *o.CsvFileUrl
}

// GetCsvFileUrlOk returns a tuple with the CsvFileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPlayersSuccessResponse) GetCsvFileUrlOk() (*string, bool) {
	if o == nil || o.CsvFileUrl == nil {
		return nil, false
	}
	return o.CsvFileUrl, true
}

// HasCsvFileUrl returns a boolean if a field has been set.
func (o *ExportPlayersSuccessResponse) HasCsvFileUrl() bool {
	if o != nil && o.CsvFileUrl != nil {
		return true
	}

	return false
}

// SetCsvFileUrl gets a reference to the given string and assigns it to the CsvFileUrl field.
func (o *ExportPlayersSuccessResponse) SetCsvFileUrl(v string) {
	o.CsvFileUrl = &v
}

func (o ExportPlayersSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsvFileUrl != nil {
		toSerialize["csv_file_url"] = o.CsvFileUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportPlayersSuccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varExportPlayersSuccessResponse := _ExportPlayersSuccessResponse{}

	if err = json.Unmarshal(bytes, &varExportPlayersSuccessResponse); err == nil {
		*o = ExportPlayersSuccessResponse(varExportPlayersSuccessResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csv_file_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportPlayersSuccessResponse struct {
	value *ExportPlayersSuccessResponse
	isSet bool
}

func (v NullableExportPlayersSuccessResponse) Get() *ExportPlayersSuccessResponse {
	return v.value
}

func (v *NullableExportPlayersSuccessResponse) Set(val *ExportPlayersSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExportPlayersSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExportPlayersSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportPlayersSuccessResponse(val *ExportPlayersSuccessResponse) *NullableExportPlayersSuccessResponse {
	return &NullableExportPlayersSuccessResponse{value: val, isSet: true}
}

func (v NullableExportPlayersSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportPlayersSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


