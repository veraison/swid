// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// File models CoSWID file-entry
type File struct {
	GlobalAttributes

	FileExtension

	FileSystemItem

	// The file's size in bytes.
	Size *int64 `cbor:"20,keyasint,omitempty" json:"size,omitempty" xml:"size,attr,omitempty"`

	// The file's version as reported by querying information on the file from
	// the operating system.
	FileVersion string `cbor:"21,keyasint,omitempty" json:"file-version,omitempty" xml:"version,attr,omitempty"`

	// Files are expected to include a hash value that provides different level
	// of confidence that if the filename, file size and file hash code all
	// match, the the file has not been modified in any fashion.
	Hash *HashEntry `cbor:"7,keyasint,omitempty" json:"hash,omitempty" xml:"hash,attr,omitempty"`
}
