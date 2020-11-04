// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// HashEntry models
type HashEntry struct {
	_ struct{} `cbor:",toarray"`

	// The number used as a value for HashAlgID MUST refer an ID in the IANA
	// "Name Information Hash Algorithm Registry". Other hash algorithms MUST
	// NOT be used.
	HashAlgID uint64

	// The HashValue MUST represent the raw hash value of the hashed resource
	// generated using the hash algorithm indicated by the HashAlgID
	HashValue []byte
}

// Named Information Hash Algorithm Registry
// https://www.iana.org/assignments/named-information/named-information.xhtml#hash-alg
const (
	sha256 = uint64(iota + 1)
	sha256_128
	sha256_120
	sha256_96
	sha256_64
	sha256_32
	sha384
	sha512
	sha3_224
	sha3_256
	sha3_384
	sha3_512
)

var (
	algToString = map[uint64]string{
		sha256:     "sha-256",
		sha256_128: "sha-256-128",
		sha256_120: "sha-256-120",
		sha256_96:  "sha-256-96",
		sha256_64:  "sha-256-64",
		sha256_32:  "sha-256-32",
		sha384:     "sha-384",
		sha512:     "sha-512",
		sha3_224:   "sha3-224",
		sha3_256:   "sha3-256",
		sha3_384:   "sha3-384",
		sha3_512:   "sha3-512",
	}

	stringToAlg = map[string]uint64{
		"sha-256":     sha256,
		"sha-256-128": sha256_128,
		"sha-256-120": sha256_120,
		"sha-256-96":  sha256_96,
		"sha-256-64":  sha256_64,
		"sha-256-32":  sha256_32,
		"sha-384":     sha384,
		"sha-512":     sha512,
		"sha3-224":    sha3_224,
		"sha3-256":    sha3_256,
		"sha3-384":    sha3_384,
		"sha3-512":    sha3_512,
	}
)

// Set assigns the supplied algID and hash value to the HashEntry receiver
func (h *HashEntry) Set(algID uint64, value []byte) error {
	h.HashAlgID = algID
	h.HashValue = value
	return nil
}

func (h HashEntry) stringify() (string, error) {
	sAlg, ok := algToString[h.HashAlgID]
	if !ok {
		return "", fmt.Errorf("unknown hash algorithm ID %d", h.HashAlgID)
	}

	//sVal := hex.EncodeToString(h.HashValue)
	sVal := base64.StdEncoding.EncodeToString(h.HashValue)
	if len(sVal) == 0 {
		return "", fmt.Errorf("empty hash value")
	}

	s := sAlg + ":" + sVal

	return s, nil
}

func (h *HashEntry) codify(v string) error {
	// expected format is <hash-alg-string>:<hash-value>
	s := strings.Split(v, ":")

	if len(s) != 2 {
		return fmt.Errorf("bad format: expecting <hash-alg-string>:<hash-value>")
	}

	sAlg := strings.TrimSpace(s[0])
	sVal := strings.TrimSpace(s[1])

	if sAlg == "" || sVal == "" {
		return fmt.Errorf("bad format: expecting <hash-alg-string>:<hash-value>")
	}

	algID, ok := stringToAlg[strings.ToLower(sAlg)]
	if !ok {
		return fmt.Errorf("unknown hash algorithm %s", sAlg)
	}

	//value, err := hex.DecodeString(sVal)
	value, err := base64.StdEncoding.DecodeString(sVal)
	if err != nil {
		return err
	}

	h.HashAlgID = algID
	h.HashValue = value

	return nil
}

// MarshalJSON provides the custom JSON marshaler for the HashEntry type
func (h HashEntry) MarshalJSON() ([]byte, error) {
	s, err := h.stringify()
	if err != nil {
		return nil, err
	}
	return json.Marshal(&s)
}

// UnmarshalJSON provides the custom JSON unmarshaler for the HashEntry type
func (h *HashEntry) UnmarshalJSON(data []byte) error {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch t := v.(type) {
	case string:
		return h.codify(t)
	default:
		return fmt.Errorf("expecting string, found %T instead", t)
	}
}

// MarshalXMLAttr provides the custom XML attribute marshaler for the HashEntry type
func (h HashEntry) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s, err := h.stringify()
	if err != nil {
		return xml.Attr{}, err
	}

	return xml.Attr{Name: name, Value: s}, nil
}

// UnmarshalXMLAttr provides the custom XML attribute unmarshaler for the HashEntry type
func (h *HashEntry) UnmarshalXMLAttr(attr xml.Attr) error {
	return h.codify(attr.Value)
}
