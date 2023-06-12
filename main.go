package main

import (
	"github.com/joho/godotenv"
	"github.com/yonisaka/xcrypto/cmd"
	"github.com/yonisaka/xcrypto/config"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.New()
	command := cmd.NewCommand(
		cmd.WithConfig(cfg),
	)

	app := cmd.NewCLI()
	app.Commands = command.Build()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
