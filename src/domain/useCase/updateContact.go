package useCase

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/repository"
)

func UpdateContact(
	contactQueryRepo repository.ContactQueryRepo,
	contactCmdRepo repository.ContactCmdRepo,
	dto dto.UpdateContact,
) error {
	_, err := contactQueryRepo.ReadById(dto.Id)
	if err != nil {
		return errors.New("ContactNotFound")
	}

	err = contactCmdRepo.Update(dto)
	if err != nil {
		log.Printf("UpdateContactError: %s", err)
		return errors.New("UpdateContactInfraError")
	}

	log.Printf("ContactId '%v' updated.", dto.Id)

	return nil
}
