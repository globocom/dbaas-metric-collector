package model

import (
	"gopkg.in/mgo.v2"

	"github.com/globocom/dbaas-metric-collector/settings"
)

type Connection struct {
	Session  *mgo.Session
	Database *mgo.Database
}

func GetConnection() Connection {
	session, err := mgo.Dial(settings.MONGODB_ENDPOINT)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return Connection{session, session.DB("dbaas_metric_collector")}
}
