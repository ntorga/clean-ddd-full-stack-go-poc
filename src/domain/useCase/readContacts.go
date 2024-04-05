package useCase

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/repository"
)

func ReadContacts(
	contactQueryRepo repository.ContactQueryRepo,
) ([]entity.Contact, error) {
	contactEntities, err := contactQueryRepo.Read()
	if err != nil {
		log.Printf("GetContactsError: %v", err)
		return []entity.Contact{}, errors.New("GetContactsInfraError")
	}

	return contactEntities, nil
}
