package TinyJsonDB

import (
	"encoding/json"
	"io/ioutil"
	"os"
	// "runtime"
	"sync"
	"time"
)

var refresh time.Duration = 60

func New() *TinyJsonDB {
	var tjdb TinyJsonDB
	tjdb.Mutex = &sync.Mutex{}
	tjdb.Data = make(map[string]map[string]interface{})
	tjdb.Load()
	go func(tjdb *TinyJsonDB) {
		defer tjdb.Save()
		for true {
			// if time.Now().Unix()-t > refresh {
				tjdb.Save()
				time.Sleep(refresh*time.Second)
			// } else {
				// runtime.Gosched()
			// }
		}
	}(&tjdb)
	return &tjdb
}
func (tjdb *TinyJsonDB) CreateTable(t string) {
	tjdb.Mutex.Lock()
	tjdb.Data[t] = make(map[string]interface{})
	tjdb.Mutex.Unlock()
}
func (tjdb *TinyJsonDB) IsSetTable(t string) bool {
	_, ok := tjdb.Data[t]
	return ok
}
func (tjdb *TinyJsonDB) Load() error {
	fc, err := os.Open("TinyJsonDB.json")
	if err != nil {
		return err
	}
	tjdb.Mutex.Lock()
	err = json.NewDecoder(fc).Decode(&tjdb.Data)
	tjdb.Mutex.Unlock()
	fc.Close()
	if err != nil {
		return err
	}
	return nil
}
func (tjdb *TinyJsonDB) Save() error {
	tjdb.Mutex.Lock()
	b, err := json.Marshal(tjdb.Data)
	tjdb.Mutex.Unlock()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("TinyJsonDB.json", b, 0777)
	if err != nil {
		return err
	}
	return nil
}

type (
	TinyJsonDB struct {
		Data  map[string]map[string]interface{}
		Mutex *sync.Mutex
	}
)
