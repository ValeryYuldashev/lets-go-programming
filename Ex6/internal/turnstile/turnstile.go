package turnstile

import (
	"ex6/internal/statistics"
	"fmt"
	"sync"
	"time"
)

func Work(id int, channel chan bool, mutex *sync.Mutex, stats *statistics.Stats) {

	for stats.Count < stats.Max {
		mutex.Lock()
		stats.Count++
		fmt.Println(id, ": ", stats.Count)
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
	channel <- true
}
