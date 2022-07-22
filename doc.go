// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

/*
Package swid provides an API for creating and interacting with Software
Identification (SWID) Tags as defined by ISO/IEC 19770-2:2015 as well as by
their "concise" counterpart (CoSWID) defined by draft-ietf-sacm-coswid.

The library aims at using the most space-efficient encoding when using CBOR
and the most expressive one when using XML and JSON, preferring to serialize
strings rather tham of their equivalent code-points. When decoding, the most
space efficient representation is used. In dealing with unknown code-points,
we follow the Postel principle: refusing to encode unknown protocol entities,
while accepting unknown values - provided they fit the underlying type
system.

Creating Tags

A tag can be created with a call to NewTag() specifying a tag ID,
the name of the software being described and its version:

	tag, err := NewTag(
		"com.acme.rrd2013-ce-sp1-v4-1-5-0",
		"ACME Roadrunner Detector 2013 Coyote Edition SP1",
		"4.1.5",
	)

This will generate a Tag with a minimal structure. You can then use the API
to add additional information and meta data to the tag.

You will need to add one or more "entity" entries, representing the
organization(s) responsible for the information contained in the tag.
All entities have an associated "role" and a recommended "registration id":

	entity, err := NewEntity(
		"The ACME Corporation",
		RoleTagCreator, RoleSoftwareCreator,
	)
	err = entity.SetRegID("acme.com")

The newly created entity can be attached to the parent tag using the
AddEntity method:

	err = tag.AddEntity(*entity)

Next any number of files, directories as well as other kinds of resources can
be collected, e.g.:

	fileSize := int64(532712)
	fileHash, _ := hex.DecodeString("a314fc2dc663ae7a6b6bc6787594057396e6b3f569cd50fd5ddb4d1bbafd2b6a")

	dir := Directory{
		FileSystemItem: FileSystemItem{
			Root:   "%programdata%",
			FsName: "rrdetector",
		},
		PathElements: PathElements{
			Files: &Files{
				File{
					FileSystemItem: FileSystemItem{
						FsName: "rrdetector.exe",
					},
					Size: &fileSize,
					Hash: &HashEntry{
						HashAlgID: 1,
						HashValue: fileHash,
					},
				},
			},
		},
	}

And subsequently added to the tag's "payload":

	payload := NewPayload()
	if err := payload.AddDirectory(dir); err != nil { ... }
	if err := tag.AddPayload(*payload); err != nil { ... }

Note that the same data structures could be added to an "evidence" instead,
were the tag describing a "live" system rather than a software package.

Once the tag is complete, it can be serialized using one of the CBOR, XML or
JSON marshalers:

	data, err := tag.ToXML() // or tag.ToCBOR(), or tag.ToJSON()


Consuming Tags

A tag can be de-serialized using one of the "From" interfaces. For example,
to decode a CoSWID tag from a memory buffer:

	var tag SoftwareIdentity

	data := []byte{ 0xa6, 0x00, 0x78, 0x21, 0x65, 0x78, ... }

	if err := tag.FromCBOR(data); err != nil { ... }

Similarly, for a SWID tag:

	var tag SoftwareIdentity

	data := []btye(`<SoftwareIdentity xmlns="...`)

	if err := tag.FromXML(data); err != nil { ... }

Or a CoSWID/JSON tag:

	var tag SoftwareIdentity

	data := []btye(`{"tag-id":"example.acme.roadrunner-sw-bl-v1-0-0", ...`)

	if err := tag.FromJSON(data); err != nil { ... }

Note that all nested fields are accessible from outside the swid package, so
(for now) no special getters are provided by the API.

Enjoy!
*/
package swid
