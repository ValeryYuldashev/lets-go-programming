package main

import (
	"ex6/internal/statistics"
	"ex6/internal/turnstile"
	"sync"
)

func main() {
	var metroStatistics statistics.Stats

	metroStatistics.Max = 80
	metroStatistics.NumberOfTurnstile = 8

	channel := make(chan bool)
	var mutex sync.Mutex
	for i := 0; i < metroStatistics.NumberOfTurnstile; i++ {
		go turnstile.Work(i, channel, &mutex, &metroStatistics)
	}

	for i := 0; i < metroStatistics.NumberOfTurnstile; i++ {
		<-channel
	}
}
