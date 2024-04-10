package cli

import (
	"fmt"

	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation"
	cliController "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/cli/controller"
	"github.com/spf13/cobra"
)

type Router struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewRouter(
	persistentDbSvc *db.PersistentDatabaseService,
) *Router {
	return &Router{
		persistentDbSvc: persistentDbSvc,
	}
}

func (router *Router) contactRoutes() {
	var contactCmd = &cobra.Command{
		Use:   "contact",
		Short: "ContactManagement",
	}

	contactController := cliController.NewContactController(router.persistentDbSvc)
	contactCmd.AddCommand(contactController.Read())
	contactCmd.AddCommand(contactController.Create())
	contactCmd.AddCommand(contactController.Update())
	contactCmd.AddCommand(contactController.Delete())
	rootCmd.AddCommand(contactCmd)
}

func (router *Router) systemRoutes() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "PrintVersion",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Clean DDD TAGHS PoC Contacts v0.0.1")
		},
	}

	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "ServeHttpServer",
		Run: func(cmd *cobra.Command, args []string) {
			presentation.HttpServerInit(router.persistentDbSvc)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)
}

func (router *Router) RegisterRoutes() {
	router.contactRoutes()
	router.systemRoutes()
}
