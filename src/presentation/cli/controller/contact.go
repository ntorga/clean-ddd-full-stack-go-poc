package cliController

import (
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	cliHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/cli/helper"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/service"
	"github.com/spf13/cobra"
)

type ContactController struct {
	contactService *service.ContactService
}

func NewContactController(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactController {
	return &ContactController{
		contactService: service.NewContactService(persistentDbSvc),
	}
}

func (controller *ContactController) Read() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "ReadContacts",
		Run: func(cmd *cobra.Command, args []string) {
			cliHelper.ResponseWrapper(controller.contactService.Read())
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

			cliHelper.ResponseWrapper(controller.contactService.Create(requestBody))
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
				"id": idStr,
			}

			if nameStr != "" {
				requestBody["name"] = nameStr
			}

			if nicknameStr != "" {
				requestBody["nickname"] = nicknameStr
			}

			if phoneStr != "" {
				requestBody["phone"] = phoneStr
			}

			cliHelper.ResponseWrapper(controller.contactService.Update(requestBody))
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

			cliHelper.ResponseWrapper(controller.contactService.Delete(requestBody))
		},
	}

	cmd.Flags().StringVarP(&idStr, "id", "i", "", "ContactId")
	cmd.MarkFlagRequired("id")
	return cmd
}
