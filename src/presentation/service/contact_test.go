package service

import (
	"testing"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/dev"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/entity"
)

func getLatestContact(contactService *ContactService) (entity.Contact, error) {
	var contact entity.Contact

	serviceOutput := contactService.Read()
	if serviceOutput.Status != Success {
		return contact, nil
	}

	contactsList := serviceOutput.Body.([]entity.Contact)
	lastItemIndex := len(contactsList) - 1
	return contactsList[lastItemIndex], nil
}

func TestContactService(t *testing.T) {
	persistentDbSvc := dev.GetPersistentDbSvc()
	contactService := NewContactService(persistentDbSvc)

	t.Run("ReadContacts", func(t *testing.T) {
		serviceOutput := contactService.Read()
		if serviceOutput.Status != Success {
			t.Error(serviceOutput.Body)
		}
	})

	t.Run("CreateContact", func(t *testing.T) {
		input := map[string]interface{}{
			"name":     "John Doe",
			"nickname": "JD",
			"phone":    "123-456",
		}

		serviceOutput := contactService.Create(input)
		if serviceOutput.Status != Created {
			t.Error(serviceOutput.Body)
		}
	})

	t.Run("UpdateContact", func(t *testing.T) {
		contactEntity, err := getLatestContact(contactService)
		if err != nil {
			t.Error(err)
		}

		input := map[string]interface{}{
			"id":   contactEntity.Id.String(),
			"name": "Jane Doe",
		}

		serviceOutput := contactService.Update(input)
		if serviceOutput.Status != Success {
			t.Error(serviceOutput.Body)
		}
	})

	t.Run("DeleteContact", func(t *testing.T) {
		contactEntity, err := getLatestContact(contactService)
		if err != nil {
			t.Error(err)
		}

		input := map[string]interface{}{
			"id": contactEntity.Id.String(),
		}

		serviceOutput := contactService.Delete(input)
		if serviceOutput.Status != Success {
			t.Error(serviceOutput.Body)
		}
	})
}
