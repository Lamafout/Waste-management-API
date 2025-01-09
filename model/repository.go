package model

import model "waste_management/model/entities"

type repository struct {
	Client *connection
}

func NewRepository(client *connection) *repository {
	return &repository{
		Client: client,
	}
}

func (r *repository) CreateProducer(producer model.Producer) error {
	err := r.Client.Create("producers", producer)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetProducers() ([]*model.Producer, error) {
	producersMap, err := r.Client.ReadAll("producers")

	if err != nil {
		return nil, err
	}

	var producers []*model.Producer  

	for _, producer := range producersMap {
		producers = append(producers, model.NewProducerFromMap(producer))
	}

	return producers, nil
}

func (r *repository) CreateTechnology(technology model.Technology) error {
	err := r.Client.Create("technologies", technology)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetTechnology(id int64) (*model.Technology, error) {
	technologyMap, err := r.Client.Read("technologies", id)

	if err != nil {
		return nil, err
	}

	return model.NewTechnologyFromMap(technologyMap), nil
}

func (r *repository) GetTechnologies() ([]*model.TechnologyShort, error) {
	technologiesMap, err := r.Client.ReadAll("technologies")

	if err != nil {
		return nil, err
	}

	var technologies []*model.TechnologyShort

	for _, technology := range technologiesMap {
		technologyBase := model.NewTechnologyFromMap(technology)
		technologies = append(technologies, model.NewTechnologyShort(technologyBase))
	}

	return technologies, nil
}

func (r *repository) GetFkkos(filter string) ([]*model.Fkko, error) {
	fkkosMap, err := r.Client.ReadFiltered("fkkos", filter, "name")

	if err != nil {
		return nil, err
	}

	var fkkos []*model.Fkko

	for _, fkko := range fkkosMap {
		fkkos = append(fkkos, model.NewFkkoFromMap(fkko))
	}

	return fkkos, nil
}

func (r *repository) GetOkpds(filter string) ([]*model.Okpd, error) {
	okpdsMap, err := r.Client.ReadFiltered("okpds", filter, "name")

	if err != nil {
		return nil, err
	}

	var okpds []*model.Okpd

	for _, okpd := range okpdsMap {
		okpds = append(okpds, model.NewOkpdFromMap(okpd))
	}

	return okpds, nil
}