package settings

import (
	"github.com/abmpio/mongodbr"
)

type Setting struct {
	mongodbr.AuditedEntity `bson:",inline"`

	NameValue `bson:",inline"`
}

func (s *Setting) Value() interface{} {
	return s.NameValue.Value
}
