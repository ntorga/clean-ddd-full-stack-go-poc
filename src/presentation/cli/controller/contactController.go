package cliController

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/useCase"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	cliHelper "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

type ContactController struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewContactController(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactController {
	return &ContactController{
		persistentDbSvc: persistentDbSvc,
	}
}

func (controller *ContactController) Read() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "ReadContacts",
		Run: func(cmd *cobra.Command, args []string) {
			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactsList, err := useCase.ReadContacts(contactQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, contactsList)
		},
	}

	return cmd
}

func (controller *ContactController) Create() *cobra.Command {
	var nameStr string
	var nicknameStr string
	var phoneStr string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateNewContact",
		Run: func(cmd *cobra.Command, args []string) {
			name := valueObject.NewPersonNamePanic(nameStr)
			nickname := valueObject.NewNicknamePanic(nicknameStr)
			phone := valueObject.NewPhoneNumberPanic(phoneStr)

			createContactDto := dto.NewCreateContact(
				name,
				nickname,
				phone,
			)

			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactCmdRepo := infra.NewContactCmdRepo(controller.persistentDbSvc)

			err := useCase.CreateContact(
				contactQueryRepo,
				contactCmdRepo,
				createContactDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "ContactCreated")
		},
	}

	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "Name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&nicknameStr, "nickname", "k", "", "Nickname")
	cmd.MarkFlagRequired("nickname")
	cmd.Flags().StringVarP(&phoneStr, "phone", "p", "", "Phone")
	cmd.MarkFlagRequired("phone")
	return cmd
}

func (controller *ContactController) Update() *cobra.Command {
	var idStr string
	var nameStr string
	var nicknameStr string
	var phoneStr string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "UpdateContact",
		Run: func(cmd *cobra.Command, args []string) {
			id := valueObject.NewContactIdPanic(idStr)

			var namePtr *valueObject.PersonName
			if nameStr != "" {
				name := valueObject.NewPersonNamePanic(nameStr)
				namePtr = &name
			}

			var nicknamePtr *valueObject.Nickname
			if nicknameStr != "" {
				nickname := valueObject.NewNicknamePanic(nicknameStr)
				nicknamePtr = &nickname
			}

			var phonePtr *valueObject.PhoneNumber
			if phoneStr != "" {
				phone := valueObject.NewPhoneNumberPanic(phoneStr)
				phonePtr = &phone
			}

			updateContactDto := dto.NewUpdateContact(
				id,
				namePtr,
				nicknamePtr,
				phonePtr,
			)

			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactCmdRepo := infra.NewContactCmdRepo(controller.persistentDbSvc)

			err := useCase.UpdateContact(
				contactQueryRepo,
				contactCmdRepo,
				updateContactDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "ContactUpdated")
		},
	}

	cmd.Flags().StringVarP(&idStr, "id", "i", "", "Id")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&nicknameStr, "nickname", "k", "", "Nickname")
	cmd.Flags().StringVarP(&phoneStr, "phone", "p", "", "Phone")
	return cmd
}
