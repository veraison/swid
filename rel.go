package swid

import (
	"encoding/xml"
)

// Rel models the rel type.
type Rel struct {
	val interface{}
}

/*
   $rel /= ancestor
   $rel /= component
   $rel /= feature
   $rel /= installationmedia
   $rel /= packageinstaller
   $rel /= parent
   $rel /= patches
   $rel /= requires
   $rel /= see-also
   $rel /= supersedes
   $rel /= supplemental
   $rel /= uint / text
   ancestor=1
   component=2
   feature=3
   installationmedia=4
   packageinstaller=5
   parent=6
   patches=7
   requires=8
   see-also=9
   supersedes=10
   supplemental=11
*/

// Rel constants
const (
	RelAncestor = uint64(iota + 1)
	RelComponent
	RelFeature
	RelInstallationMedia
	RelPackageInstaller
	RelParent
	RelPatches
	RelRequires
	RelSeeAlso
	RelSupersedes
	RelSupplemental
	RelUnknown = ^uint64(0)
)

var (
	relToString = map[uint64]string{
		RelAncestor:          "ancestor",
		RelComponent:         "component",
		RelFeature:           "feature",
		RelInstallationMedia: "installation media",
		RelPackageInstaller:  "package installer",
		RelParent:            "parent",
		RelPatches:           "patches",
		RelRequires:          "requires",
		RelSeeAlso:           "see also",
		RelSupersedes:        "supersedes",
		RelSupplemental:      "supplemental",
	}

	stringToRel = map[string]uint64{
		"ancestor":           RelAncestor,
		"component":          RelComponent,
		"feature":            RelFeature,
		"installation media": RelInstallationMedia,
		"package installer":  RelPackageInstaller,
		"parent":             RelParent,
		"patches":            RelPatches,
		"requires":           RelRequires,
		"see also":           RelSeeAlso,
		"supersedes":         RelSupersedes,
		"supplemental":       RelSupplemental,
	}
)

// String returns the value of the Rel receiver as a string
func (r Rel) String() string {
	return codeStringer(r.val, relToString, "rel")
}

// Check returns nil if the Rel receiver is of type string or code-point (i.e.,
// uint)
func (r Rel) Check() error {
	return isStringOrCode(r.val, "rel")
}

// MarshalCBOR encodes the Rel receiver as code-point if possible, otherwise as
// string
func (r Rel) MarshalCBOR() ([]byte, error) {
	return codeToCBOR(r.val, stringToRel)
}

// UnmarshalCBOR decodes the supplied data into a Rel code-point if possible,
// otherwise as string
func (r *Rel) UnmarshalCBOR(data []byte) error {
	return cborToCode(data, stringToRel, &r.val)
}

// MarshalJSON encodes the Rel receiver as string
func (r Rel) MarshalJSON() ([]byte, error) {
	return codeToJSON(r.val, relToString)
}

// UnmarshalJSON decodes the supplied data into a Rel code-point if possible,
// otherwise as string
func (r *Rel) UnmarshalJSON(data []byte) error {
	return jsonToCode(data, stringToRel, &r.val)
}

// MarshalXMLAttr encodes the Rel receiver as XML attribute
func (r Rel) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return codeToXMLAttr(name, r.val, relToString)
}

// UnmarshalXMLAttr decodes the supplied XML attribute into a Rel code-point if
// possible, otherwise as string
func (r *Rel) UnmarshalXMLAttr(attr xml.Attr) error {
	return xmlAttrToCode(attr, stringToRel, &r.val)
}
