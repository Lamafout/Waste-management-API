package responses

import model "waste_management/model/entities"

type GetProducersResponse struct {
	Amount    int64             `json:"amount"`
	Producers []*model.Producer `json:"producers"`
}

func NewGetProducersResponse(amount int64, producers []*model.Producer) *GetProducersResponse {
	return &GetProducersResponse{
		Amount:    amount,
		Producers: producers,
	}
}
