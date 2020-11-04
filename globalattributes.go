// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// GlobalAttributes models CoSWID global-attributes
type GlobalAttributes struct {
	Lang string `cbor:"15,keyasint,omitempty" json:"lang,omitempty" xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// no any-attribute's registered
}
