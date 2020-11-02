package swid

import (
	"encoding/xml"
	"testing"

	cbor "github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
)

func TestGlobalAttributes_MixingExtensionsEncode(t *testing.T) {
	type DummyExtension struct {
		FooFoo string `cbor:"32768,keyasint" xml:"http://localhost/tns foofoo,attr"`
	}

	type X struct {
		GlobalAttributes
		DummyExtension
	}

	x := X{
		GlobalAttributes{
			Lang: "babel",
		},
		DummyExtension{
			FooFoo: "bing",
		},
	}

	expected := `<X xml:lang="babel" xmlns:tns="http://localhost/tns" tns:foofoo="bing"></X>`

	data, err := xml.Marshal(x)

	assert.Nil(t, err)
	assert.Equal(t, expected, string(data))

	data, err = cbor.Marshal(x)
	assert.Nil(t, err)
	t.Logf("%x", data)
	assert.Equal(t,
		[]byte{
			/*
			   a2               # map(2)
			      0f            # unsigned(15)
			      65            # text(5)
			         626162656c # "babel"
			      19 8000       # unsigned(32768)
			      64            # text(4)
			   	  62696e67   # "bing"
			*/
			0xa2, 0x0f, 0x65, 0x62, 0x61, 0x62, 0x65, 0x6c, 0x19, 0x80, 0x00,
			0x64, 0x62, 0x69, 0x6e, 0x67,
		},
		data,
	)
}
