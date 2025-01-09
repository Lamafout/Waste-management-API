package model

type Technology struct {
	Name           string         `json:"name"`
	Assignment     string         `json:"assignment"`
	Characteristic string         `json:"characteristic"`
	Resources      Resources      `json:"resources"`
	Fkko           []Fkko         `json:"fkko"`
	Okpd           []Okpd         `json:"okpd"`
	Performance    float64        `json:"performance"`
	SecondaryWaste SecondaryWaste `json:"secondaryWaste"`
	Developer      Contacts       `json:"developer"`
	Users          []Contacts     `json:"users"`
	UseCase        string         `json:"useCase"`
	ExpertInfo
}

type Resources struct {
	Energy       float64 `json:"energy"`
	Water        float64 `json:"water"`
	UsingPerYear float64 `json:"usingPerYear"`
}

type SecondaryWaste struct {
	Fkko   Fkko    `json:"fkko"`
	Mass   float64 `json:"mass"`
	Volume float64 `json:"volume"`
}

type Contacts struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Fax   string `json:"fax"`
	Site  string `json:"site"`
}

type ExpertInfo struct {
	Conclusion string `json:"conclusion"`
	Date       int64  `json:"date"`
	Name       string `json:"name"`
}
