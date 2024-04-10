package cliController

import (
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	cliHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/cli/helper"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/liaison"
	"github.com/spf13/cobra"
)

type ContactController struct {
	contactLiaison *liaison.ContactLiaison
}

func NewContactController(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactController {
	return &ContactController{
		contactLiaison: liaison.NewContactLiaison(persistentDbSvc),
	}
}

func (controller *ContactController) Read() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "ReadContacts",
		Run: func(cmd *cobra.Command, args []string) {
			cliHelper.ResponseWrapper(controller.contactLiaison.Read())
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
			requestBody := map[string]interface{}{
				"name":     nameStr,
				"nickname": nicknameStr,
				"phone":    phoneStr,
			}

			cliHelper.ResponseWrapper(controller.contactLiaison.Create(requestBody))
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
			requestBody := map[string]interface{}{
				"id":       idStr,
				"name":     nameStr,
				"nickname": nicknameStr,
				"phone":    phoneStr,
			}

			cliHelper.ResponseWrapper(controller.contactLiaison.Update(requestBody))
		},
	}

	cmd.Flags().StringVarP(&idStr, "id", "i", "", "Id")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&nicknameStr, "nickname", "k", "", "Nickname")
	cmd.Flags().StringVarP(&phoneStr, "phone", "p", "", "Phone")
	return cmd
}

func (controller *ContactController) Delete() *cobra.Command {
	var idStr string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteContact",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"id": idStr,
			}

			cliHelper.ResponseWrapper(controller.contactLiaison.Delete(requestBody))
		},
	}

	cmd.Flags().StringVarP(&idStr, "id", "i", "", "ContactId")
	cmd.MarkFlagRequired("id")
	return cmd
}
