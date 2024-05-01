package db

import (
	"errors"
	"os"
	"reflect"

	"github.com/glebarez/sqlite"
	dbModel "github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db/model"
	"gorm.io/gorm"
)

var DatabaseFilePath = os.TempDir() + "/app.db"

type PersistentDatabaseService struct {
	Handler *gorm.DB
}

func NewPersistentDatabaseService() (*PersistentDatabaseService, error) {
	ormSvc, err := gorm.Open(
		sqlite.Open(DatabaseFilePath),
		&gorm.Config{},
	)
	if err != nil {
		return nil, errors.New("DatabaseConnectionError")
	}

	dbSvc := &PersistentDatabaseService{Handler: ormSvc}
	err = dbSvc.dbMigrate()
	if err != nil {
		return nil, err
	}

	return dbSvc, nil
}

func (dbSvc *PersistentDatabaseService) isTableEmpty(model interface{}) (bool, error) {
	var count int64
	err := dbSvc.Handler.Model(&model).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (dbSvc *PersistentDatabaseService) seedDatabase(seedModels ...interface{}) error {
	for _, seedModel := range seedModels {
		isTableEmpty, err := dbSvc.isTableEmpty(seedModel)
		if err != nil {
			return err
		}

		if !isTableEmpty {
			continue
		}

		seedModelType := reflect.TypeOf(seedModel).Elem()
		seedModelFieldsAndMethods := reflect.ValueOf(seedModel)

		seedModelInitialEntriesMethod := seedModelFieldsAndMethods.MethodByName(
			"InitialEntries",
		)
		seedModelInitialEntriesMethodResults := seedModelInitialEntriesMethod.Call(
			[]reflect.Value{},
		)
		initialEntries := seedModelInitialEntriesMethodResults[0].Interface()

		for _, entry := range initialEntries.([]interface{}) {
			entryInnerStructure := reflect.ValueOf(entry)

			entryFormatHandlerWillAccept := reflect.New(seedModelType)
			entryFormatHandlerWillAccept.Elem().Set(entryInnerStructure)
			adjustedEntry := entryFormatHandlerWillAccept.Interface()

			err = dbSvc.Handler.Create(adjustedEntry).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (dbSvc *PersistentDatabaseService) dbMigrate() error {
	err := dbSvc.Handler.AutoMigrate(
		&dbModel.Contact{},
	)
	if err != nil {
		return errors.New("DatabaseMigrationError")
	}

	modelsWithInitialEntries := []interface{}{
		&dbModel.Contact{},
	}

	err = dbSvc.seedDatabase(modelsWithInitialEntries...)
	if err != nil {
		return errors.New("CreateDefaultDatabaseEntriesError")
	}

	return nil
}
