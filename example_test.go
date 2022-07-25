// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/hex"
	"fmt"
)

func Example_links() {
	// make new tag
	tag, _ := NewTag("example.acme.roadrunner-sw-v1-0-0", "Roadrunner software bundle", "1.0.0")

	// make entity and add it to the tag
	entity, _ := NewEntity("ACME Ltd", RoleTagCreator, RoleSoftwareCreator, RoleAggregator)
	_ = entity.SetRegID("acme.example")
	_ = tag.AddEntity(*entity)

	// make links and append them to tag
	link, _ := NewLink("example.acme.roadrunner-hw-v1-0-0", *NewRel("psa-rot-compound"))
	_ = tag.AddLink(*link)

	link, _ = NewLink("example.acme.roadrunner-sw-bl-v1-0-0", *NewRel(RelComponent))
	_ = tag.AddLink(*link)

	link, _ = NewLink("example.acme.roadrunner-sw-prot-v1-0-0", *NewRel(RelComponent))
	_ = tag.AddLink(*link)

	link, _ = NewLink("example.acme.roadrunner-sw-arot-v1-0-0", *NewRel(RelComponent))
	_ = tag.AddLink(*link)

	// encode tag to JSON
	data, _ := tag.ToJSON()
	fmt.Println(string(data))

	// encode tag to XML
	data, _ = tag.ToXML()
	fmt.Println(string(data))

	// Output:
	// {"tag-id":"example.acme.roadrunner-sw-v1-0-0","tag-version":0,"software-name":"Roadrunner software bundle","software-version":"1.0.0","entity":[{"entity-name":"ACME Ltd","reg-id":"acme.example","role":["tagCreator","softwareCreator","aggregator"]}],"link":[{"href":"example.acme.roadrunner-hw-v1-0-0","rel":"psa-rot-compound"},{"href":"example.acme.roadrunner-sw-bl-v1-0-0","rel":"component"},{"href":"example.acme.roadrunner-sw-prot-v1-0-0","rel":"component"},{"href":"example.acme.roadrunner-sw-arot-v1-0-0","rel":"component"}]}
	// <SoftwareIdentity xmlns="http://standards.iso.org/iso/19770/-2/2015/schema.xsd" tagId="example.acme.roadrunner-sw-v1-0-0" name="Roadrunner software bundle" version="1.0.0"><Entity name="ACME Ltd" regid="acme.example" role="tagCreator softwareCreator aggregator"></Entity><Link href="example.acme.roadrunner-hw-v1-0-0" rel="psa-rot-compound"></Link><Link href="example.acme.roadrunner-sw-bl-v1-0-0" rel="component"></Link><Link href="example.acme.roadrunner-sw-prot-v1-0-0" rel="component"></Link><Link href="example.acme.roadrunner-sw-arot-v1-0-0" rel="component"></Link></SoftwareIdentity>
}

func Example_completePrimaryTag() {
	tag, _ := NewTag(
		"com.acme.rrd2013-ce-sp1-v4-1-5-0",
		"ACME Roadrunner Detector 2013 Coyote Edition SP1",
		"4.1.5",
	)

	entity, _ := NewEntity("The ACME Corporation", RoleTagCreator, RoleSoftwareCreator)
	_ = entity.SetRegID("acme.com")
	_ = tag.AddEntity(*entity)

	entity, _ = NewEntity("Coyote Services, Inc.", RoleDistributor)
	_ = entity.SetRegID("mycoyote.com")
	_ = tag.AddEntity(*entity)

	link, _ := NewLink("www.gnu.org/licenses/gpl.txt", *NewRel("license"))
	_ = tag.AddLink(*link)

	meta := SoftwareMeta{
		ActivationStatus:  "trial",
		Product:           "Roadrunner Detector",
		ColloquialVersion: "2013",
		Edition:           "coyote",
		Revision:          "sp1",
	}
	_ = tag.AddSoftwareMeta(meta)

	fileSize := int64(532712)
	fileHash, _ := hex.DecodeString("a314fc2dc663ae7a6b6bc6787594057396e6b3f569cd50fd5ddb4d1bbafd2b6a")

	dir := Directory{
		FileSystemItem: FileSystemItem{
			Root:   "%programdata%",
			FsName: "rrdetector",
		},
		PathElements: &PathElements{
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

	file := File{
		FileSystemItem: FileSystemItem{
			FsName: "test.exe",
		},
		Size: &fileSize,
		Hash: &HashEntry{
			HashAlgID: 1,
			HashValue: fileHash,
		},
	}

	payload := NewPayload()
	_ = payload.AddDirectory(dir)
	_ = payload.AddFile(file)
	tag.Payload = payload

	// encode tag to XML
	data, _ := tag.ToXML()
	fmt.Println(string(data))

	// Output:
	// <SoftwareIdentity xmlns="http://standards.iso.org/iso/19770/-2/2015/schema.xsd" tagId="com.acme.rrd2013-ce-sp1-v4-1-5-0" name="ACME Roadrunner Detector 2013 Coyote Edition SP1" version="4.1.5"><Meta activationStatus="trial" colloquialVersion="2013" edition="coyote" product="Roadrunner Detector" revision="sp1"></Meta><Entity name="The ACME Corporation" regid="acme.com" role="tagCreator softwareCreator"></Entity><Entity name="Coyote Services, Inc." regid="mycoyote.com" role="distributor"></Entity><Link href="www.gnu.org/licenses/gpl.txt" rel="license"></Link><Payload><Directory name="rrdetector" root="%programdata%"><File name="rrdetector.exe" size="532712" hash="sha-256:oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2o="></File></Directory><File name="test.exe" size="532712" hash="sha-256:oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2o="></File></Payload></SoftwareIdentity>
}
