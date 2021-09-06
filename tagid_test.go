// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagID_NewFromUUIDString(t *testing.T) {
	tv := "00010001-0001-0001-0001-000100010001"

	expected := "00010001-0001-0001-0001-000100010001"

	actual := NewTagID(tv)

	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual.String())
}

func TestTagID_NewTagID_empty(t *testing.T) {
	tv := ""

	actual := NewTagID(tv)

	assert.Nil(t, actual)
}

func TestTagID_16Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	}

	expected := "00010001-0001-0001-0001-000100010001"

	actual := NewTagID(tv)

	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual.String())
}

func TestTagID_15Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
	}

	_, err := NewTagIDFromUUIDBytes(tv)

	assert.EqualError(t, err, "invalid UUID (got 15 bytes)")
}

func TestTagID_17Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00,
	}

	_, err := NewTagIDFromUUIDBytes(tv)

	assert.EqualError(t, err, "invalid UUID (got 17 bytes)")
}

func TestTagID_String(t *testing.T) {
	tv := "example.acme.roadrunner-sw-v1-0-0"

	actual := NewTagID(tv)

	assert.NotNil(t, actual)
	assert.Equal(t, tv, actual.String())
}

func TestTagID_UnhandledType(t *testing.T) {
	tv := struct {
		a int
		b string
	}{
		a: 1,
		b: "one",
	}

	actual := NewTagID(tv)

	assert.Nil(t, actual)
}

func TestTagID_UnmarshalXMLAttrString_empty(t *testing.T) {
	v := ""

	tv := xml.Attr{
		Name:  xml.Name{Local: "tagId"},
		Value: v,
	}

	expectedErr := `error unmarshaling tag-id "": tag-id is neither a UUID nor a valid string`

	var actual TagID

	err := actual.UnmarshalXMLAttr(tv)

	assert.EqualError(t, err, expectedErr)
}

func TestTagID_UnmarshalXMLAttrString(t *testing.T) {
	v := "example.acme.roadrunner-sw-v1-0-0"

	tv := xml.Attr{
		Name:  xml.Name{Local: "tagId"},
		Value: v,
	}

	expected := NewTagID(v)
	require.NotNil(t, expected)

	var actual TagID

	err := actual.UnmarshalXMLAttr(tv)

	assert.Nil(t, err)
	assert.Equal(t, *expected, actual)
}

func TestTagID_MarshalXMLAttrString(t *testing.T) {
	v := "example.acme.roadrunner-sw-v1-0-0"

	tv := NewTagID(v)
	require.NotNil(t, tv)

	expected := xml.Attr{
		Name:  xml.Name{Local: "tagId"},
		Value: v,
	}

	actual, err := tv.MarshalXMLAttr(xml.Name{Local: "tagId"})

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestTagID_MarshalXMLAttrBytes(t *testing.T) {
	v := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	}

	tv := NewTagID(v)
	require.NotNil(t, tv)

	expected := "00010001-0001-0001-0001-000100010001"

	actual, err := tv.MarshalXMLAttr(xml.Name{Local: "tagId"})

	assert.Nil(t, err)
	assert.Equal(t, expected, actual.Value)
}

func TestTagID_MarshalJSONBytes(t *testing.T) {
	v := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	}

	tv := NewTagID(v)
	require.NotNil(t, tv)

	expected := `"00010001-0001-0001-0001-000100010001"`

	actual, err := tv.MarshalJSON()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestTagID_UnmarshalJSONUnhandled(t *testing.T) {
	tv := []byte(`{ "k": "0" }`)

	var actual TagID

	expectedErr := "error unmarshaling tag-id: json: cannot unmarshal object into Go value of type string"

	err := actual.UnmarshalJSON(tv)

	assert.EqualError(t, err, expectedErr)
}

func TestTagID_UnmarshalJSON_empty(t *testing.T) {
	tv := []byte(`""`)

	expectedErr := `error unmarshaling tag-id "": tag-id is neither a UUID nor a valid string`

	var actual TagID

	err := actual.UnmarshalJSON(tv)

	assert.EqualError(t, err, expectedErr)
}

func TestTagID_UnmarshalCBOR_EOF(t *testing.T) {
	tv := []byte{}

	expectedErr := `EOF`

	var actual TagID

	err := actual.UnmarshalCBOR(tv)

	assert.EqualError(t, err, expectedErr)
}

func TestTagID_UnmarshalCBOR_unhandled_type(t *testing.T) {
	tv := []byte{0xf6} // null

	expectedErr := `error unmarshaling tag-id: tag-id MUST be []byte or string; got <nil>`

	var actual TagID

	err := actual.UnmarshalCBOR(tv)

	assert.EqualError(t, err, expectedErr)
}

func TestTagID_UnmarshalCBOR_empty_bytes(t *testing.T) {
	tv := []byte{0x40} // bytes(0)

	expectedErr := `error unmarshaling tag-id: invalid UUID (got 0 bytes)`

	var actual TagID

	err := actual.UnmarshalCBOR(tv)

	assert.EqualError(t, err, expectedErr)
}
