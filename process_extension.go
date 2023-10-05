// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// ProcessExtension is a placeholder for $$process-extension
type ProcessExtension struct {
	// TCG RIM extension
	TCGRIMHashEntry *HashEntry `cbor:"7,keyasint,omitempty" json:"tgc-rim:hash-entry,omitempty" xml:"-"`
}
