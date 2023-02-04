// Package google implements GCE-based persistence for cached RPC calls.
//
// Use the NewDatastore() constructor to initialize a GCE-backed Datastore that
// can be used to construct a cached client.
package google

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/Tilo-K/riot/cachedclient"

	"cloud.google.com/go/datastore"
)

var (
	zeroTime time.Time

	// terminalTime is the last supported timestamp. Values are sorted ascending
	// in time until terminalTime. In principle, we could store values descending
	// in time to get the most recent observation, but this requires additional
	// index configuration.
	terminalTime = time.Date(2300, time.December, 0, 0, 0, 0, 0, time.UTC)
)

type googleDatastore struct {
	client    *datastore.Client
	namespace string
}

func (g *googleDatastore) Get(ctx context.Context, key string, dest interface{}, t time.Time) (time.Time, error) {
	filterKey := datastore.Key{
		Kind:      key,
		ID:        int64(terminalTime.Sub(t).Seconds()),
		Namespace: g.namespace,
	}
	query := datastore.NewQuery(key).Namespace(g.namespace).Filter("__key__ >=", &filterKey).Order("__key__").Limit(1)
	entities := reflect.New(reflect.SliceOf(reflect.TypeOf(dest)))
	keys, err := g.client.GetAll(ctx, query, entities.Interface())
	if err != nil {
		return zeroTime, err
	}
	if len(keys) == 0 {
		return zeroTime, fmt.Errorf("key %q is missing", key)
	}
	// dest is type *T so entities is type *([]*T).
	reflect.ValueOf(dest).Elem().Set(entities.Elem().Index(0).Elem())
	return time.Unix(terminalTime.Unix()-keys[0].ID, 0), nil
}

func (g *googleDatastore) Put(ctx context.Context, key string, val interface{}, t time.Time) error {
	k := datastore.Key{
		Kind:      key,
		ID:        int64(terminalTime.Sub(t).Seconds()),
		Namespace: g.namespace,
	}
	_, err := g.client.Put(ctx, &k, val)
	return err
}

func (g *googleDatastore) Purge(ctx context.Context, key string, keep int) error {
	query := datastore.NewQuery(key).Namespace(g.namespace).Order("__key__").KeysOnly()
	keys, err := g.client.GetAll(ctx, query, nil)
	if err != nil {
		return err
	}
	var toDelete []*datastore.Key
	for i := keep; i < len(keys); i++ {
		toDelete = append(toDelete, keys[i])
	}
	if len(toDelete) > 0 {
		return g.client.DeleteMulti(ctx, toDelete)
	}
	return nil
}

func NewDatastore(ctx context.Context, project, namespace string) (cachedclient.Datastore, error) {
	ds, err := datastore.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	return &googleDatastore{
		client:    ds,
		namespace: namespace,
	}, nil
}
