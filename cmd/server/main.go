package server

import (
	"os"

	"github.com/mikeempire/allesapi/cmd/app"
)

func main() {
	s := app.App{}
	s.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	s.Run(":8010")
}
