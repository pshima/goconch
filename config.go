package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Endpoint struct {
	CheckType string `json:"type"`
	Address   string `json:"endpoint"`
	Timeout   int    `json:"timeout"`
}

type Endpoints struct {
	Checklist []Endpoint
}

func (e *Endpoints) FromJSONFile(args []string) error {
	if len(args) != 2 {
		usage()
		return nil
	}

	path := args[1]
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var data = &e.Checklist
	return json.Unmarshal(file, data)
}

func usage() {
	fmt.Println("Usage: %s <configfile.json>", os.Args[0])
}
