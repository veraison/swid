// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// DirectoryExtension is a placeholder for $$directory-extension
type DirectoryExtension struct {
	// TCG RIM extension
	TCGRIMHashEntry *HashEntry `cbor:"7,keyasint,omitempty" json:"tgc-rim:hash-entry,omitempty" xml:"-"`
}
