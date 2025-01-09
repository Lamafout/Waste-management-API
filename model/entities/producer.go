package model

type Producer struct {
	Location    string `json:"location" bson:"location"`
	Name        string `json:"name" bson:"name"`
	Fkko        *Fkko  `json:"fkko" bson:"fkko"`
	HazardClass string `json:"hazardClass" bson:"hazardClass"`
}

func NewProducerFromMap(m map[string]interface{}) *Producer {
	return &Producer{
		Location: m["location"].(string),
		Name:     m["name"].(string),
		Fkko:     NewFkkoFromMap(m["fkko"].(map[string]interface{})),
	}
}
