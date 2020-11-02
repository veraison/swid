package swid

// Payload models a payload-entry
type Payload struct {
	PayloadExtension
	GlobalAttributes
	ResourceCollection
}

// NewPayload instantiates a new empty Payload object
func NewPayload() *Payload {
	return &Payload{}
}

// AddDirectory adds the supplied Directory to the ResourceCollection of the
// Payload receiver
func (p *Payload) AddDirectory(d Directory) error {
	if p.Directories == nil {
		p.Directories = new(Directories)
	}

	*p.Directories = append(*p.Directories, d)

	return nil
}

// AddResource adds the supplied Directory to the ResourceCollection of the
// Payload receiver
func (p *Payload) AddResource(r Resource) error {
	if p.Resources == nil {
		p.Resources = new(Resources)
	}

	*p.Resources = append(*p.Resources, r)

	return nil
}
