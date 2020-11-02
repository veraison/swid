package swid

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUse_MarshalJSON(t *testing.T) {
	type TestVector struct {
		val interface{}
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    string
		expectedErr error
	}{
		{
			name: "known codepoint optional",
			testVector: TestVector{
				val: UseOptional,
			},
			expected:    `"optional"`,
			expectedErr: nil,
		},
		{
			name: "known codepoint required",
			testVector: TestVector{
				val: UseRequired,
			},
			expected:    `"required"`,
			expectedErr: nil,
		},
		{
			name: "known codepoint recommended",
			testVector: TestVector{
				val: UseRecommended,
			},
			expected:    `"recommended"`,
			expectedErr: nil,
		},
		{
			name: "unknown codepoint 1024",
			testVector: TestVector{
				val: uint64(1024),
			},
			expected:    `1024`,
			expectedErr: nil,
		},
		{
			name: "a random string",
			testVector: TestVector{
				val: "rostrombolicius",
			},
			expected:    `"rostrombolicius"`,
			expectedErr: nil,
		},
		{
			name: "an unknown type for Use",
			testVector: TestVector{
				val: float64(4.343),
			},
			expected:    `__ignored__`,
			expectedErr: errors.New("unhandled type: float64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := Use{
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

func TestUse_UnmarshalJSON(t *testing.T) {
	type TestVector struct {
		val string
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    Use
		expectedErr error
	}{
		{
			name: "known string optional",
			testVector: TestVector{
				val: `"optional"`,
			},
			expected:    Use{UseOptional},
			expectedErr: nil,
		},
		{
			name: "known string required",
			testVector: TestVector{
				val: `"required"`,
			},
			expected:    Use{UseRequired},
			expectedErr: nil,
		},
		{
			name: "known string recommended",
			testVector: TestVector{
				val: `"recommended"`,
			},
			expected:    Use{UseRecommended},
			expectedErr: nil,
		},
		{
			name: "unknown string rostrombolicious",
			testVector: TestVector{
				val: `"rostrombolicious"`,
			},
			expected:    Use{"rostrombolicious"},
			expectedErr: nil,
		},
		{
			name: "integral codepoint",
			testVector: TestVector{
				val: `1024`,
			},
			expected:    Use{uint64(1024)},
			expectedErr: nil,
		},
		{
			name: "floating codepoint",
			testVector: TestVector{
				val: `1024.1024`,
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("number 1024.1024 is not uint64"),
		},
		{
			name: "JSON type object",
			testVector: TestVector{
				val: `{}`,
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: map[string]interface {}"),
		},
		{
			name: "JSON type array",
			testVector: TestVector{
				val: `[]`,
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: []interface {}"),
		},
		{
			name: "JSON type true",
			testVector: TestVector{
				val: `true`,
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "JSON type false",
			testVector: TestVector{
				val: `false`,
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "JSON type null",
			testVector: TestVector{
				val: `null`,
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: <nil>"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Use{}
			err := actual.UnmarshalJSON([]byte(test.testVector.val))
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestUse_MarshalCBOR(t *testing.T) {
	type TestVector struct {
		val interface{}
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    []byte
		expectedErr error
	}{
		{
			name: "uint64 codepoint (known)",
			testVector: TestVector{
				val: UseOptional,
			},
			/*
				01 # unsigned(1)
			*/
			expected:    []byte{0x01},
			expectedErr: nil,
		},
		{
			name: "uint64 codepoint (unknown)",
			testVector: TestVector{
				val: uint64(1024),
			},
			/*
				19 0400 # unsigned(1024)
			*/
			expected:    []byte{0x19, 0x04, 0x00},
			expectedErr: nil,
		},
		{
			name: "a random string",
			testVector: TestVector{
				val: "rostrombolicius",
			},
			/*
				6f                                # text(15)
				   726f7374726f6d626f6c6963697573 # "rostrombolicius"
			*/
			expected: []byte{
				0x6f, 0x72, 0x6f, 0x73, 0x74, 0x72, 0x6f, 0x6d, 0x62, 0x6f,
				0x6c, 0x69, 0x63, 0x69, 0x75, 0x73,
			},
			expectedErr: nil,
		},
		{
			name: "an unknown type for Use",
			testVector: TestVector{
				val: float64(4.343),
			},
			expected:    []byte{0x00},
			expectedErr: errors.New("number 4.343 is not uint64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := Use{
				val: test.testVector.val,
			}
			actual, err := u.MarshalCBOR()
			t.Logf("CBOR(hex): %x\n", actual)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestUse_UnmarshalCBOR(t *testing.T) {
	type TestVector struct {
		val []byte
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    Use
		expectedErr error
	}{
		{
			name: "known string optional",
			testVector: TestVector{
				/*
					68                  # text(8)
					   6f7074696f6e616c # "optional"
				*/
				val: []byte{
					0x68, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
				},
			},
			expected:    Use{UseOptional},
			expectedErr: nil,
		},
		{
			name: "known string required",
			testVector: TestVector{
				/*
					68                  # text(8)
					   7265717569726564 # "required"
				*/
				val: []byte{
					0x68, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
				},
			},
			expected:    Use{UseRequired},
			expectedErr: nil,
		},
		{
			name: "known string recommended",
			testVector: TestVector{
				/*
					6b                        # text(11)
					   7265636f6d6d656e646564 # "recommended"
				*/
				val: []byte{
					0x6b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
					0x64, 0x65, 0x64,
				},
			},
			expected:    Use{UseRecommended},
			expectedErr: nil,
		},
		{
			name: "unknown string rostrombolicious",
			testVector: TestVector{
				/*
					70                                  # text(16)
					   726f7374726f6d626f6c6963696f7573 # "rostrombolicious"
				*/
				val: []byte{
					0x70, 0x72, 0x6f, 0x73, 0x74, 0x72, 0x6f, 0x6d, 0x62,
					0x6f, 0x6c, 0x69, 0x63, 0x69, 0x6f, 0x75, 0x73,
				},
			},
			expected:    Use{"rostrombolicious"},
			expectedErr: nil,
		},
		{
			name: "integral codepoint",
			testVector: TestVector{
				/*
					19 0400 # unsigned(1024)
				*/
				val: []byte{
					0x19, 0x04, 0x00,
				},
			},
			expected:    Use{uint64(1024)},
			expectedErr: nil,
		},
		{
			name: "floating codepoint",
			testVector: TestVector{
				/*
					fb 40900068db8bac71 # primitive(4652218865433685105)
				*/
				val: []byte{
					0xfb, 0x40, 0x90, 0x00, 0x68, 0xdb, 0x8b, 0xac, 0x71,
				},
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("number 1024.1024 is not uint64"),
		},
		{
			name: "CBOR type map",
			testVector: TestVector{
				/*
					a0 # map(0)
				*/
				val: []byte{0xa0},
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: map[interface {}]interface {}"),
		},
		{
			name: "CBOR type array",
			testVector: TestVector{
				/*
					80 # array(0)
				*/
				val: []byte{0x80},
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: []interface {}"),
		},
		{
			name: "CBOR type true",
			testVector: TestVector{
				/*
					f5 # primitive(21)
				*/
				val: []byte{0xf5},
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "CBOR type false",
			testVector: TestVector{
				/*
					f4 # primitive(20)
				*/
				val: []byte{0xf4},
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: bool"),
		},
		{
			name: "CBOR type null",
			testVector: TestVector{
				/*
					f4 # primitive(22)
				*/
				val: []byte{0xf6},
			},
			expected:    Use{"__ignored__"},
			expectedErr: errors.New("unhandled type: <nil>"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Use{}
			err := actual.UnmarshalCBOR(test.testVector.val)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestUse_String(t *testing.T) {
	tests := []struct {
		name       string
		testVector Use
		expected   string
	}{
		{
			name:       "known code optional",
			testVector: Use{UseOptional},
			expected:   "optional",
		},
		{
			name:       "known code required",
			testVector: Use{UseRequired},
			expected:   "required",
		},
		{
			name:       "known code recommended",
			testVector: Use{UseRecommended},
			expected:   "recommended",
		},
		{
			name:       "unknown code 1024",
			testVector: Use{uint64(1024)},
			expected:   "use(1024)",
		},
		{
			name:       "unknown string code",
			testVector: Use{"random stuff"},
			expected:   "random stuff",
		},
		{
			name:       "unknown type is ignored",
			testVector: Use{float32(12.23)},
			expected:   "",
		},
		{
			name:       "empty code is ignored",
			testVector: Use{},
			expected:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.testVector.String()
			assert.Equal(t, test.expected, actual)
		})
	}
}
