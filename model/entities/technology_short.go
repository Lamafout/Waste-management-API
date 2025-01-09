package model

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
	return &TechnologyShort{
		Id:               technology.Id.Hex(),
		Name:             technology.Name,
		Assignment:       technology.Assignment,
		Characteristic:   technology.Characteristic,
		Fkko:             technology.Fkko,
		UseCase:          technology.UseCase,
		ExpertConclusion: technology.ExpertInfo.Conclusion,
	}
}
