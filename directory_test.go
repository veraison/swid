// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import "testing"

func TestDirectory_RoundtripFlat(t *testing.T) {
	tv := Directory{
		FileSystemItem: testFileSystemItemMinSet,
		PathElements: PathElements{
			Files: &Files{testFileFull},
		},
	}
	/*
		a3                                        # map(3)
			17                                    # unsigned(23)
			70                                    # text(16)
				62696e2f6669726d776172652e62696e  # "bin/firmware.bin"
			18 18                                 # unsigned(24)
			6c                                    # text(12)
				6669726d776172652e62696e          # "firmware.bin"
			18 1a                                 # unsigned(26)
			a1                                    # map(1)
				11                                # unsigned(17)
				a7                                # map(7)
					16                            # unsigned(22)
					f5                            # primitive(21)
					17                            # unsigned(23)
					70                            # text(16)
						62696e2f6669726d776172652e62696e # "bin/firmware.bin"
					18 18                         # unsigned(24)
					6c                            # text(12)
						6669726d776172652e62696e  # "firmware.bin"
					18 19                         # unsigned(25)
					61                            # text(1)
						2f                        # "/"
					14                            # unsigned(20)
					19 0400                       # unsigned(1024)
					15                            # unsigned(21)
					64                            # text(4)
						76312e32                  # "v1.2"
					07                            # unsigned(7)
					82                            # array(2)
						01                        # unsigned(1)
						42                        # bytes(2)
						0001                      # "\x00\x01"
	*/
	expectedCBOR := []byte{
		0xa3, 0x17, 0x70, 0x62, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x72, 0x6d,
		0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18, 0x18, 0x6c,
		0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69,
		0x6e, 0x18, 0x1a, 0xa1, 0x11, 0xa7, 0x16, 0xf5, 0x17, 0x70, 0x62,
		0x69, 0x6e, 0x2f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
		0x2e, 0x62, 0x69, 0x6e, 0x18, 0x18, 0x6c, 0x66, 0x69, 0x72, 0x6d,
		0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18, 0x19, 0x61,
		0x2f, 0x14, 0x19, 0x04, 0x00, 0x15, 0x64, 0x76, 0x31, 0x2e, 0x32,
		0x07, 0x82, 0x01, 0x42, 0x00, 0x01,
	}

	roundTripper(t, tv, expectedCBOR)
}

func TestDirectory_RoundtripNested(t *testing.T) {
	tv := Directory{
		FileSystemItem: testFileSystemItemMinSet,
		PathElements: PathElements{
			Directories: &Directories{
				Directory{
					FileSystemItem: testFileSystemItemMinSet,
					PathElements: PathElements{
						Files: &Files{testFileMinSet},
					},
				},
			},
		},
	}
	/*
		a3                                       # map(3)
			17                                   # unsigned(23)
			70                                   # text(16)
				62696e2f6669726d776172652e62696e # "bin/firmware.bin"
			18 18                                # unsigned(24)
			6c                                   # text(12)
				6669726d776172652e62696e         # "firmware.bin"
			18 1a                                # unsigned(26)
			a1                                   # map(1)
				10                               # unsigned(16)
				a3                               # map(3)
					17                           # unsigned(23)
					70                           # text(16)
						62696e2f6669726d776172652e62696e # "bin/firmware.bin"
					18 18                        # unsigned(24)
					6c                           # text(12)
						6669726d776172652e62696e # "firmware.bin"
					18 1a                        # unsigned(26)
					a1                           # map(1)
						11                       # unsigned(17)
						a2                       # map(2)
						17                       # unsigned(23)
						70                       # text(16)
							62696e2f6669726d776172652e62696e # "bin/firmware.bin"
						18 18                    # unsigned(24)
						6c                       # text(12)
							6669726d776172652e62696e # "firmware.bin"
	*/
	expectedCBOR := []byte{
		0xa3, 0x17, 0x70, 0x62, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x72, 0x6d,
		0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18, 0x18, 0x6c,
		0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69,
		0x6e, 0x18, 0x1a, 0xa1, 0x10, 0xa3, 0x17, 0x70, 0x62, 0x69, 0x6e,
		0x2f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62,
		0x69, 0x6e, 0x18, 0x18, 0x6c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
		0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18, 0x1a, 0xa1, 0x11, 0xa2,
		0x17, 0x70, 0x62, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x72, 0x6d, 0x77,
		0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x18, 0x18, 0x6c, 0x66,
		0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x69, 0x6e,
	}

	roundTripper(t, tv, expectedCBOR)
}
