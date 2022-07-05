// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

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

func (p *Payload) AddFile(f File) error {
	if p.Files == nil {
		p.Files = new(Files)
	}

	*p.Files = append(*p.Files, f)

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
