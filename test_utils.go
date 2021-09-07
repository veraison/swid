// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func MustHexDecode(t *testing.T, s string) []byte {
	data, err := hex.DecodeString(s)
	if t != nil {
		require.Nil(t, err)
	} else if err != nil {
		panic(err)
	}
	return data
}
