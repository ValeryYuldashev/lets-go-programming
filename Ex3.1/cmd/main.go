package main

import (
	"flag"
	"os"
	"strconv"

	calc "ex3.1/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	logFlag := flag.Bool("log", false, "Use for logging")
	flag.Parse()

	n, err := strconv.Atoi(os.Args[len(os.Args)-1])

	if err != nil {
		log.Error(err)
		return
	}

	log.Info(calc.Calculate(int64(n), *logFlag))
}
