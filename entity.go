// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// Entity models CoSWID's entity-entry map.
type Entity struct {
	// EntityExtension corresponds to the $$entity-extension CDDL socket which
	// is used to extend the entity-entry group model.
	EntityExtension

	GlobalAttributes

	// The textual name of the organizational entity claiming the roles
	// specified by the role item for the CoSWID tag.
	EntityName string `cbor:"31,keyasint" json:"entity-name" xml:"name,attr"`

	// The registration id value is intended to uniquely identify a naming
	// authority in a given scope (e.g. global, organization, vendor, customer,
	// administrative domain, etc.) for the referenced entity. The value of a
	// registration ID MUST be a RFC 3986 URI; it is not intended to be
	// dereferenced The scope SHOULD be the scope of an organization.
	// In a given scope, the registration id MUST be used
	// consistently for CoSWID tag production.
	RegID string `cbor:"32,keyasint,omitempty" json:"reg-id,omitempty" xml:"regid,attr"`

	// An integer or textual value representing the relationship(s) between the
	// entity, and this tag or the referenced software component. If an integer
	// value is used it MUST be an index value in the range -256 to 255. Integer
	// values in the range -256 to -1 are reserved for testing and use in closed
	// environments (see Section 6.2.2 of I-D.ietf-sacm-coswid). Integer values
	// in the range 0 to 255 correspond to registered entries in the IANA
	// "SWID/CoSWID Entity Role Value" registry (see Section 6.2.5 of
	// I-D.ietf-sacm-coswid).
	// The following additional requirements exist for the use of the "role"
	// item:
	// * An entity item MUST be provided with the role of "tag-creator" for
	//   every CoSWID tag. This indicates the organization that created the
	//   CoSWID tag.
	// * An entity item SHOULD be provided with the role of "software-creator"
	//   for every CoSWID tag, if this information is known to the tag creator.
	//   This indicates the organization that created the referenced software
	//   component.
	Roles Roles `cbor:"33,keyasint" json:"role" xml:"role,attr"`

	// The value of the Thumbprint field provides a hash value
	// (i.e. the thumbprint) of the signing entity's public key certificate.
	// This provides an indicator of which entity signed the CoSWID tag,
	// which will typically be the tag creator. See Section 2.9.1 of
	// I-D.ietf-sacm-coswid for more details on the use of the
	// hash-entry data structure.
	Thumbprint *HashEntry `cbor:"34,keyasint,omitempty" json:"thumbprint,omitempty" xml:"thumbprint,omitempty"`
}

// NewEntity instantiates a new Entity object initialized with the given
// entityName and roles
func NewEntity(entityName string, roles ...interface{}) (*Entity, error) {
	e := Entity{
		EntityName: entityName,
	}

	if err := e.SetRoles(roles...); err != nil {
		return nil, err
	}

	return &e, nil
}

// SetRoles sets the roles on the Entity receiver
func (e *Entity) SetRoles(roles ...interface{}) error {
	return e.Roles.Set(roles...)
}

// SetEntityName sets the name on the Entity receiver
func (e *Entity) SetEntityName(name string) error {
	e.EntityName = name
	return nil
}

// SetLang sets the language in the embedded GlobalAttributes of the Entity
// receiver
func (e *Entity) SetLang(lang string) error {
	e.Lang = lang
	return nil
}

// SetRegID sets the registration id on the Entity receiver
func (e *Entity) SetRegID(regID string) error {
	e.RegID = regID
	return nil
}

// SetThumbprint sets the signing certificate thumbprint on the Entity receiver
func (e *Entity) SetThumbprint(algID uint64, value []byte) error {
	e.Thumbprint = new(HashEntry)
	return e.Thumbprint.Set(algID, value)
}

// GetLang gets the language from the embedded GlobalAttributes of the Entity
// receiver
func (e Entity) GetLang() string {
	return e.Lang
}
