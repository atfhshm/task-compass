package main

import (
	"fmt"

	"github.com/atfhshm/task-compass/config"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load()

	conf, err := config.NewConfig()

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
}
