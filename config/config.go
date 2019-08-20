package api

import (
	"fmt"

	"github.com/narup/gconfig"
)

func ReadConfig() (*gconfig.GConfig, error) {
	//load configuration
	config , err := gconfig.Load()
	if err != nil {
		fmt.Println("Error loading configuration files.")
		return nil, err
	}
	fmt.Println("Success loading configuration files.")
	return config, nil
}

func SetupDatabase() {
	//setup database
	// hostURL := util.Config("db.main.url")
	// cfg := gmgo.DbConfig{HostURL: hostURL, DBName: util.Config("db.main.name")}
	// _, err := data.SetupMongoDb(cfg)
	// if err != nil {
	// 	notification.PostOnMonitoringChannel(fmt.Sprintf("[ERROR][LOAD]:: %s", err))
	// 	log.Panicf("Error connecting to the database %s", err.Error())
	// 	return
	// }
}
