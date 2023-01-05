package thirdparty

import (
	"capstone-alta1/utils/helper"
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "EO-Bozz <eobozz01@gmail.com>"
const CONFIG_AUTH_EMAIL = "eobozz01@gmail.com"
const CONFIG_AUTH_PASSWORD = "vodofmmlitkoyieb"

func SendMailConfirmedOrder(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Confirmed Email from EO-Bozz")
	mailer.SetBody("text/html", "Hello World!, <p>Thanks for Order, see you later. :) </p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
		log.Fatal(err.Error())
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailWaitingPayment(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Waiting for Payment from EO-Bozz")
	mailer.SetBody("text/html", "Hello World!, <p>Waiting for Payment, please payout your transaction.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
		log.Fatal(err.Error())
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailWaitingConfirmation(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Payout Success, please wait your EO Confirmation")
	mailer.SetBody("text/html", "Hello World!, <p>Thanks for your paid, your EO is checking payment. Please wait for confirmation from EO.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
		log.Fatal(err.Error())
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailCompleteOrder(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Order Done")
	mailer.SetBody("text/html", "Hello World!, <p>Thank you for your cooperation, see you next time.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
		log.Fatal(err.Error())
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailPayoutSuccess(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Payout Done")
	mailer.SetBody("text/html", "Hello World!, <p>Thank you for using EO-Bozz as your business advice.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
		log.Fatal(err.Error())
	} else {
		helper.LogDebug("Success sent email. ")
	}
}
