package settings

type NameValue struct {
	Name  string      `json:"name" bson:"name"`
	Value interface{} `json:"value" bson:"value"`
}
