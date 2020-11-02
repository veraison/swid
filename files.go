package swid

import "reflect"

// Files models file => file-entry / [ 2* file-entry ]
type Files []File

// MarshalCBOR provides the custom CBOR marshaler for file entries
func (fa Files) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(fa))
}

// UnmarshalCBOR provides the custom CBOR marshaler for file entries
func (fa *Files) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]File)(fa))
	}

	var f File

	if err := dm.Unmarshal(data, &f); err != nil {
		return err
	}

	*fa = append(*fa, f)

	return nil
}
