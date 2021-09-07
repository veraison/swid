// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntity_RoundtripOneRoleUint(t *testing.T) {
	tv := Entity{}

	err := tv.SetRoles(RoleTagCreator)
	assert.Nil(t, err)

	err = tv.SetEntityName("ACME Ltd")
	assert.Nil(t, err)

	err = tv.SetRegID("https://acme.example")
	assert.Nil(t, err)

	err = tv.SetThumbprint(Sha256_32, []byte{0x00, 0x01, 0x02, 0x03})
	assert.Nil(t, err)

	// When only one element is present role does not use the array wrap:
	// a4                     # map(4)
	//    18 1f               # unsigned(31)
	//    68                  # text(8)
	//       41434d45204c7464 # "ACME Ltd"
	//    18 20               # unsigned(32)
	//    74                  # text(20)
	//       68747470733a2f2f61636d652e6578616d706c65 # "https://acme.example"
	//    18 21               # unsigned(33)
	//    01                  # unsigned(1)
	//	  18 22               # unsigned(34)
	//	  82                  # array(2)
	//	     06               # unsigned(6)
	//       44               # bytes(4)
	//          00010203      # "\x00\x01\x02\x03"
	expectedCBOR := []byte{
		0xa4, 0x18, 0x1f, 0x68, 0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74,
		0x64, 0x18, 0x20, 0x74, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
		0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
		0x6c, 0x65, 0x18, 0x21, 0x01, 0x18, 0x22, 0x82, 0x06, 0x44, 0x00,
		0x01, 0x02, 0x03,
	}

	actual := roundTripper(t, tv, expectedCBOR).(Entity)

	assert.Equal(t, actual.Roles.String(), "tagCreator")
}

func TestEntity_RoundtripOneRoleText(t *testing.T) {
	tv := Entity{}

	err := tv.SetEntityName("ACME Ltd")
	assert.Nil(t, err)

	err = tv.SetRoles("trivirgate")
	assert.Nil(t, err)

	// When only one element is present role does not use the array wrap:
	// a2                         # map(2)
	//    18 1f                   # unsigned(31)
	//    68                      # text(8)
	//       41434d45204c7464     # "ACME Ltd"
	//    18 21                   # unsigned(33)
	//    6a                      # text(10)
	//       74726976697267617465 # "trivirgate"
	expectedCBOR := []byte{
		0xa2, 0x18, 0x1f, 0x68, 0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74,
		0x64, 0x18, 0x21, 0x6a, 0x74, 0x72, 0x69, 0x76, 0x69, 0x72, 0x67,
		0x61, 0x74, 0x65,
	}

	actual := roundTripper(t, tv, expectedCBOR).(Entity)

	assert.Equal(t, actual.Roles.String(), "trivirgate")
}

func TestEntity_RoundtripMultipleRoles(t *testing.T) {
	tv := Entity{}

	err := tv.SetRoles(RoleTagCreator, RoleAggregator, "weird-new-role", uint64(20))
	assert.Nil(t, err)

	err = tv.SetEntityName("ACME Ltd")
	assert.Nil(t, err)

	// When more than one element is present, use the array wrap:
	// a2                                    # map(2)
	//    18 1f                              # unsigned(31)
	//    68                                 # text(8)
	//       41434d45204c7464                # "ACME Ltd"
	//    18 21                              # unsigned(33)
	//    84                                 # array(4)
	//       01                              # unsigned(1)
	//       03                              # unsigned(3)
	//       6e                              # text(14)
	//          77656972642d6e65772d726f6c65 # "weird-new-role"
	//       14                              # unsigned(20)
	expectedCBOR := []byte{
		0xa2, 0x18, 0x1f, 0x68, 0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64,
		0x18, 0x21, 0x84, 0x01, 0x03, 0x6e, 0x77, 0x65, 0x69, 0x72, 0x64, 0x2d,
		0x6e, 0x65, 0x77, 0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x14,
	}

	actual := roundTripper(t, tv, expectedCBOR).(Entity)

	assert.Equal(t, actual.Roles.String(), "tagCreator aggregator weird-new-role role(20)")
}

