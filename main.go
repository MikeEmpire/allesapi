package main

import "os"

func main() {
	s := App{}
	s.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	s.Run(":8010")
}
