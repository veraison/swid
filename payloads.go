// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import "reflect"

// Payloads models CoSWID payload-entry / [ 2* payload-entry ]
type Payloads []Payload

// MarshalCBOR provides the custom CBOR marshaler for payload entries
func (pa Payloads) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(pa))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for payload entries
func (pa *Payloads) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Payload)(pa))
	}

	var p Payload

	if err := dm.Unmarshal(data, &p); err != nil {
		return err
	}

	*pa = append(*pa, p)

	return nil
}
