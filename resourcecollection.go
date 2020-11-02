package swid

// ResourceCollection models a resource-collection
type ResourceCollection struct {
	ResourceCollectionExtension

	PathElements

	// A process item allows details to be provided about the runtime behavior
	// of the software component, such as information that will appear in a
	// process listing on an endpoint
	Processes *Processes `cbor:"18,keyasint,omitempty" json:"process,omitempty" xml:"Process,omitempty"`

	// A resource item can be used to provide details about an artifact or
	// capability expected to be found on an endpoint or evidence collected
	// related to the software component. This can be used to represent concepts
	// not addressed directly by the directory, file, or process items. Examples
	// include: registry keys, bound ports, etc. The equivalent construct in
	// [SWID] is currently under specified. As a result, this item might be
	// further defined through extension in the future.
	Resources *Resources `cbor:"19,keyasint,omitempty" json:"resource,omitempty" xml:"Resource,omitempty"`
}
