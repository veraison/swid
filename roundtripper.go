// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"reflect"
	"testing"

	cbor "github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// marshal + unmarshal
func roundTripper(t *testing.T, tv interface{}, expectedCBOR []byte) interface{} {
	encMode, err := cbor.EncOptions{TimeTag: cbor.EncTagRequired}.EncMode()
	require.Nil(t, err)

	data, err := encMode.Marshal(tv)

	assert.Nil(t, err)
	t.Logf("CBOR(hex): %x\n", data)
	assert.Equal(t, expectedCBOR, data)

	decMode, err := cbor.DecOptions{TimeTag: cbor.DecTagOptional}.DecMode()
	require.Nil(t, err)

	actual := reflect.New(reflect.TypeOf(tv))
	err = decMode.Unmarshal(data, actual.Interface())

	assert.Nil(t, err)
	assert.Equal(t, tv, actual.Elem().Interface())

	// Return the an interface wrapping the roundtripped test vector.
	// In case it's needed for further processing it can be extracted
	// with a type assertion.
	return actual.Elem().Interface()
}
