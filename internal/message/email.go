package message

import (
	"fmt"
	"github.com/keighl/postmark"
)

func SendEmail(name, message, emailAddress string) {
	client := postmark.NewClient("e9884e47-892c-4ce6-8283-9749f2a701be", "10cf34a6-00d7-4896-848d-2cc88626b544")

	email := postmark.Email{
		From:       "hello@comsma.com",
		To:         "cmsmallegan@comsma.com",
		ReplyTo:    emailAddress,
		Subject:    fmt.Sprintf("New inquiry - %s", name),
		TextBody:   message,
		Tag:        "new-inquiry",
		TrackOpens: true,
	}
	_, err := client.SendEmail(email)
	if err != nil {
		{
			panic(err)
		}
	}
}
