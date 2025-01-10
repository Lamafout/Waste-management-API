package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TechnologyShort struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Assignment       string  `json:"assignment"`
	Characteristic   string  `json:"characteristic"`
	Fkko             []*Fkko `json:"fkko"`
	UseCase          string  `json:"useCase"`
	ExpertConclusion string  `json:"expertConclusion"`
}

func NewTechnologyShort(technology *Technology) *TechnologyShort {
	var id string
	if technology.Id != primitive.NilObjectID {
		id = technology.Id.Hex()
	} else {
		id = "unknown"
	}

	return &TechnologyShort{
		Id:               id,
		Name:             technology.Name,
		Assignment:       technology.Assignment,
		Characteristic:   technology.Characteristic,
		Fkko:             technology.Fkko,
		UseCase:          technology.UseCase,
		ExpertConclusion: technology.ExpertInfo.Conclusion,
	}
}
