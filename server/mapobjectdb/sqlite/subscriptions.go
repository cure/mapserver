package sqlite

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"
)

func (db *Sqlite3Accessor) GetSubscriptionsByMapblockPos(*coords.MapBlockCoords) ([]mapobjectdb.Subscription, error) {
	return nil, nil
}

func (db *Sqlite3Accessor) GetSubscriptionsByEndpoint(endpoint string) ([]mapobjectdb.Subscription, error) {
	return nil, nil
}

func (db *Sqlite3Accessor) AddSubscription(sub *mapobjectdb.Subscription) error {
	return nil
}

func (db *Sqlite3Accessor) RemoveSubscription(sub *mapobjectdb.Subscription) error {
	return nil
}
