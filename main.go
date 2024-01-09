package main

import (
	"fmt"
	"github.com/atfhshm/task-compass/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
}
