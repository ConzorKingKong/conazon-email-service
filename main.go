package main

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/conzorkingkong/conazon-email-service/types"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQURL, RabbitMQURLExists = "", false
var EmailPassword, EmailPasswordExists = os.LookupEnv("EMAILPASSWORD")

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func sendEmail(to string, subject string, body string) {
	from := "connor@connorpeshek.me"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, EmailPassword, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Printf("message sent to %s", to)
}

func main() {

	godotenv.Load()

	if !RabbitMQURLExists {
		RabbitMQURL = "amqp://guest:guest@rabbitmq"
	}

	conn, err := amqp.Dial(RabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var email types.Email
			err := json.Unmarshal(d.Body, &email)
			if err != nil {
				log.Printf("Error unmarshalling email: %v", err)
			}

			// get checkout ID from message body
			sendEmail(email.User.Email, "Conazon Purchase: "+strconv.Itoa(email.Checkout.Id), "Thank you for your purchase! \n Total price: "+email.Checkout.TotalPrice+"\n Your tracking number is: "+email.Checkout.TrackingNumber)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
