package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	_ "test-task-filikr/docs"
	"test-task-filikr/internal/app"
	"test-task-filikr/internal/config"
	"test-task-filikr/internal/database"
	"test-task-filikr/internal/logger"

	_ "github.com/lib/pq"
)

func main() {
	// config, err := config.LoadConfigYAML("./configs/config.yaml")
	config, err := config.LoadConfigENV()
	if err != nil {
		fmt.Print("failed load config: %w", err)
		return
	}
	fmt.Println("Load config")
	fmt.Println("Port app: ", config.App.PortApp)
	fmt.Println("Env levl: ", config.Env)

	log := logger.SetupLogger(config.Env)
	log.Info("Setup Logger")

	db, err := database.SetupDB(config.DB, log)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("Setup Database")

	go app.SetupApp(db, log, config.App.PortApp)
	log.Info("Start App")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Info("Stop App")

	db.Close()

	log.Info("Stop DataBase")
}
