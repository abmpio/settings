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
