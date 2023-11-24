package config

import (
	"encoding/json"
	"io"
	"os"
)

type ConnStructServ struct {
	Port string `json:"port"`
}

type ConnStructDB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type ConnKafka struct {
	IP string `json:"broker"`
}

func ReadServConf(path string) (*ConnStructServ, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var connStr ConnStructServ
	err = json.Unmarshal(byteValue, &connStr)

	if err != nil {
		return nil, err
	}

	return &connStr, nil
}

func ReadConfDBConn(path string) (*ConnStructDB, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var connStr ConnStructDB
	err = json.Unmarshal(byteValue, &connStr)

	if err != nil {
		return nil, err
	}

	return &connStr, nil
}

func ReadConfKafkaConn(path string) ([]string, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var connStr []ConnKafka
	err = json.Unmarshal(byteValue, &connStr)

	if err != nil {
		return nil, err
	}

	n := len(connStr)
	brokersIp := make([]string, n)
	for i, ip := range connStr {
		brokersIp[i] = ip.IP
	}
	return brokersIp, nil
}
