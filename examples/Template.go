package main

import (
	"context"
	"time"

	"github.com/siprtcio/email-microservice/examples/src"
)

func sendEmailWithTemplate(req *src.SendEmailWithTemplateRequest) (*src.ResponseMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.SendEmailWithTemplate(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func TemplateEmail() (*src.ResponseMessage, error) {
	// Template request
	emailRequest := src.SendEmailWithTemplateRequest{
		Recipients: &src.Recipients{
			To: []string{
				"Surendra Tiwari<surendra@siprtc.io>",
			},
		},
		TemplateName: "email-confirm",
		TemplateParams: map[string]string{
			"UserName":           "Aditya Agrawal",
			"ConfirmAccountLink": "https://google.com",
			"UnsubscribeLink":    "https://google.com/unsubscribe",
		},
		Subject: "This is from a template",
	}

	return sendEmailWithTemplate(&emailRequest)
}
