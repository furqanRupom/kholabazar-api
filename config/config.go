package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var config *Config

type DBConfig struct {
	Host          string
	Name          string
	Port          int
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
	DB          *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables :", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required !")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is Required !")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is Required !")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port number should be a Number!")
	}
	JwtSecret := os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		fmt.Println("Jwt secret key is required")
	}

	host := os.Getenv("HOST")
	if host == "" {
		fmt.Println("Db host is required")
	}

	DBPort := os.Getenv("PORT")
	if DBPort == "" {
		fmt.Println("Db host should be a number")
	}
	dbPrt, err := strconv.ParseInt(DBPort, 10, 64)
	if err != nil {
		fmt.Println("Db port is required")
	}

	user := os.Getenv("USER")
	if user == "" {
		fmt.Println("Db user is required")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		fmt.Println("Db password is required")
	}
	DBName := os.Getenv("NAME")
	if DBName == "" {
		fmt.Println("Db name is required")
	}
	sslMode := os.Getenv("ENABLE_SSL_MODE")
	if sslMode == "" {
		fmt.Println("Enable ssl mode is required")
	}
	EnableSSLMode, err := strconv.ParseBool(sslMode)
	if err != nil {
		fmt.Println("Enable ssl mode should be boolean value")
	}

	dbConfig := &DBConfig{
		Host:          host,
		Port:          int(dbPrt),
		Name:          DBName,
		User:          user,
		Password:      password,
		EnableSSLMode: EnableSSLMode,
	}
	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JWTSecret:   JwtSecret,
		DB:          dbConfig,
	}
}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}
