package settings

type NameValue struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}
