package model

type Okpd struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

func NewOkpdFromMap(m map[string]interface{}) *Okpd {
	return &Okpd{
		Name: m["name"].(string),
		Code: m["code"].(string),
	}
}