// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/google/uuid"
)

// TagID is the type of a tag identifier. Allowed formats are string or
// a valid universally unique identifier (UUID) as defined by RFC4122.
type TagID struct {
	val interface{}
}

// NewTagID takes a UUID as either a string, byte array or google.uuid.UUID
// and returns a TagID
func NewTagID(v interface{}) *TagID {
	switch t := v.(type) {
	case string:
		tagID, _ := string2TagID(t)
		return tagID
	case []byte:
		tagID, _ := NewTagIDFromUUIDBytes(t)
		return tagID
	case uuid.UUID:
		return &TagID{v}
	default:
		return nil
	}
}

func string2TagID(s string) (*TagID, error) {
	if tagID, err := NewTagIDFromUUIDString(s); err == nil {
		return tagID, nil
	}

	if tagID, err := NewTagIDFromString(s); err == nil {
		return tagID, nil
	}

	return nil, errors.New("tag-id is neither a UUID nor a valid string")
}

// NewTagIDFromString takes an untyped string and returns a TagID
func NewTagIDFromString(s string) (*TagID, error) {
	if s == "" {
		return nil, errors.New("empty string")
	}
	return &TagID{s}, nil
}

// NewTagIDFromUUIDString takes an UUID in string form and returns a TagID
func NewTagIDFromUUIDString(s string) (*TagID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return nil, err
	}

	return &TagID{u}, nil
}

// NewTagIDFromUUIDBytes takes an UUID as byte array and returns a TagID
func NewTagIDFromUUIDBytes(b []byte) (*TagID, error) {
	u, err := uuid.FromBytes(b)
	if err != nil {
		return nil, err
	}

	return &TagID{u}, nil
}

// String returns the value of the TagID as string. If the TagID has type UUID,
// the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, is returned
func (t TagID) String() string {
	switch v := t.val.(type) {
	case string:
		return v
	case uuid.UUID:
		return v.String()
	default:
		return "unknown type for tag-id"
	}
}

// Returns TagID in URI representation according to CoSWID Spec
// useful for URI fields like link->href
func (t TagID) URI() string {
	switch v := t.val.(type) {
	case string:
		return v
	case uuid.UUID:
		return "swid:" + v.String()
	default:
		return "unknown type for tag-id"
	}
}

// MarshalXMLAttr encodes the TagID receiver as XML attribute
func (t TagID) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: t.String()}, nil
}

// UnmarshalXMLAttr decodes the supplied XML attribute into a TagID
// Note that this can only unmarshal to string.
func (t *TagID) UnmarshalXMLAttr(attr xml.Attr) error {
	tagID, err := string2TagID(attr.Value)
	if err != nil {
		return fmt.Errorf("error unmarshaling tag-id %q: %w", attr.Value, err)
	}

	*t = *tagID

	return nil
}

// MarshalJSON encodes the TagID receiver as JSON string
func (t TagID) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON decodes the supplied JSON data into a TagID.  If TagID is of
// type UUID, the string form, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, is
// expected.
func (t *TagID) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling tag-id: %w", err)
	}

	tagID, err := string2TagID(s)
	if err != nil {
		return fmt.Errorf("error unmarshaling tag-id %q: %w", s, err)
	}

	*t = *tagID

	return nil
}

// MarshalCBOR encodes the TagID receiver to CBOR
func (t TagID) MarshalCBOR() ([]byte, error) {
	return em.Marshal(t.val)
}

// UnmarshalCBOR decodes the supplied data into a TagID
func (t *TagID) UnmarshalCBOR(data []byte) error {
	var (
		v     interface{}
		err   error
		tagID *TagID
	)

	if err = cbor.Unmarshal(data, &v); err != nil {
		return err
	}

	switch typ := v.(type) {
	case string:
		tagID, err = NewTagIDFromString(typ)
	case []byte:
		tagID, err = NewTagIDFromUUIDBytes(typ)
	default:
		tagID, err = nil, fmt.Errorf("tag-id MUST be []byte or string; got %T", typ)
	}

	if err != nil {
		return fmt.Errorf("error unmarshaling tag-id: %w", err)
	}

	*t = *tagID

	return nil
}
