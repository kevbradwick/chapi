package main

import (
	"fmt"
	"github.com/kevbradwick/chapi"
)

func main() {

	c, err := chapi.CompanyProfile("03977902")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.Type)
}
