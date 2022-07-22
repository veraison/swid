// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// Resource models a resource-entry
type Resource struct {
	ResourceExtension
	GlobalAttributes
	Type string `cbor:"29,keyasint" json:"type" xml:"type,attr"`
}
