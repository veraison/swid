package swid

import "reflect"

// Directories models directory-entry / [ 2* directory-entry ],
type Directories []Directory

// MarshalCBOR provides the custom CBOR marshaler for directory entries
func (da Directories) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(da))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for directory entries
func (da *Directories) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Directory)(da))
	}

	var d Directory

	if err := dm.Unmarshal(data, &d); err != nil {
		return err
	}

	*da = append(*da, d)

	return nil
}
