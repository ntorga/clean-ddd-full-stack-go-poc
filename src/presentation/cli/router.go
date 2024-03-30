package cli

import (
	"fmt"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation"
	cliController "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/cli/controller"
	"github.com/spf13/cobra"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (router Router) contactRoutes() {
	var contactCmd = &cobra.Command{
		Use:   "contact",
		Short: "ContactManagement",
	}

	contactController := cliController.NewContactController()
	contactCmd.AddCommand(contactController.GetContacts())
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
			presentation.HttpServerInit()
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)
}

func (router *Router) RegisterRoutes() {
	router.contactRoutes()
	router.systemRoutes()
}
