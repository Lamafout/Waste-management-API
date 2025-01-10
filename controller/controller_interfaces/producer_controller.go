package controller

import model "waste_management/model/entities"

type ProducerController interface {
	GetProducers(filter string, page string) ([]*model.Producer, error)
	PostProducer(m map[string]interface{}) error
}