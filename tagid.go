// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/fxamacker/cbor"
)

// TagID is the type of a tag identifier. Allowed formats (enforced via
// checkTagID) are string or [16]byte
type TagID struct {
	val interface{}
}

// NewTagID returns a TagID initialized with the supplied value v
// v is either a string or a [16]byte
func NewTagID(v interface{}) *TagID {
	if checkTagID(v) != nil {
		return nil
	}
	return &TagID{v}
}

// String returns the value of the TagID as string. If the TagID has type
// [16]byte the Base 16 encoding is returned
func (t TagID) String() string {
	switch v := t.val.(type) {
	case string:
		return v
	case []byte:
		return hex.EncodeToString(v)
	default:
		return "unknown type for tag-id"
	}
}

func checkTagID(v interface{}) error {
	switch t := v.(type) {
	case string:
	case []byte:
		if len(t) != 16 {
			return errors.New("binary tag-id MUST be 16 bytes")
		}
	default:
		return fmt.Errorf("tag-id MUST be []byte or string; got %T", v)
	}

	return nil
}

func (t TagID) isString() bool {
	switch t.val.(type) {
	case string:
		return true
	}
	return false
}

// MarshalXMLAttr encodes the TagID receiver as XML attribute
func (t TagID) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if !t.isString() {
		return xml.Attr{}, errors.New("only tag-id of type string can be serialized to XML")
	}
	return xml.Attr{Name: name, Value: t.String()}, nil
}

// UnmarshalXMLAttr decodes the supplied XML attribute into a TagID
// Note that this can only unmarshal to string.
func (t *TagID) UnmarshalXMLAttr(attr xml.Attr) error {
	t.val = attr.Value
	return nil
}

// MarshalJSON encodes the TagID receiver as JSON string
func (t TagID) MarshalJSON() ([]byte, error) {
	if !t.isString() {
		return nil, errors.New("only tag-id of type string can be serialized to JSON")
	}

	return json.Marshal(t.val)
}

// UnmarshalJSON decodes the supplied JSON data into a TagID
// Note that this can only unmarshal to string.
func (t *TagID) UnmarshalJSON(data []byte) error {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch s := v.(type) {
	case string:
		t.val = s
		return nil
	default:
		return fmt.Errorf("expecting string, found %T instead", s)
	}
}

// MarshalCBOR encodes the TagID receiver to CBOR
func (t TagID) MarshalCBOR() ([]byte, error) {
	return em.Marshal(t.val)
}

// UnmarshalCBOR decodes the supplied data into a TagID
func (t *TagID) UnmarshalCBOR(data []byte) error {
	var v interface{}

	if err := cbor.Unmarshal(data, &v); err != nil {
		return err
	}

	if err := checkTagID(v); err != nil {
		return err
	}

	t.val = v

	return nil
}
