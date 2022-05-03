package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	DbPGURL   string
	JaegerURL string
	SentryURL string
	Kafka     string
	AppId     string
	AppKey    string
	User      string
	HomePath  string
}

var (
	Host = flag.String("app-host", "127.0.0.1", "сервер приложения")
	Port = flag.Int("app-port", 8080, "порт приложения")
	Conf = Config{}
)

func ReadConfig() {

	flag.Parse()

	Conf.DbPGURL = os.Getenv("DBPGURL")
	_, err := url.ParseRequestURI(Conf.DbPGURL)
	if err != nil {
		Conf.DbPGURL = ""
		os.Unsetenv("DBPGURL")
		fmt.Println(err)
	}

	Conf.JaegerURL = os.Getenv("JAEGERURL")
	_, err = url.ParseRequestURI(Conf.JaegerURL)
	if err != nil {
		Conf.JaegerURL = ""
		os.Unsetenv("JAEGERURL")
		fmt.Println(err)
	}

	Conf.SentryURL = os.Getenv("SENTRYURL")
	_, err = url.ParseRequestURI(Conf.SentryURL)
	if err != nil {
		Conf.SentryURL = ""
		os.Unsetenv("SENTRYURL")
		fmt.Println(err)
	}

	Conf.Kafka = os.Getenv("KAFKA")
	_, err = url.ParseRequestURI(Conf.Kafka)
	if err != nil {
		Conf.Kafka = ""
		os.Unsetenv("KAFKA")
		fmt.Println(err)
	}

	Conf.AppId = os.Getenv("APPID")
	Conf.AppKey = os.Getenv("APPKEY")

	Conf.User = os.Getenv("LOGNAME")
	Conf.HomePath = os.Getenv("HOME")

}
