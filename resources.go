package swid

import "reflect"

// Resources models CoSWID resource-entry / [ 2* resource-entry ]
type Resources []Resource

// MarshalCBOR provides the custom CBOR marshaler for resource entries
func (ra Resources) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(ra))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for resource entries
func (ra *Resources) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Resource)(ra))
	}

	var r Resource

	if err := dm.Unmarshal(data, &r); err != nil {
		return err
	}

	*ra = append(*ra, r)

	return nil
}
