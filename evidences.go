package swid

import "reflect"

// Evidences models evidence => evidence-entry // evidence => [ 2* evidence-entry ]
type Evidences []Evidence

// MarshalCBOR provides the custom CBOR marshaler for evidence entries
func (ea Evidences) MarshalCBOR() ([]byte, error) {
	return arrayToCBOR(reflect.ValueOf(ea))
}

// UnmarshalCBOR provides the custom CBOR unmarshaler for evidence entries
func (ea *Evidences) UnmarshalCBOR(data []byte) error {
	if (data[0] & 0xe0) == 0x80 {
		return dm.Unmarshal(data, (*[]Evidence)(ea))
	}

	var e Evidence

	if err := dm.Unmarshal(data, &e); err != nil {
		return err
	}

	*ea = append(*ea, e)

	return nil
}
