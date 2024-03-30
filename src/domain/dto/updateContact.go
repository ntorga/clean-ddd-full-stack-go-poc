package dto

import "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"

type UpdateContact struct {
	Id       valueObject.ContactId    `json:"id"`
	Name     *valueObject.PersonName  `json:"name"`
	Nickname *valueObject.Nickname    `json:"nickname"`
	Phone    *valueObject.PhoneNumber `json:"phone"`
}

func NewUpdateContact(
	id valueObject.ContactId,
	name *valueObject.PersonName,
	nickname *valueObject.Nickname,
	phone *valueObject.PhoneNumber,
) UpdateContact {
	return UpdateContact{
		Id:       id,
		Name:     name,
		Nickname: nickname,
		Phone:    phone,
	}
}
