// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

// FileSystemItem models CoSWID filesystem-item
type FileSystemItem struct {
	GlobalAttributes

	// A boolean value indicating if a file or directory is significant or
	// required for the software component to execute or function properly.
	// These are files or directories that can be used to affirmatively
	// determine if the software component is installed on an endpoint.
	Key *bool `cbor:"22,keyasint,omitempty" json:"key,omitempty" xml:"key,attr,omitempty"`

	// The filesystem path where a file is expected to be located when installed
	// or copied. The location MUST be either relative to the location of the
	// parent directory item (preferred) or relative to the location of the
	// CoSWID tag if no parent is defined. The location MUST NOT include a
	// file's name, which is provided by the fs-name item.
	Location string `cbor:"23,keyasint,omitempty" json:"location,omitempty" xml:"location,attr,omitempty"`

	// The name of the directory or file without any path information.
	FsName string `cbor:"24,keyasint" json:"fs-name" xml:"name,attr,omitempty"`

	// A filesystem-specific name for the root of the filesystem. The location
	// item is considered relative to this location if specified. If not
	// provided, the value provided by the location item is expected to be
	// relative to its parent or the location of the CoSWID tag if no parent is
	// provided.
	Root string `cbor:"25,keyasint,omitempty" json:"root,omitempty" xml:"root,attr,omitempty"`
}
