package eth

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type Data string
type Data8 Data
type Data20 Data
type Data32 Data
type Data256 Data

// Aliases
type Hash = Data32
type Topic = Data32

func NewData(value string) (*Data, error) {
	parsed, err := parseHex(value, 0, "data")
	if err != nil {
		return nil, err
	}

	d := Data(parsed)
	return &d, nil
}

func NewData8(value string) (*Data8, error) {
	parsed, err := parseHex(value, 8, "data")
	if err != nil {
		return nil, err
	}

	d := Data8(parsed)
	return &d, nil
}

func NewData20(value string) (*Data20, error) {
	parsed, err := parseHex(value, 20, "data")
	if err != nil {
		return nil, err
	}

	d := Data20(parsed)
	return &d, nil
}

func NewData32(value string) (*Data32, error) {
	parsed, err := parseHex(value, 32, "data")
	if err != nil {
		return nil, err
	}

	d := Data32(parsed)
	return &d, nil
}

func NewHash(value string) (*Hash, error) {
	return NewData32(value)
}

func NewTopic(value string) (*Hash, error) {
	return NewData32(value)
}

func NewData256(value string) (*Data256, error) {
	parsed, err := parseHex(value, 256, "data")
	if err != nil {
		return nil, err
	}

	d := Data256(parsed)
	return &d, nil
}

func MustData(value string) *Data {
	d, err := NewData(value)
	if err != nil {
		panic(err)
	}

	return d
}

func MustData8(value string) *Data8 {
	d, err := NewData8(value)
	if err != nil {
		panic(err)
	}

	return d
}

func MustData20(value string) *Data20 {
	d, err := NewData20(value)
	if err != nil {
		panic(err)
	}

	return d
}

func MustData32(value string) *Data32 {
	d, err := NewData32(value)
	if err != nil {
		panic(err)
	}

	return d
}

func MustHash(value string) *Hash {
	return MustData32(value)
}

func MustTopic(value string) *Hash {
	return MustData32(value)
}

func MustData256(value string) *Data256 {
	d, err := NewData256(value)
	if err != nil {
		panic(err)
	}

	return d
}

func (d Data) String() string {
	return string(d)
}

func (d Data8) String() string {
	return string(d)
}

func (d Data20) String() string {
	return string(d)
}

func (d Data32) String() string {
	return string(d)
}

func (d Data256) String() string {
	return string(d)
}

func (d *Data) UnmarshalJSON(data []byte) error {
	str, err := unmarshalHex(data, 0, "data")
	if err != nil {
		return err
	}
	*d = Data(str)
	return nil
}

func (d *Data8) UnmarshalJSON(data []byte) error {
	str, err := unmarshalHex(data, 8, "data")
	if err != nil {
		return err
	}
	*d = Data8(str)
	return nil
}

func (d *Data20) UnmarshalJSON(data []byte) error {
	str, err := unmarshalHex(data, 20, "data")
	if err != nil {
		return err
	}
	*d = Data20(str)
	return nil
}

func (d *Data32) UnmarshalJSON(data []byte) error {
	str, err := unmarshalHex(data, 32, "data")
	if err != nil {
		return err
	}
	*d = Data32(str)
	return nil
}

func (d *Data256) UnmarshalJSON(data []byte) error {
	str, err := unmarshalHex(data, 256, "data")
	if err != nil {
		return err
	}
	*d = Data256(str)
	return nil
}

func (d *Data) MarshalJSON() ([]byte, error) {
	s := string(*d)
	return json.Marshal(&s)
}

func unmarshalHex(data []byte, size int, typ string) (string, error) {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return "", err
	}

	return parseHex(str, size, typ)
}

func parseHex(value string, size int, typ string) (string, error) {
	if !strings.HasPrefix(value, "0x") {
		return "", errors.Errorf("%s types must start with 0x", typ)
	}

	if size != 0 {
		dataSize := (len(value) - 2) / 2

		if size != dataSize {
			return "", errors.Errorf("%s type size mismatch, expected %d got %d", typ, size, dataSize)
		}
	}

	return value, nil
}