package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// type Config struct {
// 	Host     string `mapstructure:"DB_HOST"`
// 	User     string `mapstructure:"DB_USER"`
// 	Password string `mapstructure:"DB_PASS"`
// 	DBName   string `mapstructure:"DB_NAME"`
// }

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }

func New() (*sql.DB, error) {
	// setup, err := LoadConfig("/home/wahyu/go/src/github.com/wsaefulloh/go-solid-principle")
	DB_HOST := "database"
	DB_USER := "devops"
	DB_PASS := "abcd1234"
	DB_NAME := "golang"
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME)

	// config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", setup.Host, setup.User, setup.Password, setup.DBName)

	db, err := connect(config)

	if err != nil {
		return nil, err
	}

	fmt.Println("Database ready")

	return db, nil
}

func connect(config string) (*sql.DB, error) {
	db, err := sql.Open("postgres", config)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
