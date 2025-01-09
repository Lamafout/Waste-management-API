package model

type Producer struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Fkko        Fkko   `json:"fkko"`
	HazardClass string `json:"hazardClass"`
}
