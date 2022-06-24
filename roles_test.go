// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO(tho) coalesce marshal tests for CBOR and JSON
// TODO(tho) coalesce unmarshal tests for CBOR and JSON

func TestRoles_Check(t *testing.T) {
	type TestVector struct {
		val []interface{}
	}
	tests := []struct {
		name       string
		testVector TestVector
		expected   error
	}{
		{
			name: "all known codepoints",
			testVector: TestVector{
				val: []interface{}{
					RoleAggregator,
					RoleDistributor,
					RoleLicensor,
					RoleMaintainer,
					RoleSoftwareCreator,
					RoleTagCreator,
				},
			},
			expected: nil,
		},
		{
			name: "unknown codepoints",
			testVector: TestVector{
				val: []interface{}{
					int64(1024),
					int64(8192),
				},
			},
			expected: nil,
		},
		{
			name: "strings only",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					"yourRole",
					"theirRole",
				},
			},
			expected: nil,
		},
		{
			name: "mixed",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					int64(1024),
				},
			},
			expected: nil,
		},
		{
			name: "singleton",
			testVector: TestVector{
				val: []interface{}{
					RoleMaintainer,
				},
			},
			expected: nil,
		},
		{
			name: "unhandled float in the mix",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					1024.1024,
				},
			},
			expected: errors.New("role MUST be int64 or string; got float64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := Roles{
				val: test.testVector.val,
			}
			actual := u.Check()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestRoles_String(t *testing.T) {
	type TestVector struct {
		val []interface{}
	}
	tests := []struct {
		name       string
		testVector TestVector
		expected   string
	}{
		{
			name: "all known codepoints",
			testVector: TestVector{
				val: []interface{}{
					RoleAggregator,
					RoleDistributor,
					RoleLicensor,
					RoleMaintainer,
					RoleSoftwareCreator,
					RoleTagCreator,
				},
			},
			expected: `aggregator distributor licensor maintainer softwareCreator tagCreator`,
		},
		{
			name: "unknown codepoints",
			testVector: TestVector{
				val: []interface{}{
					int64(1024),
					int64(8192),
				},
			},
			expected: `role(1024) role(8192)`,
		},
		{
			name: "strings only",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					"yourRole",
					"theirRole",
				},
			},
			expected: `myRole yourRole theirRole`,
		},
		{
			name: "mixed",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					int64(1024),
				},
			},
			expected: `myRole maintainer role(1024)`,
		},
		{
			name: "singleton",
			testVector: TestVector{
				val: []interface{}{
					RoleMaintainer,
				},
			},
			expected: `maintainer`,
		},
		{
			name: "unhandled float in the mix",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					1024.1024,
				},
			},
			expected: `myRole maintainer`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := Roles{
				val: test.testVector.val,
			}
			actual := u.String()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestRoles_MarshalJSON(t *testing.T) {
	type TestVector struct {
		val []interface{}
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    string
		expectedErr error
	}{
		{
			name: "all known codepoints",
			testVector: TestVector{
				val: []interface{}{
					RoleAggregator,
					RoleDistributor,
					RoleLicensor,
					RoleMaintainer,
					RoleSoftwareCreator,
					RoleTagCreator,
				},
			},
			expected:    `[ "aggregator", "distributor", "licensor", "maintainer", "softwareCreator", "tagCreator" ]`,
			expectedErr: nil,
		},
		{
			name: "unknown codepoints",
			testVector: TestVector{
				val: []interface{}{
					int64(1024),
					int64(8192),
				},
			},
			expected:    `[ 1024, 8192 ]`,
			expectedErr: nil,
		},
		{
			name: "strings only",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					"yourRole",
					"theirRole",
				},
			},
			expected:    `[ "myRole", "yourRole", "theirRole" ]`,
			expectedErr: nil,
		},
		{
			name: "mixed",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					int64(1024),
				},
			},
			expected:    `[ "myRole", "maintainer", 1024 ]`,
			expectedErr: nil,
		},
		{
			name: "singleton",
			testVector: TestVector{
				val: []interface{}{
					RoleMaintainer,
				},
			},
			expected:    `"maintainer"`,
			expectedErr: nil,
		},
		{
			name: "unhandled float in the mix",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					1024.1024,
				},
			},
			expected:    `__ignored__`,
			expectedErr: errors.New("unhandled type: float64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := Roles{
				val: test.testVector.val,
			}
			actual, err := u.MarshalJSON()
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.JSONEq(t, test.expected, string(actual))
			}
		})
	}
}

