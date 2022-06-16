package main

import (
	"context"
	"github.com/MahmoudMekki/ChatSystem/database"
	"github.com/MahmoudMekki/ChatSystem/migration"
	"github.com/MahmoudMekki/ChatSystem/router"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := database.CreateDBConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	migration.RunMigration()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	routerInterface := router.NewRouter(engine)
	engine = routerInterface.SetRouter()
	srv := &http.Server{Addr: ":8080", Handler: engine}
	go func() {
		log.Err(srv.ListenAndServe())
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	<-sigChan
	log.Info().Msg("Received a terminate signal, Gracefully shutdown the server")
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
