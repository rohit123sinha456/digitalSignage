package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// type EnvConfig struct {
// 	APPDB             string `mapstructure:"APPDB"`
// 	APPOSURL          string `mapstructure:"APPOSURL"`
// 	APPOSACCKEY       string `mapstructure:"APPOSACCKEY"`
// 	APPOSSECRETACCKEY string `mapstructure:"APPOSSECRETACCKEY"`
// 	APPRABBITURL      string `mapstructure:"APPRABBITURL"`
// 	APPRABBITADMIN    string `mapstructure:"APPRABBITADMIN"`
// 	APPRABBITADMINPWD string `mapstructure:"APPRABBITADMINPWD"`
// }

func GetEnvbyKey(key string) string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Cant get pwd")
		panic(err)
	}
	log.Printf("%s", pwd)
	err = godotenv.Load(filepath.Join(pwd, ".env"))
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
	return os.Getenv(key)
}
