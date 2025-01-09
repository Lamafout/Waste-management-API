package model

type Fkko struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

func NewFkkoFromMap(m map[string]interface{}) *Fkko {
	return &Fkko{
		Name: m["name"].(string),
		Code: m["code"].(string),
	}
}