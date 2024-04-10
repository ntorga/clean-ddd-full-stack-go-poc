package dto

import "github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"

type CreateContact struct {
	Name     valueObject.PersonName  `json:"name"`
	Nickname valueObject.Nickname    `json:"nickname"`
	Phone    valueObject.PhoneNumber `json:"phone"`
}

func NewCreateContact(
	name valueObject.PersonName,
	nickname valueObject.Nickname,
	phone valueObject.PhoneNumber,
) CreateContact {
	return CreateContact{
		Name:     name,
		Nickname: nickname,
		Phone:    phone,
	}
}
