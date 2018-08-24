package model

import (
	"time"	
	"gopkg.in/mgo.v2/bson"
	
)

func DateTimeFilter(from time.Time, to time.Time) bson.M {
	return bson.M{"moment": bson.M{"$gte": from, "$lt": to}}
}