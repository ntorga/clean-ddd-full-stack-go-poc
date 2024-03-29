package useCase

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/repository"
)

func GetContacts(
	contactQueryRepo repository.ContactQueryRepo,
) ([]entity.Contact, error) {
	return contactQueryRepo.Get()
}
