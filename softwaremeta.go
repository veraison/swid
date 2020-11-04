// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// SoftwareMeta models CoSWID's software-meta-entry map
type SoftwareMeta struct {
	SoftwareMetaExtension

	GlobalAttributes

	// A textual value that identifies how the software component has been
	// activated, which might relate to specific terms and conditions for its
	// use (e.g. Trial, Serialized, Licensed, Unlicensed, etc) and relate to an
	// entitlement. This attribute is typically used in supplemental tags as it
	// contains information that might be selected during a specific install.
	ActivationStatus string `cbor:"43,keyasint,omitempty" json:"activation-status,omitempty" xml:"activationStatus,attr,omitempty"`

	// A textual value that identfies which sales, licensing, or marketing
	// channel the software component has been targeted for (e.g. Volume,
	// Retail, OEM, Academic, etc). This attribute is typically used in
	// supplemental tags as it contains information that might be selected
	// during a specific install.
	ChannelType string `cbor:"44,keyasint,omitempty" json:"channel-type,omitempty" xml:"channelType,attr,omitempty"`

	// A textual value for the software component's informal or colloquial
	// version. Examples may include a year value, a major version number, or
	// similar value that are used to identify a group of specific software
	// component releases that are part of the same release/support cycle. This
	// version can be the same through multiple releases of a software
	// component, while the software-version specified in the concise-swid-tag
	// group is much more specific and will change for each software component
	// release. This version is intended to be used for string comparison only
	// and is not intended to be used to determine if a specific value is
	// earlier or later in a sequence.
	ColloquialVersion string `cbor:"45,keyasint,omitempty" json:"colloquial-version,omitempty" xml:"colloquialVersion,attr,omitempty"`

	// A textual value that provides a detailed description of the software
	// component. This value MAY be multiple paragraphs separated by CR LF
	// characters as described by [RFC5198].
	Description string `cbor:"46,keyasint,omitempty" json:"description,omitempty" xml:"description,attr,omitempty"`

	// A textual value indicating that the software component represents a
	// functional variation of the code base used to support multiple software
	// components. For example, this item can be used to differentiate
	// enterprise, standard, or professional variants of a software component.
	Edition string `cbor:"47,keyasint,omitempty" json:"edition,omitempty" xml:"edition,attr,omitempty"`

	// A boolean value that can be used to determine if accompanying proof of
	// entitlement is needed when a software license reconciliation process is
	// performed.
	EntitlementDataRequired *bool `cbor:"48,keyasint,omitempty" json:"entitlement-data-required,omitempty" xml:"entitlementDataRequired,attr,omitempty"`

	// A vendor-specific textual key that can be used to identify and establish
	// a relationship to an entitlement. Examples of an entitlement-key might
	// include a serial number, product key, or license key. For values that
	// relate to a given software component install (i.e., license key), a
	// supplemental tag will typically contain this information. In other cases,
	// where a general-purpose key can be provided that applies to all possible
	// installs of the software component on different endpoints, a primary tag
	// will typically contain this information.
	EntitlementKey string `cbor:"49,keyasint,omitempty" json:"entitlement-key,omitempty" xml:"entitlementKey,attr,omitempty"`

	// The name (or tag-id) of the software component that created the CoSWID
	// tag. If the generating software component has a SWID or CoSWID tag, then
	// the tag-id for the generating software component SHOULD be provided.
	Generator string `cbor:"50,keyasint,omitempty" json:"generator,omitempty" xml:"generator,attr,omitempty"`

	// A globally unique identifier used to identify a set of software
	// components that are related. Software components sharing the same
	// persistent-id can be different versions. This item can be used to relate
	// software components, released at different points in time or through
	// different release channels, that may not be able to be related through
	// use of the link item.
	PersistentID string `cbor:"51,keyasint,omitempty" json:"persistent-id,omitempty" xml:"persistentId,attr,omitempty"`

	// A basic name for the software component that can be common across
	// multiple tagged software components (e.g., Apache HTTPD).
	Product string `cbor:"52,keyasint,omitempty" json:"product,omitempty" xml:"product,attr,omitempty"`

	// A textual value indicating the software components overall product
	// family. This should be used when multiple related software components
	// form a larger capability that is installed on multiple different
	// endpoints. For example, some software families may consist of server,
	// client, and shared service components that are part of a larger
	// capability. Email systems, enterprise applications, backup services, web
	// conferencing, and similar capabilities are examples of families. Use of
	// this item is not intended to represent groups of software that are
	// bundled or installed together. The persistent-id or link items SHOULD be
	// used to relate bundled software components.
	ProductFamily string `cbor:"53,keyasint,omitempty" json:"product-family,omitempty" xml:"productFamily,attr,omitempty"`

	// A string value indicating an informal or colloquial release version of
	// the software. This value can provide a different version value as
	// compared to the software-version specified in the concise-swid-tag group.
	// This is useful when one or more releases need to have an informal version
	// label that differs from the specific exact version value specified by
	// software-version. Examples can include SP1, RC1, Beta, etc.
	Revision string `cbor:"54,keyasint,omitempty" json:"revision,omitempty" xml:"revision,attr,omitempty"`

	// A short description of the software component. This MUST be a single
	// sentence suitable for display in a user interface.
	Summary string `cbor:"55,keyasint,omitempty" json:"summary,omitempty" xml:"summary,attr,omitempty"`

	// An 8 digit UNSPSC classification code for the software component. For
	// more information see https://www.unspsc.org/
	UnspscCode string `cbor:"56,keyasint,omitempty" json:"unspsc-code,omitempty" xml:"unspscCode,attr,omitempty"`

	// The version of UNSPSC used to define the unspsc-code value.
	UnspscVersion string `cbor:"57,keyasint,omitempty" json:"unspsc-version,omitempty" xml:"unspscVersion,attr,omitempty"`
}
