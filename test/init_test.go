//go:build integration
// +build integration

package tests

import (
	"homework-6/config"
	"homework-6/internal/app/producer"
	"homework-6/internal/infrastructure/kafka"
	"homework-6/test/postgres"
	"log"
)

var (
	db            *postgres.TDB
	answerService *producer.Service
)

func init() {
	conStrDB, err := config.ReadConfDBConn("../config/dbConf.json")
	if err != nil {
		log.Fatal(err)
	}
	db = postgres.NewFromEnv(conStrDB)

	brokers, err := config.ReadConfKafkaConn("../config/kafkaConf.json")
	if err != nil {
		log.Fatal(err)
	}

	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	answerService = producer.NewService(
		producer.NewKafkaSender(kafkaProducer, "URL"),
	)
}
