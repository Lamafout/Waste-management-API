package controller

import (
	"log"
	"strconv"
	"waste_management/model"
	entities "waste_management/model/entities"
)

type Controller struct {
	repository *model.Repository
}

func NewController(repository *model.Repository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) GetProducers(filter string, page string) ([]*entities.Producer, int64, error){
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error converting page number: ", err)
		return nil, 0, err
	}

	producers, amount, err := c.repository.GetProducers(filter, pageNum)

	if err != nil {
		log.Println("Error getting producers: ", err)
		return nil, 0, err
	}

	return producers, amount, nil
}

func (c *Controller) PostProducer(m map[string]interface{}) error {
	err := c.repository.CreateProducer(*entities.NewProducerFromMap(m))

	if err != nil {
		log.Println("Error creating producer: ", err)
		return err
	}

	return nil
}

func (c *Controller) GetTechnology(id string) (*entities.Technology, error) {
	technology, err := c.repository.GetTechnology(id)

	if err != nil {
		log.Println("Error getting technology: ", err)
		return nil, err
	}

	return technology, nil
}

func (c *Controller) GetTechnologies(filter string, page string) ([]*entities.TechnologyShort, int64, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error converting page number: ", err)
		return nil, 0, err
	}

	technologies, amount, err := c.repository.GetTechnologies(filter, pageNum)

	if err != nil {
		log.Println("Error getting technologies: ", err)
		return nil, 0, err
	}

	return technologies, amount, nil
}

func (c *Controller) PostTechnology(m map[string]interface{}) error {
	err := c.repository.CreateTechnology(*entities.NewTechnologyFromMap(m))

	if err != nil {
		log.Println("Error creating technology: ", err)
		return err
	}

	return nil
}

func (c *Controller) GetFkkos(filter string) ([]*entities.Fkko, error) {
	fkkos, err := c.repository.GetFkkos(filter)

	if err != nil {
		log.Println("Error getting fkkos: ", err)
		return nil, err
	}

	return fkkos, nil
}

func (c *Controller) GetOkpds(filter string) ([]*entities.Okpd, error) {
	okpds, err := c.repository.GetOkpds(filter)

	if err != nil {
		log.Println("Error getting okpds: ", err)
		return nil, err
	}

	return okpds, nil
}