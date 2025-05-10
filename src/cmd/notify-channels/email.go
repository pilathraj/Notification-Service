package notifychannels

import "log"

func SendEmailNotification(message string) {
	log.Printf("Sending email notification: %s", message)
	// Add email sending logic here (e.g., using an SMTP client or email API)
}
