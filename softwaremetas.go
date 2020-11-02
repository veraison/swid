package swid

import "reflect"

// SoftwareMetas models CoSWID software-meta-entry / [ 2* software-meta-entry ]
type SoftwareMetas []SoftwareMeta

// MarshalCBOR provides the custom CBOR marshaler for software-meta entries
func (sma SoftwareMetas) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(sma))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for software-meta entries
func (sma *SoftwareMetas) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]SoftwareMeta)(sma))
	}

	var sm SoftwareMeta

	if err := dm.Unmarshal(data, &sm); err != nil {
		return err
	}

	*sma = append(*sma, sm)

	return nil
}
