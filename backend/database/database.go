package database

import (
	"context"

	"github.com/globalsign/mgo"
)

var connections = make(map[string]*database)

type DataStore interface {
	FindAll(ctx context.Context, query interface{}, v interface{}) error
}
type database struct {
	dbname string
	sess   *mgo.Session
}

type datastore struct {
	db         *database
	collection string
}

func (db *database) withCollection(collection string, query func(*mgo.Collection) error) error {
	sess := db.sess.Copy()
	defer sess.Close()

	col := sess.DB(db.dbname).C(collection)
	return query(col)
}

func (ds *datastore) FindAll(ctx context.Context, query interface{}, v interface{}) error {
	return ds.db.withCollection(ds.collection, func(col *mgo.Collection) error {
		return col.Find(query).All(v)
	})
}

func NewDataStore(name, collection string) *datastore {
	if db, ok := connections[name]; ok {
		return &datastore{
			db:         db,
			collection: collection,
		}
	}
	return nil
}
