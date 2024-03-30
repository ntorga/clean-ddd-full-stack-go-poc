package cliController

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/useCase"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra"
	cliHelper "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

type ContactController struct {
}

func NewContactController() *ContactController {
	return &ContactController{}
}

func (controller *ContactController) GetContacts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "GetContacts",
		Run: func(cmd *cobra.Command, args []string) {
			contactQueryRepo := infra.NewContactQueryRepo()
			contactsList, err := useCase.GetContacts(contactQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, contactsList)
		},
	}

	return cmd
}
