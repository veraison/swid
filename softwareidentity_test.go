// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testXML  = []byte(`<SoftwareIdentity tagId="f432dc99-2e06-434d-b9ad-2b22e35b6fa4" name="Roadrunner software bundle" version="1.0.0"><Entity name="ACME Ltd" regid="acme.example" role="tagCreator softwareCreator"></Entity><Link href="d84fb5e2-d198-49b4-9d65-3a82421bf180" rel="parent"></Link></SoftwareIdentity>`)
	testJSON = []byte(`{
		"tag-id": "f432dc99-2e06-434d-b9ad-2b22e35b6fa4",
		"tag-version": 0,
		"software-name": "Roadrunner software bundle",
		"software-version": "1.0.0",
		"entity": [
		  {
			"entity-name": "ACME Ltd",
			"reg-id": "acme.example",
			"role": [
			  "tagCreator",
			  "softwareCreator"
			]
		  }
		],
		"link": [
		  {
			"href": "d84fb5e2-d198-49b4-9d65-3a82421bf180",
			"rel": "parent"
		  }
		]
	}`)
	testCBOR = []byte{
		0xa6, 0x00, 0x50, 0xf4, 0x32, 0xdc, 0x99, 0x2e,
		0x06, 0x43, 0x4d, 0xb9, 0xad, 0x2b, 0x22, 0xe3,
		0x5b, 0x6f, 0xa4, 0x0c, 0x00, 0x01, 0x78, 0x1a,
		0x52, 0x6f, 0x61, 0x64, 0x72, 0x75, 0x6e, 0x6e,
		0x65, 0x72, 0x20, 0x73, 0x6f, 0x66, 0x74, 0x77,
		0x61, 0x72, 0x65, 0x20, 0x62, 0x75, 0x6e, 0x64,
		0x6c, 0x65, 0x0d, 0x65, 0x31, 0x2e, 0x30, 0x2e,
		0x30, 0x02, 0xa3, 0x18, 0x1f, 0x68, 0x41, 0x43,
		0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x18, 0x20,
		0x6c, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x65, 0x78,
		0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x21, 0x82,
		0x01, 0x02, 0x04, 0xa2, 0x18, 0x26, 0x78, 0x24,
		0x64, 0x38, 0x34, 0x66, 0x62, 0x35, 0x65, 0x32,
		0x2d, 0x64, 0x31, 0x39, 0x38, 0x2d, 0x34, 0x39,
		0x62, 0x34, 0x2d, 0x39, 0x64, 0x36, 0x35, 0x2d,
		0x33, 0x61, 0x38, 0x32, 0x34, 0x32, 0x31, 0x62,
		0x66, 0x31, 0x38, 0x30, 0x18, 0x28, 0x06,
	}
)

func makeACMEEntityWithRoles(t *testing.T, roles ...interface{}) Entity {
	e := Entity{
		EntityName: "ACME Ltd",
		RegID:      "acme.example",
	}

	require.Nil(t, e.SetRoles(roles...))

	return e
}

func TestTag_FromCBOR_ok(t *testing.T) {
	expected := SoftwareIdentity{
		TagID:           *NewTagID("f432dc99-2e06-434d-b9ad-2b22e35b6fa4"),
		SoftwareName:    "Roadrunner software bundle",
		SoftwareVersion: "1.0.0",
		Entities: Entities{
			makeACMEEntityWithRoles(t,
				RoleTagCreator,
				RoleSoftwareCreator,
			),
		},
		Links: &Links{
			Link{
				Href: "d84fb5e2-d198-49b4-9d65-3a82421bf180",
				Rel:  *NewRel(RelParent),
			},
		},
	}

	var tv SoftwareIdentity

	err := tv.FromCBOR(testCBOR)
	assert.NoError(t, err)
	assert.Equal(t, expected, tv)
}

func TestTag_FromJSON_ok(t *testing.T) {
	expected := SoftwareIdentity{
		TagID:           *NewTagID("f432dc99-2e06-434d-b9ad-2b22e35b6fa4"),
		SoftwareName:    "Roadrunner software bundle",
		SoftwareVersion: "1.0.0",
		Entities: Entities{
			makeACMEEntityWithRoles(t,
				RoleTagCreator,
				RoleSoftwareCreator,
			),
		},
		Links: &Links{
			Link{
				Href: "d84fb5e2-d198-49b4-9d65-3a82421bf180",
				Rel:  *NewRel(RelParent),
			},
		},
	}

	var tv SoftwareIdentity

	err := tv.FromJSON(testJSON)
	assert.NoError(t, err)
	assert.Equal(t, expected, tv)
}

