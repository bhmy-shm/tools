package main

import (
	"log"
)

func main() {

	err := CpuCmd.Execute()
	if err != nil {
		log.Println(err)
	}

}
