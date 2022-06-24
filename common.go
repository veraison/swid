// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type codeDictionary map[int64]string

type stringDictionary map[string]int64

func stringifyCode(v *interface{}, dict codeDictionary, codeName string) error {
	switch t := (*v).(type) {
	case string:
		*v = t
		return nil
	// uint64 is only for compatibility with code that uses this library and used
	// the uint64 type before (as was default for CoSWID spec v16 and earlier)
	case uint64:
		if t > math.MaxInt64 {
			return fmt.Errorf("%s should never be above max of int64", codeName)
		}
		if s, ok := dict[int64(t)]; ok {
			*v = s
		} else if codeName != "" {
			*v = fmt.Sprintf("%s(%d)", codeName, t)
		}
		return nil
	case int64:
		if s, ok := dict[t]; ok {
			*v = s
		} else if codeName != "" {
			*v = fmt.Sprintf("%s(%d)", codeName, t)
		}
		return nil
	default:
		return fmt.Errorf("unhandled type: %T", t)
	}
}

func codifyString(v *interface{}, dict stringDictionary) error {
	switch t := (*v).(type) {
	case string:
		// try mapping to code and replace if successful
		if ui, ok := dict[t]; ok {
			*v = ui
		}
		return nil
	// CBOR library returns uint64 type for positive integers, but we need int64
	// for all types since CoSWID Spec v17
	case uint64:
		if t > math.MaxInt64 {
			return fmt.Errorf("there are no dictionary values above max of int64 in CoSWID")
		}
		*v = int64(t)
		return nil
	case int64:
		return nil
	case float64:
		// check that the JSON number is integer (i.e., no fraction / exponent)
		// if so, convert and replace
		if t == float64(int64(t)) {
			*v = int64(t)
			return nil
		}
		return fmt.Errorf("number %s is not int64", strconv.FormatFloat(t, 'f', -1, 64))
	default:
		return fmt.Errorf("unhandled type: %T", t)
	}
}

func isStringOrCode(v interface{}, codeName string) error {
	switch t := v.(type) {
	// uint64 is only for compatibility with code that uses this library and used
	// the uint64 type before (as was default for CoSWID spec v16 and earlier)
	case uint64:
		return nil
	case int64:
		return nil
	case string:
		return nil
	default:
		return fmt.Errorf("%s MUST be int64 or string; got %T", codeName, t)
	}
}

func codeStringer(code interface{}, dict codeDictionary, codeName string) string {
	v := code

	if err := stringifyCode(&v, dict, codeName); err != nil {
		return ""
	}

	return v.(string)
}

func codeToCBOR(code interface{}, dict stringDictionary) ([]byte, error) {
	v := code

	// always try to minimize bandwidth
	if err := codifyString(&v, dict); err != nil {
		return nil, err
	}

	return em.Marshal(v)
}

func codeToJSON(code interface{}, dict codeDictionary) ([]byte, error) {
	v := code // make a copy we can clobber

	// always try to maximize expressiveness
	// however, avoid encoding unknown codes
	if err := stringifyCode(&v, dict, ""); err != nil {
		return nil, err
	}

	return json.Marshal(v)
}

func codeToXMLAttr(attrName xml.Name, code interface{}, dict codeDictionary) (xml.Attr, error) {
	v := code // make a copy we can clobber

	// always try to maximize expressiveness
	// however, avoid encoding unknown codes
	if err := stringifyCode(&v, dict, ""); err != nil {
		return xml.Attr{}, err
	}

	return xml.Attr{Name: attrName, Value: v.(string)}, nil
}

type encoder func([]byte, interface{}) error

func xToCode(enc encoder, from []byte, dict stringDictionary, to *interface{}) error {
	if err := enc(from, to); err != nil {
		return err
	}

	// try to make internal representation as homogeneous as possible
	if err := codifyString(to, dict); err != nil {
		return err
	}

	return nil
}

func cborToCode(from []byte, dict stringDictionary, to *interface{}) error {
	return xToCode(dm.Unmarshal, from, dict, to)
}

func jsonToCode(from []byte, dict stringDictionary, to *interface{}) error {
	return xToCode(json.Unmarshal, from, dict, to)
}

func xmlAttrToCode(from xml.Attr, dict stringDictionary, to *interface{}) error {
	*to = from.Value

	if err := codifyString(to, dict); err != nil {
		return err
	}

	return nil
}

func arrayToCBOR(a reflect.Value) ([]byte, error) {
	switch a.Kind() {
	case reflect.Array, reflect.Slice:
	default:
		return nil, errors.New("expecting array or slice")
	}

	switch a.Len() {
	case 0:
		return nil, fmt.Errorf("array of %s MUST NOT be 0-length", a.Type().Name())
	case 1:
		return em.Marshal(a.Index(0).Interface())
	default:
		return em.Marshal(
			// this slight contortion is done to handle conversion from e.g.,
			// Processes to []Process. it is needed to steer the (e.g.,)
			// Processes marshaler to the array of Process marshaler -- if we
			// don't do that, we'd re-enter this same function and be trapped
			// in here until the stack blows up.
			a.Convert(
				reflect.SliceOf(
					a.Index(0).Type(),
				),
			).Interface(),
		)
	}
}
