package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetAccountData reads json file and returns it
func GetAccountData(account string) (string, error) {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, fmt.Sprintf("../data/saves/%s.json", account))
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), err
}

// SetAccountData writes json to file
func SetAccountData(account, data string) error {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, fmt.Sprintf("../data/saves/%s.json", account))

	err := ioutil.WriteFile(path, []byte(data), 0644)

	return err
}
