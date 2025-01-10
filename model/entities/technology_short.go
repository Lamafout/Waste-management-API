package model

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		ExpertConclusion: NewExpertInfoShort(technology.ExpertInfo).ExpertConclusion,
	}
}

type ExpertInfoShort struct {
	expertInfo       *ExpertInfo
	ExpertConclusion string `json:"expertConclusion"`
}

func NewExpertInfoShort(expertInfo *ExpertInfo) *ExpertInfoShort {
	conclusion := &ExpertInfoShort{
		expertInfo: expertInfo,
	}

	conclusion.ExpertConclusion = conclusion.ConstructConclusion()

	return conclusion
}

func (e *ExpertInfoShort) ConstructConclusion() string {
	conclusionDate := time.Unix(e.expertInfo.Date, 0).Format("2.01.2006")

	result := fmt.Sprintf(
		`%s
	Экспертное 
	заключение №%s
	от %s`, e.expertInfo.Conclusion, e.expertInfo.Number, conclusionDate)

	return result
}
