// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
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
		expectedErr string
	}{
		{
			name: "good stuff",
			testVector: TestVector{
				val: `"sha-256;3q2+7w=="`,
			},
			expected: HashEntry{
				HashAlgID: 1,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expectedErr: "",
		},
		{
			name: "good stuff (legacy)",
			testVector: TestVector{
				val: `"sha-256:3q2+7w=="`,
			},
			expected: HashEntry{
				HashAlgID: 1,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expectedErr: "",
		},
		{
			name: "no match for alg",
			testVector: TestVector{
				val: `"sha0-512;3q2+7w=="`,
			},
			expected:    HashEntry{},
			expectedErr: "unknown hash algorithm sha0-512",
		},
		{
			name: "empty string",
			testVector: TestVector{
				val: `""`,
			},
			expected:    HashEntry{},
			expectedErr: "bad format: expecting <hash-alg-string>;<hash-value>",
		},
		{
			name: "empty algo and hash value",
			testVector: TestVector{
				val: `";"`,
			},
			expected:    HashEntry{},
			expectedErr: "bad format: expecting <hash-alg-string>;<hash-value>",
		},
		{
			name: "whitespaces",
			testVector: TestVector{
				val: `" ; "`,
			},
			expected:    HashEntry{},
			expectedErr: "bad format: expecting <hash-alg-string>;<hash-value>",
		}, {
			name: "excess material",
			testVector: TestVector{
				val: `"sha-256;3q2+7w==;EXCESS MATERIAL"`,
			},
			expected:    HashEntry{},
			expectedErr: "bad format: expecting <hash-alg-string>;<hash-value>",
		},
		{
			name: "missing algo",
			testVector: TestVector{
				val: `";3q2+7w=="`,
			},
			expected:    HashEntry{},
			expectedErr: "bad format: expecting <hash-alg-string>;<hash-value>",
		},
		{
			name: "missing hash value",
			testVector: TestVector{
				val: `"sha-256;"`,
			},
			expected:    HashEntry{},
			expectedErr: "bad format: expecting <hash-alg-string>;<hash-value>",
		},
		{
			name: "corrupt base64 for value",
			testVector: TestVector{
				val: `"sha-256;....Caligula...."`,
			},
			expected:    HashEntry{},
			expectedErr: "illegal base64 data at input byte 0",
		},
		{
			name: "unexpected container",
			testVector: TestVector{
				val: `[ "sha-256", "3q2+7w==" ]`,
			},
			expected:    HashEntry{},
			expectedErr: "expecting string, found []interface {} instead",
		},
		{
			name: "invalid json",
			testVector: TestVector{
				val: `[ `,
			},
			expected:    HashEntry{},
			expectedErr: "unexpected end of JSON input",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := HashEntry{}
			err := actual.UnmarshalJSON([]byte(test.testVector.val))
			assert.Equal(t, test.expected, actual)
			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
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
				HashAlgID: 6,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expected:    `"sha-256-32;3q2+7w=="`,
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

func TestHashEntry_AlgIDToString(t *testing.T) {
	tests := []struct {
		name       string
		testVector HashEntry
		expected   string
	}{
		{
			name: "good stuff",
			testVector: HashEntry{
				HashAlgID: 1,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expected: "sha-256",
		},
		{
			name: "unknown hash algo",
			testVector: HashEntry{
				HashAlgID: 1000,
				HashValue: []byte{0xde, 0xad, 0xbe, 0xef},
			},
			expected: "alg-id(1000)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := test.testVector
			actual := u.AlgIDToString()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestHashEntry_Set_OK(t *testing.T) {
	tvs := []struct {
		alg uint64
		val []byte
	}{
		{
			alg: Sha256,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75"),
		},
		{
			alg: Sha256_128,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f36"),
		},
		{
			alg: Sha256_120,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f"),
		},
		{
			alg: Sha256_96,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3a"),
		},
		{
			alg: Sha256_64,
			val: MustHexDecode(t, "e45b72f5c0c0b572"),
		},
		{
			alg: Sha256_32,
			val: MustHexDecode(t, "e45b72ab"),
		},
		{
			alg: Sha384,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36"),
		},
		{
			alg: Sha512,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75"),
		},
		{
			alg: Sha3_224,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f36e45b72f5c0c0b572db4d8d3a"),
		},
		{
			alg: Sha3_256,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75"),
		},
		{
			alg: Sha3_384,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36"),
		},
		{
			alg: Sha3_512,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75"),
		},
	}

	for _, tv := range tvs {
		var h HashEntry
		assert.Nil(t, h.Set(tv.alg, tv.val))
	}
}

func TestHashEntry_Set_mismatched_input(t *testing.T) {
	tvs := []struct {
		alg uint64
		val []byte
		exp string
	}{
		{
			alg: Sha256,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d"),
			exp: "length mismatch for hash algorithm sha-256: want 32 bytes, got 31",
		},
		{
			alg: Sha256_128,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f"),
			exp: "length mismatch for hash algorithm sha-256-128: want 16 bytes, got 15",
		},
		{
			alg: Sha256_120,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e9"),
			exp: "length mismatch for hash algorithm sha-256-120: want 15 bytes, got 14",
		},
		{
			alg: Sha256_96,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d"),
			exp: "length mismatch for hash algorithm sha-256-96: want 12 bytes, got 11",
		},
		{
			alg: Sha256_64,
			val: MustHexDecode(t, "e45b72f5c0c0b5"),
			exp: "length mismatch for hash algorithm sha-256-64: want 8 bytes, got 7",
		},
		{
			alg: Sha256_32,
			val: MustHexDecode(t, "e45b72"),
			exp: "length mismatch for hash algorithm sha-256-32: want 4 bytes, got 3",
		},
		{
			alg: Sha384,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f"),
			exp: "length mismatch for hash algorithm sha-384: want 48 bytes, got 47",
		},
		{
			alg: Sha512,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d"),
			exp: "length mismatch for hash algorithm sha-512: want 64 bytes, got 63",
		},
		{
			alg: Sha3_224,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f36e45b72f5c0c0b572db4d8d"),
			exp: "length mismatch for hash algorithm sha3-224: want 28 bytes, got 27",
		},
		{
			alg: Sha3_256,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d"),
			exp: "length mismatch for hash algorithm sha3-256: want 32 bytes, got 31",
		},
		{
			alg: Sha3_384,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f"),
			exp: "length mismatch for hash algorithm sha3-384: want 48 bytes, got 47",
		},
		{
			alg: Sha3_512,
			val: MustHexDecode(t, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d"),
			exp: "length mismatch for hash algorithm sha3-512: want 64 bytes, got 63",
		},
	}

	for _, tv := range tvs {
		var h HashEntry
		assert.EqualError(t, h.Set(tv.alg, tv.val), tv.exp)
	}
}

func TestHashEntry_ValidHashEntry_unknown_algo(t *testing.T) {
	var unknownAlgID uint64 = 0
	err := ValidHashEntry(unknownAlgID, []byte{})
	assert.EqualError(t, err, "unknown hash algorithm 0")
}
