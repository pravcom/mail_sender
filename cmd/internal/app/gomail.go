package app

import (
	"fmt"
	"log"
	"mail_sender/cmd/internal/app/config"
	"math/rand"
	"time"

	"gopkg.in/mail.v2"
)

func sendMailGoMail(cfg config.Config) error {

	m := mail.NewMessage()

	log.Printf("Начало работы джоба: %v \n", time.Now().Format("2006-01-02 15:04:05"))

	// Set email headers
	m.SetHeader("From", cfg.Mail.From)
	m.SetHeader("To", cfg.Mail.To...)
	m.SetHeader("Subject", cfg.Mail.Subject)

	// Set email body as HTML
	bodyText := getBody(cfg)
	m.SetBody("text/html", bodyText)

	// Set up the SMTP dialer
	// Replace with your SMTP server details (host, port, username, password)
	d := mail.NewDialer("smtp.mail.ru", 465, cfg.Mail.From, cfg.Mail.Password)

	// Optional: Use this for self-signed certificates or development servers
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")

	return nil
}

func getBody(cfg config.Config) string {

	// team := map[int]string{
	// 	0: "Ахтямов Денис",
	// 	1: "Барышев Денис",
	// 	2: "Николай Саранов",
	// 	3: "Елена Бударова",
	// 	4: "Дмитрий Стерлядов",
	// }

	team := getRandomSlice(cfg.Mail.Team)
	var body string
	for i, v := range team {
		if body == "" {
			body = fmt.Sprintf("%d. %s <br>", i+1, v)
		} else {
			text := fmt.Sprintf("%d. %s <br>", i+1, v)
			body = body + text
		}
	}

	return body

}

func getRandomSlice(team []string) []string {
	rand.Shuffle(len(team), func(i, j int) {
		team[i], team[j] = team[j], team[i]
	})

	return team
}
