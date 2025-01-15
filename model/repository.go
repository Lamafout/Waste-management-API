package model

import (
	model "waste_management/model/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fields := []string{"name", "fkko.code", "fkko.name"}

	var uniqueProducersMap = make(map[string]map[string]interface{})

	for _, field := range fields {
		results, err := r.Client.ReadFiltered("producers", filter, field, page)
		if err != nil {
			return nil, 0, err
		}

		for _, result := range results {
			resultId := result["_id"].(primitive.ObjectID).Hex()
			uniqueProducersMap[resultId] = result
		}
	}

	var producers []*model.Producer  

	for _, producer := range uniqueProducersMap {
		producers = append(producers, model.NewProducerFromMap(producer))
	}

	amount, err := r.Client.Count("producers", fields, filter)

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
	fields := []string{"fkko.name", "fkko.code", "name"}

	var uniqueTechnologiesMap = make(map[string]map[string]interface{})

	for _, field := range fields {
		results, err := r.Client.ReadFiltered("technologies", filter, field, page)
		if err != nil {
			return nil, 0, err
		}

		for _, result := range results {
			resultId := result["_id"].(primitive.ObjectID).Hex()
			uniqueTechnologiesMap[resultId] = result
		}
	}
	var technologies []*model.TechnologyShort

	for _, technology := range uniqueTechnologiesMap {
		technologyBase := model.NewTechnologyFromMap(technology)
		technologies = append(technologies, model.NewTechnologyShort(technologyBase))
	}

	amount, err := r.Client.Count("technologies", fields, filter)

	if err != nil {
		return nil, 0, err
	}

	return technologies, amount, nil
}

func (r *Repository) GetFkkos(filter string) ([]*model.Fkko, error) {
	fkkosByCodeMap , err := r.Client.ReadFiltered("fkkos", filter, "code", -1)
	if err != nil {
		return nil, err
	}

	fkkosByNameMap, err := r.Client.ReadFiltered("fkkos", filter, "name", -1)
	if err != nil {
		return nil, err
	}

	var fkkos []*model.Fkko

	for _, fkko := range fkkosByCodeMap {
		fkkos = append(fkkos, model.NewFkkoFromMap(fkko))
	}

	for _, fkko := range fkkosByNameMap {
		fkkos = append(fkkos, model.NewFkkoFromMap(fkko))
	}

	return fkkos, nil
}

func (r *Repository) GetOkpds(filter string) ([]*model.Okpd, error) {
	okpdsByCodeMap, err := r.Client.ReadFiltered("okpds", filter, "code", -1)
	if err != nil {
		return nil, err
	}

	okpdsByNameMap, err := r.Client.ReadFiltered("okpds", filter, "name", -1)
	if err != nil {
		return nil, err
	}

	var okpds []*model.Okpd

	for _, okpd := range okpdsByCodeMap {
		okpds = append(okpds, model.NewOkpdFromMap(okpd))
	}

	for _, okpd := range okpdsByNameMap {
		okpds = append(okpds, model.NewOkpdFromMap(okpd))
	}

	return okpds, nil
}