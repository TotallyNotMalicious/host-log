package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	webhook "github.com/Harvey1717/go-discord-hooks"
)

func main() {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	send(response)
}

func send(response []byte) {

	host := string(response) // convert response to a string, because otherwise it goes ballistic

	embed := webhook.NewEmbed() // change everything to your liking, I have labeled all the fields
	embed.Title = "Title"
	embed.SetColour("42f5b9")
	embed.AddField("Data", ""+host+"\n", false)
	embed.SetFooter("Footer", "Footer Image URL")

	sent := embed.Send(
		"Webhook URL",
		"",
		"Webhook Name",
		"Webhook Pfp Image URL")

	if sent {
		fmt.Print("Data Sent To Webhook")
	}
}
