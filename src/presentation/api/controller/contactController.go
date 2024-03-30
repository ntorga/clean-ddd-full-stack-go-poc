package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/useCase"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	apiHelper "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/helper"
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

// GetContacts	 godoc
// @Summary      GetContacts
// @Description  List contacts.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Success      200 {array} entity.Contact
// @Router       /v1/contact/ [get]
func (controller *ContactController) GetContacts(c echo.Context) error {
	contactsQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
	contactsList, err := useCase.GetContacts(contactsQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, contactsList)
}
