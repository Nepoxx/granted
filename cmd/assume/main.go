package main

import (
	"fmt"
	"os"

	"github.com/common-fate/granted/internal/build"
	"github.com/common-fate/granted/pkg/assume"
	"github.com/common-fate/updatecheck"
	"github.com/fatih/color"
)

func main() {
	updatecheck.Check(updatecheck.GrantedCLI, build.Version, !build.IsDev())
	defer updatecheck.Print()

	app := assume.GetCliApp()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(color.Error, "%s\n", err)
		os.Exit(1)
	}
}
