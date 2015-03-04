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

var cmdRemove = &cmd.Command{
	Name:  "remove",
	Usage: "deis accounts:remove",
	Run:   runRemove,
}

func runRemove(cmd *cmd.Command, args []string) {
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
	for i, account := range client.Accounts {
		if args[0] == account.Username {
			found = true
			client.Accounts = append(client.Accounts[:i], client.Accounts[i+1:]...)
		}
	}

	if !found {
		fmt.Printf("username '%s' not found\n", args[0])
	}

	b, err := json.Marshal(client)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile(clientFilepath, b, 0644); err != nil {
		log.Fatalf("could not write to client file: %v\n", err)
	}

	fmt.Println("Account removed.")
}
