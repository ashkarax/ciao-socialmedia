package gosmtp

import (
	"fmt"
	"net/smtp"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
)

type SmtpCredentials struct {
	SmtpConfig config.Smtp
}

var smtpCredentials SmtpCredentials

func SmtpConfigsForEmailOtp(smtpConfigs config.Smtp) {
	smtpCredentials.SmtpConfig = smtpConfigs
}

func SendVerificationEmailWithOtp(otp int, recieverEmail string, recieverName string) error {
	from := smtpCredentials.SmtpConfig.SmtpSender
	password := smtpCredentials.SmtpConfig.SmtpPassword
	to := []string{recieverEmail}
	smtpHost := smtpCredentials.SmtpConfig.SmtpHost
	smtpPort := smtpCredentials.SmtpConfig.SmtpPort

	subject := "Verify Your Email Address for Ciao"
	body := fmt.Sprintf("Hello,%s\n\nThank you for signing up for Ciao. To complete your registration and ensure the security of your account, please verify your email address by entering the One-Time Password (OTP) provided below:\n\nOTP: %d\n\nPlease use the OTP to verify your email address on our platform within the next 10 minutes. After this time, the OTP will expire, and you will need to request a new one.\n\nIf you did not request this verification, please disregard this email.\n\nIf you need any assistance or have questions, feel free to reach out to our support team at support@example.com.\n\nThank you for choosing Ciao.\n\nBest regards,\nThe Ciao Team", recieverName, otp)
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("-----", err)
		return err
	}
	return nil
}

func SendRestPasswordEmailOtp(otp int, recieverEmail string) error {
	from := smtpCredentials.SmtpConfig.SmtpSender
	password := smtpCredentials.SmtpConfig.SmtpPassword
	to := []string{recieverEmail}
	smtpHost := smtpCredentials.SmtpConfig.SmtpHost
	smtpPort := smtpCredentials.SmtpConfig.SmtpPort

	subject := "Reset Your Password"
	body := fmt.Sprintf("Dear %s,\n\nYou recently requested to reset your password for your Ciao account. To complete the process, please use the following One-Time Password (OTP):\n\nOTP: %d\n\nThis OTP is valid for 10 minutes. Please do not share this OTP with anyone for security reasons. If you did not request a password reset, please ignore this email.\n\nThank you,\nThe Ciao Team", recieverEmail, otp)

	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("-----", err)
		return err
	}
	return nil
}
