// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

type AnyURI string

type TcgRimReferenceMeasurementEntry struct {
	PayloadType                    *TcgRimPayloadType `cbor:"59,keyasint,omitempty" json:"payload-type,omitempty"`
	PlatformConfigurationURIGlobal *AnyURI            `cbor:"60,keyasint,omitempty" json:"platform-configuration-uri-global,omitempty"`
	PlatformConfigurationURILocal  *AnyURI            `cbor:"61,keyasint,omitempty" json:"platform-configuration-uri-local,omitempty"`
	BindingSpecName                string             `cbor:"62,keyasint" json:"binding-spec-name"`
	BindingSpecVersion             string             `cbor:"63,keyasint" json:"binding-spec-version"`
	PlatformManufacturerID         *uint64            `cbor:"64,keyasint,omitempty" json:"platform-manufacturer-id,omitempty"`
	PlatformManufacturerName       string             `cbor:"65,keyasint" json:"platform-manufacturer-name"`
	PlatformModelName              string             `cbor:"66,keyasint" json:"platform-model-name"`
	PlatformVersion                *uint64            `cbor:"67,keyasint,omitempty" json:"platform-version,omitempty"`
	FirmwareManufacturerID         *uint64            `cbor:"68,keyasint,omitempty" json:"firmware-manufacturer-id,omitempty"`
	FirmwareManufacturerName       *string            `cbor:"69,keyasint,omitempty" json:"firmware-manufacturer-name,omitempty"`
	FirmwareModelName              *string            `cbor:"70,keyasint,omitempty" json:"firmware-model-name,omitempty"`
	FirmwareVersion                *uint64            `cbor:"71,keyasint,omitempty" json:"firmware-version,omitempty"`
	RIMLinkHash                    []byte             `cbor:"72,keyasint" json:"rim-link-hash"`
}

type TcgRimPayloadType uint64

const (
	TcgRimPayloadTypeDirect = TcgRimPayloadType(iota)
	TcgRimPayloadTypeIndirect
	TcgRimPayloadTypeHybrid
)
