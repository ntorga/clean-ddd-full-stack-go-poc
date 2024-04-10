package useCase

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/repository"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"
)

func DeleteContact(
	contactQueryRepo repository.ContactQueryRepo,
	contactCmdRepo repository.ContactCmdRepo,
	id valueObject.ContactId,
) error {
	_, err := contactQueryRepo.ReadById(id)
	if err != nil {
		return errors.New("ContactNotFound")
	}

	err = contactCmdRepo.Delete(id)
	if err != nil {
		log.Printf("DeleteContactError: %s", err)
		return errors.New("DeleteContactInfraError")
	}

	log.Printf("ContactId '%v' deleted.", id)

	return nil
}
