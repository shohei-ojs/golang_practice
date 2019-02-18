package main


import (
	"os"

	"etc/vue-golang-payment-app/backent-api/infrastructure"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}