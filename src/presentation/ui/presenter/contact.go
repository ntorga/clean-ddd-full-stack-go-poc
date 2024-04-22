package presenter

import (
	"github.com/labstack/echo/v4"

	"net/http"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/entity"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/service"
	uiHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/ui/helper"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/ui/page"
)

type ContactPresenter struct {
	contactService *service.ContactService
}

func NewContactPresenter(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactPresenter {
	return &ContactPresenter{
		contactService: service.NewContactService(persistentDbSvc),
	}
}

func (presenter *ContactPresenter) Handler(c echo.Context) error {
	responseOutput := presenter.contactService.Read()
	if responseOutput.Status != service.Success {
		return nil
	}

	contactEntities, assertOk := responseOutput.Body.([]entity.Contact)
	if !assertOk {
		return nil
	}

	return uiHelper.Render(c, page.ContactIndex(contactEntities), http.StatusOK)
}
