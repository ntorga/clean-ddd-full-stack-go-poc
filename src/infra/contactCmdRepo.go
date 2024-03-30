package infra

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"
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

func (repo *ContactCmdRepo) Update(dto dto.UpdateContact) error {
	updateMap := map[string]interface{}{}

	if dto.Name != nil {
		updateMap["name"] = dto.Name.String()
	}

	if dto.Nickname != nil {
		updateMap["nickname"] = dto.Nickname.String()
	}

	if dto.Phone != nil {
		updateMap["phone"] = dto.Phone.String()
	}

	return repo.persistentDbSvc.Handler.
		Model(&dbModel.Contact{}).
		Where("id = ?", dto.Id.String()).
		Updates(updateMap).Error
}

func (repo *ContactCmdRepo) Delete(id valueObject.ContactId) error {
	return repo.persistentDbSvc.Handler.Delete(&dbModel.Contact{}, id.Uint()).Error
}
