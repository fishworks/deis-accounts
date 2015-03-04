package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fishworks/dis/cmd"
)

const (
	// Usage is the usage string of this client, which is displayed on os.Stdout when requested.
	Usage string = `
usage: deis accounts:<command> [args]

Basic Commands:
        help    display help information and exit
        list    list all accounts
        add     add an account
        remove  remove an account
        set     switch the account
        version display version information about this plugin
`
)

var (
	commands = []*cmd.Command{
		cmdHelp,
		cmdList,
		cmdAdd,
		cmdRemove,
		cmdSet,
		cmdVersion,
	}
)

func usage() {
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(Usage))
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	for _, cmd := range commands {
		if cmd.Name == args[0] {
			cmd.Flag.Parse(args[1:])
			cmd.Run(cmd, cmd.Flag.Args())
			return
		}
	}

	fmt.Fprintf(os.Stderr, "command not found\n")
	usage()
	os.Exit(1)
}
