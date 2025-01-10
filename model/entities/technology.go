package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Technology struct {
	Id             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Assignment     string             `json:"assignment" bson:"assignment"`
	Characteristic string             `json:"characteristic" bson:"characteristic"`
	Resources      *Resources         `json:"resources" bson:"resources"`
	Fkko           []*Fkko            `json:"fkko" bson:"fkko"`
	Okpd           []*Okpd            `json:"okpd" bson:"okpd"`
	Performance    float64            `json:"performance" bson:"performance"`
	SecondaryWaste []*SecondaryWaste  `json:"secondaryWaste" bson:"secondaryWaste"`
	Developer      *Contacts          `json:"developer" bson:"developer"`
	Users          []*Contacts        `json:"users" bson:"users"`
	UseCase        string             `json:"useCase" bson:"useCase"`
	ExpertInfo     *ExpertInfo        `json:"expertInfo" bson:"expertInfo"`
}

func NewTechnologyFromMap(m map[string]interface{}) *Technology {
	technology := &Technology{
		Id:             m["_id"].(primitive.ObjectID),
		Name:           m["name"].(string),
		Assignment:     m["assignment"].(string),
		Characteristic: m["characteristic"].(string),
		Resources:      NewResourcesFromMap(m["resources"].(map[string]interface{})),
		// fkko will be added later
		// okpd will be added later
		Performance: m["performance"].(float64),
		// secondaryWaste will be added later
		Developer: NewContactsFromMap(m["developer"].(map[string]interface{})),
		// users will be added later
		UseCase:    m["useCase"].(string),
		ExpertInfo: NewExpertInfoFromMap(m["expertInfo"].(map[string]interface{})),
	}

	for _, fkko := range m["fkko"].(primitive.A) {
		technology.Fkko = append(technology.Fkko, NewFkkoFromMap(fkko.(map[string]interface{})))
	}

	for _, okpd := range m["okpd"].(primitive.A) {
		technology.Okpd = append(technology.Okpd, NewOkpdFromMap(okpd.(map[string]interface{})))
	}

	for _, secondaryWaste := range m["secondaryWaste"].(primitive.A) {
		technology.SecondaryWaste = append(technology.SecondaryWaste, NewSecondaryWasteFromMap(secondaryWaste.(map[string]interface{})))
	}

	for _, user := range m["users"].(primitive.A) {
		technology.Users = append(technology.Users, NewContactsFromMap(user.(map[string]interface{})))
	}

	return technology
}

type Resources struct {
	Energy       float64 `json:"energy"`
	Water        float64 `json:"water"`
	UsingPerYear float64 `json:"usingPerYear"`
}

func NewResourcesFromMap(m map[string]interface{}) *Resources {
	return &Resources{
		Energy:       m["energy"].(float64),
		Water:        m["water"].(float64),
		UsingPerYear: m["usingPerYear"].(float64),
	}
}

type SecondaryWaste struct {
	Fkko   *Fkko   `json:"fkko"`
	Mass   float64 `json:"mass"`
	Volume float64 `json:"volume"`
}

func NewSecondaryWasteFromMap(m map[string]interface{}) *SecondaryWaste {
	return &SecondaryWaste{
		Fkko:   NewFkkoFromMap(m["fkko"].(map[string]interface{})),
		Mass:   m["mass"].(float64),
		Volume: m["volume"].(float64),
	}
}

type Contacts struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Fax     string `json:"fax"`
	Site    string `json:"site"`
}

func NewContactsFromMap(m map[string]interface{}) *Contacts {
	return &Contacts{
		Address: m["address"].(string),
		Phone:   m["phone"].(string),
		Fax:     m["fax"].(string),
		Site:    m["site"].(string),
	}
}

type ExpertInfo struct {
	Conclusion string `json:"conclusion"`
	Date       int64  `json:"date"`
	Name       string `json:"name"`
}

func NewExpertInfoFromMap(m map[string]interface{}) *ExpertInfo {
	return &ExpertInfo{
		Conclusion: m["conclusion"].(string),
		Date:       m["date"].(int64),
		Name:       m["name"].(string),
	}
}
