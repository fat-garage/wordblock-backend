package cmd

import (
	"context"
	"github.com/fat-garage/wordblock-backend/api"
	"github.com/fat-garage/wordblock-backend/internal/service/block"
	"github.com/fat-garage/wordblock-backend/pkg/conf"
	"github.com/fat-garage/wordblock-backend/pkg/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var apiCmdName = "start-api"

func startApiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   apiCmdName,
		Short: "Start api server",
		Run: func(cmd *cobra.Command, _ []string) {
			// service init
			block.Init(database.Db)
			// gin init
			r := api.Init()
			srv := http.Server{
				Addr:    conf.Cfg.Gin.Listen,
				Handler: r,
			}
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("server listen err:%s", err)
				}
			}()
			logrus.Infof("Server started at %s", conf.Cfg.Gin.Listen)

			quit := make(chan os.Signal, 1)

			// kill (no param) default send syscall.SIGTERM
			// kill -2 is syscall.SIGINT
			// kill -9 is syscall.SIGKILL but can"t be caught, so don't need to add it
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			logrus.Info("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				logrus.Fatalf("Server Shutdown: %v", err)
			}

			logrus.Info("Server exiting")
		},
	}

	return cmd
}
