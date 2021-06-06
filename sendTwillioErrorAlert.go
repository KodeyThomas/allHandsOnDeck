package allHandsOnDeck

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendTwillioErrorAlert(err error) bool {

	// Set Up Twilio Stuff
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	numberTo := os.Getenv("NUMBER_TO")
	numberFrom := os.Getenv("NUMBER_FROM")
	twilioAPI := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Body Of Our Text Message
	textData := "Error: " + err.Error() + " - Sent with allHandsOnDeck <3"

	// Package Message Data so Twilio can do it's magic
	msg := url.Values{}
	msg.Set("To", numberTo)
	msg.Set("From", numberFrom)
	msg.Set("Body", textData)
	msgEncoded := *strings.NewReader(msg.Encode())

	// Set Up HTTP Request
	client := &http.Client{}
	req, _ := http.NewRequest("POST", twilioAPI, &msgEncoded)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send a POST Request and return if it was successful
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true // Success
	} else {
		return false // Failure
	}
}