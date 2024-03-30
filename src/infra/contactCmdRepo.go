package infra

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	dbModel "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db/model"
)

type ContactCmdRepo struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewContactCmdRepo(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactCmdRepo {
	return &ContactCmdRepo{
		persistentDbSvc: persistentDbSvc,
	}
}

func (repo *ContactCmdRepo) Create(dto dto.CreateContact) error {
	model := dbModel.Contact{
		Name:     dto.Name.String(),
		Nickname: dto.Nickname.String(),
		Phone:    dto.Phone.String(),
	}

	return repo.persistentDbSvc.Handler.Create(&model).Error
}
