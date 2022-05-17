package config

import (
	"encoding/json"
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

// https://mholt.github.io/json-to-go/
type AppConfig struct {
	Host      string      `json:"host"`
	Port      int         `json:"port"`
	DbPGURL   string      `json:"dbpgurl"`
	JaegerURL string      `json:"jaegerurl"`
	SentryURL string      `json:"sentryurl"`
	Kafka     string      `json:"kafka"`
	Appsets   Appsettings `json:"appsettings"`
}

type Appsettings struct {
	Appid    string `json:"appid"`
	Apptoken string `json:"apptoken"`
}

var (
	Host = flag.String("app-host", "127.0.0.1", "сервер приложения")
	Port = flag.Int("app-port", 8080, "порт приложения")
	Conf = Config{}
	appcfg  = AppConfig{}
)

func ReadConfig() {

	flag.Parse()

	Conf.DbPGURL = os.Getenv("DBPGURL")
	Conf.JaegerURL = os.Getenv("JAEGERURL")
	Conf.SentryURL = os.Getenv("SENTRYURL")
	Conf.Kafka = os.Getenv("KAFKA")
	Conf.AppId = os.Getenv("APPID")
	Conf.AppKey = os.Getenv("APPKEY")
	Conf.User = os.Getenv("LOGNAME")
	Conf.HomePath = os.Getenv("HOME")
	Conf.isValid()
}

func ReadConfigJSON() (AppConfig, error) {
	cfgFile, err := os.ReadFile("config/appconfig.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(cfgFile, &appcfg)
	if err != nil {
		fmt.Println(err)
	}
	return appcfg, appcfg.isValid()
}

func (conf *Config) isValid() {

	_, err := url.ParseRequestURI(conf.DbPGURL)
	if err != nil {
		Conf.DbPGURL = ""
		os.Unsetenv("DBPGURL")
		fmt.Println(err)
	}

	_, err = url.ParseRequestURI(conf.JaegerURL)
	if err != nil {
		Conf.JaegerURL = ""
		os.Unsetenv("JAEGERURL")
		fmt.Println(err)
	}

	_, err = url.ParseRequestURI(conf.SentryURL)
	if err != nil {
		Conf.SentryURL = ""
		os.Unsetenv("SENTRYURL")
		fmt.Println(err)
	}

	_, err = url.ParseRequestURI(conf.Kafka)
	if err != nil {
		Conf.Kafka = ""
		os.Unsetenv("KAFKA")
		fmt.Println(err)
	}
}

func (conf *AppConfig) isValid() error {

	_, err := url.ParseRequestURI(conf.DbPGURL)
	if err != nil {
		conf.DbPGURL = ""
		return fmt.Errorf("%w", err)
	}
	_, err = url.ParseRequestURI(conf.JaegerURL)
	if err != nil {
		conf.JaegerURL = ""
		return fmt.Errorf("%w", err)
	}
	_, err = url.ParseRequestURI(conf.SentryURL)
	if err != nil {
		conf.SentryURL = ""
		return fmt.Errorf("%w", err)
	}
	_, err = url.ParseRequestURI(conf.Kafka)
	if err != nil {
		conf.Kafka = ""
		return fmt.Errorf("%w", err)
	}
	return nil
}
