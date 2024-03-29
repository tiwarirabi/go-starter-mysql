package api

import (
	"os"
	"fmt"
	"errors"
)

type Config struct {
	AppPort		string
}

// ReadConfig reads configuration values from environment variables.
func ReadConfig() (*Config, error) {

	c := new(Config)

	c.AppPort = os.Getenv("APP_PORT")

	if c.AppPort == "" {
		fmt.Printf("Error loading configuration. %s \n", os.Getenv("APP_PORT"))
		return nil, errors.New("Error loading configuration. Please check your environment variables.")
	}

	fmt.Println("Success loading configuration.")
	return c, nil
}

func SetupDatabase() {

}
