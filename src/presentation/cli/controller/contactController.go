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

func (controller *ContactController) Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "GetContacts",
		Run: func(cmd *cobra.Command, args []string) {
			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactsList, err := useCase.GetContacts(contactQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, contactsList)
		},
	}

	return cmd
}

func (controller *ContactController) Add() *cobra.Command {
	var nameStr string
	var nicknameStr string
	var phoneStr string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "AddNewContact",
		Run: func(cmd *cobra.Command, args []string) {
			name := valueObject.NewPersonNamePanic(nameStr)
			nickname := valueObject.NewNicknamePanic(nicknameStr)
			phone := valueObject.NewPhoneNumberPanic(phoneStr)

			addContactDto := dto.NewAddContact(
				name,
				nickname,
				phone,
			)

			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactCmdRepo := infra.NewContactCmdRepo(controller.persistentDbSvc)

			err := useCase.AddContact(
				contactQueryRepo,
				contactCmdRepo,
				addContactDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "ContactAdded")
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
