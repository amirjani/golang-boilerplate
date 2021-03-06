package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"go.uber.org/automaxprocs/maxprocs"
	"golang-boilerplate/Config"
	"golang-boilerplate/Router"
	"golang-boilerplate/Service"
)

func main() {
	logger, loggerError := Service.NewLogger("Polaris Storage Service")
	if loggerError != nil {
		fmt.Errorf("error at start %w", loggerError)
	}
	defer logger.Sync()

	if _, maxProcessorError := maxprocs.Set(); maxProcessorError != nil {
		fmt.Errorf("failed to set maxprocs: %w", maxProcessorError)
	}

	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig(".env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}
	fmt.Printf("%+v\n", config)

	// =====================================================
	// Open Database Connection
	database, _ := Service.DatabaseOpen(Service.DatabaseConfig{
		User:         config.DB.User,
		Password:     config.DB.Password,
		Host:         config.DB.Host,
		Name:         config.DB.Name,
		MaxIdleConns: config.DB.MaxIdleConns,
		MaxOpenConns: config.DB.MaxOpenConns,
		DisableTLS:   config.DB.DisableTLS,
	})
	logger.Infow("Project database, ", "database", database)

	defer func() {
		logger.Infow("shutdown", "status", "here", "host", config.DB.Host)
		database.Close()
	}()

	// App Starting
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20
	app.Static("/assets/", "./public")
	Router.Routes(app, logger, database)

	errorChannel := make(chan error)
	func() {
		logger.Infow("Project Running On PORT")
		errorChannel <- app.Run(config.Api.ApiHost)
	}()
}
