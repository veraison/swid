// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import "encoding/xml"

// Use models the use type.
type Use struct {
	val interface{}
}

/*
   $use /= optional
   $use /= required
   $use /= recommended
   $use /= uint / text
   optional=1
   required=2
   recommended=3
*/

// Use constants
const (
	UseOptional = uint64(iota + 1)
	UseRequired
	UseRecommended
	UseUnknown = ^uint64(0)
)

var (
	useToString = map[uint64]string{
		UseOptional:    "optional",
		UseRequired:    "required",
		UseRecommended: "recommended",
	}

	stringToUse = map[string]uint64{
		"optional":    UseOptional,
		"required":    UseRequired,
		"recommended": UseRecommended,
	}
)

// String returns the value of the Use receiver as a string
func (u Use) String() string {
	return codeStringer(u.val, useToString, "use")
}

// Check returns nil if the Use receiver is of type string or code-point
// (i.e., uint)
func (u Use) Check() error {
	return isStringOrCode(u.val, "use")
}

// MarshalCBOR encodes the Use receiver as code-point if possible,
// otherwise as string
func (u Use) MarshalCBOR() ([]byte, error) {
	return codeToCBOR(u.val, stringToUse)
}

// UnmarshalCBOR decodes the supplied data into a Use code-point if
// possible, otherwise as string
func (u *Use) UnmarshalCBOR(data []byte) error {
	return cborToCode(data, stringToUse, &u.val)
}

// MarshalJSON encodes the Use receiver as string
func (u Use) MarshalJSON() ([]byte, error) {
	return codeToJSON(u.val, useToString)
}

// UnmarshalJSON decodes the supplied data into a Use code-point if
// possible, otherwise as string
func (u *Use) UnmarshalJSON(data []byte) error {
	return jsonToCode(data, stringToUse, &u.val)
}

// MarshalXMLAttr encodes the Use receiver as XML attribute
func (u Use) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return codeToXMLAttr(name, u.val, useToString)
}

// UnmarshalXMLAttr decodes the supplied XML attribute into a Use
// code-point if possible, otherwise as string
func (u *Use) UnmarshalXMLAttr(attr xml.Attr) error {
	return xmlAttrToCode(attr, stringToUse, &u.val)
}
