// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// CoSWIDExtension is a placeholder for $$coswid-extension
type CoSWIDExtension struct {
	TcgRimReferenceMeasurementEntry *TcgRimReferenceMeasurementEntry `cbor:"58,keyasint,omitempty" json:"tcg-rim:reference-measurement-entry,omitempty" xml:"-"`
}
