// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import "reflect"

// Entities models entity-entry / [ 2* entity-entry ]
type Entities []Entity

// MarshalCBOR provides the custom CBOR marshaler for entity entries
func (ea Entities) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(ea))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for entity entries
func (ea *Entities) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Entity)(ea))
	}

	var e Entity

	if err := dm.Unmarshal(data, &e); err != nil {
		return err
	}

	*ea = append(*ea, e)

	return nil
}
