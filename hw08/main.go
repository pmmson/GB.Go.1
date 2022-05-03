package main

import (
	"fmt"
	"os"

	"GB.Go.1/hw08/config"
)

func main() {
	// Setup configuration for check before
	setupConfiguration()

	// read configuration
	config.ReadConfig()

	// output configuartion
	fmt.Printf("Host: %s; port: %d\n", *config.Host, *config.Port)
	fmt.Printf("Connection string DataBase:  %s\n", config.Conf.DbPGURL)
	fmt.Printf("Connection string Jaeger:  %s\n", config.Conf.JaegerURL)
	fmt.Printf("Connection string Sentry:  %s\n", config.Conf.SentryURL)
	fmt.Printf("Connection string Kafka:  %s\n", config.Conf.Kafka)
	fmt.Printf("AppID:  %s\n", config.Conf.AppId)
	fmt.Printf("AppKey:  %s\n", config.Conf.AppKey)
	fmt.Printf("User:  %s\n", config.Conf.User)
	fmt.Printf("User Home Path:  %s\n", config.Conf.HomePath)
}

func setupConfiguration() {
	os.Setenv("DBPGURL", "postgres://testuser:testpass@storedb:5432/petstore?sslmode=disable")
	os.Setenv("JAEGERURL", "http://jaeger:16234")
	os.Setenv("SENTRYURL", "http//sentry:9000")
	os.Setenv("KAFKA", "kafka:9092")
	os.Setenv("APPID", "TRY")
	os.Setenv("APPKEY", "kejdjwkKWNDkwjJEFKJSEndaKAwjnda")
}
