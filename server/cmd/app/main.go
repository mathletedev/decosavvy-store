package main

import (
	"flag"
	"log"

	"github.com/mathletedev/decosavvy/internal/auth"
	"github.com/mathletedev/decosavvy/internal/db"
	"github.com/mathletedev/decosavvy/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	prod := flag.Bool("dev", false, "run in development mode")

	flag.Parse()

	*prod = !*prod

	auth.NewAuth(prod)

	s := server.NewServer(prod)
	d := db.NewDatabase()
	defer d.Close()

	log.Println("Server started! 🚀")
	log.Fatal(s.ListenAndServe())
}