func TestRoles_UnmarshalJSON(t *testing.T) {
	type TestVector struct {
		val string
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    Roles
		expectedErr error
	}{
		{
			name: "all known strings",
			testVector: TestVector{
				val: `[ "aggregator", "distributor", "licensor", "maintainer", "softwareCreator", "tagCreator" ]`,
			},
			expected: Roles{
				val: []interface{}{
					RoleAggregator,
					RoleDistributor,
					RoleLicensor,
					RoleMaintainer,
					RoleSoftwareCreator,
					RoleTagCreator,
				},
			},
			expectedErr: nil,
		},
		{
			name: "all unknown strings",
			testVector: TestVector{
				val: `[ "hallelujatic", "myriagramme", "anterointernal" ]`,
			},
			expected: Roles{
				val: []interface{}{
					"hallelujatic",
					"myriagramme",
					"anterointernal",
				},
			},
			expectedErr: nil,
		},
		{
			name: "mixed",
			testVector: TestVector{
				val: `[ "hallelujatic", 1024, 1, "distributor" ]`,
			},
			expected: Roles{
				val: []interface{}{
					"hallelujatic",
					int64(1024),
					RoleTagCreator,
					RoleDistributor,
				},
			},
			expectedErr: nil,
		},
		{
			name: "unknown inner type (bool)",
			testVector: TestVector{
				val: `[ "tagCreator", false ]`,
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "silly singleton optimisations",
			testVector: TestVector{
				val: `"tagCreator"`,
			},
			expected: Roles{
				val: []interface{}{
					RoleTagCreator,
				},
			},
			expectedErr: nil,
		},
		{
			name: "accept singleton encased in array too",
			testVector: TestVector{
				val: `[ "tagCreator" ]`,
			},
			expected: Roles{
				val: []interface{}{
					RoleTagCreator,
				},
			},
			expectedErr: nil,
		},
		{
			name: "singleton with unhandled type",
			testVector: TestVector{
				val: `true`,
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "empty array",
			testVector: TestVector{
				val: `[]`,
			},
			expected: Roles{
				val: []interface{}{},
			},
			expectedErr: nil,
		},
		{
			name: "wrong outer type (object)",
			testVector: TestVector{
				val: `{}`,
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("unhandled type: map[string]interface {}"),
		},
		{
			name: "bad number (float)",
			testVector: TestVector{
				val: `[ 1024.1024 ]`,
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("number 1024.1024 is not int64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Roles{}
			err := actual.UnmarshalJSON([]byte(test.testVector.val))
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestRoles_MarshalCBOR(t *testing.T) {
	type TestVector struct {
		val []interface{}
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    []byte
		expectedErr error
	}{
		{
			name: "all known codepoints",
			testVector: TestVector{
				val: []interface{}{
					RoleTagCreator,
					RoleSoftwareCreator,
					RoleAggregator,
					RoleDistributor,
					RoleLicensor,
					RoleMaintainer,
				},
			},
			/*
				86    # array(6)
				   01 # unsigned(1)
				   02 # unsigned(2)
				   03 # unsigned(3)
				   04 # unsigned(4)
				   05 # unsigned(5)
				   06 # unsigned(6)
			*/
			expected:    []byte{0x86, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
			expectedErr: nil,
		},
		{
			name: "unknown codepoints",
			testVector: TestVector{
				val: []interface{}{
					int64(1024),
					int64(8192),
				},
			},
			/*
				82         # array(2)
				   19 0400 # unsigned(1024)
				   19 2000 # unsigned(8192)
			*/
			expected:    []byte{0x82, 0x19, 0x04, 0x00, 0x19, 0x20, 0x00},
			expectedErr: nil,
		},
		{
			name: "strings only",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					"yourRole",
					"theirRole",
				},
			},
			/*
				83                       # array(3)
				   66                    # text(6)
				      6d79526f6c65       # "myRole"
				   68                    # text(8)
				      796f7572526f6c65   # "yourRole"
				   69                    # text(9)
					  7468656972526f6c65 # "theirRole"
			*/
			expected: []byte{
				0x83, 0x66, 0x6d, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x68, 0x79,
				0x6f, 0x75, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x69, 0x74, 0x68,
				0x65, 0x69, 0x72, 0x52, 0x6f, 0x6c, 0x65,
			},
			expectedErr: nil,
		},
		{
			name: "mixed",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					int64(1024),
				},
			},
			/*
			   83                 # array(3)
			      66              # text(6)
			         6d79526f6c65 # "myRole"
			      06              # unsigned(6)
			      19 0400         # unsigned(1024)
			*/
			expected: []byte{
				0x83, 0x66, 0x6d, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x06, 0x19,
				0x04, 0x00,
			},
			expectedErr: nil,
		},
		{
			name: "singleton",
			testVector: TestVector{
				val: []interface{}{
					RoleMaintainer,
				},
			},
			/*
			   06 # unsigned(6)
			*/
			expected:    []byte{0x06},
			expectedErr: nil,
		},
		{
			name: "unhandled float in the mix",
			testVector: TestVector{
				val: []interface{}{
					"myRole",
					RoleMaintainer,
					1024.1024,
				},
			},
			expected:    []byte("__ignored__"),
			expectedErr: errors.New("number 1024.1024 is not int64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := Roles{
				val: test.testVector.val,
			}
			actual, err := u.MarshalCBOR()
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestRoles_UnmarshalCBOR(t *testing.T) {
	type TestVector struct {
		val []byte
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    Roles
		expectedErr error
	}{
		{
			name: "all known strings",
			testVector: TestVector{
				val: []byte{
					0x86, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
				},
			},
			expected: Roles{
				val: []interface{}{
					RoleTagCreator,
					RoleSoftwareCreator,
					RoleAggregator,
					RoleDistributor,
					RoleLicensor,
					RoleMaintainer,
				},
			},
			expectedErr: nil,
		},
		{
			name: "all unknown strings",
			testVector: TestVector{
				/*
				   83                                 # array(3)
				      6c                              # text(12)
				         68616c6c656c756a61746963     # "hallelujatic"
				      6b                              # text(11)
				         6d797269616772616d6d65       # "myriagramme"
				      6e                              # text(14)
				   	  616e7465726f696e7465726e616c # "anterointernal"
				*/
				val: []byte{
					0x83, 0x6c, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x75,
					0x6a, 0x61, 0x74, 0x69, 0x63, 0x6b, 0x6d, 0x79, 0x72,
					0x69, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x6e,
					0x61, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x69, 0x6e, 0x74,
					0x65, 0x72, 0x6e, 0x61, 0x6c,
				},
			},
			expected: Roles{
				val: []interface{}{
					"hallelujatic",
					"myriagramme",
					"anterointernal",
				},
			},
			expectedErr: nil,
		},
		{
			name: "mixed",
			testVector: TestVector{
				/*
				   84                             # array(4)
				      6c                          # text(12)
				         68616c6c656c756a61746963 # "hallelujatic"
				      19 0400                     # unsigned(1024)
				      01                          # unsigned(1)
				      04                          # unsigned(4)
				*/
				val: []byte{
					0x84, 0x6c, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x75,
					0x6a, 0x61, 0x74, 0x69, 0x63, 0x19, 0x04, 0x00, 0x01,
					0x04,
				},
			},
			expected: Roles{
				val: []interface{}{
					"hallelujatic",
					int64(1024),
					RoleTagCreator,
					RoleDistributor,
				},
			},
			expectedErr: nil,
		},
		{
			name: "unknown inner type (bool)",
			testVector: TestVector{
				/*
				   82    # array(2)
				      01 # unsigned(1)
				      f4 # primitive(20)
				*/
				val: []byte{0x82, 0x01, 0xf4},
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "silly singleton optimisations",
			testVector: TestVector{
				/* 01 # unsigned(1) */
				val: []byte{0x01},
			},
			expected: Roles{
				val: []interface{}{
					RoleTagCreator,
				},
			},
			expectedErr: nil,
		},
		{
			name: "accept singleton encased in array too",
			testVector: TestVector{
				/*
				   81    # array(1)
				      01 # unsigned(1)
				*/
				val: []byte{0x81, 0x01},
			},
			expected: Roles{
				val: []interface{}{
					RoleTagCreator,
				},
			},
			expectedErr: nil,
		},
		{
			name: "singleton with unhandled type",
			testVector: TestVector{
				/* f5 # primitive(21) */
				val: []byte{0xf5},
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "empty array",
			testVector: TestVector{
				/* 80 # array(0) */
				val: []byte{0x80},
			},
			expected: Roles{
				val: []interface{}{},
			},
			expectedErr: nil,
		},
		{
			name: "wrong outer type (map)",
			testVector: TestVector{
				/* a0 # map(0) */
				val: []byte{0xa0},
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			/* NOTE: this differs from its JSON counterpart */
			expectedErr: errors.New("unhandled type: map[interface {}]interface {}"),
		},
		{
			name: "bad number (float)",
			testVector: TestVector{
				/* fb 40900068db8bac71 # primitive(4652218865433685105) */
				val: []byte{
					0xfb, 0x40, 0x90, 0x00, 0x68, 0xdb, 0x8b, 0xac, 0x71,
				},
			},
			expected: Roles{
				val: []interface{}{
					"__ignored__",
				},
			},
			expectedErr: errors.New("number 1024.1024 is not int64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Roles{}
			err := actual.UnmarshalCBOR(test.testVector.val)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}
