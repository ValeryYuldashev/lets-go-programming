package turnstile

import (
	"ex6/internal/statistics"
	"fmt"
	"sync"
	"time"
)

func Work(id int, channel chan bool, mutex *sync.Mutex, stats *statistics.Stats) {

	for {
		mutex.Lock()
		if stats.Count >= stats.Max {
			mutex.Unlock()
			break
		}
		stats.Count++
		fmt.Printf("%d: %d\n", id, stats.Count)
		mutex.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
	channel <- true
}
