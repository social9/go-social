package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	TWConsumerKey    string `env:"TWConsumerKey"`
	TWConsumerSecret string `env:"TWConsumerSecret"`
}

var instance Config
var once sync.Once

func init() {
	once.Do(func() {
		instance = Config{}

		loadFromEnvFile(".env")
		if err := env.Parse(&instance); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("TWConsumerKey", instance.TWConsumerKey)
			os.Setenv("TWConsumerSecret", instance.TWConsumerSecret)
		}
	})
}

func loadFromEnvFile(file string) {
	envFile, err := ioutil.ReadFile(file)
	if err == nil {
		formatted := strings.Split(string(envFile), "\n")
		for _, val := range formatted {
			val = strings.Trim(val, "\n")
			if val != "" && !strings.HasPrefix(val, "#") {
				envValue := strings.Split(val, "=")
				if len(envValue) == 2 {
					os.Setenv(envValue[0], envValue[1])
				}
			}
		}
	}
}

// Env returns a config instance
func Env() Config {
	return instance
}
