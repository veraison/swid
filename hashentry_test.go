package swid

import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashEntry_UnmarshalJSON(t *testing.T) {
	type TestVector struct {
		val string
	}
	tests := []struct {
		name        string
		testVector  TestVector
		expected    HashEntry
		expectedErr error
	}{
		{
			name: "good stuff",
			testVector: TestVector{
				val: `"sha-256:3q2+7w=="`,
			},
			expected: HashEntry{
				HashAlgID: 1,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expectedErr: nil,
		},
		{
			name: "no match for alg",
			testVector: TestVector{
				val: `"sha0-512:3q2+7w=="`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("unknown hash algorithm sha0-512"),
		},
		{
			name: "empty string",
			testVector: TestVector{
				val: `""`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("bad format: expecting <hash-alg-string>:<hash-value>"),
		},
		{
			name: "empty algo and hash value",
			testVector: TestVector{
				val: `":"`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("bad format: expecting <hash-alg-string>:<hash-value>"),
		},
		{
			name: "whitespaces",
			testVector: TestVector{
				val: `" : "`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("bad format: expecting <hash-alg-string>:<hash-value>"),
		}, {
			name: "excess material",
			testVector: TestVector{
				val: `"sha-256:3q2+7w==:EXCESS MATERIAL"`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("bad format: expecting <hash-alg-string>:<hash-value>"),
		},
		{
			name: "missing algo",
			testVector: TestVector{
				val: `":3q2+7w=="`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("bad format: expecting <hash-alg-string>:<hash-value>"),
		},
		{
			name: "missing hash value",
			testVector: TestVector{
				val: `"sha-256:"`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("bad format: expecting <hash-alg-string>:<hash-value>"),
		},
		{
			name: "corrupt base64 for value",
			testVector: TestVector{
				val: `"sha-256:....Caligula...."`,
			},
			expected:    HashEntry{},
			expectedErr: base64.CorruptInputError(0),
		},
		{
			name: "unexpected container",
			testVector: TestVector{
				val: `[ "sha-256", "3q2+7w==" ]`,
			},
			expected:    HashEntry{},
			expectedErr: errors.New("expecting string, found []interface {} instead"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := HashEntry{}
			err := actual.UnmarshalJSON([]byte(test.testVector.val))
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestHashEntry_MarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		testVector  HashEntry
		expected    string
		expectedErr error
	}{
		{
			name: "good stuff",
			testVector: HashEntry{
				HashAlgID: 1,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expected:    `"sha-256:3q2+7w=="`,
			expectedErr: nil,
		},
		{
			name: "unknown hash algo",
			testVector: HashEntry{
				HashAlgID: 1024,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expected:    `__ignored__`,
			expectedErr: errors.New("unknown hash algorithm ID 1024"),
		},
		{
			name: "empty hash value",
			testVector: HashEntry{
				HashAlgID: 1,
				HashValue: []byte{},
			},
			expected:    `__ignored__`,
			expectedErr: errors.New("empty hash value"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := test.testVector
			actual, err := u.MarshalJSON()
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.JSONEq(t, test.expected, string(actual))
			}
		})
	}
}
