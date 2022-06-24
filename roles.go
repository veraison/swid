// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

// Roles models the role type
type Roles struct {
	val []interface{}
}

/*
   entity-entry = {
     ...
     role => $role / [ 2* $role ],
     ...
   }

   $role /= tag-creator
   $role /= software-creator
   $role /= aggregator
   $role /= distributor
   $role /= licensor
   $role /= maintainer
   $role /= int / text
   tag-creator=1
   software-creator=2
   aggregator=3
   distributor=4
   licensor=5
   maintainer=6
*/

// Role constants
const (
	RoleTagCreator = int64(iota + 1)
	RoleSoftwareCreator
	RoleAggregator
	RoleDistributor
	RoleLicensor
	RoleMaintainer
	RoleUnknown = ^int64(0)
)

var (
	roleToString = map[int64]string{
		RoleTagCreator:      "tagCreator",
		RoleSoftwareCreator: "softwareCreator",
		RoleAggregator:      "aggregator",
		RoleDistributor:     "distributor",
		RoleLicensor:        "licensor",
		RoleMaintainer:      "maintainer",
	}

	stringToRole = map[string]int64{
		"tagCreator":      RoleTagCreator,
		"softwareCreator": RoleSoftwareCreator,
		"aggregator":      RoleAggregator,
		"distributor":     RoleDistributor,
		"licensor":        RoleLicensor,
		"maintainer":      RoleMaintainer,
	}
)

func (r Roles) stringer(skipUnknown bool) string {
	v := r.val // make a copy that we can clobber

	s := []string{}

	codeName := "role"
	if skipUnknown {
		codeName = ""
	}

	for i := range v {
		if err := stringifyCode(&v[i], roleToString, codeName); err != nil {
			continue
		}
		// after a successful stringifyCode the type assertion on v[i] is safe
		s = append(s, v[i].(string))
	}

	return strings.Join(s, " ")
}

// String returns the value of the Roles receiver as a string
// Unknown roles are returned as "role(" + code-point + ")"
func (r Roles) String() string {
	skipUnknownRoles := false

	return r.stringer(skipUnknownRoles)
}

// Check returns nil if all the roles stored in the Roles receiver are of type
// string or code-point (i.e., uint).
func (r Roles) Check() error {
	// $role /= uint / text
	for _, part := range r.val {
		if err := isStringOrCode(part, "role"); err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON provides the custom JSON marshaler for the Roles type
// that takes care of the $role / [ 2* $role ] variants
func (r Roles) MarshalJSON() ([]byte, error) {
	v := r.val // make a copy that we can clobber

	// handle singleton
	if len(v) == 1 {
		if err := stringifyCode(&v[0], roleToString, ""); err != nil {
			return nil, err
		}
		return json.Marshal(&v[0])
	}

	// handle array
	for i := range v {
		if err := stringifyCode(&v[i], roleToString, ""); err != nil {
			return nil, err
		}
	}

	return json.Marshal(v)
}

func (r *Roles) postprocess(a []interface{}) error {
	// 'a' is mostly good already, modulo some type checking and mapping
	// of strings into codepoints
	for i := range a {
		if err := codifyString(&a[i], stringToRole); err != nil {
			return err
		}
	}

	// at this point we know it is safe to use 'a'
	r.val = a

	return nil
}

// UnmarshalJSON provides the custom JSON unmarshaler for the Roles type that
// takes care of the $role / [ 2* $role ] variants. (Note that we accept [ 1*
// $role ] as valid too.
func (r *Roles) UnmarshalJSON(data []byte) error {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	return r.postUnmarshal(v)
}

// MarshalCBOR provides the custom CBOR marshaler for the Roles type
// that takes care of the $role / [ 2* $role ] variants
func (r Roles) MarshalCBOR() ([]byte, error) {
	v := r.val // make a copy that we can clobber

	if len(v) == 1 {
		if err := codifyString(&v[0], stringToRole); err != nil {
			return nil, err
		}
		return em.Marshal(&v[0])
	}

	for i := range v {
		if err := codifyString(&v[i], stringToRole); err != nil {
			return nil, err
		}
	}

	return em.Marshal(v)
}

func (r *Roles) postUnmarshal(v interface{}) error {
	// Handle '$role / [ 2* $role ]'.
	// (Note that we accept '[ 1* $role ]' too as valid.)
	switch t := v.(type) {
	case []interface{}:
		return r.postprocess(t)
	default:
		// if it did't decode to an array, there's a chance it's
		// a singleton: wrap it into an array and try to run the
		// post-processor on it.
		return r.postprocess([]interface{}{t})
	}
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for the Roles type that
// takes care of the $role / [ 2* $role ] variants. (Note that we accept [ 1*
// $role ] as valid too.
func (r *Roles) UnmarshalCBOR(data []byte) error {
	var v interface{}

	if err := dm.Unmarshal(data, &v); err != nil {
		return err
	}

	return r.postUnmarshal(v)
}

// Set the supplied roles in the Roles receiver. The set operation fails if any
// of the supplied roles does not validate (see Check)
func (r *Roles) Set(roles ...interface{}) error {
	var v Roles

	v.val = append(v.val, roles...)

	if err := v.Check(); err != nil {
		return err
	}

	// steal it
	r.val = v.val

	return nil
}

// MarshalXMLAttr provides a custom XML attribute marshaler for the Roles type
func (r Roles) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	skipUnknownRoles := true

	return xml.Attr{Name: name, Value: r.stringer(skipUnknownRoles)}, nil
}

// UnmarshalXMLAttr provides a custom XML attribute unmarshaler for the Roles
// type
func (r *Roles) UnmarshalXMLAttr(attr xml.Attr) error {
	var v []interface{}
	for _, role := range strings.Fields(attr.Value) {
		v = append(v, role)
	}
	return r.Set(v...)
}
