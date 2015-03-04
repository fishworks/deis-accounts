package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fishworks/deis-accounts/deis"
	"github.com/fishworks/dis/cmd"
)

var cmdSet = &cmd.Command{
	Name:  "set",
	Usage: "deis accounts:set",
	Run:   runSet,
}

func runSet(cmd *cmd.Command, args []string) {
	var client deis.Client

	if len(args) != 1 {
		cmd.UsageExit()
	}

	clientFile, err := os.Open(clientFilepath)
	if err != nil {
		log.Fatalln(err)
	}

	jsonParser := json.NewDecoder(clientFile)
	if err = jsonParser.Decode(&client); err != nil {
		log.Fatalf("could not parse config file: %v\n", err)
	}

	var found bool
	for _, account := range client.Accounts {
		if args[0] == account.Username {
			found = true
			client.Username = account.Username
			client.Controller = account.Controller
			client.Token = account.Token
		}
	}
	if !found {
		fmt.Printf("could not find '%s'\n", args[0])
		os.Exit(1)
	}

	b, err := json.Marshal(client)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile(clientFilepath, b, 0644); err != nil {
		log.Fatalf("could not write to client file: %v\n", err)
	}
}
