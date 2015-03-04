package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/fishworks/dis/cmd"
)

var cmdHelp = &cmd.Command{
	Name:  "help",
	Usage: "deis accounts:help [COMMAND]",
	Run:   runHelp,
}

func runHelp(cmd *cmd.Command, args []string) {
	if len(args) != 1 {
		usage()
		os.Exit(0)
	}

	binary, lookErr := exec.LookPath("man")
	if lookErr != nil {
		log.Println(lookErr)
		os.Exit(1)
	}

	// if `deis accounts:help accounts` was called, we want to retrieve deis-accounts(1)
	if args[0] == "accounts" {
		args[0] = "deis-accounts"
	} else {
		args[0] = fmt.Sprintf("deis-accounts-%s", args[0])
	}

	args = append([]string{"man"}, args[0])

	execErr := syscall.Exec(binary, args, os.Environ())
	if execErr != nil {
		log.Println(execErr)
		os.Exit(1)
	}
}
