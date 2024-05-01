package infra

import (
	"testing"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/dev"
)

func TestContactQueryRepo(t *testing.T) {
	persistentDbSvc := dev.GetPersistentDbSvc()
	contactQueryRepo := NewContactQueryRepo(persistentDbSvc)

	t.Run("ReadContacts", func(t *testing.T) {
		contactList, err := contactQueryRepo.Read()
		if err != nil {
			t.Error(err)
		}

		if len(contactList) == 0 {
			t.Error("NoContactsFound")
		}
	})

	t.Run("ReadContactById", func(t *testing.T) {
		contactList, err := contactQueryRepo.Read()
		if err != nil {
			t.Error(err)
		}

		if len(contactList) == 0 {
			t.Error("NoContactsFound")
		}

		contactId := contactList[0].Id
		contact, err := contactQueryRepo.ReadById(contactId)
		if err != nil {
			t.Error(err)
		}

		if contact.Id != contactId {
			t.Error("ContactIdMismatch")
		}
	})
}
