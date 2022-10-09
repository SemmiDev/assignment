package db

import (
	"os"
)

type JsonDB struct{}

func NewJsonDB() *JsonDB {
	return &JsonDB{}
}

func (db *JsonDB) Load(dbName DBName) ([]byte, error) {
	jsonData, err := os.ReadFile("data/" + dbName + ".json")
	return jsonData, err
}

func (db *JsonDB) Save(dbName DBName, data Data) error {
	err := os.WriteFile("data/"+dbName+".json", data, 0644)
	return err
}

func (db *JsonDB) Reset(fileDb string, defVal []byte) error {
	err := os.WriteFile("data/"+fileDb+".json", defVal, 0644)
	if err != nil {
		return err
	}
	return nil
}
