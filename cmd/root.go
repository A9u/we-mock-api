package cmd

import (
	"flag"
	"github.com/a9u/we-mock-api/config"
	"github.com/a9u/we-mock-api/internal/db"
	"github.com/a9u/we-mock-api/internal/server"
	"github.com/a9u/we-mock-api/pkg/wlog"
	"github.com/spf13/viper"
)

var conf = &config.Conf{}

func Execute() {
	svc := server.New(conf)

	if err := svc.ListenAndServe(); err != nil {
		wlog.Error("failed: server shutdown", err)
		panic(err)
	}

	return
}

func initConfig(configFile string) {
	wlog.Info("received config file")
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		wlog.Error("error while reading config file", err)
		panic(err)
	}

	if err = viper.Unmarshal(conf); err != nil {
		wlog.Error("error while unmarshalling config file", err)
		panic(err)
	}

	wlog.Info("successfully read config file")
}

func init() {
	var configFile = flag.String("config", config.DefaultConfigFile, "config file name")
	flag.Parse()

	initConfig(*configFile)
	initDb()
}

func initDb() {
	dbCon := db.InitDb(conf)
	wlog.Print(dbCon)
	wlog.Info("connected successfully")
}
