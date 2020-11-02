package swid

import (
	"errors"
	"fmt"
)

// TagID is the type of a tag identifier. Allowed formats (enforced via
// checkTagID) are string or [16]byte
type TagID interface{}

func checkTagID(v interface{}) (TagID, error) {
	switch t := v.(type) {
	case string:
	case []byte:
		if len(t) != 16 {
			return nil, errors.New("binary tag-id MUST be 16 bytes")
		}
	default:
		return nil, fmt.Errorf("tag-id MUST be []byte or string; got %T", v)
	}

	return TagID(v), nil
}
