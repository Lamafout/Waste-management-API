package controller

import model "waste_management/model/entities"

type OkpdController interface {
	GetOkpds(filter string) ([]*model.Okpd, error)
}