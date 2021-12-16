package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/a9u/we-mock-api/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conf = &config.Conf{}

func Execute() {

}

func initConfig(configFile string) {
	fmt.Println("received config file", configFile)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msg("error while reading config file")
		panic(err)
	}

	if err = viper.Unmarshal(conf); err != nil {
		log.Error().Err(err).Msg("error while unmarshalling config file")
		panic(err)
	}

	log.Info().Msg("successfully read config file")
}

func init() {
	var configFile = flag.String("config", config.DefaultConfigFile, "config file name")
	flag.Parse()

	initConfig(*configFile)
	initDb()
}

func initDb() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Database.Uri()))
	if err != nil {
		log.Error().Err(err).Msg("error connecting to db")
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed ping")
		panic(err)
	}

	log.Info().Msg("successfully connected to db")
}
