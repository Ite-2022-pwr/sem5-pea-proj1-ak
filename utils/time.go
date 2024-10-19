package utils

import (
	"fmt"
	"log"
	"time"
)

// PrintTimeElapsed wypisuje, ile czasu minęło od danego momentu w mikrosekundach
func PrintTimeElapsed(start time.Time, prompt string) float64 {
	elapsed := CalculateTimeElapsed(start)
	log.Printf("[+] %s zajęło %s\n", prompt, GreenColor(fmt.Sprintf("%.3fms", elapsed)))
	return elapsed
}

// CalculateTimeElapsed oblicza, ile czasu upłynęło od danej chwili i zwraca ten czas w mikrosekundach
func CalculateTimeElapsed(start time.Time) float64 {
	return float64(time.Since(start).Microseconds()) / 1000
}
