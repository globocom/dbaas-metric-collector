package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type EngineCount struct {
	Name  string
	Count int
}

type EngineMoment struct {
	Moment  time.Time
	Engines []EngineCount
}

func EngineCounterAdd(connection Connection, moment time.Time, engines map[string]int) {
	enginesCount := []EngineCount{}
	for name, count := range engines {
		enginesCount = append(enginesCount, EngineCount{name, count})
	}

	collection := connection.Database.C("EngineCounterMoment")
	err := collection.Insert(&EngineMoment{moment, enginesCount})
	if err != nil {
		panic(err)
	}
}

func EngineCounterGet(connection Connection) []EngineMoment {
	counters := []EngineMoment{}
	collection := connection.Database.C("EngineCounterMoment")
	err := collection.Find(bson.M{}).All(&counters)
	if err != nil {
		panic(err)
	}

	diff := len(counters) - 15
	if diff > 0 {
		counters = counters[diff:]
	}

	return counters
}
