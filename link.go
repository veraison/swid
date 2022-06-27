// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// Link models CoSWID's link-entry map
type Link struct {
	LinkExtension

	GlobalAttributes

	// To be used with rel="installation-media", this item's value provides the
	// path to the installer executable or script that can be run to launch the
	// referenced installation. Links with the same artifact name MUST be
	// considered mirrors of each other, allowing the installation media to be
	// acquired from any of the described sources.
	Artifact string `cbor:"37,keyasint,omitempty" json:"artifact,omitempty" xml:"artifact,attr,omitempty"`

	// A URI for the referenced resource. The "href" item's value can be, but is
	// not limited to, the following (which is a slightly modified excerpt from
	// [SWID]):
	// * If no URI scheme is provided, then the URI is to be interpreted as
	//   being relative to the base URI of the CoSWID tag, i.e., the URI
	//   under which the CoSWID tag was provided. For example,
	//   "folder/supplemental.coswid"
	// * a physical resource location with any acceptable URI scheme (e.g., file://
	//   http:// https:// ftp://)
	// * a URI with "swid:" as the scheme refers to another SWID or CoSWID by
	//   the referenced tag's tag-id. This URI needs to be resolved in the context
	//   of the endpoint by software that can lookup other SWID or CoSWID tags.
	//   For example, "swid:2df9de35-0aff-4a86-ace6-f7dddd1ade4c" references the
	//   tag with the tag-id value "2df9de35-0aff-4a86-ace6-f7dddd1ade4c".
	// * a URI with "swidpath:" as the scheme, which refers to another CoSIWD
	//   via an XPATH query [W3C.REC-xpath20-20101214] that matches items in that
	//   tag (Section 5.2). This URI would need to be resolved in the context of
	//   the system entity via software components that can lookup other CoSWID
	//   tags and select the appropriate tag based on an XPATH query
	//   [W3C.REC-xpath20-20101214].
	//   Examples include:
	//   * swidpath://SoftwareIdentity[Entity/@regid='http://contoso.com'] would
	//    retrieve all SWID or CoSWID tags that include an entity where the regid
	//    is "Contoso"
	//   * swidpath://SoftwareIdentity[Meta/@persistentId='b0c55172-38e9-4e36-be86-92206ad8eddb']
	//   would match all SWID or CoSWID tags with the persistent-id value
	//   "b0c55172-38e9-4e36-be86-92206ad8eddb"
	Href string `cbor:"38,keyasint" json:"href" xml:"href,attr"`

	// A hint to the consumer of the link to what target platform the link is
	// applicable to. This item represents a query as defined by the W3C Media
	// Queries Recommendation (see [W3C.REC-css3-mediaqueries-20120619]).
	Media string `cbor:"10,keyasint,omitempty" json:"media,omitempty" xml:"media,attr,omitempty"`

	// An integer or textual value (integer label with text escape,
	// see Section 2, for the "Software Tag Link Ownership Values"
	// registry Section 4.3) used when the "href" item references another
	// software component to indicate the degree of ownership between the
	// software component referenced by the CoSWID tag and the software
	// component referenced by the link. If an integer value is used it MUST
	// be an index value in the range -256 to 255. Integer values in the range
	// -256 to -1 are reserved for testing and use in closed environments
	// (see Section 6.2.2). Integer values in the range 0 to 255 correspond
	// to registered entries in the "Software Tag Link Ownership Values" registry.
	Ownership *Ownership `cbor:"39,keyasint,omitempty" json:"ownership,omitempty" xml:"ownership,attr,omitempty"`

	// An integer or textual value that (integer label with text escape,
	// see Section 2, for the "Software Tag Link Link Relationship Values"
	// registry Section 4.3) identifies the relationship between this CoSWID
	// and the target resource identified by the "href" item. If an integer
	// value is used it MUST be an index value in the range -256 to 65535.
	// Integer values in the range -256 to -1 are reserved for testing and
	// use in closed environments (see Section 6.2.2). Integer values in the
	// range 0 to 65535 correspond to registered entries in the IANA
	// "Software Tag Link Relationship Values" registry (see Section 6.2.7.
	// If a string value is used it MUST be either a private use name as defined
	// in Section 6.2.2 or a "Relation Name" from the IANA "Link Relation Types"
	// registry: https://www.iana.org/assignments/link-relations/link-relations.xhtml
	// as defined by [RFC8288]. When a string value defined in the IANA
	// "Software Tag Link Relationship Values" registry matches a Relation
	// Name defined in the IANA "Link Relation Types" registry, the index
	// value in the IANA "Software Tag Link Relationship Values" registry
	// MUST be used instead, as this relationship has a specialized meaning
	// in the context of a CoSWID tag. String values correspond to registered
	// entries in the "Software Tag Link Relationship Values" registry.
	Rel Rel `cbor:"40,keyasint" json:"rel" xml:"rel,attr"`

	// A link can point to arbitrary resources on the endpoint, local network,
	// or Internet using the href item. Use of this item supplies the resource
	// consumer with a hint of what type of resource to expect. (This is a hint:
	// There is no obligation for the server hosting the target of the URI to
	// use the indicated media type when the URI is dereferenced.) Media types
	// are identified by referencing a "Name" from the IANA "Media Types"
	// registry: http://www.iana.org/assignments/media-types/media-types.xhtml.
	// This item maps to '/SoftwareIdentity/Link/@type' in [SWID].
	MediaType string `cbor:"41,keyasint,omitempty" json:"media-type,omitempty" xml:"type,attr,omitempty"`

	// An integer or textual value (integer label with text escape, see
	// Section 2, for the "Software Tag Link Link Relationship Values"
	// registry Section 4.3) used to determine if the referenced software
	// component has to be installed before installing the software component
	// identified by the COSWID tag. If an integer value is used it MUST be an
	// index value in the range -256 to 255. Integer values in the range -256
	// to -1 are reserved for testing and use in closed environments
	// (see Section 6.2.2). Integer values in the range 0 to 255 correspond to
	// registered entries in the IANA "Link Use Values" registry (see
	// Section 6.2.8. If a string value is used it MUST be a private use name
	// as defined in Section 6.2.2. String values correspond to registered
	// entries in the "Software Tag Link Use Values" registry
	Use *Use `cbor:"42,keyasint,omitempty" json:"use,omitempty" xml:"use,attr,omitempty"`
}

// NewLink instantiates a new Link object initialized with the supplied href and
// link relation
func NewLink(href string, rel Rel) (*Link, error) {
	l := Link{
		Href: href,
	}

	if err := l.SetRel(rel); err != nil {
		return nil, err
	}

	return &l, nil
}

// SetRel assigns the supplied link relation to the Link receiver
func (l *Link) SetRel(rel Rel) error {
	if err := rel.Check(); err != nil {
		return err
	}

	l.Rel = rel

	return nil
}

// GetUseAsString returns the use attribute of a Link object as a string
func (l Link) GetUseAsString() string {
	return l.Use.String()
}

// GetOwnershipAsString returns the ownership attribute of a Link object as a
// string
func (l Link) GetOwnershipAsString() string {
	return l.Ownership.String()
}

// GetRelAsString returns the relation of a Link object as a string
func (l Link) GetRelAsString() string {
	return l.Rel.String()
}
