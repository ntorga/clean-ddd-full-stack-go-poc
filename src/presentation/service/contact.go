package service

import (
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/dto"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/useCase"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	serviceHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/service/helper"
)

type ContactService struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewContactService(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactService {
	return &ContactService{
		persistentDbSvc: persistentDbSvc,
	}
}

func (service *ContactService) Read() ServiceOutput {
	contactsQueryRepo := infra.NewContactQueryRepo(service.persistentDbSvc)
	contactsList, err := useCase.ReadContacts(contactsQueryRepo)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, contactsList)
}

func (service *ContactService) Create(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"name", "nickname", "phone"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	name, err := valueObject.NewPersonName(input["name"].(string))
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	nickname, err := valueObject.NewNickname(input["nickname"].(string))
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	phone, err := valueObject.NewPhoneNumber(input["phone"].(string))
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	createContactDto := dto.NewCreateContact(name, nickname, phone)

	contactQueryRepo := infra.NewContactQueryRepo(service.persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(service.persistentDbSvc)

	err = useCase.CreateContact(
		contactQueryRepo,
		contactCmdRepo,
		createContactDto,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Created, "ContactCreated")
}

func (service *ContactService) Update(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"id"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	id, err := valueObject.NewContactId(input["id"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var namePtr *valueObject.PersonName
	if input["name"] != nil {
		name, err := valueObject.NewPersonName(input["name"].(string))
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		namePtr = &name
	}

	var nickNamePtr *valueObject.Nickname
	if input["nickname"] != nil {
		nickname, err := valueObject.NewNickname(input["nickname"].(string))
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		nickNamePtr = &nickname
	}

	var phonePtr *valueObject.PhoneNumber
	if input["phone"] != nil {
		phone, err := valueObject.NewPhoneNumber(input["phone"].(string))
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		phonePtr = &phone
	}

	updateContactDto := dto.NewUpdateContact(
		id,
		namePtr,
		nickNamePtr,
		phonePtr,
	)

	contactQueryRepo := infra.NewContactQueryRepo(service.persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(service.persistentDbSvc)

	err = useCase.UpdateContact(
		contactQueryRepo,
		contactCmdRepo,
		updateContactDto,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "ContactUpdated")
}

func (service *ContactService) Delete(input map[string]interface{}) ServiceOutput {
	id, err := valueObject.NewContactId(input["id"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	contactQueryRepo := infra.NewContactQueryRepo(service.persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(service.persistentDbSvc)

	err = useCase.DeleteContact(
		contactQueryRepo,
		contactCmdRepo,
		id,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "ContactDeleted")
}
