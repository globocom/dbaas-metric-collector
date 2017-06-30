package cron

import (
	"fmt"
	"time"

	"github.com/otherpirate/dbaas-metric-collector/collector"
	"github.com/otherpirate/dbaas-metric-collector/settings"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

func DailyLoading() {
	ticker := updateTicker(settings.LOADING_HOUR, settings.LOADING_MINUTE, settings.LOADING_SECOND, INTERVAL_PERIOD)
	for {
		<-ticker.C
		fmt.Println("Loading databases...")
		collector.GetDatabases()
		fmt.Println("  Done")
		ticker = updateTicker(settings.LOADING_HOUR, settings.LOADING_MINUTE, settings.LOADING_SECOND, INTERVAL_PERIOD)
	}
}

func updateTicker(hour int, minute int, second int, period time.Duration) *time.Ticker {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "- next request")
	diff := nextTick.Sub(time.Now())
	return time.NewTicker(diff)
}
