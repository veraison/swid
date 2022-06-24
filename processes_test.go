// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcesses_MarshalCBOR(t *testing.T) {
	tests := []struct {
		name        string
		tv          Processes
		expected    []byte
		expectedErr error
	}{
		{
			name:        "zero elements",
			tv:          Processes{},
			expected:    []byte("__ignored__"),
			expectedErr: errors.New("array of Processes MUST NOT be 0-length"),
		},
		{
			name: "one element (encodes to scalar)",
			tv: Processes{
				Process{ProcessName: "a"},
			},
			/*
			   a1       # map(1)
			      18 1b # unsigned(27)
			      61    # text(1)
			         61 # "a"
			*/
			expected: []byte{
				0xa1, 0x18, 0x1b, 0x61, 0x61,
			},
			expectedErr: nil,
		},
		{
			name: "two elements (encodes to array)",
			tv: Processes{
				Process{ProcessName: "a"},
				Process{ProcessName: "b"},
			},
			/*
			   82          # array(2)
			      a1       # map(1)
			         18 1b # unsigned(27)
			         61    # text(1)
			            61 # "a"
			      a1       # map(1)
			         18 1b # unsigned(27)
			         61    # text(1)
			            62 # "b"
			*/
			expected: []byte{
				0x82, 0xa1, 0x18, 0x1b, 0x61, 0x61, 0xa1, 0x18, 0x1b, 0x61,
				0x62,
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.tv.MarshalCBOR()
			t.Logf("hex: %x\n", actual)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestProcesses_UnmarshalCBOR(t *testing.T) {
	tests := []struct {
		name        string
		tv          []byte
		expected    Processes
		expectedErr error
	}{
		{
			name:        "zero elements array",
			tv:          []byte{0x80},
			expected:    Processes{},
			expectedErr: nil,
		},
		{
			name: "accept array with one element (non optimized)",
			/*
			   81          # array(1)
			      a1       # map(1)
			         18 1b # unsigned(27)
			         61    # text(1)
			            61 # "a"
			*/
			tv: []byte{
				0x81, 0xa1, 0x18, 0x1b, 0x61, 0x61,
			},
			expected: Processes{
				Process{ProcessName: "a"},
			},
			expectedErr: nil,
		},
		{
			name: "array with two elements",
			/*
			   82          # array(2)
			      a1       # map(1)
			         18 1b # unsigned(27)
			         61    # text(1)
			            61 # "a"
			      a1       # map(1)
			         18 1b # unsigned(27)
			         61    # text(1)
			            62 # "b"
			*/
			tv: []byte{
				0x82, 0xa1, 0x18, 0x1b, 0x61, 0x61, 0xa1, 0x18, 0x1b, 0x61,
				0x62,
			},
			expected: Processes{
				Process{ProcessName: "a"},
				Process{ProcessName: "b"},
			},
			expectedErr: nil,
		},
		{
			name: "one scalar element (optimized encoding)",
			/*
			   a1       # map(1)
			      18 1b # unsigned(27)
			      61    # text(1)
			         61 # "a"
			*/
			tv: []byte{
				0xa1, 0x18, 0x1b, 0x61, 0x61,
			},
			expected: Processes{
				Process{ProcessName: "a"},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Processes{}
			err := actual.UnmarshalCBOR(test.tv)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func benchmarkProcessesMarshal(i int, b *testing.B) {
	ps := makeProcesses(i)

	for n := 0; n < b.N; n++ {
		if _, e := ps.MarshalCBOR(); e != nil {
			b.Fatalf("marshaling failed: %v", e)
		}
	}
}

func BenchmarkProcesses_Marshal10(b *testing.B)    { benchmarkProcessesMarshal(10, b) }
func BenchmarkProcesses_Marshal100(b *testing.B)   { benchmarkProcessesMarshal(100, b) }
func BenchmarkProcesses_Marshal1000(b *testing.B)  { benchmarkProcessesMarshal(1000, b) }
func BenchmarkProcesses_Marshal10000(b *testing.B) { benchmarkProcessesMarshal(10000, b) }

func benchmarkProcessesUnmarshal(i int, b *testing.B) {
	ps := makeProcesses(i)
	tv, e := ps.MarshalCBOR()
	if e != nil {
		b.Fatal("creating test vector")
	}

	for n := 0; n < b.N; n++ {
		if e := ps.UnmarshalCBOR(tv); e != nil {
			b.Fatalf("unmarshalling failed: %v", e)
		}
	}
}

func BenchmarkProcesses_Unmarshal10(b *testing.B)    { benchmarkProcessesUnmarshal(10, b) }
func BenchmarkProcesses_Unmarshal100(b *testing.B)   { benchmarkProcessesUnmarshal(100, b) }
func BenchmarkProcesses_Unmarshal1000(b *testing.B)  { benchmarkProcessesUnmarshal(1000, b) }
func BenchmarkProcesses_Unmarshal10000(b *testing.B) { benchmarkProcessesUnmarshal(10000, b) }

func makeProcesses(nProcs int) *Processes {
	ps := Processes{}

	for i := 0; i < nProcs; i++ {
		pID := i
		pName := fmt.Sprintf("process-%d", i)
		ps = append(ps, Process{ProcessName: pName, Pid: &pID})
	}

	return &ps

}
