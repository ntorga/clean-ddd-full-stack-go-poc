package useCase

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/repository"
)

func CreateContact(
	contactQueryRepo repository.ContactQueryRepo,
	contactCmdRepo repository.ContactCmdRepo,
	createContact dto.CreateContact,
) error {
	err := contactCmdRepo.Create(createContact)
	if err != nil {
		log.Printf("CreateContactError: %s", err)
		return errors.New("CreateContactInfraError")
	}

	log.Printf("Contact '%v' created.", createContact.Name.String())

	return nil
}
