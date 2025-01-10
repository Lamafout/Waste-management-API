package controller

import model "waste_management/model/entities"

type FkkoController interface {
	GetFkkos() ([]*model.Fkko, error)
}