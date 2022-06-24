// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionScheme_Set(t *testing.T) {
	tests := []struct {
		name     string
		tv       int64
		expected error
	}{
		{
			name:     "known codepoints (1)",
			tv:       VersionSchemeMultipartNumeric,
			expected: nil,
		},
		{
			name:     "known codepoints (2)",
			tv:       VersionSchemeMultipartNumericSuffix,
			expected: nil,
		},
		{
			name:     "known codepoints (3)",
			tv:       VersionSchemeAlphaNumeric,
			expected: nil,
		},
		{
			name:     "known codepoints (4)",
			tv:       VersionSchemeDecimal,
			expected: nil,
		},
		{
			name:     "known codepoints (5)",
			tv:       VersionSchemeSemVer,
			expected: nil,
		},
		{
			name:     "unknown codepoint",
			tv:       19283937,
			expected: errors.New("unknown version scheme 19283937"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var u VersionScheme
			actual := u.SetCode(test.tv)
			assert.Equal(t, test.expected, actual)
		})
	}
}
