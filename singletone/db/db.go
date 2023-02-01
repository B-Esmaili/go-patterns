package db

import "sync"

type Db struct{}

var fakeDb *Db
var lock = &sync.Mutex{}

func GetDb() *Db {
	if fakeDb == nil {
		lock.Lock()
		defer lock.Unlock()
		if fakeDb == nil {
			fakeDb = &Db{}
		}
	}
	return fakeDb
}
