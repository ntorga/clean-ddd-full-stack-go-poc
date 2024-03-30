package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/useCase"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"
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
func (controller *ContactController) Get(c echo.Context) error {
	contactsQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
	contactsList, err := useCase.GetContacts(contactsQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, contactsList)
}

// CreateContact	 godoc
// @Summary      CreateNewContact
// @Description  Create a new contact.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Param        createContactDto 	  body    dto.CreateContact  true  "NewContact"
// @Success      201 {object} object{} "ContactCreated"
// @Router       /v1/contact/ [post]
func (controller *ContactController) Create(c echo.Context) error {
	requiredParams := []string{"name", "nickname", "phone"}
	requestBody, _ := apiHelper.GetRequestBody(c)

	apiHelper.CheckMissingParams(requestBody, requiredParams)

	createContactDto := dto.NewCreateContact(
		valueObject.NewPersonNamePanic(requestBody["name"].(string)),
		valueObject.NewNicknamePanic(requestBody["nickname"].(string)),
		valueObject.NewPhoneNumberPanic(requestBody["phone"].(string)),
	)

	persistentDbSvc := c.Get("persistentDbSvc").(*db.PersistentDatabaseService)
	contactQueryRepo := infra.NewContactQueryRepo(persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(persistentDbSvc)

	err := useCase.CreateContact(
		contactQueryRepo,
		contactCmdRepo,
		createContactDto,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusCreated, "ContactCreated")
}
