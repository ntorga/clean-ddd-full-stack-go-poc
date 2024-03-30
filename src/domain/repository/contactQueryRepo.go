package repository

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
)

type ContactQueryRepo interface {
	Read() ([]entity.Contact, error)
}
