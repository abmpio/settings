package settings

import (
	"encoding/json"
	"strconv"

	"github.com/abmpio/mongodbr"
)

type Setting struct {
	mongodbr.AuditedEntity `bson:",inline"`

	NameValue `bson:",inline"`
}

func (s *Setting) Value() string {
	return s.NameValue.Value
}

func (s *Setting) ValueAsBool() bool {
	v := s.NameValue.Value
	if v == "" {
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}

func (s *Setting) ValueAsInt32() int32 {
	v := s.NameValue.Value
	if v == "" {
		return 0
	}
	b, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}
	return int32(b)
}

func (s *Setting) ValueAsInt64() int64 {
	v := s.NameValue.Value
	if v == "" {
		return 0
	}
	b, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return b
}

func (s *Setting) ValueToPtr(v interface{}) error {
	vString := s.NameValue.Value
	if vString == "" {
		return nil
	}
	return json.Unmarshal([]byte(vString), v)
}
