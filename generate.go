package main

import (
	"encoding/json"
	"github.com/samarec1812/auth-lab/backend"
	"io/ioutil"
	"log"
)

var Users = []backend.Users{
	{
		UserName: "admin",
		Password: "1234",
	},
	{
		UserName: "admin-main",
		Password: "0000",
	},
}
const file = "users.json"

func main() {
	data, err := json.MarshalIndent(Users, "", "     ")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ioutil.WriteFile(file, data, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

}