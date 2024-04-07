package cliHelper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/liaison"
	"golang.org/x/term"
)

func ResponseWrapper(liaisonOutput liaison.LiaisonOutput) {
	exitCode := 0
	switch liaisonOutput.Status {
	case liaison.MultiStatus:
		exitCode = 1
	case liaison.UserError:
		exitCode = 1
	case liaison.InfraError:
		exitCode = 1
	}

	stdoutFileDescriptor := int(os.Stdout.Fd())
	isNonInteractive := !term.IsTerminal(stdoutFileDescriptor)
	if isNonInteractive {
		standardJsonBytes, err := json.Marshal(liaisonOutput)
		if err != nil {
			fmt.Println("ResponseEncodingError")
			os.Exit(1)
		}

		fmt.Println(string(standardJsonBytes))
		os.Exit(exitCode)
	}

	prettyJsonBytes, err := json.MarshalIndent(liaisonOutput, "", "  ")
	if err != nil {
		fmt.Println("PrettyResponseEncodingError")
		os.Exit(1)
	}

	fmt.Println(string(prettyJsonBytes))
	os.Exit(exitCode)
}
