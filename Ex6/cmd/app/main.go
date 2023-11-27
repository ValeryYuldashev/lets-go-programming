package main

import (
	"ex6/internal/statistics"
	"ex6/internal/turnstile"
	"fmt"
	"sync"
)

func main() {
	var metroStatistics statistics.Stats

	metroStatistics.Max = 30
	metroStatistics.NumberOfTurnstile = 15

	channel := make(chan bool)
	var mutex sync.Mutex
	for i := 0; i < metroStatistics.NumberOfTurnstile; i++ {
		go turnstile.Work(i, channel, &mutex, &metroStatistics)
	}

	for i := 0; i < metroStatistics.NumberOfTurnstile; i++ {
		<-channel
	}

	fmt.Printf("Количество человек: %d", metroStatistics.Count)
}
