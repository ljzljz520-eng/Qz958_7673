package store

import (
	"bakery34/model"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"sync"
)

var buckets = [][]byte{[]byte("records"), []byte("users"), []byte("events"), []byte("audits")}

type DB struct {
	raw *bbolt.DB
	mu  sync.RWMutex
}

func Open(path string) (*DB, error) {
	d, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	x := &DB{raw: d}
	e = d.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, z := tx.CreateBucketIfNotExists(b); z != nil {
				return z
			}
		}
		return nil
	})
	if e != nil {
		d.Close()
		return nil, e
	}
	return x, nil
}
func (d *DB) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.raw == nil {
		return nil
	}
	e := d.raw.Close()
	d.raw = nil
	return e
}
func (d *DB) Put(bucket, key string, v any) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.raw == nil {
		return fmt.Errorf("database closed")
	}
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return d.raw.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (d *DB) Get(bucket, key string, v any) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.raw == nil {
		return fmt.Errorf("database closed")
	}
	return d.raw.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		x := b.Get([]byte(key))
		if x == nil {
			return fmt.Errorf("not found")
		}
		return json.Unmarshal(x, v)
	})
}
func (d *DB) Delete(bucket, key string) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Delete([]byte(key)) })
}
func (d *DB) List(bucket string) ([][]byte, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var out [][]byte
	e := d.raw.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error { out = append(out, append([]byte(nil), v...)); return nil })
	})
	return out, e
}
func EncodeRecord(r model.Record) ([]byte, error) { return json.Marshal(r) }
