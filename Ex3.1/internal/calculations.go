package calculations

import (
	log "github.com/sirupsen/logrus"
)

func Calculate(n int64, flag bool) int64 {
	if flag {
		log.Info("Start calculations...")
		log.Info("Calculate ", n, "!")
	}
	var i int64
	var result int64

	i = 1
	result = 1

	for i < n {
		i++
		result *= i
	}
	if flag {
		log.Info("Calculations complete!")
	}
	return result
}
