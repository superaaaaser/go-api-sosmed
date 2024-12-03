package cmd

import (
	"log"

	"github.com/superaaaaser/go-api-sosmed/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}