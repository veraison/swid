// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import "reflect"

// Processes models CoSWID process-entry / [ 2* process-entry ]
type Processes []Process

// MarshalCBOR provides the custom CBOR marshaler for process entries
func (pa Processes) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(pa))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for process entries
func (pa *Processes) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Process)(pa))
	}

	var p Process

	if err := dm.Unmarshal(data, &p); err != nil {
		return err
	}

	*pa = append(*pa, p)

	return nil
}
