package message

import (
	"fmt"
	"github.com/keighl/postmark"
)

func SendEmail(name, message, emailAddress string) {
	client := postmark.NewClient()
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
