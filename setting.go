package settings

import (
	"github.com/abmpio/mongodbr"
)

type ValueFieldType string

const (
	ValueFieldType_String  = "string"
	ValueFieldType_Boolean = "boolean"
	ValueFieldType_Int     = "int"
)

type Setting struct {
	mongodbr.AuditedEntity `bson:",inline"`

	AppName   string         `json:"appName" bson:"appName"`
	ValueType ValueFieldType `json:"valueType" bson:"valueType"`
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
	stringValue, ok := s.NameValue.Value.(string)
	if ok {
		return stringValue
	}
	return ""
}
