package dto

import "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"

type AddContact struct {
	Name     valueObject.PersonName  `json:"name"`
	Nickname valueObject.Nickname    `json:"nickname"`
	Phone    valueObject.PhoneNumber `json:"phone"`
}

func NewAddContact(
	name valueObject.PersonName,
	nickname valueObject.Nickname,
	phone valueObject.PhoneNumber,
) AddContact {
	return AddContact{
		Name:     name,
		Nickname: nickname,
		Phone:    phone,
	}
}
