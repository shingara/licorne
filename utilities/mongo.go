package utilities

import (
	"gopkg.in/mgo.v2"
	// "log"
	// "os"
)

type (
	Service struct {
		MongoSession *mgo.Session
		Database string
	}
)

var (
	mgoSession *mgo.Session
	config *Config
)

// Get a MongoDB Session and memoize it if already save
func getSession() *mgo.Session {
	// mgo.SetDebug(true)
	// var aLogger *log.Logger
	// aLogger = log.New(os.Stderr, "", log.LstdFlags)
	// mgo.SetLogger(aLogger)

	if config == nil {
		config = GetConfig()
	}

	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(config.Host + ":" + config.Port)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func DropDatabase() error{
	session := getSession()
	defer session.Close()
	return session.DB(config.Database).DropDatabase()
}

func WithCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(config.Database).C(collection)
	return s(c)
}
