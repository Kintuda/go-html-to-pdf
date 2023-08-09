package cmd

import (
	"context"
	"log"

	"github.com/kintuda/go-html-to-pdf/pkg/config"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewServerCmd(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Serve HTTP application",
		RunE:  StartServer,
	}

	return command
}

func StartServer(cmd *cobra.Command, arg []string) error {
	ctx := context.Background()
	app := fx.New(
		config.Module,
		http.Module,
		notification.Module,
		fx.Invoke(runHttpServer()),
	)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := app.Stop(ctx); err != nil {
			logrus.Errorf("error while closing app %v", err)
		}
	}()

	return nil
}

func runHttpServer() any {
	return func(
		// m *middleware.AuthorizationMiddleware, t *middleware.TransactionMiddleware
		lifecycle fx.Lifecycle, router *http.Router, cfg *config.AppConfig) {
		lifecycle.Append(fx.Hook{
			OnStart: func(context.Context) error {
				router.RegisterRoutes()
				return router.Start()
			},
			OnStop: func(context.Context) error {
				return router.Stop()
			},
		})
	}
}
