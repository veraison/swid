// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagID_16Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	}

	expected := "00010001000100010001000100010001"

	actual := NewTagID(tv)

	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual.String())
}

func TestTagID_15Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
	}

	err := checkTagID(tv)

	assert.EqualError(t, err, "binary tag-id MUST be 16 bytes")
}

func TestTagID_17Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00,
	}

	err := checkTagID(tv)

	assert.EqualError(t, err, "binary tag-id MUST be 16 bytes")
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

	err := checkTagID(tv)

	assert.EqualError(t, err, "tag-id MUST be []byte or string; got struct { a int; b string }")
}

func TestTagID_UnmarshalXMLAttrString(t *testing.T) {
	v := "example.acme.roadrunner-sw-v1-0-0"

	tv := xml.Attr{
		Name:  xml.Name{Local: "tagId"},
		Value: v,
	}

	expected := *NewTagID(v)

	var actual TagID

	err := actual.UnmarshalXMLAttr(tv)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestTagID_MarshalXMLAttrString(t *testing.T) {
	v := "example.acme.roadrunner-sw-v1-0-0"

	tv := *NewTagID(v)

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

	tv := *NewTagID(v)

	_, err := tv.MarshalXMLAttr(xml.Name{Local: "tagId"})

	assert.EqualError(t, err, "only tag-id of type string can be serialized to XML")
}

func TestTagID_MarshalJSONBytes(t *testing.T) {
	v := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	}

	tv := *NewTagID(v)

	_, err := tv.MarshalJSON()

	assert.EqualError(t, err, "only tag-id of type string can be serialized to JSON")
}

func TestTagID_UnMarshalJSONUnhandled(t *testing.T) {
	tv := []byte(`{ "k": "0" }`)

	var actual TagID

	err := actual.UnmarshalJSON(tv)

	assert.EqualError(t, err, "expecting string, found map[string]interface {} instead")
}
