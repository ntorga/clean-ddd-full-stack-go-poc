package infra

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	dbModel "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db/model"
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

func (repo *ContactQueryRepo) Read() ([]entity.Contact, error) {
	var entities []entity.Contact

	var models []dbModel.Contact
	err := repo.persistentDbSvc.Handler.
		Model(&dbModel.Contact{}).
		Find(&models).Error
	if err != nil {
		return entities, errors.New("GetDatabaseEntriesError")
	}

	for _, model := range models {
		entity, err := model.ToEntity()
		if err != nil {
			log.Printf("ModelToEntityError: %v", err.Error())
			continue
		}

		entities = append(entities, entity)
	}

	return entities, nil
}
