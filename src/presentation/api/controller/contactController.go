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

// ReadContacts	 godoc
// @Summary      ReadContacts
// @Description  List contacts.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Success      200 {array} entity.Contact
// @Router       /v1/contact/ [get]
func (controller *ContactController) Read(c echo.Context) error {
	contactsQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
	contactsList, err := useCase.ReadContacts(contactsQueryRepo)
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
	requestBody, _ := apiHelper.ReadRequestBody(c)

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

// UpdateContact godoc
// @Summary      UpdateContact
// @Description  Update a contact.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Param        updateContactDto 	  body dto.UpdateContact  true  "UpdateContact (Only id is required.)"
// @Success      200 {object} object{} "ContactUpdated"
// @Router       /v1/contact/ [put]
func (controller *ContactController) Update(c echo.Context) error {
	requiredParams := []string{"id"}
	requestBody, _ := apiHelper.ReadRequestBody(c)

	apiHelper.CheckMissingParams(requestBody, requiredParams)

	id := valueObject.NewContactIdPanic(requestBody["id"])

	var namePtr *valueObject.PersonName
	if requestBody["name"] != nil {
		name := valueObject.NewPersonNamePanic(requestBody["name"].(string))
		namePtr = &name
	}

	var nickNamePtr *valueObject.Nickname
	if requestBody["nickname"] != nil {
		nickname := valueObject.NewNicknamePanic(requestBody["nickname"].(string))
		nickNamePtr = &nickname
	}

	var phonePtr *valueObject.PhoneNumber
	if requestBody["phone"] != nil {
		phone := valueObject.NewPhoneNumberPanic(requestBody["phone"].(string))
		phonePtr = &phone
	}

	updateContactDto := dto.NewUpdateContact(
		id,
		namePtr,
		nickNamePtr,
		phonePtr,
	)

	persistentDbSvc := c.Get("persistentDbSvc").(*db.PersistentDatabaseService)
	contactQueryRepo := infra.NewContactQueryRepo(persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(persistentDbSvc)

	err := useCase.UpdateContact(
		contactQueryRepo,
		contactCmdRepo,
		updateContactDto,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(
			c, http.StatusInternalServerError, err.Error(),
		)
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, "ContactUpdated")
}

// DeleteContact godoc
// @Summary      DeleteContact
// @Description  Delete a contact.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Param        id 	  path   string  true  "ContactId"
// @Success      200 {object} object{} "ContactDeleted"
// @Router       /v1/contact/{id}/ [delete]
func (controller *ContactController) Delete(c echo.Context) error {
	id := valueObject.NewContactIdPanic(c.Param("id"))

	persistentDbSvc := c.Get("persistentDbSvc").(*db.PersistentDatabaseService)
	contactQueryRepo := infra.NewContactQueryRepo(persistentDbSvc)
	contactCmdRepo := infra.NewContactCmdRepo(persistentDbSvc)

	err := useCase.DeleteContact(
		contactQueryRepo,
		contactCmdRepo,
		id,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, "ContactDeleted")
}
