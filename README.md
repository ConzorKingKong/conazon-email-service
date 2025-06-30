# Conazon Email Service

This is the email microservice for the ecommerce project. It listens for a message from RabbitMQ to send an email

## Quickstart

To test locally, setup a `.env` file in the root directory with the following variables:

```
EMAILPASSWORD - App password (not regular password) to gmail account — https://support.google.com/accounts/answer/185833?visit_id=638613322705524102-924909150&p=InvalidSecondFactor&rd=1 REQUIRED
RABBITMQURL - Url that RabbitMQ instance is running. Defaults to a rabbitmq container inside of docker-compose
```