package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// singleTickerDemo()
	// demoMultipleTickers()
	demoMultipleTickersWithStop()

	for {
	}
}

func singleTickerDemo() {
	item := "timber"
	reminderInterval := priceReminderInterval()[item]

	ticker := time.NewTicker(time.Duration(reminderInterval) * time.Second)

	receiveTicker(item, ticker)
}

func demoMultipleTickers() {
	for item, reminderInterval := range priceReminderInterval() {
		reminderTicker := time.NewTicker(time.Duration(reminderInterval) * time.Second)

		go receiveTicker(item, reminderTicker)
	}
}

func demoMultipleTickersWithStop() {
	for item, reminderInterval := range priceReminderInterval() {
		reminderTicker := time.NewTicker(time.Duration(reminderInterval) * time.Second)

		go stopTicker(item, reminderTicker)
		go receiveTicker(item, reminderTicker)
	}
}

func stopTicker(item string, ticker *time.Ticker) {
	time.Sleep(8 * time.Second)

	fmt.Println("Timer stop for", item)
	ticker.Stop()
}

func receiveTicker(item string, ticker *time.Ticker) {
	for range ticker.C {
		price := rand.Intn(100)

		fmt.Println(item + ", price " + strconv.Itoa(price))
	}
}

func priceReminderInterval() map[string]int {
	reminderInterval := map[string]int{
		"gold":   1,
		"timber": 3,
		"wood":   5,
		"iron":   2,
	}
	return reminderInterval
}
