package apiController

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	apiHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/api/helper"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/service"
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

// ReadContacts	 godoc
// @Summary      ReadContacts
// @Description  List contacts.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Success      200 {array} entity.Contact
// @Router       /v1/contact/ [get]
func (controller *ContactController) Read(c echo.Context) error {
	return apiHelper.ResponseWrapper(c, controller.contactService.Read())
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
	requestBody, err := apiHelper.ReadRequestBody(c)
	if err != nil {
		return err
	}

	return apiHelper.ResponseWrapper(
		c,
		controller.contactService.Create(requestBody),
	)
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
	requestBody, err := apiHelper.ReadRequestBody(c)
	if err != nil {
		return err
	}

	return apiHelper.ResponseWrapper(
		c,
		controller.contactService.Update(requestBody),
	)
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
	requestBody := map[string]interface{}{
		"id": c.Param("id"),
	}

	return apiHelper.ResponseWrapper(
		c,
		controller.contactService.Delete(requestBody),
	)
}
