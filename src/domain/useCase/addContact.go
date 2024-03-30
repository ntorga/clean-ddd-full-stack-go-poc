package useCase

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/repository"
)

func AddContact(
	contactQueryRepo repository.ContactQueryRepo,
	contactCmdRepo repository.ContactCmdRepo,
	addContact dto.AddContact,
) error {
	err := contactCmdRepo.Add(addContact)
	if err != nil {
		log.Printf("AddContactError: %s", err)
		return errors.New("AddContactInfraError")
	}

	log.Printf("Contact '%v' added.", addContact.Name.String())

	return nil
}