func TestEntity_BadRoleType(t *testing.T) {
	tv := Entity{}

	err := tv.SetRoles(float32(1.23))
	assert.EqualError(t, err, "role MUST be uint64 or string; got float32")

	type XYZ struct{ uint64 }
	xyz := XYZ{1}

	err = tv.SetRoles(xyz)
	assert.EqualError(t, err, "role MUST be uint64 or string; got swid.XYZ")
}

func TestEntity_RoundtripWithGlobalAttributesLang(t *testing.T) {
	tv := Entity{}

	err := tv.SetEntityName("ACME Ltd")
	assert.Nil(t, err)

	err = tv.SetRoles(RoleTagCreator)
	assert.Nil(t, err)

	err = tv.SetLang("en-GB")
	assert.Nil(t, err)

	// a3                     # map(3)
	//    0f                  # unsigned(15)
	//    65                  # text(5)
	//       656e2d4742       # "en-GB"
	//    18 1f               # unsigned(31)
	//    68                  # text(8)
	//       41434d45204c7464 # "ACME Ltd"
	//    18 21               # unsigned(33)
	//    01                  # unsigned(1)
	expectedCBOR := []byte{
		0xa3, 0x0f, 0x65, 0x65, 0x6e, 0x2d, 0x47, 0x42, 0x18, 0x1f, 0x68,
		0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x18, 0x21, 0x01,
	}

	roundTripper(t, tv, expectedCBOR)
}

func TestEntity_GlobalAttributesUnknownSkipped(t *testing.T) {
	// a3                           # map(3)
	//    0f                        # unsigned(15)
	//    65                        # text(5)
	//       656e2d4742             # "en-GB"
	//    6b                        # text(11)        -.
	//       616e792d72756262697368 # "any-rubbish"    |
	//    82                        # array(2)         | any-attributes skipped
	//       19 07ee                # unsigned(2030)   |
	//       39 0113                # negative(275)   -'
	//    18 1f                     # unsigned(31)
	//    68                        # text(8)
	//       41434d45204c7464       # "ACME Ltd"
	//    18 21                     # unsigned(33)
	//    82                        # array(2)
	//       05                     # unsigned(5)
	//       04                     # unsigned(4)
	tv := []byte{
		0xa4, 0x0f, 0x65, 0x65, 0x6e, 0x2d, 0x47, 0x42, 0x6b, 0x61, 0x6e,
		0x79, 0x2d, 0x72, 0x75, 0x62, 0x62, 0x69, 0x73, 0x68, 0x82, 0x19,
		0x07, 0xee, 0x39, 0x01, 0x13, 0x18, 0x1f, 0x68, 0x41, 0x43, 0x4d,
		0x45, 0x20, 0x4c, 0x74, 0x64, 0x18, 0x21, 0x82, 0x05, 0x04,
	}

	actual := Entity{}
	err := dm.Unmarshal(tv, &actual)

	assert.Nil(t, err)
	assert.NotNil(t, actual.GetLang())
	assert.Equal(t, "en-GB", actual.Lang)
	assert.Equal(t, actual.Roles.String(), "licensor distributor")

	// Check that when re-encoding the extra "any-rubbish" part is not forwarded
	data, err := em.Marshal(actual)

	assert.Nil(t, err)
	assert.Equal(t,
		// a3                     # map(3)
		//    0f                  # unsigned(15)
		//    65                  # text(5)
		//       656e2d4742       # "en-GB"
		//    18 1f               # unsigned(31)
		//    68                  # text(8)
		//       41434d45204c7464 # "ACME Ltd"
		//    18 21               # unsigned(33)
		//    82                  # array(2)
		//       05               # unsigned(5)
		//       04               # unsigned(4)
		[]byte{
			0xa3, 0x0f, 0x65, 0x65, 0x6e, 0x2d, 0x47, 0x42, 0x18, 0x1f, 0x68,
			0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x18, 0x21, 0x82,
			0x05, 0x04,
		},
		data,
	)
}

