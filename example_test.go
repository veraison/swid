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
	// <SoftwareIdentity xmlns="http://standards.iso.org/iso/19770/-2/2015/schema.xsd" tagId="com.acme.rrd2013-ce-sp1-v4-1-5-0" name="ACME Roadrunner Detector 2013 Coyote Edition SP1" version="4.1.5"><Meta activationStatus="trial" colloquialVersion="2013" edition="coyote" product="Roadrunner Detector" revision="sp1"></Meta><Entity name="The ACME Corporation" regid="acme.com" role="tagCreator softwareCreator"></Entity><Entity name="Coyote Services, Inc." regid="mycoyote.com" role="distributor"></Entity><Link href="www.gnu.org/licenses/gpl.txt" rel="license"></Link><Payload><Directory name="rrdetector" root="%programdata%"><File name="rrdetector.exe" size="532712" hash="sha-256;oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2o="></File></Directory><File name="test.exe" size="532712" hash="sha-256;oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2o="></File></Payload></SoftwareIdentity>
}

func Example_tcgRimExtension() {
	tag, _ := NewTag("94f6b457-9ac9-4d35-9b3f-78804173b65as", "ACME IoTCore", "1.0.0")

	entity, _ := NewEntity("ACME Ltd", RoleTagCreator, RoleSoftwareCreator)
	_ = entity.SetRegID("acme.example")
	_ = tag.AddEntity(*entity)

	softwareMeta := SoftwareMeta{
		ColloquialVersion: "Firmware_2019",
		Edition:           "IoT",
		Product:           "ProductA",
		Revision:          "r2",
	}
	_ = tag.AddSoftwareMeta(softwareMeta)

	mID, fID := uint64(201234), uint64(213022)
	uriLocal := AnyURI("/boot/tcg/manifest/swidtag")
	fwVendor := "BIOSVendorA"
	tcgRimReferenceMeasurementEntry := &TcgRimReferenceMeasurementEntry{
		PlatformConfigurationURILocal: &uriLocal,
		BindingSpecName:               "IoT RIM",
		BindingSpecVersion:            "1.2",
		PlatformManufacturerID:        &mID,
		PlatformManufacturerName:      "ACME",
		PlatformModelName:             "ProductA",
		FirmwareManufacturerName:      &fwVendor,
		FirmwareManufacturerID:        &fID,
		RIMLinkHash: []byte{
			0x88, 0xf2, 0x1d, 0x8e, 0x44, 0xd4, 0x27, 0x11, 0x49, 0x29, 0x74,
			0x04, 0xdf, 0x91, 0xca, 0xf2, 0x07, 0x13, 0x0b, 0xfa, 0x11, 0x65,
			0x82, 0x40, 0x8a, 0xbd, 0x04, 0xed, 0xe6, 0xdb, 0x7f, 0x51,
		},
	}
	tag.TcgRimReferenceMeasurementEntry = tcgRimReferenceMeasurementEntry

	fSz1, fSz2 := int64(25400), int64(1024)
	dir := Directory{
		FileSystemItem: FileSystemItem{
			Location: "/boot/iot",
			FsName:   "iotBase",
		},
		PathElements: &PathElements{
			Files: &Files{
				File{
					FileSystemItem: FileSystemItem{
						FsName: "ACME-iotBase.bin",
					},
					FileVersion: "01.00",
					Size:        &fSz1,
					Hash: &HashEntry{
						HashAlgID: Sha256,
						HashValue: []byte{
							0xa3, 0x14, 0xfc, 0x2d, 0xc6, 0x63, 0xae, 0x7a,
							0x6b, 0x6b, 0xc6, 0x78, 0x75, 0x94, 0x05, 0x73,
							0x96, 0xe6, 0xb3, 0xf5, 0x69, 0xcd, 0x50, 0xfd,
							0x5d, 0xdb, 0x4d, 0x1b, 0xba, 0xfd, 0x2b, 0x6a,
						},
					},
				},
				File{
					FileSystemItem: FileSystemItem{
						FsName: "iotExec.bin",
					},
					FileVersion: "01.00",
					Size:        &fSz2,
					Hash: &HashEntry{
						HashAlgID: Sha256,
						HashValue: []byte{
							0x53, 0x2e, 0xaa, 0xbd, 0x95, 0x74, 0x88, 0x0d,
							0xbf, 0x76, 0xb9, 0xb8, 0xcc, 0x00, 0x83, 0x2c,
							0x20, 0xa6, 0xec, 0x11, 0x3d, 0x68, 0x22, 0x99,
							0x55, 0x0d, 0x7a, 0x6e, 0x0f, 0x34, 0x5e, 0x25,
						},
					},
				},
			},
		},
	}
	payload := NewPayload()
	_ = payload.AddDirectory(dir)
	tag.Payload = payload

	jdata, _ := tag.ToJSON()
	fmt.Printf("%s\n", string(jdata))

	cdata, _ := tag.ToCBOR()
	fmt.Printf("%x\n", cdata)

	// Output:
	// {"tag-id":"94f6b457-9ac9-4d35-9b3f-78804173b65as","tag-version":0,"software-name":"ACME IoTCore","software-version":"1.0.0","software-meta":[{"colloquial-version":"Firmware_2019","edition":"IoT","product":"ProductA","revision":"r2"}],"entity":[{"entity-name":"ACME Ltd","reg-id":"acme.example","role":["tagCreator","softwareCreator"]}],"payload":{"directory":[{"location":"/boot/iot","fs-name":"iotBase","path-elements":{"file":[{"fs-name":"ACME-iotBase.bin","size":25400,"file-version":"01.00","hash":"sha-256:oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2o="},{"fs-name":"iotExec.bin","size":1024,"file-version":"01.00","hash":"sha-256:Uy6qvZV0iA2/drm4zACDLCCm7BE9aCKZVQ16bg80XiU="}]}}]},"tcg-rim:reference-measurement-entry":{"platform-configuration-uri-local":"/boot/tcg/manifest/swidtag","binding-spec-name":"IoT RIM","binding-spec-version":"1.2","platform-manufacturer-id":201234,"platform-manufacturer-name":"ACME","platform-model-name":"ProductA","firmware-manufacturer-id":213022,"firmware-manufacturer-name":"BIOSVendorA","rim-link-hash":"iPIdjkTUJxFJKXQE35HK8gcTC/oRZYJAir0E7ebbf1E="}}
	// a800782539346636623435372d396163392d346433352d396233662d373838303431373362363561730c00016c41434d4520496f54436f72650d65312e302e3005a4182d6d4669726d776172655f32303139182f63496f5418346850726f6475637441183662723202a3181f6841434d45204c746418206c61636d652e6578616d706c65182182010206a110a317692f626f6f742f696f74181867696f7442617365181aa11182a418187041434d452d696f74426173652e62696e14196338156530312e30300782015820a314fc2dc663ae7a6b6bc6787594057396e6b3f569cd50fd5ddb4d1bbafd2b6aa418186b696f74457865632e62696e14190400156530312e30300782015820532eaabd9574880dbf76b9b8cc00832c20a6ec113d682299550d7a6e0f345e25183aa9183d781a2f626f6f742f7463672f6d616e69666573742f73776964746167183e67496f542052494d183f63312e3218401a0003121218416441434d4518426850726f647563744118441a0003401e18456b42494f5356656e646f72411848582088f21d8e44d4271149297404df91caf207130bfa116582408abd04ede6db7f51
}
