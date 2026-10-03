// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // shared boilerplate
package swid

import "encoding/xml"

// Ownership models the ownership type
type Ownership struct {
	val interface{}
}

/*
   $ownership /= shared
   $ownership /= private
   $ownership /= abandon
   $ownership /= int / text
   abandon=1
   private=2
   shared=3
*/

// Ownership constants
const (
	OwnershipAbandon = int64(iota + 1)
	OwnershipPrivate
	OwnershipShared
	OwnershipUnknown = ^int64(0)
)

var (
	ownershipToString = map[int64]string{
		OwnershipShared:  "shared",
		OwnershipPrivate: "private",
		OwnershipAbandon: "abandon",
	}

	stringToOwnership = map[string]int64{
		"shared":  OwnershipShared,
		"private": OwnershipPrivate,
		"abandon": OwnershipAbandon,
	}
)

// String returns the value of the Ownership receiver as a string
func (o Ownership) String() string {
	return codeStringer(o.val, ownershipToString, "ownership")
}

// Check returns nil if the Ownership receiver is of type string or code-point
// (i.e., uint)
func (o Ownership) Check() error {
	return isStringOrCode(o.val, "ownership")
}

// MarshalCBOR encodes the Ownership receiver as code-point if possible,
// otherwise as string
func (o Ownership) MarshalCBOR() ([]byte, error) {
	return codeToCBOR(o.val, stringToOwnership)
}

// UnmarshalCBOR decodes the supplied data into an Ownership code-point if
// possible, otherwise as string
func (o *Ownership) UnmarshalCBOR(data []byte) error {
	v, err := cborToCode(data, stringToOwnership)
	if err != nil {
		return err
	}

	o.val = v

	return nil
}

// MarshalJSON encodes the Ownership receiver as string
func (o Ownership) MarshalJSON() ([]byte, error) {
	return codeToJSON(o.val, ownershipToString)
}

// UnmarshalJSON decodes the supplied data into an Ownership code-point if
// possible, otherwise as string
func (o *Ownership) UnmarshalJSON(data []byte) error {
	v, err := jsonToCode(data, stringToOwnership)
	if err != nil {
		return err
	}

	o.val = v

	return nil
}

// MarshalXMLAttr encodes the Ownership receiver as XML attribute
func (o Ownership) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return codeToXMLAttr(name, o.val, ownershipToString)
}

// UnmarshalXMLAttr decodes the supplied XML attribute into an Ownership
// code-point if possible, otherwise as string
func (o *Ownership) UnmarshalXMLAttr(attr xml.Attr) error {
	v, err := xmlAttrToCode(attr, stringToOwnership)
	if err != nil {
		return err
	}

	o.val = v

	return nil
}
