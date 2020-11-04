// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// These tests reuse variables defined in file_test.go

func TestFiles_DoNotEncodeZeroLength(t *testing.T) {
	tv := Files{}

	_, err := em.Marshal(tv)
	assert.EqualError(t, err, "array MUST NOT be 0-length")
}

func TestFiles_RoundtripEncodeOne(t *testing.T) {
	tv := Files{testFileFull}
	/*
	   a7                                   # map(7)
	      16                                # unsigned(22)
	      f5                                # primitive(21)
	      17                                # unsigned(23)
	      70                                # text(16)
	         62696e2f6669726d776172652e62696e # "bin/firmware.bin"
	      18 18                             # unsigned(24)
	      6c                                # text(12)
	         6669726d776172652e62696e       # "firmware.bin"
	      18 19                             # unsigned(25)
	      61                                # text(1)
	         2f                             # "/"
	      14                                # unsigned(20)
	      19 0400                           # unsigned(1024)
	      15                                # unsigned(21)
	      64                                # text(4)
	         76312e32                       # "v1.2"
	      07                                # unsigned(7)
	      82                                # array(2)
	         01                             # unsigned(1)
	         42                             # bytes(2)
	            0001                        # "\x00\x01"
	*/
	expectedCBOR := []byte{
		0xa7, 0x16, 0xf5, 0x17, 0x70, 0x62, 0x69, 0x6e, 0x2f, 0x66, 0x69,
		0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18,
		0x18, 0x6c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e,
		0x62, 0x69, 0x6e, 0x18, 0x19, 0x61, 0x2f, 0x14, 0x19, 0x04, 0x00,
		0x15, 0x64, 0x76, 0x31, 0x2e, 0x32, 0x07, 0x82, 0x01, 0x42, 0x00,
		0x01,
	}

	roundTripper(t, tv, expectedCBOR)
}

func TestFiles_RoundtripEncodeMany(t *testing.T) {
	tv := Files{testFileFull, testFileFull}
	/*
	   82                                   # array(2)
	      a7                                # map(7)
	         16                             # unsigned(22)
	         f5                             # primitive(21)
	         17                             # unsigned(23)
	         70                             # text(16)
	            62696e2f6669726d776172652e62696e # "bin/firmware.bin"
	         18 18                          # unsigned(24)
	         6c                             # text(12)
	            6669726d776172652e62696e    # "firmware.bin"
	         18 19                          # unsigned(25)
	         61                             # text(1)
	            2f                          # "/"
	         14                             # unsigned(20)
	         19 0400                        # unsigned(1024)
	         15                             # unsigned(21)
	         64                             # text(4)
	            76312e32                    # "v1.2"
	         07                             # unsigned(7)
	         82                             # array(2)
	            01                          # unsigned(1)
	            42                          # bytes(2)
	               0001                     # "\x00\x01"
	      a7                                # map(7)
	         16                             # unsigned(22)
	         f5                             # primitive(21)
	         17                             # unsigned(23)
	         70                             # text(16)
	            62696e2f6669726d776172652e62696e # "bin/firmware.bin"
	         18 18                          # unsigned(24)
	         6c                             # text(12)
	            6669726d776172652e62696e    # "firmware.bin"
	         18 19                          # unsigned(25)
	         61                             # text(1)
	            2f                          # "/"
	         14                             # unsigned(20)
	         19 0400                        # unsigned(1024)
	         15                             # unsigned(21)
	         64                             # text(4)
	            76312e32                    # "v1.2"
	         07                             # unsigned(7)
	         82                             # array(2)
	            01                          # unsigned(1)
	            42                          # bytes(2)
	               0001                     # "\x00\x01"
	*/
	expectedCBOR := []byte{
		0x82, 0xa7, 0x16, 0xf5, 0x17, 0x70, 0x62, 0x69, 0x6e, 0x2f, 0x66,
		0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e,
		0x18, 0x18, 0x6c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
		0x2e, 0x62, 0x69, 0x6e, 0x18, 0x19, 0x61, 0x2f, 0x14, 0x19, 0x04,
		0x00, 0x15, 0x64, 0x76, 0x31, 0x2e, 0x32, 0x07, 0x82, 0x01, 0x42,
		0x00, 0x01, 0xa7, 0x16, 0xf5, 0x17, 0x70, 0x62, 0x69, 0x6e, 0x2f,
		0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69,
		0x6e, 0x18, 0x18, 0x6c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
		0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18, 0x19, 0x61, 0x2f, 0x14, 0x19,
		0x04, 0x00, 0x15, 0x64, 0x76, 0x31, 0x2e, 0x32, 0x07, 0x82, 0x01,
		0x42, 0x00, 0x01,
	}

	roundTripper(t, tv, expectedCBOR)
}
