package liaison

import (
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/dto"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/useCase"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	liaisonHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/liaison/helper"
)

type ContactLiaison struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewContactLiaison(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactLiaison {
	return &ContactLiaison{
		persistentDbSvc: persistentDbSvc,
	}
}

func (liaison *ContactLiaison) Read() LiaisonOutput {
	contactsQueryRepo := infra.NewContactQueryRepo(liaison.persistentDbSvc)
	contactsList, err := useCase.ReadContacts(contactsQueryRepo)
	if err != nil {
		return NewLiaisonOutput(InfraError, err.Error())
	}

	return NewLiaisonOutput(Success, contactsList)
}

func (liaison *ContactLiaison) Create(input map[string]interface{}) LiaisonOutput {
	requiredParams := []string{"name", "nickname", "phone"}

	err := liaisonHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	name, err := valueObject.NewPersonName(input["name"].(string))
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	nickname, err := valueObject.NewNickname(input["nickname"].(string))
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	phone, err := valueObject.NewPhoneNumber(input["phone"].(string))
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	createContactDto := dto.NewCreateContact(name, nickname, phone)

	contactQueryRepo := infra.NewContactQueryRepo(liaison.persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(liaison.persistentDbSvc)

	err = useCase.CreateContact(
		contactQueryRepo,
		contactCmdRepo,
		createContactDto,
	)
	if err != nil {
		return NewLiaisonOutput(InfraError, err.Error())
	}

	return NewLiaisonOutput(Created, "ContactCreated")
}

func (liaison *ContactLiaison) Update(input map[string]interface{}) LiaisonOutput {
	requiredParams := []string{"id"}

	err := liaisonHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	id, err := valueObject.NewContactId(input["id"])
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	var namePtr *valueObject.PersonName
	if input["name"] != nil {
		name, err := valueObject.NewPersonName(input["name"].(string))
		if err != nil {
			return NewLiaisonOutput(UserError, err.Error())
		}
		namePtr = &name
	}

	var nickNamePtr *valueObject.Nickname
	if input["nickname"] != nil {
		nickname, err := valueObject.NewNickname(input["nickname"].(string))
		if err != nil {
			return NewLiaisonOutput(UserError, err.Error())
		}
		nickNamePtr = &nickname
	}

	var phonePtr *valueObject.PhoneNumber
	if input["phone"] != nil {
		phone, err := valueObject.NewPhoneNumber(input["phone"].(string))
		if err != nil {
			return NewLiaisonOutput(UserError, err.Error())
		}
		phonePtr = &phone
	}

	updateContactDto := dto.NewUpdateContact(
		id,
		namePtr,
		nickNamePtr,
		phonePtr,
	)

	contactQueryRepo := infra.NewContactQueryRepo(liaison.persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(liaison.persistentDbSvc)

	err = useCase.UpdateContact(
		contactQueryRepo,
		contactCmdRepo,
		updateContactDto,
	)
	if err != nil {
		return NewLiaisonOutput(InfraError, err.Error())
	}

	return NewLiaisonOutput(Success, "ContactUpdated")
}

func (liaison *ContactLiaison) Delete(input map[string]interface{}) LiaisonOutput {
	id, err := valueObject.NewContactId(input["id"])
	if err != nil {
		return NewLiaisonOutput(UserError, err.Error())
	}

	contactQueryRepo := infra.NewContactQueryRepo(liaison.persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(liaison.persistentDbSvc)

	err = useCase.DeleteContact(
		contactQueryRepo,
		contactCmdRepo,
		id,
	)
	if err != nil {
		return NewLiaisonOutput(InfraError, err.Error())
	}

	return NewLiaisonOutput(Success, "ContactDeleted")
}
