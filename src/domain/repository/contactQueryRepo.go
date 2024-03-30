package repository

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"
)

type ContactQueryRepo interface {
	Read() ([]entity.Contact, error)
	ReadById(id valueObject.ContactId) (entity.Contact, error)
}
