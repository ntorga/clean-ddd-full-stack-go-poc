package infra

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
)

type ContactQueryRepo struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewContactQueryRepo(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactQueryRepo {
	return &ContactQueryRepo{
		persistentDbSvc: persistentDbSvc,
	}
}

func (repo *ContactQueryRepo) Get() ([]entity.Contact, error) {
	return []entity.Contact{}, nil
}
