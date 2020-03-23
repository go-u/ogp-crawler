package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Options struct {
	RateLimit struct {
		Twitter int
		Ogp     int
	}
}

func getOptions() Options {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	// UnmarshalしてCにマッピング
	var options Options
	err = viper.Unmarshal(&options)
	if err != nil {
		log.Fatal("config file Unmarshal error: ", err)
	}

	return options
}

func getProjectId() string {
	PROJECT_ID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if PROJECT_ID == "" {
		log.Fatalln("Failed to Get PROJECT_ID 'GOOGLE_CLOUD_PROJECT'\n If this is local test, Set 'appname-local' as GOOGLE_CLOUD_PROJECT")
	}
	log.Println("PROJECT_ID: ", PROJECT_ID)
	return PROJECT_ID
}

func getHostName() string {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		log.Println(err)
		log.Fatalln("Failed to Get hostname")
	}
	log.Println("Host: ", hostname)
	return hostname
}
