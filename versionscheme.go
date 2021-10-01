// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/xml"
	"fmt"
)

// VersionScheme models the version-scheme type.
type VersionScheme struct {
	val interface{}
}

/*
  $version-scheme /= multipartnumeric
  $version-scheme /= multipartnumeric-suffix
  $version-scheme /= alphanumeric
  $version-scheme /= decimal
  $version-scheme /= semver
  $version-scheme /= uint / text
  multipartnumeric = 1
  multipartnumeric-suffix = 2
  alphanumeric = 3
  decimal = 4
  semver = 16384
*/

// VersionScheme constants
const (
	VersionSchemeMultipartNumeric = uint64(iota + 1)
	VersionSchemeMultipartNumericSuffix
	VersionSchemeAlphaNumeric
	VersionSchemeDecimal
	VersionSchemeSemVer  = 16384
	VersionSchemeUnknown = ^uint64(0)
)

var (
	versionSchemeToString = map[uint64]string{
		VersionSchemeMultipartNumeric:       "multipartnumeric",
		VersionSchemeMultipartNumericSuffix: "multipartnumeric+suffix",
		VersionSchemeAlphaNumeric:           "alphanumeric",
		VersionSchemeDecimal:                "decimal",
		VersionSchemeSemVer:                 "semver",
	}

	stringToVersionScheme = map[string]uint64{
		"multipartnumeric":        VersionSchemeMultipartNumeric,
		"multipartnumeric+suffix": VersionSchemeMultipartNumericSuffix,
		"alphanumeric":            VersionSchemeAlphaNumeric,
		"decimal":                 VersionSchemeDecimal,
		"semver":                  VersionSchemeSemVer,
	}
)

// String returns the value of the VersionScheme receiver as a string
func (vs VersionScheme) String() string {
	return codeStringer(vs.val, versionSchemeToString, "version-scheme")
}

// Check returns nil if the VersionScheme receiver is of type string or code-point
// (i.e., uint)
func (vs VersionScheme) Check() error {
	return isStringOrCode(vs.val, "version-scheme")
}

// MarshalCBOR encodes the VersionScheme receiver as code-point if possible,
// otherwise as string
func (vs VersionScheme) MarshalCBOR() ([]byte, error) {
	return codeToCBOR(vs.val, stringToVersionScheme)
}

// UnmarshalCBOR decodes the supplied data into a VersionScheme code-point if
// possible, otherwise as string
func (vs *VersionScheme) UnmarshalCBOR(data []byte) error {
	return cborToCode(data, stringToVersionScheme, &vs.val)
}

// MarshalJSON encodes the VersionScheme receiver as string
func (vs VersionScheme) MarshalJSON() ([]byte, error) {
	return codeToJSON(vs.val, versionSchemeToString)
}

// UnmarshalJSON decodes the supplied data into an VersionScheme code-point if
// possible, otherwise as string
func (vs *VersionScheme) UnmarshalJSON(data []byte) error {
	return jsonToCode(data, stringToVersionScheme, &vs.val)
}

// MarshalXMLAttr encodes the VersionScheme receiver as XML attribute
func (vs VersionScheme) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return codeToXMLAttr(name, vs.val, versionSchemeToString)
}

// UnmarshalXMLAttr decodes the supplied XML attribute into a VersionScheme
// code-point if possible, otherwise as string
func (vs *VersionScheme) UnmarshalXMLAttr(attr xml.Attr) error {
	return xmlAttrToCode(attr, stringToVersionScheme, &vs.val)
}

func (vs *VersionScheme) SetCode(v uint64) error {
	if _, ok := versionSchemeToString[v]; ok {
		vs.val = v
		return nil
	}

	return fmt.Errorf("unknown version scheme %d", v)
}
