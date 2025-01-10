package controller

import model "waste_management/model/entities"

type TechnologyController interface {
	GetTechnologies(filter string, page string) ([]*model.TechnologyShort, error)
	GetTechnology(id string) (*model.Technology, error)
	PostTechnology(m map[string]interface{}) error
}