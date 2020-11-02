package swid

// Resource models a resource-entry
type Resource struct {
	GlobalAttributes
	Type string `cbor:"29,keyasint" json:"type" xml:"type,attr"`
	ResourceExtension
}
