package useCase

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/entity"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/repository"
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
