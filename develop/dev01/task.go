package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка получения времени: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Точное время: %v\n", time)
}
