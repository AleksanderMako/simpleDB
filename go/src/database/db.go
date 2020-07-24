package database

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// SimpleDB database struct
type SimpleDB struct {
	index map[string]int64
}

// NewSimpleDB is the constructor
func NewSimpleDB() *SimpleDB {
	index := make(map[string]int64)
	return &SimpleDB{
		index: index,
	}
}

func (db *SimpleDB) Write(key, val string) error {

	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	if stats, err := file.Stat(); err != nil {
		fmt.Println(err)
	} else {
		db.index[key] = stats.Size()
	}

	_, err = file.Write(append([]byte(key+":"+val), '\n'))
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (db *SimpleDB) Get(key string) (string, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := ""
	for scanner.Scan() {
		data := scanner.Text()
		colonIndex := strings.Index(data, ":")
		currentKey := data[0:colonIndex]
		if currentKey == key {
			result = data[colonIndex+1:]
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return result, nil
}

func (db *SimpleDB) FastGet(key string) (string, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		return "", err
	}
	result := ""
	defer file.Close()
	if row, ok := db.index[key]; ok {
		file.Seek(row, 0)
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		data := scanner.Text()

		colonIndex := strings.Index(data, ":")
		result = data[colonIndex+1:]
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return result, nil
	}

	return "", errors.New("Key not found ")
}
