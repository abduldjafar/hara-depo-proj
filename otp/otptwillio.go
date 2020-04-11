package otp

import (
	"github.com/sfreiberg/gotwilio"
	"hara-depo-proj/config"
)

func OtpTwillio(to string, randomString string) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	// Set account keys & information
	accountSid := baseConfig.Twillio.AccountSid
	authToken := baseConfig.Twillio.AuthToken

	// Create possible message bodies
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := baseConfig.Twillio.From

	message := "Terimakasih telah login. Berikut Token anda :" + randomString
	twilio.SendSMS(from, to, message, "", "")
}
