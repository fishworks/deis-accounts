package main

import (
	"fmt"

	"github.com/fishworks/deis-accounts/version"
	"github.com/fishworks/dis/cmd"
)

var cmdVersion = &cmd.Command{
	Name:  "version",
	Usage: "deis accounts:version",
	Run: func(cmd *cmd.Command, args []string) {
		fmt.Printf("v%s\n", version.Version)
	},
}
