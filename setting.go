package settings

import (
	"fmt"
	"time"

	"github.com/abmpio/libx/lang"
	"github.com/abmpio/mongodbr"
)

type ValueFieldType string

const (
	// string
	ValueFieldType_String ValueFieldType = "string"
	// boolean
	ValueFieldType_Boolean ValueFieldType = "boolean"
	// float64
	ValueFieldType_Float64 ValueFieldType = "float64"
	// time.Time
	ValueFieldType_DateTime ValueFieldType = "dateTime"
)

type Setting struct {
	mongodbr.AuditedEntity `bson:",inline"`

	AppName     string         `json:"appName" bson:"appName"`
	ValueType   ValueFieldType `json:"valueType" bson:"valueType"`
	Tag         string         `json:"tag" bson:"tag"`
	Description string         `json:"description" bson:"description"`
	//当值为true时,表示此值是一个受保护的属性值,服务器不会将数据发送到前端
	ProtectedInUi  bool `json:"protectedInUi" bson:"protectedInUi"`
	lang.NameValue `bson:",inline"`
}

// NormalizeValue ensures that the NameValue.Value matches the specified ValueType.
// If the value does not match the expected type, it sets a default value based on the ValueType.
// For ValueFieldType_DateTime, it attempts to parse the string representation into a time.Time object.
// Returns an error if the value cannot be normalized to the expected type.
func (s *Setting) NormalizeValue() error {
	switch s.ValueType {
	case ValueFieldType_DateTime:
		_, ok := s.NameValue.Value.(time.Time)
		if ok {
			return nil
		}
		v, err := time.Parse(time.RFC3339, s.ValueAsString())
		if err != nil {
			return fmt.Errorf("无效的时间值,%s", err.Error())
		}
		s.NameValue.Value = v
	case ValueFieldType_Boolean:
		_, ok := s.NameValue.Value.(bool)
		if ok {
			return nil
		}
		s.NameValue.Value = false
	case ValueFieldType_Float64:
		_, ok := s.NameValue.Value.(float64)
		if ok {
			return nil
		}
		s.NameValue.Value = float64(0)
	case ValueFieldType_String:
		_, ok := s.NameValue.Value.(string)
		if ok {
			return nil
		}
		s.NameValue.Value = ""
	}
	return nil
}

func (s *Setting) Value() interface{} {
	return s.NameValue.Value
}

func (s *Setting) ValueAsString() string {
	if s.NameValue.Value == nil {
		return ""
	}
	value, ok := s.NameValue.Value.(string)
	if !ok {
		return ""
	}
	return value
}

// if Value is nil,return false
func (s *Setting) ValueAsBoolean() bool {
	if s.NameValue.Value == nil {
		return false
	}
	value, ok := s.NameValue.Value.(bool)
	if !ok {
		return false
	}
	return value
}

// if Value is nil,return 0
func (s *Setting) ValueAsFloat64() float64 {
	if s.NameValue.Value == nil {
		return 0
	}
	value, ok := s.NameValue.Value.(float64)
	if !ok {
		return 0
	}
	return value
}

// if Value is nil,return nil
func (s *Setting) ValueAsDateTime() *time.Time {
	if s.NameValue.Value == nil {
		return nil
	}
	value, ok := s.NameValue.Value.(time.Time)
	if !ok {
		return nil
	}
	return &value

}

// ValueIsMatchType checks if the given value matches the specified ValueFieldType
// Returns true if the value matches the type, false otherwise
// For ValueFieldType_DateTime, nil is considered a valid value
// For ValueFieldType_DateTime, both time.Time and *time.Time (non-nil) are considered valid types
// For other types, only the exact type match is considered valid
// Example:
//
//	ValueIsMatchType("example", ValueFieldType_String) => true
//	ValueIsMatchType(123, ValueFieldType_String) => false
//	ValueIsMatchType(true, ValueFieldType_Boolean) => true
//	ValueIsMatchType(3.14, ValueFieldType_Float64) => true
//	ValueIsMatchType(nil, ValueFieldType_DateTime) => true
//	ValueIsMatchType(time.Now(), ValueFieldType_DateTime) => true
//	ValueIsMatchType(&time.Time{}, ValueFieldType_DateTime) => true
//	ValueIsMatchType(123, ValueFieldType_DateTime) => false
func ValueIsMatchType(value interface{}, valueType ValueFieldType) bool {
	switch valueType {
	case ValueFieldType_String:
		_, ok := value.(string)
		return ok
	case ValueFieldType_Boolean:
		_, ok := value.(bool)
		return ok
	case ValueFieldType_Float64:
		_, ok := value.(float64)
		return ok
	case ValueFieldType_DateTime:
		if value == nil {
			// nil
			return true
		}
		switch v := value.(type) {
		case time.Time:
			return true
		case *time.Time:
			// non-nil pointer
			return v != nil
		default:
			return false
		}
	default:
		return false
	}
}
