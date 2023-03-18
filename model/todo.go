package model

import (
	"encoding/json"
	"io/ioutil"
)

type Todo struct {
	Title   string
	Checked bool
}

func ReadAll(filename string) ([]*Todo, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []*Todo{}, err
	}
	var items []*Todo
	err = json.Unmarshal(b, &items)
	if err != nil {
		return []*Todo{}, err
	}

	return items, nil
}

func SaveAll(filename string, todos []*Todo) error {
	b, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
