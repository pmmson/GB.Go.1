package main

import (
	"fmt"
	"log"
	"os"

	"GB.Go.1/hw08/config"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	// Setup configuration for check before
	setupConfiguration()

	// read configuration
	config.ReadConfig()

	// output configuartion
	fmt.Printf("\n=== FLAG & ENV Configuration ===\n\n")
	fmt.Printf("Host: %s; port: %d\n", *config.Host, *config.Port)
	fmt.Printf("Connection string DataBase:  %s\n", config.Conf.DbPGURL)
	fmt.Printf("Connection string Jaeger:  %s\n", config.Conf.JaegerURL)
	fmt.Printf("Connection string Sentry:  %s\n", config.Conf.SentryURL)
	fmt.Printf("Connection string Kafka:  %s\n", config.Conf.Kafka)
	fmt.Printf("AppID:  %s\n", config.Conf.AppId)
	fmt.Printf("AppKey:  %s\n", config.Conf.AppKey)
	fmt.Printf("User:  %s\n", config.Conf.User)
	fmt.Printf("User Home Path:  %s\n", config.Conf.HomePath)

	log.Printf("=== JSON Configuration ===\n\n")
	// read configuration JSON
	cfg, err := config.ReadConfigJSON()
	if err != nil {
		panic(err)
	}
	// output configuartion
	log.Printf("Host: %s; port: %d\n", cfg.Host, cfg.Port)
	log.Printf("Connection string DataBase:  %s\n", cfg.DbPGURL)
	log.Printf("Connection string Jaeger:  %s\n", cfg.JaegerURL)
	log.Printf("Connection string Sentry:  %s\n", cfg.SentryURL)
	log.Printf("Connection string Kafka:  %s\n", cfg.Kafka)
	log.Printf("AppID:  %s\n", cfg.Appsets.Appid)
	log.Printf("AppKey:  %s\n", cfg.Appsets.Apptoken)
}

func setupConfiguration() {
	os.Setenv("DBPGURL", "postgres://testuser:testpass@storedb:5432/petstore?sslmode=disable")
	os.Setenv("JAEGERURL", "http://jaeger:16234")
	os.Setenv("SENTRYURL", "http//sentry:9000")
	os.Setenv("KAFKA", "kafka:9092")
	os.Setenv("APPID", "TRY")
	os.Setenv("APPKEY", "kejdjwkKWNDkwjJEFKJSEndaKAwjnda")
}
