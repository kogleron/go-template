package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		panic(err)
	}

	gotemplate, err := initGoTemplate()
	if err != nil {
		panic(err)
	}

	gotemplate.Run()
}
