package main

import (
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("https://yourmorpheus.com")
	client.SetUsernameAndPassword("username", "password")
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)

	// Get guidance settings
	req := &morpheus.Request{}
	guidanceSettingsResponse, err := client.GetGuidanceSettings(req)
	if err != nil {
		log.Fatal(err)
	}
	result := guidanceSettingsResponse.Result.(*morpheus.GetGuidanceSettingsResult)
	log.Println(result.GuidanceSettings)
}
