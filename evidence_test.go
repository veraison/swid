// Copyright 2020-2025 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEvidence_Roundtrip(t *testing.T) {
	date, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00Z")

	tv := Evidence{
		Date:     date,
		DeviceID: "BAD809B1-7032-43D9-8F94-BF128E5D061D",
	}
	/*
	   a2                                      # map(2)
	      18 23                                # unsigned(35)
	      c1                                   # tag(1) <- CoSWID wants time = #6.1(number)
	         00                                # unsigned(0)
	      18 24                                # unsigned(36)
	      78 24                                # text(36)
	         42414438303942312d373033322d343344392d384639342d424631323845354430363144 # "BAD809B1-7032-43D9-8F94-BF128E5D061D"
	*/
	expectedCBOR := []byte{
		0xa2, 0x18, 0x23, 0xc1, 0x00, 0x18, 0x24, 0x78, 0x24, 0x42, 0x41,
		0x44, 0x38, 0x30, 0x39, 0x42, 0x31, 0x2d, 0x37, 0x30, 0x33, 0x32,
		0x2d, 0x34, 0x33, 0x44, 0x39, 0x2d, 0x38, 0x46, 0x39, 0x34, 0x2d,
		0x42, 0x46, 0x31, 0x32, 0x38, 0x45, 0x35, 0x44, 0x30, 0x36, 0x31,
		0x44,
	}

	data, err := em.Marshal(tv)

	assert.Nil(t, err)
	t.Logf("CBOR(hex): %x\n", data)
	assert.Equal(t, expectedCBOR, data)

	actual := Evidence{}
	err = dm.Unmarshal(data, &actual)

	assert.Nil(t, err)
	assert.Equal(t, tv.Date.UTC(), actual.Date.UTC()) // compare as UTC
	assert.Equal(t, tv.DeviceID, actual.DeviceID)
}

func TestEvidence_Valid_empty_device_id(t *testing.T) {
	evidence := Evidence{
		DeviceID: "",
		Date:     time.Now(),
	}

	err := evidence.Valid()

	assert.EqualError(t, err, "evidence device-id is empty")
}

func TestEvidence_Valid_zero_date(t *testing.T) {
	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Time{},
	}

	err := evidence.Valid()

	assert.EqualError(t, err, "evidence date is zero")
}

func TestEvidence_Valid_simple_valid(t *testing.T) {
	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Now(),
	}

	err := evidence.Valid()

	assert.NoError(t, err)
}

func TestEvidence_Valid_with_valid_files(t *testing.T) {
	files := Files{
		File{
			FileSystemItem: FileSystemItem{
				FsName: "test.exe",
			},
		},
		File{
			FileSystemItem: FileSystemItem{
				FsName: "config.ini",
			},
		},
	}

	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Now(),
		ResourceCollection: ResourceCollection{
			PathElements: PathElements{
				Files: &files,
			},
		},
	}

	err := evidence.Valid()

	assert.NoError(t, err)
}

func TestEvidence_Valid_with_invalid_files(t *testing.T) {
	files := Files{
		File{
			FileSystemItem: FileSystemItem{
				FsName: "test.exe",
			},
		},
		File{
			FileSystemItem: FileSystemItem{
				FsName: "", // empty fs-name
			},
		},
	}

	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Now(),
		ResourceCollection: ResourceCollection{
			PathElements: PathElements{
				Files: &files,
			},
		},
	}

	err := evidence.Valid()

	assert.EqualError(t, err, "evidence file[1] invalid: file fs-name is empty")
}

func TestEvidence_Valid_with_valid_processes(t *testing.T) {
	processes := Processes{
		Process{
			ProcessName: "test.exe",
		},
		Process{
			ProcessName: "service.exe",
		},
	}

	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Now(),
		ResourceCollection: ResourceCollection{
			Processes: &processes,
		},
	}

	err := evidence.Valid()

	assert.NoError(t, err)
}

func TestEvidence_Valid_with_invalid_processes(t *testing.T) {
	processes := Processes{
		Process{
			ProcessName: "test.exe",
		},
		Process{
			ProcessName: "", // empty process name
		},
	}

	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Now(),
		ResourceCollection: ResourceCollection{
			Processes: &processes,
		},
	}

	err := evidence.Valid()

	assert.EqualError(t, err, "evidence process[1] process-name is empty")
}

func TestEvidence_Valid_with_mixed_valid_resources(t *testing.T) {
	files := Files{
		File{
			FileSystemItem: FileSystemItem{
				FsName: "test.exe",
			},
		},
	}

	processes := Processes{
		Process{
			ProcessName: "service.exe",
		},
	}

	evidence := Evidence{
		DeviceID: "test-device",
		Date:     time.Now(),
		ResourceCollection: ResourceCollection{
			PathElements: PathElements{
				Files: &files,
			},
			Processes: &processes,
		},
	}

	err := evidence.Valid()

	assert.NoError(t, err)
}
