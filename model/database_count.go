package model

import (
	"time"
)

type DatabaseCount struct {
	Moment time.Time
	Count  int
}

func DatabaseCounterAdd(connection Connection, moment time.Time, count int) {
	collection := connection.Database.C("DatabaseCounter")
	err := collection.Insert(&DatabaseCount{moment, count})
	if err != nil {
		panic(err)
	}
}
