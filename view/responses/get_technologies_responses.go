package responses

import model "waste_management/model/entities"

type GetTechnologiesResponse struct {
	Amount       int64                    `json:"amount"`
	Technologies []*model.TechnologyShort `json:"technologies"`
}

func NewGetTechnologiesResponse(amount int64, technologies []*model.TechnologyShort) *GetTechnologiesResponse {
	return &GetTechnologiesResponse{
		Amount:       amount,
		Technologies: technologies,
  
	}
}