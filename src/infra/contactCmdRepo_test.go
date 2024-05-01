package infra

import (
	"testing"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/dev"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/dto"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/entity"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"
)

func getLatestContact(contactQueryRepo *ContactQueryRepo) (entity.Contact, error) {
	var contact entity.Contact

	contactList, err := contactQueryRepo.Read()
	if err != nil {
		return contact, err
	}
	return contactList[len(contactList)-1], nil
}

func TestContactCmdRepo(t *testing.T) {
	persistentDbSvc := dev.GetPersistentDbSvc()
	contactQueryRepo := NewContactQueryRepo(persistentDbSvc)
	contactCmdRepo := NewContactCmdRepo(persistentDbSvc)

	t.Run("CreateContact", func(t *testing.T) {
		name, _ := valueObject.NewPersonName("John Doe")
		nickname, _ := valueObject.NewNickname("JD")
		phone, _ := valueObject.NewPhoneNumber("123-456")

		dto := dto.NewCreateContact(name, nickname, phone)

		err := contactCmdRepo.Create(dto)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("UpdateContact", func(t *testing.T) {
		lastContactEntity, err := getLatestContact(contactQueryRepo)
		if err != nil {
			t.Error(err)
		}

		contactId := lastContactEntity.Id
		name, _ := valueObject.NewPersonName("Jane Doe")

		dto := dto.NewUpdateContact(contactId, &name, nil, nil)

		err = contactCmdRepo.Update(dto)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteContact", func(t *testing.T) {
		lastContactEntity, err := getLatestContact(contactQueryRepo)
		if err != nil {
			t.Error(err)
		}

		contactId := lastContactEntity.Id

		err = contactCmdRepo.Delete(contactId)
		if err != nil {
			t.Error(err)
		}
	})
}
