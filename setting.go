package settings

import (
	"github.com/abmpio/mongodbr"
)

type ValueFieldType string

type Setting struct {
	mongodbr.AuditedEntity `bson:",inline"`

	ValueType ValueFieldType `json:"valueType" bson:"valueType"`
	NameValue `bson:",inline"`
}

func (s *Setting) Value() interface{} {
	return s.NameValue.Value
}

func (s *Setting) ValueAsString() string {
	if s.NameValue.Value == nil {
		return ""
	}
	stringValue, ok := s.NameValue.Value.(string)
	if ok {
		return stringValue
	}
	return ""
}
