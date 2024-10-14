package main

import (
	"log"
	"tty28/huiwan28"
)

func main() {
	if err := huiwan28.Run(); err != nil {
		log.Printf("huiwan28.Run() Failure :: %s \n", err.Error())
	}
}
