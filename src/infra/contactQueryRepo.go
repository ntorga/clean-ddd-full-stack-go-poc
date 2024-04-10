package infra

import (
	"errors"
	"log"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/entity"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	dbModel "github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db/model"
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
		return entities, errors.New("ReadDatabaseEntriesError")
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

func (repo *ContactQueryRepo) ReadById(id valueObject.ContactId) (entity.Contact, error) {
	var entity entity.Contact

	var model dbModel.Contact
	err := repo.persistentDbSvc.Handler.
		Model(model).
		Where("id = ?", id.Uint()).
		First(&model).Error
	if err != nil {
		return entity, errors.New("ReadDatabaseEntryError")
	}

	entity, err = model.ToEntity()
	if err != nil {
		return entity, errors.New("ModelToEntityError")
	}

	return entity, nil
}
