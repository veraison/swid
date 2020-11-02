package swid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagID_16Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
	}

	expected := tv

	actual, err := checkTagID(tv)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestTagID_15Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
	}

	_, err := checkTagID(tv)

	assert.EqualError(t, err, "binary tag-id MUST be 16 bytes")
}

func TestTagID_17Bytes(t *testing.T) {
	tv := []byte{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00,
	}

	_, err := checkTagID(tv)

	assert.EqualError(t, err, "binary tag-id MUST be 16 bytes")
}

func TestTagID_String(t *testing.T) {
	tv := "example.acme.roadrunner-sw-v1-0-0"

	actual, err := checkTagID(tv)

	assert.Nil(t, err)
	assert.Equal(t, tv, actual)
}

func TestTagID_UnhandledType(t *testing.T) {
	tv := struct {
		a int
		b string
	}{
		a: 1,
		b: "one",
	}

	_, err := checkTagID(tv)

	assert.EqualError(t, err, "tag-id MUST be []byte or string; got struct { a int; b string }")
}
