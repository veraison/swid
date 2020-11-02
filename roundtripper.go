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
	// don't allow TZ info to get lost in the encoding / decoding process
	// em, err := cbor.EncOptions{Time: cbor.TimeRFC3339, TimeTag: cbor.EncTagRequired}.EncMode()
	em, err := cbor.EncOptions{TimeTag: cbor.EncTagRequired}.EncMode()
	require.Nil(t, err)

	data, err := em.Marshal(tv)

	assert.Nil(t, err)
	t.Logf("CBOR(hex): %x\n", data)
	assert.Equal(t, expectedCBOR, data)

	dm, err := cbor.DecOptions{TimeTag: cbor.DecTagOptional}.DecMode()
	require.Nil(t, err)

	actual := reflect.New(reflect.TypeOf(tv))
	err = dm.Unmarshal(data, actual.Interface())

	assert.Nil(t, err)
	assert.Equal(t, tv, actual.Elem().Interface())

	// Return the an interface wrapping the roundtripped test vector.
	// In case it's needed for further processing it can be extracted
	// with a type assertion.
	return actual.Elem().Interface()
}
