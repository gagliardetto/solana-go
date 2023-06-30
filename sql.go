package solana

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"fmt"
)

// This file defines Value and Scan methods for a bunch of the exported types
// so they can be written to/from a DB.

func textValue[T encoding.TextMarshaler](v T) (driver.Value, error) {
	bytes, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

func scanText[T encoding.TextUnmarshaler](v T, src any) error {
	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("%T: cannot scan, expected string got '%T'", v, src)
	}
	if err := v.UnmarshalText([]byte(s)); err != nil {
		return fmt.Errorf("%T: failed to scan value '%s': %w", v, s, err)
	}
	return nil
}

type NullPublicKey struct {
	PublicKey
	Valid bool
}

func (v NullPublicKey) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.PublicKey.Value()
}

func (v *NullPublicKey) Scan(src any) error {
	if src == nil {
		return nil
	}
	if err := v.PublicKey.Scan(src); err != nil {
		return nil
	}
	v.Valid = true
	return nil
}

var (
	_ driver.Valuer = PublicKey{}
	_ sql.Scanner   = &PublicKey{}
)

func (v PublicKey) Value() (driver.Value, error) {
	return textValue(v)
}

func (v *PublicKey) Scan(src interface{}) error {
	return scanText(v, src)
}

type NullSignature struct {
	Signature
	Valid bool
}

func (v NullSignature) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Signature.Value()
}

func (v *NullSignature) Scan(src any) error {
	if src == nil {
		return nil
	}
	if err := v.Signature.Scan(src); err != nil {
		return nil
	}
	v.Valid = true
	return nil
}

var (
	_ driver.Valuer = Signature{}
	_ sql.Scanner   = &Signature{}
)

func (v Signature) Value() (driver.Value, error) {
	return textValue(v)
}

func (v *Signature) Scan(src interface{}) error {
	return scanText(v, src)
}

type NullHash struct {
	Hash
	Valid bool
}

func (v NullHash) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Hash.Value()
}

func (v *NullHash) Scan(src any) error {
	if src == nil {
		return nil
	}
	if err := v.Hash.Scan(src); err != nil {
		return nil
	}
	v.Valid = true
	return nil
}

var (
	_ driver.Valuer = Hash{}
	_ sql.Scanner   = &Hash{}
)

func (v Hash) Value() (driver.Value, error) {
	return textValue(v)
}

func (v *Hash) Scan(src interface{}) error {
	return scanText(v, src)
}
