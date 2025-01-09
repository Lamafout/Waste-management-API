package controller

import (
	"log"
	"waste_management/model"
	entities "waste_management/model/entities"
)

type Controller struct {
	repository *model.Repository
}

func NewController(repository *model.Repository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) GetProducers() ([]*entities.Producer, error){
	producers, err := c.repository.GetProducers()

	if err != nil {
		log.Println("Error getting producers: ", err)
		return nil, err
	}

	return producers, nil
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

func (c *Controller) GetTechnologies() ([]*entities.TechnologyShort, error) {
	technologies, err := c.repository.GetTechnologies()

	if err != nil {
		log.Println("Error getting technologies: ", err)
		return nil, err
	}

	return technologies, nil
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