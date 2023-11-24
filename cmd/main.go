package main

import (
	"context"
	"fmt"
	"homework-6/config"
	"homework-6/internal/app/consumer"
	"homework-6/internal/app/producer"
	"homework-6/internal/infrastructure/kafka"
	"homework-6/internal/serv/core"
	"homework-6/internal/serv/db"
	"homework-6/internal/serv/repository/postgres"
	"homework-6/internal/serv/server"
	"log"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conStrDB, err := config.ReadConfDBConn("config/dbConf.json")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewDB(ctx, conStrDB)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	articleRepo := postgres.NewArticles(database)

	conStrServ, err := config.ReadServConf("config/serverConf.json")
	if err != nil {
		log.Fatal(err)
	}

	brokers, err := config.ReadConfKafkaConn("config/kafkaConf.json")

	if err != nil {
		log.Fatal(err)
	}

	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	answerService := producer.NewService(
		producer.NewKafkaSender(kafkaProducer, "URL"),
	)

	kafkaConsumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	readService := consumer.NewService(
		consumer.NewReceiver(kafkaConsumer),
	)

	go readService.ConsumerRun()

	implementation := core.FacadeServer{&server.Server{Repo: articleRepo}, answerService}
	http.Handle("/", core.CreateRouter(implementation))
	if err := http.ListenAndServe(conStrServ.Port, nil); err != nil {
		log.Fatal(err)
	}
}
