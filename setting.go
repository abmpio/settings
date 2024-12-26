package settings

import (
	"time"

	"github.com/abmpio/mongodbr"
)

type ValueFieldType string

const (
	// string
	ValueFieldType_String = "string"
	// boolean
	ValueFieldType_Boolean = "boolean"
	// float64
	ValueFieldType_Float64 = "float64"
	// time.Time
	ValueFieldType_DateTime = "dateTime"
)

type Setting struct {
	mongodbr.AuditedEntity `bson:",inline"`

	AppName     string         `json:"appName" bson:"appName"`
	ValueType   ValueFieldType `json:"valueType" bson:"valueType"`
	Tag         string         `json:"tag" bson:"tag"`
	Description string         `json:"description" bson:"description"`
	//当值为true时,表示此值是一个受保护的属性值,服务器不会将数据发送到前端
	ProtectedInUi bool `json:"protectedInUi" bson:"protectedInUi"`
	NameValue     `bson:",inline"`
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
