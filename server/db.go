package server

import (
	"fmt"
	"io/ioutil"
	"os"
)

type JsonDb struct {
	connectionString string
	file             *os.File
}

func (db *JsonDb) OpenFile() error {
	jsonFile, err := os.Open(db.connectionString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.file = jsonFile
	return nil
}

func (db *JsonDb) ReadFile() ([]byte, error) {
	body, err := ioutil.ReadAll(db.file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

func (db *JsonDb) WriteFile(newBody []byte) error {
	err := ioutil.WriteFile(db.connectionString, newBody, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db *JsonDb) CloseFile() {
	db.file.Close()
}
