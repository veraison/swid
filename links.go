package swid

import "reflect"

// Links models link-entry / [ 2* link-entry ]
type Links []Link

// MarshalCBOR provides the custom CBOR marshaler for link entries
func (la Links) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(la))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for link entries
func (la *Links) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Link)(la))
	}

	var l Link

	if err := dm.Unmarshal(data, &l); err != nil {
		return err
	}

	*la = append(*la, l)

	return nil
}
