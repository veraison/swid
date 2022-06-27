// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// Directory models CoSWID directory-entry
type Directory struct {
	DirectoryExtension
	FileSystemItem
	*PathElements `cbor:"26,keyasint" json:"path-elements"`
}

// PathElements models CoSWID path-elements-group
type PathElements struct {
	Directories *Directories `cbor:"16,keyasint,omitempty" json:"directory,omitempty" xml:"Directory,omitempty"`
	Files       *Files       `cbor:"17,keyasint,omitempty" json:"file,omitempty" xml:"File,omitempty"`
}