func TestTag_FromXML_ok(t *testing.T) {
	expected := SoftwareIdentity{
		XMLName: xml.Name{
			Local: "SoftwareIdentity",
		},
		TagID:           *NewTagID("f432dc99-2e06-434d-b9ad-2b22e35b6fa4"),
		SoftwareName:    "Roadrunner software bundle",
		SoftwareVersion: "1.0.0",
		Entities: Entities{
			makeACMEEntityWithRoles(t,
				"tagCreator",
				"softwareCreator",
			),
		},
		Links: &Links{
			Link{
				Href: "d84fb5e2-d198-49b4-9d65-3a82421bf180",
				Rel:  *NewRel(RelParent),
			},
		},
	}

	var tv SoftwareIdentity

	err := tv.FromXML(testXML)
	assert.NoError(t, err)
	assert.Equal(t, expected, tv)
}

func TestTag_encodings_ok(t *testing.T) {
	tv := SoftwareIdentity{
		TagID:           *NewTagID("f432dc99-2e06-434d-b9ad-2b22e35b6fa4"),
		SoftwareName:    "Roadrunner software bundle",
		SoftwareVersion: "1.0.0",
		Entities: Entities{
			makeACMEEntityWithRoles(t,
				RoleTagCreator,
				RoleSoftwareCreator,
			),
		},
		Links: &Links{
			Link{
				Href: "d84fb5e2-d198-49b4-9d65-3a82421bf180",
				Rel:  *NewRel(RelParent),
			},
		},
	}

	actualCBOR, err := tv.ToCBOR()
	assert.NoError(t, err)
	assert.Equal(t, testCBOR, actualCBOR)

	actualJSON, err := tv.ToJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, string(testJSON), string(actualJSON))

	actualXML, err := tv.ToXML()
	assert.NoError(t, err)
	assert.Equal(t, testXML, actualXML)
}

