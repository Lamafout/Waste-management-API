package model

import (
	model "waste_management/model/entities"
)

type Repository struct {
	Client *connection
}

func NewRepository(client *connection) *Repository {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) CreateProducer(producer model.Producer) error {
	err := r.Client.Create("producers", producer)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetProducers(filter string, page int) ([]*model.Producer, int64,  error) {
	producersMap, err := r.Client.ReadFiltered("producers", filter, "name", page)

	if err != nil {
		return nil, 0, err
	}

	var producers []*model.Producer  

	for _, producer := range producersMap {
		producers = append(producers, model.NewProducerFromMap(producer))
	}

	amount, err := r.Client.Count("producers")

	if err != nil {
		return nil, 0, err
	}

	return producers, amount, nil
}

func (r *Repository) CreateTechnology(technology model.Technology) error {
	err := r.Client.Create("technologies", technology)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetTechnology(id string) (*model.Technology, error) {
	technologyMap, err := r.Client.Read("technologies", id)

	if err != nil {
		return nil, err
	}

	return model.NewTechnologyFromMap(technologyMap), nil
}

func (r *Repository) GetTechnologies(filter string, page int) ([]*model.TechnologyShort, int64,  error) {
	technologiesMap, err := r.Client.ReadFiltered("technologies", filter, "name", page)

	if err != nil {
		return nil, 0, err
	}

	var technologies []*model.TechnologyShort

	for _, technology := range technologiesMap {
		technologyBase := model.NewTechnologyFromMap(technology)
		technologies = append(technologies, model.NewTechnologyShort(technologyBase))
	}

	amount, err := r.Client.Count("technologies")

	if err != nil {
		return nil, 0, err
	}

	return technologies, amount, nil
}

func (r *Repository) GetFkkos(filter string) ([]*model.Fkko, error) {
	fkkosMap, err := r.Client.ReadFiltered("fkkos", filter, "name", -1)

	if err != nil {
		return nil, err
	}

	var fkkos []*model.Fkko

	for _, fkko := range fkkosMap {
		fkkos = append(fkkos, model.NewFkkoFromMap(fkko))
	}

	return fkkos, nil
}

func (r *Repository) GetOkpds(filter string) ([]*model.Okpd, error) {
	okpdsMap, err := r.Client.ReadFiltered("okpds", filter, "name", -1)

	if err != nil {
		return nil, err
	}

	var okpds []*model.Okpd

	for _, okpd := range okpdsMap {
		okpds = append(okpds, model.NewOkpdFromMap(okpd))
	}

	return okpds, nil
}