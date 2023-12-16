package main

import (
	"context"
	"log"
	"time"

	"github.com/siprtcio/email-microservice/examples/src"
)

func sendEmailStd(req *src.SendEmailRequest) (*src.ResponseMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.SendEmail(ctx, req)
	if err != nil {
		log.Printf("request failed: %v", err)
		return nil, err
	}

	return res, nil
}

func StdEmail() (*src.ResponseMessage, error) {
	emailRequest := src.SendEmailRequest{
		Recipients: &src.Recipients{
			To: []string{
				"surendra <surendratiwari3@gmail.com>",
			},
			Cc: []string{
				"surendra <surendra@siprtc.io>",
			},
			Bcc: []string{
				"admin@siprtc.io",
			},
		},
		Subject:     "Hi there. I hope you're good",
		ContentType: "text/html",
		Body:        "<h1>This is heading</h1><p>This is text</p>",
	}

	return sendEmailStd(&emailRequest)
}
