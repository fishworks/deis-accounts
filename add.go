package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/fishworks/deis-accounts/deis"
	"github.com/fishworks/dis/cmd"
)

var cmdAdd = &cmd.Command{
	Name:  "add",
	Usage: "deis accounts:add",
	Run:   runAdd,
}

var (
	clientFilepath = path.Join(os.Getenv("HOME"), ".deis", "client.json")
)

func prompt(prompt string) string {
	var response string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func runAdd(cmd *cmd.Command, args []string) {
	var (
		client        deis.Client
		controller    string
		username      string
		password      string
		loginResponse struct {
			token string `json:"token"`
		}
	)
	clientFile, err := os.Open(clientFilepath)
	if err != nil {
		log.Fatalf("could not open client file: %v\n", err)
	}

	jsonParser := json.NewDecoder(clientFile)
	if err = jsonParser.Decode(&client); err != nil {
		log.Fatalf("could not parse config file: %v\n", err)
	}

	controller = prompt("login URL: ")
	username = prompt("username: ")
	password = prompt("password: ")

	payload := bytes.NewBufferString(fmt.Sprintf("{'username': '%s', 'password': '%s'}", username, password))

	resp, err := http.Post(controller+"/v1/auth/login/", "application/json", payload)
	if err != nil {
		log.Fatalf("could not log in to '%s': %v\n", controller, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(body, &loginResponse); err != nil {
		log.Fatalln(err)
	}

	client.Accounts = append(client.Accounts, &deis.Account{Controller: controller, Username: username, Token: loginResponse.token})

	b, err := json.Marshal(client)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile(clientFilepath, b, 0644); err != nil {
		log.Fatalf("could not write to client file: %v\n", err)
	}
}
