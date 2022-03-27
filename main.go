package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/suppakorn-dev/lmwn-covid-19/handler"
	"github.com/suppakorn-dev/lmwn-covid-19/repository"
	"github.com/suppakorn-dev/lmwn-covid-19/service"
)

func main() {
	initViperConfig()

	URL := viper.GetString("resource.connection")
	PORT := viper.GetString("app.port")

	covidRepo := repository.NewCovidRepository(URL)
	covidSrv := service.NewCovidService(covidRepo)
	covidHandler := handler.NewCovidHandler(covidSrv)

	router := gin.Default()
	router.GET("/covid/summary", covidHandler.GetCovidPatientSummary)

	srv := &http.Server{
		Addr:         ":" + PORT,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("URL : ", URL)
	log.Println("PORT : ", PORT)
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}

func initViperConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Config file not found, You should set env variable.")
	}
}
