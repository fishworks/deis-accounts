package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/fishworks/deis-accounts/deis"
	"github.com/fishworks/dis/cmd"
)

var cmdList = &cmd.Command{
	Name:  "list",
	Usage: "deis accounts:list",
	Run:   runList,
}

func runList(cmd *cmd.Command, args []string) {
	var client deis.Client
	clientFile, err := os.Open(path.Join(os.Getenv("HOME"), ".deis", "client.json"))
	if err != nil {
		log.Fatalf("could not open client file: %v\n", err)
	}

	jsonParser := json.NewDecoder(clientFile)
	if err = jsonParser.Decode(&client); err != nil {
		log.Fatalf("could not parse config file: %v\n", err)
	}

	fmt.Println("=== Accounts")
	for _, account := range client.Accounts {
		fmt.Println(account.Username)
	}
}
