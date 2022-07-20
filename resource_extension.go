// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// ResourceExtension is a placeholder for $$resource-extension
type ResourceExtension struct {
	// TCG RIM extension
	TCGRIMHashEntry *HashEntry `cbor:"7,keyasint,omitempty" json:"tgc-rim:hash-entry,omitempty" xml:"-"`
}
