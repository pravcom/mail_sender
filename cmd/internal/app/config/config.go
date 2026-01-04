package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8088"`
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"localhost"`
	}

	App struct {
		TimeLocation string `yaml:"time_loc" env:"APP_TIME_LOC" env-default:"Asia/Yekaterinburg"`
		CronExpr     string `yaml:"expr" env:"APP_EXPR" env-default:"00 10 * * 1-5"`
	}

	Mail struct {
		Subject  string   `yaml:"subject" env:"MAIL_SUBJECT" env-defult:"Test header"`
		From     string   `yaml:"from" env:"MAIL_FROM" env-default:"ahtyamovden@mail.ru"`
		To       []string `yaml:"to" env:"MAIL_TO" env-defualt:"pravcom2@gmail.com"`
		Password string   `yaml:"password" env:"MAIL_PASSWORD" env-default:"password"`
		Team     []string `yaml:"team" env:"MAIL_TEAM" env-default:"Ахтямов Денис, Стерлядов Дмитрий"`
	}
}

func New() (Config, error) {
	var cfg Config
	var err error

	//Перезаписываем переменные окружения из файла .en
	err = godotenv.Load()
	if err != nil {
		log.Printf("Error read ENV: %v", err)
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		err = fmt.Errorf("ReadConfig error: %v", err)
		return cfg, err
	}

	log.Printf("%v", cfg)

	return cfg, nil
}
