package entity

import "github.com/google/uuid"

type Car struct {
	ID     uuid.UUID `json:"id,omitempty"`
	RegNum string    `json:"regNum,omitempty"`
	Mark   string    `json:"mark,omitempty"`
	Model  string    `json:"model,omitempty"`
	Year   int       `json:"year,omitempty"`
	Owner  *Owner    `json:"owner,omitempty"`
}

type Owner struct {
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
}

type RegNums struct {
	Nums []string `json:"regNum,omitempty"`
}

type CarUpdate struct {
	RegNum *string      `json:"regNum,omitempty"`
	Mark   *string      `json:"mark,omitempty"`
	Model  *string      `json:"model,omitempty"`
	Year   *int         `json:"year,omitempty"`
	Owner  *OwnerUpdate `json:"owner,omitempty"`
}

type OwnerUpdate struct {
	Name       *string `json:"name,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Patronymic *string `json:"patronymic,omitempty"`
}
