package model

import (
	"time"	
	"fmt"
	"gopkg.in/mgo.v2/bson"
	
)

func DateTimeFilter(dateFrom string, dateTo string) bson.M {
	var filter bson.M
	queryFromTime := "T00:00:00.000Z"
	queryToTime := "T23:59:59.000Z"

	if (len(dateFrom) == 0) && (len(dateTo) > 0) {
		dateTo += queryToTime
		to, err1 := time.Parse("2006-01-02T15:04:05.000Z", dateTo)
		if err1 != nil {
			fmt.Println(err1)
		}
		return bson.M{"moment": bson.M{"$lt": to}}
	}
	if (len(dateFrom) > 0) && (len(dateTo) == 0) {
		dateFrom += queryFromTime
		from, err1 := time.Parse("2006-01-02T15:04:05.000Z", dateFrom)
		if err1 != nil {
			fmt.Println(err1)
		}
		return bson.M{"moment": bson.M{"$gte": from}}
	}
	if ((len(dateFrom) > 0) && (len(dateTo) > 0)){
		dateFrom += queryFromTime
		dateTo += queryToTime
		from, err1 := time.Parse("2006-01-02T15:04:05.000Z", dateFrom)
		to, err2 := time.Parse("2006-01-02T15:04:05.000Z", dateTo)
		if err1 != nil {
			fmt.Println(err1)
		}
		if err2 != nil {
			fmt.Println(err2)
		}
		return bson.M{"moment": bson.M{"$gte": from, "$lt": to}}
	}

	return filter
}