func TestTag_RoundtripPSABundle(t *testing.T) {
	tv := SoftwareIdentity{
		TagID:           *NewTagID("example.acme.roadrunner-sw-v1-0-0"),
		SoftwareName:    "Roadrunner software bundle",
		SoftwareVersion: "1.0.0",
		Entities: Entities{
			makeACMEEntityWithRoles(t,
				RoleTagCreator,
				RoleSoftwareCreator,
				RoleAggregator,
			),
		},
		Links: &Links{
			Link{
				Href: "example.acme.roadrunner-hw-v1-0-0",
				Rel:  Rel{"psa-rot-compound"},
			},
			Link{
				Href: "example.acme.roadrunner-sw-bl-v1-0-0",
				Rel:  Rel{RelComponent},
			},
			Link{
				Href: "example.acme.roadrunner-sw-prot-v1-0-0",
				Rel:  Rel{RelComponent},
			},
			Link{
				Href: "example.acme.roadrunner-sw-arot-v1-0-0",
				Rel:  Rel{RelComponent},
			},
			Link{
				Href: NewTagID([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10}).URI(),
				Rel:  Rel{RelComponent},
			},
		},
	}
	/*
		a6                                      # map(6)
		   00                                   # unsigned(0)
		   78 21                                # text(33)
		      6578616d706c652e61636d652e726f616472756e6e65722d73772d76312d302d30 # "example.acme.roadrunner-sw-v1-0-0"
		   0c                                   # unsigned(12)
		   00                                   # unsigned(0)
		   01                                   # unsigned(1)
		   78 1a                                # text(26)
		      526f616472756e6e657220736f6674776172652062756e646c65 # "Roadrunner software bundle"
		   0d                                   # unsigned(13)
		   65                                   # text(5)
		      312e302e30                        # "1.0.0"
		   02                                   # unsigned(2)
		   a3                                   # map(3)
		      18 1f                             # unsigned(31)
		      68                                # text(8)
		         41434d45204c7464               # "ACME Ltd"
		      18 20                             # unsigned(32)
		      6c                                # text(12)
		         61636d652e6578616d706c65       # "acme.example"
		      18 21                             # unsigned(33)
		      83                                # array(3)
		         01                             # unsigned(1)
		         02                             # unsigned(2)
		         03                             # unsigned(3)
		   04                                   # unsigned(4)
		   85                                   # array(5)
		      a2                                # map(2)
		         18 26                          # unsigned(38)
		         78 21                          # text(33)
		            6578616d706c652e61636d652e726f616472756e6e65722d68772d76312d302d30 # "example.acme.roadrunner-hw-v1-0-0"
		         18 28                          # unsigned(40)
		         70                             # text(16)
		            7073612d726f742d636f6d706f756e64 # "psa-rot-compound"
		      a2                                # map(2)
		         18 26                          # unsigned(38)
		         78 24                          # text(36)
		            6578616d706c652e61636d652e726f616472756e6e65722d73772d626c2d76312d302d30 # "example.acme.roadrunner-sw-bl-v1-0-0"
		         18 28                          # unsigned(40)
		         02                             # unsigned(2)
		      a2                                # map(2)
		         18 26                          # unsigned(38)
		         78 26                          # text(38)
		            6578616d706c652e61636d652e726f616472756e6e65722d73772d70726f742d76312d302d30 # "example.acme.roadrunner-sw-prot-v1-0-0"
		         18 28                          # unsigned(40)
		         02                             # unsigned(2)
		      a2                                # map(2)
		         18 26                          # unsigned(38)
		         78 26                          # text(38)
		            6578616d706c652e61636d652e726f616472756e6e65722d73772d61726f742d76312d302d30 # "example.acme.roadrunner-sw-arot-v1-0-0"
		         18 28                          # unsigned(40)
		         02                             # unsigned(2)
		      a2                                # map(2)
		         18 26                          # unsigned(38)
		         78 29                          # text(41)
					77736469303a3031303230332d3435303630302d30372d3839306130302d306230633064306531660a30 # swid:01020304-0506-0708-090a-0b0c0d0e0f10
		         18 28                          # unsigned(40)
		         02                             # unsigned(2)
	*/
	expectedCBOR := []byte{
		0xa6, 0x00, 0x78, 0x21, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
		0x61, 0x63, 0x6d, 0x65, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x72, 0x75, 0x6e,
		0x6e, 0x65, 0x72, 0x2d, 0x73, 0x77, 0x2d, 0x76, 0x31, 0x2d, 0x30, 0x2d,
		0x30, 0x0c, 0x00, 0x01, 0x78, 0x1a, 0x52, 0x6f, 0x61, 0x64, 0x72, 0x75,
		0x6e, 0x6e, 0x65, 0x72, 0x20, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
		0x65, 0x20, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x0d, 0x65, 0x31, 0x2e,
		0x30, 0x2e, 0x30, 0x02, 0xa3, 0x18, 0x1f, 0x68, 0x41, 0x43, 0x4d, 0x45,
		0x20, 0x4c, 0x74, 0x64, 0x18, 0x20, 0x6c, 0x61, 0x63, 0x6d, 0x65, 0x2e,
		0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x21, 0x83, 0x01, 0x02,
		0x03, 0x04, 0x85, 0xa2, 0x18, 0x26, 0x78, 0x21, 0x65, 0x78, 0x61, 0x6d,
		0x70, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x72, 0x6f, 0x61,
		0x64, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x68, 0x77, 0x2d, 0x76,
		0x31, 0x2d, 0x30, 0x2d, 0x30, 0x18, 0x28, 0x70, 0x70, 0x73, 0x61, 0x2d,
		0x72, 0x6f, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64,
		0xa2, 0x18, 0x26, 0x78, 0x24, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
		0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x72, 0x75,
		0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x73, 0x77, 0x2d, 0x62, 0x6c, 0x2d, 0x76,
		0x31, 0x2d, 0x30, 0x2d, 0x30, 0x18, 0x28, 0x02, 0xa2, 0x18, 0x26, 0x78,
		0x26, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x6d,
		0x65, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
		0x2d, 0x73, 0x77, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x2d, 0x76, 0x31, 0x2d,
		0x30, 0x2d, 0x30, 0x18, 0x28, 0x02, 0xa2, 0x18, 0x26, 0x78, 0x26, 0x65,
		0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e,
		0x72, 0x6f, 0x61, 0x64, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x73,
		0x77, 0x2d, 0x61, 0x72, 0x6f, 0x74, 0x2d, 0x76, 0x31, 0x2d, 0x30, 0x2d,
		0x30, 0x18, 0x28, 0x02, 0xa2, 0x18, 0x26, 0x78, 0x29, 0x73, 0x77, 0x69,
		0x64, 0x3a, 0x30, 0x31, 0x30, 0x32, 0x30, 0x33, 0x30, 0x34, 0x2d, 0x30,
		0x35, 0x30, 0x36, 0x2d, 0x30, 0x37, 0x30, 0x38, 0x2d, 0x30, 0x39, 0x30,
		0x61, 0x2d, 0x30, 0x62, 0x30, 0x63, 0x30, 0x64, 0x30, 0x65, 0x30, 0x66,
		0x31, 0x30, 0x18, 0x28, 0x02,
	}

	roundTripper(t, tv, expectedCBOR)
}
