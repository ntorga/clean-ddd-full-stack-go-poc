package cli

import (
	"fmt"
	"os"
	"path/filepath"

	cliMiddleware "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/cli/middleware"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: "Clean DDD TAGHS PoC Contacts",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func RunRootCmd() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func CliInit() {
	defer cliMiddleware.PanicHandler()

	router := NewRouter()
	router.RegisterRoutes()

	RunRootCmd()
}