func TestEntity_UnnknownGlobalAttributesAndExtensionsToo(t *testing.T) {
	/*
		a6                                 # map(6)
		   38 67                           # negative(103)
		   69                              # text(9)
		      706f7263686c657373           # "porchless"
		   2c                              # negative(12)
		   18 f3                           # unsigned(243)
		   19 1254                         # unsigned(4692)
		   83                              # array(3)
		      19 0e1e                      # unsigned(3614)
		      19 08bf                      # unsigned(2239)
		      19 084e                      # unsigned(2126)
		   18 1f                           # unsigned(31)
		   69                              # text(9)
		      6175746f6d61746963           # "automatic"
		   18 21                           # unsigned(33)
		   04                              # unsigned(4)
		   6a                              # text(10)
		      646972656374696f6e73         # "directions"
		   a4                              # map(4)
		      64                           # text(4)
		         6c656674                  # "left"
		      a2                           # map(2)
		         6a                        # text(10)
		            76656e65726174697665   # "venerative"
		         39 0103                   # negative(259)
		         6b                        # text(11)
		            756e636f6e6365726e6564 # "unconcerned"
		         19 0995                   # unsigned(2453)
		      65                           # text(5)
		         7269676874                # "right"
		      19 10f8                      # unsigned(4344)
		      62                           # text(2)
		         7570                      # "up"
		      fb 3fe1952bea18002e          # primitive(4603124309992931374)
		      64                           # text(4)
		         646f776e                  # "down"
			  fb 3fe5be02d10cdd5f          # primitive(4604295113362693471)
	*/
	tv := []byte{
		0xa6, 0x38, 0x67, 0x69, 0x70, 0x6f, 0x72, 0x63, 0x68, 0x6c, 0x65,
		0x73, 0x73, 0x2c, 0x18, 0xf3, 0x19, 0x12, 0x54, 0x83, 0x19, 0x0e,
		0x1e, 0x19, 0x08, 0xbf, 0x19, 0x08, 0x4e, 0x18, 0x1f, 0x69, 0x61,
		0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x21, 0x04,
		0x6a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
		0xa4, 0x64, 0x6c, 0x65, 0x66, 0x74, 0xa2, 0x6a, 0x76, 0x65, 0x6e,
		0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x39, 0x01, 0x03, 0x6b,
		0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x65, 0x64,
		0x19, 0x09, 0x95, 0x65, 0x72, 0x69, 0x67, 0x68, 0x74, 0x19, 0x10,
		0xf8, 0x62, 0x75, 0x70, 0xfb, 0x3f, 0xe1, 0x95, 0x2b, 0xea, 0x18,
		0x00, 0x2e, 0x64, 0x64, 0x6f, 0x77, 0x6e, 0xfb, 0x3f, 0xe5, 0xbe,
		0x02, 0xd1, 0x0c, 0xdd, 0x5f,
	}

	actual := Entity{}
	err := dm.Unmarshal(tv, &actual)

	assert.Nil(t, err)

	// Check that when re-encoding the extra unknown extension and general
	// options part is not forwarded
	data, err := em.Marshal(actual)

	assert.Nil(t, err)
	assert.Equal(t,
		/*
			a2                       # map(2)
			   18 1f                 # unsigned(31)
			   69                    # text(9)
			      6175746f6d61746963 # "automatic"
			   18 21                 # unsigned(33)
			   04                    # unsigned(4)
		*/
		[]byte{
			0xa2, 0x18, 0x1f, 0x69, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
			0x69, 0x63, 0x18, 0x21, 0x04,
		},
		data,
	)
}
